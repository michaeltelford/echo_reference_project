.PHONY: help build run test convey lint dep generate migrate alias_migrate

# If VERSION isn't in ENV then use 'latest' docker tag.
VERSION ?= latest

help:
	@echo ""
	@echo "Commands"
	@echo "--------"
	@echo "build         - Builds go binary and docker image"
	@echo "run           - Run the api in docker (with live reloads)"
	@echo "test          - Run the tests"
	@echo "convey        - Run the tests in a UI (retest on file change)"
	@echo "lint          - Run the linter"
	@echo "dep           - Update your dependancies"
	@echo "generate      - Generate mock interfaces etc."
	@echo "migrate       - Run all database migrations"
	@echo "alias_migrate - Prints command to alias migrate for use with api database"
	@echo ""

build:
	GOOS=linux GOARCH=amd64 go build -o bin/api main.go
	docker build -t echo_reference_project:latest .

run: build
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

migrate:
	migrate -database "postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -path "db/migrations" up

alias_migrate:
	@echo "Run this command to alias migrate for easy use with the api DB:"
	@echo "alias migrate='migrate -database postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable -path db/migrations'"
