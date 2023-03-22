package steps

import (
	"context"
	"fmt"

	services_interfaces "subscription-report/interfaces/services"
	steps_interfaces "subscription-report/interfaces/steps"
	"subscription-report/internal/entities"
	options_subscription_report "subscription-report/internal/subscription_report/options"
)

type _fetchAppSubscriptionStep struct {
	client services_interfaces.DeveloperApiClient
}

func NewFetchAppSubscriptionStep(client services_interfaces.DeveloperApiClient) steps_interfaces.FetchAppSubscriptionStep {
	return &_fetchAppSubscriptionStep{client: client}
}

func (s *_fetchAppSubscriptionStep) Exec(ctx context.Context, reports []*entities.SubscriptionReport, option *options_subscription_report.ReportOption) []*entities.SubscriptionReport {
	resp, err := s.client.GetApplications(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}
