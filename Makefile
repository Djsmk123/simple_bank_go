postgres:
	docker run --name postgress-db -p 5432:5432 -e POSTGRES_PASSWORD=1235789 postgres

createdb:
	docker exec -it postgress-db createdb --username=postgres --owner=postgres simple_bank

dropdb:
	docker exec -it postgress-db dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgres://postgres:1235789@localhost:5432/simple_bank?sslmode=disable" up

migrateDown:
	.migrate -path db/migration -database "postgres://postgres:1235789@localhost:5432/simple_bank?sslmode=disable" down
test: 
	go test -v -cover ./...
openDB:
	docker exec -it postgress-db psql -U postgres -d simple_bank

mockdb: 
	mockgen --destination db/mock/store.go github.com/djsmk123/simplebank/db/sqlc Store

run:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migrateDown test openDB mockdb 