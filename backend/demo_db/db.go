package demo_db

import (
	"backend/db"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func Insert() {
	client := db.GetDB()
	collection := client.Database("HelloWorld").Collection("HelloWorld")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection.InsertOne(ctx, bson.D{{"HelloWorld", "Yes"}})
}
