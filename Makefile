.DEFAULT_GOAL := build

fmt:
		go fmt ./...
.PHONY:fmt

lint: fmt
		golint ./...
.PHONY:lint

vet: fmt lint
		go vet ./...
		go vet -vettool=$(which shadow) ./...
.PHONY:vet

build: vet
		go build greeting.go
.PHONY:build