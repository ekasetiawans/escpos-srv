#!/bin/sh

# Build the project
echo "Building [Linux x64] the project..."
GOOS=linux GOARCH=amd64 go build -o ./build/svc-linux-x64 ./

echo "Building [MacOS x64] the project..."
GOOS=darwin GOARCH=amd64 go build -o ./build/svc-darwin-x64 ./

echo "Building [Windows x64] the project..."
GOOS=windows GOARCH=amd64 go build -o ./build/svc-windows-x64.exe ./

echo "Completed"