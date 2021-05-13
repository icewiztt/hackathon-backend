postgres:
	docker run --name postgres1 -p 5432:5432 -e POSTGRES_PASSWORD=Mrcheat2002 -e  POSTGRES_USER=root -d postgres:13-alpine

createdb:
	docker exec -it postgres1 createdb --username=root --owner=root hackathon

dropdb:
	docker exec -it postgres1 dropdb hackathon

migrateup:
	migrate -path db/migration -database "postgresql://root:Mrcheat2002@localhost:5432/hackathon?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:Mrcheat2002@localhost:5432/hackathon?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./db/sqlc/...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test
