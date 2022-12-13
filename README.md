# gqlgen template

## 構成

- Go : 1.19
- MySQL : 8.0.31

## ローカル開発環境設定

- cp .env.example .env
- docker compose up でコンテナの立ち上げ
  - ポート番号
    - API : 8082
    - MySQL : 3308
- 下記コマンド で Go アプリケーションコンテナへ入る

  `docker compose exec api bash`
