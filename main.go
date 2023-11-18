package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Weather struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}

type Main struct {
	Temperature float64 `json:"temp"`
}

type Current struct {
	Name string `json:"name"`
	Sys  struct {
		Country string `json:"country"`
	} `json:"sys"`
	Weather []Weather `json:"weather"`
	Main    Main      `json:"main"`
}

type Forecast struct {
	List []struct {
		Weather []Weather `json:"weather"`
		Main    Main      `json:"main"`
	} `json:"list"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file!")
	}

	openweatherApiUrl := os.Getenv("OPENWEATHER_API_ENDPOINT")
	openweatherApiKey := os.Getenv("OPENWEATHER_API_KEY")
	lat := os.Getenv("LOCATION_LAT")
	long := os.Getenv("LOCATION_LONG")

	urlCurrent := openweatherApiUrl + "weather?lat=" + lat + "&lon=" + long + "&appid=" + openweatherApiKey + "&units=metric"
	// urlForecast := openweatherApiUrl + "forecast?lat=" + lat + "&lon=" + long + "&appid=" + openweatherApiKey + "&units=metric"

	// fetch current data
	res, err := http.Get(urlCurrent)
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
