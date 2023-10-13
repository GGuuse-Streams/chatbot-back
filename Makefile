# Migrations
migrate-up:
	migrate -path libs/sqlc/migrations -database "postgres://root:secret@localhost:5432/gs_chatbot?sslmode=disable" -verbose up

migrate-down:
	migrate -path libs/sqlc/migrations -database "postgres://root:secret@localhost:5432/gs_chatbot?sslmode=disable" -verbose down
###

## * Deprecated, postgres service moved to docker-compose
## Postgres
#pg-docker:
#	docker run --name postgres15 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine
#
#pg-createdb:
#	docker exec -it postgres15 createdb gs_chatbot
#
#pg-init: pg-docker pg-createdb migrate-up
####

# CRUD Generation using sqlc
sqlc-init:
	docker run --rm -v ".:/src" -w /src sqlc/sqlc init

sqlc-generate:
	docker run --rm -v ".:/src" -w /src sqlc/sqlc generate
###

# GRPC
proto-gen:
	cd libs/grpc && go run gen.go

###

## Simple build-run
#build:
#	go build -o ./bin/server.exe ./cmd/main.go
#
#run: build
#	./bin/server.exe
####
#
#run-m:
#	go run ./apps/bot/cmd/main.go

# Air
live-reload:
	air ./bin/main.go
###

.PHONY: migrate-up migrate-down sqlc-init sqlc-generate live-reload proto-gen