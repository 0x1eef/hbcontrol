.PHONY: all fmt build release

all: build

fmt:
	for go in cmd/control/*.go internal/*/*.go; do \
		go fmt $$go; \
	done; \

build:
	go build -o bin/control ./cmd/control

release:
	go build -ldflags="-s -w" -o bin/control ./cmd/control