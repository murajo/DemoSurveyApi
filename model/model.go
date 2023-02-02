package model

import (
	"database/sql"
	"main/config"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", config.DBUser+":"+config.DBPass+"@tcp("+config.DBHost+":"+config.DBPort+")/"+config.DBName+"?utf8mb4_bin")
	if err != nil {
		panic(err.Error())
	}
}
