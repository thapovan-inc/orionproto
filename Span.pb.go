// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Span.proto

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

type Span struct {
	TraceContext *Trace `protobuf:"bytes,1,opt,name=trace_context,json=traceContext,proto3" json:"trace_context,omitempty"`
	SpanId       string `protobuf:"bytes,2,opt,name=span_id,json=spanId,proto3" json:"span_id,omitempty"`
	// Types that are valid to be assigned to Event:
	//	*Span_StartEvent
	//	*Span_EndEvent
	//	*Span_LogEvent
	Event                 isSpan_Event `protobuf_oneof:"event"`
	Timestamp             uint64       `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	ServiceName           string       `protobuf:"bytes,7,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	EventLocation         string       `protobuf:"bytes,8,opt,name=event_location,json=eventLocation,proto3" json:"event_location,omitempty"`
	ParentSpanId          string       `protobuf:"bytes,9,opt,name=parent_span_id,json=parentSpanId,proto3" json:"parent_span_id,omitempty"`
	InternalSpanRefNumber uint64       `protobuf:"varint,10000,opt,name=internal_span_ref_number,json=internalSpanRefNumber,proto3" json:"internal_span_ref_number,omitempty"`
	InternalMeta          []byte       `protobuf:"bytes,100001,opt,name=internal_meta,json=internalMeta,proto3" json:"internal_meta,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}     `json:"-"`
	XXX_unrecognized      []byte       `json:"-"`
	XXX_sizecache         int32        `json:"-"`
}

func (m *Span) Reset()         { *m = Span{} }
func (m *Span) String() string { return proto.CompactTextString(m) }
func (*Span) ProtoMessage()    {}
func (*Span) Descriptor() ([]byte, []int) {
	return fileDescriptor_Span_334f9dc70dfb4c18, []int{0}
}
func (m *Span) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Span.Unmarshal(m, b)
}
func (m *Span) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Span.Marshal(b, m, deterministic)
}
func (dst *Span) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Span.Merge(dst, src)
}
func (m *Span) XXX_Size() int {
	return xxx_messageInfo_Span.Size(m)
}
func (m *Span) XXX_DiscardUnknown() {
	xxx_messageInfo_Span.DiscardUnknown(m)
}

var xxx_messageInfo_Span proto.InternalMessageInfo

type isSpan_Event interface {
	isSpan_Event()
}

type Span_StartEvent struct {
	StartEvent *StartEvent `protobuf:"bytes,3,opt,name=start_event,json=startEvent,proto3,oneof"`
}
type Span_EndEvent struct {
	EndEvent *EndEvent `protobuf:"bytes,4,opt,name=end_event,json=endEvent,proto3,oneof"`
}
type Span_LogEvent struct {
	LogEvent *LogEvent `protobuf:"bytes,5,opt,name=log_event,json=logEvent,proto3,oneof"`
}

func (*Span_StartEvent) isSpan_Event() {}
func (*Span_EndEvent) isSpan_Event()   {}
func (*Span_LogEvent) isSpan_Event()   {}

func (m *Span) GetEvent() isSpan_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *Span) GetTraceContext() *Trace {
	if m != nil {
		return m.TraceContext
	}
	return nil
}

func (m *Span) GetSpanId() string {
	if m != nil {
		return m.SpanId
	}
	return ""
}

func (m *Span) GetStartEvent() *StartEvent {
	if x, ok := m.GetEvent().(*Span_StartEvent); ok {
		return x.StartEvent
	}
	return nil
}

func (m *Span) GetEndEvent() *EndEvent {
	if x, ok := m.GetEvent().(*Span_EndEvent); ok {
		return x.EndEvent
	}
	return nil
}

func (m *Span) GetLogEvent() *LogEvent {
	if x, ok := m.GetEvent().(*Span_LogEvent); ok {
		return x.LogEvent
	}
	return nil
}

func (m *Span) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Span) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *Span) GetEventLocation() string {
	if m != nil {
		return m.EventLocation
	}
	return ""
}

func (m *Span) GetParentSpanId() string {
	if m != nil {
		return m.ParentSpanId
	}
	return ""
}

func (m *Span) GetInternalSpanRefNumber() uint64 {
	if m != nil {
		return m.InternalSpanRefNumber
	}
	return 0
}

func (m *Span) GetInternalMeta() []byte {
	if m != nil {
		return m.InternalMeta
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Span) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Span_OneofMarshaler, _Span_OneofUnmarshaler, _Span_OneofSizer, []interface{}{
		(*Span_StartEvent)(nil),
		(*Span_EndEvent)(nil),
		(*Span_LogEvent)(nil),
	}
}

func _Span_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Span)
	// event
	switch x := m.Event.(type) {
	case *Span_StartEvent:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.StartEvent); err != nil {
			return err
		}
	case *Span_EndEvent:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.EndEvent); err != nil {
			return err
		}
	case *Span_LogEvent:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.LogEvent); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Span.Event has unexpected type %T", x)
	}
	return nil
}

func _Span_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Span)
	switch tag {
	case 3: // event.start_event
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StartEvent)
		err := b.DecodeMessage(msg)
		m.Event = &Span_StartEvent{msg}
		return true, err
	case 4: // event.end_event
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EndEvent)
		err := b.DecodeMessage(msg)
		m.Event = &Span_EndEvent{msg}
		return true, err
	case 5: // event.log_event
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LogEvent)
		err := b.DecodeMessage(msg)
		m.Event = &Span_LogEvent{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Span_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Span)
	// event
	switch x := m.Event.(type) {
	case *Span_StartEvent:
		s := proto.Size(x.StartEvent)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Span_EndEvent:
		s := proto.Size(x.EndEvent)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Span_LogEvent:
		s := proto.Size(x.LogEvent)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Span)(nil), "com.thapovan.orion.proto.Span")
}

func init() { proto.RegisterFile("Span.proto", fileDescriptor_Span_334f9dc70dfb4c18) }

var fileDescriptor_Span_334f9dc70dfb4c18 = []byte{
	// 370 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x4f, 0x6b, 0xdb, 0x40,
	0x10, 0xc5, 0xab, 0x56, 0xfe, 0xa3, 0x91, 0xec, 0xc3, 0x42, 0xe9, 0x52, 0x0a, 0x75, 0x8d, 0x0d,
	0x3e, 0xe9, 0xd0, 0x5e, 0x7a, 0xad, 0x5b, 0xd3, 0x04, 0x1c, 0x13, 0xe4, 0x9c, 0x72, 0x11, 0x6b,
	0x69, 0xec, 0x08, 0xa4, 0x5d, 0xb1, 0xda, 0x98, 0x7c, 0x89, 0x40, 0xbe, 0x42, 0xee, 0xf9, 0x90,
	0x61, 0x47, 0x2b, 0xe7, 0x64, 0x72, 0xd2, 0xce, 0xdb, 0xf7, 0x7e, 0x3b, 0x7a, 0x00, 0xdb, 0x5a,
	0xc8, 0xb8, 0xd6, 0xca, 0x28, 0xc6, 0x33, 0x55, 0xc5, 0xe6, 0x4e, 0xd4, 0xea, 0x28, 0x64, 0xac,
	0x74, 0xa1, 0xdc, 0xcd, 0xd7, 0xf0, 0x46, 0x8b, 0x0c, 0xdd, 0x10, 0xad, 0x8e, 0x28, 0x4d, 0xd3,
	0x4e, 0xd3, 0x17, 0x1f, 0x7c, 0xcb, 0x60, 0xff, 0x60, 0x64, 0xac, 0x2b, 0xcd, 0x94, 0x34, 0xf8,
	0x60, 0xb8, 0x37, 0xf1, 0x16, 0xe1, 0xcf, 0xef, 0xf1, 0x39, 0x6a, 0x4c, 0xd0, 0x24, 0xa2, 0xd4,
	0xdf, 0x36, 0xc4, 0xbe, 0xc0, 0xa0, 0xa9, 0x85, 0x4c, 0x8b, 0x9c, 0x7f, 0x9c, 0x78, 0x8b, 0x20,
	0xe9, 0xdb, 0xf1, 0x32, 0x67, 0xff, 0x21, 0x6c, 0x8c, 0xd0, 0x26, 0x45, 0xfb, 0x3a, 0xff, 0x44,
	0xf0, 0xd9, 0x79, 0xf8, 0xd6, 0x9a, 0x69, 0xd3, 0x8b, 0x0f, 0x09, 0x34, 0xa7, 0x89, 0xfd, 0x81,
	0x00, 0x65, 0xee, 0x30, 0x3e, 0x61, 0xa6, 0xe7, 0x31, 0x2b, 0x99, 0x77, 0x90, 0x21, 0xba, 0xb3,
	0x45, 0x94, 0xea, 0xe0, 0x10, 0xbd, 0xf7, 0x10, 0x6b, 0x75, 0x38, 0x21, 0x4a, 0x77, 0x66, 0xdf,
	0x20, 0x30, 0x45, 0x85, 0x8d, 0x11, 0x55, 0xcd, 0xfb, 0x13, 0x6f, 0xe1, 0x27, 0x6f, 0x02, 0xfb,
	0x01, 0x51, 0x83, 0xfa, 0x58, 0x64, 0x98, 0x4a, 0x51, 0x21, 0x1f, 0x50, 0x15, 0xa1, 0xd3, 0x36,
	0xa2, 0x42, 0x36, 0x87, 0x31, 0xbd, 0x9f, 0x96, 0x2a, 0x13, 0xa6, 0x50, 0x92, 0x0f, 0xc9, 0x34,
	0x22, 0x75, 0xed, 0x44, 0x36, 0x83, 0x71, 0x2d, 0xb4, 0xf5, 0x75, 0xb5, 0x06, 0x64, 0x8b, 0x5a,
	0x75, 0xdb, 0x96, 0xfb, 0x1b, 0x78, 0x21, 0x0d, 0x6a, 0x29, 0xca, 0xd6, 0xa7, 0x71, 0x9f, 0xca,
	0xfb, 0x6a, 0x87, 0x9a, 0x3f, 0x6d, 0x68, 0xbb, 0xcf, 0x9d, 0xc1, 0x46, 0x12, 0xdc, 0x6f, 0xe8,
	0x96, 0xcd, 0x61, 0x74, 0x4a, 0x56, 0x68, 0x04, 0x7f, 0x7e, 0xb4, 0x7f, 0x13, 0x25, 0x51, 0x27,
	0x5f, 0xa1, 0x11, 0xcb, 0x01, 0xf4, 0x68, 0xaf, 0xe5, 0xf8, 0xda, 0xbb, 0x05, 0x6a, 0x87, 0xca,
	0xd9, 0xf5, 0xe9, 0xf3, 0xeb, 0x35, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x5d, 0x0a, 0x4c, 0x88, 0x02,
	0x00, 0x00,
}
