.PHONY: help
help: ## Displays this list of targets with descriptions
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: devenv-build
devenv-build: ## Build the development environment
	@docker compose --projec-directory .devenv build --pull

.PHONY: devenv-start
devenv-start: ## Starts the development environment and logs into main container
	@docker compose --project-directory .devenv up -d
	@docker compose --project-directory .devenv exec main zsh

.PHONY: devenv-stop
devenv-stop: ## Stops the development environment
	@docker compose --project-directory .devenv stop

.PHONY: devenv-down ## Stops and removes the development environment containers
	@docker compose --project-directory .devenv down --remove-orphans


