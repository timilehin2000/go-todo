package models

type Todo struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	IsComplete bool   `json:"isDone"`
}
