package models

import (
	"database/sql"
	"log"
)

var db *sql.DB

func Setup() {
	var err error
	db, err = sql.Open(DRIVER_NAME, DATABASE_CONNECT)

	if err != nil {
		log.Println(err)
	}
}
func CloseDB() {
	defer db.Close()
}
