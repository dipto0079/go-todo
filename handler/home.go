package handler

import (
	"encoding/json"
	"net/http"
)

type Todo struct {
	Task        string `json:"task"`
	IsCompleted bool   `json:"isCompleted"`
}

func (h *Handler) Home(rw http.ResponseWriter, r *http.Request) {
	//rw.Header().Add("Content-Type", "text/html")
	rw.Header().Add("Content-Type", "application/json")
	//fmt.Fprintf(rw, "Hi All Ok")

	todos := []Todo{
		{
			Task:        "this is task 1",
			IsCompleted: false,
		},
		{
			Task:        "this is task 2",
			IsCompleted: true,
		},
	}
	json.NewEncoder(rw).Encode(todos)
}
