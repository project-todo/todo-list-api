package task

import (
	"fmt"
	"sync"
)

// MemoryRepository TODO
type MemoryRepository struct {
	Repository
	sync.Mutex
	tasks   map[uint64]Task
	counter uint64
}

// NewMemoryRepository TODO
func NewMemoryRepository() MemoryRepository {
	return MemoryRepository{
		tasks:   make(map[uint64]Task),
		counter: 0,
	}
}

// Find TODO
func (m *MemoryRepository) Find(id uint64) *Task {
	defer m.Unlock()
	m.Lock()
	task, ok := m.tasks[id]
	if ok {
		return &task
	}
	return nil
}

// Create TODO
func (m *MemoryRepository) Create(task Task) (*Task, error) {
	defer m.Unlock()
	m.Lock()
	task.ID = m.counter
	m.tasks[task.ID] = task
	m.counter++
	return &task, nil
}

// Update TODO
func (m *MemoryRepository) Update(task Task) error {
	defer m.Unlock()
	m.Lock()
	_, ok := m.tasks[task.ID]
	if !ok {
		return fmt.Errorf("Task with id '%d' not found", task.ID)
	}
	m.tasks[task.ID] = task
	return nil
}

// Delete TODO
func (m *MemoryRepository) Delete(id uint64) error {
	defer m.Unlock()
	m.Lock()
	_, ok := m.tasks[id]
	if !ok {
		return fmt.Errorf("Task with id '%d' not found", id)
	}
	delete(m.tasks, id)
	return nil
}
