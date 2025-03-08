package main

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt" //* для хеширования
)

// * Ввод текста → Добавление в базу данных..
func addNote(reader *bufio.Reader) {
	fmt.Print("Введите текст: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}
	text = strings.TrimSpace(text)

	_, err = Database.Exec("INSERT INTO notes (text) VALUES ($1)", text)
	if err != nil {
		log.Fatal("Ошибка вставки данных:", err)
	}
	fmt.Println("Сохранено!")
}

// Читает все заметки из базы → Выводит их на экран.

func listNotes() {
	rows, err := Database.Query("SELECT id, text FROM notes")
	if err != nil {
		log.Fatal("Ошибка чтения данных:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var text string
		rows.Scan(&id, &text)
		fmt.Println(id, text)
	}
}

func main()

// * Ввод ID заметки → Удаление её из базы.
func deleteNote(reader *bufio.Reader) {
	fmt.Print("Какую заметку удалить, укажите ID заметки: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}
	input = strings.TrimSpace(input)     // ?
	deleteId, err := strconv.Atoi(input) // ?
	if err != nil {
		fmt.Println("Введите корректный числовой ID")
		return
	}

	_, err = Database.Exec("DELETE FROM notes WHERE id = $1", deleteId)
	if err != nil {
		log.Fatal("Ошибка удаления данных:", err)
		return
	}
	fmt.Println("Успешное удаление")
}

// *АВТОРИЗАЦИЯ*
func autorisation() {
	fmt.Println("Вход в аккаунт")
	for {
		fmt.Println("Введите логин")
		var login string
		fmt.Scan(&login)
		for {
			if checkLogin(login) {
				break
			} else {
				fmt.Println("Неверный логин, попробуйте снова:")
			}
		}
	}
}

func checkPassword(password string, login string) bool {
	hashedPassword := getHashedPassword(login)
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		fmt.Println("Пароль неверный")
	}
	return err == nil

}

func checkLogin(login string) bool {
	var exists bool
	err := Database.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE login = $1)", login).Scan(&exists)
	return err == nil && exists
}

func getHashedPassword(login string) string {
	var hashedPassword string
	err := Database.QueryRow("SELECT password_hash FROM users WHERE login = $1", login).Scan(&hashedPassword)
	if err != nil {
		fmt.Println("Такого пользователя не существует")
	}
	return hashedPassword
}

// *РЕГИСТРАЦИЯ*
func registration() {
	fmt.Println("РЕГИСТРАЦИЯ")
	var login string
	for {
		fmt.Println("Введите логин:")
		fmt.Scan(&login)

		if uniqueLogin(login) {
			break
		} else {
			fmt.Println("Введите другой логин:")
		}
	}

	var password string
	fmt.Println("Введите новый пароль:")
	fmt.Scan(&password)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Ошибка хеширования:", err)
		return
	}
	_, err = Database.Exec("INSERT INTO users (login, password_hash) VALUES ($1, $2)", login, hashedPassword)
	if err != nil {
		fmt.Println("Ошибка регистрации", err)
	}
}

func uniqueLogin(login string) bool {
	var exists bool //* exists будет true, если логин уже есть в базе.
	err := Database.QueryRow("SELECT EXISTS (SELECT login FROM users WHERE login = $1)", login).Scan(&exists)
	if err != nil {
		return false
	}
	return !exists // * Если exist=false, то значит логин свободен
}
