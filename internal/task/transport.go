package task

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// MakeHTTPHandler TODO
func MakeHTTPHandler(service Service, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	createTaskHandler := kithttp.NewServer(
		makeCreateTaskEndpoint(service),
		descodeCreateTaskRequest,
		encodeResponse,
		opts...,
	)

	findTaskHandler := kithttp.NewServer(
		makeFindTaskEndpoint(service),
		descodeCreateTaskRequest,
		encodeResponse,
		opts...,
	)

	router := mux.NewRouter()
	router.Path("/tasks").Methods("POST").Handler(createTaskHandler)
	router.Path("/tasks/{id}").Methods("GET").Handler(findTaskHandler)
	return router
}

func descodeCreateTaskRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var body struct {
		Description string `json:"description"`
	}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	return createTaskRequest{
		Description: body.Description,
	}, nil
}

func descodeFindTaskRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := strconv.ParseUint(idString, 10, 65)

	if err != nil {
		return nil, err
	}

	return findTaskRequest{
		ID: id,
	}, nil
}

type errorer interface {
	error() error
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	err, ok := response.(errorer)
	if ok && err.error() != nil {
		encodeError(ctx, err.error(), w)
	}

	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	switch err {
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
