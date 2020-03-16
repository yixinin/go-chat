// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account.proto

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

type SignUpReq struct {
	Header               *ReqHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Username             string     `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password             string     `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	ValidCode            string     `protobuf:"bytes,4,opt,name=validCode,proto3" json:"validCode,omitempty"`
	ValidId              string     `protobuf:"bytes,5,opt,name=validId,proto3" json:"validId,omitempty"`
	DeviceCode           string     `protobuf:"bytes,6,opt,name=deviceCode,proto3" json:"deviceCode,omitempty"`
	DeviceType           int32      `protobuf:"varint,7,opt,name=deviceType,proto3" json:"deviceType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SignUpReq) Reset()         { *m = SignUpReq{} }
func (m *SignUpReq) String() string { return proto.CompactTextString(m) }
func (*SignUpReq) ProtoMessage()    {}
func (*SignUpReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{0}
}

func (m *SignUpReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignUpReq.Unmarshal(m, b)
}
func (m *SignUpReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignUpReq.Marshal(b, m, deterministic)
}
func (m *SignUpReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignUpReq.Merge(m, src)
}
func (m *SignUpReq) XXX_Size() int {
	return xxx_messageInfo_SignUpReq.Size(m)
}
func (m *SignUpReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SignUpReq.DiscardUnknown(m)
}

var xxx_messageInfo_SignUpReq proto.InternalMessageInfo

func (m *SignUpReq) GetHeader() *ReqHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SignUpReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *SignUpReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *SignUpReq) GetValidCode() string {
	if m != nil {
		return m.ValidCode
	}
	return ""
}

func (m *SignUpReq) GetValidId() string {
	if m != nil {
		return m.ValidId
	}
	return ""
}

func (m *SignUpReq) GetDeviceCode() string {
	if m != nil {
		return m.DeviceCode
	}
	return ""
}

func (m *SignUpReq) GetDeviceType() int32 {
	if m != nil {
		return m.DeviceType
	}
	return 0
}

type SignUpAck struct {
	Header               *AckHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Token                string     `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SignUpAck) Reset()         { *m = SignUpAck{} }
func (m *SignUpAck) String() string { return proto.CompactTextString(m) }
func (*SignUpAck) ProtoMessage()    {}
func (*SignUpAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{1}
}

func (m *SignUpAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignUpAck.Unmarshal(m, b)
}
func (m *SignUpAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignUpAck.Marshal(b, m, deterministic)
}
func (m *SignUpAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignUpAck.Merge(m, src)
}
func (m *SignUpAck) XXX_Size() int {
	return xxx_messageInfo_SignUpAck.Size(m)
}
func (m *SignUpAck) XXX_DiscardUnknown() {
	xxx_messageInfo_SignUpAck.DiscardUnknown(m)
}

var xxx_messageInfo_SignUpAck proto.InternalMessageInfo

func (m *SignUpAck) GetHeader() *AckHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SignUpAck) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type SignInReq struct {
	Header               *ReqHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Username             string     `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password             string     `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	DeviceType           int32      `protobuf:"varint,7,opt,name=deviceType,proto3" json:"deviceType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SignInReq) Reset()         { *m = SignInReq{} }
func (m *SignInReq) String() string { return proto.CompactTextString(m) }
func (*SignInReq) ProtoMessage()    {}
func (*SignInReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{2}
}

func (m *SignInReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignInReq.Unmarshal(m, b)
}
func (m *SignInReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignInReq.Marshal(b, m, deterministic)
}
func (m *SignInReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignInReq.Merge(m, src)
}
func (m *SignInReq) XXX_Size() int {
	return xxx_messageInfo_SignInReq.Size(m)
}
func (m *SignInReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SignInReq.DiscardUnknown(m)
}

var xxx_messageInfo_SignInReq proto.InternalMessageInfo

func (m *SignInReq) GetHeader() *ReqHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SignInReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *SignInReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *SignInReq) GetDeviceType() int32 {
	if m != nil {
		return m.DeviceType
	}
	return 0
}

type SignInAck struct {
	Header               *AckHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Token                string     `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SignInAck) Reset()         { *m = SignInAck{} }
func (m *SignInAck) String() string { return proto.CompactTextString(m) }
func (*SignInAck) ProtoMessage()    {}
func (*SignInAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{3}
}

func (m *SignInAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignInAck.Unmarshal(m, b)
}
func (m *SignInAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignInAck.Marshal(b, m, deterministic)
}
func (m *SignInAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignInAck.Merge(m, src)
}
func (m *SignInAck) XXX_Size() int {
	return xxx_messageInfo_SignInAck.Size(m)
}
func (m *SignInAck) XXX_DiscardUnknown() {
	xxx_messageInfo_SignInAck.DiscardUnknown(m)
}

var xxx_messageInfo_SignInAck proto.InternalMessageInfo

func (m *SignInAck) GetHeader() *AckHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SignInAck) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type SignOutReq struct {
	Header               *ReqHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SignOutReq) Reset()         { *m = SignOutReq{} }
func (m *SignOutReq) String() string { return proto.CompactTextString(m) }
func (*SignOutReq) ProtoMessage()    {}
func (*SignOutReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{4}
}

func (m *SignOutReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignOutReq.Unmarshal(m, b)
}
func (m *SignOutReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignOutReq.Marshal(b, m, deterministic)
}
func (m *SignOutReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignOutReq.Merge(m, src)
}
func (m *SignOutReq) XXX_Size() int {
	return xxx_messageInfo_SignOutReq.Size(m)
}
func (m *SignOutReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SignOutReq.DiscardUnknown(m)
}

var xxx_messageInfo_SignOutReq proto.InternalMessageInfo

func (m *SignOutReq) GetHeader() *ReqHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type SignOutAck struct {
	Header               *AckHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SignOutAck) Reset()         { *m = SignOutAck{} }
func (m *SignOutAck) String() string { return proto.CompactTextString(m) }
func (*SignOutAck) ProtoMessage()    {}
func (*SignOutAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{5}
}

func (m *SignOutAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignOutAck.Unmarshal(m, b)
}
func (m *SignOutAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignOutAck.Marshal(b, m, deterministic)
}
func (m *SignOutAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignOutAck.Merge(m, src)
}
func (m *SignOutAck) XXX_Size() int {
	return xxx_messageInfo_SignOutAck.Size(m)
}
func (m *SignOutAck) XXX_DiscardUnknown() {
	xxx_messageInfo_SignOutAck.DiscardUnknown(m)
}

var xxx_messageInfo_SignOutAck proto.InternalMessageInfo

func (m *SignOutAck) GetHeader() *AckHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type DeleteReq struct {
	Header               *ReqHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Password             string     `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *DeleteReq) Reset()         { *m = DeleteReq{} }
func (m *DeleteReq) String() string { return proto.CompactTextString(m) }
func (*DeleteReq) ProtoMessage()    {}
func (*DeleteReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{6}
}

func (m *DeleteReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteReq.Unmarshal(m, b)
}
func (m *DeleteReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteReq.Marshal(b, m, deterministic)
}
func (m *DeleteReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteReq.Merge(m, src)
}
func (m *DeleteReq) XXX_Size() int {
	return xxx_messageInfo_DeleteReq.Size(m)
}
func (m *DeleteReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteReq proto.InternalMessageInfo

func (m *DeleteReq) GetHeader() *ReqHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DeleteReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type DeleteAck struct {
	Header               *AckHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *DeleteAck) Reset()         { *m = DeleteAck{} }
func (m *DeleteAck) String() string { return proto.CompactTextString(m) }
func (*DeleteAck) ProtoMessage()    {}
func (*DeleteAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{7}
}

func (m *DeleteAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteAck.Unmarshal(m, b)
}
func (m *DeleteAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteAck.Marshal(b, m, deterministic)
}
func (m *DeleteAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteAck.Merge(m, src)
}
func (m *DeleteAck) XXX_Size() int {
	return xxx_messageInfo_DeleteAck.Size(m)
}
func (m *DeleteAck) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteAck.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteAck proto.InternalMessageInfo

func (m *DeleteAck) GetHeader() *AckHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type ChangePasswordReq struct {
	Header               *ReqHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	OldPwd               string     `protobuf:"bytes,2,opt,name=oldPwd,proto3" json:"oldPwd,omitempty"`
	NewPwd               string     `protobuf:"bytes,3,opt,name=newPwd,proto3" json:"newPwd,omitempty"`
	ValidCode            string     `protobuf:"bytes,4,opt,name=validCode,proto3" json:"validCode,omitempty"`
	ValidId              string     `protobuf:"bytes,5,opt,name=validId,proto3" json:"validId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ChangePasswordReq) Reset()         { *m = ChangePasswordReq{} }
func (m *ChangePasswordReq) String() string { return proto.CompactTextString(m) }
func (*ChangePasswordReq) ProtoMessage()    {}
func (*ChangePasswordReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{8}
}

func (m *ChangePasswordReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangePasswordReq.Unmarshal(m, b)
}
func (m *ChangePasswordReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangePasswordReq.Marshal(b, m, deterministic)
}
func (m *ChangePasswordReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangePasswordReq.Merge(m, src)
}
func (m *ChangePasswordReq) XXX_Size() int {
	return xxx_messageInfo_ChangePasswordReq.Size(m)
}
func (m *ChangePasswordReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangePasswordReq.DiscardUnknown(m)
}

var xxx_messageInfo_ChangePasswordReq proto.InternalMessageInfo

func (m *ChangePasswordReq) GetHeader() *ReqHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ChangePasswordReq) GetOldPwd() string {
	if m != nil {
		return m.OldPwd
	}
	return ""
}

func (m *ChangePasswordReq) GetNewPwd() string {
	if m != nil {
		return m.NewPwd
	}
	return ""
}

func (m *ChangePasswordReq) GetValidCode() string {
	if m != nil {
		return m.ValidCode
	}
	return ""
}

func (m *ChangePasswordReq) GetValidId() string {
	if m != nil {
		return m.ValidId
	}
	return ""
}

type ChangePasswordAck struct {
	Header               *AckHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ChangePasswordAck) Reset()         { *m = ChangePasswordAck{} }
func (m *ChangePasswordAck) String() string { return proto.CompactTextString(m) }
func (*ChangePasswordAck) ProtoMessage()    {}
func (*ChangePasswordAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{9}
}

func (m *ChangePasswordAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangePasswordAck.Unmarshal(m, b)
}
func (m *ChangePasswordAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangePasswordAck.Marshal(b, m, deterministic)
}
func (m *ChangePasswordAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangePasswordAck.Merge(m, src)
}
func (m *ChangePasswordAck) XXX_Size() int {
	return xxx_messageInfo_ChangePasswordAck.Size(m)
}
func (m *ChangePasswordAck) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangePasswordAck.DiscardUnknown(m)
}

var xxx_messageInfo_ChangePasswordAck proto.InternalMessageInfo

func (m *ChangePasswordAck) GetHeader() *AckHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type ResetPasswordReq struct {
	Header               *AckHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Username             string     `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password             string     `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	ValidCode            string     `protobuf:"bytes,4,opt,name=validCode,proto3" json:"validCode,omitempty"`
	ValidId              string     `protobuf:"bytes,5,opt,name=validId,proto3" json:"validId,omitempty"`
	DeviceType           int32      `protobuf:"varint,7,opt,name=deviceType,proto3" json:"deviceType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ResetPasswordReq) Reset()         { *m = ResetPasswordReq{} }
func (m *ResetPasswordReq) String() string { return proto.CompactTextString(m) }
func (*ResetPasswordReq) ProtoMessage()    {}
func (*ResetPasswordReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{10}
}

func (m *ResetPasswordReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResetPasswordReq.Unmarshal(m, b)
}
func (m *ResetPasswordReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResetPasswordReq.Marshal(b, m, deterministic)
}
func (m *ResetPasswordReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetPasswordReq.Merge(m, src)
}
func (m *ResetPasswordReq) XXX_Size() int {
	return xxx_messageInfo_ResetPasswordReq.Size(m)
}
func (m *ResetPasswordReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetPasswordReq.DiscardUnknown(m)
}

var xxx_messageInfo_ResetPasswordReq proto.InternalMessageInfo

func (m *ResetPasswordReq) GetHeader() *AckHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ResetPasswordReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ResetPasswordReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *ResetPasswordReq) GetValidCode() string {
	if m != nil {
		return m.ValidCode
	}
	return ""
}

func (m *ResetPasswordReq) GetValidId() string {
	if m != nil {
		return m.ValidId
	}
	return ""
}

func (m *ResetPasswordReq) GetDeviceType() int32 {
	if m != nil {
		return m.DeviceType
	}
	return 0
}

type ResetPasswordAck struct {
	Header               *AckHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ResetPasswordAck) Reset()         { *m = ResetPasswordAck{} }
func (m *ResetPasswordAck) String() string { return proto.CompactTextString(m) }
func (*ResetPasswordAck) ProtoMessage()    {}
func (*ResetPasswordAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{11}
}

func (m *ResetPasswordAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResetPasswordAck.Unmarshal(m, b)
}
func (m *ResetPasswordAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResetPasswordAck.Marshal(b, m, deterministic)
}
func (m *ResetPasswordAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetPasswordAck.Merge(m, src)
}
func (m *ResetPasswordAck) XXX_Size() int {
	return xxx_messageInfo_ResetPasswordAck.Size(m)
}
func (m *ResetPasswordAck) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetPasswordAck.DiscardUnknown(m)
}

var xxx_messageInfo_ResetPasswordAck proto.InternalMessageInfo

func (m *ResetPasswordAck) GetHeader() *AckHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func init() {
	proto.RegisterType((*SignUpReq)(nil), "protocol.SignUpReq")
	proto.RegisterType((*SignUpAck)(nil), "protocol.SignUpAck")
	proto.RegisterType((*SignInReq)(nil), "protocol.SignInReq")
	proto.RegisterType((*SignInAck)(nil), "protocol.SignInAck")
	proto.RegisterType((*SignOutReq)(nil), "protocol.SignOutReq")
	proto.RegisterType((*SignOutAck)(nil), "protocol.SignOutAck")
	proto.RegisterType((*DeleteReq)(nil), "protocol.DeleteReq")
	proto.RegisterType((*DeleteAck)(nil), "protocol.DeleteAck")
	proto.RegisterType((*ChangePasswordReq)(nil), "protocol.ChangePasswordReq")
	proto.RegisterType((*ChangePasswordAck)(nil), "protocol.ChangePasswordAck")
	proto.RegisterType((*ResetPasswordReq)(nil), "protocol.ResetPasswordReq")
	proto.RegisterType((*ResetPasswordAck)(nil), "protocol.ResetPasswordAck")
}

func init() { proto.RegisterFile("account.proto", fileDescriptor_8e28828dcb8d24f0) }

var fileDescriptor_8e28828dcb8d24f0 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x54, 0x41, 0x4e, 0xc2, 0x40,
	0x14, 0xcd, 0xa0, 0x14, 0xfa, 0xd5, 0x44, 0xab, 0x31, 0x13, 0x62, 0x0c, 0xe9, 0x8a, 0xc4, 0x84,
	0x85, 0x6e, 0x74, 0xa5, 0x04, 0x17, 0xb2, 0x51, 0x52, 0xf1, 0x00, 0xb5, 0xf3, 0x03, 0x84, 0x3a,
	0x53, 0xda, 0x02, 0xf1, 0x16, 0x1e, 0xc3, 0xdb, 0x78, 0x0d, 0x8f, 0x61, 0x3a, 0x9d, 0x01, 0xac,
	0xd1, 0x30, 0xc4, 0xe8, 0xaa, 0x79, 0xef, 0xf5, 0xff, 0xbc, 0xf7, 0xda, 0x19, 0xd8, 0xf1, 0x83,
	0x40, 0x4c, 0x78, 0xda, 0x8c, 0x62, 0x91, 0x0a, 0xa7, 0x2a, 0x1f, 0x81, 0x08, 0x6b, 0xdb, 0x03,
	0xf4, 0x19, 0xc6, 0x39, 0xef, 0xbe, 0x13, 0xb0, 0xef, 0x87, 0x7d, 0xfe, 0x10, 0x79, 0x38, 0x76,
	0x4e, 0xc0, 0xca, 0x55, 0x4a, 0xea, 0xa4, 0xb1, 0x75, 0xba, 0xdf, 0xd4, 0x63, 0x4d, 0x0f, 0xc7,
	0x37, 0x52, 0xf2, 0xd4, 0x2b, 0x4e, 0x0d, 0xaa, 0x93, 0x04, 0x63, 0xee, 0x3f, 0x21, 0x2d, 0xd5,
	0x49, 0xc3, 0xf6, 0xe6, 0x38, 0xd3, 0x22, 0x3f, 0x49, 0x66, 0x22, 0x66, 0x74, 0x23, 0xd7, 0x34,
	0x76, 0x8e, 0xc0, 0x9e, 0xfa, 0xe1, 0x90, 0xb5, 0x05, 0x43, 0xba, 0x29, 0xc5, 0x05, 0xe1, 0x50,
	0xa8, 0x48, 0xd0, 0x61, 0xb4, 0x2c, 0x35, 0x0d, 0x9d, 0x63, 0x00, 0x86, 0xd3, 0x61, 0x80, 0x72,
	0xd0, 0x92, 0xe2, 0x12, 0xb3, 0xd0, 0x7b, 0xcf, 0x11, 0xd2, 0x4a, 0x9d, 0x34, 0xca, 0xde, 0x12,
	0xe3, 0xde, 0xea, 0xa4, 0xad, 0x60, 0xf4, 0x53, 0xd2, 0x56, 0x30, 0x2a, 0x24, 0x3d, 0x80, 0x72,
	0x2a, 0x46, 0xc8, 0x55, 0xcc, 0x1c, 0xb8, 0x2f, 0xaa, 0xba, 0x0e, 0xff, 0xb3, 0xea, 0x56, 0x8c,
	0xd8, 0xe1, 0xbf, 0x14, 0xf1, 0x02, 0x20, 0xdb, 0x77, 0x37, 0x49, 0x4d, 0x23, 0x2e, 0x8d, 0x9a,
	0x7a, 0x71, 0x7b, 0x60, 0x5f, 0x63, 0x88, 0x29, 0xae, 0xd3, 0xeb, 0xbc, 0xbb, 0xd2, 0xe7, 0xee,
	0xdc, 0x73, 0xbd, 0xd5, 0xd8, 0xcf, 0x2b, 0x81, 0xbd, 0xf6, 0xc0, 0xe7, 0x7d, 0xec, 0xaa, 0x65,
	0xc6, 0xc6, 0x0e, 0xc1, 0x12, 0x21, 0xeb, 0xce, 0xb4, 0x2d, 0x85, 0x32, 0x9e, 0xe3, 0x2c, 0xe3,
	0xf3, 0x4f, 0xad, 0xd0, 0xba, 0x67, 0xc4, 0xbd, 0x2a, 0x3a, 0x35, 0x0e, 0xfb, 0x46, 0x60, 0xd7,
	0xc3, 0x04, 0xd3, 0x15, 0xb3, 0x7e, 0xfd, 0x95, 0xfe, 0xed, 0x5e, 0xf8, 0xe6, 0x50, 0x5c, 0x16,
	0x02, 0x99, 0x56, 0xf2, 0x68, 0x49, 0xed, 0xec, 0x23, 0x00, 0x00, 0xff, 0xff, 0x66, 0x17, 0x96,
	0x04, 0x53, 0x05, 0x00, 0x00,
}
