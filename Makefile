# Build the application
all: build

setup:
	@go mod download

build:
	@echo "Building..."
	@go build -o bin/wowforge-cli main.go

# Run the application
run:
	@go run main.go

# Test the application
test:
	@echo "Testing..."
	@go test ./...

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f bin/wowforge-cli

.PHONY: all build run test clean
