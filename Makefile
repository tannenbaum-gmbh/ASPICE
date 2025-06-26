.PHONY: build test clean run

# Default target
all: build

# Build the calculator binary
build:
	go build -o calc ./cmd/calc

# Run tests with verbose output
test:
	go test -v ./...

# Run tests with coverage
test-coverage:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...

# Run linting
lint:
	golangci-lint run

# Clean build artifacts
clean:
	rm -f calc
	rm -f coverage.txt

# Run the calculator (usage: make run ARGS="add 5 3")
run: build
	./calc $(ARGS)

# Install dependencies
deps:
	go mod download
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/securego/gosec/v2/cmd/gosec@latest

# Security scan
security:
	gosec -fmt=json -out=results.json ./...

# Format code
fmt:
	goimports -w .
	gofmt -w .
