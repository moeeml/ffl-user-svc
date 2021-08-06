// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package user

import (
	context "context"
	fmt "fmt"
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

type UserAddReq struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Avatar               string   `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserAddReq) Reset()         { *m = UserAddReq{} }
func (m *UserAddReq) String() string { return proto.CompactTextString(m) }
func (*UserAddReq) ProtoMessage()    {}
func (*UserAddReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *UserAddReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserAddReq.Unmarshal(m, b)
}
func (m *UserAddReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserAddReq.Marshal(b, m, deterministic)
}
func (m *UserAddReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserAddReq.Merge(m, src)
}
func (m *UserAddReq) XXX_Size() int {
	return xxx_messageInfo_UserAddReq.Size(m)
}
func (m *UserAddReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserAddReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserAddReq proto.InternalMessageInfo

func (m *UserAddReq) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *UserAddReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserAddReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserAddReq) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

type UserAddRsp struct {
	Code                 int64    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Uid                  uint64   `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserAddRsp) Reset()         { *m = UserAddRsp{} }
func (m *UserAddRsp) String() string { return proto.CompactTextString(m) }
func (*UserAddRsp) ProtoMessage()    {}
func (*UserAddRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *UserAddRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserAddRsp.Unmarshal(m, b)
}
func (m *UserAddRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserAddRsp.Marshal(b, m, deterministic)
}
func (m *UserAddRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserAddRsp.Merge(m, src)
}
func (m *UserAddRsp) XXX_Size() int {
	return xxx_messageInfo_UserAddRsp.Size(m)
}
func (m *UserAddRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_UserAddRsp.DiscardUnknown(m)
}

var xxx_messageInfo_UserAddRsp proto.InternalMessageInfo

func (m *UserAddRsp) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UserAddRsp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *UserAddRsp) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type UserUpdateReq struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Avatar               string   `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Status               string   `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserUpdateReq) Reset()         { *m = UserUpdateReq{} }
func (m *UserUpdateReq) String() string { return proto.CompactTextString(m) }
func (*UserUpdateReq) ProtoMessage()    {}
func (*UserUpdateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *UserUpdateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserUpdateReq.Unmarshal(m, b)
}
func (m *UserUpdateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserUpdateReq.Marshal(b, m, deterministic)
}
func (m *UserUpdateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserUpdateReq.Merge(m, src)
}
func (m *UserUpdateReq) XXX_Size() int {
	return xxx_messageInfo_UserUpdateReq.Size(m)
}
func (m *UserUpdateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserUpdateReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserUpdateReq proto.InternalMessageInfo

func (m *UserUpdateReq) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserUpdateReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserUpdateReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserUpdateReq) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *UserUpdateReq) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type UserInfo struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Account              string   `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Avatar               string   `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Status               string   `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (m *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(m, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserInfo) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *UserInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserInfo) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *UserInfo) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type UserDetail struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Account              string   `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Avatar               string   `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Status               string   `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Password             string   `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserDetail) Reset()         { *m = UserDetail{} }
func (m *UserDetail) String() string { return proto.CompactTextString(m) }
func (*UserDetail) ProtoMessage()    {}
func (*UserDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *UserDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserDetail.Unmarshal(m, b)
}
func (m *UserDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserDetail.Marshal(b, m, deterministic)
}
func (m *UserDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserDetail.Merge(m, src)
}
func (m *UserDetail) XXX_Size() int {
	return xxx_messageInfo_UserDetail.Size(m)
}
func (m *UserDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_UserDetail.DiscardUnknown(m)
}

var xxx_messageInfo_UserDetail proto.InternalMessageInfo

func (m *UserDetail) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserDetail) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *UserDetail) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserDetail) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *UserDetail) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *UserDetail) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type UserInfoReq struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Account              string   `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfoReq) Reset()         { *m = UserInfoReq{} }
func (m *UserInfoReq) String() string { return proto.CompactTextString(m) }
func (*UserInfoReq) ProtoMessage()    {}
func (*UserInfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *UserInfoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoReq.Unmarshal(m, b)
}
func (m *UserInfoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoReq.Marshal(b, m, deterministic)
}
func (m *UserInfoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoReq.Merge(m, src)
}
func (m *UserInfoReq) XXX_Size() int {
	return xxx_messageInfo_UserInfoReq.Size(m)
}
func (m *UserInfoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoReq proto.InternalMessageInfo

func (m *UserInfoReq) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserInfoReq) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

type UserInfoRsp struct {
	Code                 int64       `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string      `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	User                 *UserDetail `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UserInfoRsp) Reset()         { *m = UserInfoRsp{} }
func (m *UserInfoRsp) String() string { return proto.CompactTextString(m) }
func (*UserInfoRsp) ProtoMessage()    {}
func (*UserInfoRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{6}
}

func (m *UserInfoRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoRsp.Unmarshal(m, b)
}
func (m *UserInfoRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoRsp.Marshal(b, m, deterministic)
}
func (m *UserInfoRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoRsp.Merge(m, src)
}
func (m *UserInfoRsp) XXX_Size() int {
	return xxx_messageInfo_UserInfoRsp.Size(m)
}
func (m *UserInfoRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoRsp.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoRsp proto.InternalMessageInfo

func (m *UserInfoRsp) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UserInfoRsp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *UserInfoRsp) GetUser() *UserDetail {
	if m != nil {
		return m.User
	}
	return nil
}

type UserListReq struct {
	Id                   []uint64 `protobuf:"varint,1,rep,packed,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserListReq) Reset()         { *m = UserListReq{} }
func (m *UserListReq) String() string { return proto.CompactTextString(m) }
func (*UserListReq) ProtoMessage()    {}
func (*UserListReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{7}
}

func (m *UserListReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListReq.Unmarshal(m, b)
}
func (m *UserListReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListReq.Marshal(b, m, deterministic)
}
func (m *UserListReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListReq.Merge(m, src)
}
func (m *UserListReq) XXX_Size() int {
	return xxx_messageInfo_UserListReq.Size(m)
}
func (m *UserListReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserListReq proto.InternalMessageInfo

func (m *UserListReq) GetId() []uint64 {
	if m != nil {
		return m.Id
	}
	return nil
}

type UserListRsp struct {
	Code                 int64       `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string      `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	User                 []*UserInfo `protobuf:"bytes,3,rep,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UserListRsp) Reset()         { *m = UserListRsp{} }
func (m *UserListRsp) String() string { return proto.CompactTextString(m) }
func (*UserListRsp) ProtoMessage()    {}
func (*UserListRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{8}
}

func (m *UserListRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListRsp.Unmarshal(m, b)
}
func (m *UserListRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListRsp.Marshal(b, m, deterministic)
}
func (m *UserListRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListRsp.Merge(m, src)
}
func (m *UserListRsp) XXX_Size() int {
	return xxx_messageInfo_UserListRsp.Size(m)
}
func (m *UserListRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListRsp.DiscardUnknown(m)
}

var xxx_messageInfo_UserListRsp proto.InternalMessageInfo

func (m *UserListRsp) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UserListRsp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *UserListRsp) GetUser() []*UserInfo {
	if m != nil {
		return m.User
	}
	return nil
}

type Response struct {
	Code                 int64    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{9}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*UserAddReq)(nil), "user.UserAddReq")
	proto.RegisterType((*UserAddRsp)(nil), "user.UserAddRsp")
	proto.RegisterType((*UserUpdateReq)(nil), "user.UserUpdateReq")
	proto.RegisterType((*UserInfo)(nil), "user.UserInfo")
	proto.RegisterType((*UserDetail)(nil), "user.UserDetail")
	proto.RegisterType((*UserInfoReq)(nil), "user.UserInfoReq")
	proto.RegisterType((*UserInfoRsp)(nil), "user.UserInfoRsp")
	proto.RegisterType((*UserListReq)(nil), "user.UserListReq")
	proto.RegisterType((*UserListRsp)(nil), "user.UserListRsp")
	proto.RegisterType((*Response)(nil), "user.Response")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x25, 0x1f, 0xb6, 0x75, 0x8a, 0xa5, 0xae, 0x20, 0x21, 0x20, 0xc8, 0xe2, 0xa1, 0x17, 0x8b,
	0xd6, 0x83, 0x5e, 0x05, 0x2f, 0x42, 0x4f, 0x81, 0x9e, 0x65, 0xed, 0xae, 0x12, 0xb0, 0xc9, 0x36,
	0xb3, 0x51, 0x6f, 0xfe, 0x0b, 0x7f, 0x95, 0x3f, 0x4a, 0x76, 0x37, 0x1f, 0xdb, 0xa8, 0x87, 0x42,
	0xf1, 0x36, 0xf3, 0x96, 0x99, 0xf7, 0xe6, 0x65, 0x26, 0x00, 0x25, 0x8a, 0x62, 0x2a, 0x8b, 0x5c,
	0xe5, 0x24, 0xd4, 0x31, 0xcd, 0x00, 0x16, 0x28, 0x8a, 0x5b, 0xce, 0x13, 0xb1, 0x26, 0x11, 0xf4,
	0xd9, 0x72, 0x99, 0x97, 0x99, 0x8a, 0xbc, 0x53, 0x6f, 0xb2, 0x9f, 0xd4, 0x29, 0x89, 0x61, 0x20,
	0x19, 0xe2, 0x5b, 0x5e, 0xf0, 0xc8, 0x37, 0x4f, 0x4d, 0x4e, 0x08, 0x84, 0x19, 0x5b, 0x89, 0x28,
	0x30, 0xb8, 0x89, 0xc9, 0x31, 0xf4, 0xd8, 0x2b, 0x53, 0xac, 0x88, 0x42, 0x83, 0x56, 0x19, 0x9d,
	0xb7, 0x7c, 0x28, 0x75, 0xe5, 0x32, 0xe7, 0xc2, 0x90, 0x05, 0x89, 0x89, 0xb5, 0x86, 0x95, 0x40,
	0x64, 0xcf, 0xa2, 0x22, 0xaa, 0x53, 0x32, 0x86, 0xa0, 0x4c, 0xb9, 0xa1, 0x09, 0x13, 0x1d, 0xd2,
	0x0f, 0x38, 0xd0, 0xdd, 0x16, 0x92, 0x33, 0x25, 0xf4, 0x00, 0x23, 0xf0, 0x53, 0x6e, 0xda, 0x85,
	0x89, 0x9f, 0xf2, 0x5d, 0xc9, 0xd6, 0x38, 0x2a, 0xa6, 0x4a, 0x8c, 0xf6, 0x2c, 0x6e, 0x33, 0xfa,
	0x0e, 0x03, 0x2d, 0xe0, 0x3e, 0x7b, 0xca, 0x7f, 0x70, 0x3b, 0x66, 0xfa, 0x9b, 0x66, 0xee, 0x82,
	0xf9, 0xd3, 0xb3, 0x4e, 0xde, 0x09, 0xc5, 0xd2, 0x97, 0xff, 0x25, 0xdf, 0xb0, 0xb5, 0xb7, 0x69,
	0x2b, 0xbd, 0x86, 0x61, 0x6d, 0xc9, 0x6f, 0x5f, 0xe4, 0x4f, 0x61, 0x94, 0x39, 0x85, 0x5b, 0xef,
	0xc6, 0x19, 0x98, 0x7d, 0x36, 0x53, 0x0d, 0x67, 0xe3, 0xa9, 0x59, 0xf4, 0xd6, 0x9f, 0xc4, 0x6e,
	0xfb, 0x89, 0xa5, 0x98, 0xa7, 0xa8, 0x5c, 0x6d, 0x81, 0xd5, 0x46, 0x1f, 0x9c, 0xe7, 0xad, 0x15,
	0xd0, 0x46, 0x41, 0x30, 0x19, 0xce, 0x46, 0xad, 0x02, 0x33, 0x90, 0xe5, 0xbf, 0x81, 0x41, 0x22,
	0x50, 0xe6, 0x19, 0x8a, 0xed, 0xba, 0xcf, 0xbe, 0x3c, 0x08, 0x75, 0x33, 0x72, 0x0e, 0xfd, 0xea,
	0x80, 0x88, 0x33, 0xa5, 0xbd, 0xdf, 0xb8, 0x83, 0xa0, 0x24, 0x97, 0x76, 0x4b, 0xec, 0x85, 0x90,
	0xa3, 0xf6, 0xbd, 0xb9, 0x99, 0xb8, 0x92, 0xda, 0x08, 0xbb, 0x70, 0x76, 0xfa, 0xb0, 0x33, 0x86,
	0x58, 0xc7, 0x5d, 0x08, 0x65, 0x5d, 0xa1, 0x7d, 0x73, 0x2b, 0x2a, 0x9b, 0xe3, 0x2e, 0x84, 0xf2,
	0xb1, 0x67, 0xfe, 0x41, 0x57, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x19, 0x25, 0xe0, 0x33, 0x91,
	0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	UserAdd(ctx context.Context, in *UserAddReq, opts ...grpc.CallOption) (*UserAddRsp, error)
	UserUpdate(ctx context.Context, in *UserUpdateReq, opts ...grpc.CallOption) (*Response, error)
	UserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoRsp, error)
	UserList(ctx context.Context, in *UserListReq, opts ...grpc.CallOption) (*UserListRsp, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) UserAdd(ctx context.Context, in *UserAddReq, opts ...grpc.CallOption) (*UserAddRsp, error) {
	out := new(UserAddRsp)
	err := c.cc.Invoke(ctx, "/user.User/UserAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserUpdate(ctx context.Context, in *UserUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/user.User/UserUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoRsp, error) {
	out := new(UserInfoRsp)
	err := c.cc.Invoke(ctx, "/user.User/UserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserList(ctx context.Context, in *UserListReq, opts ...grpc.CallOption) (*UserListRsp, error) {
	out := new(UserListRsp)
	err := c.cc.Invoke(ctx, "/user.User/UserList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	UserAdd(context.Context, *UserAddReq) (*UserAddRsp, error)
	UserUpdate(context.Context, *UserUpdateReq) (*Response, error)
	UserInfo(context.Context, *UserInfoReq) (*UserInfoRsp, error)
	UserList(context.Context, *UserListReq) (*UserListRsp, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) UserAdd(ctx context.Context, req *UserAddReq) (*UserAddRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserAdd not implemented")
}
func (*UnimplementedUserServer) UserUpdate(ctx context.Context, req *UserUpdateReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserUpdate not implemented")
}
func (*UnimplementedUserServer) UserInfo(ctx context.Context, req *UserInfoReq) (*UserInfoRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfo not implemented")
}
func (*UnimplementedUserServer) UserList(ctx context.Context, req *UserListReq) (*UserListRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserList not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_UserAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/UserAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserAdd(ctx, req.(*UserAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/UserUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserUpdate(ctx, req.(*UserUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/UserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserInfo(ctx, req.(*UserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/UserList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserList(ctx, req.(*UserListReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserAdd",
			Handler:    _User_UserAdd_Handler,
		},
		{
			MethodName: "UserUpdate",
			Handler:    _User_UserUpdate_Handler,
		},
		{
			MethodName: "UserInfo",
			Handler:    _User_UserInfo_Handler,
		},
		{
			MethodName: "UserList",
			Handler:    _User_UserList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
