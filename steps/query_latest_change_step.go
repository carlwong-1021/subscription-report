package steps

import (
	"subscription-report/repositories"
	"subscription-report/services"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type _queryLatestChangeStep struct {
	repo repositories.CommentRepository
}

func NewQueryLatestChangeStep(repo repositories.CommentRepository) services.Step {
	return &_queryLatestChangeStep{repo: repo}
}

func (q *_queryLatestChangeStep) Exec(input any, option *services.ReportOption) (any, *services.ReportOption) {
	layout := "2006-01-02"
	from, _ := time.Parse(layout, option.From)
	to, _ := time.Parse(layout, option.To)

	filter := bson.M{"date": bson.M{"$gte": from, "$lte": to}}
	comments, err := q.repo.GetComments(filter)
	if err != nil {
		panic(err)
	}
	return comments, option
}
