package main

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Registration(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "" || password == "" {
		return c.JSON(400, "Заполните все поля")
	}

	// Хешируем пароль перед сохранением
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(500, "Ошибка при хешировании пароля")
	}

	// Сохраняем пользователя в БД
	_, err = Database.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", username, string(hashedPassword))
	if err != nil {
		return c.JSON(500, "Ошибка при регистрации")
	}

	return c.JSON(201, "Пользователь зарегистрирован")
}
