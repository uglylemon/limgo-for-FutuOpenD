// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Trd_Notify/Trd_Notify.proto

package Trd_Notify

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "limgo/futupb/Common"
import Trd_Common "limgo/futupb/Trd_Common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type S2C struct {
	Header               *Trd_Common.TrdHeader `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`
	Type                 *int32                `protobuf:"varint,2,req,name=type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *S2C) Reset()         { *m = S2C{} }
func (m *S2C) String() string { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()    {}
func (*S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_Trd_Notify_f3da7110feb3d391, []int{0}
}
func (m *S2C) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2C.Unmarshal(m, b)
}
func (m *S2C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2C.Marshal(b, m, deterministic)
}
func (dst *S2C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2C.Merge(dst, src)
}
func (m *S2C) XXX_Size() int {
	return xxx_messageInfo_S2C.Size(m)
}
func (m *S2C) XXX_DiscardUnknown() {
	xxx_messageInfo_S2C.DiscardUnknown(m)
}

var xxx_messageInfo_S2C proto.InternalMessageInfo

func (m *S2C) GetHeader() *Trd_Common.TrdHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *S2C) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

type Response struct {
	// 以下3个字段每条协议都有，注释说明在InitConnect.proto中
	RetType              *int32   `protobuf:"varint,1,req,name=retType,def=-400" json:"retType,omitempty"`
	RetMsg               *string  `protobuf:"bytes,2,opt,name=retMsg" json:"retMsg,omitempty"`
	ErrCode              *int32   `protobuf:"varint,3,opt,name=errCode" json:"errCode,omitempty"`
	S2C                  *S2C     `protobuf:"bytes,4,opt,name=s2c" json:"s2c,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_Trd_Notify_f3da7110feb3d391, []int{1}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

const Default_Response_RetType int32 = -400

func (m *Response) GetRetType() int32 {
	if m != nil && m.RetType != nil {
		return *m.RetType
	}
	return Default_Response_RetType
}

func (m *Response) GetRetMsg() string {
	if m != nil && m.RetMsg != nil {
		return *m.RetMsg
	}
	return ""
}

func (m *Response) GetErrCode() int32 {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *Response) GetS2C() *S2C {
	if m != nil {
		return m.S2C
	}
	return nil
}

func init() {
	proto.RegisterType((*S2C)(nil), "Trd_Notify.S2C")
	proto.RegisterType((*Response)(nil), "Trd_Notify.Response")
}

func init() {
	proto.RegisterFile("Trd_Notify/Trd_Notify.proto", fileDescriptor_Trd_Notify_f3da7110feb3d391)
}

var fileDescriptor_Trd_Notify_f3da7110feb3d391 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x3b, 0x4f, 0xc3, 0x30,
	0x14, 0x85, 0x65, 0x27, 0x69, 0xe1, 0x76, 0x40, 0xba, 0x08, 0x64, 0x15, 0x09, 0x99, 0x4e, 0x5e,
	0x9a, 0x56, 0x16, 0x13, 0xab, 0x97, 0x2e, 0x30, 0xb8, 0xd9, 0x11, 0xc2, 0x97, 0xc7, 0xd0, 0x3a,
	0xba, 0xf6, 0xd2, 0x89, 0xbf, 0x8e, 0xf2, 0xa8, 0xd2, 0xc9, 0xe7, 0x9c, 0xef, 0xb3, 0x07, 0xc3,
	0x43, 0xc3, 0xe1, 0xfd, 0x2d, 0xe6, 0xdf, 0xaf, 0xd3, 0x66, 0x8a, 0x75, 0xcb, 0x31, 0x47, 0x84,
	0x69, 0x59, 0xde, 0xba, 0x78, 0x38, 0xc4, 0xe3, 0x66, 0x38, 0x06, 0x61, 0xd9, 0xdf, 0x1e, 0xc1,
	0x14, 0x07, 0xb8, 0xda, 0x41, 0xb1, 0xb7, 0x0e, 0xd7, 0x30, 0xfb, 0xa1, 0x8f, 0x40, 0xac, 0x84,
	0x96, 0x66, 0x61, 0xef, 0xea, 0x0b, 0xb3, 0xe1, 0xb0, 0xeb, 0xa1, 0x1f, 0x25, 0x44, 0x28, 0xf3,
	0xa9, 0x25, 0x25, 0xb5, 0x34, 0x95, 0xef, 0xf3, 0xea, 0x0f, 0xae, 0x3c, 0xa5, 0x36, 0x1e, 0x13,
	0xe1, 0x23, 0xcc, 0x99, 0x72, 0xd3, 0x29, 0xdd, 0x7b, 0xd5, 0x4b, 0xb9, 0x7e, 0xde, 0x6e, 0xfd,
	0x79, 0xc4, 0x7b, 0x98, 0x31, 0xe5, 0xd7, 0xf4, 0xad, 0xa4, 0x16, 0xe6, 0xda, 0x8f, 0x0d, 0x15,
	0xcc, 0x89, 0xd9, 0xc5, 0x40, 0xaa, 0xd0, 0xc2, 0x54, 0xfe, 0x5c, 0xf1, 0x09, 0x8a, 0x64, 0x3f,
	0x55, 0xa9, 0x85, 0x59, 0xd8, 0x9b, 0xfa, 0xe2, 0x17, 0xf6, 0xd6, 0xf9, 0x8e, 0xfd, 0x07, 0x00,
	0x00, 0xff, 0xff, 0xe5, 0x41, 0x1a, 0xef, 0x26, 0x01, 0x00, 0x00,
}