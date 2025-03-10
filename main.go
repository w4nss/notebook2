package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"

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

	e.Logger.Fatal(e.Start(":8081"))
}
