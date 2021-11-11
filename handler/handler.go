package handler

import "text/template"

type Todo struct {
	Task        string `json:"task"`
	IsCompleted bool   `json:"isCompleted"`
}
type Handler struct {
	templates *template.Template
	todos     []Todo
}

func New(todos []Todo) *Handler {
	h := &Handler{
		todos: todos,
	}
	h.parseTemplate()
	return h
}
func (h *Handler) parseTemplate() {
	h.templates = template.Must(template.ParseFiles(
		"template/create-todo.html",
		"template/list-todo.html",
		"template/edit-todo.html",
	))
}
