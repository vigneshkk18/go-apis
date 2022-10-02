package initializers

import (
	"log"
	"os"
)

var DB_URI string

func LoadEnvVariables() {
	DB_URI = os.Getenv("MONGODB_URI")

	if DB_URI == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable.")
	}
}
