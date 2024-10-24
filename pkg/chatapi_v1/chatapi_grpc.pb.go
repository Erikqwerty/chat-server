// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.28.2
// source: chatapi.proto

package chatapi_v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ChatAPIV1Client is the client API for ChatAPIV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatAPIV1Client interface {
	// Метод для создания нового чата
	CreateChat(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// Метод для удаления чата
	DeleteChat(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Метод для присоединения к существующему чату
	JoinChat(ctx context.Context, in *JoinChatRequest, opts ...grpc.CallOption) (*JoinChatResponse, error)
	// Метод для отправки сообщения в чат
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type chatAPIV1Client struct {
	cc grpc.ClientConnInterface
}

func NewChatAPIV1Client(cc grpc.ClientConnInterface) ChatAPIV1Client {
	return &chatAPIV1Client{cc}
}

func (c *chatAPIV1Client) CreateChat(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/chatapi_v1.ChatAPIV1/CreateChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatAPIV1Client) DeleteChat(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/chatapi_v1.ChatAPIV1/DeleteChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatAPIV1Client) JoinChat(ctx context.Context, in *JoinChatRequest, opts ...grpc.CallOption) (*JoinChatResponse, error) {
	out := new(JoinChatResponse)
	err := c.cc.Invoke(ctx, "/chatapi_v1.ChatAPIV1/JoinChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatAPIV1Client) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/chatapi_v1.ChatAPIV1/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatAPIV1Server is the server API for ChatAPIV1 service.
// All implementations must embed UnimplementedChatAPIV1Server
// for forward compatibility
type ChatAPIV1Server interface {
	// Метод для создания нового чата
	CreateChat(context.Context, *CreateRequest) (*CreateResponse, error)
	// Метод для удаления чата
	DeleteChat(context.Context, *DeleteRequest) (*emptypb.Empty, error)
	// Метод для присоединения к существующему чату
	JoinChat(context.Context, *JoinChatRequest) (*JoinChatResponse, error)
	// Метод для отправки сообщения в чат
	SendMessage(context.Context, *SendMessageRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedChatAPIV1Server()
}

// UnimplementedChatAPIV1Server must be embedded to have forward compatible implementations.
type UnimplementedChatAPIV1Server struct {
}

func (UnimplementedChatAPIV1Server) CreateChat(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChat not implemented")
}
func (UnimplementedChatAPIV1Server) DeleteChat(context.Context, *DeleteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteChat not implemented")
}
func (UnimplementedChatAPIV1Server) JoinChat(context.Context, *JoinChatRequest) (*JoinChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinChat not implemented")
}
func (UnimplementedChatAPIV1Server) SendMessage(context.Context, *SendMessageRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedChatAPIV1Server) mustEmbedUnimplementedChatAPIV1Server() {}

// UnsafeChatAPIV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatAPIV1Server will
// result in compilation errors.
type UnsafeChatAPIV1Server interface {
	mustEmbedUnimplementedChatAPIV1Server()
}

func RegisterChatAPIV1Server(s grpc.ServiceRegistrar, srv ChatAPIV1Server) {
	s.RegisterService(&ChatAPIV1_ServiceDesc, srv)
}

func _ChatAPIV1_CreateChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatAPIV1Server).CreateChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatapi_v1.ChatAPIV1/CreateChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatAPIV1Server).CreateChat(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatAPIV1_DeleteChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatAPIV1Server).DeleteChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatapi_v1.ChatAPIV1/DeleteChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatAPIV1Server).DeleteChat(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatAPIV1_JoinChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatAPIV1Server).JoinChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatapi_v1.ChatAPIV1/JoinChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatAPIV1Server).JoinChat(ctx, req.(*JoinChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatAPIV1_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatAPIV1Server).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatapi_v1.ChatAPIV1/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatAPIV1Server).SendMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatAPIV1_ServiceDesc is the grpc.ServiceDesc for ChatAPIV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatAPIV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chatapi_v1.ChatAPIV1",
	HandlerType: (*ChatAPIV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateChat",
			Handler:    _ChatAPIV1_CreateChat_Handler,
		},
		{
			MethodName: "DeleteChat",
			Handler:    _ChatAPIV1_DeleteChat_Handler,
		},
		{
			MethodName: "JoinChat",
			Handler:    _ChatAPIV1_JoinChat_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _ChatAPIV1_SendMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chatapi.proto",
}
