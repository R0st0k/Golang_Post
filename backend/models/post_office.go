package models

import (
	"backend/db"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Address struct {
	Postcode   string `bson:"postcode" json:"postcode" example:"453870"`
	Region     string `bson:"region,omitempty" json:"region,omitempty" example:"Республика Башкортостан"`
	District   string `bson:"district,omitempty" json:"district,omitempty" example:"Мелеузовский район"`
	Settlement string `bson:"settlement" json:"settlement" example:"пос. Нугуш"`
	Street     string `bson:"street" json:"street" example:"ул. Ленина"`
	Building   string `bson:"building" json:"building" example:"42"`
	Apartment  string `bson:"apartment,omitempty" json:"apartment,omitempty" example:"1"`
}

type PostOffice struct {
	ID        primitive.ObjectID   `bson:"_id,omitempty" json:"_id,omitempty"`
	Type      string               `bson:"type" json:"type" example:"Отделение связи"`
	Address   Address              `bson:"address" json:"address"`
	Employees []primitive.ObjectID `bson:"employees" json:"employees"`
}

func (po *PostOffice) GetSettlementByPostcode() (map[string]string, error) {
	client := db.GetDB()
	postOfficeCollection := client.Database("post").Collection("postOffices")

	var postOffices []PostOffice

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	projection := bson.D{
		{"address.postcode", 1},
		{"address.settlement", 1},
		{"_id", 0}}
	opts := options.Find().SetProjection(projection)
	cursor, err := postOfficeCollection.Find(ctx, bson.D{{}}, opts)
	if err != nil {
		return nil, fmt.Errorf("GetPostOffices: %v", err)
	}
	if err = cursor.All(ctx, &postOffices); err != nil {
		return nil, fmt.Errorf("GetPostOffices: %v", err)
	}

	settlementByPostcode := make(map[string]string)

	for i := range postOffices {
		settlementByPostcode[postOffices[i].Address.Postcode] = postOffices[i].Address.Settlement
	}

	return settlementByPostcode, nil
}

func (po *PostOffice) GetPostcodesBySettlement(types map[string]interface{}) (map[string][]string, error) {
	client := db.GetDB()
	postOfficeCollection := client.Database("post").Collection("postOffices")

	matchPipeline := mongo.Pipeline{}
	projectPipeline := mongo.Pipeline{
		bson.D{{
			"$project", bson.D{
				{"address.postcode", 1},
				{"address.settlement", 1},
				{"_id", 0}},
		}},
	}

	matchStringArray(types, "type", "type", &matchPipeline)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	findPipeline := mongo.Pipeline{}
	findPipeline = append(findPipeline, matchPipeline...)
	findPipeline = append(findPipeline, projectPipeline...)

	var postOffices []PostOffice

	cursor, err := postOfficeCollection.Aggregate(ctx, findPipeline)

	if err != nil {
		return nil, fmt.Errorf("GetPostOffices: %v", err)
	}
	if err = cursor.All(ctx, &postOffices); err != nil {
		return nil, fmt.Errorf("GetPostOffices: %v", err)
	}

	postcodesBySettlement := make(map[string][]string)

	for i := range postOffices {
		if _, inMap := postcodesBySettlement[postOffices[i].Address.Settlement]; inMap {
			postcodesBySettlement[postOffices[i].Address.Settlement] = append(postcodesBySettlement[postOffices[i].Address.Settlement], postOffices[i].Address.Postcode)
		} else {
			newSettlement := []string{postOffices[i].Address.Postcode}
			postcodesBySettlement[postOffices[i].Address.Settlement] = newSettlement
		}
	}

	return postcodesBySettlement, nil
}

func (po *PostOffice) GetPostWorkerByPostcode(postcode string) (primitive.ObjectID, error) {
	client := db.GetDB()
	employeesCollection := client.Database("post").Collection("postOffices")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var employees []Employee

	matchPostcodeStage := bson.D{{
		"$match", bson.D{{
			"address.postcode", postcode,
		}},
	}}
	firstUnwindStage := bson.D{{
		"$unwind", "$employees",
	}}
	lookupStage := bson.D{{
		"$lookup", bson.D{{
			"from", "employees",
		},
			{
				"localField", "employees",
			},
			{
				"foreignField", "_id",
			},
			{
				"as", "employee",
			}},
	}}
	matchPositionStage := bson.D{{
		"$match", bson.D{{
			"employee.position", "Сотрудник отделения связи",
		}},
	}}
	secondUnwindStage := bson.D{{
		"$unwind", "$employee",
	}}
	replaceWithStage := bson.D{{
		"$replaceWith", "$employee",
	}}

	cursor, err := employeesCollection.Aggregate(ctx, mongo.Pipeline{matchPostcodeStage, firstUnwindStage, lookupStage, matchPositionStage, secondUnwindStage, replaceWithStage})
	if err != nil {
		return primitive.NewObjectID(), fmt.Errorf("FindEmployees: %v", err)
	}
	if err = cursor.All(ctx, &employees); err != nil {
		return primitive.NewObjectID(), fmt.Errorf("FindEmployees: %v", err)
	}

	if len(employees) == 0 {
		return primitive.NewObjectID(), fmt.Errorf("FindEmployees: There are no post worker in sender's post office to start stage")
	}

	return employees[0].ID, nil
}
