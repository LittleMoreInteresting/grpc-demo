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
	cli, err := clientv3.NewFromURL("http://127.0.0.1:2379")
	etcdResolver, err := resolver.NewBuilder(cli)

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithResolvers(etcdResolver),
	}
	taget := "etcd:///grpc-demo/grpc/discover-demo"
	//taget := "127.0.0.1:8000"
	conn, err := grpc.Dial(taget, opts...)

	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := discover.NewDiscoverDemoClient(conn)
	reply, err := client.Discover(context.Background(), &discover.Request{Name: "DDDD"})
	fmt.Println(err)
	fmt.Println(reply)
}
