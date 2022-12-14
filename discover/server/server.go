package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	discover "github.com/grpc-demo/discover/pb"
	"github.com/grpc-demo/discover/register"
	"github.com/grpc-demo/discover/types"
	etcd3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type DiscoverDemo struct {
	discover.UnimplementedDiscoverDemoServer
}

func (ds *DiscoverDemo) Discover(ctx context.Context, in *discover.Request) (*discover.Reply, error) {
	return &discover.Reply{
		Content: "request content from" + port + " => " + in.Name,
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

	taget := fmt.Sprintf("grpc-demo/grpc/%s", types.SERVER_NAME)
	client, err := etcd3.New(etcd3.Config{
		Endpoints: []string{"http://127.0.0.1:2379"},
	})
	if err != nil {
		panic(err)
	}
	addr := "127.0.0.1:" + port
	err = register.EtcdAdd(client, taget, "127.0.0.1:"+port)
	//err = register.EtcdKeepAlive(client, taget, addr, 15)
	if err != nil {
		return
	}
	lis, _ := net.Listen("tcp", addr)
	go server.Serve(lis)

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shotdown Server ...")
	_ = register.EtcdDelete(client, taget, addr)
}
