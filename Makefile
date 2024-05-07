build:
	docker-compose build chat

run:
	docker-compose up chat

test:
	go test -v ./...

migrate:
	migrate -path ./db/migrations -database 'postgres://postgres:qwerty@0.0.0.0:5436/postgres?sslmode=disable' up

swag:
	swag init -g cmd/main.go


