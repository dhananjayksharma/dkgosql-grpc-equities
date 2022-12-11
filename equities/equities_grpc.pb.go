// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package equities

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

// OrderClient is the client API for Order service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderClient interface {
	// Sends a order
	ProcessOrder(ctx context.Context, opts ...grpc.CallOption) (Order_ProcessOrderClient, error)
}

type orderClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderClient(cc grpc.ClientConnInterface) OrderClient {
	return &orderClient{cc}
}

func (c *orderClient) ProcessOrder(ctx context.Context, opts ...grpc.CallOption) (Order_ProcessOrderClient, error) {
	stream, err := c.cc.NewStream(ctx, &Order_ServiceDesc.Streams[0], "/equities.Order/ProcessOrder", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderProcessOrderClient{stream}
	return x, nil
}

type Order_ProcessOrderClient interface {
	Send(*OrderRequest) error
	Recv() (*OrderResponse, error)
	grpc.ClientStream
}

type orderProcessOrderClient struct {
	grpc.ClientStream
}

func (x *orderProcessOrderClient) Send(m *OrderRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *orderProcessOrderClient) Recv() (*OrderResponse, error) {
	m := new(OrderResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OrderServer is the server API for Order service.
// All implementations must embed UnimplementedOrderServer
// for forward compatibility
type OrderServer interface {
	// Sends a order
	ProcessOrder(Order_ProcessOrderServer) error
	mustEmbedUnimplementedOrderServer()
}

// UnimplementedOrderServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServer struct {
}

func (UnimplementedOrderServer) ProcessOrder(Order_ProcessOrderServer) error {
	return status.Errorf(codes.Unimplemented, "method ProcessOrder not implemented")
}
func (UnimplementedOrderServer) mustEmbedUnimplementedOrderServer() {}

// UnsafeOrderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServer will
// result in compilation errors.
type UnsafeOrderServer interface {
	mustEmbedUnimplementedOrderServer()
}

func RegisterOrderServer(s grpc.ServiceRegistrar, srv OrderServer) {
	s.RegisterService(&Order_ServiceDesc, srv)
}

func _Order_ProcessOrder_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OrderServer).ProcessOrder(&orderProcessOrderServer{stream})
}

type Order_ProcessOrderServer interface {
	Send(*OrderResponse) error
	Recv() (*OrderRequest, error)
	grpc.ServerStream
}

type orderProcessOrderServer struct {
	grpc.ServerStream
}

func (x *orderProcessOrderServer) Send(m *OrderResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *orderProcessOrderServer) Recv() (*OrderRequest, error) {
	m := new(OrderRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Order_ServiceDesc is the grpc.ServiceDesc for Order service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Order_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "equities.Order",
	HandlerType: (*OrderServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ProcessOrder",
			Handler:       _Order_ProcessOrder_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "equities/equities.proto",
}
