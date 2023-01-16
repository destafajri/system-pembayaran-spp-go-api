package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/destafajri/system-pembayaran-spp-go-api/exception"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	postgresUp()
	mysqlUp()
}

func postgresUp() {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	exception.PanicIfNeeded(err)
	
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	exception.PanicIfNeeded(err)

    m, err := migrate.NewWithDatabaseInstance(
        "file://migrations/postgres",
        "postgres", driver)
	exception.PanicIfNeeded(err)
	m.Up()
	
	log.Println("Migration Postgres Successfully!")
}

func mysqlUp() {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	
	database, err := sql.Open("mysql", os.Getenv("MYSQL_URL"))
	exception.PanicIfNeeded(err)

	driver, err := mysql.WithInstance(database, &mysql.Config{})
	exception.PanicIfNeeded(err)

    m, err := migrate.NewWithDatabaseInstance(
        "file://migrations/mysql",
        "mysql", driver)
	exception.PanicIfNeeded(err)
	m.Up()

	log.Println("Migration MySQL Successfully!")
}