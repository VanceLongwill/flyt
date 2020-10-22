.PHONY: test build

test:
	go test -v ./...

build:
	go build -o flyt cmd/main.go
