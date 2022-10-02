package initializers

import (
	"log"
	"os"
)

var DB_URI string
var PORT string

func LoadEnvVariables() {
	DB_URI = os.Getenv("MONGODB_URI")
	PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "5000"
	}

	if DB_URI == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable.")
	}
}
