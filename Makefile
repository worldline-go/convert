BINARY    := convert
MAIN_FILE := cmd/$(BINARY)/main.go

.DEFAULT_GOAL := help

.PHONY: run
run: ## Run the application
	go run $(MAIN_FILE)

.PHONY: help
help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
