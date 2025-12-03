#!make
include .env

PROFILE ?= dev
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

## generate-queries: generate SQLC typed queries
.PHONY: generate-queries
generate-queries:
	sqlc generate

## vet: vet code
.PHONY: vet
vet:
	go vet $(PACKAGES)

## generate: compile templ files
.PHONY: generate
generate: generate-queries
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

## docker-publish: build project into a docker container image
.PHONY: docker-publish
docker-publish: test
	GOPROXY=direct docker buildx build -t $(DOCKERHUB_REPO) .
	docker tag $(DOCKERHUB_REPO) $(DOCKERHUB_URL)
	docker push $(DOCKERHUB_URL)

## start: build and run local project
.PHONY: start
start: build
	air -d

## css: build tailwindcss
.PHONY: css
css:
	tailwindcss -i css/input.css -o css/output.css --minify

## css-watch: watch build tailwindcss
.PHONY: css-watch
css-watch:
	tailwindcss -i css/input.css -o css/output.css --watch --verbose
