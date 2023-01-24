build:
	@go build -o bin/toronto-pizza ./cmd/toronto-pizza

run: build
	@./bin/toronto-pizza

test:
	@go test -v ./...
