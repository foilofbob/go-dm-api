package main

import (
	"flag"
	"fmt"
	"go_dm_api/auth"
)

func main() {
	// 1. Define your flags (Name, Default Value, Description)
	password := flag.String("password", "", "Password to hash")

	// 2. You must call flag.Parse() to actually read the CLI inputs
	flag.Parse()

	hash := auth.GenerateArgon2Hash(*password)

	// 3. De-reference the pointers to get the values
	fmt.Print(hash)
}
