// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fleetspeak/src/client/stdinservice/proto/fleetspeak_stdinservice/config.proto

/*
Package fleetspeak_stdinservice is a generated protocol buffer package.

It is generated from these files:
	fleetspeak/src/client/stdinservice/proto/fleetspeak_stdinservice/config.proto
	fleetspeak/src/client/stdinservice/proto/fleetspeak_stdinservice/messages.proto

It has these top-level messages:
	Config
	InputMessage
	OutputMessage
*/
package fleetspeak_stdinservice

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The configuration information expected by stdinservice.Factory
// in ClientServiceConfig.config.
type Config struct {
	Cmd string `protobuf:"bytes,1,opt,name=cmd" json:"cmd,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Config) GetCmd() string {
	if m != nil {
		return m.Cmd
	}
	return ""
}

func init() {
	proto.RegisterType((*Config)(nil), "fleetspeak.stdinservice.Config")
}

func init() {
	proto.RegisterFile("fleetspeak/src/client/stdinservice/proto/fleetspeak_stdinservice/config.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 116 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xf2, 0x4d, 0xcb, 0x49, 0x4d,
	0x2d, 0x29, 0x2e, 0x48, 0x4d, 0xcc, 0xd6, 0x2f, 0x2e, 0x4a, 0xd6, 0x4f, 0xce, 0xc9, 0x4c, 0xcd,
	0x2b, 0xd1, 0x2f, 0x2e, 0x49, 0xc9, 0xcc, 0x2b, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2f,
	0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x47, 0x28, 0x8c, 0x47, 0x91, 0x4d, 0xce, 0xcf, 0x4b, 0xcb, 0x4c,
	0xd7, 0x03, 0x2b, 0x12, 0x12, 0x47, 0xa8, 0xd2, 0x43, 0x56, 0xa5, 0x24, 0xc5, 0xc5, 0xe6, 0x0c,
	0x56, 0x28, 0x24, 0xc0, 0xc5, 0x9c, 0x9c, 0x9b, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04,
	0x62, 0x26, 0xb1, 0x81, 0xf5, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa2, 0xa8, 0x77, 0x6f,
	0x8c, 0x00, 0x00, 0x00,
}
