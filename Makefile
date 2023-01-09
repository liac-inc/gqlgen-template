# マイグレーションUp実行
include .env
.EXPORT_ALL_VARIABLES:
migrate-up:
	migrate -database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" \
	-path db/migration up

# マイグレーションDown実行
include .env
.EXPORT_ALL_VARIABLES:
migrate-down:
	migrate -database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" \
	-path db/migration down -all

# マイグレーションClear実行
include .env
.EXPORT_ALL_VARIABLES:
migrate-clear:
	migrate -database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" \
	-path db/migration force 1

# マイグレーションファイルの作成(tableName = テーブル名(複数形)として引数に渡す)
.PHONY: create-migrate-file
create-migrate-file:
	migrate create -ext sql -dir db/migration -seq ${tableName}

# Seedデータの投入
.PHONY: data-seeding
data-seeding:
	go run src/cmd/script/seeder.go

# SQLファイルからアプリケーションコードの生成
.PHONY: sqlc-generate
sqlc-generate:
	sqlc generate

# GraphQLのSchemaファイルからアプリケーションコードを生成
.PHONY: gqlgen-generate
gqlgen-generate:
	go run -mod=mod github.com/99designs/gqlgen generate

.PHONY: wire-gen
wire-gen:
	wire src/registry/wire.go


