package main

type Note struct {
	ID   int    `json:"id" db:"id"`     //указывает, как поля будут выглядеть в JSON-ответе.
	Text string `json:"text" db:"text"` // показывает, как они называются в БД.
}
