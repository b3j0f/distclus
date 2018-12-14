all: build test

build:
	go get ./...
	go build ./...
	go install ./...

test:
	go test ./...