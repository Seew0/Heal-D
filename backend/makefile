# Go parameters
BINARY_NAME=heal-d

startdb:
	@echo "Starting DB Instance...."
	docker compose up -d
	@echo "DB Instance Started...."

run: 
	@echo "Running the application..."
	go run internal/cmd/main.go

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

# Clean build files
clean:
	@echo "Cleaning up..."
	rm -f $(BINARY_NAME)

# Run tests
test:
	@echo "Running tests..."
	go test ./...

seed:
	@echo "Seeding the database..."
	go run seeder.go
	@echo "Database seeded successfully."