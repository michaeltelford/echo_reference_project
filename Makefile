.PHONY: help run debug test lint dep

SHELL = /bin/sh

help:
	@echo "Commands"
	@echo "--------"
	@echo "run   - Run the app in production"
	@echo "debug - Run the app in debug via watcher"
	@echo "test  - Run the tests"
	@echo "lint  - Run the linter"
	@echo "dep   - Run the go dep tool"

run:
	DEBUG=false go run cmd/main.go

debug:
	cd cmd && \
	DEBUG=true watcher main.go \
	-watch github.com/michaeltelford/echo_reference_project

test:
	go test ./...

lint:
	golint `go list ./...`

dep:
	dep ensure
