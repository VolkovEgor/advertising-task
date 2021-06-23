APP=cmd/main.go

build:
	docker-compose build app

run:
	docker-compose up app

local_build:
	go build -o bin/app.out $(APP)

local_run:
	go run $(APP) local_config

# WARNING: before running tests need to create database 'postgres_test' in postgres localhost
run_test:
	go test ./... -cover
	go test -tags=e2e

lint:
	go fmt ./...
	golangci-lint run

swag:
	swag init --parseDependency -d ./internal/handler -o ./docs/swagger -g handler.go

SCHEMA=./migrations
DB='postgres://postgres:1234@127.0.0.1:5436/postgres?sslmode=disable'

migrate_up:
	migrate -path $(SCHEMA) -database $(DB) up

migrate_down:
	migrate -path $(SCHEMA) -database $(DB) down

create_test_db:
	pgpassword=1234 psql -h localhost -p 5436 -U postgres -tc "CREATE DATABASE postgres_test"

insert_test_data:
	pgpassword=1234 psql -h localhost -p 5436 -U postgres -d postgres -f ./scripts/insert_test_data.sql