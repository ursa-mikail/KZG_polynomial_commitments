# Project configuration
BINARY_NAME=kzg-demo
MAIN=main.go
DOCKER_IMAGE=kzg-demo:latest

# Default target
all: build

# Build the Go binary
build:
	go build -o $(BINARY_NAME) $(MAIN)

# Run the program locally
run:
	go run $(MAIN)

# Run with test args
run-debug:
	go run $(MAIN) -z 3 -deg 4 -seed 123 -json

# Clean up
clean:
	rm -f $(BINARY_NAME)

# Init and tidy deps
init:
	go mod init kzgdemo
	go get github.com/arnaucube/kzg-commitments-study
	go mod tidy

# Install just dependencies
deps:
	go get github.com/arnaucube/kzg-commitments-study

# Format
fmt:
	go fmt ./...

# Static analysis
vet:
	go vet ./...

# Unit test
test:
	go test -v ./...

# Build Docker image
docker-build:
	docker build -t $(DOCKER_IMAGE) .

# Run container (basic)
docker-run:
	docker run --rm $(DOCKER_IMAGE)

# Run container with params
docker-run-debug:
	docker run --rm $(DOCKER_IMAGE) -z 4 -deg 5 -seed 123 -json

docker-clean:
	docker system prune -a --volumes -f
	
.PHONY: all build run run-debug clean init deps fmt vet test docker-build docker-run docker-run-debug

