package models

import (
	"backend/db"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/mongodb/mongo-tools/mongoexport"
	"github.com/mongodb/mongo-tools/mongoimport"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

type Client struct {
	Name       string  `bson:"name" json:"name" example:"Налсур"`
	Surname    string  `bson:"surname" json:"surname" example:"Мулюков"`
	MiddleName string  `bson:"middle_name,omitempty" json:"middle_name,omitempty" example:"Рустэмович"`
	Address    Address `bson:"address" json:"address"`
}

type Size struct {
	Length int64 `bson:"length" json:"length" example:"100"`
	Width  int64 `bson:"width" json:"width" example:"73"`
	Height int64 `bson:"height" json:"height" example:"42"`
}

type Stage struct {
	Name       string             `bson:"name" json:"name" example:"Принято в отделении связи"`
	Timestamp  time.Time          `bson:"timestamp" json:"timestamp"`
	Postcode   string             `bson:"postcode" json:"postcode" example:"453870"`
	EmployeeID primitive.ObjectID `bson:"employee_id" json:"employee_id"`
}

type Sending struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	OrderID          string             `bson:"order_id" json:"order_id"`
	RegistrationDate time.Time          `bson:"registration_date" json:"registration_date"`
	Sender           Client             `bson:"sender" json:"sender"`
	Receiver         Client             `bson:"receiver" json:"receiver"`
	Type             string             `bson:"type" json:"type" example:"Посылка"`
	Size             Size               `bson:"size" json:"size"`
	Weight           int64              `bson:"weight" json:"weight" example:"1000"`
	Stages           []Stage            `bson:"stages" json:"stages"`
	Status           string             `bson:"status" json:"status" example:"Доставлено"`
}

func (s *Sending) GetSendingByOrderID(orderID uuid.UUID) (Sending, error) {
	client := db.GetDB()
	sendingCollection := client.Database("post").Collection("sendings")

	var sending Sending

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	projection := bson.D{
		{"type", 1},
		{"status", 1},
		{"stages", 1},
		{"_id", 0}}
	opts := options.FindOne().SetProjection(projection)
	err := sendingCollection.FindOne(ctx, bson.D{{"order_id", orderID.String()}}, opts).Decode(&sending)
	if err != nil {
		return sending, fmt.Errorf("FindSending: %v", err)
	}

	return sending, nil
}

func (s *Sending) InsertNewSending(newSendingOptions map[string]interface{}) (uuid.UUID, error) {
	po := new(PostOffice)
	employeeID, err := po.GetPostWorkerByPostcode(newSendingOptions["Sender"].(map[string]string)["Postcode"])
	if err != nil {
		return uuid.New(), err
	}

	client := db.GetDB()
	sendingCollection := client.Database("post").Collection("sendings")

	newSending := Sending{
		RegistrationDate: time.Now(),
		OrderID:          uuid.NewString(),
		Sender: Client{
			Name:       newSendingOptions["Sender"].(map[string]string)["Name"],
			Surname:    newSendingOptions["Sender"].(map[string]string)["Surname"],
			MiddleName: newSendingOptions["Sender"].(map[string]string)["MiddleName"],
			Address: Address{
				Postcode:   newSendingOptions["Sender"].(map[string]string)["Postcode"],
				Region:     newSendingOptions["Sender"].(map[string]string)["Region"],
				District:   newSendingOptions["Sender"].(map[string]string)["District"],
				Settlement: newSendingOptions["Sender"].(map[string]string)["Settlement"],
				Street:     newSendingOptions["Sender"].(map[string]string)["Street"],
				Building:   newSendingOptions["Sender"].(map[string]string)["Building"],
				Apartment:  newSendingOptions["Sender"].(map[string]string)["Apartment"],
			},
		},
		Receiver: Client{
			Name:       newSendingOptions["Receiver"].(map[string]string)["Name"],
			Surname:    newSendingOptions["Receiver"].(map[string]string)["Surname"],
			MiddleName: newSendingOptions["Receiver"].(map[string]string)["MiddleName"],
			Address: Address{
				Postcode:   newSendingOptions["Receiver"].(map[string]string)["Postcode"],
				Region:     newSendingOptions["Receiver"].(map[string]string)["Region"],
				District:   newSendingOptions["Receiver"].(map[string]string)["District"],
				Settlement: newSendingOptions["Receiver"].(map[string]string)["Settlement"],
				Street:     newSendingOptions["Receiver"].(map[string]string)["Street"],
				Building:   newSendingOptions["Receiver"].(map[string]string)["Building"],
				Apartment:  newSendingOptions["Receiver"].(map[string]string)["Apartment"],
			},
		},
		Type: newSendingOptions["Type"].(string),
		Size: Size{
			Length: newSendingOptions["Length"].(int64),
			Width:  newSendingOptions["Width"].(int64),
			Height: newSendingOptions["Height"].(int64),
		},
		Weight: newSendingOptions["Weight"].(int64),
		Stages: []Stage{
			{
				Name:       "Принято в отделении связи",
				Timestamp:  time.Now(),
				Postcode:   newSendingOptions["Sender"].(map[string]string)["Postcode"],
				EmployeeID: employeeID,
			},
		},
		Status: "В пути",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = sendingCollection.InsertOne(ctx, newSending)
	if err != nil {
		return uuid.New(), fmt.Errorf("InsertSending: %v", err)
	}

	return uuid.MustParse(newSending.OrderID), nil
}

func (s *Sending) FilterSending(sendingFilter map[string]interface{}) (int64, []Sending, error) {
	client := db.GetDB()
	sendingCollection := client.Database("post").Collection("sendings")

	var err error

	pagePipeline := getPagePipeline(sendingFilter, "page", "elems")
	matchPipeline := mongo.Pipeline{}
	sortPipeline := mongo.Pipeline{}

	err = matchRegex(sendingFilter, "order_id", "order_id", "i", &matchPipeline)
	if err != nil {
		return 0, nil, err
	}
	err = matchArray(sendingFilter, "type", "type", "string", &matchPipeline)
	if err != nil {
		return 0, nil, err
	}
	err = matchArray(sendingFilter, "status", "status", "string", &matchPipeline)
	if err != nil {
		return 0, nil, err
	}
	err = matchWithCompare(sendingFilter, "date_start", "registration_date", "time", "$gte", &matchPipeline)
	if err != nil {
		return 0, nil, err
	}
	err = matchWithCompare(sendingFilter, "date_finish", "registration_date", "time", "$lte", &matchPipeline)
	if err != nil {
		return 0, nil, err
	}
	err = matchRegex(sendingFilter, "sender_settlement", "sender.address.settlement", "i", &matchPipeline)
	if err != nil {
		return 0, nil, err
	}
	err = matchRegex(sendingFilter, "receiver_settlement", "receiver.address.settlement", "i", &matchPipeline)
	if err != nil {
		return 0, nil, err
	}
	err = matchWithCompare(sendingFilter, "length_min", "size.length", "int64", "$gte", &matchPipeline)
	if err != nil {
		return 0, nil, err
	}
	err = matchWithCompare(sendingFilter, "length_max", "size.length", "int64", "$lte", &matchPipeline)
	if err != nil {
		return 0, nil, err
	}
	err = matchWithCompare(sendingFilter, "width_min", "size.width", "int64", "$gte", &matchPipeline)
	if err != nil {
		return 0, nil, err
	}
	err = matchWithCompare(sendingFilter, "width_max", "size.width", "int64", "$lte", &matchPipeline)
	if err != nil {
		return 0, nil, err
	}
	err = matchWithCompare(sendingFilter, "height_min", "size.height", "int64", "$gte", &matchPipeline)
	if err != nil {
		return 0, nil, err
	}
	err = matchWithCompare(sendingFilter, "height_max", "size.height", "int64", "$lte", &matchPipeline)
	if err != nil {
		return 0, nil, err
	}
	err = matchWithCompare(sendingFilter, "weight_min", "weight", "int64", "$gte", &matchPipeline)
	if err != nil {
		return 0, nil, err
	}
	err = matchWithCompare(sendingFilter, "weight_max", "weight", "int64", "$lte", &matchPipeline)
	if err != nil {
		return 0, nil, err
	}

	{
		sortTypeInMap, okType := sendingFilter["sort_type"]
		sortFieldInMap, okField := sendingFilter["sort_field"]
		if okType && okField {
			sortType := 0
			switch sortTypeInMap.(string) {
			case "asc":
				sortType = 1
			case "desc":
				sortType = -1
			}
			sortField := ""
			switch sortFieldInMap.(string) {
			case "order_id":
				sortField = "order_id"
			case "type":
				sortField = "type"
			case "status":
				sortField = "status"
			case "date":
				sortField = "registration_date"
			case "sender_settlement":
				sortField = "sender.address.settlement"
			case "receiver_settlement":
				sortField = "receiver.address.settlement"
			case "weight":
				sortField = "weight"
			}
			sort := bson.D{{
				"$sort", bson.D{{
					sortField, sortType,
				},
					{
						"_id", 1,
					}},
			}}
			sortPipeline = append(sortPipeline, sort)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	countPipeline := mongo.Pipeline{}
	countStage := bson.D{{"$count", "order_id"}}
	countPipeline = append(countPipeline, matchPipeline...)
	countPipeline = append(countPipeline, countStage)

	cursor, err := sendingCollection.Aggregate(ctx, countPipeline)
	if err != nil {
		return 0, nil, fmt.Errorf("FindSendings: %v", err)
	}
	var count []bson.D
	if err = cursor.All(ctx, &count); err != nil {
		return 0, nil, fmt.Errorf("FindSendings: %v", err)
	}
	total := int64(0)
	if len(count) > 0 {
		total = int64(count[0][0].Value.(int32))
	}

	resultPipeline := mongo.Pipeline{}
	resultPipeline = append(resultPipeline, matchPipeline...)
	resultPipeline = append(resultPipeline, sortPipeline...)
	resultPipeline = append(resultPipeline, pagePipeline...)

	cursor, err = sendingCollection.Aggregate(ctx, resultPipeline)
	if err != nil {
		return 0, nil, fmt.Errorf("FindSendings: %v", err)
	}
	var results []Sending
	if err = cursor.All(ctx, &results); err != nil {
		return 0, nil, fmt.Errorf("FindSendings: %v", err)
	}

	return total, results, nil
}

func (s *Sending) Export() (string, error) {
	f, err := os.CreateTemp("", "sendings")
	if err != nil {
		return "", err
	}
	defer f.Close()

	RawArgs := []string{
		fmt.Sprintf("--uri=%s", db.GetDBURI()),
		"--db=post",
		"--collection=sendings",
		fmt.Sprintf("--out=%s", f.Name()),
		"--jsonArray",
		"--pretty",
	}

	Options, err := mongoexport.ParseOptions(RawArgs, "", "")
	if err != nil {
		return "", err
	}

	MongoExport, err := mongoexport.New(Options)
	if err != nil {
		return "", err
	}
	defer MongoExport.Close()

	_, err = MongoExport.Export(f)
	if err != nil {
		return "", err
	}

	return f.Name(), nil
}

func (s *Sending) Import(filename string) error {
	client := db.GetDB()
	sendingCollection := client.Database("post").Collection("sendings")
	testCollection := client.Database("post").Collection("test")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	RawArgsTest := []string{
		fmt.Sprintf("--uri=%s", db.GetDBURI()),
		"--db=post",
		"--collection=test",
		fmt.Sprintf("--file=%s", filename),
		"--jsonArray",
	}

	RawArgs := []string{
		fmt.Sprintf("--uri=%s", db.GetDBURI()),
		"--db=post",
		"--collection=sendings",
		fmt.Sprintf("--file=%s", filename),
		"--jsonArray",
	}
	{
		Options, err := mongoimport.ParseOptions(RawArgsTest, "", "")
		if err != nil {
			return err
		}
		MongoImport, err := mongoimport.New(Options)
		if err != nil {
			return err
		}
		defer MongoImport.Close()
		success, failed, err := MongoImport.ImportDocuments()
		if err != nil {
			return err
		}
		if failed > 0 {
			return errors.New(fmt.Sprintf("%d imported and %d aborted", success, failed))
		}
		if err = testCollection.Drop(ctx); err != nil {
			return err
		}
	}
	Options, err := mongoimport.ParseOptions(RawArgs, "", "")
	if err != nil {
		return err
	}
	MongoImport, err := mongoimport.New(Options)
	if err != nil {
		return err
	}
	defer MongoImport.Close()

	if err := sendingCollection.Drop(ctx); err != nil {
		return err
	}
	success, failed, err := MongoImport.ImportDocuments()
	if err != nil {
		return err
	}
	if failed > 0 {
		return errors.New(fmt.Sprintf("%d imported and %d aborted", success, failed))
	}

	return nil
}
