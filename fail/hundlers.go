package main

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
)

//TODO: Перенести логику в обработчики
func ListNotes(c echo.Context) error {
	rows, err := Database.Query("SELECT * FROM notes WHERE user_id = $1", user_id)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Не удалось получить заметки"})
	}

	defer rows.Close()
	var notes 

	jsonData, _ := json.Marshal(rows)

	return jsonData

}

//Логика:
//Проверяет уникальность email
//Хеширует пароль (bcrypt)
//Сохраняет пользователя в БД (таблица users)
//Ответ: 201 Created + JSON с ID пользователя

