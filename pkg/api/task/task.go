package task

import (
	"time"
)

type task struct {
	ID          uint64 `json:"id"`
	Completed   bool   `json:"completed"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Created     string `json:"created"`
	Updated     string `json:"updated"`
	Expires     string `json:"expires"`
}

func (t *task) Update(task task) {
	updated := false

	if task.Completed != defaultCompleted {
		t.Completed = task.Completed
		updated = true
	}

	if task.Title != "" && task.Title != t.Title {
		t.Title = task.Title
		updated = true
	}

	if task.Description != "" && task.Description != t.Description {
		t.Description = task.Description
		updated = true
	}

	if task.Expires != "" && task.Expires != t.Expires {
		t.Expires = task.Expires
		updated = true
	}

	if updated {
		t.Updated = time.Now().In(time.UTC).Format(time.RFC3339)
	}
}

var (
	defaultCompleted = false
)

func NewTask(title, description, expiration string) task {
	t := time.Now().In(time.UTC).Format(time.RFC3339)

	return task{
		ID:          0,
		Completed:   defaultCompleted,
		Title:       title,
		Description: description,
		Created:     t,
		Updated:     t,
		Expires:     expiration,
	}
}
