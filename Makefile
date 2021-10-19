.DEFAULT_GOAL := build

fmt:
		go fmt ./...
.PHONY:fmt

lint: fmt
		golint ./...
.PHONY:lint

vet: fmt lint
		go vet ./...
		shadow ./...
.PHONY:vet

build: vet
		go build greeting.go
.PHONY:build