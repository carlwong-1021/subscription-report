package models

import "time"

type Orders struct {
	OrderNumber    uint       `db:"orderNumber"`
	OrderDate      *time.Time `db:"orderDate"`
	RequiredDate   *time.Time `db:"requiredDate"`
	ShippedDate    *time.Time `db:"shippedDate"`
	Status         *string    `db:"status"`
	Comments       *string    `db:"comments"`
	CustomerNumber uint       `db:"customerNumber"`
}
