package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file!")
	}

	openweatherApiKey := os.Getenv("OPENWEATHER_API_KEY")
}
