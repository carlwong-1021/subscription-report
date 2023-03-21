package steps_interfaces

import (
	"context"
	options_subscription_report "subscription-report/internal/subscription_report/options"
)

type FetchAppSubscriptionStep interface {
	Exec(ctx context.Context, option *options_subscription_report.ReportOption)
}
