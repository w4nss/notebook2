package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

var Database *sql.DB

//func main() {

	
	e := echo.New() //Ты создаешь "рацию" (сервер), которая будет слушать запросы

	// Говорим: "Когда придет GET-запрос на адрес /hello — сделай это"
	//e.GET("/hello", func(c echo.Context) error {
        	return c.String(200, "Привет, мир!") // Отправляем текст обратно
   	 })

	e.Start(":8080") // Запускаем сервер на порту 8080 (как сказать "Слушаю!")

e := echo.New()

//func bye(c echo.Context) error { 
	return c.String(200, "Пока")
} 			

e.GET("/bye", bye ) //* -> В браузере ты увидишь белый экран с текстом "Пока"

e.Start(":8080")


	connection := "host=localhost port=5432 user=postgres password=1303 dbname=notebook sslmode=disable"
	var err error
	Database, err = sql.Open("postgres", connection)
	if err != nil {
		log.Fatal("Ошибка подключения:", err)
	}
	defer Database.Close()

}

// ФУНКЦИИ
func menu() {
	reader := bufio.NewReader(os.Stdin)

	// * Пока сервер работает, можно управлять заметками через консоль!
	for {
		fmt.Println("\n📔 Консольный блокнот")
		fmt.Println("1. Добавить заметку")
		fmt.Println("2. Показать заметки")
		fmt.Println("3. Удалить заметку")
		fmt.Println("4. Выйти")
		fmt.Print("👉 Выбери действие: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка чтения ввода:", err)
			continue
		}
		input = strings.TrimSpace(input)

		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("🚫 Введите корректное число")
			continue
		}

		switch choice {
		case 1:
			addNote(reader)
		case 2:
			listNotes()
		case 3:
			deleteNote(reader)
		case 4:
			fmt.Println("👋 До встречи!")
			return
		default:
			fmt.Println("🚫 Неверный выбор")
		}
	}
}
