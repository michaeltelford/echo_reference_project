.PHONY: help run debug test

SHELL = /bin/bash

help:
	@echo "Commands"
	@echo "--------"
	@echo "run   - Run the main.go file in production"
	@echo "debug - Run the main.go file in debug via watcher"
	@echo "test  - Run the tests"

run:
	go run cmd/main.go

debug:
	cd cmd && \
	watcher main.go -watch github.com/michaeltelford/echo_reference_project

test:
	go test ./...
