// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.0--rc1
// source: proto/SparkService.proto

package grpc

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

const (
	AnalysisService_StartSparkJob_FullMethodName = "/AnalysisService/StartSparkJob"
	AnalysisService_StopSparkJob_FullMethodName  = "/AnalysisService/StopSparkJob"
	AnalysisService_CheckJob_FullMethodName      = "/AnalysisService/CheckJob"
)

// AnalysisServiceClient is the client API for AnalysisService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnalysisServiceClient interface {
	StartSparkJob(ctx context.Context, in *SparkJobRequest, opts ...grpc.CallOption) (*JobInfo, error)
	StopSparkJob(ctx context.Context, in *PingJob, opts ...grpc.CallOption) (*JobInfo, error)
	CheckJob(ctx context.Context, in *PingJob, opts ...grpc.CallOption) (*JobInfo, error)
}

type analysisServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAnalysisServiceClient(cc grpc.ClientConnInterface) AnalysisServiceClient {
	return &analysisServiceClient{cc}
}

func (c *analysisServiceClient) StartSparkJob(ctx context.Context, in *SparkJobRequest, opts ...grpc.CallOption) (*JobInfo, error) {
	out := new(JobInfo)
	err := c.cc.Invoke(ctx, AnalysisService_StartSparkJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analysisServiceClient) StopSparkJob(ctx context.Context, in *PingJob, opts ...grpc.CallOption) (*JobInfo, error) {
	out := new(JobInfo)
	err := c.cc.Invoke(ctx, AnalysisService_StopSparkJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analysisServiceClient) CheckJob(ctx context.Context, in *PingJob, opts ...grpc.CallOption) (*JobInfo, error) {
	out := new(JobInfo)
	err := c.cc.Invoke(ctx, AnalysisService_CheckJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnalysisServiceServer is the server API for AnalysisService service.
// All implementations must embed UnimplementedAnalysisServiceServer
// for forward compatibility
type AnalysisServiceServer interface {
	StartSparkJob(context.Context, *SparkJobRequest) (*JobInfo, error)
	StopSparkJob(context.Context, *PingJob) (*JobInfo, error)
	CheckJob(context.Context, *PingJob) (*JobInfo, error)
	mustEmbedUnimplementedAnalysisServiceServer()
}

// UnimplementedAnalysisServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAnalysisServiceServer struct {
}

func (UnimplementedAnalysisServiceServer) StartSparkJob(context.Context, *SparkJobRequest) (*JobInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartSparkJob not implemented")
}
func (UnimplementedAnalysisServiceServer) StopSparkJob(context.Context, *PingJob) (*JobInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopSparkJob not implemented")
}
func (UnimplementedAnalysisServiceServer) CheckJob(context.Context, *PingJob) (*JobInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckJob not implemented")
}
func (UnimplementedAnalysisServiceServer) mustEmbedUnimplementedAnalysisServiceServer() {}

// UnsafeAnalysisServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnalysisServiceServer will
// result in compilation errors.
type UnsafeAnalysisServiceServer interface {
	mustEmbedUnimplementedAnalysisServiceServer()
}

func RegisterAnalysisServiceServer(s grpc.ServiceRegistrar, srv AnalysisServiceServer) {
	s.RegisterService(&AnalysisService_ServiceDesc, srv)
}

func _AnalysisService_StartSparkJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SparkJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalysisServiceServer).StartSparkJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnalysisService_StartSparkJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalysisServiceServer).StartSparkJob(ctx, req.(*SparkJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnalysisService_StopSparkJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingJob)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalysisServiceServer).StopSparkJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnalysisService_StopSparkJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalysisServiceServer).StopSparkJob(ctx, req.(*PingJob))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnalysisService_CheckJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingJob)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalysisServiceServer).CheckJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnalysisService_CheckJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalysisServiceServer).CheckJob(ctx, req.(*PingJob))
	}
	return interceptor(ctx, in, info, handler)
}

// AnalysisService_ServiceDesc is the grpc.ServiceDesc for AnalysisService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AnalysisService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AnalysisService",
	HandlerType: (*AnalysisServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartSparkJob",
			Handler:    _AnalysisService_StartSparkJob_Handler,
		},
		{
			MethodName: "StopSparkJob",
			Handler:    _AnalysisService_StopSparkJob_Handler,
		},
		{
			MethodName: "CheckJob",
			Handler:    _AnalysisService_CheckJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/SparkService.proto",
}
