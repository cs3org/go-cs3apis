// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs3/shareregistry/v0alpha/resources.proto

package shareregistryv0alphapb

import (
	fmt "fmt"
	types "github.com/cernbox/go-cs3apis/cs3/types"
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

// The differentes type of share types.
type ShareType int32

const (
	ShareType_SHARE_TYPE_INVALID     ShareType = 0
	ShareType_SHARE_TYPE_USER        ShareType = 1
	ShareType_SHARE_TYPE_GROUP       ShareType = 2
	ShareType_SHARE_TYPE_PUBLIC_LINK ShareType = 3
	ShareType_SHARE_TYPE_OCM         ShareType = 4
)

var ShareType_name = map[int32]string{
	0: "SHARE_TYPE_INVALID",
	1: "SHARE_TYPE_USER",
	2: "SHARE_TYPE_GROUP",
	3: "SHARE_TYPE_PUBLIC_LINK",
	4: "SHARE_TYPE_OCM",
}

var ShareType_value = map[string]int32{
	"SHARE_TYPE_INVALID":     0,
	"SHARE_TYPE_USER":        1,
	"SHARE_TYPE_GROUP":       2,
	"SHARE_TYPE_PUBLIC_LINK": 3,
	"SHARE_TYPE_OCM":         4,
}

func (x ShareType) String() string {
	return proto.EnumName(ShareType_name, int32(x))
}

func (ShareType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a12331ba744ea1d9, []int{0}
}

// Represents the information of the share provider.
type ProviderInfo struct {
	// OPTIONAL.
	// Opaque information.
	Opaque *types.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The share type that will become part of the
	// share id.
	// For example, if the share_type is "user", shares obtained
	// from this storage provider will have a share id like "user:1234"
	// or "ocm:1234" or "group:8838".
	ShareType ShareType `protobuf:"varint,2,opt,name=share_type,json=shareType,proto3,enum=cs3.shareregistryv0alpha.ShareType" json:"share_type,omitempty"`
	// REQUIRED.
	// The address where the share provider can be reached.
	// For example, tcp://localhost:1099.
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	// OPTIONAL.
	// Information to describe the functionalities
	// offered by the share provider. Meant to be read
	// by humans.
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProviderInfo) Reset()         { *m = ProviderInfo{} }
func (m *ProviderInfo) String() string { return proto.CompactTextString(m) }
func (*ProviderInfo) ProtoMessage()    {}
func (*ProviderInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a12331ba744ea1d9, []int{0}
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

func (m *ProviderInfo) GetOpaque() *types.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *ProviderInfo) GetShareType() ShareType {
	if m != nil {
		return m.ShareType
	}
	return ShareType_SHARE_TYPE_INVALID
}

func (m *ProviderInfo) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ProviderInfo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterEnum("cs3.shareregistryv0alpha.ShareType", ShareType_name, ShareType_value)
	proto.RegisterType((*ProviderInfo)(nil), "cs3.shareregistryv0alpha.ProviderInfo")
}

func init() {
	proto.RegisterFile("cs3/shareregistry/v0alpha/resources.proto", fileDescriptor_a12331ba744ea1d9)
}

var fileDescriptor_a12331ba744ea1d9 = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xc1, 0x6e, 0xa2, 0x40,
	0x1c, 0xc6, 0x77, 0xd0, 0xb8, 0x71, 0xdc, 0xb8, 0xec, 0xb8, 0x6b, 0x88, 0xd9, 0x03, 0x69, 0x2f,
	0xda, 0x03, 0x1a, 0x79, 0x02, 0xa0, 0xc4, 0x92, 0x5a, 0x21, 0x03, 0x9a, 0xda, 0x34, 0x21, 0x08,
	0xd3, 0x4a, 0xd2, 0x3a, 0x74, 0x06, 0x4d, 0xbc, 0xf4, 0x39, 0x7a, 0xee, 0xb1, 0x6f, 0xd0, 0x57,
	0xe8, 0x53, 0x35, 0x80, 0x36, 0x34, 0xd1, 0x0b, 0x61, 0xbe, 0xdf, 0xf7, 0x7d, 0x33, 0xf3, 0x1f,
	0xd8, 0x0b, 0xb9, 0xda, 0xe7, 0xcb, 0x80, 0x11, 0x46, 0xee, 0x63, 0x9e, 0xb2, 0x6d, 0x7f, 0x33,
	0x08, 0x1e, 0x92, 0x65, 0xd0, 0x67, 0x84, 0xd3, 0x35, 0x0b, 0x09, 0x57, 0x12, 0x46, 0x53, 0x8a,
	0xa4, 0x90, 0xab, 0xca, 0x37, 0xeb, 0xce, 0xd9, 0xf9, 0x97, 0x95, 0xa4, 0xdb, 0x84, 0xf0, 0xe2,
	0x5b, 0x04, 0x4e, 0xde, 0x01, 0xfc, 0xe5, 0x30, 0xba, 0x89, 0x23, 0xc2, 0xac, 0xd5, 0x1d, 0x45,
	0x3d, 0x58, 0xa3, 0x49, 0xf0, 0xb4, 0x26, 0x12, 0x90, 0x41, 0xb7, 0x31, 0xfc, 0xa3, 0x64, 0x95,
	0x45, 0xc4, 0xce, 0x01, 0xde, 0x19, 0x90, 0x0e, 0x61, 0xbe, 0x95, 0x9f, 0x51, 0x49, 0x90, 0x41,
	0xb7, 0x39, 0x3c, 0x55, 0x8e, 0x9d, 0x40, 0x71, 0x33, 0xd1, 0xdb, 0x26, 0x04, 0xd7, 0xf9, 0xfe,
	0x17, 0x49, 0xf0, 0x67, 0x10, 0x45, 0x8c, 0x70, 0x2e, 0x55, 0x64, 0xd0, 0xad, 0xe3, 0xfd, 0x12,
	0xc9, 0xb0, 0x11, 0x11, 0x1e, 0xb2, 0x38, 0x49, 0x63, 0xba, 0x92, 0xaa, 0x39, 0x2d, 0x4b, 0x67,
	0xcf, 0xb0, 0xfe, 0xd5, 0x89, 0xda, 0x10, 0xb9, 0x17, 0x1a, 0x36, 0x7d, 0x6f, 0xee, 0x98, 0xbe,
	0x35, 0x99, 0x69, 0x63, 0xeb, 0x5c, 0xfc, 0x81, 0x5a, 0xf0, 0x77, 0x49, 0x9f, 0xba, 0x26, 0x16,
	0x01, 0xfa, 0x0b, 0xc5, 0x92, 0x38, 0xc2, 0xf6, 0xd4, 0x11, 0x05, 0xd4, 0x81, 0xed, 0x92, 0xea,
	0x4c, 0xf5, 0xb1, 0x65, 0xf8, 0x63, 0x6b, 0x72, 0x29, 0x56, 0x10, 0x82, 0xcd, 0x12, 0xb3, 0x8d,
	0x2b, 0xb1, 0xaa, 0xbf, 0x00, 0xf8, 0x3f, 0xa4, 0x8f, 0x47, 0x6f, 0xac, 0x23, 0xb7, 0xac, 0x3a,
	0xd9, 0xc0, 0x1d, 0x70, 0xd3, 0x3e, 0xe4, 0x4d, 0x16, 0xaf, 0x42, 0xcb, 0xd0, 0xed, 0x6b, 0xd7,
	0xb3, 0xb1, 0x36, 0x32, 0xb1, 0x39, 0xb2, 0x5c, 0x0f, 0xcf, 0xdf, 0x04, 0xc9, 0x70, 0xd5, 0x62,
	0x76, 0x78, 0x17, 0x99, 0x0d, 0xb4, 0x2c, 0xf2, 0x91, 0xa3, 0xdb, 0x43, 0x68, 0x51, 0xcb, 0x5f,
	0x57, 0xfd, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x61, 0x2d, 0x72, 0xe9, 0x3b, 0x02, 0x00, 0x00,
}