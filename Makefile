build:
	go build ./cmd/example/main.go

test:
	go test

fetch:
	go mod download
	go mod verify

.DEFAULT_GOAL := build