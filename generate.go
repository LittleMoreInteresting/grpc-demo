package main

//hello
//go:generate protoc -I=./hello/proto/ --go_out=./hello/ ./hello/proto/helloworld.proto  --go-grpc_out=./hello/

//streaming
//go:generate protoc -I=./streaming/pb/ --go_out=./streaming/ ./streaming/pb/streaming.proto  --go-grpc_out=./streaming/

//interceptor
//go:generate protoc -I=./interceptor/pb/ --go_out=./interceptor/ --go-grpc_out=./interceptor/ ./interceptor/pb/speaker.proto

//discover
//go:generate protoc -I=./discover/pb/ --go_out=./discover/ --go-grpc_out=./discover/ ./discover/pb/discover.proto

//gateway
//go:generate protoc -I=./gateway/pb/ -I=./third_party/ --go_out=./gateway/ --go-grpc_out=./gateway/ -grpc-gateway_out=./gateway/ ./gateway/pb/gateway.proto

//safety
//go:generate protoc -I=./safety/pb/  --go_out=./safety/ --go-grpc_out=./safety/ ./safety/pb/safety.proto
