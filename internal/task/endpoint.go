package task

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type createTaskRequest struct {
	Description string
}

type createTaskResponse struct {
	ID  uint64 `json:"id"`
	Err error  `json:"error,omitempty"`
}

func makeCreateTaskEndpoint(service Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createTaskRequest)
		task, err := service.Create(Task{Description: req.Description})
		return createTaskResponse{ID: task.ID, Err: err}, nil
	}
}

type findTaskRequest struct {
	ID uint64
}

type findTaskResponse struct {
	Task Task  `json:"task"`
	Err  error `json:"error"`
}

func makeFindTaskEndpoint(service Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(findTaskRequest)
		task := service.Find(req.ID)
		return findTaskResponse{Task: *task, Err: nil}, nil
	}
}
