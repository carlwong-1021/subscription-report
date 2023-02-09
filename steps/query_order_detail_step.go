package steps

import "subscription-report/services"

type _queryOrderDetailStep struct{}

func NewQueryOrderDetailStep() services.Step {
	return &_queryOrderDetailStep{}
}

func (q _queryOrderDetailStep) Exec(input any, option *services.ReportOption) (any, *services.ReportOption) {
	return input, option
}
