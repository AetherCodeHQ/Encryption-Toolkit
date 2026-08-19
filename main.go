package main

import (
	"fmt"
	"os"
)

// encryption_toolkit - Encryption utilities
func encryption_toolkit(path string) {
	fmt.Println("========================================")
	fmt.Println("  Encryption-Toolkit")
	fmt.Println("  Encryption utilities")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	encryption_toolkit(path)
}
