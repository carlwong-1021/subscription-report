package steps_interfaces

import (
	"subscription-report/internal/entities"
	options_subscription_report "subscription-report/internal/subscription_report/options"
)

type GenerateReportStep interface {
	Exec(reports []*entities.SubscriptionReport, option *options_subscription_report.ReportOption)
}
