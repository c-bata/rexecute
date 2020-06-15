NAME := rexecute
VERSION := $(shell git describe --tags --abbrev=0)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -X 'main.version=$(VERSION)' \
           -X 'main.revision=$(REVISION)'
.DEFAULT_GOAL := help

setup:  ## Setup for required tools.
	go get golang.org/x/tools/cmd/goimports
	go get github.com/tcnksm/ghr

fmt: main.go ## Formatting source codes.
	@goimports -w main.go

build: main.go  ## Build a binary.
	go build -ldflags "$(LDFLAGS)"

cross-compile: main.go  ## Build binaries for cross platform.
	mkdir -p pkg
	@for os in "darwin" "linux"; do \
		GOOS=$${os} GOARC=amd64 make build; \
		zip pkg/rexecute_$(VERSION)_$${os}_amd64.zip rexecute; \
	done

release: cross-compile  ## Upload to github releases
	ghr -t $${GITHUB_TOKEN} $(VERSION) pkg/

help: ## Show help text
	@echo "Commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: setup fmt help build cross-compile release
