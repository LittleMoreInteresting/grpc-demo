package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"

	"github.com/grpc-demo/safety/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {

	certificate, err := tls.LoadX509KeyPair("keys/client.crt", "keys/client.key")
	if err != nil {
		log.Fatal(err)
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("keys/ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal(" failed to append certs ")
	}
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ServerName:   "www.test.com", // NOTE: 需要与生成证书的配置匹配
		RootCAs:      certPool,
	})
	opt := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}
	conn, _ := grpc.Dial(":9090", opt...)
	defer conn.Close()

	client := pb.NewSafetyDemoClient(conn)

	speak, err := client.Secret(context.Background(), &pb.Request{Name: "golang"})
	if err != nil {
		log.Println("error")
		log.Fatal(err)
		return
	}
	log.Println(speak)
}
