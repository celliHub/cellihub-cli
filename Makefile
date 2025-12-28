all: fmt build
	@echo "All tasks completed."

build:
	@echo "Building application..."
	go build -o cellihub-cli main.go

fmt:
	@echo "Formatting code..."
	go fmt ./...

macos-arm64:
	@echo "Building for macOS (arm64)..."
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -o cellihub-cli-darwin-arm64 main.go

macos-amd64:
	@echo "Building for macOS (amd64)..."
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o cellihub-cli-darwin-amd64 main.go

macos: macos-arm64
	@echo "Built macOS (arm64) binary by default. Use macos-amd64 for Intel macs."

clean:
	@echo "Cleaning up..."
	rm -f cellihub-cli

install:
	@echo "Installing application..."
	@cp -v cellihub-cli /usr/local/bin/cellihub-cli

test:
	@echo "Running tests..."
	go test ./...

.PHONY: all clean install test