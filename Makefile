all:
	@echo "Building all targets..."
	go build -o cellihub-cli main.go

fmt:
	@echo "Formatting code..."
	go fmt ./...

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