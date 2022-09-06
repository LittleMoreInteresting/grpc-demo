package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"

	"github.com/grpc-demo/safety/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

type SafetyServer struct {
	pb.UnsafeSafetyDemoServer
}

func (s SafetyServer) Secret(ctx context.Context, in *pb.Request) (*pb.Reply, error) {
	return &pb.Reply{
		Content: "Secret Content： " + in.Name,
	}, nil
}

func main() {

	certificate, err := tls.LoadX509KeyPair("keys/server.crt", "keys/server.key")
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
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})

	server := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterSafetyDemoServer(server, &SafetyServer{})
	reflection.Register(server)
	lis, _ := net.Listen("tcp", ":9090")
	server.Serve(lis)
}
