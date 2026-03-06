BINARY     := harbor
VERSION    := $(shell cat VERSION 2>/dev/null || echo "dev")
BUILD_DIR  := ./build
INSTALL_DIR := /usr/local/bin

LDFLAGS := -ldflags "-X main.version=$(VERSION) -s -w"

.PHONY: all build install uninstall clean

all: build

## build: compile the binary for the current platform (Linux/amd64 assumed)
build:
	@echo "🔨 Building $(BINARY) $(VERSION)..."
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY) .
	@echo "✅ Binary at $(BUILD_DIR)/$(BINARY)"

## install: build and copy binary to /usr/local/bin (requires sudo)
install: build
	@echo "📦 Installing $(BINARY) to $(INSTALL_DIR)..."
	sudo cp $(BUILD_DIR)/$(BINARY) $(INSTALL_DIR)/$(BINARY)
	sudo chmod +x $(INSTALL_DIR)/$(BINARY)
	@echo "✅ $(BINARY) installed. Run: harbor --help"

## uninstall: remove the binary from /usr/local/bin
uninstall:
	@echo "🗑  Removing $(BINARY) from $(INSTALL_DIR)..."
	sudo rm -f $(INSTALL_DIR)/$(BINARY)
	@echo "✅ Uninstalled."

## clean: remove build artifacts
clean:
	rm -rf $(BUILD_DIR)

## tidy: tidy go modules
tidy:
	go mod tidy
