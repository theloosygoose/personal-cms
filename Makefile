build:
	@go build -o ./bin/main ./cmd/cms/main.go

migrate_up:
	@migrate -path migrations/ -database "sqlite3://stores/main" up
