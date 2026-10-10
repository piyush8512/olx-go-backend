.PHONY: build run

buid:
	@go build -o bin/api ./cmd/api

run:buid
	@./bin/api	

migrate-up:
	@go run ./cmd/migrate up

migrate-down:
	@go run ./cmd/migrate down