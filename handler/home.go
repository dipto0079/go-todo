package handler

import (
	"net/http"
)

type ListData struct {
	Todos []Todo
}

func (h *Handler) Home(rw http.ResponseWriter, r *http.Request) {
	//rw.Header().Add("Content-Type", "text/html")
	//rw.Header().Add("Content-Type", "application/json")
	//fmt.Fprintf(rw, "Hi All Ok")
	//json.NewEncoder(rw).Encode(h.todos)
	lt := ListData{
		Todos: h.todos,
	}
	if err := h.templates.ExecuteTemplate(rw, "list-todo.html", lt); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}
