package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/grpc-demo/streaming/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var port string

func init() {
	flag.StringVar(&port, "p", "8000", "启动端口号")
	flag.Parse()
}

type StreamServer struct {
	pb.UnimplementedStreamingServer
}

func (ss StreamServer) ServerStream(in *pb.Request, stream pb.Streaming_ServerStreamServer) error {
	for i := 0; i < 20; i++ {
		err := stream.Send(&pb.Reply{Type: "Server-Side stream", Value: fmt.Sprintf("val:%d", i)})
		if err != nil {
			return err
		}
	}
	return nil
}

func (ss StreamServer) ClientStream(stream pb.Streaming_ClientStreamServer) error {
	for {
		recv, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Reply{
				Type:  "Client-Side stream",
				Value: "close",
			})
		}
		if err != nil {
			return err
		}
		log.Printf("%v", recv)
	}
}
func (ss StreamServer) Bidirectional(stream pb.Streaming_BidirectionalServer) error {
	i := 0
	for {
		_ = stream.Send(&pb.Reply{Type: "Bidirectional stream", Value: fmt.Sprintf("val:%d", i)})
		recv, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("%v", recv)
		i++
	}
}

func main() {
	server := grpc.NewServer()
	pb.RegisterStreamingServer(server, &StreamServer{})
	reflection.Register(server)
	lis, _ := net.Listen("tcp", ":"+port)
	_ = server.Serve(lis)
}
