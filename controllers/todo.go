package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/timilehin2000/go-todo/config"
	"github.com/timilehin2000/go-todo/models"
)

var todo models.Todo

func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos := []models.Todo{}

	db := config.LoadDB()

	sqlStatement := `SELECT * FROM todos`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		err := rows.Scan(&todo.ID, &todo.Item, &todo.IsComplete)

		if err != nil {
			log.Println(fmt.Sprintf("error occured while doing this, :%s", err))
		}

		todos = append(todos, todo)
	}

	buf, err := json.Marshal(todos)

	if err != nil {
		fmt.Println(err)
	}

	w.Write(buf)

	defer db.Close()

	defer rows.Close()
}

func GetTodo(w http.ResponseWriter, r *http.Request) {

}

func AddTodo(w http.ResponseWriter, r *http.Request) {
	db := config.LoadDB()

	json.NewDecoder(r.Body).Decode(&todo)

	defer db.Close()

	sqlStatement := `INSERT INTO todos(item, isComplete) VALUES($1, $2) returning *`

	err := db.QueryRow(sqlStatement, todo.Item, todo.IsComplete).Scan(&todo.ID, &todo.Item, &todo.IsComplete)

	if err != nil {
		log.Println(fmt.Sprintf("error occured while doing this: %s", err))
	}

	json.NewEncoder(w).Encode(todo)
}

func CompleteTodo() {

}

func DeleteTodo() {

}
