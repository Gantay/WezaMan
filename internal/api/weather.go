package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Weather struct {
	Location struct {
		Name      string `json:"name"`
		Country   string `json:"country"`
		TimeLocal int64  `json:"localtime_epoch"`
	} `json:"Location"`

	Current struct {
		TimeOfUpdate int64   `json:"last_updated_epoch"`
		TemC         float32 `json:"temp_c"`
		Humidity     int8    `json:"humidity"`
		WindSpeed    float32 `json:"wind_kph"`
		Gust         float32 `json:"gust_kph"`
		WindDegree   float32 `json:"wind_degree"`
		FeelsLike    float32 `json:"feelslike_c"`
		HeatIndex    float32 `json:"heatindex_c"`
		WindChill    float32 `json:"windchill_c"`
		Uv           float32 `json:"uv"`
		DewPoint     float32 `json:"dewpoint_c"`
		Visibility   float32 `json:"vis_km"`
		IsDay        int8    `json:"is_day"`
		Rain         float32 `json:"precip_mm"`
		Condition    struct {
			Text string `json:"text"`
			Code int16  `json:"code"`
		} `json:"condition"`
		AirQuality struct {
			Co    float32 `json:"co"`
			No2   float32 `json:"no2"`
			O3    float32 `json:"o3"`
			So2   float32 `json:"so2"`
			Pm2_5 float32 `json:"pm2_5"`
			Pm10  float32 `json:"pm10"`
			AQI   int8    `json:"us-epa-index"`
		} `json:"air_quality"`
	} `json:"current"`

	Forecast struct {
		Forecastday []struct {
			Hour []struct {
				TimeEpoch int64   `json:"time_epoch"`
				TempC     float64 `json:"temp_c"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
				AirQuality struct {
					Co    float64 `json:"co"`
					No2   float64 `json:"no2"`
					O3    float64 `json:"o3"`
					So2   float64 `json:"so2"`
					Pm2_5 float64 `json:"pm2_5"`
					Pm10  float64 `json:"pm10"`
					Defra float64 `json:"gb-defra-index"`
				} `json:"air_quality"`
				ChanceOfRain float64 `json:"chance_of_rain"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`

	Alerts struct {
		Alert []struct{} `json:"alert"`
	} `json:"alerts"`
}

func FetchCurrentWeather(query string, apiKey string) (*Weather, error) {

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

		resp.Body.Close() // Avoid leaking resources
		fmt.Printf("Unexpected status code: %d. Retrying...\n", resp.StatusCode)

		time.Sleep(time.Duration(retries+1) * 5 * time.Second)

	}

	if resp == nil {
		return nil, fmt.Errorf("failed to fetch weather data after retries: %v", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		return nil, fmt.Errorf("failed to parse weather data: %w", err)
	}

	return &weather, nil
}
