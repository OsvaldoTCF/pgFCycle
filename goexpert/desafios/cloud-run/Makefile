conf ?= .env
include $(conf)
export $(shell sed 's/=.*//' $(conf))



## ---------- UTILS
.PHONY: help
help: ## Show this menu
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-10s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: clean
clean: ## Clean all temp files
	@rm -f coverage.*



## ----- BUILD & PUSH & DEPLOY
.PHONY: all
all: build push deploy


.PHONY: build
build: ## Build the container image
	@docker build -t us-east1-docker.pkg.dev/goexpert-462214/desafios/cloud-run:v1 .

.PHONY: push
push: ## Push the container image to image registry
	@docker push us-east1-docker.pkg.dev/goexpert-462214/desafios/cloud-run:v1

.PHONY: deploy
deploy: ## Deploy
	@gcloud run deploy cloud-run --project goexpert-462214 --image us-east1-docker.pkg.dev/goexpert-462214/desafios/cloud-run:v1 --platform managed --region us-east1



## ----- COMPOSE
.PHONY: up
up: ## Put the compose containers up
	@docker compose up -d --build

.PHONY: down
down: ## Put the compose containers down
	@docker compose down



## ----- MAIN
.PHONY: serve
serve: ## Run the server
	@cd cmd/app && go run main.go
	@cd -

.PHONY: run
run: ## Make some requests
	@echo -n "422: "; curl -s "http://localhost:8080/cep/2345678"
	@echo -n "404: "; curl -s "http://localhost:8080/cep/68840420"
	@echo -n "200: "; curl -s "http://localhost:8080/cep/58900000"

.PHONY: test
test: ## Run the tests
	@go test -v ./... -coverprofile=coverage.out
	@go tool cover -html coverage.out -o coverage.html
