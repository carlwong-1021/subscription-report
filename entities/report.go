package entities

import "time"

type Report struct {
	OrderNumber    uint
	OrderDate      time.Time
	RequiredDate   time.Time
	ShippedDate    time.Time
	Status         string
	Comments       string
	CustomerNumber string
	CustomerName   string
	Phone          string
	City           string
}
