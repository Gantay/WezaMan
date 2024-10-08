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

func fetchForecastWeather(query string, apiKey string) Weather {

	res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=" + apiKey + "&q=" + query + "&days=10&aqi=yes&alerts=yes")
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

	config, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(filepath.Join(config, "WeatherMan", "ww.json"), body, 0740)
	if err != nil {
		panic(err)
	}

	return weather
}

func printForecastWeather(weather Weather) {

	location, current, hours := weather.Location, weather.Current, weather.Forecast.Forecastday[0].Hour

	fmt.Printf("%s, %s: %.0fC, %s, Time is: %d\n",
		location.Name,
		location.Country,
		current.TemC,
		current.Condition.Text,
		location.TimeLocal,
	)
	for _, hour := range hours {
		date := time.Unix(hour.TimeEpoch, 0)

		fmt.Printf(
			"%s - %.01fC, %.0f, %s\n",
			date.Format("15:04"),
			hour.TempC,
			hour.ChanceOfRain,
			hour.Condition.Text,
		)
	}

	airQ := weather.Forecast.Forecastday[0].Hour[0].AirQuality

	fmt.Printf("%.0f Co, %.0f No2, %.0f O3, %.0f So2, %.0f Pm2_5, %.0f Pm10, %.0f Defra\n",
		airQ.Co,
		airQ.No2,
		airQ.O3,
		airQ.So2,
		airQ.Pm2_5,
		airQ.Pm10,
		airQ.Defra,
	)

}
