package subscription_report

import (
	services_interfaces "subscription-report/interfaces/services"
	steps_interfaces "subscription-report/interfaces/steps"
	"subscription-report/internal/entities"
	options_subscription_report "subscription-report/internal/subscription_report/options"
	"time"

	"golang.org/x/net/context"
)

type _reportService struct {
	options *options_subscription_report.ReportOption
}

func NewReportService(options *options_subscription_report.ReportOption) services_interfaces.ReportService {
	return &_reportService{
		options: options,
	}
}

func (r *_reportService) Exec(fetchAppSubscriptionStep steps_interfaces.FetchAppSubscriptionStep) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	reports := []*entities.SubscriptionReport{}
	reports = fetchAppSubscriptionStep.Exec(ctx, reports, r.options)
}
