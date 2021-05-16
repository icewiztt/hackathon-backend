postgres:
	docker run --name postgres1 -p 5432:5432 -e POSTGRES_PASSWORD=hackathon2021 -e  POSTGRES_USER=root -d postgres:13-alpine

createdb:
	docker exec -it postgres1 createdb --username=root --owner=root hackathon

dropdb:
	docker exec -it postgres1 dropdb hackathon

migrateup:
	migrate -path db/migration -database "postgresql://root:hackathon2021@localhost:5432/hackathon?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:hackathon2021@localhost:5432/hackathon?sslmode=disable" -verbose down

migrateup1:
	migrate -path db/migration -database "postgresql://root:hackathon2021@localhost:5432/hackathon?sslmode=disable" -verbose up 1

migratedown1:
	migrate -path db/migration -database "postgresql://root:hackathon2021@localhost:5432/hackathon?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server migrate1 migratedown1
