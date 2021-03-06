// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fleetspeak/src/common/proto/fleetspeak/common.proto

/*
Package fleetspeak is a generated protocol buffer package.

It is generated from these files:
	fleetspeak/src/common/proto/fleetspeak/common.proto
	fleetspeak/src/common/proto/fleetspeak/system.proto

It has these top-level messages:
	Address
	ValidationInfo
	Message
	MessageResult
	Label
	Signature
	WrappedContactData
	ContactData
	EmptyMessage
	MessageAckData
	MessageErrorData
	ClientInfoData
	RemoveServiceData
	SignedClientServiceConfig
	ClientServiceConfig
	ClientServiceConfigs
	RevokedCertificateList
*/
package fleetspeak

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/any"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The message priority. The primary effect is on the ordering of messages
// sent from the client to the server.
type Message_Priority int32

const (
	Message_MEDIUM Message_Priority = 0
	Message_LOW    Message_Priority = 1
	Message_HIGH   Message_Priority = 2
)

var Message_Priority_name = map[int32]string{
	0: "MEDIUM",
	1: "LOW",
	2: "HIGH",
}
var Message_Priority_value = map[string]int32{
	"MEDIUM": 0,
	"LOW":    1,
	"HIGH":   2,
}

func (x Message_Priority) String() string {
	return proto.EnumName(Message_Priority_name, int32(x))
}
func (Message_Priority) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

// An Address identifies the source or destination of a message.
type Address struct {
	// The client_id, if the address refers to a service on a client. If unset,
	// the address refers to a service on the server.
	ClientId []byte `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// The name of the fleetspeak service which sent or should receive the
	// message.  Required.
	ServiceName string `protobuf:"bytes,2,opt,name=service_name,json=serviceName" json:"service_name,omitempty"`
}

func (m *Address) Reset()                    { *m = Address{} }
func (m *Address) String() string            { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()               {}
func (*Address) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Address) GetClientId() []byte {
	if m != nil {
		return m.ClientId
	}
	return nil
}

func (m *Address) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

type ValidationInfo struct {
	Tags map[string]string `protobuf:"bytes,1,rep,name=tags" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ValidationInfo) Reset()                    { *m = ValidationInfo{} }
func (m *ValidationInfo) String() string            { return proto.CompactTextString(m) }
func (*ValidationInfo) ProtoMessage()               {}
func (*ValidationInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ValidationInfo) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type Message struct {
	// A hash based on origin and origin_message_id. It is set by the fleetspeak
	// system on message intake and may be used for deduplication.
	MessageId []byte `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	// The source of the messages. Required.
	Source *Address `protobuf:"bytes,2,opt,name=source" json:"source,omitempty"`
	// An sequence of bytes set by the source in a way to uniquely identify the
	// message among all messages with the same origin.
	SourceMessageId []byte `protobuf:"bytes,3,opt,name=source_message_id,json=sourceMessageId,proto3" json:"source_message_id,omitempty"`
	// The destination of the message. Required.
	Destination *Address `protobuf:"bytes,4,opt,name=destination" json:"destination,omitempty"`
	// The type of message. This field is mostly opaque to the Fleetspeak system,
	// but can be used for some statistics gathering. It is recommended that each
	// service define a static collection of short readable message types and
	// dispatch according to this when processing messages. e.g. "ResourceUsage",
	// "StdOutputData".
	MessageType string `protobuf:"bytes,5,opt,name=message_type,json=messageType" json:"message_type,omitempty"`
	// Set when the message enters the FS system.
	CreationTime *google_protobuf1.Timestamp `protobuf:"bytes,6,opt,name=creation_time,json=creationTime" json:"creation_time,omitempty"`
	// The data associated with this request, accepted types are determined by the
	// service and may depend on message_type. Not typically stored after the
	// message is processed.
	Data *google_protobuf.Any `protobuf:"bytes,7,opt,name=data" json:"data,omitempty"`
	// Additional validation information, set by on the server by the (optional)
	// authorizer component based on WrappedContactData.validators, etc.
	ValidationInfo *ValidationInfo `protobuf:"bytes,8,opt,name=validation_info,json=validationInfo" json:"validation_info,omitempty"`
	// The result of processing the message. Set once processing has finished.
	Result   *MessageResult   `protobuf:"bytes,9,opt,name=result" json:"result,omitempty"`
	Priority Message_Priority `protobuf:"varint,10,opt,name=priority,enum=fleetspeak.Message_Priority" json:"priority,omitempty"`
	// A background message does not count as activity when deciding how fast to
	// poll the server. This flag should be set on messages which are unlikely to
	// trigger additional activity.
	Background bool `protobuf:"varint,11,opt,name=background" json:"background,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Message) GetMessageId() []byte {
	if m != nil {
		return m.MessageId
	}
	return nil
}

func (m *Message) GetSource() *Address {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *Message) GetSourceMessageId() []byte {
	if m != nil {
		return m.SourceMessageId
	}
	return nil
}

func (m *Message) GetDestination() *Address {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (m *Message) GetMessageType() string {
	if m != nil {
		return m.MessageType
	}
	return ""
}

func (m *Message) GetCreationTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreationTime
	}
	return nil
}

func (m *Message) GetData() *google_protobuf.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Message) GetValidationInfo() *ValidationInfo {
	if m != nil {
		return m.ValidationInfo
	}
	return nil
}

func (m *Message) GetResult() *MessageResult {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *Message) GetPriority() Message_Priority {
	if m != nil {
		return m.Priority
	}
	return Message_MEDIUM
}

func (m *Message) GetBackground() bool {
	if m != nil {
		return m.Background
	}
	return false
}

type MessageResult struct {
	// The time that processing finished.
	ProcessedTime *google_protobuf1.Timestamp `protobuf:"bytes,2,opt,name=processed_time,json=processedTime" json:"processed_time,omitempty"`
	// Set when processing ended with a permanent failure.
	Failed bool `protobuf:"varint,3,opt,name=failed" json:"failed,omitempty"`
	// A human readable error message, normally set when failed is true.
	FailedReason string `protobuf:"bytes,4,opt,name=failed_reason,json=failedReason" json:"failed_reason,omitempty"`
}

func (m *MessageResult) Reset()                    { *m = MessageResult{} }
func (m *MessageResult) String() string            { return proto.CompactTextString(m) }
func (*MessageResult) ProtoMessage()               {}
func (*MessageResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *MessageResult) GetProcessedTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.ProcessedTime
	}
	return nil
}

func (m *MessageResult) GetFailed() bool {
	if m != nil {
		return m.Failed
	}
	return false
}

func (m *MessageResult) GetFailedReason() string {
	if m != nil {
		return m.FailedReason
	}
	return ""
}

// A Label is a tag assigned to a client by a plugin. Primary use is to limit
// broadcasts to specific clients.
type Label struct {
	// The service which set this label.
	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName" json:"service_name,omitempty"`
	// A free form tag choosen by the setting plugin.
	Label string `protobuf:"bytes,2,opt,name=label" json:"label,omitempty"`
}

func (m *Label) Reset()                    { *m = Label{} }
func (m *Label) String() string            { return proto.CompactTextString(m) }
func (*Label) ProtoMessage()               {}
func (*Label) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Label) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *Label) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

type Signature struct {
	// A chain of ASN.1 DER encoded x509 certificates.
	Certificate [][]byte `protobuf:"bytes,1,rep,name=certificate,proto3" json:"certificate,omitempty"`
	// Indicates the choice of signature algorithm, a constant from
	// https://golang.org/pkg/crypto/x509/#SignatureAlgorithm
	Algorithm int32 `protobuf:"varint,2,opt,name=algorithm" json:"algorithm,omitempty"`
	// A signature of the validated data, it should be consistent with both the
	// algorithm choice and the first element of the certificate chain.
	Signature []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *Signature) Reset()                    { *m = Signature{} }
func (m *Signature) String() string            { return proto.CompactTextString(m) }
func (*Signature) ProtoMessage()               {}
func (*Signature) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Signature) GetCertificate() [][]byte {
	if m != nil {
		return m.Certificate
	}
	return nil
}

func (m *Signature) GetAlgorithm() int32 {
	if m != nil {
		return m.Algorithm
	}
	return 0
}

func (m *Signature) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

// A WrappedContactData is provided by the client to the server with every
// contact.
type WrappedContactData struct {
	ContactData []byte       `protobuf:"bytes,1,opt,name=contact_data,json=contactData,proto3" json:"contact_data,omitempty"`
	Signatures  []*Signature `protobuf:"bytes,2,rep,name=signatures" json:"signatures,omitempty"`
	// contact_data.
	ClientLabels []string `protobuf:"bytes,3,rep,name=client_labels,json=clientLabels" json:"client_labels,omitempty"`
}

func (m *WrappedContactData) Reset()                    { *m = WrappedContactData{} }
func (m *WrappedContactData) String() string            { return proto.CompactTextString(m) }
func (*WrappedContactData) ProtoMessage()               {}
func (*WrappedContactData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *WrappedContactData) GetContactData() []byte {
	if m != nil {
		return m.ContactData
	}
	return nil
}

func (m *WrappedContactData) GetSignatures() []*Signature {
	if m != nil {
		return m.Signatures
	}
	return nil
}

func (m *WrappedContactData) GetClientLabels() []string {
	if m != nil {
		return m.ClientLabels
	}
	return nil
}

// On every contact, the client and server exchange ContactData messages.
type ContactData struct {
	// During every contact, the server passes a random sequencing_nonce to the
	// client, and the client provides the sequencing_nonce to the server during
	// the next contact.
	SequencingNonce uint64     `protobuf:"varint,1,opt,name=sequencing_nonce,json=sequencingNonce" json:"sequencing_nonce,omitempty"`
	Messages        []*Message `protobuf:"bytes,2,rep,name=messages" json:"messages,omitempty"`
	// Records the client's current time setting, as of the creation of this
	// ContactData. Only set by the client.
	ClientClock *google_protobuf1.Timestamp `protobuf:"bytes,3,opt,name=client_clock,json=clientClock" json:"client_clock,omitempty"`
}

func (m *ContactData) Reset()                    { *m = ContactData{} }
func (m *ContactData) String() string            { return proto.CompactTextString(m) }
func (*ContactData) ProtoMessage()               {}
func (*ContactData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ContactData) GetSequencingNonce() uint64 {
	if m != nil {
		return m.SequencingNonce
	}
	return 0
}

func (m *ContactData) GetMessages() []*Message {
	if m != nil {
		return m.Messages
	}
	return nil
}

func (m *ContactData) GetClientClock() *google_protobuf1.Timestamp {
	if m != nil {
		return m.ClientClock
	}
	return nil
}

// An empty message, typically used as a trivial RPC response.
type EmptyMessage struct {
}

func (m *EmptyMessage) Reset()                    { *m = EmptyMessage{} }
func (m *EmptyMessage) String() string            { return proto.CompactTextString(m) }
func (*EmptyMessage) ProtoMessage()               {}
func (*EmptyMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func init() {
	proto.RegisterType((*Address)(nil), "fleetspeak.Address")
	proto.RegisterType((*ValidationInfo)(nil), "fleetspeak.ValidationInfo")
	proto.RegisterType((*Message)(nil), "fleetspeak.Message")
	proto.RegisterType((*MessageResult)(nil), "fleetspeak.MessageResult")
	proto.RegisterType((*Label)(nil), "fleetspeak.Label")
	proto.RegisterType((*Signature)(nil), "fleetspeak.Signature")
	proto.RegisterType((*WrappedContactData)(nil), "fleetspeak.WrappedContactData")
	proto.RegisterType((*ContactData)(nil), "fleetspeak.ContactData")
	proto.RegisterType((*EmptyMessage)(nil), "fleetspeak.EmptyMessage")
	proto.RegisterEnum("fleetspeak.Message_Priority", Message_Priority_name, Message_Priority_value)
}

func init() {
	proto.RegisterFile("fleetspeak/src/common/proto/fleetspeak/common.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 775 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xd1, 0x6e, 0xe3, 0x44,
	0x14, 0xc5, 0x49, 0x9a, 0xc6, 0xd7, 0x69, 0x1a, 0x66, 0x17, 0xe4, 0x2d, 0x0b, 0x64, 0x0d, 0x0f,
	0x59, 0x90, 0x12, 0xd1, 0xd5, 0x8a, 0x0a, 0x09, 0x41, 0xd5, 0xad, 0xd8, 0x48, 0xdb, 0x05, 0x0d,
	0x85, 0x7d, 0xb4, 0xa6, 0xe3, 0x89, 0x19, 0xd5, 0x9e, 0x31, 0x33, 0x93, 0x4a, 0x79, 0xe6, 0x03,
	0x78, 0xe0, 0x27, 0xf8, 0x0f, 0x7e, 0x0c, 0x79, 0x66, 0xec, 0xb8, 0x74, 0xcb, 0xbe, 0xcd, 0x9c,
	0x7b, 0xee, 0xf1, 0x9d, 0x73, 0xef, 0x35, 0x3c, 0x5b, 0x17, 0x8c, 0x19, 0x5d, 0x31, 0x72, 0xbd,
	0xd4, 0x8a, 0x2e, 0xa9, 0x2c, 0x4b, 0x29, 0x96, 0x95, 0x92, 0x46, 0x2e, 0x3b, 0x31, 0x87, 0x2f,
	0x2c, 0x8e, 0x60, 0x17, 0x38, 0x7a, 0x94, 0x4b, 0x99, 0x17, 0xcc, 0x65, 0x5c, 0x6d, 0xd6, 0x4b,
	0x22, 0xb6, 0x8e, 0x76, 0xf4, 0xe9, 0x7f, 0x43, 0x86, 0x97, 0x4c, 0x1b, 0x52, 0x56, 0x8e, 0x90,
	0xac, 0x60, 0xff, 0x34, 0xcb, 0x14, 0xd3, 0x1a, 0x7d, 0x04, 0x21, 0x2d, 0x38, 0x13, 0x26, 0xe5,
	0x59, 0x1c, 0xcc, 0x82, 0xf9, 0x18, 0x8f, 0x1c, 0xb0, 0xca, 0xd0, 0x13, 0x18, 0x6b, 0xa6, 0x6e,
	0x38, 0x65, 0xa9, 0x20, 0x25, 0x8b, 0x7b, 0xb3, 0x60, 0x1e, 0xe2, 0xc8, 0x63, 0xaf, 0x49, 0xc9,
	0x92, 0x3f, 0x02, 0x98, 0xfc, 0x4a, 0x0a, 0x9e, 0x11, 0xc3, 0xa5, 0x58, 0x89, 0xb5, 0x44, 0x27,
	0x30, 0x30, 0x24, 0xd7, 0x71, 0x30, 0xeb, 0xcf, 0xa3, 0xe3, 0xcf, 0x17, 0xbb, 0xa2, 0x17, 0xb7,
	0x99, 0x8b, 0x4b, 0x92, 0xeb, 0x73, 0x61, 0xd4, 0x16, 0xdb, 0x8c, 0xa3, 0xaf, 0x21, 0x6c, 0x21,
	0x34, 0x85, 0xfe, 0x35, 0xdb, 0xda, 0x9a, 0x42, 0x5c, 0x1f, 0xd1, 0x43, 0xd8, 0xbb, 0x21, 0xc5,
	0xa6, 0xa9, 0xc3, 0x5d, 0xbe, 0xe9, 0x9d, 0x04, 0xc9, 0x3f, 0x03, 0xd8, 0xbf, 0x60, 0x5a, 0x93,
	0x9c, 0xa1, 0x8f, 0x01, 0x4a, 0x77, 0xdc, 0x3d, 0x29, 0xf4, 0xc8, 0x2a, 0x43, 0x5f, 0xc2, 0x50,
	0xcb, 0x8d, 0xa2, 0x4e, 0x25, 0x3a, 0x7e, 0xd0, 0xad, 0xcf, 0xbb, 0x82, 0x3d, 0x05, 0x7d, 0x01,
	0xef, 0xbb, 0x53, 0xda, 0x91, 0xec, 0x5b, 0xc9, 0x43, 0x17, 0xb8, 0x68, 0x85, 0x9f, 0x43, 0x94,
	0x31, 0x6d, 0xb8, 0xb0, 0xef, 0x8b, 0x07, 0xf7, 0xab, 0x77, 0x79, 0xb5, 0xc7, 0x8d, 0xb6, 0xd9,
	0x56, 0x2c, 0xde, 0x73, 0x1e, 0x7b, 0xec, 0x72, 0x5b, 0x31, 0xf4, 0x1d, 0x1c, 0x50, 0xc5, 0x2c,
	0x3d, 0xad, 0x5b, 0x19, 0x0f, 0xad, 0xf6, 0xd1, 0xc2, 0xf5, 0x79, 0xd1, 0xf4, 0x79, 0x71, 0xd9,
	0xf4, 0x19, 0x8f, 0x9b, 0x84, 0x1a, 0x42, 0x73, 0x18, 0x64, 0xc4, 0x90, 0x78, 0xdf, 0xe6, 0x3d,
	0xbc, 0x93, 0x77, 0x2a, 0xb6, 0xd8, 0x32, 0xd0, 0x19, 0x1c, 0xde, 0xb4, 0x3d, 0x4a, 0xb9, 0x58,
	0xcb, 0x78, 0xe4, 0x3f, 0x76, 0x6f, 0x1b, 0xf1, 0xe4, 0xe6, 0xf6, 0x00, 0x7c, 0x05, 0x43, 0xc5,
	0xf4, 0xa6, 0x30, 0x71, 0x68, 0x73, 0x1f, 0x75, 0x73, 0xbd, 0x61, 0xd8, 0x12, 0xb0, 0x27, 0xa2,
	0x13, 0x18, 0x55, 0x8a, 0x4b, 0xc5, 0xcd, 0x36, 0x86, 0x59, 0x30, 0x9f, 0x1c, 0x3f, 0x7e, 0x4b,
	0xd2, 0xe2, 0x27, 0xcf, 0xc1, 0x2d, 0x1b, 0x7d, 0x02, 0x70, 0x45, 0xe8, 0x75, 0xae, 0xe4, 0x46,
	0x64, 0x71, 0x34, 0x0b, 0xe6, 0x23, 0xdc, 0x41, 0x92, 0xa7, 0x30, 0x6a, 0xb2, 0x10, 0xc0, 0xf0,
	0xe2, 0xfc, 0xc5, 0xea, 0x97, 0x8b, 0xe9, 0x7b, 0x68, 0x1f, 0xfa, 0xaf, 0x7e, 0x7c, 0x33, 0x0d,
	0xd0, 0x08, 0x06, 0x2f, 0x57, 0x3f, 0xbc, 0x9c, 0xf6, 0x92, 0x3f, 0x03, 0x38, 0xb8, 0x55, 0x1e,
	0x3a, 0x85, 0x49, 0xa5, 0x24, 0x65, 0x5a, 0xb3, 0xcc, 0x59, 0xdf, 0x7b, 0xa7, 0xf5, 0x07, 0x6d,
	0x86, 0xf5, 0xfe, 0x43, 0x18, 0xae, 0x09, 0x2f, 0x98, 0x9b, 0x9b, 0x11, 0xf6, 0x37, 0xf4, 0x19,
	0x1c, 0xb8, 0x53, 0xaa, 0x18, 0xd1, 0x7e, 0x60, 0x42, 0x3c, 0x76, 0x20, 0xb6, 0x58, 0xf2, 0x3d,
	0xec, 0xbd, 0x22, 0x57, 0xac, 0xb8, 0xb3, 0x89, 0xc1, 0x9d, 0x4d, 0xac, 0xb7, 0xa3, 0xa8, 0xb9,
	0xcd, 0x76, 0xd8, 0x4b, 0xc2, 0x21, 0xfc, 0x99, 0xe7, 0x82, 0x98, 0x8d, 0x62, 0x68, 0x06, 0x11,
	0x65, 0xca, 0xf0, 0x35, 0xa7, 0xc4, 0x30, 0xbb, 0xa0, 0x63, 0xdc, 0x85, 0xd0, 0x63, 0x08, 0x49,
	0x91, 0xd7, 0x6e, 0xfd, 0x56, 0x5a, 0xa1, 0x3d, 0xbc, 0x03, 0xea, 0xa8, 0x6e, 0xc4, 0xfc, 0x1a,
	0xec, 0x80, 0xe4, 0xaf, 0x00, 0xd0, 0x1b, 0x45, 0xaa, 0x8a, 0x65, 0x67, 0x52, 0x18, 0x42, 0xcd,
	0x8b, 0x7a, 0xa4, 0x9e, 0xc0, 0x98, 0xba, 0x6b, 0x6a, 0x87, 0xd0, 0x6d, 0x64, 0x44, 0x3b, 0x94,
	0xe7, 0x00, 0xad, 0x8c, 0x8e, 0x7b, 0xf6, 0xbf, 0xf1, 0x41, 0xb7, 0xff, 0xed, 0x13, 0x70, 0x87,
	0x58, 0x5b, 0xe8, 0xff, 0x5d, 0xf6, 0xad, 0x3a, 0xee, 0xcf, 0xfa, 0xb5, 0x85, 0x0e, 0xb4, 0xc6,
	0xe9, 0xe4, 0xef, 0x00, 0xa2, 0x6e, 0x39, 0x4f, 0x61, 0xaa, 0xd9, 0xef, 0x1b, 0x26, 0x28, 0x17,
	0x79, 0x2a, 0xa4, 0xa0, 0xce, 0xcd, 0x01, 0x3e, 0xdc, 0xe1, 0xaf, 0x6b, 0x18, 0x2d, 0x61, 0xe4,
	0xd7, 0xb0, 0x29, 0xea, 0xc1, 0xdb, 0x26, 0xb9, 0x25, 0xa1, 0x6f, 0xc1, 0x7f, 0x3b, 0xa5, 0x85,
	0xa4, 0xd7, 0xd6, 0xa2, 0xff, 0x1f, 0x96, 0xc8, 0xf1, 0xcf, 0x6a, 0x7a, 0x32, 0x81, 0xf1, 0x79,
	0x59, 0x99, 0xad, 0x17, 0xbe, 0x1a, 0xda, 0x84, 0x67, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xe3,
	0x90, 0xd9, 0x82, 0x2c, 0x06, 0x00, 0x00,
}
