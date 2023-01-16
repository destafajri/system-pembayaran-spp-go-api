# Go-Lang Fiber Project

Sample Go-Lang Fiber Project Structure

## How to use
- Please clone or download this repository.
- Prepare postgres database, or use docker, you can type
```
docker-compose up
```
OR
```
docker-compose up -d
```
- add .env file to setup your database connection
```
MONGO_URI=mongodb://mongo:mongo@localhost:27017
MONGO_DATABASE=golang_test
MONGO_POOL_MIN=10
MONGO_POOL_MAX=100
MONGO_MAX_IDLE_TIME_SECOND=60

POSTGRES_URL="user=postgres password=[your-password] host=[your-host] port=5432 dbname=postgres"
POSTGRES_USER=postgres
POSTGRES_PASSWORD=[your-password]
POSTGRES_HOST=[your-host]
POSTGRES_PORT=5432
POSTGRES_DB=postgres

MYSQL_URL="root:secret@tcp(localhost:3306)/sample?parseTime=true"

KEY_JWT="nafonFajriSecretKeyJWTdkdjfnfja"
```
- `installing migrator tools` download from [golang migrate](https://github.com/golang-migrate/migrate) in release page
- run
```
make migrate-up
```
- run the golang server
```
make run
```

## Framework

- Web : GoFiber
- Validation : Go-Ozzo
- Configuration : GoDotEnv
- Database : MongoDB, Postgre(Supabase), MySQL

## Architecture

Controller -> Service -> Repository

## Addition 

- Entity is representing database table
- Model is representing payload and response
- Helper is representing anything what you need to help your coding process
