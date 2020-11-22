PACKAGE   := github.com/project-todo/todo-list-api
TARGET    := todo/server
TARGETDIR := bin
ENV       := GOOS=linux

all: vet test build;

vet:
	go vet ./...

test:
	go test ./...

build: build-server;

build-server:
	$(ENV) go build -o $(TARGETDIR)/$(TARGET) $(LDFLAGS) $(PACKAGE)/cmd/$(TARGET)