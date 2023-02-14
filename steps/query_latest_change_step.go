package steps

import (
	"context"
	"fmt"
	"subscription-report/models"
	"subscription-report/services"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type _queryLatestChangeStep struct {
	db *mongo.Database
}

func NewQueryLatestChangeStep(db *mongo.Database) services.Step {
	return &_queryLatestChangeStep{db: db}
}

func (q _queryLatestChangeStep) Exec(input any, option *services.ReportOption) (any, *services.ReportOption) {
	// layout := "2006-01-02"
	// from, _ := time.Parse(layout, option.From)
	// to, _ := time.Parse(layout, option.To)

	collection := q.db.Collection("comments")
	filter := bson.D{{"name", "John Bishop"}}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
	}

	var comments []models.Comment
	if err = cursor.All(context.TODO(), &comments); err != nil {
		panic(err)
	}
	for _, comment := range comments {
		cursor.Decode(&comment)
	}
	return comments, option
}
