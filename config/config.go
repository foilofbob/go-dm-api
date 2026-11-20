package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
	JWT      JWT      `yaml:"jwt"`
}

type Server struct {
	Port   string `yaml:"port"`
	Host   string `yaml:"host"`
	Origin string `yaml:"origin"`
}

type Database struct {
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	Username string `yaml:"user"`
	Password string `yaml:"pass"`
	DB       string `yaml:"db"`
	Driver   string `yaml:"driver"`
}

type JWT struct {
	Secret string `yaml:"secret"`
}

var currentConfig = Config{
	Server: Server{
		Port:   "123",
		Host:   "localhost",
		Origin: "http://localhost:8280",
	},
	Database: Database{
		Username: "db",
		Password: "pass",
		DB:       "dm_campaign_manager",
		Driver:   "mysql",
	},
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func InitConfig() {

	env := getEnv("ENV", "local")

	if env == "local" {
		// For local dev read from config
		f, err := os.Open("config/config.yml")
		if err != nil {
			log.Println(fmt.Sprintf("Error opening config: %s", err.Error()))
		}
		defer f.Close()

		decoder := yaml.NewDecoder(f)
		err = decoder.Decode(&currentConfig)
		if err != nil {
			log.Println(fmt.Sprintf("Error decoding config: %s", err.Error()))
		}
	} else {
		// For deployment read from environment variables
		currentConfig.Server.Host = getEnv("HOST", "localhost")
		currentConfig.Server.Port = getEnv("PORT", "8090")
		currentConfig.Server.Origin = getEnv("ORIGIN", "http://localhost:8280")
		currentConfig.Database.Host = getEnv("DB_HOST", "localhost")
		currentConfig.Database.Port = getEnv("DB_PORT", "3306")
		currentConfig.Database.DB = getEnv("DB_NAME", "dm_campaign_manager")
		currentConfig.Database.Username = getEnv("DB_USERNAME", "db")
		currentConfig.Database.Password = getEnv("DB_PASSWORD", "pass")
	}
}

func Cfg() Config {
	return currentConfig
}

func Origin() string {
	return currentConfig.Server.Origin
}

func DBDriver() string {
	return currentConfig.Database.Driver
}

func DBConnectString() string {
	return currentConfig.Database.Username + ":" + currentConfig.Database.Password + "@/" + currentConfig.Database.DB
}
