package main

//hello
//go:generate protoc -I=./hello/proto/ --go_out=./hello/ ./hello/proto/helloworld.proto  --go-grpc_out=./hello/

//hello
//go:generate protoc -I=./streaming/pb/ --go_out=./streaming/ ./streaming/pb/streaming.proto  --go-grpc_out=./streaming/
