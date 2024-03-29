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
// source: cs3/storage/registry/v1beta1/registry_api.proto

package registryv1beta1

import (
	v1beta11 "github.com/cs3org/go-cs3apis/cs3/rpc/v1beta1"
	v1beta12 "github.com/cs3org/go-cs3apis/cs3/storage/provider/v1beta1"
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

type GetHomeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
}

func (x *GetHomeRequest) Reset() {
	*x = GetHomeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cs3_storage_registry_v1beta1_registry_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHomeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHomeRequest) ProtoMessage() {}

func (x *GetHomeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cs3_storage_registry_v1beta1_registry_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHomeRequest.ProtoReflect.Descriptor instead.
func (*GetHomeRequest) Descriptor() ([]byte, []int) {
	return file_cs3_storage_registry_v1beta1_registry_api_proto_rawDescGZIP(), []int{0}
}

func (x *GetHomeRequest) GetOpaque() *v1beta1.Opaque {
	if x != nil {
		return x.Opaque
	}
	return nil
}

type GetHomeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// REQUIRED.
	// The response status.
	Status *v1beta11.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The storage provider for the user home.
	Provider *ProviderInfo `protobuf:"bytes,3,opt,name=provider,proto3" json:"provider,omitempty"`
}

func (x *GetHomeResponse) Reset() {
	*x = GetHomeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cs3_storage_registry_v1beta1_registry_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHomeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHomeResponse) ProtoMessage() {}

func (x *GetHomeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cs3_storage_registry_v1beta1_registry_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHomeResponse.ProtoReflect.Descriptor instead.
func (*GetHomeResponse) Descriptor() ([]byte, []int) {
	return file_cs3_storage_registry_v1beta1_registry_api_proto_rawDescGZIP(), []int{1}
}

func (x *GetHomeResponse) GetStatus() *v1beta11.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *GetHomeResponse) GetOpaque() *v1beta1.Opaque {
	if x != nil {
		return x.Opaque
	}
	return nil
}

func (x *GetHomeResponse) GetProvider() *ProviderInfo {
	if x != nil {
		return x.Provider
	}
	return nil
}

type GetStorageProvidersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The reference for the resource.
	Ref *v1beta12.Reference `protobuf:"bytes,2,opt,name=ref,proto3" json:"ref,omitempty"`
}

func (x *GetStorageProvidersRequest) Reset() {
	*x = GetStorageProvidersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cs3_storage_registry_v1beta1_registry_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStorageProvidersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStorageProvidersRequest) ProtoMessage() {}

func (x *GetStorageProvidersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cs3_storage_registry_v1beta1_registry_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStorageProvidersRequest.ProtoReflect.Descriptor instead.
func (*GetStorageProvidersRequest) Descriptor() ([]byte, []int) {
	return file_cs3_storage_registry_v1beta1_registry_api_proto_rawDescGZIP(), []int{2}
}

func (x *GetStorageProvidersRequest) GetOpaque() *v1beta1.Opaque {
	if x != nil {
		return x.Opaque
	}
	return nil
}

func (x *GetStorageProvidersRequest) GetRef() *v1beta12.Reference {
	if x != nil {
		return x.Ref
	}
	return nil
}

type GetStorageProvidersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// REQUIRED.
	// The response status.
	Status *v1beta11.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The storage providers handling the requested storage resource.
	Providers []*ProviderInfo `protobuf:"bytes,3,rep,name=providers,proto3" json:"providers,omitempty"`
}

func (x *GetStorageProvidersResponse) Reset() {
	*x = GetStorageProvidersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cs3_storage_registry_v1beta1_registry_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStorageProvidersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStorageProvidersResponse) ProtoMessage() {}

func (x *GetStorageProvidersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cs3_storage_registry_v1beta1_registry_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStorageProvidersResponse.ProtoReflect.Descriptor instead.
func (*GetStorageProvidersResponse) Descriptor() ([]byte, []int) {
	return file_cs3_storage_registry_v1beta1_registry_api_proto_rawDescGZIP(), []int{3}
}

func (x *GetStorageProvidersResponse) GetStatus() *v1beta11.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *GetStorageProvidersResponse) GetOpaque() *v1beta1.Opaque {
	if x != nil {
		return x.Opaque
	}
	return nil
}

func (x *GetStorageProvidersResponse) GetProviders() []*ProviderInfo {
	if x != nil {
		return x.Providers
	}
	return nil
}

type ListStorageProvidersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"` // TODO(labkode): maybe add some filter?
}

func (x *ListStorageProvidersRequest) Reset() {
	*x = ListStorageProvidersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cs3_storage_registry_v1beta1_registry_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStorageProvidersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStorageProvidersRequest) ProtoMessage() {}

func (x *ListStorageProvidersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cs3_storage_registry_v1beta1_registry_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStorageProvidersRequest.ProtoReflect.Descriptor instead.
func (*ListStorageProvidersRequest) Descriptor() ([]byte, []int) {
	return file_cs3_storage_registry_v1beta1_registry_api_proto_rawDescGZIP(), []int{4}
}

func (x *ListStorageProvidersRequest) GetOpaque() *v1beta1.Opaque {
	if x != nil {
		return x.Opaque
	}
	return nil
}

type ListStorageProvidersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// REQUIRED.
	// The response status.
	Status *v1beta11.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The list of storage providers this registry knows about.
	Providers []*ProviderInfo `protobuf:"bytes,3,rep,name=providers,proto3" json:"providers,omitempty"`
}

func (x *ListStorageProvidersResponse) Reset() {
	*x = ListStorageProvidersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cs3_storage_registry_v1beta1_registry_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStorageProvidersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStorageProvidersResponse) ProtoMessage() {}

func (x *ListStorageProvidersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cs3_storage_registry_v1beta1_registry_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStorageProvidersResponse.ProtoReflect.Descriptor instead.
func (*ListStorageProvidersResponse) Descriptor() ([]byte, []int) {
	return file_cs3_storage_registry_v1beta1_registry_api_proto_rawDescGZIP(), []int{5}
}

func (x *ListStorageProvidersResponse) GetStatus() *v1beta11.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *ListStorageProvidersResponse) GetOpaque() *v1beta1.Opaque {
	if x != nil {
		return x.Opaque
	}
	return nil
}

func (x *ListStorageProvidersResponse) GetProviders() []*ProviderInfo {
	if x != nil {
		return x.Providers
	}
	return nil
}

var File_cs3_storage_registry_v1beta1_registry_api_proto protoreflect.FileDescriptor

var file_cs3_storage_registry_v1beta1_registry_api_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x63, 0x73, 0x33, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1c, 0x63, 0x73, 0x33, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a,
	0x1c, 0x63, 0x73, 0x33, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x63,
	0x73, 0x33, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x63, 0x73, 0x33,
	0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x63, 0x73, 0x33, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x48,
	0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x06, 0x6f, 0x70,
	0x61, 0x71, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x73, 0x33,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4f,
	0x70, 0x61, 0x71, 0x75, 0x65, 0x52, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x22, 0xbd, 0x01,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x31, 0x0a, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x52, 0x06, 0x6f,
	0x70, 0x61, 0x71, 0x75, 0x65, 0x12, 0x46, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22, 0x8a, 0x01,
	0x0a, 0x1a, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x06,
	0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63,
	0x73, 0x33, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x4f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x52, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x12,
	0x39, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63,
	0x73, 0x33, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x03, 0x72, 0x65, 0x66, 0x22, 0xcb, 0x01, 0x0a, 0x1b, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x73, 0x33,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x31, 0x0a, 0x06, 0x6f,
	0x70, 0x61, 0x71, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x73,
	0x33, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x4f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x52, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x12, 0x48,
	0x0a, 0x09, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x22, 0x50, 0x0a, 0x1b, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4f, 0x70, 0x61, 0x71,
	0x75, 0x65, 0x52, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x22, 0xcc, 0x01, 0x0a, 0x1c, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x73,
	0x33, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x31, 0x0a, 0x06,
	0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63,
	0x73, 0x33, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x4f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x52, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x12,
	0x48, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x32, 0x92, 0x03, 0x0a, 0x0b, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x41, 0x50, 0x49, 0x12, 0x8a, 0x01, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x73, 0x12, 0x38, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e, 0x63, 0x73,
	0x33, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x8d, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x12,
	0x39, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x63, 0x73, 0x33,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x66, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6d,
	0x65, 0x12, 0x2c, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2d, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x92,
	0x02, 0x0a, 0x20, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x42, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x41, 0x70, 0x69,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x73, 0x33, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x73,
	0x33, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x73, 0x33, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x3b, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x53, 0x52, 0xaa, 0x02, 0x1c, 0x43, 0x73, 0x33, 0x2e, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e,
	0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x1c, 0x43, 0x73, 0x33, 0x5c, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x5c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5c, 0x56,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x28, 0x43, 0x73, 0x33, 0x5c, 0x53, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x5c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5c, 0x56, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x1f, 0x43, 0x73, 0x33, 0x3a, 0x3a, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x3a, 0x3a, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cs3_storage_registry_v1beta1_registry_api_proto_rawDescOnce sync.Once
	file_cs3_storage_registry_v1beta1_registry_api_proto_rawDescData = file_cs3_storage_registry_v1beta1_registry_api_proto_rawDesc
)

func file_cs3_storage_registry_v1beta1_registry_api_proto_rawDescGZIP() []byte {
	file_cs3_storage_registry_v1beta1_registry_api_proto_rawDescOnce.Do(func() {
		file_cs3_storage_registry_v1beta1_registry_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_cs3_storage_registry_v1beta1_registry_api_proto_rawDescData)
	})
	return file_cs3_storage_registry_v1beta1_registry_api_proto_rawDescData
}

var file_cs3_storage_registry_v1beta1_registry_api_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_cs3_storage_registry_v1beta1_registry_api_proto_goTypes = []interface{}{
	(*GetHomeRequest)(nil),               // 0: cs3.storage.registry.v1beta1.GetHomeRequest
	(*GetHomeResponse)(nil),              // 1: cs3.storage.registry.v1beta1.GetHomeResponse
	(*GetStorageProvidersRequest)(nil),   // 2: cs3.storage.registry.v1beta1.GetStorageProvidersRequest
	(*GetStorageProvidersResponse)(nil),  // 3: cs3.storage.registry.v1beta1.GetStorageProvidersResponse
	(*ListStorageProvidersRequest)(nil),  // 4: cs3.storage.registry.v1beta1.ListStorageProvidersRequest
	(*ListStorageProvidersResponse)(nil), // 5: cs3.storage.registry.v1beta1.ListStorageProvidersResponse
	(*v1beta1.Opaque)(nil),               // 6: cs3.types.v1beta1.Opaque
	(*v1beta11.Status)(nil),              // 7: cs3.rpc.v1beta1.Status
	(*ProviderInfo)(nil),                 // 8: cs3.storage.registry.v1beta1.ProviderInfo
	(*v1beta12.Reference)(nil),           // 9: cs3.storage.provider.v1beta1.Reference
}
var file_cs3_storage_registry_v1beta1_registry_api_proto_depIdxs = []int32{
	6,  // 0: cs3.storage.registry.v1beta1.GetHomeRequest.opaque:type_name -> cs3.types.v1beta1.Opaque
	7,  // 1: cs3.storage.registry.v1beta1.GetHomeResponse.status:type_name -> cs3.rpc.v1beta1.Status
	6,  // 2: cs3.storage.registry.v1beta1.GetHomeResponse.opaque:type_name -> cs3.types.v1beta1.Opaque
	8,  // 3: cs3.storage.registry.v1beta1.GetHomeResponse.provider:type_name -> cs3.storage.registry.v1beta1.ProviderInfo
	6,  // 4: cs3.storage.registry.v1beta1.GetStorageProvidersRequest.opaque:type_name -> cs3.types.v1beta1.Opaque
	9,  // 5: cs3.storage.registry.v1beta1.GetStorageProvidersRequest.ref:type_name -> cs3.storage.provider.v1beta1.Reference
	7,  // 6: cs3.storage.registry.v1beta1.GetStorageProvidersResponse.status:type_name -> cs3.rpc.v1beta1.Status
	6,  // 7: cs3.storage.registry.v1beta1.GetStorageProvidersResponse.opaque:type_name -> cs3.types.v1beta1.Opaque
	8,  // 8: cs3.storage.registry.v1beta1.GetStorageProvidersResponse.providers:type_name -> cs3.storage.registry.v1beta1.ProviderInfo
	6,  // 9: cs3.storage.registry.v1beta1.ListStorageProvidersRequest.opaque:type_name -> cs3.types.v1beta1.Opaque
	7,  // 10: cs3.storage.registry.v1beta1.ListStorageProvidersResponse.status:type_name -> cs3.rpc.v1beta1.Status
	6,  // 11: cs3.storage.registry.v1beta1.ListStorageProvidersResponse.opaque:type_name -> cs3.types.v1beta1.Opaque
	8,  // 12: cs3.storage.registry.v1beta1.ListStorageProvidersResponse.providers:type_name -> cs3.storage.registry.v1beta1.ProviderInfo
	2,  // 13: cs3.storage.registry.v1beta1.RegistryAPI.GetStorageProviders:input_type -> cs3.storage.registry.v1beta1.GetStorageProvidersRequest
	4,  // 14: cs3.storage.registry.v1beta1.RegistryAPI.ListStorageProviders:input_type -> cs3.storage.registry.v1beta1.ListStorageProvidersRequest
	0,  // 15: cs3.storage.registry.v1beta1.RegistryAPI.GetHome:input_type -> cs3.storage.registry.v1beta1.GetHomeRequest
	3,  // 16: cs3.storage.registry.v1beta1.RegistryAPI.GetStorageProviders:output_type -> cs3.storage.registry.v1beta1.GetStorageProvidersResponse
	5,  // 17: cs3.storage.registry.v1beta1.RegistryAPI.ListStorageProviders:output_type -> cs3.storage.registry.v1beta1.ListStorageProvidersResponse
	1,  // 18: cs3.storage.registry.v1beta1.RegistryAPI.GetHome:output_type -> cs3.storage.registry.v1beta1.GetHomeResponse
	16, // [16:19] is the sub-list for method output_type
	13, // [13:16] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_cs3_storage_registry_v1beta1_registry_api_proto_init() }
func file_cs3_storage_registry_v1beta1_registry_api_proto_init() {
	if File_cs3_storage_registry_v1beta1_registry_api_proto != nil {
		return
	}
	file_cs3_storage_registry_v1beta1_resources_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cs3_storage_registry_v1beta1_registry_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHomeRequest); i {
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
		file_cs3_storage_registry_v1beta1_registry_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHomeResponse); i {
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
		file_cs3_storage_registry_v1beta1_registry_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStorageProvidersRequest); i {
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
		file_cs3_storage_registry_v1beta1_registry_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStorageProvidersResponse); i {
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
		file_cs3_storage_registry_v1beta1_registry_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStorageProvidersRequest); i {
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
		file_cs3_storage_registry_v1beta1_registry_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStorageProvidersResponse); i {
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
			RawDescriptor: file_cs3_storage_registry_v1beta1_registry_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cs3_storage_registry_v1beta1_registry_api_proto_goTypes,
		DependencyIndexes: file_cs3_storage_registry_v1beta1_registry_api_proto_depIdxs,
		MessageInfos:      file_cs3_storage_registry_v1beta1_registry_api_proto_msgTypes,
	}.Build()
	File_cs3_storage_registry_v1beta1_registry_api_proto = out.File
	file_cs3_storage_registry_v1beta1_registry_api_proto_rawDesc = nil
	file_cs3_storage_registry_v1beta1_registry_api_proto_goTypes = nil
	file_cs3_storage_registry_v1beta1_registry_api_proto_depIdxs = nil
}
