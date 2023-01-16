package config

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/destafajri/system-pembayaran-spp-go-api/exception"
	_ "github.com/go-sql-driver/mysql"

	"github.com/golang-migrate/migrate/v4"
    "github.com/golang-migrate/migrate/v4/database/mysql"
    _ "github.com/golang-migrate/migrate/v4/source/file"
)

func NewMySQLDatabase(configuration Config) *sql.DB {
	_, cancel := NewMySQLContext()
	defer cancel()

	database, err := sql.Open("mysql", configuration.Get("MYSQL_URL"))
	if err != nil {
		panic(err)
	}

	exception.PanicIfNeeded(err)

	err = database.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("MySQL Successfully connected!")

	driver, err := mysql.WithInstance(database, &mysql.Config{})
	exception.PanicIfNeeded(err)
    m, err := migrate.NewWithDatabaseInstance(
        "file://migrations/mysql",
        "mysql", driver)
	exception.PanicIfNeeded(err)
    m.Up()

	log.Println("Migration MySQL Successfully!")
	return database
}

func NewMySQLContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}