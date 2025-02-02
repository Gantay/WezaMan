package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func FetchCurrentWeather(query string, apiKey string) ([]byte, error) {

	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=yes&alerts=yes", apiKey, query)

	var (
		resp *http.Response
		err  error
	)

	for retries := 0; retries < 10; retries++ {
		resp, err = http.Get(url)
		if err != nil {
			fmt.Printf("HTTP request failed: %v. Retrying...\n", err)
			//TODO: Should Multiply by the n of retries
			time.Sleep(5 * time.Second)
			continue
		}

		if resp.StatusCode == http.StatusOK {
			break
		}

		fmt.Printf("Unexpected status code: %d. Retrying...\n", resp.StatusCode)
		resp.Body.Close() // Avoid leaking resources
		//TODO: Should Multiply by the n of retries
		time.Sleep(5 * time.Second)

	}

	if resp == nil {
		return nil, fmt.Errorf("failed to fetch weather data after retries: %v", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// gotWeather := &Weather
	// err = json.Unmarshal(body, &gotWeather)
	// if err != nil {
	// 	panic(err)
	// }

	// currentTime := time.Now()
	// timeString := currentTime.Format("2006-01-02:_15_")
	// fileName := fmt.Sprintf("%s.json", timeString)

	// config, err := os.UserConfigDir()
	// if err != nil {
	// 	panic(err)
	// }

	// var bodyFormated interface{}
	// err = json.Unmarshal(body, &bodyFormated)
	// if err != nil {
	// 	panic(err)
	// }
	// formattedJSON, err := json.MarshalIndent(bodyFormated, "", " ")
	// if err != nil {
	// 	panic(err)
	// }

	// err = os.WriteFile(filepath.Join(config, "WeatherMan", fileName), formattedJSON, 0740)
	// if err != nil {
	// 	panic(err)
	// }

	return body, nil
}
