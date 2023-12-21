# Simple Makefile for a Go project

# Build the application
all: build

build:
	@echo "Building..."
	@go build -o bin/main cmd/api/main.go

# Run the application
run: build
	@./bin/main

# Test the application
test:
	@echo "Testing..."
	@go test ./...

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload
watch:
	@air

.PHONY: all build run test clean
