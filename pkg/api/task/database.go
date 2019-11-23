package task

type Database interface {
	Get() ([]task, error)
	GetByID(id uint64) (task, error)
	Create(task task) (uint64, error)
	Update(id uint64, task task) error
	Delete(id uint64) error
}
