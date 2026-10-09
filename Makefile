.PHONY: build run

buid:
	@go build -o bin/api ./cmd/api

run:buid
	@./bin/api	