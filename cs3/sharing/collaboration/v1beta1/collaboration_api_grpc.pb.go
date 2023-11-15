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
// source: cs3/sharing/collaboration/v1beta1/collaboration_api.proto

package collaborationv1beta1

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
	CollaborationAPI_CreateShare_FullMethodName         = "/cs3.sharing.collaboration.v1beta1.CollaborationAPI/CreateShare"
	CollaborationAPI_RemoveShare_FullMethodName         = "/cs3.sharing.collaboration.v1beta1.CollaborationAPI/RemoveShare"
	CollaborationAPI_GetShare_FullMethodName            = "/cs3.sharing.collaboration.v1beta1.CollaborationAPI/GetShare"
	CollaborationAPI_ListShares_FullMethodName          = "/cs3.sharing.collaboration.v1beta1.CollaborationAPI/ListShares"
	CollaborationAPI_UpdateShare_FullMethodName         = "/cs3.sharing.collaboration.v1beta1.CollaborationAPI/UpdateShare"
	CollaborationAPI_ListReceivedShares_FullMethodName  = "/cs3.sharing.collaboration.v1beta1.CollaborationAPI/ListReceivedShares"
	CollaborationAPI_UpdateReceivedShare_FullMethodName = "/cs3.sharing.collaboration.v1beta1.CollaborationAPI/UpdateReceivedShare"
	CollaborationAPI_GetReceivedShare_FullMethodName    = "/cs3.sharing.collaboration.v1beta1.CollaborationAPI/GetReceivedShare"
)

// CollaborationAPIClient is the client API for CollaborationAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CollaborationAPIClient interface {
	// Creates a new share.
	// MUST return CODE_NOT_FOUND if the resource reference does not exist.
	// MUST return CODE_LOCKED if the resource reference already locked.
	// MUST return CODE_ALREADY_EXISTS if the share already exists for the 4-tuple consisting of
	// (owner, shared_resource, grantee).
	// New shares MUST be created in the state SHARE_STATE_PENDING.
	CreateShare(ctx context.Context, in *CreateShareRequest, opts ...grpc.CallOption) (*CreateShareResponse, error)
	// Removes a share.
	// MUST return CODE_NOT_FOUND if the share reference does not exist.
	RemoveShare(ctx context.Context, in *RemoveShareRequest, opts ...grpc.CallOption) (*RemoveShareResponse, error)
	// Gets share information for a single share.
	// MUST return CODE_NOT_FOUND if the share reference does not exist.
	GetShare(ctx context.Context, in *GetShareRequest, opts ...grpc.CallOption) (*GetShareResponse, error)
	// List the shares the authenticated principal has created,
	// both as owner and creator. If a filter is specified, only
	// shares satisfying the filter MUST be returned.
	ListShares(ctx context.Context, in *ListSharesRequest, opts ...grpc.CallOption) (*ListSharesResponse, error)
	// Updates a share.
	// MUST return CODE_NOT_FOUND if the share reference does not exist.
	UpdateShare(ctx context.Context, in *UpdateShareRequest, opts ...grpc.CallOption) (*UpdateShareResponse, error)
	// List all shares the authenticated principal has received.
	ListReceivedShares(ctx context.Context, in *ListReceivedSharesRequest, opts ...grpc.CallOption) (*ListReceivedSharesResponse, error)
	// Update the received share to change the share state or the display name.
	// MUST return CODE_NOT_FOUND if the share reference does not exist.
	UpdateReceivedShare(ctx context.Context, in *UpdateReceivedShareRequest, opts ...grpc.CallOption) (*UpdateReceivedShareResponse, error)
	// Get the information for the given received share reference.
	// MUST return CODE_NOT_FOUND if the received share reference does not exist.
	GetReceivedShare(ctx context.Context, in *GetReceivedShareRequest, opts ...grpc.CallOption) (*GetReceivedShareResponse, error)
}

type collaborationAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewCollaborationAPIClient(cc grpc.ClientConnInterface) CollaborationAPIClient {
	return &collaborationAPIClient{cc}
}

func (c *collaborationAPIClient) CreateShare(ctx context.Context, in *CreateShareRequest, opts ...grpc.CallOption) (*CreateShareResponse, error) {
	out := new(CreateShareResponse)
	err := c.cc.Invoke(ctx, CollaborationAPI_CreateShare_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaborationAPIClient) RemoveShare(ctx context.Context, in *RemoveShareRequest, opts ...grpc.CallOption) (*RemoveShareResponse, error) {
	out := new(RemoveShareResponse)
	err := c.cc.Invoke(ctx, CollaborationAPI_RemoveShare_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaborationAPIClient) GetShare(ctx context.Context, in *GetShareRequest, opts ...grpc.CallOption) (*GetShareResponse, error) {
	out := new(GetShareResponse)
	err := c.cc.Invoke(ctx, CollaborationAPI_GetShare_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaborationAPIClient) ListShares(ctx context.Context, in *ListSharesRequest, opts ...grpc.CallOption) (*ListSharesResponse, error) {
	out := new(ListSharesResponse)
	err := c.cc.Invoke(ctx, CollaborationAPI_ListShares_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaborationAPIClient) UpdateShare(ctx context.Context, in *UpdateShareRequest, opts ...grpc.CallOption) (*UpdateShareResponse, error) {
	out := new(UpdateShareResponse)
	err := c.cc.Invoke(ctx, CollaborationAPI_UpdateShare_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaborationAPIClient) ListReceivedShares(ctx context.Context, in *ListReceivedSharesRequest, opts ...grpc.CallOption) (*ListReceivedSharesResponse, error) {
	out := new(ListReceivedSharesResponse)
	err := c.cc.Invoke(ctx, CollaborationAPI_ListReceivedShares_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaborationAPIClient) UpdateReceivedShare(ctx context.Context, in *UpdateReceivedShareRequest, opts ...grpc.CallOption) (*UpdateReceivedShareResponse, error) {
	out := new(UpdateReceivedShareResponse)
	err := c.cc.Invoke(ctx, CollaborationAPI_UpdateReceivedShare_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaborationAPIClient) GetReceivedShare(ctx context.Context, in *GetReceivedShareRequest, opts ...grpc.CallOption) (*GetReceivedShareResponse, error) {
	out := new(GetReceivedShareResponse)
	err := c.cc.Invoke(ctx, CollaborationAPI_GetReceivedShare_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CollaborationAPIServer is the server API for CollaborationAPI service.
// All implementations must embed UnimplementedCollaborationAPIServer
// for forward compatibility
type CollaborationAPIServer interface {
	// Creates a new share.
	// MUST return CODE_NOT_FOUND if the resource reference does not exist.
	// MUST return CODE_LOCKED if the resource reference already locked.
	// MUST return CODE_ALREADY_EXISTS if the share already exists for the 4-tuple consisting of
	// (owner, shared_resource, grantee).
	// New shares MUST be created in the state SHARE_STATE_PENDING.
	CreateShare(context.Context, *CreateShareRequest) (*CreateShareResponse, error)
	// Removes a share.
	// MUST return CODE_NOT_FOUND if the share reference does not exist.
	RemoveShare(context.Context, *RemoveShareRequest) (*RemoveShareResponse, error)
	// Gets share information for a single share.
	// MUST return CODE_NOT_FOUND if the share reference does not exist.
	GetShare(context.Context, *GetShareRequest) (*GetShareResponse, error)
	// List the shares the authenticated principal has created,
	// both as owner and creator. If a filter is specified, only
	// shares satisfying the filter MUST be returned.
	ListShares(context.Context, *ListSharesRequest) (*ListSharesResponse, error)
	// Updates a share.
	// MUST return CODE_NOT_FOUND if the share reference does not exist.
	UpdateShare(context.Context, *UpdateShareRequest) (*UpdateShareResponse, error)
	// List all shares the authenticated principal has received.
	ListReceivedShares(context.Context, *ListReceivedSharesRequest) (*ListReceivedSharesResponse, error)
	// Update the received share to change the share state or the display name.
	// MUST return CODE_NOT_FOUND if the share reference does not exist.
	UpdateReceivedShare(context.Context, *UpdateReceivedShareRequest) (*UpdateReceivedShareResponse, error)
	// Get the information for the given received share reference.
	// MUST return CODE_NOT_FOUND if the received share reference does not exist.
	GetReceivedShare(context.Context, *GetReceivedShareRequest) (*GetReceivedShareResponse, error)
	mustEmbedUnimplementedCollaborationAPIServer()
}

// UnimplementedCollaborationAPIServer must be embedded to have forward compatible implementations.
type UnimplementedCollaborationAPIServer struct {
}

func (UnimplementedCollaborationAPIServer) CreateShare(context.Context, *CreateShareRequest) (*CreateShareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShare not implemented")
}
func (UnimplementedCollaborationAPIServer) RemoveShare(context.Context, *RemoveShareRequest) (*RemoveShareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveShare not implemented")
}
func (UnimplementedCollaborationAPIServer) GetShare(context.Context, *GetShareRequest) (*GetShareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShare not implemented")
}
func (UnimplementedCollaborationAPIServer) ListShares(context.Context, *ListSharesRequest) (*ListSharesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListShares not implemented")
}
func (UnimplementedCollaborationAPIServer) UpdateShare(context.Context, *UpdateShareRequest) (*UpdateShareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateShare not implemented")
}
func (UnimplementedCollaborationAPIServer) ListReceivedShares(context.Context, *ListReceivedSharesRequest) (*ListReceivedSharesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListReceivedShares not implemented")
}
func (UnimplementedCollaborationAPIServer) UpdateReceivedShare(context.Context, *UpdateReceivedShareRequest) (*UpdateReceivedShareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateReceivedShare not implemented")
}
func (UnimplementedCollaborationAPIServer) GetReceivedShare(context.Context, *GetReceivedShareRequest) (*GetReceivedShareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReceivedShare not implemented")
}
func (UnimplementedCollaborationAPIServer) mustEmbedUnimplementedCollaborationAPIServer() {}

// UnsafeCollaborationAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CollaborationAPIServer will
// result in compilation errors.
type UnsafeCollaborationAPIServer interface {
	mustEmbedUnimplementedCollaborationAPIServer()
}

func RegisterCollaborationAPIServer(s grpc.ServiceRegistrar, srv CollaborationAPIServer) {
	s.RegisterService(&CollaborationAPI_ServiceDesc, srv)
}

func _CollaborationAPI_CreateShare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateShareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollaborationAPIServer).CreateShare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CollaborationAPI_CreateShare_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollaborationAPIServer).CreateShare(ctx, req.(*CreateShareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollaborationAPI_RemoveShare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveShareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollaborationAPIServer).RemoveShare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CollaborationAPI_RemoveShare_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollaborationAPIServer).RemoveShare(ctx, req.(*RemoveShareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollaborationAPI_GetShare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollaborationAPIServer).GetShare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CollaborationAPI_GetShare_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollaborationAPIServer).GetShare(ctx, req.(*GetShareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollaborationAPI_ListShares_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSharesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollaborationAPIServer).ListShares(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CollaborationAPI_ListShares_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollaborationAPIServer).ListShares(ctx, req.(*ListSharesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollaborationAPI_UpdateShare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateShareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollaborationAPIServer).UpdateShare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CollaborationAPI_UpdateShare_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollaborationAPIServer).UpdateShare(ctx, req.(*UpdateShareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollaborationAPI_ListReceivedShares_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReceivedSharesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollaborationAPIServer).ListReceivedShares(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CollaborationAPI_ListReceivedShares_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollaborationAPIServer).ListReceivedShares(ctx, req.(*ListReceivedSharesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollaborationAPI_UpdateReceivedShare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReceivedShareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollaborationAPIServer).UpdateReceivedShare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CollaborationAPI_UpdateReceivedShare_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollaborationAPIServer).UpdateReceivedShare(ctx, req.(*UpdateReceivedShareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollaborationAPI_GetReceivedShare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReceivedShareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollaborationAPIServer).GetReceivedShare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CollaborationAPI_GetReceivedShare_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollaborationAPIServer).GetReceivedShare(ctx, req.(*GetReceivedShareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CollaborationAPI_ServiceDesc is the grpc.ServiceDesc for CollaborationAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CollaborationAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cs3.sharing.collaboration.v1beta1.CollaborationAPI",
	HandlerType: (*CollaborationAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateShare",
			Handler:    _CollaborationAPI_CreateShare_Handler,
		},
		{
			MethodName: "RemoveShare",
			Handler:    _CollaborationAPI_RemoveShare_Handler,
		},
		{
			MethodName: "GetShare",
			Handler:    _CollaborationAPI_GetShare_Handler,
		},
		{
			MethodName: "ListShares",
			Handler:    _CollaborationAPI_ListShares_Handler,
		},
		{
			MethodName: "UpdateShare",
			Handler:    _CollaborationAPI_UpdateShare_Handler,
		},
		{
			MethodName: "ListReceivedShares",
			Handler:    _CollaborationAPI_ListReceivedShares_Handler,
		},
		{
			MethodName: "UpdateReceivedShare",
			Handler:    _CollaborationAPI_UpdateReceivedShare_Handler,
		},
		{
			MethodName: "GetReceivedShare",
			Handler:    _CollaborationAPI_GetReceivedShare_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cs3/sharing/collaboration/v1beta1/collaboration_api.proto",
}
