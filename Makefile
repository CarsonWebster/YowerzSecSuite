BINARY_NAME=YowerzSecSuite

build: 
	go build -o bin/$(BINARY_NAME) cmd/$(BINARY_NAME)/main.go

run:
	go run cmd/$(BINARY_NAME)/main.go