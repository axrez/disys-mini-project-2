// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package disys_mini_project_2

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

// ChittyChatClient is the client API for ChittyChat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChittyChatClient interface {
	Publish(ctx context.Context, in *PublishMessage, opts ...grpc.CallOption) (*EmptyReturn, error)
	Subscribe(ctx context.Context, in *SubscribeInputMessasge, opts ...grpc.CallOption) (ChittyChat_SubscribeClient, error)
	Join(ctx context.Context, in *JoinMessage, opts ...grpc.CallOption) (*JoinReplyMessage, error)
	Leave(ctx context.Context, in *LeaveMessage, opts ...grpc.CallOption) (*EmptyReturn, error)
}

type chittyChatClient struct {
	cc grpc.ClientConnInterface
}

func NewChittyChatClient(cc grpc.ClientConnInterface) ChittyChatClient {
	return &chittyChatClient{cc}
}

func (c *chittyChatClient) Publish(ctx context.Context, in *PublishMessage, opts ...grpc.CallOption) (*EmptyReturn, error) {
	out := new(EmptyReturn)
	err := c.cc.Invoke(ctx, "/main.ChittyChat/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chittyChatClient) Subscribe(ctx context.Context, in *SubscribeInputMessasge, opts ...grpc.CallOption) (ChittyChat_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChittyChat_ServiceDesc.Streams[0], "/main.ChittyChat/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &chittyChatSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChittyChat_SubscribeClient interface {
	Recv() (*BroadcastMessage, error)
	grpc.ClientStream
}

type chittyChatSubscribeClient struct {
	grpc.ClientStream
}

func (x *chittyChatSubscribeClient) Recv() (*BroadcastMessage, error) {
	m := new(BroadcastMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chittyChatClient) Join(ctx context.Context, in *JoinMessage, opts ...grpc.CallOption) (*JoinReplyMessage, error) {
	out := new(JoinReplyMessage)
	err := c.cc.Invoke(ctx, "/main.ChittyChat/Join", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chittyChatClient) Leave(ctx context.Context, in *LeaveMessage, opts ...grpc.CallOption) (*EmptyReturn, error) {
	out := new(EmptyReturn)
	err := c.cc.Invoke(ctx, "/main.ChittyChat/Leave", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChittyChatServer is the server API for ChittyChat service.
// All implementations must embed UnimplementedChittyChatServer
// for forward compatibility
type ChittyChatServer interface {
	Publish(context.Context, *PublishMessage) (*EmptyReturn, error)
	Subscribe(*SubscribeInputMessasge, ChittyChat_SubscribeServer) error
	Join(context.Context, *JoinMessage) (*JoinReplyMessage, error)
	Leave(context.Context, *LeaveMessage) (*EmptyReturn, error)
	mustEmbedUnimplementedChittyChatServer()
}

// UnimplementedChittyChatServer must be embedded to have forward compatible implementations.
type UnimplementedChittyChatServer struct {
}

func (UnimplementedChittyChatServer) Publish(context.Context, *PublishMessage) (*EmptyReturn, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedChittyChatServer) Subscribe(*SubscribeInputMessasge, ChittyChat_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedChittyChatServer) Join(context.Context, *JoinMessage) (*JoinReplyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Join not implemented")
}
func (UnimplementedChittyChatServer) Leave(context.Context, *LeaveMessage) (*EmptyReturn, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Leave not implemented")
}
func (UnimplementedChittyChatServer) mustEmbedUnimplementedChittyChatServer() {}

// UnsafeChittyChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChittyChatServer will
// result in compilation errors.
type UnsafeChittyChatServer interface {
	mustEmbedUnimplementedChittyChatServer()
}

func RegisterChittyChatServer(s grpc.ServiceRegistrar, srv ChittyChatServer) {
	s.RegisterService(&ChittyChat_ServiceDesc, srv)
}

func _ChittyChat_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChittyChatServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.ChittyChat/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChittyChatServer).Publish(ctx, req.(*PublishMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChittyChat_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeInputMessasge)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChittyChatServer).Subscribe(m, &chittyChatSubscribeServer{stream})
}

type ChittyChat_SubscribeServer interface {
	Send(*BroadcastMessage) error
	grpc.ServerStream
}

type chittyChatSubscribeServer struct {
	grpc.ServerStream
}

func (x *chittyChatSubscribeServer) Send(m *BroadcastMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _ChittyChat_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChittyChatServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.ChittyChat/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChittyChatServer).Join(ctx, req.(*JoinMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChittyChat_Leave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChittyChatServer).Leave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.ChittyChat/Leave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChittyChatServer).Leave(ctx, req.(*LeaveMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// ChittyChat_ServiceDesc is the grpc.ServiceDesc for ChittyChat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChittyChat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.ChittyChat",
	HandlerType: (*ChittyChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _ChittyChat_Publish_Handler,
		},
		{
			MethodName: "Join",
			Handler:    _ChittyChat_Join_Handler,
		},
		{
			MethodName: "Leave",
			Handler:    _ChittyChat_Leave_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _ChittyChat_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "chittychat.proto",
}
