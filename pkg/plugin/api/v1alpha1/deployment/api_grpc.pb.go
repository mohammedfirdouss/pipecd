// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: pkg/plugin/api/v1alpha1/deployment/api.proto

package deployment

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

// DeploymentServiceClient is the client API for DeploymentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeploymentServiceClient interface {
	// FetchDefinedStages fetches the defined stages' name which are supported by the plugin.
	FetchDefinedStages(ctx context.Context, in *FetchDefinedStagesRequest, opts ...grpc.CallOption) (*FetchDefinedStagesResponse, error)
	// DetermineVersions determines which versions of the artifacts will be used for the given deployment.
	DetermineVersions(ctx context.Context, in *DetermineVersionsRequest, opts ...grpc.CallOption) (*DetermineVersionsResponse, error)
	// DetermineStrategy determines which strategy should be used for the given deployment.
	DetermineStrategy(ctx context.Context, in *DetermineStrategyRequest, opts ...grpc.CallOption) (*DetermineStrategyResponse, error)
	// BuildPipelineSyncStages builds the deployment pipeline stages.
	// The built pipeline includes non-rollback (defined in the application config) and rollback stages.
	// The request contains only non-rollback stages whose names are listed in FetchDefinedStages() of this plugin.
	//
	// Note about the response indexes:
	//   - For a non-rollback stage, use the index given by the request remaining the execution order.
	//   - For a rollback stage, use one of the indexes given by the request.
	//   - The indexes of the response stages must not be duplicated across non-rollback stages and rollback stages.
	//     A non-rollback stage and a rollback stage can have the same index.
	//
	// For example, given request indexes are {2,4,5}, then
	//   - Non-rollback stages indexes must be {2,4,5}
	//   - Rollback stages indexes must be selected from {2,4,5}.  For a deploymentPlugin, using only {2} is recommended.
	BuildPipelineSyncStages(ctx context.Context, in *BuildPipelineSyncStagesRequest, opts ...grpc.CallOption) (*BuildPipelineSyncStagesResponse, error)
	// BuildQuickSyncStages builds the deployment quick sync stages.
	BuildQuickSyncStages(ctx context.Context, in *BuildQuickSyncStagesRequest, opts ...grpc.CallOption) (*BuildQuickSyncStagesResponse, error)
	// ExecuteStage executes the given stage.
	ExecuteStage(ctx context.Context, in *ExecuteStageRequest, opts ...grpc.CallOption) (*ExecuteStageResponse, error)
}

type deploymentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeploymentServiceClient(cc grpc.ClientConnInterface) DeploymentServiceClient {
	return &deploymentServiceClient{cc}
}

func (c *deploymentServiceClient) FetchDefinedStages(ctx context.Context, in *FetchDefinedStagesRequest, opts ...grpc.CallOption) (*FetchDefinedStagesResponse, error) {
	out := new(FetchDefinedStagesResponse)
	err := c.cc.Invoke(ctx, "/grpc.plugin.deploymentapi.v1alpha1.DeploymentService/FetchDefinedStages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deploymentServiceClient) DetermineVersions(ctx context.Context, in *DetermineVersionsRequest, opts ...grpc.CallOption) (*DetermineVersionsResponse, error) {
	out := new(DetermineVersionsResponse)
	err := c.cc.Invoke(ctx, "/grpc.plugin.deploymentapi.v1alpha1.DeploymentService/DetermineVersions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deploymentServiceClient) DetermineStrategy(ctx context.Context, in *DetermineStrategyRequest, opts ...grpc.CallOption) (*DetermineStrategyResponse, error) {
	out := new(DetermineStrategyResponse)
	err := c.cc.Invoke(ctx, "/grpc.plugin.deploymentapi.v1alpha1.DeploymentService/DetermineStrategy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deploymentServiceClient) BuildPipelineSyncStages(ctx context.Context, in *BuildPipelineSyncStagesRequest, opts ...grpc.CallOption) (*BuildPipelineSyncStagesResponse, error) {
	out := new(BuildPipelineSyncStagesResponse)
	err := c.cc.Invoke(ctx, "/grpc.plugin.deploymentapi.v1alpha1.DeploymentService/BuildPipelineSyncStages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deploymentServiceClient) BuildQuickSyncStages(ctx context.Context, in *BuildQuickSyncStagesRequest, opts ...grpc.CallOption) (*BuildQuickSyncStagesResponse, error) {
	out := new(BuildQuickSyncStagesResponse)
	err := c.cc.Invoke(ctx, "/grpc.plugin.deploymentapi.v1alpha1.DeploymentService/BuildQuickSyncStages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deploymentServiceClient) ExecuteStage(ctx context.Context, in *ExecuteStageRequest, opts ...grpc.CallOption) (*ExecuteStageResponse, error) {
	out := new(ExecuteStageResponse)
	err := c.cc.Invoke(ctx, "/grpc.plugin.deploymentapi.v1alpha1.DeploymentService/ExecuteStage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeploymentServiceServer is the server API for DeploymentService service.
// All implementations must embed UnimplementedDeploymentServiceServer
// for forward compatibility
type DeploymentServiceServer interface {
	// FetchDefinedStages fetches the defined stages' name which are supported by the plugin.
	FetchDefinedStages(context.Context, *FetchDefinedStagesRequest) (*FetchDefinedStagesResponse, error)
	// DetermineVersions determines which versions of the artifacts will be used for the given deployment.
	DetermineVersions(context.Context, *DetermineVersionsRequest) (*DetermineVersionsResponse, error)
	// DetermineStrategy determines which strategy should be used for the given deployment.
	DetermineStrategy(context.Context, *DetermineStrategyRequest) (*DetermineStrategyResponse, error)
	// BuildPipelineSyncStages builds the deployment pipeline stages.
	// The built pipeline includes non-rollback (defined in the application config) and rollback stages.
	// The request contains only non-rollback stages whose names are listed in FetchDefinedStages() of this plugin.
	//
	// Note about the response indexes:
	//   - For a non-rollback stage, use the index given by the request remaining the execution order.
	//   - For a rollback stage, use one of the indexes given by the request.
	//   - The indexes of the response stages must not be duplicated across non-rollback stages and rollback stages.
	//     A non-rollback stage and a rollback stage can have the same index.
	//
	// For example, given request indexes are {2,4,5}, then
	//   - Non-rollback stages indexes must be {2,4,5}
	//   - Rollback stages indexes must be selected from {2,4,5}.  For a deploymentPlugin, using only {2} is recommended.
	BuildPipelineSyncStages(context.Context, *BuildPipelineSyncStagesRequest) (*BuildPipelineSyncStagesResponse, error)
	// BuildQuickSyncStages builds the deployment quick sync stages.
	BuildQuickSyncStages(context.Context, *BuildQuickSyncStagesRequest) (*BuildQuickSyncStagesResponse, error)
	// ExecuteStage executes the given stage.
	ExecuteStage(context.Context, *ExecuteStageRequest) (*ExecuteStageResponse, error)
	mustEmbedUnimplementedDeploymentServiceServer()
}

// UnimplementedDeploymentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDeploymentServiceServer struct {
}

func (UnimplementedDeploymentServiceServer) FetchDefinedStages(context.Context, *FetchDefinedStagesRequest) (*FetchDefinedStagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchDefinedStages not implemented")
}
func (UnimplementedDeploymentServiceServer) DetermineVersions(context.Context, *DetermineVersionsRequest) (*DetermineVersionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetermineVersions not implemented")
}
func (UnimplementedDeploymentServiceServer) DetermineStrategy(context.Context, *DetermineStrategyRequest) (*DetermineStrategyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetermineStrategy not implemented")
}
func (UnimplementedDeploymentServiceServer) BuildPipelineSyncStages(context.Context, *BuildPipelineSyncStagesRequest) (*BuildPipelineSyncStagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuildPipelineSyncStages not implemented")
}
func (UnimplementedDeploymentServiceServer) BuildQuickSyncStages(context.Context, *BuildQuickSyncStagesRequest) (*BuildQuickSyncStagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuildQuickSyncStages not implemented")
}
func (UnimplementedDeploymentServiceServer) ExecuteStage(context.Context, *ExecuteStageRequest) (*ExecuteStageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteStage not implemented")
}
func (UnimplementedDeploymentServiceServer) mustEmbedUnimplementedDeploymentServiceServer() {}

// UnsafeDeploymentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeploymentServiceServer will
// result in compilation errors.
type UnsafeDeploymentServiceServer interface {
	mustEmbedUnimplementedDeploymentServiceServer()
}

func RegisterDeploymentServiceServer(s grpc.ServiceRegistrar, srv DeploymentServiceServer) {
	s.RegisterService(&DeploymentService_ServiceDesc, srv)
}

func _DeploymentService_FetchDefinedStages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchDefinedStagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeploymentServiceServer).FetchDefinedStages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.plugin.deploymentapi.v1alpha1.DeploymentService/FetchDefinedStages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeploymentServiceServer).FetchDefinedStages(ctx, req.(*FetchDefinedStagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeploymentService_DetermineVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetermineVersionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeploymentServiceServer).DetermineVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.plugin.deploymentapi.v1alpha1.DeploymentService/DetermineVersions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeploymentServiceServer).DetermineVersions(ctx, req.(*DetermineVersionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeploymentService_DetermineStrategy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetermineStrategyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeploymentServiceServer).DetermineStrategy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.plugin.deploymentapi.v1alpha1.DeploymentService/DetermineStrategy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeploymentServiceServer).DetermineStrategy(ctx, req.(*DetermineStrategyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeploymentService_BuildPipelineSyncStages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildPipelineSyncStagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeploymentServiceServer).BuildPipelineSyncStages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.plugin.deploymentapi.v1alpha1.DeploymentService/BuildPipelineSyncStages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeploymentServiceServer).BuildPipelineSyncStages(ctx, req.(*BuildPipelineSyncStagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeploymentService_BuildQuickSyncStages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildQuickSyncStagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeploymentServiceServer).BuildQuickSyncStages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.plugin.deploymentapi.v1alpha1.DeploymentService/BuildQuickSyncStages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeploymentServiceServer).BuildQuickSyncStages(ctx, req.(*BuildQuickSyncStagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeploymentService_ExecuteStage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteStageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeploymentServiceServer).ExecuteStage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.plugin.deploymentapi.v1alpha1.DeploymentService/ExecuteStage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeploymentServiceServer).ExecuteStage(ctx, req.(*ExecuteStageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DeploymentService_ServiceDesc is the grpc.ServiceDesc for DeploymentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeploymentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.plugin.deploymentapi.v1alpha1.DeploymentService",
	HandlerType: (*DeploymentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchDefinedStages",
			Handler:    _DeploymentService_FetchDefinedStages_Handler,
		},
		{
			MethodName: "DetermineVersions",
			Handler:    _DeploymentService_DetermineVersions_Handler,
		},
		{
			MethodName: "DetermineStrategy",
			Handler:    _DeploymentService_DetermineStrategy_Handler,
		},
		{
			MethodName: "BuildPipelineSyncStages",
			Handler:    _DeploymentService_BuildPipelineSyncStages_Handler,
		},
		{
			MethodName: "BuildQuickSyncStages",
			Handler:    _DeploymentService_BuildQuickSyncStages_Handler,
		},
		{
			MethodName: "ExecuteStage",
			Handler:    _DeploymentService_ExecuteStage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/plugin/api/v1alpha1/deployment/api.proto",
}
