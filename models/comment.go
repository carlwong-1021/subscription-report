package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	ID      primitive.ObjectID `bson:"_id"`
	Name    string
	Date    *time.Time
	MovieId primitive.ObjectID `bson:"movie_id"`
}
