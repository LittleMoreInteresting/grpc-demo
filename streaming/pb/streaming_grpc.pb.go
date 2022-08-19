// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: streaming.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StreamingClient is the client API for Streaming service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamingClient interface {
	ServerStream(ctx context.Context, in *Request, opts ...grpc.CallOption) (Streaming_ServerStreamClient, error)
	ClientStream(ctx context.Context, opts ...grpc.CallOption) (Streaming_ClientStreamClient, error)
	Bidirectional(ctx context.Context, opts ...grpc.CallOption) (Streaming_BidirectionalClient, error)
}

type streamingClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamingClient(cc grpc.ClientConnInterface) StreamingClient {
	return &streamingClient{cc}
}

func (c *streamingClient) ServerStream(ctx context.Context, in *Request, opts ...grpc.CallOption) (Streaming_ServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Streaming_ServiceDesc.Streams[0], "/pb.Streaming/ServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamingServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Streaming_ServerStreamClient interface {
	Recv() (*Reply, error)
	grpc.ClientStream
}

type streamingServerStreamClient struct {
	grpc.ClientStream
}

func (x *streamingServerStreamClient) Recv() (*Reply, error) {
	m := new(Reply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamingClient) ClientStream(ctx context.Context, opts ...grpc.CallOption) (Streaming_ClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Streaming_ServiceDesc.Streams[1], "/pb.Streaming/ClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamingClientStreamClient{stream}
	return x, nil
}

type Streaming_ClientStreamClient interface {
	Send(*Request) error
	CloseAndRecv() (*Reply, error)
	grpc.ClientStream
}

type streamingClientStreamClient struct {
	grpc.ClientStream
}

func (x *streamingClientStreamClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamingClientStreamClient) CloseAndRecv() (*Reply, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Reply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamingClient) Bidirectional(ctx context.Context, opts ...grpc.CallOption) (Streaming_BidirectionalClient, error) {
	stream, err := c.cc.NewStream(ctx, &Streaming_ServiceDesc.Streams[2], "/pb.Streaming/Bidirectional", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamingBidirectionalClient{stream}
	return x, nil
}

type Streaming_BidirectionalClient interface {
	Send(*Request) error
	Recv() (*Reply, error)
	grpc.ClientStream
}

type streamingBidirectionalClient struct {
	grpc.ClientStream
}

func (x *streamingBidirectionalClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamingBidirectionalClient) Recv() (*Reply, error) {
	m := new(Reply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamingServer is the server API for Streaming service.
// All implementations must embed UnimplementedStreamingServer
// for forward compatibility
type StreamingServer interface {
	ServerStream(*Request, Streaming_ServerStreamServer) error
	ClientStream(Streaming_ClientStreamServer) error
	Bidirectional(Streaming_BidirectionalServer) error
	mustEmbedUnimplementedStreamingServer()
}

// UnimplementedStreamingServer must be embedded to have forward compatible implementations.
type UnimplementedStreamingServer struct {
}

func (UnimplementedStreamingServer) ServerStream(*Request, Streaming_ServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerStream not implemented")
}
func (UnimplementedStreamingServer) ClientStream(Streaming_ClientStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientStream not implemented")
}
func (UnimplementedStreamingServer) Bidirectional(Streaming_BidirectionalServer) error {
	return status.Errorf(codes.Unimplemented, "method Bidirectional not implemented")
}
func (UnimplementedStreamingServer) mustEmbedUnimplementedStreamingServer() {}

// UnsafeStreamingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamingServer will
// result in compilation errors.
type UnsafeStreamingServer interface {
	mustEmbedUnimplementedStreamingServer()
}

func RegisterStreamingServer(s grpc.ServiceRegistrar, srv StreamingServer) {
	s.RegisterService(&Streaming_ServiceDesc, srv)
}

func _Streaming_ServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamingServer).ServerStream(m, &streamingServerStreamServer{stream})
}

type Streaming_ServerStreamServer interface {
	Send(*Reply) error
	grpc.ServerStream
}

type streamingServerStreamServer struct {
	grpc.ServerStream
}

func (x *streamingServerStreamServer) Send(m *Reply) error {
	return x.ServerStream.SendMsg(m)
}

func _Streaming_ClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamingServer).ClientStream(&streamingClientStreamServer{stream})
}

type Streaming_ClientStreamServer interface {
	SendAndClose(*Reply) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type streamingClientStreamServer struct {
	grpc.ServerStream
}

func (x *streamingClientStreamServer) SendAndClose(m *Reply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamingClientStreamServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Streaming_Bidirectional_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamingServer).Bidirectional(&streamingBidirectionalServer{stream})
}

type Streaming_BidirectionalServer interface {
	Send(*Reply) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type streamingBidirectionalServer struct {
	grpc.ServerStream
}

func (x *streamingBidirectionalServer) Send(m *Reply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamingBidirectionalServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Streaming_ServiceDesc is the grpc.ServiceDesc for Streaming service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Streaming_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Streaming",
	HandlerType: (*StreamingServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerStream",
			Handler:       _Streaming_ServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ClientStream",
			Handler:       _Streaming_ClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Bidirectional",
			Handler:       _Streaming_Bidirectional_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "streaming.proto",
}