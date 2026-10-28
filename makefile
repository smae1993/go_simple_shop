.PHONY: run build debug test test-race vet fmt tidy install-tools clean tailwind-build tailwind-watch

APP_NAME := web
CMD_PATH := ./cmd/web
BIN_DIR := bin

TAILWIND_BIN := ./tailwindcss
TAILWIND_INPUT := ./web/static/css/input.css
TAILWIND_OUTPUT := ./web/static/css/output.css

# Tailwind: Build and minify CSS for production
tailwind-build:
	$(TAILWIND_BIN) -i $(TAILWIND_INPUT) -o $(TAILWIND_OUTPUT) --minify

# Tailwind: Watch files and rebuild on changes for development
tailwind-watch:
	$(TAILWIND_BIN) -i $(TAILWIND_INPUT) -o $(TAILWIND_OUTPUT) --watch

# Run in development mode (builds CSS first)
run: tailwind-build
	go run $(CMD_PATH)

# Build production binary (builds CSS first)
build: tailwind-build
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
	rm -rf $(BIN_DIR) $(TAILWIND_OUTPUT)
