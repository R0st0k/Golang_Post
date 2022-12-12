package models

import (
	"backend/db"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

type EmployeeDemonstration struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name        string             `bson:"name" json:"name" example:"Райан"`
	Surname     string             `bson:"surname" json:"surname" example:"Гослинг"`
	MiddleName  string             `bson:"middle_name,omitempty" json:"middle_name,omitempty" example:"Томасович"`
	Gender      string             `bson:"gender" json:"gender" example:"М"`
	BirthDate   time.Time          `bson:"birth_date" json:"birth_date"`
	Position    string             `bson:"position" json:"position" example:"Водитель"`
	PhoneNumber string             `bson:"phone_number" json:"phone_number" example:"88005553535"`
	Settlement  string             `bson:"settlement" json:"settlement" example:"пос. Нугуш"`
	Postcode    string             `bson:"postcode" json:"postcode" example:"453870"`
}

func (e *Employee) FilterEmployee(employeeFilter map[string]interface{}) (int64, []EmployeeDemonstration, error) {
	client := db.GetDB()
	employeeCollection := client.Database("post").Collection("employees")

	var err error

	pagePipeline := getPagePipeline(employeeFilter, "page", "elems")
	structurePipeline := mongo.Pipeline{
		bson.D{
			{
				"$lookup", bson.D{
					{"from", "postOffices"},
					{"localField", "_id"},
					{"foreignField", "employees"},
					{"as", "office"},
				},
			},
		},
		bson.D{
			{
				"$unwind", "$office",
			},
		},
		bson.D{
			{
				"$addFields", bson.D{
					{"settlement", "$office.address.settlement"},
					{"postcode", "$office.address.postcode"},
				},
			},
		},
		bson.D{
			{
				"$unset", bson.A{"office", "_id"},
			},
		},
	}
	matchPipeline := mongo.Pipeline{}
	sortPipeline := mongo.Pipeline{}

	err = matchRegex(employeeFilter, "name", "name", "i", &matchPipeline)
	if err != nil {
		return 0, nil, err
	}
	err = matchRegex(employeeFilter, "surname", "surname", "i", &matchPipeline)
	if err != nil {
		return 0, nil, err
	}
	err = matchRegex(employeeFilter, "middle_name", "middle_name", "i", &matchPipeline)
	if err != nil {
		return 0, nil, err
	}
	err = matchArray(employeeFilter, "gender", "gender", "string", &matchPipeline)
	if err != nil {
		return 0, nil, err
	}
	err = matchWithCompare(employeeFilter, "birth_date_start", "birth_date", "time", "$gte", &matchPipeline)
	if err != nil {
		return 0, nil, err
	}
	err = matchWithCompare(employeeFilter, "birth_date_finish", "birth_date", "time", "$lte", &matchPipeline)
	if err != nil {
		return 0, nil, err
	}
	err = matchArray(employeeFilter, "position", "position", "string", &matchPipeline)
	if err != nil {
		return 0, nil, err
	}
	err = matchRegex(employeeFilter, "phone_number", "phone_number", "i", &matchPipeline)
	if err != nil {
		return 0, nil, err
	}
	err = matchRegex(employeeFilter, "settlement", "settlement", "i", &matchPipeline)
	if err != nil {
		return 0, nil, err
	}
	err = matchRegex(employeeFilter, "postcode", "postcode", "i", &matchPipeline)
	if err != nil {
		return 0, nil, err
	}

	{
		sortTypeInMap, okType := employeeFilter["sort_type"]
		sortFieldInMap, okField := employeeFilter["sort_field"]
		if okType && okField {
			sortType := 0
			switch sortTypeInMap.(string) {
			case "asc":
				sortType = 1
			case "desc":
				sortType = -1
			}
			/*[ full_name, settlement, postcode, position, birth_date, gender, phone_number ]*/
			sortField := ""
			if sortFieldInMap.(string) != "full_name" {
				switch sortFieldInMap.(string) {
				case "settlement":
					sortField = "settlement"
				case "postcode":
					sortField = "postcode"
				case "position":
					sortField = "position"
				case "birth_date":
					sortField = "birth_date"
				case "gender":
					sortField = "gender"
				case "phone_number":
					sortField = "phone_number"
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
			} else {
				sort := bson.D{{
					"$sort", bson.D{
						{"surname", sortType},
						{"name", sortType},
						{"middle_name", sortType},
						{"_id", 1},
					},
				}}
				sortPipeline = append(sortPipeline, sort)
			}
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	countPipeline := mongo.Pipeline{}
	countStage := bson.D{{"$count", "name"}}
	countPipeline = append(countPipeline, structurePipeline...)
	countPipeline = append(countPipeline, matchPipeline...)
	countPipeline = append(countPipeline, countStage)

	cursor, err := employeeCollection.Aggregate(ctx, countPipeline)
	if err != nil {
		return 0, nil, fmt.Errorf("FindEmployees: %v", err)
	}
	var count []bson.D
	if err = cursor.All(ctx, &count); err != nil {
		return 0, nil, fmt.Errorf("FindEmployees: %v", err)
	}
	total := int64(0)
	if len(count) > 0 {
		total = int64(count[0][0].Value.(int32))
	}

	resultPipeline := mongo.Pipeline{}
	resultPipeline = append(resultPipeline, structurePipeline...)
	resultPipeline = append(resultPipeline, matchPipeline...)
	resultPipeline = append(resultPipeline, sortPipeline...)
	resultPipeline = append(resultPipeline, pagePipeline...)

	cursor, err = employeeCollection.Aggregate(ctx, resultPipeline)
	if err != nil {
		return 0, nil, fmt.Errorf("FindEmployees: %v", err)
	}
	var results []EmployeeDemonstration
	if err = cursor.All(ctx, &results); err != nil {
		return 0, nil, fmt.Errorf("FindEmployees: %v", err)
	}

	return total, results, nil

}
