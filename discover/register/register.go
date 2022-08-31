package register

import (
	"fmt"
	"log"
	"time"

	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
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
			_, err = c.Get(c.Ctx(), target)
			em, _ := endpoints.NewManager(c, service)
			if err != nil {
				if err == rpctypes.ErrKeyNotFound { // create
					err := em.AddEndpoint(c.Ctx(), service+"/"+addr, endpoints.Endpoint{Addr: addr}, clientv3.WithLease(grant.ID))
					if err != nil {
						log.Println(err)
					}
				} else {
					log.Println(err)
				}
				log.Println("add")
			} else { //update
				update := endpoints.NewAddUpdateOpts(target, endpoints.Endpoint{Addr: addr}, clientv3.WithLease(grant.ID))
				err := em.Update(c.Ctx(), []*endpoints.UpdateWithOpts{update})
				if err != nil {
					log.Println(err)
				}
				log.Println("update")
			}
			select {
			case <-stopSignal:
				return
			case <-ticker.C:
			}
		}

	}()

	return nil
}

//func Register(name string, host string, port string, target string, interval time.Duration, ttl int) error {
//	serviceValue := fmt.Sprintf("%s:%s", host, port)
//	serviceKey = fmt.Sprintf("/%s/%s/%s", Prefix, name, serviceValue)
//	fmt.Printf(serviceKey)
//	// get endpoints for register dial address
//	var err error
//	client, err := etcd3.New(etcd3.Config{
//		Endpoints: strings.Split(target, ","),
//	})
//	if err != nil {
//		return fmt.Errorf("grpclb: create etcd3 client failed: %v", err)
//	}
//	go func() {
//		// invoke self-register with ticker
//		ticker := time.NewTicker(interval)
//		for {
//			// minimum lease TTL is ttl-second
//			resp, _ := client.Grant(context.TODO(), int64(ttl))
//			// should get first, if not exist, set it
//			_, err := client.Get(context.Background(), serviceKey)
//			if err != nil {
//				if err == rpctypes.ErrKeyNotFound {
//					if _, err := client.Put(context.TODO(), serviceKey, serviceValue, etcd3.WithLease(resp.ID)); err != nil {
//						log.Printf("grpclb: set service '%s' with ttl to etcd3 failed: %s", name, err.Error())
//					}
//				} else {
//					log.Printf("grpclb: service '%s' connect to etcd3 failed: %s", name, err.Error())
//				}
//			} else {
//				// refresh set to true for not notifying the watcher
//				if _, err := client.Put(context.Background(), serviceKey, serviceValue, etcd3.WithLease(resp.ID)); err != nil {
//					log.Printf("grpclb: refresh service '%s' with ttl to etcd3 failed: %s", name, err.Error())
//				}
//			}
//			select {
//			case <-stopSignal:
//				return
//			case <-ticker.C:
//			}
//		}
//	}()
//	return nil
//}
//
// Dial an RPC service using the etcd gRPC resolver and a gRPC Balancer:
//
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
func etcdDelete(c *clientv3.Client, service, addr string) error {
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
