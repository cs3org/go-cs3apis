// Copyright 2018-2019 CERN
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// In applying this license, CERN does not waive the privileges and immunities
// granted to it by virtue of its status as an Intergovernmental Organization
// or submit itself to any jurisdiction.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cs3/identity/user/v1beta1/user_api.proto

package userv1beta1

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
	UserAPI_GetUser_FullMethodName        = "/cs3.identity.user.v1beta1.UserAPI/GetUser"
	UserAPI_GetUserByClaim_FullMethodName = "/cs3.identity.user.v1beta1.UserAPI/GetUserByClaim"
	UserAPI_GetUserGroups_FullMethodName  = "/cs3.identity.user.v1beta1.UserAPI/GetUserGroups"
	UserAPI_FindUsers_FullMethodName      = "/cs3.identity.user.v1beta1.UserAPI/FindUsers"
)

// UserAPIClient is the client API for UserAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserAPIClient interface {
	// Gets the information about a user by the user id.
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	// Gets the information about a user based on a specified claim.
	GetUserByClaim(ctx context.Context, in *GetUserByClaimRequest, opts ...grpc.CallOption) (*GetUserByClaimResponse, error)
	// Gets the groups of a user.
	GetUserGroups(ctx context.Context, in *GetUserGroupsRequest, opts ...grpc.CallOption) (*GetUserGroupsResponse, error)
	// Finds users by any attribute of the user.
	// TODO(labkode): to define the filters that make more sense.
	FindUsers(ctx context.Context, in *FindUsersRequest, opts ...grpc.CallOption) (*FindUsersResponse, error)
}

type userAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewUserAPIClient(cc grpc.ClientConnInterface) UserAPIClient {
	return &userAPIClient{cc}
}

func (c *userAPIClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, UserAPI_GetUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAPIClient) GetUserByClaim(ctx context.Context, in *GetUserByClaimRequest, opts ...grpc.CallOption) (*GetUserByClaimResponse, error) {
	out := new(GetUserByClaimResponse)
	err := c.cc.Invoke(ctx, UserAPI_GetUserByClaim_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAPIClient) GetUserGroups(ctx context.Context, in *GetUserGroupsRequest, opts ...grpc.CallOption) (*GetUserGroupsResponse, error) {
	out := new(GetUserGroupsResponse)
	err := c.cc.Invoke(ctx, UserAPI_GetUserGroups_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAPIClient) FindUsers(ctx context.Context, in *FindUsersRequest, opts ...grpc.CallOption) (*FindUsersResponse, error) {
	out := new(FindUsersResponse)
	err := c.cc.Invoke(ctx, UserAPI_FindUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserAPIServer is the server API for UserAPI service.
// All implementations should embed UnimplementedUserAPIServer
// for forward compatibility
type UserAPIServer interface {
	// Gets the information about a user by the user id.
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	// Gets the information about a user based on a specified claim.
	GetUserByClaim(context.Context, *GetUserByClaimRequest) (*GetUserByClaimResponse, error)
	// Gets the groups of a user.
	GetUserGroups(context.Context, *GetUserGroupsRequest) (*GetUserGroupsResponse, error)
	// Finds users by any attribute of the user.
	// TODO(labkode): to define the filters that make more sense.
	FindUsers(context.Context, *FindUsersRequest) (*FindUsersResponse, error)
}

// UnimplementedUserAPIServer should be embedded to have forward compatible implementations.
type UnimplementedUserAPIServer struct {
}

func (UnimplementedUserAPIServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserAPIServer) GetUserByClaim(context.Context, *GetUserByClaimRequest) (*GetUserByClaimResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByClaim not implemented")
}
func (UnimplementedUserAPIServer) GetUserGroups(context.Context, *GetUserGroupsRequest) (*GetUserGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserGroups not implemented")
}
func (UnimplementedUserAPIServer) FindUsers(context.Context, *FindUsersRequest) (*FindUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUsers not implemented")
}

// UnsafeUserAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserAPIServer will
// result in compilation errors.
type UnsafeUserAPIServer interface {
	mustEmbedUnimplementedUserAPIServer()
}

func RegisterUserAPIServer(s grpc.ServiceRegistrar, srv UserAPIServer) {
	s.RegisterService(&UserAPI_ServiceDesc, srv)
}

func _UserAPI_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAPIServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserAPI_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAPIServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAPI_GetUserByClaim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByClaimRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAPIServer).GetUserByClaim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserAPI_GetUserByClaim_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAPIServer).GetUserByClaim(ctx, req.(*GetUserByClaimRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAPI_GetUserGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAPIServer).GetUserGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserAPI_GetUserGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAPIServer).GetUserGroups(ctx, req.(*GetUserGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAPI_FindUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAPIServer).FindUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserAPI_FindUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAPIServer).FindUsers(ctx, req.(*FindUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserAPI_ServiceDesc is the grpc.ServiceDesc for UserAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cs3.identity.user.v1beta1.UserAPI",
	HandlerType: (*UserAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _UserAPI_GetUser_Handler,
		},
		{
			MethodName: "GetUserByClaim",
			Handler:    _UserAPI_GetUserByClaim_Handler,
		},
		{
			MethodName: "GetUserGroups",
			Handler:    _UserAPI_GetUserGroups_Handler,
		},
		{
			MethodName: "FindUsers",
			Handler:    _UserAPI_FindUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cs3/identity/user/v1beta1/user_api.proto",
}
