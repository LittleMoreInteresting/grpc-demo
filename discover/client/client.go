package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	discover "github.com/grpc-demo/discover/pb"
	"github.com/grpc-demo/discover/register"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var port string

func init() {
	flag.StringVar(&port, "p", "8000", "启动端口号")
	flag.Parse()
}
func main() {
	cli, _ := clientv3.NewFromURL("http://127.0.0.1:2379")

	taget := "grpc-demo/grpc/discover-demo"

	n := 0
	for n < 50 {
		conn, err := register.EtcdDial(cli, taget)
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		client := discover.NewDiscoverDemoClient(conn)
		reply, err := client.Discover(context.Background(), &discover.Request{Name: "DDDD"})
		if err != nil {
			panic(err)
		}
		fmt.Println(reply)
		n++
		time.Sleep(time.Second)
	}
}
