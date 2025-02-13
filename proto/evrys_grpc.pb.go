// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	pb "github.com/cloudevents/sdk-go/binding/format/protobuf/v2/pb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EvrysClient is the client API for Evrys service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EvrysClient interface {
	RecordEvent(ctx context.Context, in *pb.CloudEvent, opts ...grpc.CallOption) (*RecordEventResponse, error)
	GetEvent(ctx context.Context, in *GetEventRequest, opts ...grpc.CallOption) (*pb.CloudEvent, error)
	SliceEvents(ctx context.Context, in *SliceEventsRequest, opts ...grpc.CallOption) (Evrys_SliceEventsClient, error)
}

type evrysClient struct {
	cc grpc.ClientConnInterface
}

func NewEvrysClient(cc grpc.ClientConnInterface) EvrysClient {
	return &evrysClient{cc}
}

func (c *evrysClient) RecordEvent(ctx context.Context, in *pb.CloudEvent, opts ...grpc.CallOption) (*RecordEventResponse, error) {
	out := new(RecordEventResponse)
	err := c.cc.Invoke(ctx, "/evrys.Evrys/RecordEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evrysClient) GetEvent(ctx context.Context, in *GetEventRequest, opts ...grpc.CallOption) (*pb.CloudEvent, error) {
	out := new(pb.CloudEvent)
	err := c.cc.Invoke(ctx, "/evrys.Evrys/GetEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evrysClient) SliceEvents(ctx context.Context, in *SliceEventsRequest, opts ...grpc.CallOption) (Evrys_SliceEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Evrys_ServiceDesc.Streams[0], "/evrys.Evrys/SliceEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &evrysSliceEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Evrys_SliceEventsClient interface {
	Recv() (*pb.CloudEvent, error)
	grpc.ClientStream
}

type evrysSliceEventsClient struct {
	grpc.ClientStream
}

func (x *evrysSliceEventsClient) Recv() (*pb.CloudEvent, error) {
	m := new(pb.CloudEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EvrysServer is the server API for Evrys service.
// All implementations must embed UnimplementedEvrysServer
// for forward compatibility
type EvrysServer interface {
	RecordEvent(context.Context, *pb.CloudEvent) (*RecordEventResponse, error)
	GetEvent(context.Context, *GetEventRequest) (*pb.CloudEvent, error)
	SliceEvents(*SliceEventsRequest, Evrys_SliceEventsServer) error
	mustEmbedUnimplementedEvrysServer()
}

// UnimplementedEvrysServer must be embedded to have forward compatible implementations.
type UnimplementedEvrysServer struct {
}

func (UnimplementedEvrysServer) RecordEvent(context.Context, *pb.CloudEvent) (*RecordEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecordEvent not implemented")
}
func (UnimplementedEvrysServer) GetEvent(context.Context, *GetEventRequest) (*pb.CloudEvent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvent not implemented")
}
func (UnimplementedEvrysServer) SliceEvents(*SliceEventsRequest, Evrys_SliceEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method SliceEvents not implemented")
}
func (UnimplementedEvrysServer) mustEmbedUnimplementedEvrysServer() {}

// UnsafeEvrysServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EvrysServer will
// result in compilation errors.
type UnsafeEvrysServer interface {
	mustEmbedUnimplementedEvrysServer()
}

func RegisterEvrysServer(s grpc.ServiceRegistrar, srv EvrysServer) {
	s.RegisterService(&Evrys_ServiceDesc, srv)
}

func _Evrys_RecordEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.CloudEvent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvrysServer).RecordEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evrys.Evrys/RecordEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvrysServer).RecordEvent(ctx, req.(*pb.CloudEvent))
	}
	return interceptor(ctx, in, info, handler)
}

func _Evrys_GetEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvrysServer).GetEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evrys.Evrys/GetEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvrysServer).GetEvent(ctx, req.(*GetEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Evrys_SliceEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SliceEventsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EvrysServer).SliceEvents(m, &evrysSliceEventsServer{stream})
}

type Evrys_SliceEventsServer interface {
	Send(*pb.CloudEvent) error
	grpc.ServerStream
}

type evrysSliceEventsServer struct {
	grpc.ServerStream
}

func (x *evrysSliceEventsServer) Send(m *pb.CloudEvent) error {
	return x.ServerStream.SendMsg(m)
}

// Evrys_ServiceDesc is the grpc.ServiceDesc for Evrys service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Evrys_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "evrys.Evrys",
	HandlerType: (*EvrysServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RecordEvent",
			Handler:    _Evrys_RecordEvent_Handler,
		},
		{
			MethodName: "GetEvent",
			Handler:    _Evrys_GetEvent_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SliceEvents",
			Handler:       _Evrys_SliceEvents_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "evrys.proto",
}
