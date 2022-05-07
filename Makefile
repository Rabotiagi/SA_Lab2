build:
	go build ./cmd/example/main.go

test:
	go test

.DEFAULT_GOAL := build