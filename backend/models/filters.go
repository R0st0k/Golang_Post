package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func matchStringArray(options map[string]interface{}, option, field string, pipeline *mongo.Pipeline) {
	if filter, ok := options[option]; ok {
		result := bson.A{}
		for _, data := range filter.([]string) {
			result = append(result, data)
		}
		match := bson.D{{
			"$match", bson.D{{
				field, bson.D{{
					"$in", result,
				}},
			}},
		}}
		*pipeline = append(*pipeline, match)
	}
}
