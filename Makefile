compile:
	sqlc generate
	mockery

docker-compose-up:
	docker compose up -d

docker-compose-down:
	docker compose down

build-docker-image:
	docker build -t sokol111/educational-go:latest .

migrate:
	migrate -source file://sql/migration -database 'postgres://localhost:5432/postgres?user=postgres&password=password&sslmode=disable' up

test:
	go test ./... -v -cover

compose-up:
	docker compose up -d

compose-down:
	docker compose down