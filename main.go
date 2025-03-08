package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

var Database *sql.DB

func main() {
	connection := "host=localhost port=5432 user=postgres password=1303 dbname=notebook sslmode=disable"
	var err error
	Database, err = sql.Open("postgres", connection)
	if err != nil {
		log.Fatal("Ошибка подключения:", err)
	}
	defer Database.Close()

	// * Теперь запускаем консольное меню
	menu()
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
