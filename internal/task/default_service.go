package task

// service TODO
type service struct {
	Service
	repository Repository
}

// NewDefaultManager TODO
func NewDefaultManager(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

// Find TODO
func (s *service) Find(id uint64) *Task {
	return s.repository.Find(id)
}

// Create TODO
func (s *service) Create(task Task) (*Task, error) {
	return s.repository.Create(task)
}

// Update TODO
func (s *service) Update(task Task) error {
	return s.Update(task)
}

// Delete TODO
func (s *service) Delete(id uint64) error {
	return s.repository.Delete(id)
}
