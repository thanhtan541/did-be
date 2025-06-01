# List of all Go modules
MODULES := core api

.PHONY: all format lint test

## Format code using gofmt across all modules
format:
	@echo "Formatting Go code..."
	@gofmt -s -w $(shell find . -type f -name '*.go')

## Run linter (golangci-lint) on all modules
lint:
	@echo "Linting Go code..."
	@for dir in $(MODULES); do \
		echo "-> Linting $$dir"; \
		(cd $$dir && golangci-lint run); \
	done

## Run tests in all modules
test:
	@echo "Running tests..."
	@for dir in $(MODULES); do \
		echo "-> Testing $$dir"; \
		go test ./$$dir/...; \
	done

show-todos:
	grep -rni ./README.md -e 'todo'
	grep -rni ./api -e 'todo'
	grep -rni ./core -e 'todo'

## All: run format, lint, test
all: format lint test
