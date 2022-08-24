package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// route :=
	log.Fatal(http.ListenAndServe(":8080", route))
	fmt.Println("Server started on port 8080")

}
