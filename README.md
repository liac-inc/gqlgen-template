# gqlgen template

## 構成

- Go : 1.19
- PostgreSQL : 15.1

## ローカル開発環境設定

- cp .env.example .env
- docker compose up でコンテナの立ち上げ
  - ポート番号
    - API : 4000
    - PostgreSQL : 5432
- 下記コマンド で Go アプリケーションコンテナへ入る

  `docker compose exec api bash`
