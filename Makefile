APP=cmd/main.go

build:
	go build -o bin/app.out $(APP)

run:
	go run $(APP)

swag:
	swag init --parseDependency -d ./internal/handler -o ./docs/swagger -g handler.go