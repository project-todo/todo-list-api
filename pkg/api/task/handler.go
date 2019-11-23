package task

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type taskHandler struct {
	database Database
}

func NewHandler(db Database) taskHandler {
	return taskHandler{
		database: db,
	}
}

func (th *taskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
		return
	}

	task, err := th.database.GetByID(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
		return
	}

	data, err := json.Marshal(&task)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(data))
}

func (th *taskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// TODO error handling
	tasks, _ := th.database.Get()
	data, _ := json.Marshal(&tasks)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(data))
}

func (th *taskHandler) PostTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)

	task := NewTask("", "", "")

	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		panic(err)
	}

	err = th.database.Update(id, task)

	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "{\"id\":%d}", id)
}
