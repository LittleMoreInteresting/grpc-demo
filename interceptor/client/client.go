package main

import (
	"context"
	"flag"
	"log"

	"github.com/grpc-demo/interceptor/interceptors"
	pb "github.com/grpc-demo/interceptor/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var port string

func init() {
	flag.StringVar(&port, "p", "9090", "启动端口号")
	flag.Parse()
}
func main() {
	opt := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(interceptors.TimeoutInterceptor()),
	}
	conn, _ := grpc.Dial(":"+port, opt...)
	defer conn.Close()

	client := pb.NewSpeakerClient(conn)
	speak, err := client.Speak(context.Background(), &pb.Request{Name: "golang", Content: "gRPC is great"})
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(speak)
}
