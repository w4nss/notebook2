package main

import "github.com/labstack/echo/v4"

func ShowForm(c echo.Context) error {
	html := `
    <!DOCTYPE html>
    <html lang="ru">
    <head>
        <meta charset="UTF-8">
        <title>Добавить заметку</title>
    </head>
    <body>ф
        <h2>Добавить заметку</h2>
        <form action="/notes" method="POST">
            <label for="content">Текст заметки:</label><br>
            <textarea id="content" name="content"></textarea><br>
            <button type="submit">Добавить заметку</button>
        </form>
    </body>
    </html>`
	return c.HTML(200, html)
}
