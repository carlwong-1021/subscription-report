package entities

type FieldDefinition struct {
	GetValue func(*SubscriptionReport) any
	EN       string
	CH       string
}
