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
// source: cs3/admin/group/v1beta1/group_api.proto

package groupv1beta1

import (
	v1beta11 "github.com/cs3org/go-cs3apis/cs3/identity/group/v1beta1"
	v1beta13 "github.com/cs3org/go-cs3apis/cs3/identity/user/v1beta1"
	v1beta12 "github.com/cs3org/go-cs3apis/cs3/rpc/v1beta1"
	v1beta1 "github.com/cs3org/go-cs3apis/cs3/types/v1beta1"
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

type CreateGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// OPTIONAL.
	// Opaque information. Allow to send any arbitrary data a service might use that is outside the API boundaries
	Opaque *v1beta1.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The information of group to be created.
	Group *v1beta11.Group `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *CreateGroupRequest) Reset() {
	*x = CreateGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cs3_admin_group_v1beta1_group_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroupRequest) ProtoMessage() {}

func (x *CreateGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cs3_admin_group_v1beta1_group_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroupRequest.ProtoReflect.Descriptor instead.
func (*CreateGroupRequest) Descriptor() ([]byte, []int) {
	return file_cs3_admin_group_v1beta1_group_api_proto_rawDescGZIP(), []int{0}
}

func (x *CreateGroupRequest) GetOpaque() *v1beta1.Opaque {
	if x != nil {
		return x.Opaque
	}
	return nil
}

func (x *CreateGroupRequest) GetGroup() *v1beta11.Group {
	if x != nil {
		return x.Group
	}
	return nil
}

type CreateGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// REQUIRED.
	// The response status.
	Status *v1beta12.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The group information.
	Group *v1beta11.Group `protobuf:"bytes,3,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *CreateGroupResponse) Reset() {
	*x = CreateGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cs3_admin_group_v1beta1_group_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroupResponse) ProtoMessage() {}

func (x *CreateGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cs3_admin_group_v1beta1_group_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroupResponse.ProtoReflect.Descriptor instead.
func (*CreateGroupResponse) Descriptor() ([]byte, []int) {
	return file_cs3_admin_group_v1beta1_group_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateGroupResponse) GetStatus() *v1beta12.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *CreateGroupResponse) GetOpaque() *v1beta1.Opaque {
	if x != nil {
		return x.Opaque
	}
	return nil
}

func (x *CreateGroupResponse) GetGroup() *v1beta11.Group {
	if x != nil {
		return x.Group
	}
	return nil
}

type DeleteGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// OPTIONAL.
	// Opaque information. Allow to send any arbitrary data a service might use that is outside the API boundaries.
	Opaque *v1beta1.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The group to be deleted, given their ID.
	GroupId *v1beta11.GroupId `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *DeleteGroupRequest) Reset() {
	*x = DeleteGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cs3_admin_group_v1beta1_group_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGroupRequest) ProtoMessage() {}

func (x *DeleteGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cs3_admin_group_v1beta1_group_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGroupRequest.ProtoReflect.Descriptor instead.
func (*DeleteGroupRequest) Descriptor() ([]byte, []int) {
	return file_cs3_admin_group_v1beta1_group_api_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteGroupRequest) GetOpaque() *v1beta1.Opaque {
	if x != nil {
		return x.Opaque
	}
	return nil
}

func (x *DeleteGroupRequest) GetGroupId() *v1beta11.GroupId {
	if x != nil {
		return x.GroupId
	}
	return nil
}

type DeleteGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// REQUIRED.
	// The response status.
	Status *v1beta12.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"`
}

func (x *DeleteGroupResponse) Reset() {
	*x = DeleteGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cs3_admin_group_v1beta1_group_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGroupResponse) ProtoMessage() {}

func (x *DeleteGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cs3_admin_group_v1beta1_group_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGroupResponse.ProtoReflect.Descriptor instead.
func (*DeleteGroupResponse) Descriptor() ([]byte, []int) {
	return file_cs3_admin_group_v1beta1_group_api_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteGroupResponse) GetStatus() *v1beta12.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *DeleteGroupResponse) GetOpaque() *v1beta1.Opaque {
	if x != nil {
		return x.Opaque
	}
	return nil
}

type AddUserToGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// REQUIRED.
	// ID of the user to add to the group
	UserId *v1beta13.UserId `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// REQUIRED.
	// ID of the target group.
	GroupId *v1beta11.GroupId `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,3,opt,name=opaque,proto3" json:"opaque,omitempty"`
}

func (x *AddUserToGroupRequest) Reset() {
	*x = AddUserToGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cs3_admin_group_v1beta1_group_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserToGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserToGroupRequest) ProtoMessage() {}

func (x *AddUserToGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cs3_admin_group_v1beta1_group_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserToGroupRequest.ProtoReflect.Descriptor instead.
func (*AddUserToGroupRequest) Descriptor() ([]byte, []int) {
	return file_cs3_admin_group_v1beta1_group_api_proto_rawDescGZIP(), []int{4}
}

func (x *AddUserToGroupRequest) GetUserId() *v1beta13.UserId {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *AddUserToGroupRequest) GetGroupId() *v1beta11.GroupId {
	if x != nil {
		return x.GroupId
	}
	return nil
}

func (x *AddUserToGroupRequest) GetOpaque() *v1beta1.Opaque {
	if x != nil {
		return x.Opaque
	}
	return nil
}

type AddUserToGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// REQUIRED.
	// The response status.
	Status *v1beta12.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"`
}

func (x *AddUserToGroupResponse) Reset() {
	*x = AddUserToGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cs3_admin_group_v1beta1_group_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserToGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserToGroupResponse) ProtoMessage() {}

func (x *AddUserToGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cs3_admin_group_v1beta1_group_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserToGroupResponse.ProtoReflect.Descriptor instead.
func (*AddUserToGroupResponse) Descriptor() ([]byte, []int) {
	return file_cs3_admin_group_v1beta1_group_api_proto_rawDescGZIP(), []int{5}
}

func (x *AddUserToGroupResponse) GetStatus() *v1beta12.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *AddUserToGroupResponse) GetOpaque() *v1beta1.Opaque {
	if x != nil {
		return x.Opaque
	}
	return nil
}

type RemoveUserFromGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// REQUIRED.
	// ID of the user to add to the group
	UserId *v1beta13.UserId `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// REQUIRED.
	// ID of the target group.
	GroupId *v1beta11.GroupId `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,3,opt,name=opaque,proto3" json:"opaque,omitempty"`
}

func (x *RemoveUserFromGroupRequest) Reset() {
	*x = RemoveUserFromGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cs3_admin_group_v1beta1_group_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveUserFromGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveUserFromGroupRequest) ProtoMessage() {}

func (x *RemoveUserFromGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cs3_admin_group_v1beta1_group_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveUserFromGroupRequest.ProtoReflect.Descriptor instead.
func (*RemoveUserFromGroupRequest) Descriptor() ([]byte, []int) {
	return file_cs3_admin_group_v1beta1_group_api_proto_rawDescGZIP(), []int{6}
}

func (x *RemoveUserFromGroupRequest) GetUserId() *v1beta13.UserId {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *RemoveUserFromGroupRequest) GetGroupId() *v1beta11.GroupId {
	if x != nil {
		return x.GroupId
	}
	return nil
}

func (x *RemoveUserFromGroupRequest) GetOpaque() *v1beta1.Opaque {
	if x != nil {
		return x.Opaque
	}
	return nil
}

type RemoveUserFromGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// REQUIRED.
	// The response status.
	Status *v1beta12.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"`
}

func (x *RemoveUserFromGroupResponse) Reset() {
	*x = RemoveUserFromGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cs3_admin_group_v1beta1_group_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveUserFromGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveUserFromGroupResponse) ProtoMessage() {}

func (x *RemoveUserFromGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cs3_admin_group_v1beta1_group_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveUserFromGroupResponse.ProtoReflect.Descriptor instead.
func (*RemoveUserFromGroupResponse) Descriptor() ([]byte, []int) {
	return file_cs3_admin_group_v1beta1_group_api_proto_rawDescGZIP(), []int{7}
}

func (x *RemoveUserFromGroupResponse) GetStatus() *v1beta12.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *RemoveUserFromGroupResponse) GetOpaque() *v1beta1.Opaque {
	if x != nil {
		return x.Opaque
	}
	return nil
}

var File_cs3_admin_group_v1beta1_group_api_proto protoreflect.FileDescriptor

var file_cs3_admin_group_v1beta1_group_api_proto_rawDesc = []byte{
	0x0a, 0x27, 0x63, 0x73, 0x33, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x63, 0x73, 0x33, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x1a, 0x2a, 0x63, 0x73, 0x33, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29,
	0x63, 0x73, 0x33, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x63, 0x73, 0x33, 0x2f, 0x72,
	0x70, 0x63, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x63, 0x73, 0x33, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a,
	0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x63, 0x73, 0x33, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x4f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x52, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65,
	0x12, 0x37, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22, 0xb2, 0x01, 0x0a, 0x13, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x31, 0x0a, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x52, 0x06, 0x6f,
	0x70, 0x61, 0x71, 0x75, 0x65, 0x12, 0x37, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x87,
	0x01, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4f, 0x70, 0x61, 0x71, 0x75, 0x65,
	0x52, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x12, 0x3e, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x73, 0x33,
	0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x52,
	0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x79, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x31, 0x0a, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x4f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x52, 0x06, 0x6f, 0x70, 0x61,
	0x71, 0x75, 0x65, 0x22, 0xc6, 0x01, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x54,
	0x6f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x63, 0x73, 0x33, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3e, 0x0a, 0x08, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x73,
	0x33, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64,
	0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x06, 0x6f, 0x70, 0x61,
	0x71, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x73, 0x33, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4f, 0x70,
	0x61, 0x71, 0x75, 0x65, 0x52, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x22, 0x7c, 0x0a, 0x16,
	0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x31, 0x0a, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4f, 0x70, 0x61, 0x71,
	0x75, 0x65, 0x52, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x22, 0xcb, 0x01, 0x0a, 0x1a, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x73, 0x33,
	0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3e, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x52, 0x07, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4f, 0x70, 0x61, 0x71, 0x75, 0x65,
	0x52, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x22, 0x81, 0x01, 0x0a, 0x1b, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x31, 0x0a, 0x06, 0x6f, 0x70, 0x61,
	0x71, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x73, 0x33, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4f, 0x70,
	0x61, 0x71, 0x75, 0x65, 0x52, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x32, 0xd4, 0x03, 0x0a,
	0x08, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x50, 0x49, 0x12, 0x68, 0x0a, 0x0b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x2b, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x68, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x2b, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2c, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x71, 0x0a,
	0x0e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12,
	0x2e, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65,
	0x72, 0x54, 0x6f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2f, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65,
	0x72, 0x54, 0x6f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x80, 0x01, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x46,
	0x72, 0x6f, 0x6d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x33, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x46, 0x72, 0x6f,
	0x6d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e,
	0x63, 0x73, 0x33, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0xee, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x73, 0x33, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x42, 0x0d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x70, 0x69, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x73, 0x33, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x73, 0x33, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x63, 0x73, 0x33, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x41, 0x47, 0xaa, 0x02, 0x17,
	0x43, 0x73, 0x33, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x2e,
	0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x17, 0x43, 0x73, 0x33, 0x5c, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x5c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0xe2, 0x02, 0x23, 0x43, 0x73, 0x33, 0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x5c, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1a, 0x43, 0x73, 0x33, 0x3a, 0x3a, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x3a, 0x3a, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cs3_admin_group_v1beta1_group_api_proto_rawDescOnce sync.Once
	file_cs3_admin_group_v1beta1_group_api_proto_rawDescData = file_cs3_admin_group_v1beta1_group_api_proto_rawDesc
)

func file_cs3_admin_group_v1beta1_group_api_proto_rawDescGZIP() []byte {
	file_cs3_admin_group_v1beta1_group_api_proto_rawDescOnce.Do(func() {
		file_cs3_admin_group_v1beta1_group_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_cs3_admin_group_v1beta1_group_api_proto_rawDescData)
	})
	return file_cs3_admin_group_v1beta1_group_api_proto_rawDescData
}

var file_cs3_admin_group_v1beta1_group_api_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_cs3_admin_group_v1beta1_group_api_proto_goTypes = []interface{}{
	(*CreateGroupRequest)(nil),          // 0: cs3.admin.group.v1beta1.CreateGroupRequest
	(*CreateGroupResponse)(nil),         // 1: cs3.admin.group.v1beta1.CreateGroupResponse
	(*DeleteGroupRequest)(nil),          // 2: cs3.admin.group.v1beta1.DeleteGroupRequest
	(*DeleteGroupResponse)(nil),         // 3: cs3.admin.group.v1beta1.DeleteGroupResponse
	(*AddUserToGroupRequest)(nil),       // 4: cs3.admin.group.v1beta1.AddUserToGroupRequest
	(*AddUserToGroupResponse)(nil),      // 5: cs3.admin.group.v1beta1.AddUserToGroupResponse
	(*RemoveUserFromGroupRequest)(nil),  // 6: cs3.admin.group.v1beta1.RemoveUserFromGroupRequest
	(*RemoveUserFromGroupResponse)(nil), // 7: cs3.admin.group.v1beta1.RemoveUserFromGroupResponse
	(*v1beta1.Opaque)(nil),              // 8: cs3.types.v1beta1.Opaque
	(*v1beta11.Group)(nil),              // 9: cs3.identity.group.v1beta1.Group
	(*v1beta12.Status)(nil),             // 10: cs3.rpc.v1beta1.Status
	(*v1beta11.GroupId)(nil),            // 11: cs3.identity.group.v1beta1.GroupId
	(*v1beta13.UserId)(nil),             // 12: cs3.identity.user.v1beta1.UserId
}
var file_cs3_admin_group_v1beta1_group_api_proto_depIdxs = []int32{
	8,  // 0: cs3.admin.group.v1beta1.CreateGroupRequest.opaque:type_name -> cs3.types.v1beta1.Opaque
	9,  // 1: cs3.admin.group.v1beta1.CreateGroupRequest.group:type_name -> cs3.identity.group.v1beta1.Group
	10, // 2: cs3.admin.group.v1beta1.CreateGroupResponse.status:type_name -> cs3.rpc.v1beta1.Status
	8,  // 3: cs3.admin.group.v1beta1.CreateGroupResponse.opaque:type_name -> cs3.types.v1beta1.Opaque
	9,  // 4: cs3.admin.group.v1beta1.CreateGroupResponse.group:type_name -> cs3.identity.group.v1beta1.Group
	8,  // 5: cs3.admin.group.v1beta1.DeleteGroupRequest.opaque:type_name -> cs3.types.v1beta1.Opaque
	11, // 6: cs3.admin.group.v1beta1.DeleteGroupRequest.group_id:type_name -> cs3.identity.group.v1beta1.GroupId
	10, // 7: cs3.admin.group.v1beta1.DeleteGroupResponse.status:type_name -> cs3.rpc.v1beta1.Status
	8,  // 8: cs3.admin.group.v1beta1.DeleteGroupResponse.opaque:type_name -> cs3.types.v1beta1.Opaque
	12, // 9: cs3.admin.group.v1beta1.AddUserToGroupRequest.user_id:type_name -> cs3.identity.user.v1beta1.UserId
	11, // 10: cs3.admin.group.v1beta1.AddUserToGroupRequest.group_id:type_name -> cs3.identity.group.v1beta1.GroupId
	8,  // 11: cs3.admin.group.v1beta1.AddUserToGroupRequest.opaque:type_name -> cs3.types.v1beta1.Opaque
	10, // 12: cs3.admin.group.v1beta1.AddUserToGroupResponse.status:type_name -> cs3.rpc.v1beta1.Status
	8,  // 13: cs3.admin.group.v1beta1.AddUserToGroupResponse.opaque:type_name -> cs3.types.v1beta1.Opaque
	12, // 14: cs3.admin.group.v1beta1.RemoveUserFromGroupRequest.user_id:type_name -> cs3.identity.user.v1beta1.UserId
	11, // 15: cs3.admin.group.v1beta1.RemoveUserFromGroupRequest.group_id:type_name -> cs3.identity.group.v1beta1.GroupId
	8,  // 16: cs3.admin.group.v1beta1.RemoveUserFromGroupRequest.opaque:type_name -> cs3.types.v1beta1.Opaque
	10, // 17: cs3.admin.group.v1beta1.RemoveUserFromGroupResponse.status:type_name -> cs3.rpc.v1beta1.Status
	8,  // 18: cs3.admin.group.v1beta1.RemoveUserFromGroupResponse.opaque:type_name -> cs3.types.v1beta1.Opaque
	0,  // 19: cs3.admin.group.v1beta1.GroupAPI.CreateGroup:input_type -> cs3.admin.group.v1beta1.CreateGroupRequest
	2,  // 20: cs3.admin.group.v1beta1.GroupAPI.DeleteGroup:input_type -> cs3.admin.group.v1beta1.DeleteGroupRequest
	4,  // 21: cs3.admin.group.v1beta1.GroupAPI.AddUserToGroup:input_type -> cs3.admin.group.v1beta1.AddUserToGroupRequest
	6,  // 22: cs3.admin.group.v1beta1.GroupAPI.RemoveUserFromGroup:input_type -> cs3.admin.group.v1beta1.RemoveUserFromGroupRequest
	1,  // 23: cs3.admin.group.v1beta1.GroupAPI.CreateGroup:output_type -> cs3.admin.group.v1beta1.CreateGroupResponse
	3,  // 24: cs3.admin.group.v1beta1.GroupAPI.DeleteGroup:output_type -> cs3.admin.group.v1beta1.DeleteGroupResponse
	5,  // 25: cs3.admin.group.v1beta1.GroupAPI.AddUserToGroup:output_type -> cs3.admin.group.v1beta1.AddUserToGroupResponse
	7,  // 26: cs3.admin.group.v1beta1.GroupAPI.RemoveUserFromGroup:output_type -> cs3.admin.group.v1beta1.RemoveUserFromGroupResponse
	23, // [23:27] is the sub-list for method output_type
	19, // [19:23] is the sub-list for method input_type
	19, // [19:19] is the sub-list for extension type_name
	19, // [19:19] is the sub-list for extension extendee
	0,  // [0:19] is the sub-list for field type_name
}

func init() { file_cs3_admin_group_v1beta1_group_api_proto_init() }
func file_cs3_admin_group_v1beta1_group_api_proto_init() {
	if File_cs3_admin_group_v1beta1_group_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cs3_admin_group_v1beta1_group_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGroupRequest); i {
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
		file_cs3_admin_group_v1beta1_group_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGroupResponse); i {
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
		file_cs3_admin_group_v1beta1_group_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGroupRequest); i {
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
		file_cs3_admin_group_v1beta1_group_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGroupResponse); i {
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
		file_cs3_admin_group_v1beta1_group_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserToGroupRequest); i {
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
		file_cs3_admin_group_v1beta1_group_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserToGroupResponse); i {
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
		file_cs3_admin_group_v1beta1_group_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveUserFromGroupRequest); i {
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
		file_cs3_admin_group_v1beta1_group_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveUserFromGroupResponse); i {
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
			RawDescriptor: file_cs3_admin_group_v1beta1_group_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cs3_admin_group_v1beta1_group_api_proto_goTypes,
		DependencyIndexes: file_cs3_admin_group_v1beta1_group_api_proto_depIdxs,
		MessageInfos:      file_cs3_admin_group_v1beta1_group_api_proto_msgTypes,
	}.Build()
	File_cs3_admin_group_v1beta1_group_api_proto = out.File
	file_cs3_admin_group_v1beta1_group_api_proto_rawDesc = nil
	file_cs3_admin_group_v1beta1_group_api_proto_goTypes = nil
	file_cs3_admin_group_v1beta1_group_api_proto_depIdxs = nil
}
