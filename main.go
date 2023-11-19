package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
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
		DateTime                   int64     `json:"dt"`
		Weather                    []Weather `json:"weather"`
		Main                       Main      `json:"main"`
		ProbabilityOfPrecipitation float64   `json:"pop"`
	} `json:"list"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file!")
	}

	openweatherApiUrl := os.Getenv("OPENWEATHER_API_URL")
	openweatherApiKey := os.Getenv("OPENWEATHER_API_KEY")
	lat := os.Getenv("LOCATION_LAT")
	long := os.Getenv("LOCATION_LONG")

	urlCurrent := openweatherApiUrl + "weather?lat=" + lat + "&lon=" + long + "&appid=" + openweatherApiKey + "&units=metric"
	urlForecast := openweatherApiUrl + "forecast?lat=" + lat + "&lon=" + long + "&appid=" + openweatherApiKey + "&units=metric"

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

	var weatherCurrent Current
	err = json.Unmarshal(body, &weatherCurrent)
	if err != nil {
		panic(err)
	}

	// fetch forecast data
	res, err = http.Get(urlForecast)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weather API not available")
	}

	body, err = io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var weatherForecast Forecast
	err = json.Unmarshal(body, &weatherForecast)
	if err != nil {
		panic(err)
	}

	// print to terminal
	state, country, currentTemp, currentWeatherDesc := weatherCurrent.Name, weatherCurrent.Sys.Country, weatherCurrent.Main.Temperature, weatherCurrent.Weather[0].Description
	fmt.Printf("%s, %s: %.0fC, %s\n", state, country, currentTemp, currentWeatherDesc)

	yNow, mNow, dNow := time.Now().Date()

	forecastList := weatherForecast.List
	for _, f := range forecastList {
		dateTime, forecastTemp, forecastPop, forecastWeather := time.Unix(f.DateTime, 0), f.Main.Temperature, f.ProbabilityOfPrecipitation, f.Weather[0]

		if dateTime.Before(time.Now()) {
			continue
		}

		yForecast, mForecast, dForecast := dateTime.Date()
		if yNow != yForecast || mNow != mForecast || dNow != dForecast {
			continue
		}

		message := fmt.Sprintf("%s - %.0fC, %.0f%%, %s\n", dateTime.Format("15:04"), forecastTemp, forecastPop*100, forecastWeather.Description)
		if forecastWeather.Main != "Rain" {
			fmt.Print(message)
			continue
		}

		if forecastPop < 0.4 {
			fmt.Print(message)
			continue
		}

		color.Red(message)
	}
}
