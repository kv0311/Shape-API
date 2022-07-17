package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load() //Load .env file
	if err != nil {
		log.Fatal("error: failed to load the env file")
	}
}

func GetConfig(key string) string {
	return os.Getenv(key)
}
