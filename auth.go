package main

import (
	"database/sql"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"time"
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

var jwtSecret = []byte("my_secret_key") // Храним в ENV (лучше не хардкодить)

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	var storedPassword string
	err := Database.QueryRow("SELECT password FROM users WHERE username = $1", username).Scan(&storedPassword)
	if err == sql.ErrNoRows {
		return c.JSON(401, "Неверный логин или пароль")
	} else if err != nil {
		return c.JSON(500, "Ошибка сервера")
	}

	// Проверяем пароль
	if err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password)); err != nil {
		return c.JSON(401, "Неверный логин или пароль")
	}

	// Создаем токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Токен живет 24 часа
	})

	// Подписываем и выдаем токен
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.JSON(500, "Ошибка при создании токена")
	}

	return c.JSON(200, echo.Map{"token": tokenString})
}

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(401, "Отсутствует токен")
		}

		// Убираем "Bearer " перед токеном
		tokenString = tokenString[len("Bearer "):]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			return c.JSON(401, "Неверный токен")
		}

		return next(c)
	}
}
