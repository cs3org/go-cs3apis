// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs3/storagebroker/v0alpha/resources.proto

package storagebrokerv0alphapb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// TODO(labkode): maybe add capabilities field
type ProviderInfo struct {
	Location             string   `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProviderInfo) Reset()         { *m = ProviderInfo{} }
func (m *ProviderInfo) String() string { return proto.CompactTextString(m) }
func (*ProviderInfo) ProtoMessage()    {}
func (*ProviderInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_09d00d8e2bdc177e, []int{0}
}

func (m *ProviderInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProviderInfo.Unmarshal(m, b)
}
func (m *ProviderInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProviderInfo.Marshal(b, m, deterministic)
}
func (m *ProviderInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProviderInfo.Merge(m, src)
}
func (m *ProviderInfo) XXX_Size() int {
	return xxx_messageInfo_ProviderInfo.Size(m)
}
func (m *ProviderInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ProviderInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ProviderInfo proto.InternalMessageInfo

func (m *ProviderInfo) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func init() {
	proto.RegisterType((*ProviderInfo)(nil), "cs3.storagebrokerv0alpha.ProviderInfo")
}

func init() {
	proto.RegisterFile("cs3/storagebroker/v0alpha/resources.proto", fileDescriptor_09d00d8e2bdc177e)
}

var fileDescriptor_09d00d8e2bdc177e = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4c, 0x2e, 0x36, 0xd6,
	0x2f, 0x2e, 0xc9, 0x2f, 0x4a, 0x4c, 0x4f, 0x4d, 0x2a, 0xca, 0xcf, 0x4e, 0x2d, 0xd2, 0x2f, 0x33,
	0x48, 0xcc, 0x29, 0xc8, 0x48, 0xd4, 0x2f, 0x4a, 0x2d, 0xce, 0x2f, 0x2d, 0x4a, 0x4e, 0x2d, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x48, 0x2e, 0x36, 0xd6, 0x43, 0x51, 0x0a, 0x55, 0xa9,
	0xa4, 0xc5, 0xc5, 0x13, 0x50, 0x94, 0x5f, 0x96, 0x99, 0x92, 0x5a, 0xe4, 0x99, 0x97, 0x96, 0x2f,
	0x24, 0xc5, 0xc5, 0x91, 0x93, 0x9f, 0x9c, 0x58, 0x92, 0x99, 0x9f, 0x27, 0xc1, 0xa8, 0xc0, 0xa8,
	0xc1, 0x19, 0x04, 0xe7, 0x3b, 0xb5, 0x33, 0x72, 0xc9, 0x24, 0xe7, 0xe7, 0xea, 0xe1, 0x32, 0xcc,
	0x89, 0x2f, 0x08, 0x66, 0x6f, 0x00, 0xc8, 0xda, 0x00, 0xc6, 0x28, 0x31, 0x6c, 0xea, 0x0a, 0x92,
	0x16, 0x31, 0xb1, 0x39, 0x3b, 0xf9, 0x47, 0x38, 0x3a, 0xad, 0x62, 0x92, 0x70, 0x0e, 0x36, 0xd6,
	0x0b, 0x86, 0xa8, 0x72, 0x02, 0xab, 0x0a, 0x33, 0x70, 0x04, 0xa9, 0x3a, 0x05, 0x96, 0x8a, 0xc1,
	0x26, 0x95, 0xc4, 0x06, 0xf6, 0x96, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x99, 0x3a, 0x98, 0x99,
	0x03, 0x01, 0x00, 0x00,
}