# Variables
APP_NAME := loan-management-system
SRC := main.go

# Default target
.PHONY: run
run:
	go run $(SRC)

# Build the application
.PHONY: build
build:
	go build -o $(APP_NAME) $(SRC)

# Clean up build artifacts
.PHONY: clean
clean:
	rm -f $(APP_NAME)

# Run tests
.PHONY: test
test:
	go test ./...

# Install dependencies
.PHONY: deps
deps:
	go mod tidy

# Format the code
.PHONY: fmt
fmt:
	go fmt ./...

# Lint the code
.PHONY: lint
lint:
	golangci-lint run

# Docker compose build
.PHONY: docker-build
docker-build:
	docker-compose build

# Docker compose run
.PHONY: docker-up
docker-up:
	docker-compose up

# Docker compose run
.PHONY: docker-down
docker-down:
	docker-compose down

# Help message
.PHONY: help
help:
	@echo "Makefile commands:"
	@echo "  run    - Run the application"
	@echo "  build  - Build the application"
	@echo " docker-build - Docker image build"
	@echo " docker-up - Uses compose to start app with postgres container db"
	@echo " docker-down - Uses compose to stop app with postgres container db"
	@echo "  clean  - Clean up build artifacts"
	@echo "  test   - Run tests"
	@echo "  deps   - Install dependencies"
	@echo "  fmt    - Format the code"
	@echo "  lint   - Lint the code"
	@echo "  help   - Show this help message"
