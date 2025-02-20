// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// WoodpeckerClient is the client API for Woodpecker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WoodpeckerClient interface {
	Next(ctx context.Context, in *NextRequest, opts ...grpc.CallOption) (*NextReply, error)
	Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*Empty, error)
	Wait(ctx context.Context, in *WaitRequest, opts ...grpc.CallOption) (*Empty, error)
	Done(ctx context.Context, in *DoneRequest, opts ...grpc.CallOption) (*Empty, error)
	Extend(ctx context.Context, in *ExtendRequest, opts ...grpc.CallOption) (*Empty, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Empty, error)
	Upload(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*Empty, error)
	Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*Empty, error)
}

type woodpeckerClient struct {
	cc grpc.ClientConnInterface
}

func NewWoodpeckerClient(cc grpc.ClientConnInterface) WoodpeckerClient {
	return &woodpeckerClient{cc}
}

func (c *woodpeckerClient) Next(ctx context.Context, in *NextRequest, opts ...grpc.CallOption) (*NextReply, error) {
	out := new(NextReply)
	err := c.cc.Invoke(ctx, "/proto.Woodpecker/Next", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *woodpeckerClient) Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.Woodpecker/Init", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *woodpeckerClient) Wait(ctx context.Context, in *WaitRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.Woodpecker/Wait", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *woodpeckerClient) Done(ctx context.Context, in *DoneRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.Woodpecker/Done", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *woodpeckerClient) Extend(ctx context.Context, in *ExtendRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.Woodpecker/Extend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *woodpeckerClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.Woodpecker/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *woodpeckerClient) Upload(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.Woodpecker/Upload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *woodpeckerClient) Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.Woodpecker/Log", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WoodpeckerServer is the server API for Woodpecker service.
// All implementations must embed UnimplementedWoodpeckerServer
// for forward compatibility
type WoodpeckerServer interface {
	Next(context.Context, *NextRequest) (*NextReply, error)
	Init(context.Context, *InitRequest) (*Empty, error)
	Wait(context.Context, *WaitRequest) (*Empty, error)
	Done(context.Context, *DoneRequest) (*Empty, error)
	Extend(context.Context, *ExtendRequest) (*Empty, error)
	Update(context.Context, *UpdateRequest) (*Empty, error)
	Upload(context.Context, *UploadRequest) (*Empty, error)
	Log(context.Context, *LogRequest) (*Empty, error)
	mustEmbedUnimplementedWoodpeckerServer()
}

// UnimplementedWoodpeckerServer must be embedded to have forward compatible implementations.
type UnimplementedWoodpeckerServer struct {
}

func (UnimplementedWoodpeckerServer) Next(context.Context, *NextRequest) (*NextReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Next not implemented")
}
func (UnimplementedWoodpeckerServer) Init(context.Context, *InitRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (UnimplementedWoodpeckerServer) Wait(context.Context, *WaitRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Wait not implemented")
}
func (UnimplementedWoodpeckerServer) Done(context.Context, *DoneRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Done not implemented")
}
func (UnimplementedWoodpeckerServer) Extend(context.Context, *ExtendRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Extend not implemented")
}
func (UnimplementedWoodpeckerServer) Update(context.Context, *UpdateRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedWoodpeckerServer) Upload(context.Context, *UploadRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedWoodpeckerServer) Log(context.Context, *LogRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Log not implemented")
}
func (UnimplementedWoodpeckerServer) mustEmbedUnimplementedWoodpeckerServer() {}

// UnsafeWoodpeckerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WoodpeckerServer will
// result in compilation errors.
type UnsafeWoodpeckerServer interface {
	mustEmbedUnimplementedWoodpeckerServer()
}

func RegisterWoodpeckerServer(s grpc.ServiceRegistrar, srv WoodpeckerServer) {
	s.RegisterService(&Woodpecker_ServiceDesc, srv)
}

func _Woodpecker_Next_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WoodpeckerServer).Next(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Woodpecker/Next",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WoodpeckerServer).Next(ctx, req.(*NextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Woodpecker_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WoodpeckerServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Woodpecker/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WoodpeckerServer).Init(ctx, req.(*InitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Woodpecker_Wait_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WaitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WoodpeckerServer).Wait(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Woodpecker/Wait",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WoodpeckerServer).Wait(ctx, req.(*WaitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Woodpecker_Done_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WoodpeckerServer).Done(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Woodpecker/Done",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WoodpeckerServer).Done(ctx, req.(*DoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Woodpecker_Extend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExtendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WoodpeckerServer).Extend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Woodpecker/Extend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WoodpeckerServer).Extend(ctx, req.(*ExtendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Woodpecker_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WoodpeckerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Woodpecker/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WoodpeckerServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Woodpecker_Upload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WoodpeckerServer).Upload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Woodpecker/Upload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WoodpeckerServer).Upload(ctx, req.(*UploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Woodpecker_Log_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WoodpeckerServer).Log(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Woodpecker/Log",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WoodpeckerServer).Log(ctx, req.(*LogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Woodpecker_ServiceDesc is the grpc.ServiceDesc for Woodpecker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Woodpecker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Woodpecker",
	HandlerType: (*WoodpeckerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Next",
			Handler:    _Woodpecker_Next_Handler,
		},
		{
			MethodName: "Init",
			Handler:    _Woodpecker_Init_Handler,
		},
		{
			MethodName: "Wait",
			Handler:    _Woodpecker_Wait_Handler,
		},
		{
			MethodName: "Done",
			Handler:    _Woodpecker_Done_Handler,
		},
		{
			MethodName: "Extend",
			Handler:    _Woodpecker_Extend_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Woodpecker_Update_Handler,
		},
		{
			MethodName: "Upload",
			Handler:    _Woodpecker_Upload_Handler,
		},
		{
			MethodName: "Log",
			Handler:    _Woodpecker_Log_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "woodpecker.proto",
}

// HealthClient is the client API for Health service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HealthClient interface {
	Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
}

type healthClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthClient(cc grpc.ClientConnInterface) HealthClient {
	return &healthClient{cc}
}

func (c *healthClient) Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/proto.Health/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthServer is the server API for Health service.
// All implementations must embed UnimplementedHealthServer
// for forward compatibility
type HealthServer interface {
	Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
	mustEmbedUnimplementedHealthServer()
}

// UnimplementedHealthServer must be embedded to have forward compatible implementations.
type UnimplementedHealthServer struct {
}

func (UnimplementedHealthServer) Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (UnimplementedHealthServer) mustEmbedUnimplementedHealthServer() {}

// UnsafeHealthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HealthServer will
// result in compilation errors.
type UnsafeHealthServer interface {
	mustEmbedUnimplementedHealthServer()
}

func RegisterHealthServer(s grpc.ServiceRegistrar, srv HealthServer) {
	s.RegisterService(&Health_ServiceDesc, srv)
}

func _Health_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Health/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).Check(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Health_ServiceDesc is the grpc.ServiceDesc for Health service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Health_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Health",
	HandlerType: (*HealthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Health_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "woodpecker.proto",
}
