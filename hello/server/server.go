package main

import (
	"context"
	"flag"
	"fmt"
	"net"

	pb "github.com/grpc-demo/hello/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/anypb"
)

var port string

func init() {
	flag.StringVar(&port, "p", "8000", "启动端口号")
	flag.Parse()
}

type Server struct {
	pb.UnimplementedGreeterServer
}

func (gs Server) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloReply, error) {
	str := r.GetMobile()
	if len(str) == 0 {
		str = r.GetPhone()
	}
	var url string
	for k, v := range r.Role {
		url = fmt.Sprintf("%d=%s&", k, v)
	}
	a, _ := anypb.New(r)
	return &pb.HelloReply{
		Message: "Hello World " + r.Name + str + "?" + url,
		Details: []*anypb.Any{a},
	}, nil
}
func main() {
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &Server{})
	reflection.Register(server)
	lis, _ := net.Listen("tcp", ":"+port)
	server.Serve(lis)
}
