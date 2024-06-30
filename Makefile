# Go settings
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=commit-bot-go

# Build targets
all: build-linux build-windows build-darwin-amd64 build-darwin-arm64
	@echo "Built for all platforms."

build-linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o dist/linux64/$(BINARY_NAME)
	@echo "Built for Linux 64-bit."

build-windows:
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o dist/windows64/$(BINARY_NAME).exe
	@echo "Built for Windows 64-bit."

build-darwin-amd64:
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o dist/darwin-amd64/$(BINARY_NAME)
	@echo "Built for macOS Intel 64-bit."

build-darwin-arm64:
	GOOS=darwin GOARCH=arm64 $(GOBUILD) -o dist/darwin-arm64/$(BINARY_NAME)
	@echo "Built for macOS Apple Silicon 64-bit."

# Install scripts
install-linux: build-linux
	@echo "Installing on Linux..."
	sudo cp dist/linux64/$(BINARY_NAME) /usr/local/bin/
	@echo "Installed to /usr/local/bin/$(BINARY_NAME)"

install-windows: build-windows
	@echo "Installing on Windows..."
	copy dist\windows64\$(BINARY_NAME).exe "C:\Program Files\$(BINARY_NAME)\$(BINARY_NAME).exe"
	setx path "%path%;C:\Program Files\$(BINARY_NAME)"
	@echo "Installed to C:\Program Files\$(BINARY_NAME) and added to PATH"

install-darwin-amd64: build-darwin-amd64
	@echo "Installing on macOS..."
	sudo cp dist/darwin-amd64/$(BINARY_NAME) /usr/local/bin/
	@echo "Installed to /usr/local/bin/$(BINARY_NAME)"

install-darwin-arm64: build-darwin-arm64
	@echo "Installing on macOS..."
	sudo cp dist/darwin-arm64/$(BINARY_NAME) /usr/local/bin/
	@echo "Installed to /usr/local/bin/$(BINARY_NAME)"

# Clean target
clean:
	$(GOCLEAN)
	rm -rf dist/*
	@echo "Cleaned build artifacts."

# Test target
test:
	$(GOTEST) -v ./...
	@echo "Tests passed."