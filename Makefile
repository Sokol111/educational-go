compile:
	sqlc generate
	mockery

docker-compose-up:
	docker compose up -d

docker-compose-down:
	docker compose down

migrate:
	migrate -source file://sql/migration -database 'postgres://localhost:5432/postgres?user=postgres&password=password&sslmode=disable' up

test:
	go test ./... -v -cover

build:
	go build -o educational-go ./cmd

run: build
	./educational-go

compose-up:
	docker compose up -d

compose-down:
	docker compose down