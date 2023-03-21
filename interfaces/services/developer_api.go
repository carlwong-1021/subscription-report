package services_interfaces

import (
	"context"
	pb "subscription-report/developer-protos-src/go/developer_api"
)

type DeveloperApiClient interface {
	GetApplications(ctx context.Context) ([]*pb.Application, error)
}
