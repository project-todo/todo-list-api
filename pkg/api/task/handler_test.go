package task

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func createServer(tasks []task) *httptest.Server {
	r := mux.NewRouter()
	db := NewMemoryDB()

	taskHandler := NewHandler(db)
	r.HandleFunc("/tasks", taskHandler.GetTasks).Methods("GET")
	r.HandleFunc("/tasks", taskHandler.PostTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", taskHandler.GetTask).Methods("GET")

	for _, task := range tasks {
		db.Create(task)
	}

	return httptest.NewServer(r)
}

func TestTaskGetEmpty(t *testing.T) {
	expected := make([]task, 0)

	server := createServer(expected)
	defer server.Close()

	resp, err := http.Get(server.URL + "/tasks")

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Received non-200 response: %d\n", resp.StatusCode)
	}

	var got []task

	err = json.NewDecoder(resp.Body).Decode(&got)

	if err != nil {
		t.Fatal(err)
	}

	if len(expected) != len(got) {
		t.Fatalf("Expected %d, got %d", len(expected), len(got))
	}
}

func TestTaskGetID(t *testing.T) {
	tasks := make([]task, 0)
	tasks = append(tasks, NewTask("Get A", "", ""))
	tasks = append(tasks, NewTask("Do B", "", ""))
	tasks = append(tasks, NewTask("Check C", "", ""))

	server := createServer(tasks)
	defer server.Close()

	resp, err := http.Get(server.URL + "/tasks/1")

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Received non-200 response: %d\n", resp.StatusCode)
	}

	var got task

	err = json.NewDecoder(resp.Body).Decode(&got)

	if err != nil {
		t.Fatal(err)
	}

	if 1 != got.ID {
		t.Fatalf("Expected %d, got %d", 1, got.ID)
	}

	if tasks[1].Title != got.Title {
		t.Fatalf("Expected %s, got %s", tasks[1].Title, got.Title)
	}
}

func TestTaskGetInvalidID(t *testing.T) {
	tasks := make([]task, 0)

	server := createServer(tasks)
	defer server.Close()

	resp, err := http.Get(server.URL + "/tasks/999999")

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Fatalf("Expected %d, got %d", http.StatusNotFound, resp.StatusCode)
	}
}

func TestTaskGetAll(t *testing.T) {
	expected := make([]task, 0)
	expected = append(expected, NewTask("Get A", "", ""))
	expected = append(expected, NewTask("Do B", "", ""))
	expected = append(expected, NewTask("Check C", "", ""))

	server := createServer(expected)
	defer server.Close()

	resp, err := http.Get(server.URL + "/tasks")

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Received non-200 response: %d\n", resp.StatusCode)
	}

	var got []task

	err = json.NewDecoder(resp.Body).Decode(&got)

	if err != nil {
		t.Fatal(err)
	}

	if len(expected) != len(got) {
		t.Fatalf("Expected %d, got %d", len(expected), len(got))
	}
}

func TestTaskPost(t *testing.T) {
	t.SkipNow()
}

func TestTaskDelete(t *testing.T) {
	t.SkipNow()
}
