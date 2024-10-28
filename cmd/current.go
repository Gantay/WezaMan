package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func FetchCurrentWeather(query string, apiKey string) Weather {

	weatherApi := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=yes&alerts=yes", apiKey, query)

	//is retryCount being reset to 0 after the loop is broken??????????
	maxRetries := 10
	retryCount := 0
	resp, err := http.Get(weatherApi)
	if err != nil && resp.StatusCode == 200 {
		for resp.StatusCode != 200 {

			switch resp.StatusCode {
			case 404:
				fmt.Printf("400 Bad Request. retry in 10 seconds. counter:%d", retryCount)
			case 505:
				fmt.Printf("503 Service Unavailable. retry in 10 seconds. counter:%d", retryCount)
			default:
				fmt.Printf("Error code %d. retry in 10 seconds. counter:%d", resp.StatusCode, retryCount)

			}
			time.Sleep(10 * time.Second)
			resp, err = http.Get(weatherApi)

			retryCount++
			if retryCount >= maxRetries {
				fmt.Println("Max retrys reached. Exiting.")
				panic(err)
			}

		}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}

	currentTime := time.Now()
	timeString := currentTime.Format("2006-01-02:_15_")
	fileName := fmt.Sprintf("%s.json", timeString)

	config, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}

	var bodyFormated interface{}
	err = json.Unmarshal(body, &bodyFormated)
	if err != nil {
		panic(err)
	}
	formattedJSON, err := json.MarshalIndent(bodyFormated, "", " ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(filepath.Join(config, "WeatherMan", fileName), formattedJSON, 0740)
	if err != nil {
		panic(err)
	}

	return weather
}

func PrintCurrentWeather(weather Weather) {

	time := time.Unix(weather.Current.TimeOfUpdate, 0)
	ftime := time.Format("15:04")

	fmt.Printf("Location: %s, "+"Temp: %0.fC, "+"Humidity: %d, "+"FeelsLike: %0.fC, "+"UV: %0.f, "+"AQI: %d, "+"TimeOfUpdate: %s \n",
		weather.Location.Name,
		weather.Current.TemC,
		weather.Current.Humidity,
		weather.Current.FeelsLike,
		weather.Current.Uv,
		weather.Current.AirQuality.AQI,
		ftime,
	)
}
