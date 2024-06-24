## Datings
A simple profile swipes backend service

## Project structure
### `cmd/app/main.go`
Configuration and initialization. Then the main function "continues" in
`internal/app/app.go`.

### `config`
The config structure is in the `config.go` that load from `.env` file

### `integration-test`
Integration tests.

### `internal/app`
Main function after `cmd/app/main.go`

### `internal/controller`
Server handler layer using http (Gin framework)

#### `internal/controller/http/response`
Base JSON response for structured and uniformed response.

#### `internal/controller/http/v1`
Simple REST versioning.

#### `internal/controller/http/validator`
Request body or query validator. Processing validation rules from `github.com/go-playground/validator/v10` and return it as `map[string]string` that show key and value of missing parameter.

### `internal/entity`
Entities of business logic (models) can be used in any layer.

### `internal/usecase`
Business logic.

#### `internal/usecase/repo`
A repository is an abstract storage (database) that business logic works with. Grouped by package used.

### `pkg`
All helper and utility for reusable code.

## Makefile
For convenience, this project uses a Makefile. To see the available commands you can check it by run `make help` to see them.

## Migration
This project has a executable files called `./migration`. Run `./migration help` to see the available commands.

## How to run this project
### Using Docker
#### Please make sure your computer has docker already installed
- After clone this project, setup `.env` file based on the `.env.example`
- Migrate the migrations by using `./migration` binary in this project
- After previous step is already working properly, then run `make compose-up` to start the docker container
### Without Docker
- After clone this project, setup `.env` file based on the `.env.example`
- Migrate the migrations by using `./migration` binary in this project
- Run `make run` to start the app

## Integration Test
- Run `go test ./integration-test/...` to run the integration test from `./integration-test` directory.