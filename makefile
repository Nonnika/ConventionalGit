.PHONY: run-test build

run-test:
	go test ./...

build:
	mkdir -p build && go build -ldflags "-X main.version=$(shell git describe --tags --always --dirty)" -o ./build/ .