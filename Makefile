all: build test

build:
	go get -v -t -d ./...
	go build ./...
	go install ./...

test:
	go test -v ./...
