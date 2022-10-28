package models

import (
	"backend/db"
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Client struct {
	Name       string  `bson:"name" json:"name" example:"Налсур"`
	Surname    string  `bson:"surname" json:"surname" example:"Мулюков"`
	MiddleName string  `bson:"middle_name,omitempty" json:"middle_name,omitempty" example:"Рустэмович"`
	Address    Address `bson:"address" json:"address"`
}

type Size struct {
	Length int `bson:"length" json:"length" example:"100"`
	Width  int `bson:"width" json:"width" example:"73"`
	Height int `bson:"height" json:"height" example:"42"`
}

type Stage struct {
	Name       string             `bson:"name" json:"name" example:"Принято в отделении связи"`
	Date       time.Time          `bson:"date" json:"date"`
	Postcode   string             `bson:"postcode" json:"postcode" example:"453870"`
	EmployeeID primitive.ObjectID `bson:"employee_id" json:"employee_id"`
}

type Sending struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	OrderID          uuid.UUID          `bson:"order_id" json:"order_id"`
	RegistrationDate time.Time          `bson:"registration_date" json:"registration_date"`
	Sender           Client             `bson:"sender" json:"sender"`
	Receiver         Client             `bson:"receiver" json:"receiver"`
	Type             string             `bson:"type" json:"type" example:"Посылка"`
	Size             Size               `bson:"size" json:"size"`
	Weight           int                `bson:"weight" json:"weight" example:"1000"`
	Stages           []Stage            `bson:"stages" json:"stages"`
	Status           string             `bson:"status" json:"status" example:"Доставлено"`
}

func (s *Sending) InsertExample() error {
	e := new(Employee)
	employees, err := e.FindExample()
	if err != nil {
		return fmt.Errorf("InsertExample: %v", err)
	}

	client := db.GetDB()
	sendingCollection := client.Database("Post").Collection("Sending")

	sending := Sending{
		RegistrationDate: time.Now(),
		OrderID:          uuid.New(),
		Sender: Client{
			Name:       "Налсур",
			Surname:    "Мулюков",
			MiddleName: "Рустэмович",
			Address: Address{
				Postcode:   "453870",
				Region:     "Республика Башкортостан",
				District:   "Мелеузовский район",
				Settlement: "пос. Нугуш",
				Street:     "ул. Ленина",
				Building:   "42",
			},
		},
		Receiver: Client{
			Name:    "Екатерина",
			Surname: "Феминисткова",
			Address: Address{
				Postcode:  "123456",
				Region:    "г. Москва",
				Street:    "ул. Мира",
				Building:  "1",
				Apartment: "1",
			},
		},
		Type: "Посылка",
		Size: Size{
			Length: 100,
			Width:  73,
			Height: 42,
		},
		Weight: 1000,
		Stages: []Stage{
			{
				Name:       "Принято в отделении связи",
				Date:       time.Now(),
				Postcode:   "453870",
				EmployeeID: employees[0].ID,
			},
			{
				Name:       "Вручено адресату",
				Date:       time.Now(),
				Postcode:   "123456",
				EmployeeID: employees[1].ID,
			},
		},
		Status: "Доставлено",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = sendingCollection.InsertOne(ctx, sending)
	if err != nil {
		return fmt.Errorf("InsertExample: %v", err)
	}

	return nil
}

func (s *Sending) FindExample() ([]Sending, error) {
	client := db.GetDB()
	sendingCollection := client.Database("Post").Collection("Sending")

	var sendings []Sending

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := sendingCollection.Find(ctx, bson.M{"registration_date": bson.D{{"$lt", time.Now()}}})
	if err != nil {
		return nil, fmt.Errorf("FindExample: %v", err)
	}
	if err = cursor.All(ctx, &sendings); err != nil {
		return nil, fmt.Errorf("FindExample: %v", err)
	}

	return sendings, nil
}
