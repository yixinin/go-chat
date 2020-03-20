// Code generated by protoc-gen-go. DO NOT EDIT.
// source: echo.proto

package protocol

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

type EchoReq struct {
	Header               *ReqHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Message              string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *EchoReq) Reset()         { *m = EchoReq{} }
func (m *EchoReq) String() string { return proto.CompactTextString(m) }
func (*EchoReq) ProtoMessage()    {}
func (*EchoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_08134aea513e0001, []int{0}
}

func (m *EchoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoReq.Unmarshal(m, b)
}
func (m *EchoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoReq.Marshal(b, m, deterministic)
}
func (m *EchoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoReq.Merge(m, src)
}
func (m *EchoReq) XXX_Size() int {
	return xxx_messageInfo_EchoReq.Size(m)
}
func (m *EchoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoReq.DiscardUnknown(m)
}

var xxx_messageInfo_EchoReq proto.InternalMessageInfo

func (m *EchoReq) GetHeader() *ReqHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *EchoReq) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type EchoAck struct {
	Header               *AckHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Message              string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *EchoAck) Reset()         { *m = EchoAck{} }
func (m *EchoAck) String() string { return proto.CompactTextString(m) }
func (*EchoAck) ProtoMessage()    {}
func (*EchoAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_08134aea513e0001, []int{1}
}

func (m *EchoAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoAck.Unmarshal(m, b)
}
func (m *EchoAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoAck.Marshal(b, m, deterministic)
}
func (m *EchoAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoAck.Merge(m, src)
}
func (m *EchoAck) XXX_Size() int {
	return xxx_messageInfo_EchoAck.Size(m)
}
func (m *EchoAck) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoAck.DiscardUnknown(m)
}

var xxx_messageInfo_EchoAck proto.InternalMessageInfo

func (m *EchoAck) GetHeader() *AckHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *EchoAck) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*EchoReq)(nil), "protocol.EchoReq")
	proto.RegisterType((*EchoAck)(nil), "protocol.EchoAck")
}

func init() { proto.RegisterFile("echo.proto", fileDescriptor_08134aea513e0001) }

var fileDescriptor_08134aea513e0001 = []byte{
	// 128 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4d, 0xce, 0xc8,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0xc9, 0xf9, 0x39, 0x52, 0x3c, 0x19,
	0xa9, 0x89, 0x29, 0xa9, 0x45, 0x10, 0x71, 0xa5, 0x00, 0x2e, 0x76, 0xd7, 0xe4, 0x8c, 0xfc, 0xa0,
	0xd4, 0x42, 0x21, 0x6d, 0x2e, 0x36, 0x88, 0x94, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0xb0,
	0x1e, 0x4c, 0x8f, 0x5e, 0x50, 0x6a, 0xa1, 0x07, 0x58, 0x2a, 0x08, 0xaa, 0x44, 0x48, 0x82, 0x8b,
	0x3d, 0x37, 0xb5, 0xb8, 0x38, 0x31, 0x3d, 0x55, 0x82, 0x49, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc6,
	0x85, 0x99, 0xe8, 0x98, 0x9c, 0x8d, 0xcf, 0x44, 0xc7, 0xe4, 0x6c, 0x62, 0x4d, 0x4c, 0x62, 0x03,
	0xeb, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x87, 0x7d, 0x59, 0xd9, 0xd0, 0x00, 0x00, 0x00,
}
