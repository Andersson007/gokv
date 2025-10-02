# Makefile for the GoKV project

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOFMT=$(GOCMD) fmt
GOLINT=golangci-lint run

# Binary names
SERVER_NAME=gokv-server
CLI_NAME=gokv-cli

# Output directory for binaries
BUILD_DIR=./bin

# All source files
SRC=$(shell find . -type f -name '*.go')

.PHONY: all build test fmt lint clean help

# Default target executed when you just run `make`
all: build

# Builds all binaries
build: build-server build-cli

build-server: $(SRC)
	@echo "Building server..."
	@$(GOBUILD) -o $(BUILD_DIR)/$(SERVER_NAME) ./cmd/gokv-server

build-cli: $(SRC)
	@echo "Building client..."
	@$(GOBUILD) -o $(BUILD_DIR)/$(CLI_NAME) ./cmd/gokv-cli

# Runs all tests
test:
	@echo "Running tests..."
	@$(GOTEST) -v ./...

# Runs tests with coverage report
test-coverage:
	@echo "Running tests with coverage..."
	@$(GOTEST) -v -cover -coverprofile=coverage.out ./...
	@$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Formats the code
fmt:
	@echo "Formatting code..."
	@$(GOFMT) ./...

# Lints the code (requires golangci-lint to be installed)
lint:
	@echo "Linting code..."
	@$(GOLINT) ./...

# Cleans up build artifacts
clean:
	@echo "Cleaning up..."
	@rm -f $(BUILD_DIR)/$(SERVER_NAME)
	@rm -f $(BUILD_DIR)/$(CLI_NAME)
	@rm -f coverage.out coverage.html

# Self-documenting help
help:
	@echo "Available commands:"
	@echo "  make build          - Build the server and client binaries"
	@echo "  make test           - Run all tests"
	@echo "  make test-coverage  - Run tests and generate an HTML coverage report"
	@echo "  make fmt            - Format all Go source files"
	@echo "  make lint           - Lint the codebase"
	@echo "  make clean          - Remove build artifacts"
