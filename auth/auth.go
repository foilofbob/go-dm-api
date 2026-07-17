package auth

import (
	"encoding/json"
	"fmt"
	"go_dm_api/api"
	"go_dm_api/config"
	"go_dm_api/domain"
	"net/http"
	"runtime"
	"strconv"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

var jwtSecretKey = []byte(config.Cfg().JWT.Secret)

var argonParams = &argon2id.Params{
	Memory:      128 * 1024,
	Iterations:  4,
	Parallelism: uint8(runtime.NumCPU()),
	SaltLength:  16,
	KeyLength:   32,
}

func VerifyUserPassword(passwordInput string, storedHash string) bool {
	match, err := argon2id.ComparePasswordAndHash(passwordInput, storedHash)
	if err != nil {
		return false
	}
	return match
}

func GenerateArgon2Hash(passwordInput string) string {
	hash, err := argon2id.CreateHash(passwordInput, argonParams)
	if err != nil {
		return ""
	}
	return hash
}

type CustomClaims struct {
	UserID   string `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID, username, role string) (string, error) {
	claims := CustomClaims{
		UserID:   userID,
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(20 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "go-dm-auth",
			Subject:   userID,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}

type AuthForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		api.StandardResponse(w, nil)
		return
	}
	api.EnableCORS(&w)

	var formData AuthForm
	json.NewDecoder(r.Body).Decode(&formData)

	user, err := domain.GetUserByUsername(formData.Username)
	if err != nil || user == nil {
		fmt.Println("user not found: " + formData.Username)
		http.Error(w, "Access Denied", http.StatusForbidden)
		return
	}

	valid := VerifyUserPassword(formData.Password, user.Password)
	if valid == false {
		fmt.Println("passwords did not match")
		http.Error(w, "Access Denied", http.StatusForbidden)
		return
	}

	tokenString, err := GenerateJWT(fmt.Sprintf("%d", user.ID), user.Username, "DM")
	if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	expires := time.Now().Add(30 * time.Minute)

	http.SetCookie(w, &http.Cookie{
		Name:     "dm_token",
		Value:    tokenString,
		Expires:  expires,
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		//SameSite: http.SameSiteStrictMode,
	})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"username": formData.Username, "expires": expires.Format(time.RFC3339)})
}

func HandleLogout(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		api.StandardResponse(w, nil)
		return
	}
	api.EnableCORS(&w)

	expiredCookie := &http.Cookie{
		Name:     "dm_token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, expiredCookie)

	w.WriteHeader(http.StatusOK)
}

// UserAccessMiddleware To check that user is authenticated, but not specific permissions
func UserAccessMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api.EnableCORS(&w)

		if r.Method == "OPTIONS" {
			next.ServeHTTP(w, r)
			return
		}

		cookie, err := r.Cookie("dm_token")
		if err != nil {
			fmt.Println("(UserAccess) Could not parse cookie: ", err.Error())
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token, err := jwt.ParseWithClaims(cookie.Value, &CustomClaims{}, func(t *jwt.Token) (any, error) {
			return jwtSecretKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		_, ok := token.Claims.(*CustomClaims)

		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func CampaignAccessMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api.EnableCORS(&w)

		if r.Method == "OPTIONS" {
			next.ServeHTTP(w, r)
			return
		}

		vars := mux.Vars(r)
		campaignIDStr := vars["campaignId"]
		campaignID, err := strconv.Atoi(campaignIDStr)

		cookie, err := r.Cookie("dm_token")
		if err != nil {
			fmt.Println("(CampaignAccess) Could not parse cookie: ", err.Error())
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token, err := jwt.ParseWithClaims(cookie.Value, &CustomClaims{}, func(t *jwt.Token) (any, error) {
			return jwtSecretKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(*CustomClaims)

		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		rawUserID := claims.UserID
		userID, err := strconv.Atoi(rawUserID)

		if err != nil {
			fmt.Printf("Could not parse userID (%s): %s\n", rawUserID, err.Error())
			http.Error(w, "Forbidden: You do not have access to this campaign", http.StatusForbidden)
			return
		}

		hasAccess := err == nil && domain.HasPermission(userID, campaignID)
		if !hasAccess {
			http.Error(w, "Forbidden: You do not have access to this campaign", http.StatusForbidden)
			return
		}

		if claims.ExpiresAt.Before(time.Now()) {
			fmt.Printf("JWT expired for userID (%s)", rawUserID)
			http.Error(w, "Forbidden: Session expired", http.StatusForbidden)
			return
		}

		// Renew cookie
		// TODO: Implement refresh token pattern
		user, err := domain.GetUser(userID)
		tokenString, err := GenerateJWT(rawUserID, user.Username, "DM")
		if err != nil {
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "dm_token",
			Value:    tokenString,
			Expires:  time.Now().Add(30 * time.Minute),
			HttpOnly: true,
			Secure:   true,
			Path:     "/",
			//SameSite: http.SameSiteStrictMode,
		})

		next.ServeHTTP(w, r)
	})
}
