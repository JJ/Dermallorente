.PHONY: build lint

build:
	go build ./...

lint:
	golangci-lint run --timeout 5m ./...
