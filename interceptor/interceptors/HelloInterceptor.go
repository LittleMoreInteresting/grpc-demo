package interceptors

import (
	"context"
	"google.golang.org/grpc"
	"log"
)

func HelloInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (resp interface{}, err error) {
		log.Println("Hello")
		resp, err = handler(ctx, req)
		log.Println("Bye bye")
		return
	}
}
