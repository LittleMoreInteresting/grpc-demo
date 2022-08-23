package interceptors

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

func DurationInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (resp interface{}, err error) {
	now := time.Now()
	resp, err = handler(ctx, req)
	log.Printf("Duration:%+v", time.Since(now).Microseconds())
	return
}
