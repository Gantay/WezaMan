package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Settings struct {
	Api      string
	Location string
	ApiKey   string
}

func (s *Settings) Init() {
	config, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}

	fmt.Println("Pick one:")
	fmt.Println("1-weatherapi.com, 2.coming soon :)")
	fmt.Scan(&s.Api)

	fmt.Print("Your Location: ")
	fmt.Scan(&s.Location)

	fmt.Print("API key: ")
	fmt.Scan(&s.ApiKey)

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

func (s *Settings) Load() {
	config, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}

	data, err := os.ReadFile(filepath.Join(config, "WeatherMan", "Setting.json"))
	if err != nil {
		fmt.Println("No Setting Found...")
		s.Init()
		return
	}

	err = json.Unmarshal(data, &s)
	if err != nil {
		panic(err)
	}
}

func (s *Settings) Save() {
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

func (s *Settings) Print() {
	fmt.Printf("Weather API service:    %s\n", s.Api)
	fmt.Printf("Weather API ApiKey:     %s\n", s.ApiKey)
	fmt.Printf("Weather API Location:   %s\n", s.Location)

}
