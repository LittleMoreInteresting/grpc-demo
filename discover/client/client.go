package main

import (
	"context"
	"flag"
	"fmt"

	discover "github.com/grpc-demo/discover/pb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var port string

func init() {
	flag.StringVar(&port, "p", "8000", "启动端口号")
	flag.Parse()
}
func main() {
	cli, _ := clientv3.NewFromURL("http://localhost:2379")
	etcdResolver, _ := resolver.NewBuilder(cli)
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithResolvers(etcdResolver),
	}
	conn, _ := grpc.Dial(":"+port, opts...)
	defer conn.Close()
	client := discover.NewDiscoverDemoClient(conn)
	reply, _ := client.Discover(context.Background(), &discover.Request{Name: "DDDD"})
	fmt.Println(reply)
}
