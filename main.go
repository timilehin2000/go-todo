package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/timilehin2000/go-todo/router"
)

func main() {
	route := router.Router()
	log.Fatal(http.ListenAndServe(":8080", route))
	fmt.Println("Server started on port 8080")

}
