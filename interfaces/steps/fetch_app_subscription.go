package steps_interfaces

import (
	"context"
	"subscription-report/internal/entities"
	options_subscription_report "subscription-report/internal/subscription_report/options"
)

type FetchAppSubscriptionStep interface {
	Exec(ctx context.Context, reports []*entities.SubscriptionReport, option *options_subscription_report.ReportOption) []*entities.SubscriptionReport
}
