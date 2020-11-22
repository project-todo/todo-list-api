package task

// Repository define the contract of communication with a storage solution for tasks.
type Repository interface {
	FindAll() []Task
	Find(id uint64) *Task
	Create(task Task) (*Task, error)
	Update(task Task) error
	Delete(id uint64) error
}
