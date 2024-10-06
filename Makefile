# Variables
APP_NAME=blu-installment
DOCKER_IMAGE=$(APP_NAME):latest
DOCKERFILE_PATH=./Dockerfile
GO_FILES=$(wildcard *.go)

# Default target
.PHONY: all
all: build

# Build the Go application (Docker build)
.PHONY: build
build:
	@echo "Building the Docker image..."
	docker build -t $(DOCKER_IMAGE) -f $(DOCKERFILE_PATH) .

# Run the application inside the container
.PHONY: run
run:
	@echo "Running the application in a Docker container..."
	docker run -p 8000:8000 --rm $(DOCKER_IMAGE)

# Run the application with arguments (add custom args like --config, --secret)
.PHONY: run-args
run-args:
	@echo "Running the application in a Docker container with arguments..."
	docker run --rm $(DOCKER_IMAGE) --config /path/to/config.yaml --secret /path/to/secret.yaml

# Clean up any dangling Docker images or containers
.PHONY: clean
clean:
	@echo "Cleaning up unused Docker resources..."
	docker system prune -f
	docker rmi $(DOCKER_IMAGE) || true

# Rebuild the Docker image
.PHONY: rebuild
rebuild: clean build

# Build the Go binary locally with dependency handling
.PHONY: build-local
build-local:
	@echo "Tidy up go.mod and regenerate go.sum..."
	go mod tidy
	@echo "Building Go binary locally..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o taurus .

# Clean up the locally built binary
.PHONY: clean-local
clean-local:
	@

# Run unit tests and generate a coverage report
.PHONY: test
test:
	@echo "Running unit tests and generating coverage report..."
	go test -v -coverprofile=coverage.out ./...
	@echo "Generating HTML report from coverage..."
	go tool cover -html=coverage.out -o coverage.html
	@echo "Unit test report generated: coverage.html"

# Clean the test coverage reports
.PHONY: clean-test
clean-test:
	@echo "Cleaning up test coverage reports..."
	rm -f coverage.out coverage.html