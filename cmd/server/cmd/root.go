package cmd

import (
	"net/http"
	"os"

	"github.com/project-todo/todo-list-api/pkg/api/task"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "server [flags]",
	Run: func(cmd *cobra.Command, args []string) {
		r := mux.NewRouter()
		db := task.NewMemoryDB()

		taskHandler := task.NewHandler(db)
		r.HandleFunc("/tasks", taskHandler.GetTasks).Methods("GET")
		r.HandleFunc("/tasks", taskHandler.PostTask).Methods("POST")
		r.HandleFunc("/tasks/{id}", taskHandler.GetTask).Methods("GET")

		http.Handle("/", r)
		http.ListenAndServe(":8080", nil)
	},
}

// Execute runs root command.
func Execute() {
	if rootCmd.Execute() != nil {
		os.Exit(1)
	}
}
