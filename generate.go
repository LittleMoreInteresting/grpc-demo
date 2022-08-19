package main

//go:generate protoc -I=./hello/proto/ --go_out=./hello/ ./hello/proto/helloworld.proto  --go-grpc_out=./hello/
