PACKAGE   := github.com/project-todo/todo-list-api
TARGET    := server
TARGETDIR := bin
ENV       := GOOS=linux

all: vet test ;

vet:
	go vet ./...

test:
	go test ./...
