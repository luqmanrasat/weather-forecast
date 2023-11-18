package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
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

	apiUrl := openweatherApiEndpoint + "?lat=" + lat + "&lon=" + long + "&appid=" + openweatherApiKey

	// fetch forecast data
	res, err := http.Get(apiUrl)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weather API not available")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
