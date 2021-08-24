// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package eventsourcedentity

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

// EventSourcedTckModelClient is the client API for EventSourcedTckModel service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventSourcedTckModelClient interface {
	Process(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type eventSourcedTckModelClient struct {
	cc grpc.ClientConnInterface
}

func NewEventSourcedTckModelClient(cc grpc.ClientConnInterface) EventSourcedTckModelClient {
	return &eventSourcedTckModelClient{cc}
}

func (c *eventSourcedTckModelClient) Process(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/akkaserverless.tck.model.eventsourcedentity.EventSourcedTckModel/Process", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventSourcedTckModelServer is the server API for EventSourcedTckModel service.
// All implementations must embed UnimplementedEventSourcedTckModelServer
// for forward compatibility
type EventSourcedTckModelServer interface {
	Process(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedEventSourcedTckModelServer()
}

// UnimplementedEventSourcedTckModelServer must be embedded to have forward compatible implementations.
type UnimplementedEventSourcedTckModelServer struct {
}

func (UnimplementedEventSourcedTckModelServer) Process(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Process not implemented")
}
func (UnimplementedEventSourcedTckModelServer) mustEmbedUnimplementedEventSourcedTckModelServer() {}

// UnsafeEventSourcedTckModelServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventSourcedTckModelServer will
// result in compilation errors.
type UnsafeEventSourcedTckModelServer interface {
	mustEmbedUnimplementedEventSourcedTckModelServer()
}

func RegisterEventSourcedTckModelServer(s grpc.ServiceRegistrar, srv EventSourcedTckModelServer) {
	s.RegisterService(&EventSourcedTckModel_ServiceDesc, srv)
}

func _EventSourcedTckModel_Process_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventSourcedTckModelServer).Process(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/akkaserverless.tck.model.eventsourcedentity.EventSourcedTckModel/Process",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventSourcedTckModelServer).Process(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// EventSourcedTckModel_ServiceDesc is the grpc.ServiceDesc for EventSourcedTckModel service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventSourcedTckModel_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "akkaserverless.tck.model.eventsourcedentity.EventSourcedTckModel",
	HandlerType: (*EventSourcedTckModelServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Process",
			Handler:    _EventSourcedTckModel_Process_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eventsourcedentity/event_sourced_entity.proto",
}

// EventSourcedTwoClient is the client API for EventSourcedTwo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventSourcedTwoClient interface {
	Call(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type eventSourcedTwoClient struct {
	cc grpc.ClientConnInterface
}

func NewEventSourcedTwoClient(cc grpc.ClientConnInterface) EventSourcedTwoClient {
	return &eventSourcedTwoClient{cc}
}

func (c *eventSourcedTwoClient) Call(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/akkaserverless.tck.model.eventsourcedentity.EventSourcedTwo/Call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventSourcedTwoServer is the server API for EventSourcedTwo service.
// All implementations must embed UnimplementedEventSourcedTwoServer
// for forward compatibility
type EventSourcedTwoServer interface {
	Call(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedEventSourcedTwoServer()
}

// UnimplementedEventSourcedTwoServer must be embedded to have forward compatible implementations.
type UnimplementedEventSourcedTwoServer struct {
}

func (UnimplementedEventSourcedTwoServer) Call(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}
func (UnimplementedEventSourcedTwoServer) mustEmbedUnimplementedEventSourcedTwoServer() {}

// UnsafeEventSourcedTwoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventSourcedTwoServer will
// result in compilation errors.
type UnsafeEventSourcedTwoServer interface {
	mustEmbedUnimplementedEventSourcedTwoServer()
}

func RegisterEventSourcedTwoServer(s grpc.ServiceRegistrar, srv EventSourcedTwoServer) {
	s.RegisterService(&EventSourcedTwo_ServiceDesc, srv)
}

func _EventSourcedTwo_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventSourcedTwoServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/akkaserverless.tck.model.eventsourcedentity.EventSourcedTwo/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventSourcedTwoServer).Call(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// EventSourcedTwo_ServiceDesc is the grpc.ServiceDesc for EventSourcedTwo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventSourcedTwo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "akkaserverless.tck.model.eventsourcedentity.EventSourcedTwo",
	HandlerType: (*EventSourcedTwoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _EventSourcedTwo_Call_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eventsourcedentity/event_sourced_entity.proto",
}

// EventSourcedConfiguredClient is the client API for EventSourcedConfigured service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventSourcedConfiguredClient interface {
	Call(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type eventSourcedConfiguredClient struct {
	cc grpc.ClientConnInterface
}

func NewEventSourcedConfiguredClient(cc grpc.ClientConnInterface) EventSourcedConfiguredClient {
	return &eventSourcedConfiguredClient{cc}
}

func (c *eventSourcedConfiguredClient) Call(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/akkaserverless.tck.model.eventsourcedentity.EventSourcedConfigured/Call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventSourcedConfiguredServer is the server API for EventSourcedConfigured service.
// All implementations must embed UnimplementedEventSourcedConfiguredServer
// for forward compatibility
type EventSourcedConfiguredServer interface {
	Call(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedEventSourcedConfiguredServer()
}

// UnimplementedEventSourcedConfiguredServer must be embedded to have forward compatible implementations.
type UnimplementedEventSourcedConfiguredServer struct {
}

func (UnimplementedEventSourcedConfiguredServer) Call(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}
func (UnimplementedEventSourcedConfiguredServer) mustEmbedUnimplementedEventSourcedConfiguredServer() {
}

// UnsafeEventSourcedConfiguredServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventSourcedConfiguredServer will
// result in compilation errors.
type UnsafeEventSourcedConfiguredServer interface {
	mustEmbedUnimplementedEventSourcedConfiguredServer()
}

func RegisterEventSourcedConfiguredServer(s grpc.ServiceRegistrar, srv EventSourcedConfiguredServer) {
	s.RegisterService(&EventSourcedConfigured_ServiceDesc, srv)
}

func _EventSourcedConfigured_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventSourcedConfiguredServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/akkaserverless.tck.model.eventsourcedentity.EventSourcedConfigured/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventSourcedConfiguredServer).Call(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// EventSourcedConfigured_ServiceDesc is the grpc.ServiceDesc for EventSourcedConfigured service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventSourcedConfigured_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "akkaserverless.tck.model.eventsourcedentity.EventSourcedConfigured",
	HandlerType: (*EventSourcedConfiguredServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _EventSourcedConfigured_Call_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eventsourcedentity/event_sourced_entity.proto",
}
