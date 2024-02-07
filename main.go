package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/loftwah/soc2/cmd"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: No .env file found")
	}

	cmd.Execute()
}
