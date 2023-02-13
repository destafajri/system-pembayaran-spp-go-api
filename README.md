# Go-Lang Fiber Project

System Pembayaran SPP API

- [API Documentation](https://whimsical.com/erd-table-api-map-Umr1mzU2SY3jAqJyDa7KMX)
- [API Contract](https://www.notion.so/e785eb0dfdc245659df81b6e91bf40a5?v=2176e61ec8d744d4af0b35e613895ad9)
- [Postman Collection](https://documenter.getpostman.com/view/22138766/2s935sohZ6)
- [Github Repository](https://github.com/destafajri/system-pembayaran-spp-go-api)

## How to use
- Please clone or download this repository.
```
git clone https://github.com/destafajri/system-pembayaran-spp-go-api.git
```
- Prepare postgres database, for this project i use [SUPABASE](https://supabase.com/)
- add .env file to setup your database connection and configuration
- if you want to use docker, you can type
```
docker-compose up
```
OR
```
docker-compose up -d
```
- `migrator tools` from [golang migrate](https://github.com/golang-migrate/migrate)
- to install the migration tools, you can type
```
go install -tags 'postgres,mysql,mongodb’ github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```
- to create, drop, and alter table migration you can read this [link](https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md)
- to run your database schema
```
make migrate-up
```
- to run the golang server
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
- Database : Postgre(Supabase)

## Architecture

Controller -> Service -> Repository

## Project Structure example
    .
    ├── config
    ├── exception
    ├── helper
    │   ├── jwts
    │   ├── password
    │   └── timeutil
    ├── internal
    │   ├── controller
    │   │   ├── bayar
    │   │   ├── guru
    │   │   ├── kelas
    │   │   ├── siswa
    │   │   ├── spp
    │   │   └── user
    │   ├── domain
    │   │   ├── entity
    │   │   └── model
    │   ├── middlewares
    │   ├── repository
    │   │   └── database
    │   │       ├── mongo
    │   │       ├── mysql
    │   │       └── postgres
    │   │           ├── bayar_repository
    │   │           ├── guru_repository
    │   │           ├── kelas_repository
    │   │           ├── siswa_repository
    │   │           ├── spp_repository
    │   │           └── user_repository
    │   ├── service
    │   │   ├── bayar_service
    │   │   ├── guru_service
    │   │   ├── kelas_service
    │   │   ├── siswa_service
    │   │   ├── spp_service
    │   │   └── user_service
    │   └── validations
    ├── meta
    │   └── param
    ├── migrations
    │   ├── cmd
    │   │   ├── down
    │   │   └── up
    │   ├── mysql
    │   └── postgres
    ├── responses
    └── vendor

## Addition 

- Entity is representing database table
- Model is representing payload and response
- Helper is representing anything what you need to help your coding process
