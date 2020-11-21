package task

// Manager define the contract of communication to a manager of tasks.
type Manager interface {
	find(id uint64) Task
	create(task Task) (Task, error)
	update(task Task) error
	delete(id uint64) error
}
