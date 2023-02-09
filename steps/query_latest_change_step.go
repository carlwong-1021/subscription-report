package steps

import (
	"subscription-report/models"
	"subscription-report/services"
	"time"

	"gorm.io/gorm"
)

type _queryLatestChangeStep struct {
	db *gorm.DB
}

func NewQueryLatestChangeStep(db *gorm.DB) services.Step {
	return &_queryLatestChangeStep{db: db}
}

func (q _queryLatestChangeStep) Exec(input any, option *services.ReportOption) (any, *services.ReportOption) {
	layout := "2006-01-02"
	from, _ := time.Parse(layout, option.From)
	to, _ := time.Parse(layout, option.To)

	var orders []models.Orders
	q.db.Debug().Where("orderDate >= ? and orderDate <= ?", from, to).Order("orderDate").Find(&orders)
	return orders, option
}
