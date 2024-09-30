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

	weatherApi := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=yes&alerts=yes", apiKey, query)

	//is retryCount being reset to 0 after the loop is broken??????????
	maxRetries := 10
	retryCount := 0
	res, err := http.Get(weatherApi)
	if err != nil {
		// panic(err)
		for err != nil {
			fmt.Println("no connection retry in 10 seconds.", retryCount)
			time.Sleep(10 * time.Second)
			res, err = http.Get(weatherApi)

			retryCount++
			if retryCount >= maxRetries {
				fmt.Println("Max retrys reached. Exiting.")
				panic(err)
			}

		}
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

	location, current := weather.Location, weather.Current
	time := time.Unix(location.TimeLocal, 0)

	fmt.Printf("%s, %s: %.0fC, %s, Time is: %s\n",
		location.Name,
		location.Country,
		current.TemC,
		current.Condition.Text,
		time,
	)
}
