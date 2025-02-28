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

	"golang.org/x/crypto/bcrypt" //* для хеширования

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



Авторизация (вход в систему) 
Пользователь вводит логин и пароль
Система ищет пользователя по логину
Проверяет, совпадает ли пароль с хешем
Если все верно → создается JWT-токен
Токен используется для доступа к API
4. Использование JWT-токена
📌 При каждом запросе (например, на создание заметки) пользователь должен отправлять свой JWT-токен. 

Клиент передает токен в заголовке запроса
Сервер проверяет токен
Если токен валиден → запрос выполняется
Если нет → ошибка 401 Unauthorized
5. CRUD-операции для заметок
Создание заметки (POST /notes)
API получает текст заметки
Проверяет, что пользователь авторизован
Создает заметку и привязывает ее к user_id
Просмотр заметок (GET /notes)
API берет user_id из токена
Отправляет только заметки этого пользователя
Удаление заметки (DELETE /notes/:id)
API проверяет, принадлежит ли заметка пользователю
Если да → удаляет
Если нет → 403 Forbidden
6. Защита API (обработка ошибок)
🔒 Безопасность:

Запретить дублирующиеся логины
Проверять, передан ли токен
Ограничить количество попыток входа (чтобы защититься от брутфорса)
📌 Итог
🚀 Логика такая же, но вместо email при авторизации и регистрации используется логин.
Теперь ты можешь создать API без email, используя только логин + пароль. 🔥		