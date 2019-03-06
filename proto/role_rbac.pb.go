// Code generated by protoc-gen-go. DO NOT EDIT.
// source: role_rbac.proto

package test

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type Rep struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Rep) Reset()         { *m = Rep{} }
func (m *Rep) String() string { return proto.CompactTextString(m) }
func (*Rep) ProtoMessage()    {}
func (*Rep) Descriptor() ([]byte, []int) {
	return fileDescriptor_54744e7db0d1daef, []int{0}
}

func (m *Rep) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rep.Unmarshal(m, b)
}
func (m *Rep) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rep.Marshal(b, m, deterministic)
}
func (m *Rep) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rep.Merge(m, src)
}
func (m *Rep) XXX_Size() int {
	return xxx_messageInfo_Rep.Size(m)
}
func (m *Rep) XXX_DiscardUnknown() {
	xxx_messageInfo_Rep.DiscardUnknown(m)
}

var xxx_messageInfo_Rep proto.InternalMessageInfo

type EnforceRep struct {
	Sub                  []string `protobuf:"bytes,1,rep,name=sub,proto3" json:"sub,omitempty"`
	Obj                  string   `protobuf:"bytes,2,opt,name=obj,proto3" json:"obj,omitempty"`
	Act                  string   `protobuf:"bytes,3,opt,name=act,proto3" json:"act,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnforceRep) Reset()         { *m = EnforceRep{} }
func (m *EnforceRep) String() string { return proto.CompactTextString(m) }
func (*EnforceRep) ProtoMessage()    {}
func (*EnforceRep) Descriptor() ([]byte, []int) {
	return fileDescriptor_54744e7db0d1daef, []int{1}
}

func (m *EnforceRep) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnforceRep.Unmarshal(m, b)
}
func (m *EnforceRep) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnforceRep.Marshal(b, m, deterministic)
}
func (m *EnforceRep) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnforceRep.Merge(m, src)
}
func (m *EnforceRep) XXX_Size() int {
	return xxx_messageInfo_EnforceRep.Size(m)
}
func (m *EnforceRep) XXX_DiscardUnknown() {
	xxx_messageInfo_EnforceRep.DiscardUnknown(m)
}

var xxx_messageInfo_EnforceRep proto.InternalMessageInfo

func (m *EnforceRep) GetSub() []string {
	if m != nil {
		return m.Sub
	}
	return nil
}

func (m *EnforceRep) GetObj() string {
	if m != nil {
		return m.Obj
	}
	return ""
}

func (m *EnforceRep) GetAct() string {
	if m != nil {
		return m.Act
	}
	return ""
}

type RBACInfo struct {
	Sub                  string   `protobuf:"bytes,1,opt,name=sub,proto3" json:"sub,omitempty"`
	Obj                  string   `protobuf:"bytes,2,opt,name=obj,proto3" json:"obj,omitempty"`
	Act                  string   `protobuf:"bytes,3,opt,name=act,proto3" json:"act,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RBACInfo) Reset()         { *m = RBACInfo{} }
func (m *RBACInfo) String() string { return proto.CompactTextString(m) }
func (*RBACInfo) ProtoMessage()    {}
func (*RBACInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_54744e7db0d1daef, []int{2}
}

func (m *RBACInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RBACInfo.Unmarshal(m, b)
}
func (m *RBACInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RBACInfo.Marshal(b, m, deterministic)
}
func (m *RBACInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RBACInfo.Merge(m, src)
}
func (m *RBACInfo) XXX_Size() int {
	return xxx_messageInfo_RBACInfo.Size(m)
}
func (m *RBACInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RBACInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RBACInfo proto.InternalMessageInfo

func (m *RBACInfo) GetSub() string {
	if m != nil {
		return m.Sub
	}
	return ""
}

func (m *RBACInfo) GetObj() string {
	if m != nil {
		return m.Obj
	}
	return ""
}

func (m *RBACInfo) GetAct() string {
	if m != nil {
		return m.Act
	}
	return ""
}

type EnforceRes struct {
	Code                 int64    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnforceRes) Reset()         { *m = EnforceRes{} }
func (m *EnforceRes) String() string { return proto.CompactTextString(m) }
func (*EnforceRes) ProtoMessage()    {}
func (*EnforceRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_54744e7db0d1daef, []int{3}
}

func (m *EnforceRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnforceRes.Unmarshal(m, b)
}
func (m *EnforceRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnforceRes.Marshal(b, m, deterministic)
}
func (m *EnforceRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnforceRes.Merge(m, src)
}
func (m *EnforceRes) XXX_Size() int {
	return xxx_messageInfo_EnforceRes.Size(m)
}
func (m *EnforceRes) XXX_DiscardUnknown() {
	xxx_messageInfo_EnforceRes.DiscardUnknown(m)
}

var xxx_messageInfo_EnforceRes proto.InternalMessageInfo

func (m *EnforceRes) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *EnforceRes) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type ModActRes struct {
	Code                 int64       `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string      `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	DataInfo             []*RBACInfo `protobuf:"bytes,3,rep,name=data_info,json=dataInfo,proto3" json:"data_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ModActRes) Reset()         { *m = ModActRes{} }
func (m *ModActRes) String() string { return proto.CompactTextString(m) }
func (*ModActRes) ProtoMessage()    {}
func (*ModActRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_54744e7db0d1daef, []int{4}
}

func (m *ModActRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModActRes.Unmarshal(m, b)
}
func (m *ModActRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModActRes.Marshal(b, m, deterministic)
}
func (m *ModActRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModActRes.Merge(m, src)
}
func (m *ModActRes) XXX_Size() int {
	return xxx_messageInfo_ModActRes.Size(m)
}
func (m *ModActRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ModActRes.DiscardUnknown(m)
}

var xxx_messageInfo_ModActRes proto.InternalMessageInfo

func (m *ModActRes) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ModActRes) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *ModActRes) GetDataInfo() []*RBACInfo {
	if m != nil {
		return m.DataInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*Rep)(nil), "test.Rep")
	proto.RegisterType((*EnforceRep)(nil), "test.EnforceRep")
	proto.RegisterType((*RBACInfo)(nil), "test.RBACInfo")
	proto.RegisterType((*EnforceRes)(nil), "test.EnforceRes")
	proto.RegisterType((*ModActRes)(nil), "test.ModActRes")
}

func init() { proto.RegisterFile("role_rbac.proto", fileDescriptor_54744e7db0d1daef) }

var fileDescriptor_54744e7db0d1daef = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x1b, 0x1c, 0x4a, 0x73, 0x48, 0xb4, 0xba, 0x29, 0xea, 0x14, 0x79, 0x40, 0x95, 0x90,
	0x82, 0x14, 0xfe, 0x00, 0x2d, 0x62, 0x60, 0x60, 0xf1, 0x0f, 0x20, 0xb2, 0x1d, 0xa7, 0xa2, 0xa2,
	0xb9, 0x28, 0x36, 0xff, 0x1f, 0x5d, 0xdc, 0x12, 0xc4, 0x04, 0xdb, 0xd3, 0xe7, 0xbb, 0xe7, 0xe7,
	0x67, 0x58, 0x0e, 0xf4, 0xe1, 0xea, 0xc1, 0x68, 0x5b, 0xf6, 0x03, 0x05, 0xc2, 0x34, 0x38, 0x1f,
	0xe4, 0x25, 0x08, 0xe5, 0x7a, 0xb9, 0x03, 0x78, 0xee, 0x5a, 0x1a, 0xac, 0x53, 0xae, 0xc7, 0x15,
	0x08, 0xff, 0x69, 0xf2, 0xa4, 0x10, 0x9b, 0x4c, 0xb1, 0x64, 0x42, 0xe6, 0x90, 0x5f, 0x14, 0x09,
	0x13, 0x32, 0x07, 0x26, 0xda, 0x86, 0x5c, 0x44, 0xa2, 0x6d, 0x90, 0x8f, 0xb0, 0x50, 0xbb, 0xed,
	0xd3, 0x4b, 0xd7, 0xd2, 0xe4, 0x90, 0xfc, 0xc7, 0xa1, 0xfa, 0x91, 0xc2, 0x23, 0x42, 0x6a, 0xa9,
	0x71, 0xa3, 0x89, 0x50, 0xa3, 0xe6, 0x9d, 0xa3, 0xdf, 0x9f, 0x5d, 0x8e, 0x7e, 0x2f, 0xdf, 0x20,
	0x7b, 0xa5, 0x66, 0x6b, 0xc3, 0x9f, 0x57, 0xf0, 0x0e, 0xb2, 0x46, 0x07, 0x5d, 0xbf, 0x77, 0x2d,
	0xe5, 0xa2, 0x10, 0x9b, 0xeb, 0xea, 0xa6, 0xe4, 0x36, 0xca, 0x73, 0x7e, 0xb5, 0xe0, 0x01, 0x56,
	0x55, 0x0d, 0x29, 0x53, 0xbc, 0x85, 0x79, 0xbc, 0x07, 0xb3, 0xd3, 0xac, 0xeb, 0xd7, 0xcb, 0x28,
	0xbf, 0x03, 0xc8, 0x19, 0xde, 0xc3, 0xd5, 0xe9, 0x0d, 0xb8, 0x8a, 0xa7, 0x53, 0xb1, 0xeb, 0xdf,
	0xc4, 0xcb, 0x99, 0x99, 0x8f, 0xdf, 0xf1, 0xf0, 0x15, 0x00, 0x00, 0xff, 0xff, 0x10, 0x61, 0x11,
	0xd9, 0xa1, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RBACClient is the client API for RBAC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RBACClient interface {
	ModAct(ctx context.Context, in *Rep, opts ...grpc.CallOption) (*ModActRes, error)
	Enforce(ctx context.Context, in *EnforceRep, opts ...grpc.CallOption) (*EnforceRes, error)
}

type rBACClient struct {
	cc *grpc.ClientConn
}

func NewRBACClient(cc *grpc.ClientConn) RBACClient {
	return &rBACClient{cc}
}

func (c *rBACClient) ModAct(ctx context.Context, in *Rep, opts ...grpc.CallOption) (*ModActRes, error) {
	out := new(ModActRes)
	err := c.cc.Invoke(ctx, "/test.RBAC/ModAct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rBACClient) Enforce(ctx context.Context, in *EnforceRep, opts ...grpc.CallOption) (*EnforceRes, error) {
	out := new(EnforceRes)
	err := c.cc.Invoke(ctx, "/test.RBAC/Enforce", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RBACServer is the server API for RBAC service.
type RBACServer interface {
	ModAct(context.Context, *Rep) (*ModActRes, error)
	Enforce(context.Context, *EnforceRep) (*EnforceRes, error)
}

func RegisterRBACServer(s *grpc.Server, srv RBACServer) {
	s.RegisterService(&_RBAC_serviceDesc, srv)
}

func _RBAC_ModAct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Rep)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RBACServer).ModAct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.RBAC/ModAct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RBACServer).ModAct(ctx, req.(*Rep))
	}
	return interceptor(ctx, in, info, handler)
}

func _RBAC_Enforce_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnforceRep)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RBACServer).Enforce(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.RBAC/Enforce",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RBACServer).Enforce(ctx, req.(*EnforceRep))
	}
	return interceptor(ctx, in, info, handler)
}

var _RBAC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "test.RBAC",
	HandlerType: (*RBACServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ModAct",
			Handler:    _RBAC_ModAct_Handler,
		},
		{
			MethodName: "Enforce",
			Handler:    _RBAC_Enforce_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "role_rbac.proto",
}
