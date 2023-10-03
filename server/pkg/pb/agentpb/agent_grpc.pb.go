// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: agent.proto

package agentpb

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

// AgentClient is the client API for Agent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentClient interface {
	SubmitJob(ctx context.Context, in *JobRequest, opts ...grpc.CallOption) (*JobResponse, error)
	ClusterAdd(ctx context.Context, in *ClusterRequest, opts ...grpc.CallOption) (*JobResponse, error)
	ClusterDelete(ctx context.Context, in *ClusterRequest, opts ...grpc.CallOption) (*JobResponse, error)
	RepositoryAdd(ctx context.Context, in *RepositoryAddRequest, opts ...grpc.CallOption) (*JobResponse, error)
	RepositoryDelete(ctx context.Context, in *RepositoryDeleteRequest, opts ...grpc.CallOption) (*JobResponse, error)
	ProjectAdd(ctx context.Context, in *ProjectAddRequest, opts ...grpc.CallOption) (*JobResponse, error)
	ProjectDelete(ctx context.Context, in *ProjectDeleteRequest, opts ...grpc.CallOption) (*JobResponse, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	StoreCredential(ctx context.Context, in *StoreCredentialRequest, opts ...grpc.CallOption) (*StoreCredentialResponse, error)
	SyncApp(ctx context.Context, in *SyncAppRequest, opts ...grpc.CallOption) (*SyncAppResponse, error)
	GetClusterApps(ctx context.Context, in *GetClusterAppsRequest, opts ...grpc.CallOption) (*GetClusterAppsResponse, error)
	GetClusterAppLaunches(ctx context.Context, in *GetClusterAppLaunchesRequest, opts ...grpc.CallOption) (*GetClusterAppLaunchesResponse, error)
	ConfigureAppSSO(ctx context.Context, in *ConfigureAppSSORequest, opts ...grpc.CallOption) (*ConfigureAppSSOResponse, error)
	GetClusterAppConfig(ctx context.Context, in *GetClusterAppConfigRequest, opts ...grpc.CallOption) (*GetClusterAppConfigResponse, error)
	GetClusterAppValues(ctx context.Context, in *GetClusterAppValuesRequest, opts ...grpc.CallOption) (*GetClusterAppValuesResponse, error)
	GetClusterGlobalValues(ctx context.Context, in *GetClusterGlobalValuesRequest, opts ...grpc.CallOption) (*GetClusterGlobalValuesResponse, error)
	InstallApp(ctx context.Context, in *InstallAppRequest, opts ...grpc.CallOption) (*InstallAppResponse, error)
	UnInstallApp(ctx context.Context, in *UnInstallAppRequest, opts ...grpc.CallOption) (*UnInstallAppResponse, error)
	SetClusterGitoptsProject(ctx context.Context, in *SetClusterGitoptsProjectRequest, opts ...grpc.CallOption) (*SetClusterGitoptsProjectResponse, error)
	GetClusterGitoptsProject(ctx context.Context, in *GetClusterGitoptsProjectRequest, opts ...grpc.CallOption) (*GetClusterGitoptsProjectResponse, error)
	DeleteClusterGitoptsProject(ctx context.Context, in *DeleteClusterGitoptsProjectRequest, opts ...grpc.CallOption) (*DeleteClusterGitoptsProjectResponse, error)
	UpgradeApp(ctx context.Context, in *UpgradeAppRequest, opts ...grpc.CallOption) (*UpgradeAppResponse, error)
	UpdateAppValues(ctx context.Context, in *UpdateAppValuesRequest, opts ...grpc.CallOption) (*UpdateAppValuesResponse, error)
}

type agentClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentClient(cc grpc.ClientConnInterface) AgentClient {
	return &agentClient{cc}
}

func (c *agentClient) SubmitJob(ctx context.Context, in *JobRequest, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := c.cc.Invoke(ctx, "/agentpb.Agent/SubmitJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) ClusterAdd(ctx context.Context, in *ClusterRequest, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := c.cc.Invoke(ctx, "/agentpb.Agent/ClusterAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) ClusterDelete(ctx context.Context, in *ClusterRequest, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := c.cc.Invoke(ctx, "/agentpb.Agent/ClusterDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) RepositoryAdd(ctx context.Context, in *RepositoryAddRequest, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := c.cc.Invoke(ctx, "/agentpb.Agent/RepositoryAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) RepositoryDelete(ctx context.Context, in *RepositoryDeleteRequest, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := c.cc.Invoke(ctx, "/agentpb.Agent/RepositoryDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) ProjectAdd(ctx context.Context, in *ProjectAddRequest, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := c.cc.Invoke(ctx, "/agentpb.Agent/ProjectAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) ProjectDelete(ctx context.Context, in *ProjectDeleteRequest, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := c.cc.Invoke(ctx, "/agentpb.Agent/ProjectDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/agentpb.Agent/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) StoreCredential(ctx context.Context, in *StoreCredentialRequest, opts ...grpc.CallOption) (*StoreCredentialResponse, error) {
	out := new(StoreCredentialResponse)
	err := c.cc.Invoke(ctx, "/agentpb.Agent/StoreCredential", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) SyncApp(ctx context.Context, in *SyncAppRequest, opts ...grpc.CallOption) (*SyncAppResponse, error) {
	out := new(SyncAppResponse)
	err := c.cc.Invoke(ctx, "/agentpb.Agent/SyncApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) GetClusterApps(ctx context.Context, in *GetClusterAppsRequest, opts ...grpc.CallOption) (*GetClusterAppsResponse, error) {
	out := new(GetClusterAppsResponse)
	err := c.cc.Invoke(ctx, "/agentpb.Agent/GetClusterApps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) GetClusterAppLaunches(ctx context.Context, in *GetClusterAppLaunchesRequest, opts ...grpc.CallOption) (*GetClusterAppLaunchesResponse, error) {
	out := new(GetClusterAppLaunchesResponse)
	err := c.cc.Invoke(ctx, "/agentpb.Agent/GetClusterAppLaunches", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) ConfigureAppSSO(ctx context.Context, in *ConfigureAppSSORequest, opts ...grpc.CallOption) (*ConfigureAppSSOResponse, error) {
	out := new(ConfigureAppSSOResponse)
	err := c.cc.Invoke(ctx, "/agentpb.Agent/ConfigureAppSSO", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) GetClusterAppConfig(ctx context.Context, in *GetClusterAppConfigRequest, opts ...grpc.CallOption) (*GetClusterAppConfigResponse, error) {
	out := new(GetClusterAppConfigResponse)
	err := c.cc.Invoke(ctx, "/agentpb.Agent/GetClusterAppConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) GetClusterAppValues(ctx context.Context, in *GetClusterAppValuesRequest, opts ...grpc.CallOption) (*GetClusterAppValuesResponse, error) {
	out := new(GetClusterAppValuesResponse)
	err := c.cc.Invoke(ctx, "/agentpb.Agent/GetClusterAppValues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) GetClusterGlobalValues(ctx context.Context, in *GetClusterGlobalValuesRequest, opts ...grpc.CallOption) (*GetClusterGlobalValuesResponse, error) {
	out := new(GetClusterGlobalValuesResponse)
	err := c.cc.Invoke(ctx, "/agentpb.Agent/GetClusterGlobalValues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) InstallApp(ctx context.Context, in *InstallAppRequest, opts ...grpc.CallOption) (*InstallAppResponse, error) {
	out := new(InstallAppResponse)
	err := c.cc.Invoke(ctx, "/agentpb.Agent/InstallApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) UnInstallApp(ctx context.Context, in *UnInstallAppRequest, opts ...grpc.CallOption) (*UnInstallAppResponse, error) {
	out := new(UnInstallAppResponse)
	err := c.cc.Invoke(ctx, "/agentpb.Agent/UnInstallApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) SetClusterGitoptsProject(ctx context.Context, in *SetClusterGitoptsProjectRequest, opts ...grpc.CallOption) (*SetClusterGitoptsProjectResponse, error) {
	out := new(SetClusterGitoptsProjectResponse)
	err := c.cc.Invoke(ctx, "/agentpb.Agent/SetClusterGitoptsProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) GetClusterGitoptsProject(ctx context.Context, in *GetClusterGitoptsProjectRequest, opts ...grpc.CallOption) (*GetClusterGitoptsProjectResponse, error) {
	out := new(GetClusterGitoptsProjectResponse)
	err := c.cc.Invoke(ctx, "/agentpb.Agent/GetClusterGitoptsProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) DeleteClusterGitoptsProject(ctx context.Context, in *DeleteClusterGitoptsProjectRequest, opts ...grpc.CallOption) (*DeleteClusterGitoptsProjectResponse, error) {
	out := new(DeleteClusterGitoptsProjectResponse)
	err := c.cc.Invoke(ctx, "/agentpb.Agent/DeleteClusterGitoptsProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) UpgradeApp(ctx context.Context, in *UpgradeAppRequest, opts ...grpc.CallOption) (*UpgradeAppResponse, error) {
	out := new(UpgradeAppResponse)
	err := c.cc.Invoke(ctx, "/agentpb.Agent/UpgradeApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) UpdateAppValues(ctx context.Context, in *UpdateAppValuesRequest, opts ...grpc.CallOption) (*UpdateAppValuesResponse, error) {
	out := new(UpdateAppValuesResponse)
	err := c.cc.Invoke(ctx, "/agentpb.Agent/UpdateAppValues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServer is the server API for Agent service.
// All implementations must embed UnimplementedAgentServer
// for forward compatibility
type AgentServer interface {
	SubmitJob(context.Context, *JobRequest) (*JobResponse, error)
	ClusterAdd(context.Context, *ClusterRequest) (*JobResponse, error)
	ClusterDelete(context.Context, *ClusterRequest) (*JobResponse, error)
	RepositoryAdd(context.Context, *RepositoryAddRequest) (*JobResponse, error)
	RepositoryDelete(context.Context, *RepositoryDeleteRequest) (*JobResponse, error)
	ProjectAdd(context.Context, *ProjectAddRequest) (*JobResponse, error)
	ProjectDelete(context.Context, *ProjectDeleteRequest) (*JobResponse, error)
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	StoreCredential(context.Context, *StoreCredentialRequest) (*StoreCredentialResponse, error)
	SyncApp(context.Context, *SyncAppRequest) (*SyncAppResponse, error)
	GetClusterApps(context.Context, *GetClusterAppsRequest) (*GetClusterAppsResponse, error)
	GetClusterAppLaunches(context.Context, *GetClusterAppLaunchesRequest) (*GetClusterAppLaunchesResponse, error)
	ConfigureAppSSO(context.Context, *ConfigureAppSSORequest) (*ConfigureAppSSOResponse, error)
	GetClusterAppConfig(context.Context, *GetClusterAppConfigRequest) (*GetClusterAppConfigResponse, error)
	GetClusterAppValues(context.Context, *GetClusterAppValuesRequest) (*GetClusterAppValuesResponse, error)
	GetClusterGlobalValues(context.Context, *GetClusterGlobalValuesRequest) (*GetClusterGlobalValuesResponse, error)
	InstallApp(context.Context, *InstallAppRequest) (*InstallAppResponse, error)
	UnInstallApp(context.Context, *UnInstallAppRequest) (*UnInstallAppResponse, error)
	SetClusterGitoptsProject(context.Context, *SetClusterGitoptsProjectRequest) (*SetClusterGitoptsProjectResponse, error)
	GetClusterGitoptsProject(context.Context, *GetClusterGitoptsProjectRequest) (*GetClusterGitoptsProjectResponse, error)
	DeleteClusterGitoptsProject(context.Context, *DeleteClusterGitoptsProjectRequest) (*DeleteClusterGitoptsProjectResponse, error)
	UpgradeApp(context.Context, *UpgradeAppRequest) (*UpgradeAppResponse, error)
	UpdateAppValues(context.Context, *UpdateAppValuesRequest) (*UpdateAppValuesResponse, error)
	mustEmbedUnimplementedAgentServer()
}

// UnimplementedAgentServer must be embedded to have forward compatible implementations.
type UnimplementedAgentServer struct {
}

func (UnimplementedAgentServer) SubmitJob(context.Context, *JobRequest) (*JobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitJob not implemented")
}
func (UnimplementedAgentServer) ClusterAdd(context.Context, *ClusterRequest) (*JobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClusterAdd not implemented")
}
func (UnimplementedAgentServer) ClusterDelete(context.Context, *ClusterRequest) (*JobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClusterDelete not implemented")
}
func (UnimplementedAgentServer) RepositoryAdd(context.Context, *RepositoryAddRequest) (*JobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RepositoryAdd not implemented")
}
func (UnimplementedAgentServer) RepositoryDelete(context.Context, *RepositoryDeleteRequest) (*JobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RepositoryDelete not implemented")
}
func (UnimplementedAgentServer) ProjectAdd(context.Context, *ProjectAddRequest) (*JobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProjectAdd not implemented")
}
func (UnimplementedAgentServer) ProjectDelete(context.Context, *ProjectDeleteRequest) (*JobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProjectDelete not implemented")
}
func (UnimplementedAgentServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedAgentServer) StoreCredential(context.Context, *StoreCredentialRequest) (*StoreCredentialResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreCredential not implemented")
}
func (UnimplementedAgentServer) SyncApp(context.Context, *SyncAppRequest) (*SyncAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncApp not implemented")
}
func (UnimplementedAgentServer) GetClusterApps(context.Context, *GetClusterAppsRequest) (*GetClusterAppsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusterApps not implemented")
}
func (UnimplementedAgentServer) GetClusterAppLaunches(context.Context, *GetClusterAppLaunchesRequest) (*GetClusterAppLaunchesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusterAppLaunches not implemented")
}
func (UnimplementedAgentServer) ConfigureAppSSO(context.Context, *ConfigureAppSSORequest) (*ConfigureAppSSOResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigureAppSSO not implemented")
}
func (UnimplementedAgentServer) GetClusterAppConfig(context.Context, *GetClusterAppConfigRequest) (*GetClusterAppConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusterAppConfig not implemented")
}
func (UnimplementedAgentServer) GetClusterAppValues(context.Context, *GetClusterAppValuesRequest) (*GetClusterAppValuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusterAppValues not implemented")
}
func (UnimplementedAgentServer) GetClusterGlobalValues(context.Context, *GetClusterGlobalValuesRequest) (*GetClusterGlobalValuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusterGlobalValues not implemented")
}
func (UnimplementedAgentServer) InstallApp(context.Context, *InstallAppRequest) (*InstallAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InstallApp not implemented")
}
func (UnimplementedAgentServer) UnInstallApp(context.Context, *UnInstallAppRequest) (*UnInstallAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnInstallApp not implemented")
}
func (UnimplementedAgentServer) SetClusterGitoptsProject(context.Context, *SetClusterGitoptsProjectRequest) (*SetClusterGitoptsProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetClusterGitoptsProject not implemented")
}
func (UnimplementedAgentServer) GetClusterGitoptsProject(context.Context, *GetClusterGitoptsProjectRequest) (*GetClusterGitoptsProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusterGitoptsProject not implemented")
}
func (UnimplementedAgentServer) DeleteClusterGitoptsProject(context.Context, *DeleteClusterGitoptsProjectRequest) (*DeleteClusterGitoptsProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteClusterGitoptsProject not implemented")
}
func (UnimplementedAgentServer) UpgradeApp(context.Context, *UpgradeAppRequest) (*UpgradeAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpgradeApp not implemented")
}
func (UnimplementedAgentServer) UpdateAppValues(context.Context, *UpdateAppValuesRequest) (*UpdateAppValuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAppValues not implemented")
}
func (UnimplementedAgentServer) mustEmbedUnimplementedAgentServer() {}

// UnsafeAgentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentServer will
// result in compilation errors.
type UnsafeAgentServer interface {
	mustEmbedUnimplementedAgentServer()
}

func RegisterAgentServer(s grpc.ServiceRegistrar, srv AgentServer) {
	s.RegisterService(&Agent_ServiceDesc, srv)
}

func _Agent_SubmitJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).SubmitJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agentpb.Agent/SubmitJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).SubmitJob(ctx, req.(*JobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_ClusterAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).ClusterAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agentpb.Agent/ClusterAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).ClusterAdd(ctx, req.(*ClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_ClusterDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).ClusterDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agentpb.Agent/ClusterDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).ClusterDelete(ctx, req.(*ClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_RepositoryAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepositoryAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).RepositoryAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agentpb.Agent/RepositoryAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).RepositoryAdd(ctx, req.(*RepositoryAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_RepositoryDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepositoryDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).RepositoryDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agentpb.Agent/RepositoryDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).RepositoryDelete(ctx, req.(*RepositoryDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_ProjectAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).ProjectAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agentpb.Agent/ProjectAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).ProjectAdd(ctx, req.(*ProjectAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_ProjectDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).ProjectDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agentpb.Agent/ProjectDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).ProjectDelete(ctx, req.(*ProjectDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agentpb.Agent/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_StoreCredential_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreCredentialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).StoreCredential(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agentpb.Agent/StoreCredential",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).StoreCredential(ctx, req.(*StoreCredentialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_SyncApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).SyncApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agentpb.Agent/SyncApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).SyncApp(ctx, req.(*SyncAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_GetClusterApps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterAppsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).GetClusterApps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agentpb.Agent/GetClusterApps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).GetClusterApps(ctx, req.(*GetClusterAppsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_GetClusterAppLaunches_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterAppLaunchesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).GetClusterAppLaunches(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agentpb.Agent/GetClusterAppLaunches",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).GetClusterAppLaunches(ctx, req.(*GetClusterAppLaunchesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_ConfigureAppSSO_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigureAppSSORequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).ConfigureAppSSO(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agentpb.Agent/ConfigureAppSSO",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).ConfigureAppSSO(ctx, req.(*ConfigureAppSSORequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_GetClusterAppConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterAppConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).GetClusterAppConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agentpb.Agent/GetClusterAppConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).GetClusterAppConfig(ctx, req.(*GetClusterAppConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_GetClusterAppValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterAppValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).GetClusterAppValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agentpb.Agent/GetClusterAppValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).GetClusterAppValues(ctx, req.(*GetClusterAppValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_GetClusterGlobalValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterGlobalValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).GetClusterGlobalValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agentpb.Agent/GetClusterGlobalValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).GetClusterGlobalValues(ctx, req.(*GetClusterGlobalValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_InstallApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstallAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).InstallApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agentpb.Agent/InstallApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).InstallApp(ctx, req.(*InstallAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_UnInstallApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnInstallAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).UnInstallApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agentpb.Agent/UnInstallApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).UnInstallApp(ctx, req.(*UnInstallAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_SetClusterGitoptsProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetClusterGitoptsProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).SetClusterGitoptsProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agentpb.Agent/SetClusterGitoptsProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).SetClusterGitoptsProject(ctx, req.(*SetClusterGitoptsProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_GetClusterGitoptsProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterGitoptsProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).GetClusterGitoptsProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agentpb.Agent/GetClusterGitoptsProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).GetClusterGitoptsProject(ctx, req.(*GetClusterGitoptsProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_DeleteClusterGitoptsProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteClusterGitoptsProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).DeleteClusterGitoptsProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agentpb.Agent/DeleteClusterGitoptsProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).DeleteClusterGitoptsProject(ctx, req.(*DeleteClusterGitoptsProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_UpgradeApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpgradeAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).UpgradeApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agentpb.Agent/UpgradeApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).UpgradeApp(ctx, req.(*UpgradeAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_UpdateAppValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAppValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).UpdateAppValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agentpb.Agent/UpdateAppValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).UpdateAppValues(ctx, req.(*UpdateAppValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Agent_ServiceDesc is the grpc.ServiceDesc for Agent service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Agent_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "agentpb.Agent",
	HandlerType: (*AgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitJob",
			Handler:    _Agent_SubmitJob_Handler,
		},
		{
			MethodName: "ClusterAdd",
			Handler:    _Agent_ClusterAdd_Handler,
		},
		{
			MethodName: "ClusterDelete",
			Handler:    _Agent_ClusterDelete_Handler,
		},
		{
			MethodName: "RepositoryAdd",
			Handler:    _Agent_RepositoryAdd_Handler,
		},
		{
			MethodName: "RepositoryDelete",
			Handler:    _Agent_RepositoryDelete_Handler,
		},
		{
			MethodName: "ProjectAdd",
			Handler:    _Agent_ProjectAdd_Handler,
		},
		{
			MethodName: "ProjectDelete",
			Handler:    _Agent_ProjectDelete_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Agent_Ping_Handler,
		},
		{
			MethodName: "StoreCredential",
			Handler:    _Agent_StoreCredential_Handler,
		},
		{
			MethodName: "SyncApp",
			Handler:    _Agent_SyncApp_Handler,
		},
		{
			MethodName: "GetClusterApps",
			Handler:    _Agent_GetClusterApps_Handler,
		},
		{
			MethodName: "GetClusterAppLaunches",
			Handler:    _Agent_GetClusterAppLaunches_Handler,
		},
		{
			MethodName: "ConfigureAppSSO",
			Handler:    _Agent_ConfigureAppSSO_Handler,
		},
		{
			MethodName: "GetClusterAppConfig",
			Handler:    _Agent_GetClusterAppConfig_Handler,
		},
		{
			MethodName: "GetClusterAppValues",
			Handler:    _Agent_GetClusterAppValues_Handler,
		},
		{
			MethodName: "GetClusterGlobalValues",
			Handler:    _Agent_GetClusterGlobalValues_Handler,
		},
		{
			MethodName: "InstallApp",
			Handler:    _Agent_InstallApp_Handler,
		},
		{
			MethodName: "UnInstallApp",
			Handler:    _Agent_UnInstallApp_Handler,
		},
		{
			MethodName: "SetClusterGitoptsProject",
			Handler:    _Agent_SetClusterGitoptsProject_Handler,
		},
		{
			MethodName: "GetClusterGitoptsProject",
			Handler:    _Agent_GetClusterGitoptsProject_Handler,
		},
		{
			MethodName: "DeleteClusterGitoptsProject",
			Handler:    _Agent_DeleteClusterGitoptsProject_Handler,
		},
		{
			MethodName: "UpgradeApp",
			Handler:    _Agent_UpgradeApp_Handler,
		},
		{
			MethodName: "UpdateAppValues",
			Handler:    _Agent_UpdateAppValues_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agent.proto",
}
