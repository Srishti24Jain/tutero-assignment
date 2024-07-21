#!/bin/bash

# Build the Go application
go build -o runner/main main.go

# Check if build was successful
if [ $? -ne 0 ]; then
    echo "Build failed"
    exit 1
fi

# Build the Docker image
docker build -t learning-roadmap .

# Check if Docker build was successful
if [ $? -ne 0 ]; then
    echo "Docker build failed"
    exit 1
fi

# Run the Docker container
docker run --rm -v "$(pwd)/input.txt:/usr/src/app/input.txt" learning-roadmap
