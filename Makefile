.DEFAULT_GOAL := help
.PHONY: help files build

help:
	@echo "Usage: make [TARGET]"
	@echo "Targets:"
	@echo "files           Show files"
	@echo "build           Build cli"

files:
	@find . -path './.git' -prune -o -ls > FILES

build:
	@go mod tidy && CGO_ENABLED="0" go build -ldflags="-X 'go-cli/app/internal/config/base.APP_NAME=go-cli' -X 'go-cli/app/internal/config/base.APP_VERSION=$CI_COMMIT_REF_NAME'" -o go-cli **/*.go
