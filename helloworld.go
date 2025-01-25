

package main

import (
	"fmt"
	"os"
)

func main() {
	// Access the secret from the environment variable
	secret := os.Getenv("MY_SECRET")
	fmt.Printf("My secret is '%s'\n", secret)
}

