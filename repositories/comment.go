package repositories

import (
	"context"
	"subscription-report/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type _commentRepository struct {
	db *mongo.Database
}

type CommentRepository interface {
	GetComments(filter any) ([]models.Comment, error)
}

func NewCommentRepository(db *mongo.Database) CommentRepository {
	return &_commentRepository{db: db}
}

func (c *_commentRepository) GetComments(filter any) ([]models.Comment, error) {
	collection := c.db.Collection("comments")
	var comments []models.Comment

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return comments, err
	}
	if err = cursor.All(context.TODO(), &comments); err != nil {
		return comments, err
	}
	for _, comment := range comments {
		cursor.Decode(&comment)
	}
	return comments, nil
}
