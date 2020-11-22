package task

import (
	"time"

	"github.com/go-kit/kit/log"
)

type loggingService struct {
	Service
	logger log.Logger
}

// NewLoggingService TODO
func NewLoggingService(logger log.Logger, service Service) Service {
	return &loggingService{service, logger}
}

func (s *loggingService) Find(id uint64) *Task {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "find",
			"id", id,
			"took", time.Since(begin),
		)
	}(time.Now())
	return s.Service.Find(id)
}

func (s *loggingService) Create(task Task) (*Task, error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "create",
			"task", task,
			"took", time.Since(begin),
		)
	}(time.Now())
	return s.Service.Create(task)
}

func (s *loggingService) Update(task Task) error {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "update",
			"task", task,
			"took", time.Since(begin),
		)
	}(time.Now())
	return s.Service.Update(task)
}

func (s *loggingService) Delete(id uint64) error {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "delete",
			"id", id,
			"took", time.Since(begin),
		)
	}(time.Now())
	return s.Service.Delete(id)
}
