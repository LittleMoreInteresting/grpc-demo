package main

import (
	"context"
	"flag"
	"log"

	pb "github.com/grpc-demo/hello/proto"
	"google.golang.org/grpc"
)

var port string

func init() {
	flag.StringVar(&port, "p", "8000", "启动端口号")
	flag.Parse()
}
func main() {
	conn, _ := grpc.Dial(":"+port, grpc.WithInsecure())
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	_ = SayHello(client)
}

func SayHello(client pb.GreeterClient) error {
	reapy := &pb.HelloRequest{
		Name: "eddycjy",
		Call: &pb.HelloRequest_Mobile{Mobile: "151"},
		Role: map[int64]string{1: "888"},
	}
	resp, _ := client.SayHello(context.Background(), reapy)
	log.Printf("client.SayHello resp: %s", resp.Message)
	return nil
}
