package main

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"

	_ "github.com/lib/pq"
)

var Database *sql.DB

func main() {
	e := echo.New()

	ConnectionDB()
	defer func() {
		if err := Database.Close(); err != nil {
			log.Println("Ошибка при закрытии соединения:", err)
		}
	}()

	fmt.Println("Подключение к БД успешно!")

	e.GET("/", ShowForm)
	e.POST("/notes", CreatedNote)

}

// Подключение к БД
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
