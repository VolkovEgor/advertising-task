APP=cmd/main.go

build:
	go build -o bin/app.out $(APP)

run:
	go run $(APP)

swag:
	swag init --parseDependency -d ./internal/handler -o ./docs/swagger -g handler.go

SCHEMA=./scripts
DB='postgres://postgres:123matan123@127.0.0.1:5432/advertising_task?sslmode=disable'

migrate_up:
	migrate -path $(SCHEMA) -database $(DB) up

migrate_down:
	migrate -path $(SCHEMA) -database $(DB) down