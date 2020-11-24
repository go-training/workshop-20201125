GO ?= go
GOFILES := $(shell find . -name "*.go" -type f)
PACKAGES ?= $(shell $(GO) list ./...)

build: $(GOFILES)
	$(GO) build -v -o bin/server .

test:
	$(GO) test -v ./...
