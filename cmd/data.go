package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path/filepath"
)

// DO I need this????
func SettingsPath(segments ...string) string {
	config, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}
	segments = append([]string{config, "WeatherMan"}, segments...)
	return filepath.Join(segments...)
}

type Settings struct {
	Api    string
	Query  string
	ApiKey string
}

func (s *Settings) init() {
	config, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}

	fmt.Println("Pick one:")
	fmt.Println("1-weatherapi.com, 2.coming soon :), :")
	fmt.Scan(&s.Api)

	fmt.Print("Your Location: ")
	fmt.Scan(&s.Query)

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

// clean this up, bud!!!
func Database(weather Weather) {

	var (
		//Location And Time
		country    string = weather.Location.Country
		city       string = weather.Location.Name
		localtime  int64  = weather.Location.TimeLocal
		timeOfData int64  = weather.Current.TimeOfUpdate

		//WeatherInfo
		temp       float32 = weather.Current.TemC
		code       int16   = weather.Current.Condition.Code
		humidity   int8    = weather.Current.Humidity
		dewPoint   float32 = weather.Current.DewPoint
		windSpeed  float32 = weather.Current.WindSpeed
		windDegree float32 = weather.Current.WindDegree
		gust       float32 = weather.Current.Gust
		feelsLike  float32 = weather.Current.FeelsLike
		heatIndex  float32 = weather.Current.HeatIndex
		windChill  float32 = weather.Current.WindChill
		visibility float32 = weather.Current.Visibility
		uv         float32 = weather.Current.Uv

		//AirQuality
		co    float32 = weather.Current.AirQuality.Co
		no2   float32 = weather.Current.AirQuality.No2
		o3    float32 = weather.Current.AirQuality.O3
		so2   float32 = weather.Current.AirQuality.So2
		pm2_5 float32 = weather.Current.AirQuality.Pm2_5
		pm10  float32 = weather.Current.AirQuality.Pm10
		aqi   int8    = weather.Current.AirQuality.AQI
	)

	//this should work :)
	config, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}
	data := fmt.Sprint(filepath.Join(config, "WeatherMan", "Weather.db"))

	db, err := sql.Open("sqlite3", data)
	if err != nil {
		panic(err)
	}

	statement, err := db.Prepare(
		"CREATE TABLE IF NOT EXISTS weather (timeOfData INTEGER PRIMARY KEY UNIQUE,localtime INTEGER,city TEXT, country TEXT, temp REAL,code INTEGER, humidity INTEGER, dewPoint REAL, windSpeed REAL, windDegree REAL,gust REAL, feelsLike REAL,heatIndex REAL, windChill REAL,visibility INTEGER,uv REAL)")
	if err != nil {
		panic(err)
	}
	statement.Exec()

	statement, err = db.Prepare(
		"CREATE TABLE IF NOT EXISTS airquality (timeOfData INTEGER PRIMARY KEY UNIQUE,co REAL, no2 REAL, o3 REAL, so2 REAL, pm2_5 REAL, pm10 REAL, aqi INTEGER)")
	if err != nil {
		panic(err)
	}
	statement.Exec()

	// createAirQualityTableSQL := "CREATE TABLE IF NOT EXISTS airquality (time INTEGER PRIMARY KEY UNIQUE,co REAL, no2 REAL, o3 REAL, so2 REAL, pm2_5 REAL, pm10 REAL, defra REAL)"
	// _, err = db.Query("SELECT * FROM airquality")
	// if err != nil {
	// 	statement.Exec(createAirQualityTableSQL)
	// }

	statement, err = db.Prepare("INSERT INTO weather (timeOfData, localtime, city, country, temp,code, humidity,dewPoint, windSpeed, windDegree,gust, feelsLike,heatIndex, windChill,visibility, uv) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		panic(err)
	}
	statement.Exec(timeOfData, localtime, city, country, temp, code, humidity, dewPoint, windSpeed, windDegree, gust, feelsLike, heatIndex, windChill, visibility, uv)

	statement, err = db.Prepare("INSERT INTO airquality (timeOfData,co, no2, o3, so2, pm2_5, pm10, aqi) VALUES (?,?,?,?,?,?,?,?)")
	if err != nil {
		panic(err)
	}
	statement.Exec(timeOfData, co, no2, o3, so2, pm2_5, pm10, aqi)

	// rows, _ := db.Query("SELECT time,city,country,temp, humidity FROM weather")

	// for rows.Next() {
	// 	rows.Scan(&time, &city, &country, &temp, &humidity, &windSpeed, &windDegree, &feelsLike, &windChill, &uv)
	// 	fmt.Println(time, city, country, temp, humidity, windSpeed, windDegree, feelsLike, windChill, uv)
	// }

	db.Close()
}
