package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client

func Init() {
	client = ConnectMongo()
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func ConnectMongo() *mongo.Client {
	// Create connect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	DBHost := GetEnv("DB_HOST", "localhost")
	DBPort := GetEnv("DB_PORT", "27017")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+DBHost+":"+DBPort))
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}

func GetDB() *mongo.Client {
	return client
}
