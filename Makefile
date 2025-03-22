# Makefile for hs-backend

# Entry point
MAIN_FILE=cmd/app/main.go

# Output directory for swag
SWAG_DIR=docs

# Commands

.PHONY: all docs clean-docs run

all: run

## 🔁 Re-generate Swagger docs
docs:
	@echo "🚀 Generating Swagger docs..."
	@rm -rf $(SWAG_DIR)
	@swag init -g $(MAIN_FILE)
	@echo "✅ Swagger docs generated."

## 🧼 Clean generated Swagger files
clean-docs:
	@echo "🧹 Cleaning Swagger docs..."
	@rm -rf $(SWAG_DIR)
	@echo "✅ Cleaned."

## 🏃 Run the dev server
dev:
	@echo "🚀 Running the server..."
	go run $(MAIN_FILE)