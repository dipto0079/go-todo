package main

import (
	"example.com/m/handler"
	"log"
	"net/http"
)

func main() {

	todos := []handler.Todo{
		{
			Task:        "This is task 1",
			IsCompleted: false,
		},
		{
			Task:        "This is task 2",
			IsCompleted: true,
		},
		{
			Task:        "This is task 3",
			IsCompleted: false,
		},
		{
			Task:        "This is task 4",
			IsCompleted: false,
		},
	}
	h := handler.New(todos)
	http.HandleFunc("/", h.Home)
	http.HandleFunc("/todos/create", h.CreateTodo)
	http.HandleFunc("/todos/store", h.StoreTodo)
	http.HandleFunc("/todos/complete/", h.CompleteTodo)
	http.HandleFunc("/todos/edit/", h.EditTodo)
	http.HandleFunc("/todos/update", h.UpdateTodo)
	log.Println("Server Starting .....")
	if err := http.ListenAndServe("127.0.0.1:3000", nil); err != nil {
		log.Fatal(err)
	}

}
