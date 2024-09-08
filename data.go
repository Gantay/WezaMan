package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func Database() {

	db, err := sql.Open("sqlite3", "./w.db")
	if err != nil {
		panic(err)
	}

	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS weather (date INTEGER PRIMMARY KEY,country TEXT, name TEXT, temp FLOAT, windspeed, FLOAT)")
	if err != nil {
		panic(err)
	}
	statement.Exec()

	statement, err = db.Prepare("INSERT INTO weather (date,country,name,temp, windspeed) VALUES (?,?,?,?,?)")
	if err != nil {
		panic(err)
	}
	statement.Exec(1725554700, "Saudi Arabia", "Dammam", 35.2, 9.0)

	rows, _ := db.Query("SELECT date,country,name,temp, windspeed FROM weather")

	var date int
	var country string
	var name string
	var temp float64
	var windspeed float64

	for rows.Next() {
		rows.Scan(&date, &country, &name, &temp, &windspeed)
		fmt.Println(strconv.Itoa(date)+" : "+country, name, temp, windspeed)
	}
}
