// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Events.proto

package orionproto

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

type LogLevel int32

const (
	LogLevel_DEBUG    LogLevel = 0
	LogLevel_INFO     LogLevel = 1
	LogLevel_WARN     LogLevel = 2
	LogLevel_ERROR    LogLevel = 3
	LogLevel_CRITICAL LogLevel = 4
)

var LogLevel_name = map[int32]string{
	0: "DEBUG",
	1: "INFO",
	2: "WARN",
	3: "ERROR",
	4: "CRITICAL",
}
var LogLevel_value = map[string]int32{
	"DEBUG":    0,
	"INFO":     1,
	"WARN":     2,
	"ERROR":    3,
	"CRITICAL": 4,
}

func (x LogLevel) String() string {
	return proto.EnumName(LogLevel_name, int32(x))
}
func (LogLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_Events_0a24b42278d0690a, []int{0}
}

type StartEvent struct {
	EventId uint64 `protobuf:"varint,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	// Types that are valid to be assigned to Metadata:
	//	*StartEvent_JsonString
	Metadata             isStartEvent_Metadata `protobuf_oneof:"metadata"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *StartEvent) Reset()         { *m = StartEvent{} }
func (m *StartEvent) String() string { return proto.CompactTextString(m) }
func (*StartEvent) ProtoMessage()    {}
func (*StartEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_Events_0a24b42278d0690a, []int{0}
}
func (m *StartEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartEvent.Unmarshal(m, b)
}
func (m *StartEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartEvent.Marshal(b, m, deterministic)
}
func (dst *StartEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartEvent.Merge(dst, src)
}
func (m *StartEvent) XXX_Size() int {
	return xxx_messageInfo_StartEvent.Size(m)
}
func (m *StartEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_StartEvent.DiscardUnknown(m)
}

var xxx_messageInfo_StartEvent proto.InternalMessageInfo

type isStartEvent_Metadata interface {
	isStartEvent_Metadata()
}

type StartEvent_JsonString struct {
	JsonString string `protobuf:"bytes,3,opt,name=jsonString,proto3,oneof"`
}

func (*StartEvent_JsonString) isStartEvent_Metadata() {}

func (m *StartEvent) GetMetadata() isStartEvent_Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *StartEvent) GetEventId() uint64 {
	if m != nil {
		return m.EventId
	}
	return 0
}

func (m *StartEvent) GetJsonString() string {
	if x, ok := m.GetMetadata().(*StartEvent_JsonString); ok {
		return x.JsonString
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*StartEvent) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _StartEvent_OneofMarshaler, _StartEvent_OneofUnmarshaler, _StartEvent_OneofSizer, []interface{}{
		(*StartEvent_JsonString)(nil),
	}
}

func _StartEvent_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*StartEvent)
	// metadata
	switch x := m.Metadata.(type) {
	case *StartEvent_JsonString:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.JsonString)
	case nil:
	default:
		return fmt.Errorf("StartEvent.Metadata has unexpected type %T", x)
	}
	return nil
}

func _StartEvent_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*StartEvent)
	switch tag {
	case 3: // metadata.jsonString
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Metadata = &StartEvent_JsonString{x}
		return true, err
	default:
		return false, nil
	}
}

func _StartEvent_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*StartEvent)
	// metadata
	switch x := m.Metadata.(type) {
	case *StartEvent_JsonString:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.JsonString)))
		n += len(x.JsonString)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type EndEvent struct {
	EventId uint64 `protobuf:"varint,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	// Types that are valid to be assigned to Metadata:
	//	*EndEvent_JsonString
	Metadata             isEndEvent_Metadata `protobuf_oneof:"metadata"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *EndEvent) Reset()         { *m = EndEvent{} }
func (m *EndEvent) String() string { return proto.CompactTextString(m) }
func (*EndEvent) ProtoMessage()    {}
func (*EndEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_Events_0a24b42278d0690a, []int{1}
}
func (m *EndEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndEvent.Unmarshal(m, b)
}
func (m *EndEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndEvent.Marshal(b, m, deterministic)
}
func (dst *EndEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndEvent.Merge(dst, src)
}
func (m *EndEvent) XXX_Size() int {
	return xxx_messageInfo_EndEvent.Size(m)
}
func (m *EndEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_EndEvent.DiscardUnknown(m)
}

var xxx_messageInfo_EndEvent proto.InternalMessageInfo

type isEndEvent_Metadata interface {
	isEndEvent_Metadata()
}

type EndEvent_JsonString struct {
	JsonString string `protobuf:"bytes,3,opt,name=jsonString,proto3,oneof"`
}

func (*EndEvent_JsonString) isEndEvent_Metadata() {}

func (m *EndEvent) GetMetadata() isEndEvent_Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *EndEvent) GetEventId() uint64 {
	if m != nil {
		return m.EventId
	}
	return 0
}

func (m *EndEvent) GetJsonString() string {
	if x, ok := m.GetMetadata().(*EndEvent_JsonString); ok {
		return x.JsonString
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*EndEvent) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _EndEvent_OneofMarshaler, _EndEvent_OneofUnmarshaler, _EndEvent_OneofSizer, []interface{}{
		(*EndEvent_JsonString)(nil),
	}
}

func _EndEvent_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*EndEvent)
	// metadata
	switch x := m.Metadata.(type) {
	case *EndEvent_JsonString:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.JsonString)
	case nil:
	default:
		return fmt.Errorf("EndEvent.Metadata has unexpected type %T", x)
	}
	return nil
}

func _EndEvent_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*EndEvent)
	switch tag {
	case 3: // metadata.jsonString
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Metadata = &EndEvent_JsonString{x}
		return true, err
	default:
		return false, nil
	}
}

func _EndEvent_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*EndEvent)
	// metadata
	switch x := m.Metadata.(type) {
	case *EndEvent_JsonString:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.JsonString)))
		n += len(x.JsonString)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type LogEvent struct {
	EventId uint64   `protobuf:"varint,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	Level   LogLevel `protobuf:"varint,2,opt,name=level,proto3,enum=com.thapovan.orion.proto.LogLevel" json:"level,omitempty"`
	Message string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	// Types that are valid to be assigned to Metadata:
	//	*LogEvent_JsonString
	Metadata             isLogEvent_Metadata `protobuf_oneof:"metadata"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *LogEvent) Reset()         { *m = LogEvent{} }
func (m *LogEvent) String() string { return proto.CompactTextString(m) }
func (*LogEvent) ProtoMessage()    {}
func (*LogEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_Events_0a24b42278d0690a, []int{2}
}
func (m *LogEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogEvent.Unmarshal(m, b)
}
func (m *LogEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogEvent.Marshal(b, m, deterministic)
}
func (dst *LogEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogEvent.Merge(dst, src)
}
func (m *LogEvent) XXX_Size() int {
	return xxx_messageInfo_LogEvent.Size(m)
}
func (m *LogEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_LogEvent.DiscardUnknown(m)
}

var xxx_messageInfo_LogEvent proto.InternalMessageInfo

type isLogEvent_Metadata interface {
	isLogEvent_Metadata()
}

type LogEvent_JsonString struct {
	JsonString string `protobuf:"bytes,5,opt,name=jsonString,proto3,oneof"`
}

func (*LogEvent_JsonString) isLogEvent_Metadata() {}

func (m *LogEvent) GetMetadata() isLogEvent_Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *LogEvent) GetEventId() uint64 {
	if m != nil {
		return m.EventId
	}
	return 0
}

func (m *LogEvent) GetLevel() LogLevel {
	if m != nil {
		return m.Level
	}
	return LogLevel_DEBUG
}

func (m *LogEvent) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *LogEvent) GetJsonString() string {
	if x, ok := m.GetMetadata().(*LogEvent_JsonString); ok {
		return x.JsonString
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*LogEvent) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _LogEvent_OneofMarshaler, _LogEvent_OneofUnmarshaler, _LogEvent_OneofSizer, []interface{}{
		(*LogEvent_JsonString)(nil),
	}
}

func _LogEvent_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*LogEvent)
	// metadata
	switch x := m.Metadata.(type) {
	case *LogEvent_JsonString:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.JsonString)
	case nil:
	default:
		return fmt.Errorf("LogEvent.Metadata has unexpected type %T", x)
	}
	return nil
}

func _LogEvent_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*LogEvent)
	switch tag {
	case 5: // metadata.jsonString
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Metadata = &LogEvent_JsonString{x}
		return true, err
	default:
		return false, nil
	}
}

func _LogEvent_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*LogEvent)
	// metadata
	switch x := m.Metadata.(type) {
	case *LogEvent_JsonString:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.JsonString)))
		n += len(x.JsonString)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*StartEvent)(nil), "com.thapovan.orion.proto.StartEvent")
	proto.RegisterType((*EndEvent)(nil), "com.thapovan.orion.proto.EndEvent")
	proto.RegisterType((*LogEvent)(nil), "com.thapovan.orion.proto.LogEvent")
	proto.RegisterEnum("com.thapovan.orion.proto.LogLevel", LogLevel_name, LogLevel_value)
}

func init() { proto.RegisterFile("Events.proto", fileDescriptor_Events_0a24b42278d0690a) }

var fileDescriptor_Events_0a24b42278d0690a = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x8f, 0x41, 0x4b, 0xc3, 0x30,
	0x1c, 0xc5, 0x97, 0xad, 0x75, 0xd9, 0x9f, 0x31, 0x4a, 0x4e, 0xf1, 0x56, 0x7a, 0x2a, 0x1e, 0x72,
	0xd0, 0x8b, 0xd7, 0x75, 0x56, 0x2d, 0x94, 0x4d, 0x52, 0x87, 0xe0, 0x45, 0xa2, 0x0d, 0xb5, 0xb2,
	0xe6, 0x3f, 0xda, 0xd0, 0xaf, 0xe4, 0xd7, 0x94, 0xa6, 0x08, 0x22, 0xa8, 0x17, 0x4f, 0x79, 0x2f,
	0xef, 0xe5, 0xf1, 0x0b, 0x2c, 0xd3, 0x5e, 0x1b, 0xdb, 0x89, 0x63, 0x8b, 0x16, 0x19, 0x7f, 0xc1,
	0x46, 0xd8, 0x57, 0x75, 0xc4, 0x5e, 0x19, 0x81, 0x6d, 0x8d, 0x66, 0x4c, 0xa2, 0x3d, 0x40, 0x61,
	0x55, 0x6b, 0x5d, 0x9d, 0x9d, 0x02, 0xd5, 0x83, 0x78, 0xaa, 0x4b, 0x4e, 0x42, 0x12, 0x7b, 0x72,
	0xee, 0x7c, 0x56, 0xb2, 0x10, 0xe0, 0xad, 0x43, 0x53, 0xd8, 0xb6, 0x36, 0x15, 0x9f, 0x85, 0x24,
	0x5e, 0xdc, 0x4e, 0xe4, 0x97, 0xbb, 0x04, 0x80, 0x36, 0xda, 0xaa, 0x52, 0x59, 0x15, 0x15, 0x40,
	0x53, 0x53, 0xfe, 0xf3, 0xe8, 0x3b, 0x01, 0x9a, 0x63, 0xf5, 0xe7, 0xea, 0x25, 0xf8, 0x07, 0xdd,
	0xeb, 0x03, 0x9f, 0x86, 0x24, 0x5e, 0x9d, 0x47, 0xe2, 0xa7, 0xdf, 0x8b, 0x1c, 0xab, 0x7c, 0x68,
	0xca, 0xf1, 0x01, 0xe3, 0x30, 0x6f, 0x74, 0xd7, 0xa9, 0x4a, 0x8f, 0x30, 0xf2, 0xd3, 0x7e, 0x23,
	0xf5, 0x7f, 0x27, 0x3d, 0x4b, 0x1c, 0xa8, 0x9b, 0x66, 0x0b, 0xf0, 0xaf, 0xd2, 0x64, 0x7f, 0x13,
	0x4c, 0x18, 0x05, 0x2f, 0xdb, 0x5e, 0xef, 0x02, 0x32, 0xa8, 0x87, 0xb5, 0xdc, 0x06, 0xd3, 0x21,
	0x4e, 0xa5, 0xdc, 0xc9, 0x60, 0xc6, 0x96, 0x40, 0x37, 0x32, 0xbb, 0xcf, 0x36, 0xeb, 0x3c, 0xf0,
	0x92, 0xd5, 0x1d, 0x79, 0x04, 0x07, 0xeb, 0x58, 0x9f, 0x4f, 0xdc, 0x71, 0xf1, 0x11, 0x00, 0x00,
	0xff, 0xff, 0x44, 0xb8, 0xba, 0x97, 0xda, 0x01, 0x00, 0x00,
}
