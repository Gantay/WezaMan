package main

import (
	"database/sql"
	"fmt"

	//"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func Database(weather Weather) {

	var (
		country    string  = weather.Location.Country
		city       string  = weather.Location.Name
		timeLocal  int64   = weather.Location.TimeLocal
		temp       float32 = weather.Current.TemC
		humidity   int8    = weather.Current.Humidity
		windSpeed  float32 = weather.Current.WindSpeed
		windDegree float32 = weather.Current.WindDegree
		feelsLike  float32 = weather.Current.FeelsLike
		windChill  float32 = weather.Current.WindChill
		uv         float32 = weather.Current.Uv
		//
		co    float32 = weather.Current.AirQuality.Co
		no2   float32 = weather.Current.AirQuality.No2
		o3    float32 = weather.Current.AirQuality.O3
		so2   float32 = weather.Current.AirQuality.So2
		pm2_5 float32 = weather.Current.AirQuality.Pm2_5
		pm10  float32 = weather.Current.AirQuality.Pm10
		defra float32 = weather.Current.AirQuality.Defra
	)

	db, err := sql.Open("sqlite3", "./w.db")
	if err != nil {
		panic(err)
	}

	statement, err := db.Prepare(
		"CREATE TABLE IF NOT EXISTS weather (timeLocal INTEGER PRIMARY KEY UNIQUE,city TEXT, country TEXT, temp REAL, humidity INTEGER, windSpeed REAL, windDegree REAL,feelsLike REAL,windChill REAL,uv REAL)")
	if err != nil {
		panic(err)
	}
	statement.Exec()

	statement, err = db.Prepare(
		"CREATE TABLE IF NOT EXISTS airquality (timeLocal INTEGER PRIMARY KEY UNIQUE,co REAL, no2 REAL, o3 REAL, so2 REAL, pm2_5 REAL, pm10 REAL, defra REAL)")
	if err != nil {
		panic(err)
	}
	statement.Exec()

	// createAirQualityTableSQL := "CREATE TABLE IF NOT EXISTS airquality (timeLocal INTEGER PRIMARY KEY UNIQUE,co REAL, no2 REAL, o3 REAL, so2 REAL, pm2_5 REAL, pm10 REAL, defra REAL)"
	// _, err = db.Query("SELECT * FROM airquality")
	// if err != nil {
	// 	statement.Exec(createAirQualityTableSQL)
	// }

	statement, err = db.Prepare("INSERT INTO weather (timeLocal,city,country,temp, humidity, windSpeed, windDegree, feelsLike, windChill, uv) VALUES (?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		panic(err)
	}
	statement.Exec(timeLocal, city, country, temp, humidity, windSpeed, windDegree, feelsLike, windChill, uv)

	statement, err = db.Prepare("INSERT INTO airquality (co, no2, o3, so2, pm2_5, pm10, defra) VALUES (?,?,?,?,?,?,?)")
	if err != nil {
		panic(err)
	}
	statement.Exec(timeLocal, co, no2, o3, so2, pm2_5, pm10, defra)

	rows, _ := db.Query("SELECT timeLocal,city,country,temp, humidity FROM weather")

	for rows.Next() {
		rows.Scan(&timeLocal, &city, &country, &temp, &humidity, &windSpeed, &windDegree, &feelsLike, &windChill, &uv)
		fmt.Println(timeLocal, city, country, temp, humidity, windSpeed, windDegree, feelsLike, windChill, uv)
	}

	db.Close()
}
