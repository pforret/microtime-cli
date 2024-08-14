#!/usr/bin/env bash
env GOOS=darwin  GOARCH=amd64 go build -o bin/macos/macrotime microtime.go
env GOOS=linux   GOARCH=amd64 go build -o bin/linux/microtime microtime.go
env GOOS=windows GOARCH=amd64 go build -o bin/windows/microtime.exe microtime.go
