postgres:
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14.1-alpine 

startdb:
	docker container start postgres14

createdb:
	docker exec -it postgres14 createdb --username=root --owner=root ms_products

dropdb:
	docker exec -it postgres14 dropdb ms_products

migrateup:
	migrate -path repository/migrations -database "postgresql://root:secret@localhost:5432/ms_products?sslmode=disable" -verbose up

migratedown:
	migrate -path repository/migrations -database "postgresql://root:secret@localhost:5432/ms_products?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

# mock:
# 	mockgen -package mockdb -destination db/mock/store.go github.com/aulas/demo-bank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc startdb server # mock
