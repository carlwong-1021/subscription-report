package steps

import (
	"context"

	pb "subscription-report/developer-protos-src/go/developer_api"
	services_interfaces "subscription-report/interfaces/services"
	steps_interfaces "subscription-report/interfaces/steps"
	"subscription-report/internal/entities"
	options_subscription_report "subscription-report/internal/subscription_report/options"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type _fetchAppSubscriptionStep struct {
	client services_interfaces.DeveloperApiClient
}

func NewFetchAppSubscriptionStep(client services_interfaces.DeveloperApiClient) steps_interfaces.FetchAppSubscriptionStep {
	return &_fetchAppSubscriptionStep{client: client}
}

func (s *_fetchAppSubscriptionStep) Exec(ctx context.Context, opts *options_subscription_report.ReportOption) ([]*entities.SubscriptionReport, error) {
	from, to := timestamppb.New(opts.From), timestamppb.New(opts.To)
	req := &pb.ListSubscriptionReportsRequest{From: from, To: to}

	resp, err := s.client.GetSubscriptions(ctx, req)
	if err != nil {
		return nil, err
	}
	return reportResToEntity(resp), nil
}

func getAppPrice(isFree bool) string {
	if isFree {
		return "免費"
	}
	return "付費"
}

func reportResToEntity(reports []*pb.SubscriptionReport) []*entities.SubscriptionReport {
	results := []*entities.SubscriptionReport{}
	for _, report := range reports {
		subDate := report.CreatedAt.AsTime()

		results = append(results, &entities.SubscriptionReport{
			MerchantID:         report.MerchantId,
			ApplicationId:      report.ApplicationId,
			SubscriptionSource: report.PurchaseType,
			AppName:            report.ApplicationName,
			AppDeveloper:       report.BusinessName,
			AppPrice:           getAppPrice(report.IsFree),
			SubscriptionDate:   &subDate,
		})
	}

	return results
}
