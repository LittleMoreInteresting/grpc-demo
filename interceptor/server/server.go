package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/grpc-demo/interceptor/interceptors"
	"github.com/grpc-demo/interceptor/pb"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Speaker struct {
	pb.UnimplementedSpeakerServer
}

func (s Speaker) Speak(ctx context.Context, req *pb.Request) (*pb.Reply, error) {
	reply := &pb.Reply{
		Message: req.Name + " : " + req.Content,
	}
	log.Println("Speak>" + reply.Message)
	time.Sleep(4 * time.Second)
	return reply, nil
}

func main() {
	port := "9090"
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			interceptors.HelloInterceptor,
			interceptors.DurationInterceptor,
		)),
	}
	server := grpc.NewServer(opts...)
	pb.RegisterSpeakerServer(server, &Speaker{})
	reflection.Register(server)
	listen, _ := net.Listen("tcp", ":"+port)
	_ = server.Serve(listen)
}
