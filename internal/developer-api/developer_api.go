package developer_api

import (
	service_interfaces "calculate_api/interfaces/services"
	"calculate_api/plugins/grpc_client"
	"context"
	"time"

	pb "bitbucket.org/starlinglabs/domain-data-protos/go/domain_data_api"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/newrelic/go-agent/v3/integrations/nrgrpc"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/shoplineapp/go-app/plugins/env"
	"github.com/shoplineapp/go-app/plugins/logger"
	"google.golang.org/grpc"
)

const (
	maxRetry          = 3
	retryBaseInterval = 100 * time.Millisecond
)

func NewDeveloperApiClient(logger *logger.Logger, env *env.Env) service_interfaces.DomainDataApiClient {
	opts := []grpc_retry.CallOption{
		grpc_retry.WithBackoff(grpc_retry.BackoffExponential(retryBaseInterval)),
		grpc_retry.WithMax(uint(maxRetry)),
	}
	conn := grpc_client.NewGrpcClientConnection(logger, env, "DOMAIN_DATA_API_ENDPOINT", "DOMAIN_DATA_API_SSL", "DOMAIN_DATA_API_DIAL_TIMEOUT", []grpc.UnaryClientInterceptor{
		grpc_retry.UnaryClientInterceptor(opts...),
		grpc_client.TraceIDClientInterceptor,
		func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
			txn := newrelic.FromContext(ctx)
			ctx = newrelic.NewContext(ctx, txn.NewGoroutine())
			return nrgrpc.UnaryClientInterceptor(ctx, method, req, reply, cc, invoker, opts...)
		},
	})
	return pb.NewDomainDataAPIClient(conn)
}
