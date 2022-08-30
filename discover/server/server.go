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

	taget := fmt.Sprintf("grpc-demo/grpc/%s", types.SERVER_NAME)
	err := etcdv3.Register(taget, "127.0.0.1", port, "http://127.0.0.1:2379", 10*time.Second, 15)
	if err != nil {
		panic(err)
	}
	lis, _ := net.Listen("tcp", ":"+port)
	server.Serve(lis)
}

//	import (
//		"go.etcd.io/etcd/client/v3"
//		"go.etcd.io/etcd/client/v3/naming/endpoints"
//		"go.etcd.io/etcd/client/v3/naming/resolver"
//		"google.golang.org/grpc"
//	)
//
// First, register new endpoint addresses for a service:
//
//	func etcdAdd(c *clientv3.Client, service, addr string) error {
//		em := endpoints.NewManager(c, service)
//		return em.AddEndpoint(c.Ctx(), service+"/"+addr, endpoints.Endpoint{Addr:addr});
//	}
//
// Dial an RPC service using the etcd gRPC resolver and a gRPC Balancer:
//
//	func etcdDial(c *clientv3.Client, service string) (*grpc.ClientConn, error) {
//		etcdResolver, err := resolver.NewBuilder(c);
//		if err { return nil, err }
//		return  grpc.Dial("etcd:///" + service, grpc.WithResolvers(etcdResolver))
//	}
//
// Optionally, force delete an endpoint:
//
//	func etcdDelete(c *clientv3, service, addr string) error {
//		em := endpoints.NewManager(c, service)
//		return em.DeleteEndpoint(c.Ctx(), service+"/"+addr)
//	}
//
// Or register an expiring endpoint with a lease:
//
//	func etcdAdd(c *clientv3.Client, lid clientv3.LeaseID, service, addr string) error {
//		em := endpoints.NewManager(c, service)
//		return em.AddEndpoint(c.Ctx(), service+"/"+addr, endpoints.Endpoint{Addr:addr}, clientv3.WithLease(lid));
//	}
//
