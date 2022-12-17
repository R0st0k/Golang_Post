package models

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"regexp"
	"time"
)

func getPagePipeline(options map[string]interface{}, pageOption, elemsOnPageOption string) mongo.Pipeline {
	pagePipeline := mongo.Pipeline{}

	if page := options[pageOption].(int64); page > 1 {
		skip := bson.D{{
			"$skip", (page - 1) * options[elemsOnPageOption].(int64),
		}}
		pagePipeline = append(pagePipeline, skip)
	}
	{
		limit := bson.D{{
			"$limit", options[elemsOnPageOption].(int64),
		}}
		pagePipeline = append(pagePipeline, limit)
	}

	return pagePipeline
}

func matchArray(options map[string]interface{}, option, field, fieldType string, pipeline *mongo.Pipeline) error {
	if filter, ok := options[option]; ok {
		result := bson.A{}
		switch fieldType {
		case "string":
			filter = filter.([]string)
		default:
			return errors.New("unsupported format of option")
		}
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

	return nil
}

func matchWithCompare(options map[string]interface{}, option, field, fieldType, compOpt string, pipeline *mongo.Pipeline) error {
	if filter, ok := options[option]; ok {
		switch fieldType {
		case "int64":
			filter = filter.(int64)
		case "time":
			filter = filter.(time.Time)
		default:
			return errors.New("unsupported format of option")
		}
		match := bson.D{{
			"$match", bson.D{{
				field, bson.D{{
					compOpt, filter,
				}},
			}},
		}}
		*pipeline = append(*pipeline, match)
	}
	return nil
}

func matchRegex(options map[string]interface{}, option, field, regexOptions string, pipeline *mongo.Pipeline) error {
	if filter, ok := options[option]; ok {
		match := bson.D{{
			"$match", bson.D{{
				field, bson.D{{
					"$regex", regexp.QuoteMeta(filter.(string)),
				},
					{
						"$options", regexOptions,
					}},
			}},
		}}
		*pipeline = append(*pipeline, match)
	}

	return nil
}
