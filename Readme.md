Go Install

1. go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
2. go install github.com/pressly/goose/v3/cmd/goose@latest
3. go get github.com/google/uuid

Migration
1. create folder sql/schema
2. create new file 001_users.sql
    ```
    -- +goose Up

    CREATE TABLE users (
        id UUID PRIMARY KEY,
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL,
        name TEXT NOT NULL
    );

    -- +goose Down
    DROP TABLE users;
    ```
3. cd sql/schema
4. goose postgres postgres://postgres:postgres@localhost:5432/go-restapi up

Generate Model using Sqlc 
1. create folder in sql/queries
2. create sql file users.sql
```
    -- name: CreateUser :one
    INSERT INTO users (id, created_at, updated_at, name) VALUES ($1, $2, $3, $4)
    RETURNING *;
```
3. run command ```sqlc generate```

Running

1. module init : go mod init github.com/tedyfd/go-restapi
2. build to binary files : go build .
3. running : go run .
4. get module : go get github.com/joho/godotenv
5. download module : go mod vendor
