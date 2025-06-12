ifeq ($(wildcard .env),)
    $(shell cp .env.example .env)
endif

include .env
export


## ---------- UTILS
.PHONY: help
help: ## Show this menu
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: up
up: ## Put the compose containers up
	@docker compose up -d

.PHONY: down
down: ## Put the compose containers down
	@docker compose down
