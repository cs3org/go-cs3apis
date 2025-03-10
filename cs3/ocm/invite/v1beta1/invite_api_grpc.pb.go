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
// source: cs3/ocm/invite/v1beta1/invite_api.proto

package invitev1beta1

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
	InviteAPI_GenerateInviteToken_FullMethodName = "/cs3.ocm.invite.v1beta1.InviteAPI/GenerateInviteToken"
	InviteAPI_ListInviteTokens_FullMethodName    = "/cs3.ocm.invite.v1beta1.InviteAPI/ListInviteTokens"
	InviteAPI_ForwardInvite_FullMethodName       = "/cs3.ocm.invite.v1beta1.InviteAPI/ForwardInvite"
	InviteAPI_AcceptInvite_FullMethodName        = "/cs3.ocm.invite.v1beta1.InviteAPI/AcceptInvite"
	InviteAPI_GetAcceptedUser_FullMethodName     = "/cs3.ocm.invite.v1beta1.InviteAPI/GetAcceptedUser"
	InviteAPI_FindAcceptedUsers_FullMethodName   = "/cs3.ocm.invite.v1beta1.InviteAPI/FindAcceptedUsers"
	InviteAPI_DeleteAcceptedUser_FullMethodName  = "/cs3.ocm.invite.v1beta1.InviteAPI/DeleteAcceptedUser"
)

// InviteAPIClient is the client API for InviteAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InviteAPIClient interface {
	// Generates a new token for the user with a validity of 24 hours.
	GenerateInviteToken(ctx context.Context, in *GenerateInviteTokenRequest, opts ...grpc.CallOption) (*GenerateInviteTokenResponse, error)
	// Lists the valid tokens generated by the user.
	ListInviteTokens(ctx context.Context, in *ListInviteTokensRequest, opts ...grpc.CallOption) (*ListInviteTokensResponse, error)
	// Forwards a received invite to the remote sync'n'share system provider. The remote
	// system SHALL get an `invite-accepted` call as follows:
	// https://cs3org.github.io/OCM-API/docs.html?branch=v1.2.0&repo=OCM-API&user=cs3org#/paths/~1invite-accepted/post
	// MUST return CODE_NOT_FOUND if the token does not exist.
	// MUST return CODE_INVALID_ARGUMENT if the token expired.
	// MUST return CODE_ALREADY_EXISTS if the user already accepted an invite.
	// MUST return CODE_PERMISSION_DENIED if the remote service is not trusted to accept invitations.
	ForwardInvite(ctx context.Context, in *ForwardInviteRequest, opts ...grpc.CallOption) (*ForwardInviteResponse, error)
	// Completes an invitation acceptance.
	// MUST return CODE_NOT_FOUND if the token does not exist.
	// MUST return CODE_INVALID_ARGUMENT if the token expired.
	// MUST return CODE_ALREADY_EXISTS if the user already accepted an invite.
	AcceptInvite(ctx context.Context, in *AcceptInviteRequest, opts ...grpc.CallOption) (*AcceptInviteResponse, error)
	// Retrieves details about a remote user who has accepted an invite to share.
	// MUST return CODE_NOT_FOUND if the user does not exist.
	GetAcceptedUser(ctx context.Context, in *GetAcceptedUserRequest, opts ...grpc.CallOption) (*GetAcceptedUserResponse, error)
	// Finds users who accepted invite tokens by their attributes.
	FindAcceptedUsers(ctx context.Context, in *FindAcceptedUsersRequest, opts ...grpc.CallOption) (*FindAcceptedUsersResponse, error)
	// Delete a previously accepted remote user, that is unfriend that user.
	// MUST return CODE_NOT_FOUND if the user does not exist.
	DeleteAcceptedUser(ctx context.Context, in *DeleteAcceptedUserRequest, opts ...grpc.CallOption) (*DeleteAcceptedUserResponse, error)
}

type inviteAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewInviteAPIClient(cc grpc.ClientConnInterface) InviteAPIClient {
	return &inviteAPIClient{cc}
}

func (c *inviteAPIClient) GenerateInviteToken(ctx context.Context, in *GenerateInviteTokenRequest, opts ...grpc.CallOption) (*GenerateInviteTokenResponse, error) {
	out := new(GenerateInviteTokenResponse)
	err := c.cc.Invoke(ctx, InviteAPI_GenerateInviteToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inviteAPIClient) ListInviteTokens(ctx context.Context, in *ListInviteTokensRequest, opts ...grpc.CallOption) (*ListInviteTokensResponse, error) {
	out := new(ListInviteTokensResponse)
	err := c.cc.Invoke(ctx, InviteAPI_ListInviteTokens_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inviteAPIClient) ForwardInvite(ctx context.Context, in *ForwardInviteRequest, opts ...grpc.CallOption) (*ForwardInviteResponse, error) {
	out := new(ForwardInviteResponse)
	err := c.cc.Invoke(ctx, InviteAPI_ForwardInvite_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inviteAPIClient) AcceptInvite(ctx context.Context, in *AcceptInviteRequest, opts ...grpc.CallOption) (*AcceptInviteResponse, error) {
	out := new(AcceptInviteResponse)
	err := c.cc.Invoke(ctx, InviteAPI_AcceptInvite_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inviteAPIClient) GetAcceptedUser(ctx context.Context, in *GetAcceptedUserRequest, opts ...grpc.CallOption) (*GetAcceptedUserResponse, error) {
	out := new(GetAcceptedUserResponse)
	err := c.cc.Invoke(ctx, InviteAPI_GetAcceptedUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inviteAPIClient) FindAcceptedUsers(ctx context.Context, in *FindAcceptedUsersRequest, opts ...grpc.CallOption) (*FindAcceptedUsersResponse, error) {
	out := new(FindAcceptedUsersResponse)
	err := c.cc.Invoke(ctx, InviteAPI_FindAcceptedUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inviteAPIClient) DeleteAcceptedUser(ctx context.Context, in *DeleteAcceptedUserRequest, opts ...grpc.CallOption) (*DeleteAcceptedUserResponse, error) {
	out := new(DeleteAcceptedUserResponse)
	err := c.cc.Invoke(ctx, InviteAPI_DeleteAcceptedUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InviteAPIServer is the server API for InviteAPI service.
// All implementations should embed UnimplementedInviteAPIServer
// for forward compatibility
type InviteAPIServer interface {
	// Generates a new token for the user with a validity of 24 hours.
	GenerateInviteToken(context.Context, *GenerateInviteTokenRequest) (*GenerateInviteTokenResponse, error)
	// Lists the valid tokens generated by the user.
	ListInviteTokens(context.Context, *ListInviteTokensRequest) (*ListInviteTokensResponse, error)
	// Forwards a received invite to the remote sync'n'share system provider. The remote
	// system SHALL get an `invite-accepted` call as follows:
	// https://cs3org.github.io/OCM-API/docs.html?branch=v1.2.0&repo=OCM-API&user=cs3org#/paths/~1invite-accepted/post
	// MUST return CODE_NOT_FOUND if the token does not exist.
	// MUST return CODE_INVALID_ARGUMENT if the token expired.
	// MUST return CODE_ALREADY_EXISTS if the user already accepted an invite.
	// MUST return CODE_PERMISSION_DENIED if the remote service is not trusted to accept invitations.
	ForwardInvite(context.Context, *ForwardInviteRequest) (*ForwardInviteResponse, error)
	// Completes an invitation acceptance.
	// MUST return CODE_NOT_FOUND if the token does not exist.
	// MUST return CODE_INVALID_ARGUMENT if the token expired.
	// MUST return CODE_ALREADY_EXISTS if the user already accepted an invite.
	AcceptInvite(context.Context, *AcceptInviteRequest) (*AcceptInviteResponse, error)
	// Retrieves details about a remote user who has accepted an invite to share.
	// MUST return CODE_NOT_FOUND if the user does not exist.
	GetAcceptedUser(context.Context, *GetAcceptedUserRequest) (*GetAcceptedUserResponse, error)
	// Finds users who accepted invite tokens by their attributes.
	FindAcceptedUsers(context.Context, *FindAcceptedUsersRequest) (*FindAcceptedUsersResponse, error)
	// Delete a previously accepted remote user, that is unfriend that user.
	// MUST return CODE_NOT_FOUND if the user does not exist.
	DeleteAcceptedUser(context.Context, *DeleteAcceptedUserRequest) (*DeleteAcceptedUserResponse, error)
}

// UnimplementedInviteAPIServer should be embedded to have forward compatible implementations.
type UnimplementedInviteAPIServer struct {
}

func (UnimplementedInviteAPIServer) GenerateInviteToken(context.Context, *GenerateInviteTokenRequest) (*GenerateInviteTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateInviteToken not implemented")
}
func (UnimplementedInviteAPIServer) ListInviteTokens(context.Context, *ListInviteTokensRequest) (*ListInviteTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListInviteTokens not implemented")
}
func (UnimplementedInviteAPIServer) ForwardInvite(context.Context, *ForwardInviteRequest) (*ForwardInviteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForwardInvite not implemented")
}
func (UnimplementedInviteAPIServer) AcceptInvite(context.Context, *AcceptInviteRequest) (*AcceptInviteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptInvite not implemented")
}
func (UnimplementedInviteAPIServer) GetAcceptedUser(context.Context, *GetAcceptedUserRequest) (*GetAcceptedUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAcceptedUser not implemented")
}
func (UnimplementedInviteAPIServer) FindAcceptedUsers(context.Context, *FindAcceptedUsersRequest) (*FindAcceptedUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAcceptedUsers not implemented")
}
func (UnimplementedInviteAPIServer) DeleteAcceptedUser(context.Context, *DeleteAcceptedUserRequest) (*DeleteAcceptedUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAcceptedUser not implemented")
}

// UnsafeInviteAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InviteAPIServer will
// result in compilation errors.
type UnsafeInviteAPIServer interface {
	mustEmbedUnimplementedInviteAPIServer()
}

func RegisterInviteAPIServer(s grpc.ServiceRegistrar, srv InviteAPIServer) {
	s.RegisterService(&InviteAPI_ServiceDesc, srv)
}

func _InviteAPI_GenerateInviteToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateInviteTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InviteAPIServer).GenerateInviteToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InviteAPI_GenerateInviteToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InviteAPIServer).GenerateInviteToken(ctx, req.(*GenerateInviteTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InviteAPI_ListInviteTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInviteTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InviteAPIServer).ListInviteTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InviteAPI_ListInviteTokens_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InviteAPIServer).ListInviteTokens(ctx, req.(*ListInviteTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InviteAPI_ForwardInvite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForwardInviteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InviteAPIServer).ForwardInvite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InviteAPI_ForwardInvite_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InviteAPIServer).ForwardInvite(ctx, req.(*ForwardInviteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InviteAPI_AcceptInvite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptInviteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InviteAPIServer).AcceptInvite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InviteAPI_AcceptInvite_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InviteAPIServer).AcceptInvite(ctx, req.(*AcceptInviteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InviteAPI_GetAcceptedUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAcceptedUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InviteAPIServer).GetAcceptedUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InviteAPI_GetAcceptedUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InviteAPIServer).GetAcceptedUser(ctx, req.(*GetAcceptedUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InviteAPI_FindAcceptedUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAcceptedUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InviteAPIServer).FindAcceptedUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InviteAPI_FindAcceptedUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InviteAPIServer).FindAcceptedUsers(ctx, req.(*FindAcceptedUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InviteAPI_DeleteAcceptedUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAcceptedUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InviteAPIServer).DeleteAcceptedUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InviteAPI_DeleteAcceptedUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InviteAPIServer).DeleteAcceptedUser(ctx, req.(*DeleteAcceptedUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InviteAPI_ServiceDesc is the grpc.ServiceDesc for InviteAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InviteAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cs3.ocm.invite.v1beta1.InviteAPI",
	HandlerType: (*InviteAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateInviteToken",
			Handler:    _InviteAPI_GenerateInviteToken_Handler,
		},
		{
			MethodName: "ListInviteTokens",
			Handler:    _InviteAPI_ListInviteTokens_Handler,
		},
		{
			MethodName: "ForwardInvite",
			Handler:    _InviteAPI_ForwardInvite_Handler,
		},
		{
			MethodName: "AcceptInvite",
			Handler:    _InviteAPI_AcceptInvite_Handler,
		},
		{
			MethodName: "GetAcceptedUser",
			Handler:    _InviteAPI_GetAcceptedUser_Handler,
		},
		{
			MethodName: "FindAcceptedUsers",
			Handler:    _InviteAPI_FindAcceptedUsers_Handler,
		},
		{
			MethodName: "DeleteAcceptedUser",
			Handler:    _InviteAPI_DeleteAcceptedUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cs3/ocm/invite/v1beta1/invite_api.proto",
}
