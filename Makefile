postgres:
	docker-compose up -d

createdb:
	docker exec -it go-simple-bank-postgres createdb -U postgres go_simple_bank_db

dropdb:
	docker exec -it go-simple-bank-postgres dropdb -U postgres go_simple_bank_db

migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/go_simple_bank_db?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/go_simple_bank_db?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc