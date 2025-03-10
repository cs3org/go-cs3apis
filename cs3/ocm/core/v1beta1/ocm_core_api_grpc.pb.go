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
// source: cs3/ocm/core/v1beta1/ocm_core_api.proto

package corev1beta1

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
	OcmCoreAPI_CreateOCMCoreShare_FullMethodName = "/cs3.ocm.core.v1beta1.OcmCoreAPI/CreateOCMCoreShare"
	OcmCoreAPI_UpdateOCMCoreShare_FullMethodName = "/cs3.ocm.core.v1beta1.OcmCoreAPI/UpdateOCMCoreShare"
	OcmCoreAPI_DeleteOCMCoreShare_FullMethodName = "/cs3.ocm.core.v1beta1.OcmCoreAPI/DeleteOCMCoreShare"
)

// OcmCoreAPIClient is the client API for OcmCoreAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OcmCoreAPIClient interface {
	// Creates a new OCM share, in response to a call from remote to:
	// https://cs3org.github.io/OCM-API/docs.html?branch=v1.2.0&repo=OCM-API&user=cs3org#/paths/~1shares/post
	CreateOCMCoreShare(ctx context.Context, in *CreateOCMCoreShareRequest, opts ...grpc.CallOption) (*CreateOCMCoreShareResponse, error)
	// Updates an OCM share, in response to a notification from the remote system to:
	// https://cs3org.github.io/OCM-API/docs.html?branch=v1.2.0&repo=OCM-API&user=cs3org#/paths/~1notifications/post
	UpdateOCMCoreShare(ctx context.Context, in *UpdateOCMCoreShareRequest, opts ...grpc.CallOption) (*UpdateOCMCoreShareResponse, error)
	// Deletes an OCM share, in response to a notification from the remote system to:
	// https://cs3org.github.io/OCM-API/docs.html?branch=v1.2.0&repo=OCM-API&user=cs3org#/paths/~1notifications/post
	DeleteOCMCoreShare(ctx context.Context, in *DeleteOCMCoreShareRequest, opts ...grpc.CallOption) (*DeleteOCMCoreShareResponse, error)
}

type ocmCoreAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewOcmCoreAPIClient(cc grpc.ClientConnInterface) OcmCoreAPIClient {
	return &ocmCoreAPIClient{cc}
}

func (c *ocmCoreAPIClient) CreateOCMCoreShare(ctx context.Context, in *CreateOCMCoreShareRequest, opts ...grpc.CallOption) (*CreateOCMCoreShareResponse, error) {
	out := new(CreateOCMCoreShareResponse)
	err := c.cc.Invoke(ctx, OcmCoreAPI_CreateOCMCoreShare_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocmCoreAPIClient) UpdateOCMCoreShare(ctx context.Context, in *UpdateOCMCoreShareRequest, opts ...grpc.CallOption) (*UpdateOCMCoreShareResponse, error) {
	out := new(UpdateOCMCoreShareResponse)
	err := c.cc.Invoke(ctx, OcmCoreAPI_UpdateOCMCoreShare_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocmCoreAPIClient) DeleteOCMCoreShare(ctx context.Context, in *DeleteOCMCoreShareRequest, opts ...grpc.CallOption) (*DeleteOCMCoreShareResponse, error) {
	out := new(DeleteOCMCoreShareResponse)
	err := c.cc.Invoke(ctx, OcmCoreAPI_DeleteOCMCoreShare_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OcmCoreAPIServer is the server API for OcmCoreAPI service.
// All implementations should embed UnimplementedOcmCoreAPIServer
// for forward compatibility
type OcmCoreAPIServer interface {
	// Creates a new OCM share, in response to a call from remote to:
	// https://cs3org.github.io/OCM-API/docs.html?branch=v1.2.0&repo=OCM-API&user=cs3org#/paths/~1shares/post
	CreateOCMCoreShare(context.Context, *CreateOCMCoreShareRequest) (*CreateOCMCoreShareResponse, error)
	// Updates an OCM share, in response to a notification from the remote system to:
	// https://cs3org.github.io/OCM-API/docs.html?branch=v1.2.0&repo=OCM-API&user=cs3org#/paths/~1notifications/post
	UpdateOCMCoreShare(context.Context, *UpdateOCMCoreShareRequest) (*UpdateOCMCoreShareResponse, error)
	// Deletes an OCM share, in response to a notification from the remote system to:
	// https://cs3org.github.io/OCM-API/docs.html?branch=v1.2.0&repo=OCM-API&user=cs3org#/paths/~1notifications/post
	DeleteOCMCoreShare(context.Context, *DeleteOCMCoreShareRequest) (*DeleteOCMCoreShareResponse, error)
}

// UnimplementedOcmCoreAPIServer should be embedded to have forward compatible implementations.
type UnimplementedOcmCoreAPIServer struct {
}

func (UnimplementedOcmCoreAPIServer) CreateOCMCoreShare(context.Context, *CreateOCMCoreShareRequest) (*CreateOCMCoreShareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOCMCoreShare not implemented")
}
func (UnimplementedOcmCoreAPIServer) UpdateOCMCoreShare(context.Context, *UpdateOCMCoreShareRequest) (*UpdateOCMCoreShareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOCMCoreShare not implemented")
}
func (UnimplementedOcmCoreAPIServer) DeleteOCMCoreShare(context.Context, *DeleteOCMCoreShareRequest) (*DeleteOCMCoreShareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOCMCoreShare not implemented")
}

// UnsafeOcmCoreAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OcmCoreAPIServer will
// result in compilation errors.
type UnsafeOcmCoreAPIServer interface {
	mustEmbedUnimplementedOcmCoreAPIServer()
}

func RegisterOcmCoreAPIServer(s grpc.ServiceRegistrar, srv OcmCoreAPIServer) {
	s.RegisterService(&OcmCoreAPI_ServiceDesc, srv)
}

func _OcmCoreAPI_CreateOCMCoreShare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOCMCoreShareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcmCoreAPIServer).CreateOCMCoreShare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OcmCoreAPI_CreateOCMCoreShare_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcmCoreAPIServer).CreateOCMCoreShare(ctx, req.(*CreateOCMCoreShareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcmCoreAPI_UpdateOCMCoreShare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOCMCoreShareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcmCoreAPIServer).UpdateOCMCoreShare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OcmCoreAPI_UpdateOCMCoreShare_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcmCoreAPIServer).UpdateOCMCoreShare(ctx, req.(*UpdateOCMCoreShareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcmCoreAPI_DeleteOCMCoreShare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOCMCoreShareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcmCoreAPIServer).DeleteOCMCoreShare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OcmCoreAPI_DeleteOCMCoreShare_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcmCoreAPIServer).DeleteOCMCoreShare(ctx, req.(*DeleteOCMCoreShareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OcmCoreAPI_ServiceDesc is the grpc.ServiceDesc for OcmCoreAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OcmCoreAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cs3.ocm.core.v1beta1.OcmCoreAPI",
	HandlerType: (*OcmCoreAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOCMCoreShare",
			Handler:    _OcmCoreAPI_CreateOCMCoreShare_Handler,
		},
		{
			MethodName: "UpdateOCMCoreShare",
			Handler:    _OcmCoreAPI_UpdateOCMCoreShare_Handler,
		},
		{
			MethodName: "DeleteOCMCoreShare",
			Handler:    _OcmCoreAPI_DeleteOCMCoreShare_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cs3/ocm/core/v1beta1/ocm_core_api.proto",
}
