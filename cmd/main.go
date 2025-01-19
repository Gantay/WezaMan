package main

import (
	_ "Gantay/weather/tui"
	"fmt"
	"os"
	_ "time"
)

var settings = Settings{Query: "", ApiKey: ""}

func main() {
	//Load settings if not init.
	settings.load()
	//TODO: add help, edit conf, version print.
	if len(os.Args) >= 2 {
		settings.Query = os.Args[1]
		settings.save()
	}

	raw, err := FetchCurrentWeather(settings.Query, settings.ApiKey)
	if err != nil {
		fmt.Printf("Error fetching weather data: %v\n", err)
		return
	}

	var currentWeather Weather

	err = currentWeather.UpdateWeather(raw)
	if err != nil {
		fmt.Printf("Error updating weather: %v\n", err)
		return
	}
	currentWeather.PrintWeather()

	// ticker := time.NewTicker(2 * time.Minute)
	// defer ticker.Stop()

	// for t := range ticker.C {
	// 	weather, err := FetchCurrentWeather(settings.Query, settings.ApiKey)
	// 	if err != nil {
	// 		fmt.Printf("Error %v", err)
	// 	}
	// 	PrintCurrentWeather(weather)

	// 	fmt.Println("tick at: ", t.Format("15:04"))
	// 	//Database(weather)
	// }
}
