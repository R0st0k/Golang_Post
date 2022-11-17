package models

import (
	"backend/db"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Address struct {
	Postcode   int64  `bson:"postcode" json:"postcode" example:"453870"`
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

func (po *PostOffice) InsertExample() error {
	e := new(Employee)
	employees, err := e.FindExample()
	if err != nil {
		return fmt.Errorf("InsertExample: %v", err)
	}

	client := db.GetDB()
	postOfficeCollection := client.Database("Post").Collection("PostOffice")

	postOffice := PostOffice{
		Type: "Отделение связи",
		Address: Address{
			Postcode:   453870,
			Region:     "Республика Башкортостан",
			District:   "Мелеузовский район",
			Settlement: "пос. Нугуш",
			Street:     "ул. Ленина",
			Building:   "42",
		},
		Employees: []primitive.ObjectID{employees[0].ID, employees[1].ID},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = postOfficeCollection.InsertOne(ctx, postOffice)
	if err != nil {
		return fmt.Errorf("InsertExample: %v", err)
	}

	return nil
}

func (po *PostOffice) FindExample() ([]PostOffice, error) {
	client := db.GetDB()
	postOfficeCollection := client.Database("Post").Collection("PostOffice")

	var postOffices []PostOffice

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := postOfficeCollection.Find(ctx, bson.D{{"type", "Отделение связи"}})
	if err != nil {
		return nil, fmt.Errorf("FindExample: %v", err)
	}
	if err = cursor.All(ctx, &postOffices); err != nil {
		return nil, fmt.Errorf("FindExample: %v", err)
	}

	return postOffices, nil
}

func (po *PostOffice) FindCityByPostcode() (map[int64]string, error) {
	client := db.GetDB()
	postOfficeCollection := client.Database("Post").Collection("PostOffice")

	var postOffices []PostOffice

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	projection := bson.D{{"address.postcode", 1}, {"address.settlement", 1}, {"_id", 0}}
	opts := options.Find().SetProjection(projection)
	cursor, err := postOfficeCollection.Find(ctx, bson.D{{}}, opts)
	if err != nil {
		return nil, fmt.Errorf("FindExample: %v", err)
	}
	if err = cursor.All(ctx, &postOffices); err != nil {
		return nil, fmt.Errorf("FindExample: %v", err)
	}

	cityByCode := make(map[int64]string)

	for i := range postOffices {
		cityByCode[postOffices[i].Address.Postcode] = postOffices[i].Address.Settlement
	}

	return cityByCode, nil
}

func (po *PostOffice) FindPostcodesByCity() (map[string][]int64, error) {
	client := db.GetDB()
	postOfficeCollection := client.Database("Post").Collection("PostOffice")

	var postOffices []PostOffice

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	projection := bson.D{{"address.postcode", 1}, {"address.settlement", 1}, {"_id", 0}}
	opts := options.Find().SetProjection(projection)
	cursor, err := postOfficeCollection.Find(ctx, bson.D{{}}, opts)
	if err != nil {
		return nil, fmt.Errorf("FindExample: %v", err)
	}
	if err = cursor.All(ctx, &postOffices); err != nil {
		return nil, fmt.Errorf("FindExample: %v", err)
	}

	postcodesByCity := make(map[string][]int64)

	for i := range postOffices {
		if _, inMap := postcodesByCity[postOffices[i].Address.Settlement]; inMap {
			postcodesByCity[postOffices[i].Address.Settlement] = append(postcodesByCity[postOffices[i].Address.Settlement], postOffices[i].Address.Postcode)
		} else {
			new_city := []int64{postOffices[i].Address.Postcode}
			postcodesByCity[postOffices[i].Address.Settlement] = new_city
		}
	}

	return postcodesByCity, nil
}
