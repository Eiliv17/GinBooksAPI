package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

// loads the environment variables from the .env file
func LoadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
