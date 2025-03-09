package main

type Note struct {
	ID   int    `json:"id" db:"id"`
	Text string `json:"text" db:"text"`
}
