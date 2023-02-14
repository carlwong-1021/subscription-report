package steps

import (
	"fmt"
	"subscription-report/models"
	"subscription-report/services"
	"time"

	"github.com/jmoiron/sqlx"
)

type _queryLatestChangeStep struct {
	db *sqlx.DB
}

func NewQueryLatestChangeStep(db *sqlx.DB) services.Step {
	return &_queryLatestChangeStep{db: db}
}

func (q _queryLatestChangeStep) Exec(input any, option *services.ReportOption) (any, *services.ReportOption) {
	layout := "2006-01-02"
	from, _ := time.Parse(layout, option.From)
	to, _ := time.Parse(layout, option.To)

	var orders []models.Orders
	rows, err := q.db.Queryx("SELECT orderNumber, orderDate, status, customerNumber FROM orders WHERE orderDate >= ? AND orderDate <= ? ORDER BY orderNumber", from, to)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var order models.Orders
		err = rows.StructScan(&order)
		orders = append(orders, order)
	}
	return orders, option
}
