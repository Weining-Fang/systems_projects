// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.24.4
// source: gpu_sim.proto

package gpu_sim

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	GPUDevice_GetDeviceMetadata_FullMethodName = "/gpu_sim.GPUDevice/GetDeviceMetadata"
	GPUDevice_BeginSend_FullMethodName         = "/gpu_sim.GPUDevice/BeginSend"
	GPUDevice_BeginReceive_FullMethodName      = "/gpu_sim.GPUDevice/BeginReceive"
	GPUDevice_StreamSend_FullMethodName        = "/gpu_sim.GPUDevice/StreamSend"
	GPUDevice_StreamReceive_FullMethodName     = "/gpu_sim.GPUDevice/StreamReceive"
	GPUDevice_GetStreamStatus_FullMethodName   = "/gpu_sim.GPUDevice/GetStreamStatus"
)

// GPUDeviceClient is the client API for GPUDevice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// A service that simulates a single GPU device
type GPUDeviceClient interface {
	GetDeviceMetadata(ctx context.Context, in *GetDeviceMetadataRequest, opts ...grpc.CallOption) (*GetDeviceMetadataResponse, error)
	// Called by the GPUCoordinator to start the data transfer between two devices.
	// Begin.*() functions are "non-blocking", meaning they return immediately after initiating the operation.
	// The actual data transfer should happen in the background initiated by the devices.
	BeginSend(ctx context.Context, in *BeginSendRequest, opts ...grpc.CallOption) (*BeginSendResponse, error)
	BeginReceive(ctx context.Context, in *BeginReceiveRequest, opts ...grpc.CallOption) (*BeginReceiveResponse, error)
	// Called by the src device to send data to the dst device.
	StreamSend(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[DataChunk, StreamSendResponse], error)
	StreamReceive(ctx context.Context, in *StreamReceiveRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[DataChunk], error)
	// For the coordinator to know if a stream has completed.
	GetStreamStatus(ctx context.Context, in *GetStreamStatusRequest, opts ...grpc.CallOption) (*GetStreamStatusResponse, error)
}

type gPUDeviceClient struct {
	cc grpc.ClientConnInterface
}

func NewGPUDeviceClient(cc grpc.ClientConnInterface) GPUDeviceClient {
	return &gPUDeviceClient{cc}
}

func (c *gPUDeviceClient) GetDeviceMetadata(ctx context.Context, in *GetDeviceMetadataRequest, opts ...grpc.CallOption) (*GetDeviceMetadataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDeviceMetadataResponse)
	err := c.cc.Invoke(ctx, GPUDevice_GetDeviceMetadata_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gPUDeviceClient) BeginSend(ctx context.Context, in *BeginSendRequest, opts ...grpc.CallOption) (*BeginSendResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BeginSendResponse)
	err := c.cc.Invoke(ctx, GPUDevice_BeginSend_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gPUDeviceClient) BeginReceive(ctx context.Context, in *BeginReceiveRequest, opts ...grpc.CallOption) (*BeginReceiveResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BeginReceiveResponse)
	err := c.cc.Invoke(ctx, GPUDevice_BeginReceive_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gPUDeviceClient) StreamSend(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[DataChunk, StreamSendResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &GPUDevice_ServiceDesc.Streams[0], GPUDevice_StreamSend_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[DataChunk, StreamSendResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type GPUDevice_StreamSendClient = grpc.ClientStreamingClient[DataChunk, StreamSendResponse]

func (c *gPUDeviceClient) StreamReceive(ctx context.Context, in *StreamReceiveRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[DataChunk], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &GPUDevice_ServiceDesc.Streams[1], GPUDevice_StreamReceive_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[StreamReceiveRequest, DataChunk]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type GPUDevice_StreamReceiveClient = grpc.ServerStreamingClient[DataChunk]

func (c *gPUDeviceClient) GetStreamStatus(ctx context.Context, in *GetStreamStatusRequest, opts ...grpc.CallOption) (*GetStreamStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetStreamStatusResponse)
	err := c.cc.Invoke(ctx, GPUDevice_GetStreamStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GPUDeviceServer is the server API for GPUDevice service.
// All implementations should embed UnimplementedGPUDeviceServer
// for forward compatibility.
//
// A service that simulates a single GPU device
type GPUDeviceServer interface {
	GetDeviceMetadata(context.Context, *GetDeviceMetadataRequest) (*GetDeviceMetadataResponse, error)
	// Called by the GPUCoordinator to start the data transfer between two devices.
	// Begin.*() functions are "non-blocking", meaning they return immediately after initiating the operation.
	// The actual data transfer should happen in the background initiated by the devices.
	BeginSend(context.Context, *BeginSendRequest) (*BeginSendResponse, error)
	BeginReceive(context.Context, *BeginReceiveRequest) (*BeginReceiveResponse, error)
	// Called by the src device to send data to the dst device.
	StreamSend(grpc.ClientStreamingServer[DataChunk, StreamSendResponse]) error
	StreamReceive(*StreamReceiveRequest, grpc.ServerStreamingServer[DataChunk]) error
	// For the coordinator to know if a stream has completed.
	GetStreamStatus(context.Context, *GetStreamStatusRequest) (*GetStreamStatusResponse, error)
}

// UnimplementedGPUDeviceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGPUDeviceServer struct{}

func (UnimplementedGPUDeviceServer) GetDeviceMetadata(context.Context, *GetDeviceMetadataRequest) (*GetDeviceMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceMetadata not implemented")
}
func (UnimplementedGPUDeviceServer) BeginSend(context.Context, *BeginSendRequest) (*BeginSendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BeginSend not implemented")
}
func (UnimplementedGPUDeviceServer) BeginReceive(context.Context, *BeginReceiveRequest) (*BeginReceiveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BeginReceive not implemented")
}
func (UnimplementedGPUDeviceServer) StreamSend(grpc.ClientStreamingServer[DataChunk, StreamSendResponse]) error {
	return status.Errorf(codes.Unimplemented, "method StreamSend not implemented")
}
func (UnimplementedGPUDeviceServer) StreamReceive(*StreamReceiveRequest, grpc.ServerStreamingServer[DataChunk]) error {
	return status.Errorf(codes.Unimplemented, "method StreamReceive not implemented")
}
func (UnimplementedGPUDeviceServer) GetStreamStatus(context.Context, *GetStreamStatusRequest) (*GetStreamStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStreamStatus not implemented")
}
func (UnimplementedGPUDeviceServer) testEmbeddedByValue() {}

// UnsafeGPUDeviceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GPUDeviceServer will
// result in compilation errors.
type UnsafeGPUDeviceServer interface {
	mustEmbedUnimplementedGPUDeviceServer()
}

func RegisterGPUDeviceServer(s grpc.ServiceRegistrar, srv GPUDeviceServer) {
	// If the following call pancis, it indicates UnimplementedGPUDeviceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GPUDevice_ServiceDesc, srv)
}

func _GPUDevice_GetDeviceMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GPUDeviceServer).GetDeviceMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GPUDevice_GetDeviceMetadata_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GPUDeviceServer).GetDeviceMetadata(ctx, req.(*GetDeviceMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GPUDevice_BeginSend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BeginSendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GPUDeviceServer).BeginSend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GPUDevice_BeginSend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GPUDeviceServer).BeginSend(ctx, req.(*BeginSendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GPUDevice_BeginReceive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BeginReceiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GPUDeviceServer).BeginReceive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GPUDevice_BeginReceive_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GPUDeviceServer).BeginReceive(ctx, req.(*BeginReceiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GPUDevice_StreamSend_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GPUDeviceServer).StreamSend(&grpc.GenericServerStream[DataChunk, StreamSendResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type GPUDevice_StreamSendServer = grpc.ClientStreamingServer[DataChunk, StreamSendResponse]

func _GPUDevice_StreamReceive_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamReceiveRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GPUDeviceServer).StreamReceive(m, &grpc.GenericServerStream[StreamReceiveRequest, DataChunk]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type GPUDevice_StreamReceiveServer = grpc.ServerStreamingServer[DataChunk]

func _GPUDevice_GetStreamStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStreamStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GPUDeviceServer).GetStreamStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GPUDevice_GetStreamStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GPUDeviceServer).GetStreamStatus(ctx, req.(*GetStreamStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GPUDevice_ServiceDesc is the grpc.ServiceDesc for GPUDevice service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GPUDevice_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gpu_sim.GPUDevice",
	HandlerType: (*GPUDeviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDeviceMetadata",
			Handler:    _GPUDevice_GetDeviceMetadata_Handler,
		},
		{
			MethodName: "BeginSend",
			Handler:    _GPUDevice_BeginSend_Handler,
		},
		{
			MethodName: "BeginReceive",
			Handler:    _GPUDevice_BeginReceive_Handler,
		},
		{
			MethodName: "GetStreamStatus",
			Handler:    _GPUDevice_GetStreamStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamSend",
			Handler:       _GPUDevice_StreamSend_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamReceive",
			Handler:       _GPUDevice_StreamReceive_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "gpu_sim.proto",
}

const (
	GPUCoordinator_CommInit_FullMethodName      = "/gpu_sim.GPUCoordinator/CommInit"
	GPUCoordinator_GetCommStatus_FullMethodName = "/gpu_sim.GPUCoordinator/GetCommStatus"
	GPUCoordinator_CommFinalize_FullMethodName  = "/gpu_sim.GPUCoordinator/CommFinalize"
	GPUCoordinator_CommDestroy_FullMethodName   = "/gpu_sim.GPUCoordinator/CommDestroy"
	GPUCoordinator_GroupStart_FullMethodName    = "/gpu_sim.GPUCoordinator/GroupStart"
	GPUCoordinator_GroupEnd_FullMethodName      = "/gpu_sim.GPUCoordinator/GroupEnd"
	GPUCoordinator_AllReduceRing_FullMethodName = "/gpu_sim.GPUCoordinator/AllReduceRing"
	GPUCoordinator_Memcpy_FullMethodName        = "/gpu_sim.GPUCoordinator/Memcpy"
)

// GPUCoordinatorClient is the client API for GPUCoordinator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// A service that simulates a coordinator that manages multiple GPU devices
type GPUCoordinatorClient interface {
	CommInit(ctx context.Context, in *CommInitRequest, opts ...grpc.CallOption) (*CommInitResponse, error)
	GetCommStatus(ctx context.Context, in *GetCommStatusRequest, opts ...grpc.CallOption) (*GetCommStatusResponse, error)
	// You may choose to implement CommFinalize and CommDestroy RPCs
	CommFinalize(ctx context.Context, in *CommFinalizeRequest, opts ...grpc.CallOption) (*CommFinalizeResponse, error)
	CommDestroy(ctx context.Context, in *CommDestroyRequest, opts ...grpc.CallOption) (*CommDestroyResponse, error)
	// Group operations wrapper
	GroupStart(ctx context.Context, in *GroupStartRequest, opts ...grpc.CallOption) (*GroupStartResponse, error)
	GroupEnd(ctx context.Context, in *GroupEndRequest, opts ...grpc.CallOption) (*GroupEndResponse, error)
	// RPCs for group or peer-to-peer communication
	AllReduceRing(ctx context.Context, in *AllReduceRingRequest, opts ...grpc.CallOption) (*AllReduceRingResponse, error)
	// Host-to-device data transfer and vice versa
	// You may implement this as streaming as well
	Memcpy(ctx context.Context, in *MemcpyRequest, opts ...grpc.CallOption) (*MemcpyResponse, error)
}

type gPUCoordinatorClient struct {
	cc grpc.ClientConnInterface
}

func NewGPUCoordinatorClient(cc grpc.ClientConnInterface) GPUCoordinatorClient {
	return &gPUCoordinatorClient{cc}
}

func (c *gPUCoordinatorClient) CommInit(ctx context.Context, in *CommInitRequest, opts ...grpc.CallOption) (*CommInitResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommInitResponse)
	err := c.cc.Invoke(ctx, GPUCoordinator_CommInit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gPUCoordinatorClient) GetCommStatus(ctx context.Context, in *GetCommStatusRequest, opts ...grpc.CallOption) (*GetCommStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCommStatusResponse)
	err := c.cc.Invoke(ctx, GPUCoordinator_GetCommStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gPUCoordinatorClient) CommFinalize(ctx context.Context, in *CommFinalizeRequest, opts ...grpc.CallOption) (*CommFinalizeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommFinalizeResponse)
	err := c.cc.Invoke(ctx, GPUCoordinator_CommFinalize_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gPUCoordinatorClient) CommDestroy(ctx context.Context, in *CommDestroyRequest, opts ...grpc.CallOption) (*CommDestroyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommDestroyResponse)
	err := c.cc.Invoke(ctx, GPUCoordinator_CommDestroy_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gPUCoordinatorClient) GroupStart(ctx context.Context, in *GroupStartRequest, opts ...grpc.CallOption) (*GroupStartResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GroupStartResponse)
	err := c.cc.Invoke(ctx, GPUCoordinator_GroupStart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gPUCoordinatorClient) GroupEnd(ctx context.Context, in *GroupEndRequest, opts ...grpc.CallOption) (*GroupEndResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GroupEndResponse)
	err := c.cc.Invoke(ctx, GPUCoordinator_GroupEnd_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gPUCoordinatorClient) AllReduceRing(ctx context.Context, in *AllReduceRingRequest, opts ...grpc.CallOption) (*AllReduceRingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AllReduceRingResponse)
	err := c.cc.Invoke(ctx, GPUCoordinator_AllReduceRing_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gPUCoordinatorClient) Memcpy(ctx context.Context, in *MemcpyRequest, opts ...grpc.CallOption) (*MemcpyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MemcpyResponse)
	err := c.cc.Invoke(ctx, GPUCoordinator_Memcpy_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GPUCoordinatorServer is the server API for GPUCoordinator service.
// All implementations should embed UnimplementedGPUCoordinatorServer
// for forward compatibility.
//
// A service that simulates a coordinator that manages multiple GPU devices
type GPUCoordinatorServer interface {
	CommInit(context.Context, *CommInitRequest) (*CommInitResponse, error)
	GetCommStatus(context.Context, *GetCommStatusRequest) (*GetCommStatusResponse, error)
	// You may choose to implement CommFinalize and CommDestroy RPCs
	CommFinalize(context.Context, *CommFinalizeRequest) (*CommFinalizeResponse, error)
	CommDestroy(context.Context, *CommDestroyRequest) (*CommDestroyResponse, error)
	// Group operations wrapper
	GroupStart(context.Context, *GroupStartRequest) (*GroupStartResponse, error)
	GroupEnd(context.Context, *GroupEndRequest) (*GroupEndResponse, error)
	// RPCs for group or peer-to-peer communication
	AllReduceRing(context.Context, *AllReduceRingRequest) (*AllReduceRingResponse, error)
	// Host-to-device data transfer and vice versa
	// You may implement this as streaming as well
	Memcpy(context.Context, *MemcpyRequest) (*MemcpyResponse, error)
}

// UnimplementedGPUCoordinatorServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGPUCoordinatorServer struct{}

func (UnimplementedGPUCoordinatorServer) CommInit(context.Context, *CommInitRequest) (*CommInitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommInit not implemented")
}
func (UnimplementedGPUCoordinatorServer) GetCommStatus(context.Context, *GetCommStatusRequest) (*GetCommStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommStatus not implemented")
}
func (UnimplementedGPUCoordinatorServer) CommFinalize(context.Context, *CommFinalizeRequest) (*CommFinalizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommFinalize not implemented")
}
func (UnimplementedGPUCoordinatorServer) CommDestroy(context.Context, *CommDestroyRequest) (*CommDestroyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommDestroy not implemented")
}
func (UnimplementedGPUCoordinatorServer) GroupStart(context.Context, *GroupStartRequest) (*GroupStartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupStart not implemented")
}
func (UnimplementedGPUCoordinatorServer) GroupEnd(context.Context, *GroupEndRequest) (*GroupEndResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupEnd not implemented")
}
func (UnimplementedGPUCoordinatorServer) AllReduceRing(context.Context, *AllReduceRingRequest) (*AllReduceRingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllReduceRing not implemented")
}
func (UnimplementedGPUCoordinatorServer) Memcpy(context.Context, *MemcpyRequest) (*MemcpyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Memcpy not implemented")
}
func (UnimplementedGPUCoordinatorServer) testEmbeddedByValue() {}

// UnsafeGPUCoordinatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GPUCoordinatorServer will
// result in compilation errors.
type UnsafeGPUCoordinatorServer interface {
	mustEmbedUnimplementedGPUCoordinatorServer()
}

func RegisterGPUCoordinatorServer(s grpc.ServiceRegistrar, srv GPUCoordinatorServer) {
	// If the following call pancis, it indicates UnimplementedGPUCoordinatorServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GPUCoordinator_ServiceDesc, srv)
}

func _GPUCoordinator_CommInit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommInitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GPUCoordinatorServer).CommInit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GPUCoordinator_CommInit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GPUCoordinatorServer).CommInit(ctx, req.(*CommInitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GPUCoordinator_GetCommStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GPUCoordinatorServer).GetCommStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GPUCoordinator_GetCommStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GPUCoordinatorServer).GetCommStatus(ctx, req.(*GetCommStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GPUCoordinator_CommFinalize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommFinalizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GPUCoordinatorServer).CommFinalize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GPUCoordinator_CommFinalize_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GPUCoordinatorServer).CommFinalize(ctx, req.(*CommFinalizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GPUCoordinator_CommDestroy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommDestroyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GPUCoordinatorServer).CommDestroy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GPUCoordinator_CommDestroy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GPUCoordinatorServer).CommDestroy(ctx, req.(*CommDestroyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GPUCoordinator_GroupStart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupStartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GPUCoordinatorServer).GroupStart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GPUCoordinator_GroupStart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GPUCoordinatorServer).GroupStart(ctx, req.(*GroupStartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GPUCoordinator_GroupEnd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupEndRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GPUCoordinatorServer).GroupEnd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GPUCoordinator_GroupEnd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GPUCoordinatorServer).GroupEnd(ctx, req.(*GroupEndRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GPUCoordinator_AllReduceRing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllReduceRingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GPUCoordinatorServer).AllReduceRing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GPUCoordinator_AllReduceRing_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GPUCoordinatorServer).AllReduceRing(ctx, req.(*AllReduceRingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GPUCoordinator_Memcpy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemcpyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GPUCoordinatorServer).Memcpy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GPUCoordinator_Memcpy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GPUCoordinatorServer).Memcpy(ctx, req.(*MemcpyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GPUCoordinator_ServiceDesc is the grpc.ServiceDesc for GPUCoordinator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GPUCoordinator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gpu_sim.GPUCoordinator",
	HandlerType: (*GPUCoordinatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CommInit",
			Handler:    _GPUCoordinator_CommInit_Handler,
		},
		{
			MethodName: "GetCommStatus",
			Handler:    _GPUCoordinator_GetCommStatus_Handler,
		},
		{
			MethodName: "CommFinalize",
			Handler:    _GPUCoordinator_CommFinalize_Handler,
		},
		{
			MethodName: "CommDestroy",
			Handler:    _GPUCoordinator_CommDestroy_Handler,
		},
		{
			MethodName: "GroupStart",
			Handler:    _GPUCoordinator_GroupStart_Handler,
		},
		{
			MethodName: "GroupEnd",
			Handler:    _GPUCoordinator_GroupEnd_Handler,
		},
		{
			MethodName: "AllReduceRing",
			Handler:    _GPUCoordinator_AllReduceRing_Handler,
		},
		{
			MethodName: "Memcpy",
			Handler:    _GPUCoordinator_Memcpy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gpu_sim.proto",
}
