package router

import (
	"github.com/gorilla/mux"
	"github.com/timilehin2000/go-todo/controllers"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/todos", controllers.GetTodos).Methods("GET")
	router.HandleFunc("/todo/{id}", controllers.GetTodo).Methods("GET")
	// router.HandleFunc()

	return router
}
