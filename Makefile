# Go parameters
CONFIG = config.yaml
SRV_CONFIG = config.services.yaml
SRV_CONFIG_EX = config.services.example.yaml

# Build directory
BUILD_DIR = ./build

# Binary name
BINARY_NAME = logSpy

# Default target
.DEFAULT_GOAL := build

# Build the program
build: clean copyConfig
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/logSpy

# Clean the build
clean:
	@echo "Cleaning..."
	@go clean
	@rm -rf $(BUILD_DIR)

# Copy Config files
copyConfig:
	@echo "Copying config files..."
	@mkdir -p $(BUILD_DIR)
	cp $(CONFIG) $(BUILD_DIR)/$(CONFIG)
	cp $(SRV_CONFIG_EX) $(BUILD_DIR)/$(SRV_CONFIG)

.PHONY: build clean copy-config
