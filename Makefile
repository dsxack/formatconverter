PROJECT=formatconverter
VERSION ?= $(shell git describe --abbrev=4 --dirty --always --tags)

.PHONY: build
build:
	go version
	go env
	go build -ldflags "-X main.Version=$(VERSION)" ./cmd/$(PROJECT)
