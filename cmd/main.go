package main

import (
	_ "Gantay/weather/tui"
	"fmt"
	"os"
	_ "time"
)

var settings = Settings{Query: "", ApiKey: ""}

// var opt = [2]string{"help", "v"}

func main() {
	//Load settings if not init.
	settings.load()
	//TODO: add help, edit conf, version print.
	// if len(os.Args) >= 2 {
	// 	settings.Query = os.Args[1]
	// 	settings.save()
	// }
	// USE case/switch
	if len(os.Args) >= 2 {
		for _, v := range os.Args {
			if v == "help" {
				fmt.Println("HELP ME DOWG!!!")
				os.Exit(0)
			} else if v == "v" {
				fmt.Println("version: 0.1 beta")
				os.Exit(0)
			}

		}

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
	currentWeather.JsonWeather(raw)

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
