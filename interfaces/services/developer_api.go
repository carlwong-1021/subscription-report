package services_interfaces

import (
	"context"
	pb "subscription-report/developer-protos-src/go/developer_api"
)

type DeveloperApiClient interface {
	GetSubscriptions(ctx context.Context, req *pb.ListSubscriptionReportsRequest) ([]*pb.SubscriptionReport, error)
}
