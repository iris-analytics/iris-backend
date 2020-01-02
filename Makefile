include .env
export

.PHONY : run test up down ch

run:
	go run cmd/iris-backend.go
test:
	go test ./internal/... -cover fmt
up:
	docker-compose up -d --force-recreate --build
down:
	docker-compose down --remove-orphans
ch:
	docker-compose exec ch sh -c "clickhouse-client -u default --password changeme -d iris"