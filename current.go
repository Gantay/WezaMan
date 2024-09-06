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

	res, err := http.Get("http://api.weatherapi.com/v1/current.json?key=" + apiKey + "&q=" + query + "&aqi=yes&alerts=yes")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weather api not available")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}

	currentTime := time.Now()
	timeString := currentTime.Format("15_02-01-2006")
	fileName := fmt.Sprintf("%s.json", timeString)

	config, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(filepath.Join(config, "WeatherMan", fileName), body, 0740)
	if err != nil {
		panic(err)
	}

	return weather
}

func PrintCurrentWeather(weather Weather) {

	location, current := weather.Location, weather.Current

	fmt.Printf("%s, %s: %.0fC, %s, Time is: %s\n",
		location.Name,
		location.Country,
		current.TemC,
		current.Condition.Text,
		location.TimeLocal,
	)
}
