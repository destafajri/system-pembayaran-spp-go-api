# Go-Lang Fiber Project

System Pembayaran SPP API

- [API Documentation](https://whimsical.com/erd-table-api-map-Umr1mzU2SY3jAqJyDa7KMX)
- [API Contract](https://www.notion.so/e785eb0dfdc245659df81b6e91bf40a5?v=2176e61ec8d744d4af0b35e613895ad9)

## How to use
- Please clone or download this repository.
```
git clone https://github.com/destafajri/system-pembayaran-spp-go-api.git
```
- Prepare postgres database, for this project i use [SUPABASE](https://supabase.com/)
- if you want to use docker, you can type
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
or
```
go run main.go
```

## Framework

- Web : GoFiber
- Validation : Go-Ozzo
- Configuration : GoDotEnv
- Database : MongoDB, Postgre(Supabase), MySQL

## Architecture

Controller -> Service -> Repository

## Project Structure example
    .
    ├── Dockerfile
    ├── LICENSE
    ├── Makefile
    ├── README.md
    ├── config
    │   ├── config.go
    │   ├── fiber.go
    │   ├── mongo.go
    │   ├── mysql.go
    │   └── postgres.go
    ├── docker-compose.yml
    ├── exception
    │   ├── error.go
    │   ├── error_handler.go
    │   └── validation_error.go
    ├── go.mod
    ├── go.sum
    ├── helper
    │   ├── generate_jwt.go
    │   └── generate_password.go
    ├── internal
    │   ├── controller
    │   │   ├── controller.go
    │   │   ├── product
    │   │   │   ├── controller_test.go
    │   │   │   ├── product_controller.go
    │   │   │   ├── product_controller_test.go
    │   │   │   └── product_router.go
    │   │   └── user
    │   │       ├── user_controller.go
    │   │       └── user_router.go
    │   ├── entity
    │   │   ├── product.go
    │   │   └── user.go
    │   ├── middlewares
    │   │   ├── JWTMiddleware.go
    │   │   ├── config.go
    │   │   ├── crypto.go
    │   │   ├── jwks.go
    │   │   └── jwt_claims.go
    │   ├── model
    │   │   ├── product_model.go
    │   │   └── user_model.go
    │   ├── repository
    │   │   ├── product_repository
    │   │   │   └── product_repository_impl.go
    │   │   ├── product_repository.go
    │   │   ├── user_repository
    │   │   │   └── user_repository_impl.go
    │   │   └── user_repository.go
    │   ├── service
    │   │   ├── product_service
    │   │   │   └── product_service_impl.go
    │   │   ├── product_service.go
    │   │   ├── user_service.go
    │   │   └── user_service_impl
    │   │       └── user_service_impl.go
    │   └── validation
    │       ├── product_validation.go
    │       └── user_validation.go
    ├── main.go
    ├── migrations
    │   ├── cmd
    │   │   ├── down
    │   │   │   └── main.go
    │   │   └── up
    │   │       └── main.go
    │   ├── mysql
    │   │   ├── 000001_create_sample_tables.up.sql
    │   │   ├── 000001_down_sample_tables.down.sql
    │   │   ├── 000002_create_users_tables.up.sql
    │   │   └── 000002_down_users_table.down.sql
    │   └── postgres
    │       ├── 000001_create_sample_tables.up.sql
    │       ├── 000001_down_sample_tables.down.sql
    │       ├── 000002_create_users_tables.up.sql
    │       └── 000002_down_users_table.down.sql
    ├── responses
    │   └── web_response.go
    ├── test.http
    └── vendor

## Addition 

- Entity is representing database table
- Model is representing payload and response
- Helper is representing anything what you need to help your coding process
