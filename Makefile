.PHONY: help run debug test convey lint dep generate

SHELL = /bin/sh

help:
	@echo "Commands"
	@echo "--------"
	@echo "run      - Run the app in production"
	@echo "debug    - Run the app in debug (live reloads)"
	@echo "test     - Run the tests"
	@echo "convey   - Run the tests in a UI (retest on file change)"
	@echo "lint     - Run the linter"
	@echo "dep      - Update your dependancies"
	@echo "generate - Generate mock interfaces etc."

run:
	DEBUG=false go run main.go

debug:
	DEBUG=true watcher

test:
	go test ./...

convey:
	goconvey

lint:
	golint `go list ./...`

dep:
	dep ensure

generate:
	go generate ./...
