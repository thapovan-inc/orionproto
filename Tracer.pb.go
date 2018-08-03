// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Tracer.proto

package orionproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ControlRequest_RequestType int32

const (
	ControlRequest_AUTH       ControlRequest_RequestType = 0
	ControlRequest_END_STREAM ControlRequest_RequestType = 1
)

var ControlRequest_RequestType_name = map[int32]string{
	0: "AUTH",
	1: "END_STREAM",
}
var ControlRequest_RequestType_value = map[string]int32{
	"AUTH":       0,
	"END_STREAM": 1,
}

func (x ControlRequest_RequestType) String() string {
	return proto.EnumName(ControlRequest_RequestType_name, int32(x))
}
func (ControlRequest_RequestType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_Tracer_d7df53067f5d2587, []int{1, 0}
}

type ServerResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerResponse) Reset()         { *m = ServerResponse{} }
func (m *ServerResponse) String() string { return proto.CompactTextString(m) }
func (*ServerResponse) ProtoMessage()    {}
func (*ServerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_Tracer_d7df53067f5d2587, []int{0}
}
func (m *ServerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerResponse.Unmarshal(m, b)
}
func (m *ServerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerResponse.Marshal(b, m, deterministic)
}
func (dst *ServerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerResponse.Merge(dst, src)
}
func (m *ServerResponse) XXX_Size() int {
	return xxx_messageInfo_ServerResponse.Size(m)
}
func (m *ServerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ServerResponse proto.InternalMessageInfo

func (m *ServerResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *ServerResponse) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *ServerResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ControlRequest struct {
	RequestType ControlRequest_RequestType `protobuf:"varint,1,opt,name=request_type,json=requestType,proto3,enum=com.thapovan.orion.proto.ControlRequest_RequestType" json:"request_type,omitempty"`
	// Types that are valid to be assigned to Params:
	//	*ControlRequest_JsonString
	Params               isControlRequest_Params `protobuf_oneof:"params"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ControlRequest) Reset()         { *m = ControlRequest{} }
func (m *ControlRequest) String() string { return proto.CompactTextString(m) }
func (*ControlRequest) ProtoMessage()    {}
func (*ControlRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_Tracer_d7df53067f5d2587, []int{1}
}
func (m *ControlRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ControlRequest.Unmarshal(m, b)
}
func (m *ControlRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ControlRequest.Marshal(b, m, deterministic)
}
func (dst *ControlRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControlRequest.Merge(dst, src)
}
func (m *ControlRequest) XXX_Size() int {
	return xxx_messageInfo_ControlRequest.Size(m)
}
func (m *ControlRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ControlRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ControlRequest proto.InternalMessageInfo

type isControlRequest_Params interface {
	isControlRequest_Params()
}

type ControlRequest_JsonString struct {
	JsonString string `protobuf:"bytes,3,opt,name=jsonString,proto3,oneof"`
}

func (*ControlRequest_JsonString) isControlRequest_Params() {}

func (m *ControlRequest) GetParams() isControlRequest_Params {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *ControlRequest) GetRequestType() ControlRequest_RequestType {
	if m != nil {
		return m.RequestType
	}
	return ControlRequest_AUTH
}

func (m *ControlRequest) GetJsonString() string {
	if x, ok := m.GetParams().(*ControlRequest_JsonString); ok {
		return x.JsonString
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ControlRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ControlRequest_OneofMarshaler, _ControlRequest_OneofUnmarshaler, _ControlRequest_OneofSizer, []interface{}{
		(*ControlRequest_JsonString)(nil),
	}
}

func _ControlRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ControlRequest)
	// params
	switch x := m.Params.(type) {
	case *ControlRequest_JsonString:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.JsonString)
	case nil:
	default:
		return fmt.Errorf("ControlRequest.Params has unexpected type %T", x)
	}
	return nil
}

func _ControlRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ControlRequest)
	switch tag {
	case 3: // params.jsonString
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Params = &ControlRequest_JsonString{x}
		return true, err
	default:
		return false, nil
	}
}

func _ControlRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ControlRequest)
	// params
	switch x := m.Params.(type) {
	case *ControlRequest_JsonString:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.JsonString)))
		n += len(x.JsonString)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type UnaryRequest struct {
	AuthToken            string   `protobuf:"bytes,1,opt,name=auth_token,json=authToken,proto3" json:"auth_token,omitempty"`
	SpanData             *Span    `protobuf:"bytes,2,opt,name=span_data,json=spanData,proto3" json:"span_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnaryRequest) Reset()         { *m = UnaryRequest{} }
func (m *UnaryRequest) String() string { return proto.CompactTextString(m) }
func (*UnaryRequest) ProtoMessage()    {}
func (*UnaryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_Tracer_d7df53067f5d2587, []int{2}
}
func (m *UnaryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnaryRequest.Unmarshal(m, b)
}
func (m *UnaryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnaryRequest.Marshal(b, m, deterministic)
}
func (dst *UnaryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnaryRequest.Merge(dst, src)
}
func (m *UnaryRequest) XXX_Size() int {
	return xxx_messageInfo_UnaryRequest.Size(m)
}
func (m *UnaryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UnaryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UnaryRequest proto.InternalMessageInfo

func (m *UnaryRequest) GetAuthToken() string {
	if m != nil {
		return m.AuthToken
	}
	return ""
}

func (m *UnaryRequest) GetSpanData() *Span {
	if m != nil {
		return m.SpanData
	}
	return nil
}

type StreamRequest struct {
	// Types that are valid to be assigned to Request:
	//	*StreamRequest_ControlRequest
	//	*StreamRequest_SpanData
	Request              isStreamRequest_Request `protobuf_oneof:"request"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *StreamRequest) Reset()         { *m = StreamRequest{} }
func (m *StreamRequest) String() string { return proto.CompactTextString(m) }
func (*StreamRequest) ProtoMessage()    {}
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_Tracer_d7df53067f5d2587, []int{3}
}
func (m *StreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamRequest.Unmarshal(m, b)
}
func (m *StreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamRequest.Marshal(b, m, deterministic)
}
func (dst *StreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamRequest.Merge(dst, src)
}
func (m *StreamRequest) XXX_Size() int {
	return xxx_messageInfo_StreamRequest.Size(m)
}
func (m *StreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamRequest proto.InternalMessageInfo

type isStreamRequest_Request interface {
	isStreamRequest_Request()
}

type StreamRequest_ControlRequest struct {
	ControlRequest *ControlRequest `protobuf:"bytes,1,opt,name=control_request,json=controlRequest,proto3,oneof"`
}
type StreamRequest_SpanData struct {
	SpanData *Span `protobuf:"bytes,2,opt,name=span_data,json=spanData,proto3,oneof"`
}

func (*StreamRequest_ControlRequest) isStreamRequest_Request() {}
func (*StreamRequest_SpanData) isStreamRequest_Request()       {}

func (m *StreamRequest) GetRequest() isStreamRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *StreamRequest) GetControlRequest() *ControlRequest {
	if x, ok := m.GetRequest().(*StreamRequest_ControlRequest); ok {
		return x.ControlRequest
	}
	return nil
}

func (m *StreamRequest) GetSpanData() *Span {
	if x, ok := m.GetRequest().(*StreamRequest_SpanData); ok {
		return x.SpanData
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*StreamRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _StreamRequest_OneofMarshaler, _StreamRequest_OneofUnmarshaler, _StreamRequest_OneofSizer, []interface{}{
		(*StreamRequest_ControlRequest)(nil),
		(*StreamRequest_SpanData)(nil),
	}
}

func _StreamRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*StreamRequest)
	// request
	switch x := m.Request.(type) {
	case *StreamRequest_ControlRequest:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ControlRequest); err != nil {
			return err
		}
	case *StreamRequest_SpanData:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SpanData); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("StreamRequest.Request has unexpected type %T", x)
	}
	return nil
}

func _StreamRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*StreamRequest)
	switch tag {
	case 1: // request.control_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ControlRequest)
		err := b.DecodeMessage(msg)
		m.Request = &StreamRequest_ControlRequest{msg}
		return true, err
	case 2: // request.span_data
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Span)
		err := b.DecodeMessage(msg)
		m.Request = &StreamRequest_SpanData{msg}
		return true, err
	default:
		return false, nil
	}
}

func _StreamRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*StreamRequest)
	// request
	switch x := m.Request.(type) {
	case *StreamRequest_ControlRequest:
		s := proto.Size(x.ControlRequest)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *StreamRequest_SpanData:
		s := proto.Size(x.SpanData)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type StreamResponse struct {
	// Types that are valid to be assigned to Response:
	//	*StreamResponse_ServerResponse
	Response             isStreamResponse_Response `protobuf_oneof:"response"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *StreamResponse) Reset()         { *m = StreamResponse{} }
func (m *StreamResponse) String() string { return proto.CompactTextString(m) }
func (*StreamResponse) ProtoMessage()    {}
func (*StreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_Tracer_d7df53067f5d2587, []int{4}
}
func (m *StreamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamResponse.Unmarshal(m, b)
}
func (m *StreamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamResponse.Marshal(b, m, deterministic)
}
func (dst *StreamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamResponse.Merge(dst, src)
}
func (m *StreamResponse) XXX_Size() int {
	return xxx_messageInfo_StreamResponse.Size(m)
}
func (m *StreamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamResponse proto.InternalMessageInfo

type isStreamResponse_Response interface {
	isStreamResponse_Response()
}

type StreamResponse_ServerResponse struct {
	ServerResponse *ServerResponse `protobuf:"bytes,1,opt,name=server_response,json=serverResponse,proto3,oneof"`
}

func (*StreamResponse_ServerResponse) isStreamResponse_Response() {}

func (m *StreamResponse) GetResponse() isStreamResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *StreamResponse) GetServerResponse() *ServerResponse {
	if x, ok := m.GetResponse().(*StreamResponse_ServerResponse); ok {
		return x.ServerResponse
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*StreamResponse) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _StreamResponse_OneofMarshaler, _StreamResponse_OneofUnmarshaler, _StreamResponse_OneofSizer, []interface{}{
		(*StreamResponse_ServerResponse)(nil),
	}
}

func _StreamResponse_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*StreamResponse)
	// response
	switch x := m.Response.(type) {
	case *StreamResponse_ServerResponse:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ServerResponse); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("StreamResponse.Response has unexpected type %T", x)
	}
	return nil
}

func _StreamResponse_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*StreamResponse)
	switch tag {
	case 1: // response.server_response
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ServerResponse)
		err := b.DecodeMessage(msg)
		m.Response = &StreamResponse_ServerResponse{msg}
		return true, err
	default:
		return false, nil
	}
}

func _StreamResponse_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*StreamResponse)
	// response
	switch x := m.Response.(type) {
	case *StreamResponse_ServerResponse:
		s := proto.Size(x.ServerResponse)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type BulkRequest struct {
	AuthToken            string   `protobuf:"bytes,1,opt,name=auth_token,json=authToken,proto3" json:"auth_token,omitempty"`
	SpanData             []*Span  `protobuf:"bytes,2,rep,name=span_data,json=spanData,proto3" json:"span_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BulkRequest) Reset()         { *m = BulkRequest{} }
func (m *BulkRequest) String() string { return proto.CompactTextString(m) }
func (*BulkRequest) ProtoMessage()    {}
func (*BulkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_Tracer_d7df53067f5d2587, []int{5}
}
func (m *BulkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BulkRequest.Unmarshal(m, b)
}
func (m *BulkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BulkRequest.Marshal(b, m, deterministic)
}
func (dst *BulkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BulkRequest.Merge(dst, src)
}
func (m *BulkRequest) XXX_Size() int {
	return xxx_messageInfo_BulkRequest.Size(m)
}
func (m *BulkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BulkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BulkRequest proto.InternalMessageInfo

func (m *BulkRequest) GetAuthToken() string {
	if m != nil {
		return m.AuthToken
	}
	return ""
}

func (m *BulkRequest) GetSpanData() []*Span {
	if m != nil {
		return m.SpanData
	}
	return nil
}

func init() {
	proto.RegisterType((*ServerResponse)(nil), "com.thapovan.orion.proto.ServerResponse")
	proto.RegisterType((*ControlRequest)(nil), "com.thapovan.orion.proto.ControlRequest")
	proto.RegisterType((*UnaryRequest)(nil), "com.thapovan.orion.proto.UnaryRequest")
	proto.RegisterType((*StreamRequest)(nil), "com.thapovan.orion.proto.StreamRequest")
	proto.RegisterType((*StreamResponse)(nil), "com.thapovan.orion.proto.StreamResponse")
	proto.RegisterType((*BulkRequest)(nil), "com.thapovan.orion.proto.BulkRequest")
	proto.RegisterEnum("com.thapovan.orion.proto.ControlRequest_RequestType", ControlRequest_RequestType_name, ControlRequest_RequestType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TracerClient is the client API for Tracer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TracerClient interface {
	UploadSpanStream(ctx context.Context, opts ...grpc.CallOption) (Tracer_UploadSpanStreamClient, error)
	UploadSpan(ctx context.Context, in *UnaryRequest, opts ...grpc.CallOption) (*ServerResponse, error)
	UploadSpanBulk(ctx context.Context, in *BulkRequest, opts ...grpc.CallOption) (*ServerResponse, error)
}

type tracerClient struct {
	cc *grpc.ClientConn
}

func NewTracerClient(cc *grpc.ClientConn) TracerClient {
	return &tracerClient{cc}
}

func (c *tracerClient) UploadSpanStream(ctx context.Context, opts ...grpc.CallOption) (Tracer_UploadSpanStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Tracer_serviceDesc.Streams[0], "/com.thapovan.orion.proto.Tracer/UploadSpanStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &tracerUploadSpanStreamClient{stream}
	return x, nil
}

type Tracer_UploadSpanStreamClient interface {
	Send(*StreamRequest) error
	Recv() (*ServerResponse, error)
	grpc.ClientStream
}

type tracerUploadSpanStreamClient struct {
	grpc.ClientStream
}

func (x *tracerUploadSpanStreamClient) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tracerUploadSpanStreamClient) Recv() (*ServerResponse, error) {
	m := new(ServerResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tracerClient) UploadSpan(ctx context.Context, in *UnaryRequest, opts ...grpc.CallOption) (*ServerResponse, error) {
	out := new(ServerResponse)
	err := c.cc.Invoke(ctx, "/com.thapovan.orion.proto.Tracer/UploadSpan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tracerClient) UploadSpanBulk(ctx context.Context, in *BulkRequest, opts ...grpc.CallOption) (*ServerResponse, error) {
	out := new(ServerResponse)
	err := c.cc.Invoke(ctx, "/com.thapovan.orion.proto.Tracer/UploadSpanBulk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TracerServer is the server API for Tracer service.
type TracerServer interface {
	UploadSpanStream(Tracer_UploadSpanStreamServer) error
	UploadSpan(context.Context, *UnaryRequest) (*ServerResponse, error)
	UploadSpanBulk(context.Context, *BulkRequest) (*ServerResponse, error)
}

func RegisterTracerServer(s *grpc.Server, srv TracerServer) {
	s.RegisterService(&_Tracer_serviceDesc, srv)
}

func _Tracer_UploadSpanStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TracerServer).UploadSpanStream(&tracerUploadSpanStreamServer{stream})
}

type Tracer_UploadSpanStreamServer interface {
	Send(*ServerResponse) error
	Recv() (*StreamRequest, error)
	grpc.ServerStream
}

type tracerUploadSpanStreamServer struct {
	grpc.ServerStream
}

func (x *tracerUploadSpanStreamServer) Send(m *ServerResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tracerUploadSpanStreamServer) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Tracer_UploadSpan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TracerServer).UploadSpan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.thapovan.orion.proto.Tracer/UploadSpan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TracerServer).UploadSpan(ctx, req.(*UnaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tracer_UploadSpanBulk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BulkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TracerServer).UploadSpanBulk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.thapovan.orion.proto.Tracer/UploadSpanBulk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TracerServer).UploadSpanBulk(ctx, req.(*BulkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Tracer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.thapovan.orion.proto.Tracer",
	HandlerType: (*TracerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadSpan",
			Handler:    _Tracer_UploadSpan_Handler,
		},
		{
			MethodName: "UploadSpanBulk",
			Handler:    _Tracer_UploadSpanBulk_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadSpanStream",
			Handler:       _Tracer_UploadSpanStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "Tracer.proto",
}

func init() { proto.RegisterFile("Tracer.proto", fileDescriptor_Tracer_d7df53067f5d2587) }

var fileDescriptor_Tracer_d7df53067f5d2587 = []byte{
	// 468 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xcd, 0xb6, 0x55, 0x6a, 0x8f, 0x83, 0x1b, 0xed, 0xc9, 0xaa, 0x04, 0x8a, 0x2c, 0x41, 0x73,
	0xb2, 0x50, 0xe0, 0x86, 0x38, 0xd4, 0xb4, 0x52, 0x2e, 0x20, 0xb4, 0x76, 0x04, 0xe2, 0x80, 0x35,
	0x38, 0xab, 0x36, 0x6d, 0xb2, 0xbb, 0xdd, 0xdd, 0x54, 0xca, 0x4f, 0x71, 0xe4, 0xc8, 0xb7, 0x21,
	0xaf, 0x6d, 0xc5, 0x39, 0x04, 0x52, 0x4e, 0x9e, 0x79, 0xb3, 0x33, 0xf3, 0x66, 0xde, 0x18, 0x06,
	0xb9, 0xc6, 0x92, 0xeb, 0x44, 0x69, 0x69, 0x25, 0x8d, 0x4a, 0xb9, 0x4a, 0xec, 0x2d, 0x2a, 0xf9,
	0x88, 0x22, 0x91, 0x7a, 0x21, 0x45, 0x1d, 0x39, 0x87, 0x4c, 0x61, 0x63, 0xc7, 0x5f, 0x21, 0xcc,
	0xb8, 0x7e, 0xe4, 0x9a, 0x71, 0xa3, 0xa4, 0x30, 0x9c, 0x46, 0x70, 0x6a, 0xd6, 0x65, 0xc9, 0x8d,
	0x89, 0xc8, 0x88, 0x8c, 0x3d, 0xd6, 0xba, 0x94, 0xc2, 0x49, 0x29, 0xe7, 0x3c, 0x3a, 0x1a, 0x91,
	0xb1, 0xcf, 0x9c, 0x5d, 0xbd, 0x5e, 0x71, 0x63, 0xf0, 0x86, 0x47, 0xc7, 0x0e, 0x6e, 0xdd, 0xf8,
	0x37, 0x81, 0xf0, 0x83, 0x14, 0x56, 0xcb, 0x25, 0xe3, 0x0f, 0x6b, 0x6e, 0x2c, 0xfd, 0x02, 0x03,
	0x5d, 0x9b, 0x85, 0xdd, 0x28, 0xee, 0xea, 0x87, 0x93, 0xb7, 0xc9, 0x3e, 0xa6, 0xc9, 0x6e, 0x7e,
	0xd2, 0x7c, 0xf3, 0x8d, 0xe2, 0x2c, 0xd0, 0x5b, 0x87, 0x8e, 0x00, 0xee, 0x8c, 0x14, 0x99, 0xd5,
	0x0b, 0x71, 0x53, 0x13, 0x99, 0xf6, 0x58, 0x07, 0x8b, 0x2f, 0x20, 0xe8, 0x64, 0x53, 0x0f, 0x4e,
	0x2e, 0x67, 0xf9, 0x74, 0xd8, 0xa3, 0x21, 0xc0, 0xf5, 0xa7, 0xab, 0x22, 0xcb, 0xd9, 0xf5, 0xe5,
	0xc7, 0x21, 0x49, 0x3d, 0xe8, 0x2b, 0xd4, 0xb8, 0x32, 0xf1, 0x1d, 0x0c, 0x66, 0x02, 0xf5, 0xa6,
	0x65, 0xff, 0x1c, 0x00, 0xd7, 0xf6, 0xb6, 0xb0, 0xf2, 0x9e, 0x0b, 0xc7, 0xdd, 0x67, 0x7e, 0x85,
	0xe4, 0x15, 0x40, 0xdf, 0x81, 0x6f, 0x14, 0x8a, 0x62, 0x8e, 0x16, 0xdd, 0x8a, 0x82, 0xc9, 0x8b,
	0xfd, 0x93, 0x55, 0x12, 0x30, 0xaf, 0x4a, 0xb8, 0x42, 0x8b, 0xf1, 0x4f, 0x02, 0xcf, 0x32, 0xab,
	0x39, 0xae, 0xda, 0x6e, 0x19, 0x9c, 0x95, 0xf5, 0xf4, 0x45, 0x33, 0xa9, 0x6b, 0x19, 0x4c, 0xc6,
	0x87, 0xae, 0x6b, 0xda, 0x63, 0x61, 0xb9, 0x2b, 0xc0, 0xfb, 0x27, 0x73, 0x9c, 0xf6, 0xb6, 0x2c,
	0x53, 0x1f, 0x4e, 0x1b, 0x2e, 0xf1, 0x03, 0x84, 0x2d, 0xdf, 0xe6, 0x6e, 0x32, 0x38, 0x33, 0xee,
	0x92, 0x0a, 0xdd, 0x40, 0xff, 0x26, 0xbc, 0x7b, 0x7a, 0x15, 0x61, 0xb3, 0x83, 0xa4, 0x00, 0x5e,
	0x5b, 0x2d, 0x5e, 0x40, 0x90, 0xae, 0x97, 0xf7, 0xff, 0x27, 0xc7, 0xf1, 0x53, 0xe4, 0x98, 0xfc,
	0x3a, 0x82, 0x7e, 0xfd, 0x33, 0xd1, 0x05, 0x0c, 0x67, 0x6a, 0x29, 0x71, 0x5e, 0x3d, 0xa9, 0x47,
	0xa6, 0x17, 0x7f, 0x29, 0xd4, 0x15, 0xf1, 0xfc, 0xe0, 0xd1, 0xc7, 0xe4, 0x35, 0xa1, 0xdf, 0x01,
	0xb6, 0xad, 0xe8, 0xab, 0xfd, 0xb9, 0xdd, 0xb3, 0x3c, 0xbc, 0x07, 0x45, 0x08, 0xb7, 0xf5, 0xab,
	0x55, 0xd2, 0x97, 0xfb, 0x73, 0x3b, 0xab, 0x3e, 0xbc, 0x45, 0x1a, 0x7e, 0x26, 0xdf, 0xc0, 0xc5,
	0x5d, 0xf8, 0x47, 0xdf, 0x7d, 0xde, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0xad, 0x82, 0xee, 0xb6,
	0x9b, 0x04, 0x00, 0x00,
}