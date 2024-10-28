package main

import (
	_ "Gantay/weather/tui"
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

func (s *Settings) init() {
	config, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}

	fmt.Print("API key: ")
	fmt.Scan(&s.ApiKey)

	fmt.Print("Your Location: ")
	fmt.Scan(&s.Query)

	data, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(filepath.Join(config, "WeatherMan"), 0700)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(filepath.Join(config, "WeatherMan", "Setting.json"), data, 0700)
	if err != nil {
		panic(err)
	}

}

func (s *Settings) load() {
	config, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}

	data, err := os.ReadFile(filepath.Join(config, "WeatherMan", "Setting.json"))
	if err != nil {
		fmt.Println("No Setting Found...")
		s.init()
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

	err = os.WriteFile(filepath.Join(config, "WeatherMan", "Setting.json"), data, 0600)
	if err != nil {
		panic(err)
	}

}

var settings = Settings{Query: "", ApiKey: ""}

func main() {
	settings.load()
	// tui.Tea()
	if len(os.Args) >= 2 {
		settings.Query = os.Args[1]
		settings.save()
	}
	FetchCurrentWeather(settings.Query, settings.ApiKey)
	weather := FetchCurrentWeather(settings.Query, settings.ApiKey)
	PrintCurrentWeather(weather)
	//Database(weather)

	ticker := time.NewTicker(15 * time.Minute)
	defer ticker.Stop()

	for t := range ticker.C {
		weather := FetchCurrentWeather(settings.Query, settings.ApiKey)
		PrintCurrentWeather(weather)

		fmt.Println("tick at: ", t.Format("15:04"))
		//Database(weather)
	}
}
