include .env

.PHONY: help

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

run: ### Run app
	go build -o bin/app cmd/app/main.go
	./bin/app
.PHONY: run

watch: ### Run app in watch mode
	CompileDaemon --build="go build -o ./bin/app ./cmd/app/main.go" --command="./bin/app" --include=".env"
.PHONY: watch

compose-up: ### Run docker-compose
	docker-compose up --build -d && docker-compose logs -f datings
.PHONY: compose-up

compose-down: ### Down docker-compose
	docker-compose down --remove-orphans
.PHONY: compose-down

test: ### run test
	go test -v -cover -race ./internal/...
.PHONY: test
