# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
GOINSTALL=$(GOCMD) install

# App name
APP_NAME=cash-register-svc

# Main package path
MAIN_PATH=./

# Build output directory
BUILD_DIR=./build

# Binary output path
BINARY=$(BUILD_DIR)/$(APP_NAME)

all: build

build:
	$(GOBUILD) -o $(BINARY) $(MAIN_PATH)

clean:
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)

run:
	$(GOBUILD) -o $(BINARY) $(MAIN_PATH)
	./$(BINARY)
	$(GOINSTALL)

get:
	$(GOGET) -u $(MAIN_PATH)

.PHONY: all build clean run test get