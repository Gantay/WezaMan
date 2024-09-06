package main

type Weather struct {
	Location struct {
		Name      string `json:"name"`
		Country   string `json:"country"`
		TimeLocal string `json:"localtime"`
	} `json:"Location"`

	Current struct {
		TemC      float64 `json:"temp_c"`
		WindSpeed float64 `json:"wind_kph"`
		Uv        float64 `json:"uv"`
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
