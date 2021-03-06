// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs3/tx/v1beta1/resources.proto

package txv1beta1

import (
	fmt "fmt"
	v1beta11 "github.com/cs3org/go-cs3apis/cs3/identity/user/v1beta1"
	v1beta1 "github.com/cs3org/go-cs3apis/cs3/storage/provider/v1beta1"
	v1beta12 "github.com/cs3org/go-cs3apis/cs3/types/v1beta1"
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

// Status represents transfer status.
type TxInfo_Status int32

const (
	TxInfo_STATUS_INVALID TxInfo_Status = 0
	// The destination could not be found.
	TxInfo_STATUS_DESTINATION_NOT_FOUND TxInfo_Status = 1
	// A new data transfer
	TxInfo_STATUS_TRANSFER_NEW TxInfo_Status = 2
	// The data transfer is awaiting acceptance from the destination
	TxInfo_STATUS_TRANSFER_AWAITING_ACCEPTANCE TxInfo_Status = 3
	// The data transfer is accepted by the destination.
	TxInfo_STATUS_TRANSFER_ACCEPTED TxInfo_Status = 4
	// The data transfer has started and not yet completed.
	TxInfo_STATUS_TRANSFER_IN_PROGRESS TxInfo_Status = 5
	// The data transfer has completed.
	TxInfo_STATUS_TRANSFER_COMPLETE TxInfo_Status = 6
	// The data transfer has failed.
	TxInfo_STATUS_TRANSFER_FAILED TxInfo_Status = 7
	// The data transfer had been cancelled.
	TxInfo_STATUS_TRANSFER_CANCELLED TxInfo_Status = 8
	// The request for cancelling the data transfer has failed.
	TxInfo_STATUS_TRANSFER_CANCEL_FAILED TxInfo_Status = 9
	// The transfer has expired somewhere down the line.
	TxInfo_STATUS_TRANSFER_EXPIRED TxInfo_Status = 10
)

var TxInfo_Status_name = map[int32]string{
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

var TxInfo_Status_value = map[string]int32{
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

func (x TxInfo_Status) String() string {
	return proto.EnumName(TxInfo_Status_name, int32(x))
}

func (TxInfo_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e6e9f5a73c9e5a84, []int{1, 0}
}

// TxId uniquely identifies a transfer in the transfer provider namespace.
type TxId struct {
	// REQUIRED.
	// The internal transfer id used by the service implementor
	// to uniquely identity the transfer in the internal
	// implementation of the service.
	OpaqueId             string   `protobuf:"bytes,1,opt,name=opaque_id,json=opaqueId,proto3" json:"opaque_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TxId) Reset()         { *m = TxId{} }
func (m *TxId) String() string { return proto.CompactTextString(m) }
func (*TxId) ProtoMessage()    {}
func (*TxId) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6e9f5a73c9e5a84, []int{0}
}

func (m *TxId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxId.Unmarshal(m, b)
}
func (m *TxId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxId.Marshal(b, m, deterministic)
}
func (m *TxId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxId.Merge(m, src)
}
func (m *TxId) XXX_Size() int {
	return xxx_messageInfo_TxId.Size(m)
}
func (m *TxId) XXX_DiscardUnknown() {
	xxx_messageInfo_TxId.DiscardUnknown(m)
}

var xxx_messageInfo_TxId proto.InternalMessageInfo

func (m *TxId) GetOpaqueId() string {
	if m != nil {
		return m.OpaqueId
	}
	return ""
}

// TxInfo represents information about a transfer.
type TxInfo struct {
	// REQUIRED.
	// The transfer identifier.
	Id *TxId `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// REQUIRED.
	// Reference of the resource to be transfered.
	Ref *v1beta1.Reference `protobuf:"bytes,2,opt,name=ref,proto3" json:"ref,omitempty"`
	// REQUIRED.
	// The transfer status. Eg.: STATUS_TRANSFER_FAILED.
	// Note: the description field may provide additional information on the transfer status.
	Status TxInfo_Status `protobuf:"varint,3,opt,name=status,proto3,enum=cs3.tx.v1beta1.TxInfo_Status" json:"status,omitempty"`
	// REQUIRED.
	// The destination (receiver of the transfer)
	Grantee *v1beta1.Grantee `protobuf:"bytes,4,opt,name=grantee,proto3" json:"grantee,omitempty"`
	// REQUIRED.
	// Uniquely identifies a principal who initiates the transfer creation.
	Creator *v1beta11.UserId `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty"`
	// REQUIRED.
	// Creation time of the transfer.
	Ctime *v1beta12.Timestamp `protobuf:"bytes,6,opt,name=ctime,proto3" json:"ctime,omitempty"`
	// OPTIONAL.
	// Information to describe the transfer status.
	// Eg. may contain information about a transfer failure.
	// Meant to be human-readable.
	Description          string   `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TxInfo) Reset()         { *m = TxInfo{} }
func (m *TxInfo) String() string { return proto.CompactTextString(m) }
func (*TxInfo) ProtoMessage()    {}
func (*TxInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6e9f5a73c9e5a84, []int{1}
}

func (m *TxInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxInfo.Unmarshal(m, b)
}
func (m *TxInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxInfo.Marshal(b, m, deterministic)
}
func (m *TxInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxInfo.Merge(m, src)
}
func (m *TxInfo) XXX_Size() int {
	return xxx_messageInfo_TxInfo.Size(m)
}
func (m *TxInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TxInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TxInfo proto.InternalMessageInfo

func (m *TxInfo) GetId() *TxId {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *TxInfo) GetRef() *v1beta1.Reference {
	if m != nil {
		return m.Ref
	}
	return nil
}

func (m *TxInfo) GetStatus() TxInfo_Status {
	if m != nil {
		return m.Status
	}
	return TxInfo_STATUS_INVALID
}

func (m *TxInfo) GetGrantee() *v1beta1.Grantee {
	if m != nil {
		return m.Grantee
	}
	return nil
}

func (m *TxInfo) GetCreator() *v1beta11.UserId {
	if m != nil {
		return m.Creator
	}
	return nil
}

func (m *TxInfo) GetCtime() *v1beta12.Timestamp {
	if m != nil {
		return m.Ctime
	}
	return nil
}

func (m *TxInfo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterEnum("cs3.tx.v1beta1.TxInfo_Status", TxInfo_Status_name, TxInfo_Status_value)
	proto.RegisterType((*TxId)(nil), "cs3.tx.v1beta1.TxId")
	proto.RegisterType((*TxInfo)(nil), "cs3.tx.v1beta1.TxInfo")
}

func init() { proto.RegisterFile("cs3/tx/v1beta1/resources.proto", fileDescriptor_e6e9f5a73c9e5a84) }

var fileDescriptor_e6e9f5a73c9e5a84 = []byte{
	// 572 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x69, 0xba, 0xb6, 0xeb, 0x3b, 0xa9, 0x8a, 0x0c, 0x62, 0x61, 0x5b, 0xa1, 0xdb, 0x40,
	0x1b, 0x12, 0x4a, 0xb5, 0x55, 0x1c, 0x10, 0x07, 0x94, 0xb5, 0xde, 0x64, 0xa9, 0xa4, 0x95, 0xe3,
	0xfd, 0x01, 0x4d, 0x8a, 0xb2, 0xc4, 0x9d, 0x72, 0x68, 0x53, 0x6c, 0x77, 0xca, 0xbe, 0x0e, 0x47,
	0x3e, 0x0a, 0x17, 0xae, 0x7c, 0x1c, 0x14, 0x37, 0xa9, 0xa0, 0x2b, 0x3b, 0xe6, 0x79, 0x7e, 0xcf,
	0xeb, 0xc7, 0x76, 0x0c, 0x2f, 0x43, 0xd9, 0x69, 0xab, 0xb4, 0x7d, 0x77, 0x74, 0xc3, 0x55, 0x70,
	0xd4, 0x16, 0x5c, 0x26, 0x33, 0x11, 0x72, 0x69, 0x4f, 0x45, 0xa2, 0x12, 0xd4, 0x08, 0x65, 0xc7,
	0x56, 0xa9, 0x9d, 0xfb, 0x5b, 0x6f, 0x33, 0x3e, 0x8e, 0xf8, 0x44, 0xc5, 0xea, 0xbe, 0x3d, 0x93,
	0x5c, 0xfc, 0x2f, 0xba, 0xf5, 0x2e, 0x43, 0xa5, 0x4a, 0x44, 0x70, 0xcb, 0xdb, 0x53, 0x91, 0xdc,
	0xc5, 0xd1, 0x23, 0x74, 0x53, 0x17, 0xb9, 0x9f, 0x72, 0xb9, 0x40, 0xf4, 0xd7, 0xdc, 0xde, 0xdb,
	0x87, 0x35, 0x96, 0x92, 0x08, 0x6d, 0x43, 0x3d, 0x99, 0x06, 0xdf, 0x66, 0xdc, 0x8f, 0x23, 0xab,
	0xd4, 0x2a, 0x1d, 0xd6, 0xe9, 0xfa, 0x5c, 0x20, 0xd1, 0xde, 0xaf, 0x0a, 0x54, 0x59, 0x4a, 0x26,
	0xa3, 0x04, 0xbd, 0x06, 0x23, 0x07, 0x36, 0x8e, 0x9f, 0xd9, 0xff, 0x6e, 0xc2, 0xce, 0x26, 0x51,
	0x23, 0x8e, 0xd0, 0x07, 0x28, 0x0b, 0x3e, 0xb2, 0x0c, 0x8d, 0x1d, 0x68, 0x2c, 0x2f, 0x6c, 0x17,
	0x85, 0x17, 0x21, 0xca, 0x47, 0x5c, 0xf0, 0x49, 0xc8, 0x69, 0x96, 0x41, 0xef, 0xa1, 0x2a, 0x55,
	0xa0, 0x66, 0xd2, 0x2a, 0xb7, 0x4a, 0x87, 0x8d, 0xe3, 0xe6, 0x8a, 0x45, 0x26, 0xa3, 0xc4, 0xf6,
	0x34, 0x44, 0x73, 0x18, 0x7d, 0x82, 0xda, 0xad, 0x08, 0x26, 0x8a, 0x73, 0x6b, 0x4d, 0xaf, 0xfa,
	0xe6, 0xf1, 0x55, 0xcf, 0xe6, 0x30, 0x2d, 0x52, 0xe8, 0x23, 0xd4, 0x42, 0xc1, 0x03, 0x95, 0x08,
	0xab, 0xa2, 0x07, 0xec, 0xea, 0x01, 0xc5, 0x95, 0xd8, 0xd9, 0x95, 0x2c, 0xd2, 0xe7, 0x92, 0x0b,
	0x12, 0xd1, 0x22, 0x81, 0x8e, 0xa1, 0x12, 0xaa, 0x78, 0xcc, 0xad, 0xaa, 0x8e, 0xee, 0xcc, 0x3b,
	0xeb, 0x63, 0x5e, 0xd4, 0x8e, 0xc7, 0x5c, 0xaa, 0x60, 0x3c, 0xa5, 0x73, 0x14, 0xb5, 0x60, 0x23,
	0xe2, 0x32, 0x14, 0xf1, 0x54, 0xc5, 0xc9, 0xc4, 0xaa, 0xe9, 0x33, 0xff, 0x5b, 0xda, 0xfb, 0x6d,
	0x40, 0x75, 0xbe, 0x4d, 0x84, 0xa0, 0xe1, 0x31, 0x87, 0x9d, 0x7b, 0x3e, 0x71, 0x2f, 0x9c, 0x3e,
	0xe9, 0x99, 0x4f, 0x50, 0x0b, 0x76, 0x72, 0xad, 0x87, 0x3d, 0x46, 0x5c, 0x87, 0x91, 0x81, 0xeb,
	0xbb, 0x03, 0xe6, 0x9f, 0x0e, 0xce, 0xdd, 0x9e, 0x59, 0x42, 0x9b, 0xf0, 0x34, 0x27, 0x18, 0x75,
	0x5c, 0xef, 0x14, 0x53, 0xdf, 0xc5, 0x97, 0xa6, 0x81, 0x0e, 0x60, 0x7f, 0xd9, 0x70, 0x2e, 0x1d,
	0xc2, 0x88, 0x7b, 0xe6, 0x3b, 0xdd, 0x2e, 0x1e, 0x32, 0xc7, 0xed, 0x62, 0xb3, 0x8c, 0x76, 0xc0,
	0x7a, 0x00, 0x6a, 0x1f, 0xf7, 0xcc, 0x35, 0xf4, 0x0a, 0xb6, 0x97, 0x5d, 0xe2, 0xfa, 0x43, 0x3a,
	0x38, 0xa3, 0xd8, 0xf3, 0xcc, 0xca, 0xaa, 0x78, 0x77, 0xf0, 0x79, 0xd8, 0xc7, 0x0c, 0x9b, 0x55,
	0xb4, 0x05, 0xcf, 0x97, 0xdd, 0x53, 0x87, 0xf4, 0x71, 0xcf, 0xac, 0xa1, 0x26, 0xbc, 0x78, 0x90,
	0xcc, 0x3a, 0xf5, 0x33, 0x7b, 0x1d, 0xed, 0x42, 0x73, 0xb5, 0x5d, 0x4c, 0xa8, 0xa3, 0x6d, 0xd8,
	0x5c, 0x46, 0xf0, 0xd5, 0x90, 0x50, 0xdc, 0x33, 0xe1, 0xe4, 0x0b, 0xa0, 0x30, 0x19, 0x2f, 0xfd,
	0x5a, 0x27, 0x0d, 0x5a, 0x3c, 0x9e, 0x61, 0xf6, 0x38, 0x86, 0xa5, 0xaf, 0x75, 0x95, 0xe6, 0xe6,
	0x77, 0xa3, 0xdc, 0x65, 0x57, 0x3f, 0x8c, 0x46, 0x57, 0x76, 0x6c, 0x96, 0xda, 0x17, 0x47, 0x27,
	0x99, 0xfc, 0x53, 0x0b, 0xd7, 0x2c, 0xbd, 0xce, 0x85, 0x9b, 0xaa, 0x7e, 0x58, 0x9d, 0x3f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x66, 0xe7, 0x2d, 0x09, 0x02, 0x04, 0x00, 0x00,
}
