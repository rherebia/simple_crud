.PHONY: build run test

default: run

build:
	@go build -o bin/api cmd/api/main.go

run:
	@go run cmd/api/main.go

test:
	@go test -v ./...
