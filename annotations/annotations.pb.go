// Code generated by protoc-gen-go. DO NOT EDIT.
// source: annotations/annotations.proto

package annotations

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

type HTTP struct {
	Endpoint             string   `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Method               string   `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HTTP) Reset()         { *m = HTTP{} }
func (m *HTTP) String() string { return proto.CompactTextString(m) }
func (*HTTP) ProtoMessage()    {}
func (*HTTP) Descriptor() ([]byte, []int) {
	return fileDescriptor_21dfaf6fd39fa3b7, []int{0}
}

func (m *HTTP) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HTTP.Unmarshal(m, b)
}
func (m *HTTP) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HTTP.Marshal(b, m, deterministic)
}
func (m *HTTP) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HTTP.Merge(m, src)
}
func (m *HTTP) XXX_Size() int {
	return xxx_messageInfo_HTTP.Size(m)
}
func (m *HTTP) XXX_DiscardUnknown() {
	xxx_messageInfo_HTTP.DiscardUnknown(m)
}

var xxx_messageInfo_HTTP proto.InternalMessageInfo

func (m *HTTP) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func (m *HTTP) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

var E_Http = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*HTTP)(nil),
	Field:         50011,
	Name:          "maestro.http",
	Tag:           "bytes,50011,opt,name=http",
	Filename:      "annotations/annotations.proto",
}

func init() {
	proto.RegisterType((*HTTP)(nil), "maestro.HTTP")
	proto.RegisterExtension(E_Http)
}

func init() {
	proto.RegisterFile("annotations/annotations.proto", fileDescriptor_21dfaf6fd39fa3b7)
}

var fileDescriptor_21dfaf6fd39fa3b7 = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0xcc, 0xcb, 0xcb,
	0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0xd6, 0x47, 0x62, 0xeb, 0x15, 0x14, 0xe5, 0x97, 0xe4,
	0x0b, 0xb1, 0xe7, 0x26, 0xa6, 0x16, 0x97, 0x14, 0xe5, 0x4b, 0x29, 0xa4, 0xe7, 0xe7, 0xa7, 0xe7,
	0xa4, 0xea, 0x83, 0x85, 0x93, 0x4a, 0xd3, 0xf4, 0x53, 0x52, 0x8b, 0x93, 0x8b, 0x32, 0x0b, 0x4a,
	0xf2, 0x8b, 0x20, 0x4a, 0x95, 0xac, 0xb8, 0x58, 0x3c, 0x42, 0x42, 0x02, 0x84, 0xa4, 0xb8, 0x38,
	0x52, 0xf3, 0x52, 0x0a, 0xf2, 0x33, 0xf3, 0x4a, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xe0,
	0x7c, 0x21, 0x31, 0x2e, 0xb6, 0xdc, 0xd4, 0x92, 0x8c, 0xfc, 0x14, 0x09, 0x26, 0xb0, 0x0c, 0x94,
	0x67, 0xe5, 0xcc, 0xc5, 0x92, 0x51, 0x52, 0x52, 0x20, 0x24, 0xa7, 0x07, 0xb1, 0x46, 0x0f, 0x66,
	0x8d, 0x9e, 0x2f, 0x58, 0x81, 0x7f, 0x01, 0xd8, 0x51, 0x12, 0xb7, 0xdb, 0x98, 0x15, 0x18, 0x35,
	0xb8, 0x8d, 0x78, 0xf5, 0xa0, 0xee, 0xd2, 0x03, 0x59, 0x19, 0x04, 0xd6, 0xec, 0xa4, 0x16, 0xa5,
	0x92, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x9f, 0x95, 0x5a, 0x91, 0x99,
	0xa8, 0x0f, 0x55, 0x86, 0xec, 0xb3, 0x24, 0x36, 0xb0, 0xe1, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xc3, 0x29, 0x7e, 0x83, 0xfb, 0x00, 0x00, 0x00,
}
