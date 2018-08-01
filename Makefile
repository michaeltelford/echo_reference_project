.PHONY: help run test convey lint dep generate migrate alias_migrate

SHELL = /bin/sh

help:
	@echo ""
	@echo "Commands"
	@echo "--------"
	@echo "run           - Run the app in docker (with live reloads)"
	@echo "test          - Run the tests"
	@echo "convey        - Run the tests in a UI (retest on file change)"
	@echo "lint          - Run the linter"
	@echo "dep           - Update your dependancies"
	@echo "generate      - Generate mock interfaces etc."
	@echo "migrate       - Run all database migrations"
	@echo "alias_migrate - Prints command to alias migrate for use with app database"
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

migrate:
	migrate -database "postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -path "./db/migrations" up

alias_migrate:
	@echo "Run this command to alias migrate for easy use with the app DB:"
	@echo "alias migrate='migrate -database postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable -path ./db/migrations'"
