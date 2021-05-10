.PHONY: build test
.DEFAULT_GOAL := build

build:
	go build -v ./cmd/app

test:
	go test --short -v ./...