package common

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvVariable(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	log.Println("Environment variables loaded")

	return os.Getenv(key)
}
