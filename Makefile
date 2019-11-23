PACKAGE   := github.com/project-todo/todo-list-api
TARGET    := server
TARGETDIR := bin
ENV       := GOOS=linux

all: vet test build ;

build:
	$(ENV) go build -o $(TARGETDIR)/$(TARGET) $(LDFLAGS) $(PACKAGE)/cmd/$(TARGET)

vet:
	go vet ./...

test:
	go test ./...
