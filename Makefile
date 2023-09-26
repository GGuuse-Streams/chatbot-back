# Migrations
migrate-up:
	migrate -path internal/db/migrations -database "postgres://root:secret@localhost:5433/gs_chatbot?sslmode=disable" -verbose up

migrate-down:
	migrate -path internal/db/migrations -database "postgres://root:secret@localhost:5433/gs_chatbot?sslmode=disable" -verbose down

pg-createdb:
	docker exec -it postgres15 createdb gs_chatbot

pg-init: pg-createdb migrate-up

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

.PHONY: migrate-up migrate-down sqlc-init sqlc-generate build run live-reload pg-init pg-createdb