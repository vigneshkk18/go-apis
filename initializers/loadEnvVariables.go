package initializers

import (
	"log"
	"os"
)

var DB_URI string
var SECRET_KEY string

func LoadEnvVariables() {
	DB_URI = os.Getenv("MONGODB_URI")
	SECRET_KEY = os.Getenv("SECRET_KEY")

	if DB_URI == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable.")
	}
}
