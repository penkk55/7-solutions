# Variables
PACKAGE := ./...
COVERAGE_FILE := coverage.out
COVERAGE_HTML := coverage.html

# Run all tests
test:
	go test $(PACKAGE) -v

# Run tests with coverage
coverage:
	go test $(PACKAGE) -coverprofile=$(COVERAGE_FILE)

# Show coverage in terminal
coverage-show: coverage
	go tool cover -func=$(COVERAGE_FILE)

# Generate HTML report for coverage
css: coverage
	go tool cover -html=$(COVERAGE_FILE) -o $(COVERAGE_HTML)
	@echo "HTML coverage report generated: $(COVERAGE_HTML)"

# Clean up coverage files
clean:
	rm -f $(COVERAGE_FILE) $(COVERAGE_HTML)

# Default command: run tests
default: test

# Install dependencies
install:
	go mod tidy

dev:
	go run main.go