.PHONY: help run debug test

SHELL = /bin/sh

help:
	@echo "Commands"
	@echo "--------"
	@echo "run   - Run the app in production"
	@echo "debug - Run the app in debug via watcher"
	@echo "test  - Run the tests"

run:
	go run cmd/main.go

debug:
	cd cmd && \
	DEBUG=true watcher main.go \
	-watch github.com/michaeltelford/echo_reference_project

test:
	go test ./...
