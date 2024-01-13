#!/bin/bash

# Set the project name and binary output
PROJECT_NAME="rest-server-go"
OUTPUT_BINARY="bin/$PROJECT_NAME"

# Clean previous build artifacts
go clean

# Format the project code
gofmt -s -w .
gofumpt -w .

# Run golangci-lint
golangci-lint run

# Check if linting passed before proceeding with the build
if [ $? -eq 0 ]; then
    # Build the Go project
    go build -o $OUTPUT_BINARY cmd/rest-server-go/main.go

    # Check if the build was successful
    if [ $? -eq 0 ]; then
        echo "Build successful. Running the project..."
        ./$OUTPUT_BINARY
    else
        echo "Build failed."
    fi
else
    echo "Linting failed. Please fix the linting issues before building."
fi
