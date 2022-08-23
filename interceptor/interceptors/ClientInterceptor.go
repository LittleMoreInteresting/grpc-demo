package interceptors

import (
	"context"
	"time"

	"google.golang.org/grpc"
)

func TimeoutInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context,
		method string,
		req,
		reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption) error {
		if _, ok := ctx.Deadline(); !ok {
			timeout := 3 * time.Second
			var cancel context.CancelFunc
			ctx, cancel = context.WithTimeout(ctx, timeout)
			defer cancel()
		}
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
