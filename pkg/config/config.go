package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	env := os.Getenv("ENV")
	err := godotenv.Load()
	if err != nil && env == ""{
		log.Fatal("error loading .env file")
	}
}
