APP_NAME=beta-book-api
BUILD_DIR=bin
CMD_ENTRY=cmd/main.go
SWAG=swag

.PHONY: all swag build run dev clean

all: dev

# Generate Swagger docs
swag:
	@echo "📚 Generating Swagger docs..."
	$(SWAG) init -g $(CMD_ENTRY) -o ./docs

# Build binary
build:
	@echo "🔨 Building app binary..."
	go build -o $(BUILD_DIR)/$(APP_NAME) $(CMD_ENTRY)

# Run binary
run:
	@echo "🚀 Running app..."
	./$(BUILD_DIR)/$(APP_NAME)

# Dev: Generate Swagger + Build + Run
dev:
	@$(MAKE) swag
	@$(MAKE) build
	@$(MAKE) run

# Clean build
clean:
	@echo "🧹 Cleaning build directory..."
	rm -rf $(BUILD_DIR)
