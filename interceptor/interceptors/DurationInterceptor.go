package interceptors

import (
	"context"
	"fmt"
	"log"
	"runtime/debug"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

// go-zero 超时连接器
func UnaryTimeoutInterceptor(timeout time.Duration) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (interface{}, error) {
		var cancel context.CancelFunc
		if _, ok := ctx.Deadline(); !ok {
			ctx, cancel = context.WithTimeout(ctx, timeout)
		}
		if cancel != nil {
			defer cancel()
		}

		var resp interface{}
		var err error
		var lock sync.Mutex
		done := make(chan struct{})
		// create channel with buffer size 1 to avoid goroutine leak
		panicChan := make(chan interface{}, 1)
		go func() {
			defer func() {
				if p := recover(); p != nil {
					// attach call stack to avoid missing in different goroutine
					panicChan <- fmt.Sprintf("%+v\n\n%s", p, strings.TrimSpace(string(debug.Stack())))
				}
			}()

			lock.Lock()
			defer lock.Unlock()
			resp, err = handler(ctx, req)
			close(done)
		}()

		select {
		case p := <-panicChan:
			panic(p)
		case <-done:
			lock.Lock()
			defer lock.Unlock()
			return resp, err
		case <-ctx.Done():
			err := ctx.Err()
			if err == context.Canceled {
				err = status.Error(codes.Canceled, err.Error())
			} else if err == context.DeadlineExceeded {
				err = status.Error(codes.DeadlineExceeded, err.Error())
			}
			return nil, err
		}
	}
}
