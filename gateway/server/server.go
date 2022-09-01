package main

import (
	"context"
	"fmt"
	"net"

	gateway "github.com/grpc-demo/gateway/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Gate struct {
	gateway.UnimplementedGatewayDemoServer
}

func (g Gate) Gate(ctx context.Context, req *gateway.Request) (*gateway.Reply, error) {
	return &gateway.Reply{
		Content: fmt.Sprintf("name:%s;context:%v", req.Name, ctx),
	}, nil
}

func main() {
	server := grpc.NewServer()
	gateway.RegisterGatewayDemoServer(server, &Gate{})
	reflection.Register(server)
	lis, _ := net.Listen("tcp", ":9000")
	_ = server.Serve(lis)
}
