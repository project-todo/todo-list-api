package task

import (
	"fmt"
	"strconv"
)

type MemoryDB struct {
	Database

	data map[uint64]task
}

func NewMemoryDB() *MemoryDB {
	return &MemoryDB{
		data: make(map[uint64]task, 0),
	}
}

// DEBUG
func (m *MemoryDB) Clear() {
	m.data = make(map[uint64]task, 0)
}

func (m MemoryDB) Get() ([]task, error) {
	tasks := make([]task, 0)

	for _, task := range m.data {
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (m MemoryDB) GetByID(id uint64) (task, error) {
	t, ok := m.data[id]

	if ok {
		return t, nil
	}

	return task{}, fmt.Errorf("ID not found: " + strconv.FormatUint(id, 10))
}

func (m *MemoryDB) Create(task task) (uint64, error) {
	id := uint64(len(m.data))
	task.ID = id
	m.data[id] = task

	return id, nil
}

func (m *MemoryDB) Update(id uint64, task task) error {
	_, ok := m.data[id]

	if ok {
		m.data[id] = task
		return nil
	}

	_, err := m.Create(task)

	return err
}

func (m *MemoryDB) Delete(id uint64) error {
	_, ok := m.data[id]

	if ok {
		delete(m.data, id)
		return nil
	}

	return fmt.Errorf("ID does not exist: " + strconv.FormatUint(id, 10))
}
