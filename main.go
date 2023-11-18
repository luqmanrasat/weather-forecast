package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file!")
	}

	openweatherApiEndpoint := os.Getenv("OPENWEATHER_API_ENDPOINT")
	openweatherApiKey := os.Getenv("OPENWEATHER_API_KEY")
	lat := os.Getenv("LOCATION_LAT")
	long := os.Getenv("LOCATION_LONG")

	apiUrl := openweatherApiEndpoint + "?lat=" + lat + "&long=" + long + "&appid=" + openweatherApiKey
	fmt.Println(apiUrl)
}
