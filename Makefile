GO_BUILD = go build
GOFLAGS  = CGO_ENABLED=0
DATABASE_HOST     ?= localhost
DATABASE_PORT     ?= $(shell grep "DB_PORT" .env | cut -d '=' -f2)
DATABASE_NAME 	  ?= $(shell grep "DB_NAME" .env | cut -d '=' -f2)
DATABASE_USERNAME ?= $(shell grep "DB_USERNAME" .env | cut -d '=' -f2)
DATABASE_PASSWORD ?= $(shell grep "DB_PASSWORD" .env | cut -d '=' -f2)
DATABSE_DSN       ?= ${DATABASE_USERNAME}:${DATABASE_PASSWORD}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_NAME}

# Version of migrations - this is optionally used on goto command
V?=

# Number of migrations - this is optionally used on up and down commands
N?=



## build: Build app binary
.PHONY: build
build:
	$(GOFLAGS) $(GO_BUILD) -a -v -ldflags="-w -s" -o bin/app main.go