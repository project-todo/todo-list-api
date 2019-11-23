package main

import (
	"net/http"

	"github.com/project-todo/todo-list-api/pkg/api/task"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	db := task.NewMemoryDB()

	taskHandler := task.NewHandler(db)
	r.HandleFunc("/tasks", taskHandler.GetTasks).Methods("GET")
	r.HandleFunc("/tasks", taskHandler.PostTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", taskHandler.GetTask).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
