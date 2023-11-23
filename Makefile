test:
	go test ./... -v -cover

build:
	go build -o educational-go ./cmd

run: build
	./educational-go
