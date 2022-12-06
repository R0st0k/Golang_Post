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
	"regexp"
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

	pagePipeline := mongo.Pipeline{}
	matchPipeline := mongo.Pipeline{}
	sortPipeline := mongo.Pipeline{}

	if page := sendingFilter["page"].(int64); page > 1 {
		skip := bson.D{{
			"$skip", (page - 1) * sendingFilter["elems"].(int64),
		}}
		pagePipeline = append(pagePipeline, skip)
	}
	{
		limit := bson.D{{
			"$limit", sendingFilter["elems"].(int64),
		}}
		pagePipeline = append(pagePipeline, limit)
	}
	if filter, ok := sendingFilter["order_id"]; ok {
		match := bson.D{{
			"$match", bson.D{{
				"order_id", bson.D{{
					"$regex", regexp.QuoteMeta(filter.(string)),
				},
					{
						"$options", "i",
					}},
			}},
		}}
		matchPipeline = append(matchPipeline, match)
	}
	if filter, ok := sendingFilter["type"]; ok {
		result := bson.A{}
		for _, data := range filter.([]string) {
			result = append(result, data)
		}
		match := bson.D{{
			"$match", bson.D{{
				"type", bson.D{{
					"$in", result,
				}},
			}},
		}}
		matchPipeline = append(matchPipeline, match)
	}
	if filter, ok := sendingFilter["status"]; ok {
		result := bson.A{}
		for _, data := range filter.([]string) {
			result = append(result, data)
		}
		match := bson.D{{
			"$match", bson.D{{
				"status", bson.D{{
					"$in", result,
				}},
			}},
		}}
		matchPipeline = append(matchPipeline, match)
	}
	if filter, ok := sendingFilter["date_start"]; ok {
		match := bson.D{{
			"$match", bson.D{{
				"registration_date", bson.D{{
					"$gte", filter.(time.Time).Format(time.RFC3339),
				}},
			}},
		}}
		matchPipeline = append(matchPipeline, match)
	}
	if filter, ok := sendingFilter["date_finish"]; ok {
		match := bson.D{{
			"$match", bson.D{{
				"registration_date", bson.D{{
					"$lte", filter.(time.Time).Format(time.RFC3339),
				}},
			}},
		}}
		matchPipeline = append(matchPipeline, match)
	}
	if filter, ok := sendingFilter["sender_settlement"]; ok {
		match := bson.D{{
			"$match", bson.D{{
				"sender.address.settlement", bson.D{{
					"$regex", regexp.QuoteMeta(filter.(string)),
				},
					{
						"$options", "i",
					}},
			}},
		}}
		matchPipeline = append(matchPipeline, match)
	}
	if filter, ok := sendingFilter["receiver_settlement"]; ok {
		match := bson.D{{
			"$match", bson.D{{
				"receiver.address.settlement", bson.D{{
					"$regex", regexp.QuoteMeta(filter.(string)),
				},
					{
						"$options", "i",
					}},
			}},
		}}
		matchPipeline = append(matchPipeline, match)
	}
	if filter, ok := sendingFilter["length_min"]; ok {
		match := bson.D{{
			"$match", bson.D{{
				"size.length", bson.D{{
					"$gte", filter.(int64),
				}},
			}},
		}}
		matchPipeline = append(matchPipeline, match)
	}
	if filter, ok := sendingFilter["length_max"]; ok {
		match := bson.D{{
			"$match", bson.D{{
				"size.length", bson.D{{
					"$lte", filter.(int64),
				}},
			}},
		}}
		matchPipeline = append(matchPipeline, match)
	}
	if filter, ok := sendingFilter["width_min"]; ok {
		match := bson.D{{
			"$match", bson.D{{
				"size.width", bson.D{{
					"$gte", filter.(int64),
				}},
			}},
		}}
		matchPipeline = append(matchPipeline, match)
	}
	if filter, ok := sendingFilter["width_max"]; ok {
		match := bson.D{{
			"$match", bson.D{{
				"size.width", bson.D{{
					"$lte", filter.(int64),
				}},
			}},
		}}
		matchPipeline = append(matchPipeline, match)
	}
	if filter, ok := sendingFilter["height_min"]; ok {
		match := bson.D{{
			"$match", bson.D{{
				"size.height", bson.D{{
					"$gte", filter.(int64),
				}},
			}},
		}}
		matchPipeline = append(matchPipeline, match)
	}
	if filter, ok := sendingFilter["height_max"]; ok {
		match := bson.D{{
			"$match", bson.D{{
				"size.height", bson.D{{
					"$lte", filter.(int64),
				}},
			}},
		}}
		matchPipeline = append(matchPipeline, match)
	}
	if filter, ok := sendingFilter["weight_min"]; ok {
		match := bson.D{{
			"$match", bson.D{{
				"weight", bson.D{{
					"$gte", filter.(int64),
				}},
			}},
		}}
		matchPipeline = append(matchPipeline, match)
	}
	if filter, ok := sendingFilter["weight_max"]; ok {
		match := bson.D{{
			"$match", bson.D{{
				"weight", bson.D{{
					"$lte", filter.(int64),
				}},
			}},
		}}
		matchPipeline = append(matchPipeline, match)
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
