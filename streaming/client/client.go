package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"

	"github.com/grpc-demo/streaming/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var port string
var method int

func init() {
	flag.StringVar(&port, "p", "8000", "启动端口号")
	flag.IntVar(&method, "m", 1, "流模式")
	flag.Parse()
}
func ServerStreamClient(client pb.StreamingClient, r *pb.Request) error {
	stream, err := client.ServerStream(context.Background(), r)
	if err != nil {
		return err
	}
	for {
		recv, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Printf("response:%v", recv)
	}
	return nil
}

func ClientStreamClient(client pb.StreamingClient) error {
	stream, err := client.ClientStream(context.Background())
	if err != nil {
		return err
	}
	for i := 0; i < 10; i++ {
		_ = stream.Send(&pb.Request{Type: "ClientStreamClient", Value: fmt.Sprintf("val:%d", i)})
	}
	recv, err := stream.CloseAndRecv()

	log.Printf("resp :%v", recv)
	return err
}

func Bidirectional(client pb.StreamingClient) error {
	stream, err := client.Bidirectional(context.Background())
	if err != nil {
		return err
	}
	for i := 0; i < 10; i++ {
		_ = stream.Send(&pb.Request{Type: "BidirectionalClient", Value: fmt.Sprintf("val:%d", i)})
		recv, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("resp :%v", err)
			return err
		}
		log.Printf("resp :%v", recv)
	}
	return stream.CloseSend()
}

func main() {
	conn, _ := grpc.Dial(":"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	log.Printf("mothod %d", method)
	client := pb.NewStreamingClient(conn)
	switch method {
	case 1:
		_ = ServerStreamClient(client, &pb.Request{Type: "ss client", Value: "ok"})
	case 2:
		_ = ClientStreamClient(client)
	case 3:
		_ = Bidirectional(client)
	default:
		log.Fatal("error method")
	}
}
