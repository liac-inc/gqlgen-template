# マイグレーションUp実行
include .env
.EXPORT_ALL_VARIABLES:
migrate-up:
	migrate -database "mysql://${MYSQL_USERNAME}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DATABASE}?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=True" \
	-path db/migrations up

# マイグレーションDown実行
include .env
.EXPORT_ALL_VARIABLES:
migrate-down:
	migrate -database "mysql://${MYSQL_USERNAME}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DATABASE}?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=True" \
	-path db/migrations down -all

# マイグレーションファイルのクリア
include .env
.EXPORT_ALL_VARIABLES:
migrate-clear:
	migrate -database "mysql://${MYSQL_USERNAME}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DATABASE}?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=True" \
	-path db/migrations force 1

# マイグレーションファイルの作成(tableName = テーブル名(複数形)として引数に渡す)
.PHONY: create-migrate-file
create-migrate-file:
	migrate create -ext sql -dir db/migrations -seq ${tableName}

# Seedデータの投入
.PHONY: data-seeding
data-seeding:
	go run app/cmd/scripts/seeder.go

# SQLファイルからアプリケーションコードの生成
.PHONY: sqlc-generate
sqlc-generate:
	sqlc generate

# GraphQLのSchemaファイルからアプリケーションコードを生成
.PHONY: gqlgen-generate
gqlgen-generate:
	go run -mod=mod github.com/99designs/gqlgen generate

