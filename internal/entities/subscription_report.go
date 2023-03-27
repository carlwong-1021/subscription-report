package entities

import "time"

type SubscriptionReport struct {
	ApplicationId        string
	SubscriptionDate     *time.Time
	SubscriptionSource   string
	AppName              string
	AppPrice             string
	AppDeveloper         string
	MerchantID           string
	Handle               string
	ShopName             string
	ShopURL              string
	SubscriptionCycle    string
	SubscriptionFee      *float64
	OneOffActivationFee  *float64
	SubscriptionCurrency string
	BillNumber           string
	OrderId              string
	InvoiceType          string
	CompanyName          string
	TaxID                string
	CompanyAddress       string
	InvoiceReceiptEmail  string
}
