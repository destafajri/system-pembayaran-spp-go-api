package config

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/destafajri/system-pembayaran-spp-go-api/exception"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDatabase(configuration Config) *mongo.Database {
	ctx, cancel := NewMongoContext()
	defer cancel()

	mongoPoolMin, err := strconv.Atoi(configuration.Get("MONGO_POOL_MIN"))
	exception.PanicIfNeeded(err)

	mongoPoolMax, err := strconv.Atoi(configuration.Get("MONGO_POOL_MAX"))
	exception.PanicIfNeeded(err)

	mongoMaxIdleTime, err := strconv.Atoi(configuration.Get("MONGO_MAX_IDLE_TIME_SECOND"))
	exception.PanicIfNeeded(err)

	option := options.Client().
		ApplyURI(configuration.Get("MONGO_URI")).
		SetMinPoolSize(uint64(mongoPoolMin)).
		SetMaxPoolSize(uint64(mongoPoolMax)).
		SetMaxConnIdleTime(time.Duration(mongoMaxIdleTime) * time.Second)

	client, err := mongo.NewClient(option)
	exception.PanicIfNeeded(err)

	err = client.Connect(ctx)
	exception.PanicIfNeeded(err)

	database := client.Database(configuration.Get("MONGO_DATABASE"))
	log.Println("MongoDB Successfully connected!")
	
	return database
}

func NewMongoContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
