package interceptors

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func HelloInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Println("Hello")
	resp, err = handler(ctx, req)
	log.Println("Bye bye")
	return
}
