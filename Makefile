# Binary name
BINARY_NAME=autoDnsQuery
BIN_DIR=bin

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Build targets
.PHONY: all build clean test linux windows darwin

all: clean build

build:
	mkdir -p $(BIN_DIR)
	$(GOBUILD) -o $(BIN_DIR)/$(BINARY_NAME) -v ./

clean:
	$(GOCLEAN)
	rm -rf $(BIN_DIR)

test:
	$(GOTEST) -v ./

# Cross compilation
linux:
	mkdir -p $(BIN_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BIN_DIR)/$(BINARY_NAME)-linux -v ./

windows:
	mkdir -p $(BIN_DIR)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BIN_DIR)/$(BINARY_NAME)-windows.exe -v ./

darwin:
	mkdir -p $(BIN_DIR)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BIN_DIR)/$(BINARY_NAME)-darwin -v ./

# Build for all platforms
build-all: linux windows darwin

# Run
run:
	mkdir -p $(BIN_DIR)
	$(GOBUILD) -o $(BIN_DIR)/$(BINARY_NAME) -v ./
	./$(BIN_DIR)/$(BINARY_NAME)