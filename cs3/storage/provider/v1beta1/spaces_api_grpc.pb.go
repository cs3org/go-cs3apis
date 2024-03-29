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
// source: cs3/storage/provider/v1beta1/spaces_api.proto

package providerv1beta1

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
	SpacesAPI_CreateStorageSpace_FullMethodName = "/cs3.storage.provider.v1beta1.SpacesAPI/CreateStorageSpace"
	SpacesAPI_ListStorageSpaces_FullMethodName  = "/cs3.storage.provider.v1beta1.SpacesAPI/ListStorageSpaces"
	SpacesAPI_UpdateStorageSpace_FullMethodName = "/cs3.storage.provider.v1beta1.SpacesAPI/UpdateStorageSpace"
	SpacesAPI_DeleteStorageSpace_FullMethodName = "/cs3.storage.provider.v1beta1.SpacesAPI/DeleteStorageSpace"
)

// SpacesAPIClient is the client API for SpacesAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SpacesAPIClient interface {
	// Creates a storage space.
	CreateStorageSpace(ctx context.Context, in *CreateStorageSpaceRequest, opts ...grpc.CallOption) (*CreateStorageSpaceResponse, error)
	// Lists storage spaces.
	ListStorageSpaces(ctx context.Context, in *ListStorageSpacesRequest, opts ...grpc.CallOption) (*ListStorageSpacesResponse, error)
	// Updates a storage space.
	UpdateStorageSpace(ctx context.Context, in *UpdateStorageSpaceRequest, opts ...grpc.CallOption) (*UpdateStorageSpaceResponse, error)
	// Deletes a storage space.
	DeleteStorageSpace(ctx context.Context, in *DeleteStorageSpaceRequest, opts ...grpc.CallOption) (*DeleteStorageSpaceResponse, error)
}

type spacesAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewSpacesAPIClient(cc grpc.ClientConnInterface) SpacesAPIClient {
	return &spacesAPIClient{cc}
}

func (c *spacesAPIClient) CreateStorageSpace(ctx context.Context, in *CreateStorageSpaceRequest, opts ...grpc.CallOption) (*CreateStorageSpaceResponse, error) {
	out := new(CreateStorageSpaceResponse)
	err := c.cc.Invoke(ctx, SpacesAPI_CreateStorageSpace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spacesAPIClient) ListStorageSpaces(ctx context.Context, in *ListStorageSpacesRequest, opts ...grpc.CallOption) (*ListStorageSpacesResponse, error) {
	out := new(ListStorageSpacesResponse)
	err := c.cc.Invoke(ctx, SpacesAPI_ListStorageSpaces_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spacesAPIClient) UpdateStorageSpace(ctx context.Context, in *UpdateStorageSpaceRequest, opts ...grpc.CallOption) (*UpdateStorageSpaceResponse, error) {
	out := new(UpdateStorageSpaceResponse)
	err := c.cc.Invoke(ctx, SpacesAPI_UpdateStorageSpace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spacesAPIClient) DeleteStorageSpace(ctx context.Context, in *DeleteStorageSpaceRequest, opts ...grpc.CallOption) (*DeleteStorageSpaceResponse, error) {
	out := new(DeleteStorageSpaceResponse)
	err := c.cc.Invoke(ctx, SpacesAPI_DeleteStorageSpace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpacesAPIServer is the server API for SpacesAPI service.
// All implementations should embed UnimplementedSpacesAPIServer
// for forward compatibility
type SpacesAPIServer interface {
	// Creates a storage space.
	CreateStorageSpace(context.Context, *CreateStorageSpaceRequest) (*CreateStorageSpaceResponse, error)
	// Lists storage spaces.
	ListStorageSpaces(context.Context, *ListStorageSpacesRequest) (*ListStorageSpacesResponse, error)
	// Updates a storage space.
	UpdateStorageSpace(context.Context, *UpdateStorageSpaceRequest) (*UpdateStorageSpaceResponse, error)
	// Deletes a storage space.
	DeleteStorageSpace(context.Context, *DeleteStorageSpaceRequest) (*DeleteStorageSpaceResponse, error)
}

// UnimplementedSpacesAPIServer should be embedded to have forward compatible implementations.
type UnimplementedSpacesAPIServer struct {
}

func (UnimplementedSpacesAPIServer) CreateStorageSpace(context.Context, *CreateStorageSpaceRequest) (*CreateStorageSpaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStorageSpace not implemented")
}
func (UnimplementedSpacesAPIServer) ListStorageSpaces(context.Context, *ListStorageSpacesRequest) (*ListStorageSpacesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStorageSpaces not implemented")
}
func (UnimplementedSpacesAPIServer) UpdateStorageSpace(context.Context, *UpdateStorageSpaceRequest) (*UpdateStorageSpaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStorageSpace not implemented")
}
func (UnimplementedSpacesAPIServer) DeleteStorageSpace(context.Context, *DeleteStorageSpaceRequest) (*DeleteStorageSpaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStorageSpace not implemented")
}

// UnsafeSpacesAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SpacesAPIServer will
// result in compilation errors.
type UnsafeSpacesAPIServer interface {
	mustEmbedUnimplementedSpacesAPIServer()
}

func RegisterSpacesAPIServer(s grpc.ServiceRegistrar, srv SpacesAPIServer) {
	s.RegisterService(&SpacesAPI_ServiceDesc, srv)
}

func _SpacesAPI_CreateStorageSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStorageSpaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpacesAPIServer).CreateStorageSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SpacesAPI_CreateStorageSpace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpacesAPIServer).CreateStorageSpace(ctx, req.(*CreateStorageSpaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpacesAPI_ListStorageSpaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStorageSpacesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpacesAPIServer).ListStorageSpaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SpacesAPI_ListStorageSpaces_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpacesAPIServer).ListStorageSpaces(ctx, req.(*ListStorageSpacesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpacesAPI_UpdateStorageSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStorageSpaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpacesAPIServer).UpdateStorageSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SpacesAPI_UpdateStorageSpace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpacesAPIServer).UpdateStorageSpace(ctx, req.(*UpdateStorageSpaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpacesAPI_DeleteStorageSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStorageSpaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpacesAPIServer).DeleteStorageSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SpacesAPI_DeleteStorageSpace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpacesAPIServer).DeleteStorageSpace(ctx, req.(*DeleteStorageSpaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SpacesAPI_ServiceDesc is the grpc.ServiceDesc for SpacesAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SpacesAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cs3.storage.provider.v1beta1.SpacesAPI",
	HandlerType: (*SpacesAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStorageSpace",
			Handler:    _SpacesAPI_CreateStorageSpace_Handler,
		},
		{
			MethodName: "ListStorageSpaces",
			Handler:    _SpacesAPI_ListStorageSpaces_Handler,
		},
		{
			MethodName: "UpdateStorageSpace",
			Handler:    _SpacesAPI_UpdateStorageSpace_Handler,
		},
		{
			MethodName: "DeleteStorageSpace",
			Handler:    _SpacesAPI_DeleteStorageSpace_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cs3/storage/provider/v1beta1/spaces_api.proto",
}
