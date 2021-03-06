// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs3/ocm/invite/v1beta1/resources.proto

package invitev1beta1

import (
	fmt "fmt"
	v1beta1 "github.com/cs3org/go-cs3apis/cs3/identity/user/v1beta1"
	v1beta11 "github.com/cs3org/go-cs3apis/cs3/types/v1beta1"
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

// InviteToken is used to invite users and groups from other sync'n'share
// systems to collaborate on resources.
type InviteToken struct {
	// REQUIRED.
	// Unique ID associated with an InviteToken.
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// REQUIRED.
	// The user who created the token.
	UserId *v1beta1.UserId `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// OPTIONAL.
	// The time when the token will expire.
	Expiration *v1beta11.Timestamp `protobuf:"bytes,3,opt,name=expiration,proto3" json:"expiration,omitempty"`
	// OPTIONAL.
	// User-defined description to be forwarded to the invitees.
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InviteToken) Reset()         { *m = InviteToken{} }
func (m *InviteToken) String() string { return proto.CompactTextString(m) }
func (*InviteToken) ProtoMessage()    {}
func (*InviteToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bdba69a40d5b11f, []int{0}
}

func (m *InviteToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InviteToken.Unmarshal(m, b)
}
func (m *InviteToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InviteToken.Marshal(b, m, deterministic)
}
func (m *InviteToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InviteToken.Merge(m, src)
}
func (m *InviteToken) XXX_Size() int {
	return xxx_messageInfo_InviteToken.Size(m)
}
func (m *InviteToken) XXX_DiscardUnknown() {
	xxx_messageInfo_InviteToken.DiscardUnknown(m)
}

var xxx_messageInfo_InviteToken proto.InternalMessageInfo

func (m *InviteToken) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *InviteToken) GetUserId() *v1beta1.UserId {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *InviteToken) GetExpiration() *v1beta11.Timestamp {
	if m != nil {
		return m.Expiration
	}
	return nil
}

func (m *InviteToken) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*InviteToken)(nil), "cs3.ocm.invite.v1beta1.InviteToken")
}

func init() {
	proto.RegisterFile("cs3/ocm/invite/v1beta1/resources.proto", fileDescriptor_9bdba69a40d5b11f)
}

var fileDescriptor_9bdba69a40d5b11f = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xb1, 0x4b, 0xc4, 0x30,
	0x14, 0xc6, 0xe9, 0x9d, 0x9e, 0x98, 0x43, 0x87, 0x22, 0x47, 0x29, 0x0a, 0xd5, 0x41, 0xce, 0xe5,
	0x95, 0xda, 0x4d, 0x9c, 0x7a, 0x53, 0xa7, 0x1e, 0xe5, 0x74, 0x90, 0x82, 0xf4, 0xd2, 0x37, 0x04,
	0x49, 0x53, 0x92, 0xf4, 0xf0, 0xfe, 0x1d, 0x47, 0xff, 0x09, 0x77, 0xff, 0x2a, 0x49, 0xd2, 0xca,
	0x0d, 0x3a, 0xb5, 0x2f, 0xdf, 0xef, 0x7b, 0xf9, 0xf2, 0x91, 0x5b, 0xaa, 0xd2, 0x58, 0x50, 0x1e,
	0xb3, 0x76, 0xc7, 0x34, 0xc6, 0xbb, 0x64, 0x8b, 0xba, 0x4e, 0x62, 0x89, 0x4a, 0xf4, 0x92, 0xa2,
	0x82, 0x4e, 0x0a, 0x2d, 0xfc, 0x05, 0x55, 0x29, 0x08, 0xca, 0xc1, 0x71, 0x30, 0x70, 0xe1, 0x9d,
	0xf1, 0xb3, 0x06, 0x5b, 0xcd, 0xf4, 0x3e, 0xee, 0x15, 0xca, 0xff, 0x56, 0x84, 0x57, 0x06, 0xd5,
	0xfb, 0x0e, 0xd5, 0x2f, 0x62, 0x27, 0x27, 0xdf, 0x7c, 0x79, 0x64, 0x9e, 0xdb, 0xe5, 0x1b, 0xf1,
	0x86, 0xad, 0x7f, 0x41, 0x8e, 0xb5, 0xf9, 0x09, 0xbc, 0xc8, 0x5b, 0x9e, 0x96, 0x6e, 0xf0, 0x1f,
	0xc8, 0x89, 0xb9, 0xe4, 0x95, 0x35, 0xc1, 0x24, 0xf2, 0x96, 0xf3, 0xfb, 0x6b, 0x30, 0xc9, 0xc6,
	0x04, 0x60, 0xc4, 0x31, 0x1c, 0x3c, 0x29, 0x94, 0x79, 0x53, 0xce, 0x7a, 0xfb, 0xf5, 0x1f, 0x09,
	0xc1, 0xf7, 0x8e, 0xc9, 0x5a, 0x33, 0xd1, 0x06, 0x53, 0x6b, 0xbf, 0xb4, 0x76, 0x97, 0x63, 0xb4,
	0x6d, 0x18, 0x47, 0xa5, 0x6b, 0xde, 0x95, 0x07, 0xbc, 0x1f, 0x91, 0x79, 0x83, 0x8a, 0x4a, 0xd6,
	0x59, 0xfb, 0x91, 0x4d, 0x75, 0x78, 0x94, 0xf5, 0x24, 0xa4, 0x82, 0xc3, 0xdf, 0x4d, 0x65, 0xe7,
	0xe5, 0xd8, 0xc7, 0xda, 0xbc, 0x77, 0xed, 0xbd, 0x9c, 0x39, 0x62, 0x00, 0x3e, 0x26, 0xd3, 0x55,
	0x91, 0x7f, 0x4e, 0x16, 0x2b, 0x95, 0x42, 0x41, 0x39, 0xb8, 0x32, 0xe0, 0x39, 0xc9, 0x8c, 0xfc,
	0x6d, 0x85, 0xaa, 0xa0, 0xbc, 0x72, 0x42, 0x35, 0x08, 0xdb, 0x99, 0xed, 0x2f, 0xfd, 0x09, 0x00,
	0x00, 0xff, 0xff, 0x38, 0xee, 0x9a, 0xa2, 0xcb, 0x01, 0x00, 0x00,
}
