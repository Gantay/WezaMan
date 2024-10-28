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
	var (
		resp            *http.Response
		err             error
		maxRetries      = 10
		currentRretries = 0
	)

	for currentRretries >= maxRetries {
		resp, err := http.Get(weatherApi)
		if err != nil {
			fmt.Printf("Request faild: %q", err)
			currentRretries++
			time.Sleep(10 * time.Second)
			continue
		}

		if resp.StatusCode == http.StatusOK {
			break
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
