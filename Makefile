# Makefile for hs-backend

# Entry point
MAIN_FILE=cmd/app/main.go

# Output directory for swag
SWAG_DIR=docs

# Commands

.PHONY: all docs clean-docs run test coverage coverage-html

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

lint:
	@echo "🚀 Running lint..."
	golangci-lint run --fix
	@echo "✅ Lint passed."

check:
	@echo "🚀 Running check..."
	go fmt ./...
	go vet ./...
	go mod tidy
	golangci-lint run --fix
	@echo "✅ Check passed."


test:
	go test ./... -v

coverage:
	mkdir -p reports
	go test ./... -coverprofile=reports/coverage.out
	@echo "✅ Coverage report generated."

coverage-html:
	go tool cover -html=reports/coverage.out
	@echo "✅ Coverage report generated."
