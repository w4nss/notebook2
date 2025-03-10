package main

import (
	"github.com/labstack/echo/v4"
)

func GetNotes(c echo.Context) error {
	var notes []Note
	rows, err := Database.Query("SELECT id, text FROM notes")
	if err != nil {
		return c.JSON(500, "Ошибка получения заметок")
	}
	defer rows.Close()

	for rows.Next() {
		var note Note
		err := rows.Scan(&note.ID, &note.Text)
		if err != nil {
			return c.JSON(500, "Ошибка при обработке заметок")
		}
		notes = append(notes, note)
	}
	err = rows.Err()
	if err != nil {
		return c.JSON(500, "Ошибка при завершении запроса")
	}
	return c.JSON(200, notes)
}

func CreatedNote(c echo.Context) error {
	text := c.FormValue("text")
	if text == "" {
		return c.JSON(400, "Текст заметки не может быть пустым")
	}

	_, err := Database.Exec("INSERT INTO notes (text) VALUES ($1)", text)
	if err != nil {
		return c.JSON(500, "Ошибка при добавлении заметки")
	}

	return c.JSON(201, "Заметка добавлена")
}

func DeletedNote(c echo.Context) error {
	id := c.Param("id")
	_, err := Database.Exec("DELETE FROM notes WHERE id = $1", id)
	if err != nil {
		c.String(500, "Ошибка удаления заметки")
	}
	return c.JSON(200, "Заметка удалена")
}
