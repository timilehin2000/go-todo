package router

import "github.com/gorilla/mux"

func router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/todos").Methods("GET")
}
