package main

import (
	_ "Gantay/weather/tui"
	"fmt"
	"os"
)

var settings = Settings{Location: "", ApiKey: ""}

func main() {
	//Load settings if not then init.
	settings.load()

	switch {
	case len(os.Args) <= 1:
		break
	case os.Args[1] == "-h", os.Args[1] == "-help":
		fmt.Println("HELP ME DOWG!!!")
		os.Exit(0)
	case os.Args[1] == "-version", os.Args[1] == "-v":
		fmt.Println("version: beta")
		os.Exit(0)
	case os.Args[1] == "-location", os.Args[1] == "-l":
		if len(os.Args) <= 2 {
			fmt.Println("no location bud, LOCK IN!!!")
			os.Exit(1)
		}
		settings.Location = os.Args[2]
		fmt.Printf("location Updated to [%s].\n", os.Args[2])
	case os.Args[1] == "-env":
		settings.print()
		os.Exit(0)

	default:
		fmt.Println("Unknown command. Use -h or -help for usage instructions.")
		os.Exit(0)
	}

	raw, err := FetchCurrentWeather(settings.Location, settings.ApiKey)
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

	//TUI

	//Print the weather to stdout
	currentWeather.PrintWeather()

	// ticker := time.NewTicker(30 * time.Minute)
	// defer ticker.Stop()
	// N := 0
	// for t := range ticker.C {
	// 	N += 1
	// 	raw, err = FetchCurrentWeather(settings.Location, settings.ApiKey)
	// 	if err != nil {
	// 		fmt.Printf("Error fetching weather data: %v\n", err)
	// 		return
	// 	}

	// 	err = currentWeather.UpdateWeather(raw)
	// 	if err != nil {
	// 		fmt.Printf("Error updating weather: %v\n", err)
	// 		return
	// 	}
	// 	currentWeather.PrintWeather()
	// 	fmt.Printf("tick at: %s.  N:%d \n", t.Format("15:04"), N)

	// }
}
