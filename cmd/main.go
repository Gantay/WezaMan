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
	//Load settings if not then init.
	settings.load()

	switch {
	case os.Args[1] == "-h", os.Args[1] == "-help":
		fmt.Println("HELP ME DOWG!!!")
		os.Exit(0)
	case os.Args[1] == "-version", os.Args[1] == "-v":
		fmt.Println("version: 0.1 beta")
		os.Exit(0)
	case os.Args[1] == "-location", os.Args[1] == "-l":
		if len(os.Args) <= 2 {
			fmt.Println("no location bud, LOCK IN!!!")
			os.Exit(0)
		}
		settings.Query = os.Args[2]
		fmt.Printf("location Updated to [%s].\n", os.Args[2])
	case os.Args[1] == "env":
		fmt.Println("")

	default:
		fmt.Println("Unknow command")
		os.Exit(0)
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
