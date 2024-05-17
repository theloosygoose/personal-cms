run:
	@go run ./cmd/cms/main.go
build:
	@go build -o ./bin/main ./cmd/cms/main.go

migrate_up:
	@migrate -path database/migrations/ -database "sqlite3://stores/personal-cms.db" up
