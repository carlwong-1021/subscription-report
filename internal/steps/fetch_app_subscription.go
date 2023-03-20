package steps

import (
	"subscription-report/developer-protos-src/go/developer_api"
	steps_interfaces "subscription-report/interfaces/steps"
	"subscription-report/internal/subscription_report"
	"time"
)

type _fetchAppSubscriptionStep struct {
	client developer_api.ApplicationsClient
}

func NewFetchAppSubscriptionStep(client developer_api.ApplicationsClient) steps_interfaces.FetchAppSubscriptionStep {
	return &_fetchAppSubscriptionStep{client: client}
}

func (s *_fetchAppSubscriptionStep) Exec(option *subscription_report.ReportOption) {
	layout := "2006-01-02"
	from, _ := time.Parse(layout, option.From)
	to, _ := time.Parse(layout, option.To)

	resp, err := s.client.Index()
}
