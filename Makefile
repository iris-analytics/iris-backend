include .env
export

.PHONY : run test

run:
	go run cmd/iris-backend.go
test:
	go test ./internal/... -cover fmt
up:
	docker-compose up -d
down:
	docker-compose down --remove-orphans