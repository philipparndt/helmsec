BINARY_NAME=helmsec
VERSION?=dev
COMMIT=$(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
DATE=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LDFLAGS=-s -w -X helmsec/version.Version=$(VERSION) -X helmsec/version.GitCommit=$(COMMIT) -X helmsec/version.BuildDate=$(DATE)

.PHONY: all build install uninstall clean run tidy

all: build

build:
	go build -ldflags "$(LDFLAGS)" -o $(BINARY_NAME) .

install:
	go install -ldflags "$(LDFLAGS)" .

uninstall:
	rm -f $(shell go env GOPATH)/bin/$(BINARY_NAME)

run:
	go run . $(ARGS)

tidy:
	go mod tidy

clean:
	rm -f $(BINARY_NAME)
