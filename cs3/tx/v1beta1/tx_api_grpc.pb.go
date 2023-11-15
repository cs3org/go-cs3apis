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
// source: cs3/tx/v1beta1/tx_api.proto

package txv1beta1

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
	TxAPI_CreateTransfer_FullMethodName    = "/cs3.tx.v1beta1.TxAPI/CreateTransfer"
	TxAPI_GetTransferStatus_FullMethodName = "/cs3.tx.v1beta1.TxAPI/GetTransferStatus"
	TxAPI_CancelTransfer_FullMethodName    = "/cs3.tx.v1beta1.TxAPI/CancelTransfer"
	TxAPI_ListTransfers_FullMethodName     = "/cs3.tx.v1beta1.TxAPI/ListTransfers"
	TxAPI_RetryTransfer_FullMethodName     = "/cs3.tx.v1beta1.TxAPI/RetryTransfer"
)

// TxAPIClient is the client API for TxAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TxAPIClient interface {
	// Requests creation of a transfer.
	// Returns a CreateTransferResponse.
	CreateTransfer(ctx context.Context, in *CreateTransferRequest, opts ...grpc.CallOption) (*CreateTransferResponse, error)
	// Requests a transfer status.
	GetTransferStatus(ctx context.Context, in *GetTransferStatusRequest, opts ...grpc.CallOption) (*GetTransferStatusResponse, error)
	// Requests to cancel a transfer.
	CancelTransfer(ctx context.Context, in *CancelTransferRequest, opts ...grpc.CallOption) (*CancelTransferResponse, error)
	// Requests a list of transfers received by the authenticated principle.
	// If a filter is specified, only transfers satisfying the filter MUST be returned.
	ListTransfers(ctx context.Context, in *ListTransfersRequest, opts ...grpc.CallOption) (*ListTransfersResponse, error)
	// Requests retrying a transfer.
	RetryTransfer(ctx context.Context, in *RetryTransferRequest, opts ...grpc.CallOption) (*RetryTransferResponse, error)
}

type txAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewTxAPIClient(cc grpc.ClientConnInterface) TxAPIClient {
	return &txAPIClient{cc}
}

func (c *txAPIClient) CreateTransfer(ctx context.Context, in *CreateTransferRequest, opts ...grpc.CallOption) (*CreateTransferResponse, error) {
	out := new(CreateTransferResponse)
	err := c.cc.Invoke(ctx, TxAPI_CreateTransfer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *txAPIClient) GetTransferStatus(ctx context.Context, in *GetTransferStatusRequest, opts ...grpc.CallOption) (*GetTransferStatusResponse, error) {
	out := new(GetTransferStatusResponse)
	err := c.cc.Invoke(ctx, TxAPI_GetTransferStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *txAPIClient) CancelTransfer(ctx context.Context, in *CancelTransferRequest, opts ...grpc.CallOption) (*CancelTransferResponse, error) {
	out := new(CancelTransferResponse)
	err := c.cc.Invoke(ctx, TxAPI_CancelTransfer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *txAPIClient) ListTransfers(ctx context.Context, in *ListTransfersRequest, opts ...grpc.CallOption) (*ListTransfersResponse, error) {
	out := new(ListTransfersResponse)
	err := c.cc.Invoke(ctx, TxAPI_ListTransfers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *txAPIClient) RetryTransfer(ctx context.Context, in *RetryTransferRequest, opts ...grpc.CallOption) (*RetryTransferResponse, error) {
	out := new(RetryTransferResponse)
	err := c.cc.Invoke(ctx, TxAPI_RetryTransfer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TxAPIServer is the server API for TxAPI service.
// All implementations should embed UnimplementedTxAPIServer
// for forward compatibility
type TxAPIServer interface {
	// Requests creation of a transfer.
	// Returns a CreateTransferResponse.
	CreateTransfer(context.Context, *CreateTransferRequest) (*CreateTransferResponse, error)
	// Requests a transfer status.
	GetTransferStatus(context.Context, *GetTransferStatusRequest) (*GetTransferStatusResponse, error)
	// Requests to cancel a transfer.
	CancelTransfer(context.Context, *CancelTransferRequest) (*CancelTransferResponse, error)
	// Requests a list of transfers received by the authenticated principle.
	// If a filter is specified, only transfers satisfying the filter MUST be returned.
	ListTransfers(context.Context, *ListTransfersRequest) (*ListTransfersResponse, error)
	// Requests retrying a transfer.
	RetryTransfer(context.Context, *RetryTransferRequest) (*RetryTransferResponse, error)
}

// UnimplementedTxAPIServer should be embedded to have forward compatible implementations.
type UnimplementedTxAPIServer struct {
}

func (UnimplementedTxAPIServer) CreateTransfer(context.Context, *CreateTransferRequest) (*CreateTransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransfer not implemented")
}
func (UnimplementedTxAPIServer) GetTransferStatus(context.Context, *GetTransferStatusRequest) (*GetTransferStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransferStatus not implemented")
}
func (UnimplementedTxAPIServer) CancelTransfer(context.Context, *CancelTransferRequest) (*CancelTransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelTransfer not implemented")
}
func (UnimplementedTxAPIServer) ListTransfers(context.Context, *ListTransfersRequest) (*ListTransfersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTransfers not implemented")
}
func (UnimplementedTxAPIServer) RetryTransfer(context.Context, *RetryTransferRequest) (*RetryTransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetryTransfer not implemented")
}

// UnsafeTxAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TxAPIServer will
// result in compilation errors.
type UnsafeTxAPIServer interface {
	mustEmbedUnimplementedTxAPIServer()
}

func RegisterTxAPIServer(s grpc.ServiceRegistrar, srv TxAPIServer) {
	s.RegisterService(&TxAPI_ServiceDesc, srv)
}

func _TxAPI_CreateTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxAPIServer).CreateTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TxAPI_CreateTransfer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxAPIServer).CreateTransfer(ctx, req.(*CreateTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TxAPI_GetTransferStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransferStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxAPIServer).GetTransferStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TxAPI_GetTransferStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxAPIServer).GetTransferStatus(ctx, req.(*GetTransferStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TxAPI_CancelTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxAPIServer).CancelTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TxAPI_CancelTransfer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxAPIServer).CancelTransfer(ctx, req.(*CancelTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TxAPI_ListTransfers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTransfersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxAPIServer).ListTransfers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TxAPI_ListTransfers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxAPIServer).ListTransfers(ctx, req.(*ListTransfersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TxAPI_RetryTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetryTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxAPIServer).RetryTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TxAPI_RetryTransfer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxAPIServer).RetryTransfer(ctx, req.(*RetryTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TxAPI_ServiceDesc is the grpc.ServiceDesc for TxAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TxAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cs3.tx.v1beta1.TxAPI",
	HandlerType: (*TxAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTransfer",
			Handler:    _TxAPI_CreateTransfer_Handler,
		},
		{
			MethodName: "GetTransferStatus",
			Handler:    _TxAPI_GetTransferStatus_Handler,
		},
		{
			MethodName: "CancelTransfer",
			Handler:    _TxAPI_CancelTransfer_Handler,
		},
		{
			MethodName: "ListTransfers",
			Handler:    _TxAPI_ListTransfers_Handler,
		},
		{
			MethodName: "RetryTransfer",
			Handler:    _TxAPI_RetryTransfer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cs3/tx/v1beta1/tx_api.proto",
}
