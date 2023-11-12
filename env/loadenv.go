package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() (string, error) {
	// load env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
	ProjectId := os.Getenv("ProjectId")
	return ProjectId, err
}
