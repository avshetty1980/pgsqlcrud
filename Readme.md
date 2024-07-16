## Run docker posgres container with:

```sh
docker run --name pg-container -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres
```

## Create database with

```sh
docker exec -it pg-container createdb -U postgres gopgtest
```

## Connect to pg with

```sh
docker exec -it pg-container psql -U postgres
```

## Connect to database (default db name = postgres) with

```sh
\c gopgtest
```

## Display tables

```sh
\dt
```

## Create migrations

```sh
sqlc generate
```
