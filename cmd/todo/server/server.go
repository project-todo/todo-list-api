package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	task "github.com/project-todo/todo-list-api/internal/task"
)

func main() {
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	tr := task.NewMemoryRepository()
	ts := task.NewDefaultManager(tr)
	ts = task.NewLoggingService(log.With(logger, "component", "task"), ts)

	httpLogger := log.With(logger, "component", "http")
	mux := http.NewServeMux()
	mux.Handle("/tasks", task.MakeHTTPHandler(ts, httpLogger))
	http.Handle("/", handlerFunc(mux))

	logger.Log("transport", "http", "address", "9000", "msg", "listening")
	logger.Log("terminated", http.ListenAndServe(":9000", nil))
}

func handlerFunc(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
}
