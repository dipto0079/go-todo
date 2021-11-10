package main

import (
	"example.com/m/handler"
	"log"
	"net/http"
)

func main() {
	h := handler.New()
	http.HandleFunc("/", h.Home)
	http.HandleFunc("/todos/create", h.CreateTodo)
	log.Println("Server starting .............")
	if err := http.ListenAndServe("127.0.0.1:3000", nil); err != nil {
		log.Fatal(err)
	}
}
