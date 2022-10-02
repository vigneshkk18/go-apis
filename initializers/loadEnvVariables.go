package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var DB_URI string

func LoadEnvVariables() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	DB_URI = os.Getenv("MONGODB_URI")

	if DB_URI == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable.")
	}
}
