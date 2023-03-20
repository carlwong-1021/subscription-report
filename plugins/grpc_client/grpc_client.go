package grpc_client

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"
	"strconv"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGrpcClientConnection(endpointEnvKey string, sslEnvKey string, dialTimeoutEnvKey string, interceptors []grpc.UnaryClientInterceptor) *grpc.ClientConn {
	timeout, err := strconv.Atoi(os.Getenv(dialTimeoutEnvKey))
	if err != nil {
		fmt.Printf("wrong timeout environmental variable %s: %v", dialTimeoutEnvKey, err)
		panic("wrong timeout environmental variable")
	}

	var opt grpc.DialOption
	if os.Getenv(sslEnvKey) == "true" {
		opt = grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: true}))
	} else {
		opt = grpc.WithTransportCredentials(insecure.NewCredentials())
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		os.Getenv(endpointEnvKey),
		opt,
		grpc.WithUnaryInterceptor(
			grpc_middleware.ChainUnaryClient(
				interceptors...,
			),
		),
	)
	if err != nil {
		fmt.Printf("Fail to connect %s: %v", endpointEnvKey, err)
		panic("Fail to connect")
	}

	return conn
}
