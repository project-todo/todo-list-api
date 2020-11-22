package task

import (
	"testing"
)

type RepositoryTest struct {
	repository Repository
}

func (rt RepositoryTest) TestAll(t *testing.T) {
	t.Run("Create a task", rt.testCreateTask)
	t.Run("Find a task", rt.testFindTask)
	t.Run("Try to find a task that does not exist", rt.testFindTaskThatDoesNotExist)
	t.Run("Update a task", rt.testUpdateTaskThatExist)
	t.Run("Try to update a task that does not exist", rt.testUpdateTaskThatDoesNotExist)
	t.Run("Delete a task", rt.testDeleteOnTaskThatExist)
	t.Run("Try to delete a task that does not exist", rt.testUpdateTaskThatDoesNotExist)
}

func (rt RepositoryTest) testCreateTask(t *testing.T) {
	completed := false
	description := "task description"
	task := Task{
		Completed:   completed,
		Description: description,
	}

	createdTask, err := rt.repository.Create(task)
	if err != nil {
		t.Errorf("Expected error to be nil")
	}

	if createdTask.ID != 0 {
		t.Errorf("Expected 0 but got %d", createdTask.ID)
	}

	assertCompletedIsEqual(t, &task, createdTask)
	assertDescriptionIsEqual(t, &task, createdTask)
}

func (rt RepositoryTest) testFindTask(t *testing.T) {
	task, _ := rt.repository.Create(Task{})

	foundTask := rt.repository.Find(task.ID)
	if foundTask == nil {
		t.Errorf("Expected foundTask not to be nil, but was.")
	}

	assertTaskIsEqual(t, task, foundTask)
}

func (rt RepositoryTest) testFindTaskThatDoesNotExist(t *testing.T) {
	foundTask := rt.repository.Find(1000)
	if foundTask != nil {
		t.Errorf("Expected foundTask not to be nil, but was.")
	}
}

func (rt RepositoryTest) testUpdateTaskThatExist(t *testing.T) {
	task, _ := rt.repository.Create(Task{
		Completed:   false,
		Description: "desc",
	})

	task.Completed = true
	task.Description = "updated desc"

	rt.repository.Update(*task)
	foundTask := rt.repository.Find(task.ID)

	assertTaskIsEqual(t, task, foundTask)
}

func (rt RepositoryTest) testUpdateTaskThatDoesNotExist(t *testing.T) {
	task := Task{
		ID:          1000,
		Completed:   true,
		Description: "description",
	}
	err := rt.repository.Update(task)
	if err == nil {
		t.Error("Expected error from update on task that does not exist")
	}
}

func (rt RepositoryTest) testDeleteOnTaskThatExist(t *testing.T) {
	task, _ := rt.repository.Create(Task{
		Completed:   false,
		Description: "desc",
	})

	rt.repository.Delete(task.ID)
	foundTask := rt.repository.Find(task.ID)
	if foundTask != nil {
		t.Error("Expected find to return nil")
	}
}

func (rt RepositoryTest) testDeleteOnTaskThatDoesNotExist(t *testing.T) {
	err := rt.repository.Delete(1234567)
	if err == nil {
		t.Error("Expected delete to return error, but was nil")
	}
}

func assertTaskIsEqual(t *testing.T, expected, actual *Task) {
	assertIDIsEqual(t, expected, actual)
	assertCompletedIsEqual(t, expected, actual)
	assertDescriptionIsEqual(t, expected, actual)
}

func assertIDIsEqual(t *testing.T, expected, actual *Task) {
	expectedID := expected.ID
	actualID := actual.ID
	if actualID != expectedID {
		t.Errorf("Expected ID to be '%d', but was '%d'", expectedID, actualID)
	}
}

func assertCompletedIsEqual(t *testing.T, expected, actual *Task) {
	expectedCompleted := expected.Completed
	actualCompleted := actual.Completed
	if actualCompleted != expectedCompleted {
		t.Errorf("Expected Completed to be '%t', but was '%t'", expectedCompleted, actualCompleted)
	}
}

func assertDescriptionIsEqual(t *testing.T, expected, actual *Task) {
	expectedDescription := expected.Description
	actualDescription := actual.Description
	if actualDescription != expectedDescription {
		t.Errorf("Expected Description to be '%s', but was '%s'", expectedDescription, actualDescription)
	}
}
