NAME=rs-ssh-keys
SHELL:=/bin/bash
GOOS=linux
GOARCH=amd64

default: build

build: main.go
	rm -fr build/$(NAME)
	mkdir -p build
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o build/rs-ssh-keys .

.PHONY: build
