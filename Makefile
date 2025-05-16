.PHONY: all build clean test run lint help

# Application name and main entry point
APP_NAME=api-server
MAIN_PATH=./cmd/main/main.go

# Build output directory
BIN_DIR=./bin

# Go build flags
GOFLAGS=-ldflags="-s -w"

all: clean build test

help:
	@echo "Available commands:"
	@echo "  make build    - Build the application"
	@echo "  make clean    - Remove build artifacts"
	@echo "  make test     - Run tests"
	@echo "  make run      - Run the application"
	@echo "  make lint     - Run linters"
	@echo "  make docker   - Build Docker image"
	@echo "  make help     - Show this help message"

build:
	@echo "Building application..."
	@mkdir -p $(BIN_DIR)
	go build $(GOFLAGS) -o $(BIN_DIR)/$(APP_NAME) $(MAIN_PATH)
	@echo "Build successful: $(BIN_DIR)/$(APP_NAME)"

clean:
	@echo "Cleaning up..."
	@rm -rf $(BIN_DIR)
	@go clean
	@echo "Clean complete"

test:
	@echo "Running tests..."
	go test -v ./test/...
	@echo "Tests complete"

run:
	@echo "Running application..."
	go run $(MAIN_PATH)

lint:
	@echo "Running linters..."
	go vet ./...
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run; \
	else \
		echo "Warning: golangci-lint not installed"; \
	fi
	@echo "Lint complete"

docker:
	@echo "Building Docker image..."
	docker build -t $(APP_NAME) .
	@echo "Docker build complete"