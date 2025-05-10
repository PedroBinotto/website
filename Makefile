PROFILE ?= dev
CONTAINER_NAME=website-app
DB=db/app.db
MIGRATIONS=db/migrations
PACKAGES := $(shell go list ./...)
SCHEMA=db/query/schema.sql
NAME := $(shell basename $(PWD))

DOCKERHUB_LOGIN := pedrobinotto
DOCKERHUB_REPO := $(NAME):$(PROFILE)
DOCKERHUB_URL := $(DOCKERHUB_LOGIN)/$(DOCKERHUB_REPO)

ifeq ($(filter $(PROFILE),dev prod),)
	$(error Invalid PROFILE '$(PROFILE)'. Must be 'dev' or 'prod')
endif

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
	go mod init $(module)
	go install github.com/air-verse/air@latest
	asdf reshim golang

## migrate-up: run goose UP migrations
.PHONY: migrate-up
migrate-up:
	goose -dir $(MIGRATIONS) sqlite3 $(DB) up

## migrate-down: run goose DOWN migrations (rollback)
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

## clean: clean up generated and compiled files
.PHONY: clean
clean:
	rm -rf app css/output.css db/app.db db/query/schema.sql sqlc-generated templates/**/*.go tmp

## docker-build: build project into a docker container image
.PHONY: docker-build
docker-build:
	GOPROXY=direct docker buildx build --progress=plain -t $(DOCKERHUB_REPO) .
	docker tag $(DOCKERHUB_REPO) $(DOCKERHUB_URL)
	docker push $(DOCKERHUB_URL)

## docker-run: run project in a container
.PHONY: docker-run
docker-run:
	docker run -p 8080:8080 -v /website/data:/app/db --name $(CONTAINER_NAME) --detach $(DOCKERHUB_REPO)

## docker-rm: remove running app container
.PHONY: docker-rm
docker-rm:
	docker rm -f $(CONTAINER_NAME)

## start: build and run local project
.PHONY: start
start: build
	SQLITE_DB=$(DB) air

## css: build tailwindcss
.PHONY: css
css:
	tailwindcss -i css/input.css -o css/output.css --minify

## css-watch: watch build tailwindcss
.PHONY: css-watch
css-watch:
	tailwindcss -i css/input.css -o css/output.css --watch --verbose
