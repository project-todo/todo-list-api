package task

// Service define the contract of communication to the tasks service.
type Service interface {
	Find(id uint64) *Task
	Create(task Task) (*Task, error)
	Update(task Task) error
	Delete(id uint64) error
}
