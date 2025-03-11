package main

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	echoSwagger "github.com/swaggo/echo-swagger"
	"log"

	// Пакет для Swagger UI
	_ "notebook2/docs" // Заменить "your_project" на имя модуля
)

var Database *sql.DB

// @Summary Добавить заметку
// @Description Создает новую заметку
// @Tags notes
// @Accept json
// @Produce json
// @Param note body Note true "Данные заметки"
// @Success 201 {object} Note
// @Router /notes [post]
func CreateNote(c echo.Context) error {
	note := new(Note)
	if err := c.Bind(note); err != nil {
		return c.JSON(400, "Ошибка")
	}
	return c.JSON(201, note)
}

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

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8081"))
}
