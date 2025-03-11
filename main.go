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
	e.POST("/register", Registration)
	e.POST("/login", Login)

	e.GET("/notes", GetNotes, AuthMiddleware)
	e.POST("/notes", CreatedNote, AuthMiddleware)
	e.DELETE("/notes/:id", DeletedNote, AuthMiddleware)

	e.Logger.Fatal(e.Start(":8081"))
}
