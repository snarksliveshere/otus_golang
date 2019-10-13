package main

import (
	"fmt"
	"os"
)

func main() {
	firstEnv := os.Getenv("first_env")
	fmt.Printf(firstEnv)
}
