package register

import (
	"fmt"
	"log"
	"time"

	"go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//
// First, register new endpoint addresses for a service:
//
var stopSignal = make(chan struct{})

func EtcdAdd(c *clientv3.Client, service, addr string) error {
	em, _ := endpoints.NewManager(c, service)
	list, _ := em.List(c.Ctx())
	fmt.Println(list)
	return em.AddEndpoint(c.Ctx(), service+"/"+addr, endpoints.Endpoint{Addr: addr})
}

func EtcdKeepAlive(c *clientv3.Client, service, addr string, ttl int64) error {
	target := service + "/" + addr

	go func() {
		ticker := time.NewTicker(time.Second * 10)
		for {
			grant, err := c.Grant(c.Ctx(), ttl)
			if err != nil {
				log.Println(err)
			}
			em, _ := endpoints.NewManager(c, service)
			update := endpoints.NewAddUpdateOpts(target, endpoints.Endpoint{Addr: addr}, clientv3.WithLease(grant.ID))
			err = em.Update(c.Ctx(), []*endpoints.UpdateWithOpts{update})
			if err != nil {
				log.Println(err)
			}
			log.Println("update")
			select {
			case <-stopSignal:
				return
			case <-ticker.C:
			}
		}

	}()

	return nil
}

func EtcdDial(c *clientv3.Client, service string) (*grpc.ClientConn, error) {
	etcdResolver, err := resolver.NewBuilder(c)
	if err != nil {
		return nil, err
	}
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithResolvers(etcdResolver),
	}
	return grpc.Dial("etcd:///"+service, opts...)
}

//
// Optionally, force delete an endpoint:
//
func EtcdDelete(c *clientv3.Client, service, addr string) error {
	em, _ := endpoints.NewManager(c, service)
	return em.DeleteEndpoint(c.Ctx(), service+"/"+addr)
}

//
// Or register an expiring endpoint with a lease:
//
//func etcdAdd(c *clientv3.Client, lid clientv3.LeaseID, service, addr string) error {
//	em, _ := endpoints.NewManager(c, service)
//	return em.AddEndpoint(c.Ctx(), service+"/"+addr, endpoints.Endpoint{Addr: addr}, clientv3.WithLease(lid))
//}
