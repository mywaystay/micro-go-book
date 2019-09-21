// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package pb

import (
	context "context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"math"
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

type UserRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRequest) Reset()         { *m = UserRequest{} }
func (m *UserRequest) String() string { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()    {}
func (*UserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *UserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRequest.Unmarshal(m, b)
}
func (m *UserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRequest.Marshal(b, m, deterministic)
}
func (m *UserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRequest.Merge(m, src)
}
func (m *UserRequest) XXX_Size() int {
	return xxx_messageInfo_UserRequest.Size(m)
}
func (m *UserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRequest proto.InternalMessageInfo

func (m *UserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type UserResponse struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (m *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(m, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func (m *UserResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*UserRequest)(nil), "pb.UserRequest")
	proto.RegisterType((*UserResponse)(nil), "pb.UserResponse")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x72, 0xe5, 0xe2, 0x0e, 0x2d,
	0x4e, 0x2d, 0x0a, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe2, 0xe2, 0x00, 0x29, 0xc8,
	0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf3, 0x41, 0x72, 0x05, 0x89,
	0xc5, 0xc5, 0xe5, 0xf9, 0x45, 0x29, 0x12, 0x4c, 0x10, 0x39, 0x18, 0x5f, 0xc9, 0x82, 0x8b, 0x07,
	0x62, 0x4c, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x18, 0x17, 0x5b, 0x51, 0x6a, 0x71, 0x69,
	0x4e, 0x09, 0xd8, 0x14, 0x8e, 0x20, 0x28, 0x4f, 0x48, 0x80, 0x8b, 0x39, 0xb5, 0xa8, 0x08, 0xaa,
	0x1d, 0xc4, 0x34, 0xb2, 0x86, 0x38, 0x20, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0x48, 0x87,
	0x8b, 0xd5, 0x39, 0x23, 0x35, 0x39, 0x5b, 0x88, 0x5f, 0xaf, 0x20, 0x49, 0x0f, 0xc9, 0x69, 0x52,
	0x02, 0x08, 0x01, 0x88, 0x25, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0x8f, 0x18, 0x03, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x9d, 0xa2, 0x3d, 0x18, 0xd6, 0x00, 0x00, 0x00,
}

type UserServiceServer interface {
	Check(context.Context, *UserRequest) (*UserResponse, error)
}
