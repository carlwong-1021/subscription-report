package developer_api

import (
	"subscription-report/plugins/grpc_client"
	"time"

	pb "subscription-report/developer-protos-src/go/developer_api"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
)

const (
	maxRetry          = 3
	retryBaseInterval = 100 * time.Millisecond
)

func NewDeveloperApiClient() pb.ApplicationsClient {
	opts := []grpc_retry.CallOption{
		grpc_retry.WithBackoff(grpc_retry.BackoffExponential(retryBaseInterval)),
		grpc_retry.WithMax(uint(maxRetry)),
	}
	conn := grpc_client.NewGrpcClientConnection("DEVELOPER_API_ENDPOINT", "DEVELOPER_API_SSL", "DEVELOPER_API_DIAL_TIMEOUT", []grpc.UnaryClientInterceptor{
		grpc_retry.UnaryClientInterceptor(opts...),
	})
	return pb.NewApplicationsClient(conn)
}
