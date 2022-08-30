package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"time"

	etcdv3 "github.com/grpc-demo/discover/etctv3"
	discover "github.com/grpc-demo/discover/pb"
	"github.com/grpc-demo/discover/types"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type DiscoverDemo struct {
	discover.UnimplementedDiscoverDemoServer
}

func (ds *DiscoverDemo) Discover(ctx context.Context, in *discover.Request) (*discover.Reply, error) {
	return &discover.Reply{
		Content: "request content" + in.Name,
	}, nil
}

var port string

func init() {
	flag.StringVar(&port, "p", "8000", "启动端口号")
	flag.Parse()
}
func main() {
	server := grpc.NewServer()
	discover.RegisterDiscoverDemoServer(server, &DiscoverDemo{})
	reflection.Register(server)

	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http:127.0.0.1:2379"},
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		panic(err)
	}
	defer etcdClient.Close()
	taget := fmt.Sprintf("/etctv3://grpc-demo/grpc/%s", types.SERVER_NAME)
	err = etcdv3.Register(taget, "127.0.0.1", port, "http:127.0.0.1:2379", 10*time.Second, 15)
	if err != nil {
		panic(err)
	}
	lis, _ := net.Listen("tcp", ":"+port)
	server.Serve(lis)
}
