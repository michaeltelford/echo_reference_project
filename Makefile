.PHONY: help run debug test lint dep

SHELL = /bin/sh

help:
	@echo "Commands"
	@echo "--------"
	@echo "run   - Run the app in production"
	@echo "debug - Run the app in debug via watcher (live reloads)"
	@echo "test  - Run the tests"
	@echo "lint  - Run the linter"
	@echo "dep   - Run the go dep tool"

run:
	DEBUG=false go run main.go

debug:
	DEBUG=true watcher

test:
	go test ./...

lint:
	golint `go list ./...`

dep:
	dep ensure
