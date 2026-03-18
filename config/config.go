package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadENv() {
	err := godotenv.Load()

	if err != nil {
		log.Println("No .env file found")
	}
}

func GetPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = "9000"
	}
	return port
}
