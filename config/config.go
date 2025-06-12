package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv(filename string) {
	err := godotenv.Load(filename)
	if err != nil {
		log.Println("Error loading .env file") // Don't log fatal error (Github Actions)
	}
}
