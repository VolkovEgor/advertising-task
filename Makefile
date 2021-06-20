APP=cmd/main.go

build:
	go build -o bin/app.out $(APP)

run:
	go run $(APP)

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
DB='postgres://postgres:123matan123@127.0.0.1:5432/advertising_task?sslmode=disable'

migrate_up:
	migrate -path $(SCHEMA) -database $(DB) up

migrate_down:
	migrate -path $(SCHEMA) -database $(DB) down

insert_test_data:
	pgpassword=123matan123 psql -h localhost -U postgres -d advertising_task -f ./scripts/insert_test_data.sql