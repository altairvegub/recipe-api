MODULE = $(shell go list -m)
PACKAGES := $(shell go list ./... | grep -v /vendor/)


.PHONY: fmt
fmt: ## run "go fmt" on all Go packages
	@go fmt $(PACKAGES)

.PHONY: run
run:
	go run cmd/web/main.go
