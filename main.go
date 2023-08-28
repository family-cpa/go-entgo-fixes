package main

import (
	"log"
	"os"
)

func main() {
	path, _ := os.Getwd()
	log.Fatalf("failed to run fixer in: %s", path)
}
