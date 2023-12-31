postgres:
	docker run --name postgres --network golangapi-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine
createdb:
	docker exec -it postgres createdb --username=root --owner=root SimpleBank
dropdb:
	docker exec -it postgres dropdb SimpleBank
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/SimpleBank?sslmode=disable" -verbose up
migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/SimpleBank?sslmode=disable" -verbose up 1
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/SimpleBank?sslmode=disable" -verbose down
migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/SimpleBank?sslmode=disable" -verbose down 1
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/PerfectoZ/GoLang-API/db/sqlc Store
.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock migratedown1 migrateup1