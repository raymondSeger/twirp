// Code generated by protoc-gen-go.
// source: multiple2.proto
// DO NOT EDIT!

package multiple

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Msg2 struct {
}

func (m *Msg2) Reset()                    { *m = Msg2{} }
func (m *Msg2) String() string            { return proto.CompactTextString(m) }
func (*Msg2) ProtoMessage()               {}
func (*Msg2) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func init() {
	proto.RegisterType((*Msg2)(nil), "twirp.internal.twirptest.multiple.Msg2")
}

func init() { proto.RegisterFile("multiple2.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 114 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0x2d, 0xcd, 0x29,
	0xc9, 0x2c, 0xc8, 0x49, 0x35, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x2c, 0x29, 0xcf,
	0x2c, 0x2a, 0xd0, 0xcb, 0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x03, 0x73, 0x4b, 0x52,
	0x8b, 0x4b, 0xf4, 0x60, 0x2a, 0x95, 0xd8, 0xb8, 0x58, 0x7c, 0x8b, 0xd3, 0x8d, 0x8c, 0x12, 0xb8,
	0x58, 0x82, 0xcb, 0x92, 0x8d, 0x84, 0x22, 0xb8, 0x58, 0x82, 0x53, 0xf3, 0x52, 0x84, 0xd4, 0xf5,
	0x08, 0xea, 0xd5, 0x03, 0x69, 0x94, 0x22, 0x56, 0xa1, 0x13, 0x57, 0x14, 0x07, 0x4c, 0x20, 0x89,
	0x0d, 0xec, 0x3e, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe8, 0xf5, 0x85, 0x74, 0xb2, 0x00,
	0x00, 0x00,
}