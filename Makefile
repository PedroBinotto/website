PACKAGES := $(shell go list ./...)
DB=db/app.db
MIGRATIONS=db/migrations
SCHEMA=db/query/schema.sql
name := $(shell basename ${PWD})

all: help

.PHONY: help
help: Makefile
	@echo
	@echo " Choose a make command to run"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

## init: initialize project (make init module=github.com/user/project)
.PHONY: init
init:
	go mod init ${module}
	go install github.com/air-verse/air@latest
	asdf reshim golang

## migrate-up: run goose UP migrations
.PHONY: migrate-up
migrate-up:
	goose -dir $(MIGRATIONS) sqlite3 $(DB) up

## migrate-up: run goose DOWN migrations (rollback)
.PHONY: migrate-down
migrate-down:
	goose -dir $(MIGRATIONS) sqlite3 $(DB) down

## generate-queries: generate SQLC typed queries
.PHONY: generate-queries
generate-queries: migrate-up
	sqlite3 $(DB) .schema > $(SCHEMA)
	sqlc generate

## vet: vet code
.PHONY: vet
vet:
	go vet $(PACKAGES)

## generate: compile templ files
.PHONY: generate
generate:
	templ generate 

## test: run unit tests
.PHONY: test
test: generate generate-queries
	go test -race -cover $(PACKAGES)

## build: build a binary
.PHONY: build
build: test
	go build -o ./app -v

## docker-build: build project into a docker container image
.PHONY: docker-build
docker-build: test
	GOPROXY=direct docker buildx build -t ${name} .

## docker-run: run project in a container
.PHONY: docker-run
docker-run:
	docker run -it --rm -p 8080:8080 ${name}

## start: build and run local project
.PHONY: start
start: build
	air

## css: build tailwindcss
.PHONY: css
css:
	tailwindcss -i css/input.css -o css/output.css --minify

## css-watch: watch build tailwindcss
.PHONY: css-watch
css-watch:
	tailwindcss -i css/input.css -o css/output.css --watch --verbose
