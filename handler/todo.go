package handler

import (
	"net/http"
)

func (h *Handler) CreateTodo(rw http.ResponseWriter, r *http.Request) {
	//if err := r.ParseForm(); err != nil {
	//	log.Fatal(err)
	//}
	//taskName := r.FormValue("task")
	////log.Println("Tesk name :", taskName)
	//fmt.Fprintf(rw, "Task Name:%s", taskName)
	//rw.Header().Add("Content-Type", "text/html")
	//fmt.Fprintf(rw, ``)

	if err := h.templates.ExecuteTemplate(rw, "create-todo.html", nil); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (h *Handler) StoreTodo(rw http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	task := r.FormValue("Task")
	if task == "" {
		http.Redirect(rw, r, "/todos/create", http.StatusTemporaryRedirect)
		return
	}
	h.todos = append(h.todos, Todo{Task: task})
	http.Redirect(rw, r, "/", http.StatusTemporaryRedirect)
}

func (h *Handler) CompleteTodo(rw http.ResponseWriter, r *http.Request) {
	task := r.URL.Path[len("/todos/complete/"):]
	if task == "" {
		http.Error(rw, "invalid URL", http.StatusInternalServerError)
		return
	}
	for i, todo := range h.todos {
		if todo.Task == task {
			h.todos[i].IsCompleted = true
			break
		}
	}
	http.Redirect(rw, r, "/", http.StatusTemporaryRedirect)
}

func (h *Handler) EditTodo(rw http.ResponseWriter, r *http.Request) {
	task := r.URL.Path[len("/todos/edit/"):]
	if task == "" {
		http.Error(rw, "invalid URL", http.StatusInternalServerError)
		return
	}
	var t Todo
	for i, todo := range h.todos {
		if todo.Task == task {
			h.todos[i].IsCompleted = true
			break
		}
	}
	if err := h.templates.ExecuteTemplate(rw, "create-todo.html", t); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) UpdateTodo(rw http.ResponseWriter, r *http.Request) {
	task := r.URL.Path[len("/todos/update/"):]
	if task == "" {
		http.Error(rw, "invalid URL", http.StatusInternalServerError)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	t := r.FormValue("Task")
	if task == "" {
		http.Redirect(rw, r, "/todos/edit"+task, http.StatusTemporaryRedirect)
		return
	}
	for i, todo := range h.todos {
		if todo.Task == task {
			h.todos[i].Task = t
			break
		}
	}
	http.Redirect(rw, r, "/", http.StatusTemporaryRedirect)
}
