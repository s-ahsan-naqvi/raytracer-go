# Variables
GO_FILES := main.go

# Output binary
BINARY_NAME := raytracer-go

# Default target: Build the Go application
all: bin/$(BINARY_NAME)

# Build the Go application
$(BINARY_NAME): $(GO_FILES)
build:
	@go build -o bin/$(BINARY_NAME) $(GO_FILES)

# Run the application
run: build
	./bin/$(BINARY_NAME)

# Clean up generated files
clean:
	rm -f bin/$(BINARY_NAME)

# Rebuild everything
rebuild: clean all

# PHONY targets
.PHONY: all clean rebuild run
