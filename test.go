package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/hello", hello)
	e.Start(":8080")
}

func hello(c echo.Context) error {
	return c.String(200, "Привет")
}
