postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine
createdb:
	docker exec -it postgres createdb --username=root --owner=root SimpleBank
dropdb:
	docker exec -it postgres dropdb SimpleBank
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/SimpleBank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/SimpleBank?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
	make migratedown
	make migrateup

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test