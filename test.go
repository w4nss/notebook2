package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/hello", hello)
	e.GET("/bye", bye)
	e.Start(":8081")
}

func hello(c echo.Context) error {
	return c.String(200, "Привет")
}

func bye(c echo.Context) error {
	return c.String(200, "Пока")
}
