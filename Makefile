VERSION := $(shell git describe --tags --always)
COMMIT  := $(shell git rev-parse --short HEAD)
DATE    := $(shell date -u +%Y-%m-%dT%H:%M:%SZ)

LDFLAGS := -X 'todo/internal/version.Version=$(VERSION)' \
           -X 'todo/internal/version.Commit=$(COMMIT)' \
           -X 'todo/internal/version.BuildTime=$(DATE)'

build:
	go build -ldflags "$(LDFLAGS)" -o bin/todo .

build-macos:
	GOOS=darwin GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o bin/todo-macos .

build-linux:
	GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o bin/todo-linux .

build-windows:
	GOOS=windows GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o bin/todo.exe .

clean:
	rm -f bin/*
