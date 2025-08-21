BINARY_NAME=lura

SRC_DIR=./cmd/lura/
BUILD_DIR=./bin

GO=go
GOFLAGS=

.PHONY: all
all: build

.PHONY: build
build: clean
	$(GO) build -o $(BUILD_DIR)/$(BINARY_NAME) $(SRC_DIR)/*.go

.PHONY: run
run: build
	$(BUILD_DIR)/$(BINARY_NAME)

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)/$(BINARY_NAME)

.PHONY: install
install:
	$(GO) install -buildvcs=false

.PHONY: lint
lint:
	golint $(SRC_DIR)/...

.PHONY: test
test:
	$(GO) test -v $(SRC_DIR)/...

.PHONY: format
format:
	$(GO) fmt $(SRC_DIR)/...

.PHONY: help
help:
	@echo "Makefile for Lura"
	@echo "Usage:"
	@echo "  make        Build the application"
	@echo "  make run    Run the application"
	@echo "  make clean  Clean build artifacts"
	@echo "  make test   Run tests"
	@echo "  make lint   Lint the code"
	@echo "  make format Format the code"
	@echo "  make install Install the Go binary"
