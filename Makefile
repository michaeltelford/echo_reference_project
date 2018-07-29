.PHONY: help run test convey lint dep generate

SHELL = /bin/sh

help:
	@echo ""
	@echo "Commands"
	@echo "--------"
	@echo "run      - Run the app in docker (with live reloads)"
	@echo "test     - Run the tests"
	@echo "convey   - Run the tests in a UI (retest on file change)"
	@echo "lint     - Run the linter"
	@echo "dep      - Update your dependancies"
	@echo "generate - Generate mock interfaces etc."
	@echo ""

run:
	docker-compose up

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
