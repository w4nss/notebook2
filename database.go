package main

import (
	"database/sql"
	"log"
)

func ConnectionDB() {
	connection := "host=localhost port=5432 user=postgres password=1303 dbname=notebook sslmode=disable"
	var err error
	Database, err = sql.Open("postgres", connection)
	if err != nil {
		log.Fatal("Ошибка подключения:", err)
	}

	err = Database.Ping()
	if err != nil {
		log.Fatal("БД недоступна:", err)
	}
}
