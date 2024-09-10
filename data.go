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
	)

	db, err := sql.Open("sqlite3", "./w.db")
	if err != nil {
		panic(err)
	}

	statement, err := db.Prepare(
		"CREATE TABLE IF NOT EXISTS weather (timeLocal INTEGER PRIMARY KEY UNIQUE,city TEXT, country TEXT, temp FLOAT, humidity INTEGER, windSpeed FLOAT, windDegree FLOAT,feelsLike FLOAT,windChill FLOAT,uv FLOAT)")
	if err != nil {
		panic(err)
	}
	statement.Exec()

	statement, err = db.Prepare("INSERT INTO weather (timeLocal,city,country,temp, humidity, windSpeed, windDegree, feelsLike, windChill, uv) VALUES (?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		panic(err)
	}
	statement.Exec(timeLocal, city, country, temp, humidity, windSpeed, windDegree, feelsLike, windChill, uv)

	rows, _ := db.Query("SELECT timeLocal,city,country,temp, humidity FROM weather")

	for rows.Next() {
		rows.Scan(&timeLocal, &city, &country, &temp, &humidity, &windSpeed, &windDegree, &feelsLike, &windChill, &uv)
		fmt.Println(timeLocal, city, country, temp, humidity, windSpeed, windDegree, feelsLike, windChill, uv)
	}
	db.Close()
}
