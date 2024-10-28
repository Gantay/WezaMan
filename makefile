.DEFAULT_GOAL:= run

# Define the name of the binary
BINARY_NAME=WezaMan

# Define the main package (path to main.go or cmd if using a multi-package project)
MAIN_PKG=./cmd

# Go flags (you can set specific Go build flags here)
GO_FLAGS=

# Cross-compilation targets
OS ?= linux
ARCH ?= amd64

# Default target: build the project
.PHONY: all
all: build

# Build the Go project
.PHONY: build
build:
	go build $(GO_FLAGS) -o ./bin/$(BINARY_NAME) $(MAIN_PKG)

# Build for a specific OS/ARCH (cross-compilation)
.PHONY: build-cross
build-cross:
	GOOS=$(OS) GOARCH=$(ARCH) go build $(GO_FLAGS) -o ./bin/$(BINARY_NAME)-$(OS)-$(ARCH) $(MAIN_PKG)



# Run the Go project
.PHONY: run
run:
	go run $(MAIN_PKG)

# Test the Go project
.PHONY: test
test:
	go test ./...

# Clean up binary and other build artifacts
.PHONY: clean
clean:
	@rm -rf ./bin/$(BINARY_NAME)

# Install the Go project binary to the system
.PHONY: install
install:
	go install $(MAIN_PKG)

# Format the code using gofmt
.PHONY: fmt
fmt:
	gofmt -w .

# Vendor dependencies
.PHONY: vendor
vendor:
	go mod vendor
	
# Check for dependency updates using `go get -u` and update them
.PHONY: update-deps
update-deps:
	go get -u ./...
	go mod tidy

# Tidy up the go.mod file
.PHONY: tidy
tidy:
	go mod tidy