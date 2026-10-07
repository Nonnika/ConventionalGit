.PHONY: run-test build

run-test:
	go test ./...

build:
	go build -ldflags "-X main.version=$(shell git describe --tags --always --dirty)" .