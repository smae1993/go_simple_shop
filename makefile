
.PHONY: run build debug test lint clean install-tools

APP_NAME := web
CMD_PATH := ./cmd/web
BIN_DIR := bin

# Run in development mode
run:
	go run $(CMD_PATH)

# Build production binary
build:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(APP_NAME) $(CMD_PATH)

# Debug with Delve
debug:
	dlv debug $(CMD_PATH) --headless=false

# Run tests
test:
	go test ./...

# Run tests with race detection
test-race:
	go test -race ./...

# Static analysis
vet:
	go vet ./...

# Format code
fmt:
	go fmt ./...

# Sync dependencies
tidy:
	go mod tidy

# Install debugging tools
install-tools:
	go install github.com/go-delve/delve/cmd/dlv@latest

# Remove build artifacts
clean:
	rm -rf $(BIN_DIR)
