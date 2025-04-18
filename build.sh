#!/bin/bash

# macOS Intel (x86_64)
GOOS=darwin GOARCH=amd64 go build -o ./gotstarter-mac-amd64 ./cmd/gotstarter/main.go

# macOS Apple Silicon (arm64)
GOOS=darwin GOARCH=arm64 go build -o ./gotstarter-mac-arm64 ./cmd/gotstarter/main.go

# Linux 64
GOOS=linux GOARCH=amd64 go build -o ./gotstarter-linux ./cmd/gotstarter/main.go

# Windows 64
GOOS=windows GOARCH=amd64 go build -o ./gotstarter-windows.exe ./cmd/gotstarter/main.go

echo "Build completed."