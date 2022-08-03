GOPATH:=$(shell go env GOPATH)

.PHONY: init
# init env
init:
	go mod tidy
	go mod download

.PHONY: linux
# build
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

macos:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build

windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build	

.PHONY: all
# generate all
all:
	make linux;
	make macos;

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo '  make [env]'
	@echo ''
	@echo 'Env:'
	@echo '  macos|linux|windows'
	@echo ''


.DEFAULT_GOAL := help
