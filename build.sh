#!/bin/bash

# Build for AMD64
echo "Building for Linux AMD64..."
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o mini_dlna_updater_linux_amd64

# Build for ARM64
echo "Building for Linux ARM64..."
GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -ldflags="-s -w" -o mini_dlna_updater_linux_arm64

echo "Build complete!"
