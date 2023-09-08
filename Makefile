# Migrations
migrate-up:
	migrate -path internal/db/migrations -database "postgres://root:secret@localhost:5433/gs_chatbot?sslmode=disable" -verbose up

migrate-down:
	migrate -path internal/db/migrations -database "postgres://root:secret@localhost:5433/gs_chatbot?sslmode=disable" -verbose down

###

# CRUD Generation using sqlc
sqlc-init:
	docker run --rm -v ".:/src" -w /src sqlc/sqlc init

sqlc-generate:
	docker run --rm -v ".:/src" -w /src sqlc/sqlc generate

###

# Simple build-run
build:
	go build -o ./bin/server.exe ./cmd/main.go

run: build
	./bin/server.exe

###

# Air
live-reload:
	air ./bin/main.go

###

.PHONY: migrate-up migrate-down