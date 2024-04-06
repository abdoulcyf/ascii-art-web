.PHONY: build run test

build:
	@echo "Building the project..."
	@go build -o asciiartweb main.go

run: build
	@echo "Running the project..."
	@./asciiartweb

test:
	@echo "Running tests..."
	@go test  -v ./...

clean:
	@echo "Cleaning up..."
	rm -f asciiartweb

