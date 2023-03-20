package subscription_report

import (
	services_interfaces "subscription-report/interfaces/services"
	steps_interfaces "subscription-report/interfaces/steps"
)

type _reportService struct {
	options *ReportOption
}

type ReportOption struct {
	From string
	To   string
}

func NewReportService(options *ReportOption) services_interfaces.ReportService {
	return &_reportService{
		options: options,
	}
}

func (r *_reportService) Exec(fetchAppSubscriptionStep steps_interfaces.FetchAppSubscriptionStep) {

}
