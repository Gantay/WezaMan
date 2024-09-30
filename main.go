package main

import (
	"encoding/json"
	"fmt"

	"os"
	"path/filepath"
	"time"
)

func SettingsPath(segments ...string) string {
	config, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}
	segments = append([]string{config, "WeatherMan"}, segments...)
	return filepath.Join(segments...)
}

type Settings struct {
	Query  string
	ApiKey string
}

func (s *Settings) load() {
	config, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}

	data, err := os.ReadFile(filepath.Join(config, "WeatherMan", "Setting.json"))
	if err != nil {
		fmt.Println("No Setting Loaded...")
		return
	}

	err = json.Unmarshal(data, &s)
	if err != nil {
		panic(err)
	}
}

func (s *Settings) save() {
	config, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	err = os.MkdirAll(filepath.Join(config, "WeatherMan"), 0700)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(filepath.Join(config, "WeatherMan", "Setting.json"), data, 0600)
	if err != nil {
		panic(err)
	}

}

var settings = Settings{Query: "", ApiKey: ""}

func main() {
	fmt.Println("it's runing")
	settings.load()

	if len(os.Args) >= 2 {
		settings.Query = os.Args[1]
		settings.save()
	}

	weather := FetchCurrentWeather(settings.Query, settings.ApiKey)
	PrintCurrentWeather(weather)
	Database(weather)

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for t := range ticker.C {
		weather := FetchCurrentWeather(settings.Query, settings.ApiKey)
		PrintCurrentWeather(weather)

		fmt.Println("tick at: ", t.Format("15:04"))
		Database(weather)
	}

	// foreCast := fetchForecastWeather(settings.Query, settings.ApiKey)
	// printForecastWeather(foreCast)
	// weather := FetchCurrentWeather(settings.Query, settings.ApiKey)
	// PrintCurrentWeather(weather)

	// Database(weather)

}
