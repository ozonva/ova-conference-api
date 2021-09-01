// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ova_conference_api

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ConferencesClient is the client API for Conferences service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConferencesClient interface {
	CreateConference(ctx context.Context, in *CreateConferenceRequest, opts ...grpc.CallOption) (*ConferenceResponse, error)
	DescribeConference(ctx context.Context, in *EntityConferenceRequest, opts ...grpc.CallOption) (*ConferenceResponse, error)
	ListConference(ctx context.Context, in *ListConferenceRequest, opts ...grpc.CallOption) (*ListConferenceResponse, error)
	RemoveConference(ctx context.Context, in *EntityConferenceRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type conferencesClient struct {
	cc grpc.ClientConnInterface
}

func NewConferencesClient(cc grpc.ClientConnInterface) ConferencesClient {
	return &conferencesClient{cc}
}

func (c *conferencesClient) CreateConference(ctx context.Context, in *CreateConferenceRequest, opts ...grpc.CallOption) (*ConferenceResponse, error) {
	out := new(ConferenceResponse)
	err := c.cc.Invoke(ctx, "/ova.conference.api.Conferences/CreateConference", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conferencesClient) DescribeConference(ctx context.Context, in *EntityConferenceRequest, opts ...grpc.CallOption) (*ConferenceResponse, error) {
	out := new(ConferenceResponse)
	err := c.cc.Invoke(ctx, "/ova.conference.api.Conferences/DescribeConference", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conferencesClient) ListConference(ctx context.Context, in *ListConferenceRequest, opts ...grpc.CallOption) (*ListConferenceResponse, error) {
	out := new(ListConferenceResponse)
	err := c.cc.Invoke(ctx, "/ova.conference.api.Conferences/ListConference", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conferencesClient) RemoveConference(ctx context.Context, in *EntityConferenceRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ova.conference.api.Conferences/RemoveConference", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConferencesServer is the server API for Conferences service.
// All implementations must embed UnimplementedConferencesServer
// for forward compatibility
type ConferencesServer interface {
	CreateConference(context.Context, *CreateConferenceRequest) (*ConferenceResponse, error)
	DescribeConference(context.Context, *EntityConferenceRequest) (*ConferenceResponse, error)
	ListConference(context.Context, *ListConferenceRequest) (*ListConferenceResponse, error)
	RemoveConference(context.Context, *EntityConferenceRequest) (*empty.Empty, error)
	mustEmbedUnimplementedConferencesServer()
}

// UnimplementedConferencesServer must be embedded to have forward compatible implementations.
type UnimplementedConferencesServer struct {
}

func (UnimplementedConferencesServer) CreateConference(context.Context, *CreateConferenceRequest) (*ConferenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateConference not implemented")
}
func (UnimplementedConferencesServer) DescribeConference(context.Context, *EntityConferenceRequest) (*ConferenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeConference not implemented")
}
func (UnimplementedConferencesServer) ListConference(context.Context, *ListConferenceRequest) (*ListConferenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListConference not implemented")
}
func (UnimplementedConferencesServer) RemoveConference(context.Context, *EntityConferenceRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveConference not implemented")
}
func (UnimplementedConferencesServer) mustEmbedUnimplementedConferencesServer() {}

// UnsafeConferencesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConferencesServer will
// result in compilation errors.
type UnsafeConferencesServer interface {
	mustEmbedUnimplementedConferencesServer()
}

func RegisterConferencesServer(s grpc.ServiceRegistrar, srv ConferencesServer) {
	s.RegisterService(&Conferences_ServiceDesc, srv)
}

func _Conferences_CreateConference_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateConferenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConferencesServer).CreateConference(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.conference.api.Conferences/CreateConference",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConferencesServer).CreateConference(ctx, req.(*CreateConferenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conferences_DescribeConference_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EntityConferenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConferencesServer).DescribeConference(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.conference.api.Conferences/DescribeConference",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConferencesServer).DescribeConference(ctx, req.(*EntityConferenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conferences_ListConference_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListConferenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConferencesServer).ListConference(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.conference.api.Conferences/ListConference",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConferencesServer).ListConference(ctx, req.(*ListConferenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conferences_RemoveConference_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EntityConferenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConferencesServer).RemoveConference(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.conference.api.Conferences/RemoveConference",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConferencesServer).RemoveConference(ctx, req.(*EntityConferenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Conferences_ServiceDesc is the grpc.ServiceDesc for Conferences service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Conferences_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ova.conference.api.Conferences",
	HandlerType: (*ConferencesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateConference",
			Handler:    _Conferences_CreateConference_Handler,
		},
		{
			MethodName: "DescribeConference",
			Handler:    _Conferences_DescribeConference_Handler,
		},
		{
			MethodName: "ListConference",
			Handler:    _Conferences_ListConference_Handler,
		},
		{
			MethodName: "RemoveConference",
			Handler:    _Conferences_RemoveConference_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/api.proto",
}
