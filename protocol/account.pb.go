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
	Nickname             string     `protobuf:"bytes,8,opt,name=nickname,proto3" json:"nickname,omitempty"`
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

func (m *SignUpReq) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

type SignUpAck struct {
	Header               *AckHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Token                string     `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	DeviceCode           string     `protobuf:"bytes,3,opt,name=deviceCode,proto3" json:"deviceCode,omitempty"`
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

func (m *SignUpAck) GetDeviceCode() string {
	if m != nil {
		return m.DeviceCode
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

type SignOffReq struct {
	Header               *ReqHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Password             string     `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SignOffReq) Reset()         { *m = SignOffReq{} }
func (m *SignOffReq) String() string { return proto.CompactTextString(m) }
func (*SignOffReq) ProtoMessage()    {}
func (*SignOffReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{6}
}

func (m *SignOffReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignOffReq.Unmarshal(m, b)
}
func (m *SignOffReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignOffReq.Marshal(b, m, deterministic)
}
func (m *SignOffReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignOffReq.Merge(m, src)
}
func (m *SignOffReq) XXX_Size() int {
	return xxx_messageInfo_SignOffReq.Size(m)
}
func (m *SignOffReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SignOffReq.DiscardUnknown(m)
}

var xxx_messageInfo_SignOffReq proto.InternalMessageInfo

func (m *SignOffReq) GetHeader() *ReqHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SignOffReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type SignOffAck struct {
	Header               *AckHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SignOffAck) Reset()         { *m = SignOffAck{} }
func (m *SignOffAck) String() string { return proto.CompactTextString(m) }
func (*SignOffAck) ProtoMessage()    {}
func (*SignOffAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{7}
}

func (m *SignOffAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignOffAck.Unmarshal(m, b)
}
func (m *SignOffAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignOffAck.Marshal(b, m, deterministic)
}
func (m *SignOffAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignOffAck.Merge(m, src)
}
func (m *SignOffAck) XXX_Size() int {
	return xxx_messageInfo_SignOffAck.Size(m)
}
func (m *SignOffAck) XXX_DiscardUnknown() {
	xxx_messageInfo_SignOffAck.DiscardUnknown(m)
}

var xxx_messageInfo_SignOffAck proto.InternalMessageInfo

func (m *SignOffAck) GetHeader() *AckHeader {
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
	Header               *ReqHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
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

func (m *ResetPasswordReq) GetHeader() *ReqHeader {
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

type GetUserInfoReq struct {
	Header               *ReqHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetUserInfoReq) Reset()         { *m = GetUserInfoReq{} }
func (m *GetUserInfoReq) String() string { return proto.CompactTextString(m) }
func (*GetUserInfoReq) ProtoMessage()    {}
func (*GetUserInfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{12}
}

func (m *GetUserInfoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserInfoReq.Unmarshal(m, b)
}
func (m *GetUserInfoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserInfoReq.Marshal(b, m, deterministic)
}
func (m *GetUserInfoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserInfoReq.Merge(m, src)
}
func (m *GetUserInfoReq) XXX_Size() int {
	return xxx_messageInfo_GetUserInfoReq.Size(m)
}
func (m *GetUserInfoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserInfoReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserInfoReq proto.InternalMessageInfo

func (m *GetUserInfoReq) GetHeader() *ReqHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type GetUserInfoAck struct {
	Header               *AckHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Nickname             string     `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Avatar               string     `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetUserInfoAck) Reset()         { *m = GetUserInfoAck{} }
func (m *GetUserInfoAck) String() string { return proto.CompactTextString(m) }
func (*GetUserInfoAck) ProtoMessage()    {}
func (*GetUserInfoAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{13}
}

func (m *GetUserInfoAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserInfoAck.Unmarshal(m, b)
}
func (m *GetUserInfoAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserInfoAck.Marshal(b, m, deterministic)
}
func (m *GetUserInfoAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserInfoAck.Merge(m, src)
}
func (m *GetUserInfoAck) XXX_Size() int {
	return xxx_messageInfo_GetUserInfoAck.Size(m)
}
func (m *GetUserInfoAck) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserInfoAck.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserInfoAck proto.InternalMessageInfo

func (m *GetUserInfoAck) GetHeader() *AckHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *GetUserInfoAck) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *GetUserInfoAck) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func init() {
	proto.RegisterType((*SignUpReq)(nil), "protocol.SignUpReq")
	proto.RegisterType((*SignUpAck)(nil), "protocol.SignUpAck")
	proto.RegisterType((*SignInReq)(nil), "protocol.SignInReq")
	proto.RegisterType((*SignInAck)(nil), "protocol.SignInAck")
	proto.RegisterType((*SignOutReq)(nil), "protocol.SignOutReq")
	proto.RegisterType((*SignOutAck)(nil), "protocol.SignOutAck")
	proto.RegisterType((*SignOffReq)(nil), "protocol.SignOffReq")
	proto.RegisterType((*SignOffAck)(nil), "protocol.SignOffAck")
	proto.RegisterType((*ChangePasswordReq)(nil), "protocol.ChangePasswordReq")
	proto.RegisterType((*ChangePasswordAck)(nil), "protocol.ChangePasswordAck")
	proto.RegisterType((*ResetPasswordReq)(nil), "protocol.ResetPasswordReq")
	proto.RegisterType((*ResetPasswordAck)(nil), "protocol.ResetPasswordAck")
	proto.RegisterType((*GetUserInfoReq)(nil), "protocol.GetUserInfoReq")
	proto.RegisterType((*GetUserInfoAck)(nil), "protocol.GetUserInfoAck")
}

func init() { proto.RegisterFile("account.proto", fileDescriptor_8e28828dcb8d24f0) }

var fileDescriptor_8e28828dcb8d24f0 = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x54, 0xc1, 0x4e, 0xea, 0x40,
	0x14, 0x4d, 0xcb, 0xa3, 0xc0, 0x7d, 0x6a, 0xb4, 0x1a, 0xd3, 0x10, 0x63, 0x48, 0x57, 0x24, 0x26,
	0x2c, 0x74, 0xe5, 0xc2, 0x28, 0x61, 0xa1, 0x6c, 0x94, 0x54, 0xf9, 0x80, 0xb1, 0xbd, 0x85, 0xa6,
	0x38, 0x53, 0xda, 0x02, 0x71, 0xef, 0x07, 0xf8, 0x19, 0xfe, 0x8d, 0xbf, 0x64, 0x3a, 0xd3, 0x81,
	0xd2, 0x44, 0xcd, 0x10, 0xa3, 0xab, 0xe6, 0x9c, 0xd3, 0x3b, 0x3d, 0xf7, 0x74, 0xee, 0x85, 0x6d,
	0xe2, 0xba, 0x6c, 0x46, 0xd3, 0x4e, 0x14, 0xb3, 0x94, 0x99, 0x75, 0xfe, 0x70, 0xd9, 0xa4, 0xb9,
	0x35, 0x46, 0xe2, 0x61, 0x2c, 0x78, 0xfb, 0x45, 0x87, 0xc6, 0x7d, 0x30, 0xa2, 0xc3, 0xc8, 0xc1,
	0xa9, 0x79, 0x02, 0x86, 0x50, 0x2d, 0xad, 0xa5, 0xb5, 0xff, 0x9f, 0xee, 0x77, 0x64, 0x59, 0xc7,
	0xc1, 0xe9, 0x0d, 0x97, 0x9c, 0xfc, 0x15, 0xb3, 0x09, 0xf5, 0x59, 0x82, 0x31, 0x25, 0x4f, 0x68,
	0xe9, 0x2d, 0xad, 0xdd, 0x70, 0x96, 0x38, 0xd3, 0x22, 0x92, 0x24, 0x0b, 0x16, 0x7b, 0x56, 0x45,
	0x68, 0x12, 0x9b, 0x47, 0xd0, 0x98, 0x93, 0x49, 0xe0, 0xf5, 0x98, 0x87, 0xd6, 0x3f, 0x2e, 0xae,
	0x08, 0xd3, 0x82, 0x1a, 0x07, 0x7d, 0xcf, 0xaa, 0x72, 0x4d, 0x42, 0xf3, 0x18, 0xc0, 0xc3, 0x79,
	0xe0, 0x22, 0x2f, 0x34, 0xb8, 0x58, 0x60, 0x56, 0xfa, 0xc3, 0x73, 0x84, 0x56, 0xad, 0xa5, 0xb5,
	0xab, 0x4e, 0x81, 0xc9, 0x3c, 0xd1, 0xc0, 0x0d, 0xb9, 0xdf, 0xba, 0xf0, 0x24, 0xb1, 0x4d, 0x65,
	0x0a, 0x5d, 0x37, 0xfc, 0x2a, 0x85, 0xae, 0x1b, 0x96, 0x52, 0x38, 0x80, 0x6a, 0xca, 0x42, 0xa4,
	0x79, 0x04, 0x02, 0x94, 0xbc, 0x56, 0xca, 0x5e, 0xed, 0x57, 0x4d, 0x7c, 0xb0, 0x4f, 0x7f, 0x2d,
	0xf6, 0x6f, 0xe2, 0xb1, 0x6f, 0xa5, 0xa3, 0x9f, 0x89, 0xc0, 0x3e, 0x07, 0xc8, 0xce, 0xbb, 0x9b,
	0xa5, 0xaa, 0x2d, 0x16, 0x4a, 0x55, 0xbd, 0xd8, 0xc3, 0xbc, 0xd4, 0xf7, 0x37, 0x09, 0x76, 0x19,
	0x9e, 0xbe, 0x1e, 0xde, 0xd2, 0x91, 0xef, 0x2b, 0x3b, 0x7a, 0xd3, 0x60, 0xaf, 0x37, 0x26, 0x74,
	0x84, 0x83, 0xfc, 0x34, 0x65, 0x67, 0x87, 0x60, 0xb0, 0x89, 0x37, 0x58, 0x48, 0x5f, 0x39, 0xca,
	0x78, 0x8a, 0x8b, 0x8c, 0x17, 0x3f, 0x3b, 0x47, 0x9b, 0x4e, 0x98, 0x7d, 0x55, 0x76, 0xaa, 0xdc,
	0xec, 0xbb, 0x06, 0xbb, 0x0e, 0x26, 0x98, 0x6e, 0xdc, 0xeb, 0x9f, 0x6d, 0x95, 0x4f, 0xc6, 0xe2,
	0xb2, 0xd4, 0x90, 0x72, 0x24, 0x17, 0xb0, 0x73, 0x8d, 0xe9, 0x30, 0xc1, 0xb8, 0x4f, 0x7d, 0xa6,
	0x3c, 0x0b, 0xd3, 0xb5, 0x72, 0xe5, 0xd9, 0x2c, 0x2e, 0x3d, 0x7d, 0x7d, 0xe9, 0x65, 0xd7, 0x87,
	0xcc, 0x49, 0x4a, 0x62, 0x79, 0x7d, 0x04, 0x7a, 0x34, 0xf8, 0x79, 0x67, 0x1f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xe0, 0xd7, 0xfe, 0xc2, 0x43, 0x06, 0x00, 0x00,
}
