package main

import (
	"context"
	"net/http"

	gateway "github.com/grpc-demo/gateway/pb"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	// gRPC服务地址
	ServerAddr = "127.0.0.1:9000"

	ClientAddr = "127.0.0.1:8000"
)

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := gateway.RegisterGatewayDemoHandlerFromEndpoint(ctx, mux, ServerAddr, opts)
	if err != nil {
		panic(err)
	}
	err = http.ListenAndServe(ClientAddr, mux)
	if err != nil {
		panic(err)
	}
}
