package models

import (
	"backend/db"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Employee struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name        string             `bson:"name" json:"name" example:"Райан"`
	Surname     string             `bson:"surname" json:"surname" example:"Гослинг"`
	MiddleName  string             `bson:"middle_name,omitempty" json:"middle_name,omitempty" example:"Томасович"`
	Gender      string             `bson:"gender" json:"gender" example:"М"`
	BirthDate   time.Time          `bson:"birth_date" json:"birth_date"`
	Position    string             `bson:"position" json:"position" example:"Водитель"`
	PhoneNumber string             `bson:"phone_number" json:"phone_number" example:"88005553535"`
}

func (e *Employee) FindExample() ([]Employee, error) {
	client := db.GetDB()
	employeesCollection := client.Database("Post").Collection("Employee")

	var employees []Employee

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := employeesCollection.Find(ctx, bson.M{"birth_date": bson.D{{"$lt", time.Now()}}})
	if err != nil {
		return nil, fmt.Errorf("FindExample: %v", err)
	}
	if err = cursor.All(ctx, &employees); err != nil {
		return nil, fmt.Errorf("FindExample: %v", err)
	}

	return employees, nil
}
