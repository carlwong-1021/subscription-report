package grpc_client

import (
	"context"
	"crypto/tls"
	"strconv"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/shoplineapp/go-app/plugins/env"
	"github.com/shoplineapp/go-app/plugins/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGrpcClientConnection(logger *logger.Logger, env *env.Env, endpointEnvKey string, sslEnvKey string, dialTimeoutEnvKey string, interceptors []grpc.UnaryClientInterceptor) *grpc.ClientConn {
	timeout, err := strconv.Atoi(env.GetEnv(dialTimeoutEnvKey))
	if err != nil {
		logger.Infof("wrong timeout environmental variable %s: %v", dialTimeoutEnvKey, err)
		panic("wrong timeout environmental variable")
	}

	var opt grpc.DialOption
	if env.GetEnv(sslEnvKey) == "true" {
		opt = grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: true}))
	} else {
		opt = grpc.WithTransportCredentials(insecure.NewCredentials())
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		env.GetEnv(endpointEnvKey),
		opt,
		grpc.WithUnaryInterceptor(
			grpc_middleware.ChainUnaryClient(
				interceptors...,
			),
		),
	)
	if err != nil {
		logger.Infof("Fail to connect %s: %v", endpointEnvKey, err)
		panic("Fail to connect")
	}

	return conn
}
