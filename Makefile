createdb:
	docker exec -it postgres16 createdb --username=root --owner=root url-shortening
create-migrate-file:
	migrate create -ext sql -dir db/migration -seq init_schema
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/url-shortening?sslmode=disable" -verbose up
migrateuplastest:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/url-shortening?sslmode=disable" -verbose up 1
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/url-shortening?sslmode=disable" -verbose down
migratedownlastest:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/url-shortening?sslmode=disable" -verbose down 1
dropdb:
	docker exec -it postgres16 dropdb --username=root url-shortening
sqlc:
	sqlc generate
test:
	go test -v -cover -short ./...
server:
	go run cmd/url-shortening/main.go
.PHONY: createdb create-migrate-file migrateup migrateuplastest migratedown migratedownlastest dropdb sqlc test server