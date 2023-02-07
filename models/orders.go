package models

import "time"

type Orders struct {
	OrderNumber    uint      `gorm:"column:orderNumber;primaryKey"`
	OrderDate      time.Time `gorm:"column:orderDate"`
	RequiredDate   time.Time `gorm:"column:requiredDate"`
	ShippedDate    time.Time `gorm:"column:shippedDate"`
	Status         string    `gorm:"column:status"`
	Comments       string    `gorm:"column:comments"`
	CustomerNumber uint      `gorm:"column:customerNumber"`
}
