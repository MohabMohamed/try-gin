package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// DbURL is the url of the used DB.
var DbURL string

// LoadEnv loads enviroment vars required for th project from .env file.
func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file : ", err)
	}
	DbURL = os.Getenv("SQL_LITE_URL")
}
