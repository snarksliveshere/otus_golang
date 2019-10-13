package main

import (
	"fmt"
	"os"
)

func main() {
	firstEnv := os.Getenv("first_env")
	secondEnv := os.Getenv("second_env")
	fmt.Printf("Call prog with env var. First: %s, Second: %s", firstEnv, secondEnv)
}
