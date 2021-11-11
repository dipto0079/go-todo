package handler

import (
	"fmt"
	"log"
	"net/http"
)

func (h *Handler) CreateTodo(rw http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}
	taskName := r.FormValue("task")
	//log.Println("Tesk name :", taskName)
	fmt.Fprintf(rw, "Task Name:%s", taskName)
}
