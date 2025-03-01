# Go parameters
BINARY_NAME=heal-d


run: 
	@echo "Running the application..."
	go run internal/cmd/main.go

# Development mode (live reload)
dev:
	@echo "Running in development mode..."
	air 

# Production mode
prod: build
	@echo "Running in production mode..."
	./$(BINARY_NAME)

# Build the application
build:
	@echo "Building the application..."
	go build -o $(BINARY_NAME) .

# Generate Wire dependencies
wire:
	@echo "Generating Wire dependencies..."
	wire

# Start MongoDB (Docker)
start-db:
	@echo "Starting MongoDB..."
	docker run --rm -d \
		--name mongo \
		-p 27017:27017 \
		-e MONGO_INITDB_DATABASE=$(DBNAME) \
		mongo:latest

# Clean build files
clean:
	@echo "Cleaning up..."
	rm -f $(BINARY_NAME)

# Run tests
test:
	@echo "Running tests..."
	go test ./...

.PHONY: dev prod build wire start-db clean test
