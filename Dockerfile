FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum /app/

RUN go install github.com/cosmtrek/air@v1.40.4 &&  \
    go install github.com/99designs/gqlgen@v0.17.21 && \
    go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.2 && \
    go install github.com/kyleconroy/sqlc/cmd/sqlc@v1.16.0 && \
    go mod download

CMD ["air", "-c", ".air.toml"]
