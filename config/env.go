package config

import (
	"log"

	"github.com/joho/godotenv"
)

// load environment
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
