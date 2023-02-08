1. Setup the postgres db
`docker run --name go-postgres -p 5432:5432 -e POSTGRES_PASSWORD=secret -d postgres`