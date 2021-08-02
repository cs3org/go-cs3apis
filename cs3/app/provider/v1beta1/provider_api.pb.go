// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs3/app/provider/v1beta1/provider_api.proto

package providerv1beta1

import (
	context "context"
	fmt "fmt"
	v1beta12 "github.com/cs3org/go-cs3apis/cs3/rpc/v1beta1"
	v1beta11 "github.com/cs3org/go-cs3apis/cs3/storage/provider/v1beta1"
	v1beta1 "github.com/cs3org/go-cs3apis/cs3/types/v1beta1"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// REQUIRED.
// View mode.
type OpenInAppRequest_ViewMode int32

const (
	OpenInAppRequest_VIEW_MODE_INVALID OpenInAppRequest_ViewMode = 0
	// The resource can be opened but not downloaded.
	OpenInAppRequest_VIEW_MODE_VIEW_ONLY OpenInAppRequest_ViewMode = 1
	// The resource can be downloaded.
	OpenInAppRequest_VIEW_MODE_READ_ONLY OpenInAppRequest_ViewMode = 2
	// The resource can be downloaded and updated.
	OpenInAppRequest_VIEW_MODE_READ_WRITE OpenInAppRequest_ViewMode = 3
)

var OpenInAppRequest_ViewMode_name = map[int32]string{
	0: "VIEW_MODE_INVALID",
	1: "VIEW_MODE_VIEW_ONLY",
	2: "VIEW_MODE_READ_ONLY",
	3: "VIEW_MODE_READ_WRITE",
}

var OpenInAppRequest_ViewMode_value = map[string]int32{
	"VIEW_MODE_INVALID":    0,
	"VIEW_MODE_VIEW_ONLY":  1,
	"VIEW_MODE_READ_ONLY":  2,
	"VIEW_MODE_READ_WRITE": 3,
}

func (x OpenInAppRequest_ViewMode) String() string {
	return proto.EnumName(OpenInAppRequest_ViewMode_name, int32(x))
}

func (OpenInAppRequest_ViewMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c007b70b037097fe, []int{0, 0}
}

type OpenInAppRequest struct {
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The resourceInfo to be opened. The gateway grpc message has a ref instead.
	ResourceInfo *v1beta11.ResourceInfo    `protobuf:"bytes,2,opt,name=resource_info,json=resourceInfo,proto3" json:"resource_info,omitempty"`
	ViewMode     OpenInAppRequest_ViewMode `protobuf:"varint,3,opt,name=view_mode,json=viewMode,proto3,enum=cs3.app.provider.v1beta1.OpenInAppRequest_ViewMode" json:"view_mode,omitempty"`
	// REQUIRED.
	// The access token this application provider will use when contacting
	// the storage provider to read and write.
	// Service implementors MUST make sure that the access token only grants
	// access to the requested resource.
	// Service implementors should use a ResourceId rather than a filepath to grant access, as
	// ResourceIds MUST NOT change when a resource is renamed.
	// The access token MUST be short-lived.
	// TODO(labkode): investigate token derivation techniques.
	AccessToken          string   `protobuf:"bytes,4,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpenInAppRequest) Reset()         { *m = OpenInAppRequest{} }
func (m *OpenInAppRequest) String() string { return proto.CompactTextString(m) }
func (*OpenInAppRequest) ProtoMessage()    {}
func (*OpenInAppRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c007b70b037097fe, []int{0}
}

func (m *OpenInAppRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenInAppRequest.Unmarshal(m, b)
}
func (m *OpenInAppRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenInAppRequest.Marshal(b, m, deterministic)
}
func (m *OpenInAppRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenInAppRequest.Merge(m, src)
}
func (m *OpenInAppRequest) XXX_Size() int {
	return xxx_messageInfo_OpenInAppRequest.Size(m)
}
func (m *OpenInAppRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenInAppRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OpenInAppRequest proto.InternalMessageInfo

func (m *OpenInAppRequest) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *OpenInAppRequest) GetResourceInfo() *v1beta11.ResourceInfo {
	if m != nil {
		return m.ResourceInfo
	}
	return nil
}

func (m *OpenInAppRequest) GetViewMode() OpenInAppRequest_ViewMode {
	if m != nil {
		return m.ViewMode
	}
	return OpenInAppRequest_VIEW_MODE_INVALID
}

func (m *OpenInAppRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

type OpenInAppResponse struct {
	// REQUIRED.
	// The response status.
	Status *v1beta12.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The url that user agents will render to clients.
	// Usually the rendering happens by using HTML iframes or in separate browser tabs.
	AppUrl               string   `protobuf:"bytes,3,opt,name=app_url,json=appUrl,proto3" json:"app_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpenInAppResponse) Reset()         { *m = OpenInAppResponse{} }
func (m *OpenInAppResponse) String() string { return proto.CompactTextString(m) }
func (*OpenInAppResponse) ProtoMessage()    {}
func (*OpenInAppResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c007b70b037097fe, []int{1}
}

func (m *OpenInAppResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenInAppResponse.Unmarshal(m, b)
}
func (m *OpenInAppResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenInAppResponse.Marshal(b, m, deterministic)
}
func (m *OpenInAppResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenInAppResponse.Merge(m, src)
}
func (m *OpenInAppResponse) XXX_Size() int {
	return xxx_messageInfo_OpenInAppResponse.Size(m)
}
func (m *OpenInAppResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenInAppResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OpenInAppResponse proto.InternalMessageInfo

func (m *OpenInAppResponse) GetStatus() *v1beta12.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *OpenInAppResponse) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *OpenInAppResponse) GetAppUrl() string {
	if m != nil {
		return m.AppUrl
	}
	return ""
}

func init() {
	proto.RegisterEnum("cs3.app.provider.v1beta1.OpenInAppRequest_ViewMode", OpenInAppRequest_ViewMode_name, OpenInAppRequest_ViewMode_value)
	proto.RegisterType((*OpenInAppRequest)(nil), "cs3.app.provider.v1beta1.OpenInAppRequest")
	proto.RegisterType((*OpenInAppResponse)(nil), "cs3.app.provider.v1beta1.OpenInAppResponse")
}

func init() {
	proto.RegisterFile("cs3/app/provider/v1beta1/provider_api.proto", fileDescriptor_c007b70b037097fe)
}

var fileDescriptor_c007b70b037097fe = []byte{
	// 489 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xd1, 0x6e, 0xd3, 0x3e,
	0x14, 0xc6, 0xff, 0x49, 0xff, 0x2a, 0xab, 0x3b, 0xa0, 0x33, 0xa0, 0x86, 0x6a, 0x48, 0xa5, 0x57,
	0xd5, 0x86, 0x5c, 0xb5, 0x79, 0x82, 0x74, 0xed, 0x45, 0xa4, 0x6d, 0x89, 0xcc, 0xe8, 0x04, 0xaa,
	0x14, 0x79, 0xa9, 0x87, 0x22, 0xd6, 0xf8, 0xcc, 0x4e, 0x3a, 0x71, 0xc5, 0x2b, 0xf0, 0x0c, 0x5c,
	0xf2, 0x28, 0x3c, 0x12, 0x57, 0x28, 0x8e, 0x13, 0xca, 0xa6, 0x49, 0xbd, 0xb3, 0xcf, 0xf7, 0xfb,
	0xec, 0x73, 0x3e, 0x27, 0xe8, 0x38, 0x56, 0xee, 0x88, 0x01, 0x8c, 0x40, 0x8a, 0x4d, 0xb2, 0xe2,
	0x72, 0xb4, 0x19, 0x5f, 0xf1, 0x8c, 0x8d, 0xeb, 0x42, 0xc4, 0x20, 0x21, 0x20, 0x45, 0x26, 0xb0,
	0x13, 0x2b, 0x97, 0x30, 0x00, 0x52, 0x69, 0xc4, 0xc0, 0xbd, 0xc3, 0xe2, 0x18, 0x09, 0x71, 0xed,
	0x56, 0x19, 0xcb, 0x72, 0x55, 0xfa, 0x7a, 0xef, 0x0a, 0x55, 0x65, 0x42, 0xb2, 0xcf, 0xfc, 0xe1,
	0x45, 0x92, 0x2b, 0x91, 0xcb, 0x98, 0x57, 0xf4, 0x9b, 0x82, 0xce, 0xbe, 0x02, 0x57, 0x35, 0xa2,
	0x77, 0xa5, 0x3c, 0xf8, 0x6d, 0xa3, 0x4e, 0x00, 0x3c, 0xf5, 0x53, 0x0f, 0x80, 0xf2, 0xdb, 0x9c,
	0xab, 0x0c, 0x8f, 0x51, 0x53, 0x00, 0xbb, 0xcd, 0xb9, 0x63, 0xf5, 0xad, 0x61, 0x7b, 0xf2, 0x9a,
	0x14, 0xad, 0x96, 0x36, 0x73, 0x08, 0x09, 0x34, 0x40, 0x0d, 0x88, 0x03, 0xf4, 0xb4, 0xba, 0x39,
	0x4a, 0xd2, 0x6b, 0xe1, 0xd8, 0xda, 0x79, 0xa4, 0x9d, 0xa6, 0xd9, 0x07, 0x83, 0x12, 0x6a, 0x2c,
	0x7e, 0x7a, 0x2d, 0xe8, 0xbe, 0xdc, 0xda, 0xe1, 0x10, 0xb5, 0x36, 0x09, 0xbf, 0x8b, 0xd6, 0x62,
	0xc5, 0x9d, 0x46, 0xdf, 0x1a, 0x3e, 0x9b, 0xb8, 0xe4, 0xb1, 0xc4, 0xc8, 0xfd, 0x11, 0xc8, 0x22,
	0xe1, 0x77, 0x67, 0x62, 0xc5, 0xe9, 0xde, 0xc6, 0xac, 0xf0, 0x5b, 0xb4, 0xcf, 0xe2, 0x98, 0x2b,
	0x15, 0x65, 0xe2, 0x0b, 0x4f, 0x9d, 0xff, 0xfb, 0xd6, 0xb0, 0x45, 0xdb, 0x65, 0xed, 0xa2, 0x28,
	0x0d, 0xd6, 0x68, 0xaf, 0x32, 0xe2, 0x57, 0xe8, 0x60, 0xe1, 0xcf, 0x2f, 0xa3, 0xb3, 0x60, 0x36,
	0x8f, 0xfc, 0xf3, 0x85, 0x77, 0xea, 0xcf, 0x3a, 0xff, 0xe1, 0x2e, 0x7a, 0xf1, 0xb7, 0xac, 0x57,
	0xc1, 0xf9, 0xe9, 0xc7, 0x8e, 0xf5, 0xaf, 0x40, 0xe7, 0xde, 0xac, 0x14, 0x6c, 0xec, 0xa0, 0x97,
	0xf7, 0x84, 0x4b, 0xea, 0x5f, 0xcc, 0x3b, 0x8d, 0xc1, 0x77, 0x0b, 0x1d, 0x6c, 0x75, 0xae, 0x40,
	0xa4, 0x8a, 0xe3, 0x11, 0x6a, 0x96, 0xef, 0x6d, 0xd2, 0xef, 0xea, 0xb1, 0x25, 0xc4, 0xf5, 0xb4,
	0xef, 0xb5, 0x4c, 0x0d, 0xb6, 0xf5, 0x5c, 0xf6, 0xae, 0xcf, 0xd5, 0x45, 0x4f, 0x18, 0x40, 0x94,
	0xcb, 0x1b, 0x9d, 0x6d, 0x8b, 0x36, 0x19, 0xc0, 0x07, 0x79, 0x33, 0x51, 0xa8, 0x1d, 0x9a, 0x70,
	0xbd, 0xd0, 0xc7, 0x2b, 0xd4, 0xaa, 0x1b, 0xc4, 0x47, 0xbb, 0xe7, 0xdf, 0x3b, 0xde, 0x89, 0x2d,
	0x27, 0x9e, 0x7e, 0x43, 0x87, 0xb1, 0x58, 0x3f, 0xea, 0x98, 0x76, 0xea, 0x96, 0x20, 0x09, 0x8b,
	0xcf, 0x36, 0xb4, 0x3e, 0x3d, 0xaf, 0x28, 0x03, 0xfd, 0xb0, 0x1b, 0x27, 0x5e, 0xf8, 0xd3, 0x76,
	0x4e, 0x94, 0x4b, 0x3c, 0x00, 0x52, 0x79, 0xc8, 0x62, 0x3c, 0x2d, 0x80, 0x5f, 0x5a, 0x5a, 0x7a,
	0x00, 0xcb, 0x4a, 0x5a, 0x1a, 0xe9, 0xaa, 0xa9, 0x7f, 0x06, 0xf7, 0x4f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x7b, 0x0a, 0x82, 0xae, 0xc0, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProviderAPIClient is the client API for ProviderAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProviderAPIClient interface {
	// Returns the App provider URL
	// MUST return CODE_NOT_FOUND if the resource does not exist.
	OpenInApp(ctx context.Context, in *OpenInAppRequest, opts ...grpc.CallOption) (*OpenInAppResponse, error)
}

type providerAPIClient struct {
	cc *grpc.ClientConn
}

func NewProviderAPIClient(cc *grpc.ClientConn) ProviderAPIClient {
	return &providerAPIClient{cc}
}

func (c *providerAPIClient) OpenInApp(ctx context.Context, in *OpenInAppRequest, opts ...grpc.CallOption) (*OpenInAppResponse, error) {
	out := new(OpenInAppResponse)
	err := c.cc.Invoke(ctx, "/cs3.app.provider.v1beta1.ProviderAPI/OpenInApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProviderAPIServer is the server API for ProviderAPI service.
type ProviderAPIServer interface {
	// Returns the App provider URL
	// MUST return CODE_NOT_FOUND if the resource does not exist.
	OpenInApp(context.Context, *OpenInAppRequest) (*OpenInAppResponse, error)
}

// UnimplementedProviderAPIServer can be embedded to have forward compatible implementations.
type UnimplementedProviderAPIServer struct {
}

func (*UnimplementedProviderAPIServer) OpenInApp(ctx context.Context, req *OpenInAppRequest) (*OpenInAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenInApp not implemented")
}

func RegisterProviderAPIServer(s *grpc.Server, srv ProviderAPIServer) {
	s.RegisterService(&_ProviderAPI_serviceDesc, srv)
}

func _ProviderAPI_OpenInApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenInAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderAPIServer).OpenInApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cs3.app.provider.v1beta1.ProviderAPI/OpenInApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderAPIServer).OpenInApp(ctx, req.(*OpenInAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProviderAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cs3.app.provider.v1beta1.ProviderAPI",
	HandlerType: (*ProviderAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OpenInApp",
			Handler:    _ProviderAPI_OpenInApp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cs3/app/provider/v1beta1/provider_api.proto",
}
