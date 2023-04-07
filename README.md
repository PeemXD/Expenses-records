# expenses-gin

# DATABASE_URL=postgres://postgres:peem2544@localhost:5432/expenses PORT=:2565 go run server.go

# for window

# $env:DATABASE_URL="postgres://postgres:peem2544@localhost:5432/expenses"; $env:PORT=":2565"; go run server.go

# for container

# PORT=2565 DATABASE_URL=postgres://postgres:peem2544@db:5432/PostgresDemo docker-compose -f docker-compose.yml up

# $env:PORT=2565; $env:DATABASE_URL="postgres://postgres:peem2544@db:5432/PostgresDemo"; docker-compose -f docker-compose.yml up

# docker-compose -f docker-compose.yml up

# data in server.env

PORT=:2565
DATABASE_URL=postgres://postgres:peem2544@db:5432/PostgresDemo
