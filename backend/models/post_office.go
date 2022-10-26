package models

import (
	"backend/db"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

type Employee struct {
	Name        string    `bson:"name" json:"name" example:"Райан"`
	Surname     string    `bson:"surname" json:"surname" example:"Гослинг"`
	MiddleName  string    `bson:"middle_name,omitempty" json:"middle_name,omitempty" example:"Томасович"`
	Gender      string    `bson:"gender" json:"gender" example:"М"`
	BirthDate   time.Time `bson:"birth_date" json:"birth_date"`
	Position    string    `bson:"position" json:"position" example:"Водитель"`
	PhoneNumber string    `bson:"phone_number" json:"phone_number" example:"88005553535"`
}

type PostOffice struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Type      string             `bson:"type" json:"type" example:"Отделение связи"`
	Address   Address            `bson:"address" json:"address"`
	Employees []Employee         `bson:"employees" json:"employees"`
}

func (po *PostOffice) InsertExample() error {
	client := db.GetDB()
	sendingsCollection := client.Database("Post").Collection("PostOffice")

	postOffice := PostOffice{
		Type: "Отделение связи",
		Address: Address{
			Postcode:   "453870",
			Region:     "Республика Башкортостан",
			District:   "Мелеузовский район",
			Settlement: "пос. Нугуш",
			Street:     "ул. Ленина",
			Building:   "42",
		},
		Employees: []Employee{
			{
				Surname:     "Гослинг",
				Name:        "Райан",
				MiddleName:  "Томасович",
				Gender:      "М",
				BirthDate:   time.Now().Add(-time.Duration(200000) * time.Hour),
				Position:    "Водитель",
				PhoneNumber: "88005553535",
			},
			{
				Surname:     "Секретова",
				Name:        "Секрета",
				MiddleName:  "Секретовна",
				Gender:      "Ж",
				BirthDate:   time.Now().Add(-time.Duration(500000) * time.Hour),
				Position:    "Почтальон",
				PhoneNumber: "82233222233",
			},
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := sendingsCollection.InsertOne(ctx, postOffice)
	if err != nil {
		return fmt.Errorf("InsertExample: %v", err)
	}

	return nil
}

func (po *PostOffice) FindExample() ([]PostOffice, error) {
	client := db.GetDB()
	sendingsCollection := client.Database("Post").Collection("PostOffice")

	var postOffices []PostOffice

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := sendingsCollection.Find(ctx, bson.D{{"type", "Отделение связи"}})
	if err != nil {
		return nil, fmt.Errorf("FindExample: %v", err)
	}
	if err = cursor.All(ctx, &postOffices); err != nil {
		return nil, fmt.Errorf("FindExample: %v", err)
	}

	for _, element := range postOffices {
		json, _ := json.MarshalIndent(element, "", "\t")
		fmt.Println(string(json))
	}

	return postOffices, nil
}
