postgres:
	docker run --name postgress-db -p 5432:5432 -e POSTGRES_PASSWORD=1235789 postgres

createdb:
	docker exec -it postgress-db createdb --username=postgres --owner=postgres simple_bank

dropdb:
	docker exec -it postgress-db dropdb simple_bank

mgup:
	migrate -path db/migration -database "postgres://postgres:1235789@localhost:5432/simple_bank?sslmode=disable" up
mgup1:
	migrate -path db/migration -database "postgres://postgres:1235789@localhost:5432/simple_bank?sslmode=disable" up 1

mgd:
	migrate -path db/migration -database "postgres://postgres:1235789@localhost:5432/simple_bank?sslmode=disable" down
mgd1:
	migrate -path db/migration -database "postgres://postgres:1235789@localhost:5432/simple_bank?sslmode=disable" down 1
test: 
	go test -v -cover ./...
openDB:
	docker exec -it postgress-db psql -U postgres -d simple_bank

mockdb: 
	mockgen --destination db/mock/store.go github.com/djsmk123/simplebank/db/sqlc Store

sqlcgen:
	docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate	

run:
	go run main.go

.PHONY: postgres createdb dropdb mgup mgup1 mgd mgd1  test openDB mockdb 