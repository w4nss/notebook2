package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"github.com/labstack/echo/v4"
)

func AddNote(c echo.Context, text) error {
	err := Database.Exec("	")
}

//! Перенеси логику в обработчики

func ListNotes(c echo.Context) error {
	rows, err := Database.Query("SELECT * FROM notes WHERE user_id = $1", user_id)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Не удалось получить заметки"})
	}
	defer rows.Close()

	//Чтение данных 
	var notes []string
	for rows.Next() {
		fmt.Println(rows, "цикл")
		var id int
		var text string
		rows.Scan(&id, &text)
		notes = append(notes, fmt.Sprintf("%d: %s", id, text))
	}
	return c.String(200, strings.Join(notes, "\n"))
}

//Логика:
//Проверяет уникальность email
//Хеширует пароль (bcrypt)
//Сохраняет пользователя в БД (таблица users)
//Ответ: 201 Created + JSON с ID пользователя

