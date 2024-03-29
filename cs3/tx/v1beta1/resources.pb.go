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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cs3/tx/v1beta1/resources.proto

package txv1beta1

import (
	v1beta11 "github.com/cs3org/go-cs3apis/cs3/identity/user/v1beta1"
	v1beta13 "github.com/cs3org/go-cs3apis/cs3/sharing/ocm/v1beta1"
	v1beta1 "github.com/cs3org/go-cs3apis/cs3/storage/provider/v1beta1"
	v1beta12 "github.com/cs3org/go-cs3apis/cs3/types/v1beta1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Status represents transfer status.
type Status int32

const (
	Status_STATUS_INVALID Status = 0
	// The destination could not be found.
	Status_STATUS_DESTINATION_NOT_FOUND Status = 1
	// A new data transfer
	Status_STATUS_TRANSFER_NEW Status = 2
	// The data transfer is awaiting acceptance from the destination
	Status_STATUS_TRANSFER_AWAITING_ACCEPTANCE Status = 3
	// The data transfer is accepted by the destination.
	Status_STATUS_TRANSFER_ACCEPTED Status = 4
	// The data transfer has started and not yet completed.
	Status_STATUS_TRANSFER_IN_PROGRESS Status = 5
	// The data transfer has completed.
	Status_STATUS_TRANSFER_COMPLETE Status = 6
	// The data transfer has failed.
	Status_STATUS_TRANSFER_FAILED Status = 7
	// The data transfer had been cancelled.
	Status_STATUS_TRANSFER_CANCELLED Status = 8
	// The request for cancelling the data transfer has failed.
	Status_STATUS_TRANSFER_CANCEL_FAILED Status = 9
	// The transfer has expired somewhere down the line.
	Status_STATUS_TRANSFER_EXPIRED Status = 10
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0:  "STATUS_INVALID",
		1:  "STATUS_DESTINATION_NOT_FOUND",
		2:  "STATUS_TRANSFER_NEW",
		3:  "STATUS_TRANSFER_AWAITING_ACCEPTANCE",
		4:  "STATUS_TRANSFER_ACCEPTED",
		5:  "STATUS_TRANSFER_IN_PROGRESS",
		6:  "STATUS_TRANSFER_COMPLETE",
		7:  "STATUS_TRANSFER_FAILED",
		8:  "STATUS_TRANSFER_CANCELLED",
		9:  "STATUS_TRANSFER_CANCEL_FAILED",
		10: "STATUS_TRANSFER_EXPIRED",
	}
	Status_value = map[string]int32{
		"STATUS_INVALID":                      0,
		"STATUS_DESTINATION_NOT_FOUND":        1,
		"STATUS_TRANSFER_NEW":                 2,
		"STATUS_TRANSFER_AWAITING_ACCEPTANCE": 3,
		"STATUS_TRANSFER_ACCEPTED":            4,
		"STATUS_TRANSFER_IN_PROGRESS":         5,
		"STATUS_TRANSFER_COMPLETE":            6,
		"STATUS_TRANSFER_FAILED":              7,
		"STATUS_TRANSFER_CANCELLED":           8,
		"STATUS_TRANSFER_CANCEL_FAILED":       9,
		"STATUS_TRANSFER_EXPIRED":             10,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_cs3_tx_v1beta1_resources_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_cs3_tx_v1beta1_resources_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_cs3_tx_v1beta1_resources_proto_rawDescGZIP(), []int{0}
}

// TxId uniquely identifies a transfer in the transfer provider namespace.
type TxId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// REQUIRED.
	// The internal transfer id used by the service implementor
	// to uniquely identity the transfer in the internal
	// implementation of the service.
	OpaqueId string `protobuf:"bytes,1,opt,name=opaque_id,json=opaqueId,proto3" json:"opaque_id,omitempty"`
}

func (x *TxId) Reset() {
	*x = TxId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cs3_tx_v1beta1_resources_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxId) ProtoMessage() {}

func (x *TxId) ProtoReflect() protoreflect.Message {
	mi := &file_cs3_tx_v1beta1_resources_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxId.ProtoReflect.Descriptor instead.
func (*TxId) Descriptor() ([]byte, []int) {
	return file_cs3_tx_v1beta1_resources_proto_rawDescGZIP(), []int{0}
}

func (x *TxId) GetOpaqueId() string {
	if x != nil {
		return x.OpaqueId
	}
	return ""
}

// TxInfo represents information about a transfer.
type TxInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// REQUIRED.
	// The transfer identifier.
	Id *TxId `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// REQUIRED.
	// The transfer status. Eg.: STATUS_TRANSFER_FAILED.
	// Note: the description field may provide additional information on the transfer status.
	Status Status `protobuf:"varint,2,opt,name=status,proto3,enum=cs3.tx.v1beta1.Status" json:"status,omitempty"`
	// REQUIRED.
	// The destination (receiver of the transfer)
	Grantee *v1beta1.Grantee `protobuf:"bytes,3,opt,name=grantee,proto3" json:"grantee,omitempty"`
	// REQUIRED.
	// Uniquely identifies a principal who initiates the transfer creation.
	Creator *v1beta11.UserId `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty"`
	// REQUIRED.
	// Creation time of the transfer.
	Ctime *v1beta12.Timestamp `protobuf:"bytes,5,opt,name=ctime,proto3" json:"ctime,omitempty"`
	// OPTIONAL.
	// Information to describe the transfer status.
	// Eg. may contain information about a transfer failure.
	// Meant to be human-readable.
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	// REQUIRED.
	// Opaque unique identifier of the share on which the transfer is based.
	ShareId *v1beta13.ShareId `protobuf:"bytes,7,opt,name=share_id,json=shareId,proto3" json:"share_id,omitempty"`
}

func (x *TxInfo) Reset() {
	*x = TxInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cs3_tx_v1beta1_resources_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxInfo) ProtoMessage() {}

func (x *TxInfo) ProtoReflect() protoreflect.Message {
	mi := &file_cs3_tx_v1beta1_resources_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxInfo.ProtoReflect.Descriptor instead.
func (*TxInfo) Descriptor() ([]byte, []int) {
	return file_cs3_tx_v1beta1_resources_proto_rawDescGZIP(), []int{1}
}

func (x *TxInfo) GetId() *TxId {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *TxInfo) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_STATUS_INVALID
}

func (x *TxInfo) GetGrantee() *v1beta1.Grantee {
	if x != nil {
		return x.Grantee
	}
	return nil
}

func (x *TxInfo) GetCreator() *v1beta11.UserId {
	if x != nil {
		return x.Creator
	}
	return nil
}

func (x *TxInfo) GetCtime() *v1beta12.Timestamp {
	if x != nil {
		return x.Ctime
	}
	return nil
}

func (x *TxInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TxInfo) GetShareId() *v1beta13.ShareId {
	if x != nil {
		return x.ShareId
	}
	return nil
}

var File_cs3_tx_v1beta1_resources_proto protoreflect.FileDescriptor

var file_cs3_tx_v1beta1_resources_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x73, 0x33, 0x2f, 0x74, 0x78, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0e, 0x63, 0x73, 0x33, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x1a, 0x29, 0x63, 0x73, 0x33, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x63, 0x73, 0x33,
	0x2f, 0x73, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x6f, 0x63, 0x6d, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x63, 0x73, 0x33, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x63, 0x73, 0x33, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x23, 0x0a, 0x04, 0x54, 0x78, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x70, 0x61,
	0x71, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x70,
	0x61, 0x71, 0x75, 0x65, 0x49, 0x64, 0x22, 0xef, 0x02, 0x0a, 0x06, 0x54, 0x78, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x24, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x63, 0x73, 0x33, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54,
	0x78, 0x49, 0x64, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x74, 0x78,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3f, 0x0a, 0x07, 0x67, 0x72, 0x61, 0x6e, 0x74,
	0x65, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x65, 0x52,
	0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x65, 0x12, 0x3b, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x73, 0x33, 0x2e,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x32, 0x0a, 0x05, 0x63, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x05, 0x63, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x08, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x63, 0x73, 0x33, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x6f, 0x63, 0x6d, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x49, 0x64, 0x52,
	0x07, 0x73, 0x68, 0x61, 0x72, 0x65, 0x49, 0x64, 0x2a, 0xd8, 0x02, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x20, 0x0a, 0x1c, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x44, 0x45, 0x53, 0x54, 0x49, 0x4e, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f,
	0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x4e, 0x45, 0x57,
	0x10, 0x02, 0x12, 0x27, 0x0a, 0x23, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x54, 0x52, 0x41,
	0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x41, 0x57, 0x41, 0x49, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x41,
	0x43, 0x43, 0x45, 0x50, 0x54, 0x41, 0x4e, 0x43, 0x45, 0x10, 0x03, 0x12, 0x1c, 0x0a, 0x18, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x41,
	0x43, 0x43, 0x45, 0x50, 0x54, 0x45, 0x44, 0x10, 0x04, 0x12, 0x1f, 0x0a, 0x1b, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x49, 0x4e, 0x5f,
	0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x05, 0x12, 0x1c, 0x0a, 0x18, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x43, 0x4f,
	0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x06, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x46, 0x41, 0x49, 0x4c,
	0x45, 0x44, 0x10, 0x07, 0x12, 0x1d, 0x0a, 0x19, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x54,
	0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45,
	0x44, 0x10, 0x08, 0x12, 0x21, 0x0a, 0x1d, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x54, 0x52,
	0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x5f, 0x46, 0x41,
	0x49, 0x4c, 0x45, 0x44, 0x10, 0x09, 0x12, 0x1b, 0x0a, 0x17, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45,
	0x44, 0x10, 0x0a, 0x42, 0xb5, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x73, 0x33, 0x2e,
	0x74, 0x78, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x73, 0x33, 0x6f, 0x72, 0x67, 0x2f,
	0x67, 0x6f, 0x2d, 0x63, 0x73, 0x33, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x73, 0x33, 0x2f, 0x74,
	0x78, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x74, 0x78, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x54, 0x58, 0xaa, 0x02, 0x0e, 0x43, 0x73, 0x33, 0x2e,
	0x54, 0x78, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x0e, 0x43, 0x73, 0x33,
	0x5c, 0x54, 0x78, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x1a, 0x43, 0x73,
	0x33, 0x5c, 0x54, 0x78, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x43, 0x73, 0x33, 0x3a, 0x3a,
	0x54, 0x78, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_cs3_tx_v1beta1_resources_proto_rawDescOnce sync.Once
	file_cs3_tx_v1beta1_resources_proto_rawDescData = file_cs3_tx_v1beta1_resources_proto_rawDesc
)

func file_cs3_tx_v1beta1_resources_proto_rawDescGZIP() []byte {
	file_cs3_tx_v1beta1_resources_proto_rawDescOnce.Do(func() {
		file_cs3_tx_v1beta1_resources_proto_rawDescData = protoimpl.X.CompressGZIP(file_cs3_tx_v1beta1_resources_proto_rawDescData)
	})
	return file_cs3_tx_v1beta1_resources_proto_rawDescData
}

var file_cs3_tx_v1beta1_resources_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cs3_tx_v1beta1_resources_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cs3_tx_v1beta1_resources_proto_goTypes = []interface{}{
	(Status)(0),                // 0: cs3.tx.v1beta1.Status
	(*TxId)(nil),               // 1: cs3.tx.v1beta1.TxId
	(*TxInfo)(nil),             // 2: cs3.tx.v1beta1.TxInfo
	(*v1beta1.Grantee)(nil),    // 3: cs3.storage.provider.v1beta1.Grantee
	(*v1beta11.UserId)(nil),    // 4: cs3.identity.user.v1beta1.UserId
	(*v1beta12.Timestamp)(nil), // 5: cs3.types.v1beta1.Timestamp
	(*v1beta13.ShareId)(nil),   // 6: cs3.sharing.ocm.v1beta1.ShareId
}
var file_cs3_tx_v1beta1_resources_proto_depIdxs = []int32{
	1, // 0: cs3.tx.v1beta1.TxInfo.id:type_name -> cs3.tx.v1beta1.TxId
	0, // 1: cs3.tx.v1beta1.TxInfo.status:type_name -> cs3.tx.v1beta1.Status
	3, // 2: cs3.tx.v1beta1.TxInfo.grantee:type_name -> cs3.storage.provider.v1beta1.Grantee
	4, // 3: cs3.tx.v1beta1.TxInfo.creator:type_name -> cs3.identity.user.v1beta1.UserId
	5, // 4: cs3.tx.v1beta1.TxInfo.ctime:type_name -> cs3.types.v1beta1.Timestamp
	6, // 5: cs3.tx.v1beta1.TxInfo.share_id:type_name -> cs3.sharing.ocm.v1beta1.ShareId
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_cs3_tx_v1beta1_resources_proto_init() }
func file_cs3_tx_v1beta1_resources_proto_init() {
	if File_cs3_tx_v1beta1_resources_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cs3_tx_v1beta1_resources_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxId); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cs3_tx_v1beta1_resources_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cs3_tx_v1beta1_resources_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cs3_tx_v1beta1_resources_proto_goTypes,
		DependencyIndexes: file_cs3_tx_v1beta1_resources_proto_depIdxs,
		EnumInfos:         file_cs3_tx_v1beta1_resources_proto_enumTypes,
		MessageInfos:      file_cs3_tx_v1beta1_resources_proto_msgTypes,
	}.Build()
	File_cs3_tx_v1beta1_resources_proto = out.File
	file_cs3_tx_v1beta1_resources_proto_rawDesc = nil
	file_cs3_tx_v1beta1_resources_proto_goTypes = nil
	file_cs3_tx_v1beta1_resources_proto_depIdxs = nil
}
