package database

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

var db Database

type Database struct {
	handler *sql.DB
}

func GetDatabaseHandler() *sql.DB {
	if db.handler != nil {
		return db.handler
	}

	created, err := CreateDbHandler()

	if err != nil {
		panic(err)
	}
	if !created {
		panic("Failed to create Database handler.")
	}

	return db.handler
}

func CreateDbHandler() (bool, error) {
	cfg := mysql.Config{
		User:                 "user",
		Passwd:               "password",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "GO_MUSIC_FINDER",
		AllowNativePasswords: true,
	}

	var err error
	db.handler, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
		return false, err
	}

	pingErr := db.handler.Ping()
	if pingErr != nil {
		panic(err)
		return false, pingErr
	}
	return true, err
}
