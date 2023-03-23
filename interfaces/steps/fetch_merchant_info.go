package steps_interfaces

import (
	"context"
	"subscription-report/internal/entities"
	options_subscription_report "subscription-report/internal/subscription_report/options"
)

type FetchMerchantInfoStep interface {
	Exec(ctx context.Context, reports []*entities.SubscriptionReport, opts *options_subscription_report.ReportOption) ([]*entities.SubscriptionReport, error)
}
