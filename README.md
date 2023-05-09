# expenses-gin

DATABASE_URL=postgres://postgres:peem2544@localhost:5432/expenses PORT=:2565 go run server.go

# for window

$env:DATABASE_URL="postgres://postgres:peem2544@localhost:5432/expenses"; $env:PORT=":2565"; go run server.go

# for container

docker-compose up

# data in server.env

PORT=:2565
DATABASE_URL=postgres://postgres:peem2544@db:5432/PostgresDemo

# test

$env:CGO_ENABLED=0; $env:AUTH_TOKEN="asd"; go test --tags=integration ./...

$env:CGO_ENABLED=0; $env:AUTH_TOKEN="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODA5NjI2OTF9.iZtoIOFd_7qLiTkfyMNwJbZpY_YxCRAOyVCOlPizEiE"; go test --tags=integration ./...

# test docker-compose sandbox

$env:CGO_ENABLED=0; docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit --exit-code-from it_tests

! not work because i don't know how to set AUTH_TOKEN
