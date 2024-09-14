package main

import (
	"database/sql"
	_ "fmt"

	//"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func Database(weather Weather) {

	var (
		country    string  = weather.Location.Country
		city       string  = weather.Location.Name
		localtime  int64   = weather.Location.TimeLocal
		time       int64   = weather.Current.TimeOfUpdate
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
		defra float32 = weather.Current.AirQuality.Defra
	)

	db, err := sql.Open("sqlite3", "./w.db")
	if err != nil {
		panic(err)
	}

	statement, err := db.Prepare(
		"CREATE TABLE IF NOT EXISTS weather (time INTEGER PRIMARY KEY UNIQUE,localtime INTEGER,city TEXT, country TEXT, temp REAL,code INTEGER, humidity INTEGER, dewPoint REAL, windSpeed REAL, windDegree REAL,gust REAL, feelsLike REAL,heatIndex REAL, windChill REAL,visibility INTEGER,uv REAL)")
	if err != nil {
		panic(err)
	}
	statement.Exec()

	statement, err = db.Prepare(
		"CREATE TABLE IF NOT EXISTS airquality (time INTEGER PRIMARY KEY UNIQUE,co REAL, no2 REAL, o3 REAL, so2 REAL, pm2_5 REAL, pm10 REAL, defra REAL)")
	if err != nil {
		panic(err)
	}
	statement.Exec()

	// createAirQualityTableSQL := "CREATE TABLE IF NOT EXISTS airquality (time INTEGER PRIMARY KEY UNIQUE,co REAL, no2 REAL, o3 REAL, so2 REAL, pm2_5 REAL, pm10 REAL, defra REAL)"
	// _, err = db.Query("SELECT * FROM airquality")
	// if err != nil {
	// 	statement.Exec(createAirQualityTableSQL)
	// }

	statement, err = db.Prepare("INSERT INTO weather (time, localtime, city, country, temp,code, humidity,dewPoint, windSpeed, windDegree,gust, feelsLike,heatIndex, windChill,visibility, uv) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		panic(err)
	}
	statement.Exec(time, localtime, city, country, temp, code, humidity, dewPoint, windSpeed, windDegree, gust, feelsLike, heatIndex, windChill, visibility, uv)

	statement, err = db.Prepare("INSERT INTO airquality (time,co, no2, o3, so2, pm2_5, pm10, defra) VALUES (?,?,?,?,?,?,?,?)")
	if err != nil {
		panic(err)
	}
	statement.Exec(time, co, no2, o3, so2, pm2_5, pm10, defra)

	// rows, _ := db.Query("SELECT time,city,country,temp, humidity FROM weather")

	// for rows.Next() {
	// 	rows.Scan(&time, &city, &country, &temp, &humidity, &windSpeed, &windDegree, &feelsLike, &windChill, &uv)
	// 	fmt.Println(time, city, country, temp, humidity, windSpeed, windDegree, feelsLike, windChill, uv)
	// }

	db.Close()
}
