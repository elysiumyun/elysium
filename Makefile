PROJECT :=  elysium

all: help

.PHONY: help
help:     ## Show this help.
	@echo "Makefile Help Menu >>>\n"
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

.PHONY: tidy
tidy:     ## Go mod tidy.
	@go mod tidy

.PHONY: lint
lint:     ## Lint code.
	@go vet ./...
	@golangci-lint run ./... 
	@go test ./...

.PHONY: generate
generate: ## Go generate deps.
	@go generate -x ./...

.PHONY: check
check:    ## Check deps ...
	@make generate
	@make lint

.PHONY: clean
clean:    ## Clean build cache.
	@rm -rvf bin
	@echo "clean [ ok ]"