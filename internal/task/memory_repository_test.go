package task

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCreateTask(t *testing.T) {
	repository := NewMemoryRepository()
	completed := false
	description := "task description"
	task := Task{
		Completed:   completed,
		Description: description,
	}

	createdTask, err := repository.Create(task)
	if err != nil {
		t.Errorf("Expected error to be nil")
	}

	if createdTask.ID != 0 {
		t.Errorf("Expected 0 but got %d", createdTask.ID)
	}

	if createdTask.Completed != completed {
		t.Errorf("Expected completed to be %t but was %t", completed, createdTask.Completed)
	}

	if createdTask.Description != description {
		t.Errorf("Expected description to be %s but was %s", description, createdTask.Description)
	}
}

func TestFindTask(t *testing.T) {
	repository := NewMemoryRepository()
	task, _ := repository.Create(Task{})

	foundTask := repository.Find(task.ID)
	if foundTask == nil {
		t.Errorf("Expected foundTask not to be nil, but was.")
	}
}

func TestFindTaskThatDoesNotExist(t *testing.T) {
	repository := NewMemoryRepository()
	foundTask := repository.Find(1000)
	if foundTask != nil {
		t.Errorf("Expected foundTask to be nil, but was not.")
	}
}

func TestUpdateTask(t *testing.T) {
	repository := NewMemoryRepository()
	task, _ := repository.Create(Task{
		Completed:   false,
		Description: "desc",
	})

	task.Completed = true
	task.Description = "updated desc"

	repository.Update(*task)
	foundTask := repository.Find(task.ID)

	diff := cmp.Diff(task, foundTask)
	if diff != "" {
		t.Errorf("Difference found from expected %s", diff)
	}
}

func TestUpdateTaskThatDoesNotExist(t *testing.T) {
	repository := NewMemoryRepository()
	task := Task{
		ID:          1000,
		Completed:   true,
		Description: "description",
	}
	err := repository.Update(task)
	if err == nil {
		t.Error("Expected error from update on task that does not exist")
	}
}

func TestDeleteOnTaskThatExist(t *testing.T) {
	repository := NewMemoryRepository()
	task, _ := repository.Create(Task{
		Completed:   false,
		Description: "desc",
	})

	repository.Delete(task.ID)
	foundTask := repository.Find(task.ID)
	if foundTask != nil {
		t.Error("Expected find to return nil")
	}
}

func TestDeleteOnTaskThatDoesNotExist(t *testing.T) {
	repository := NewMemoryRepository()

	err := repository.Delete(1234567)
	if err == nil {
		t.Error("Expected delete to return error, but was nil")
	}
}
