package orionproto

import (
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"io"
)

var marshaler = &jsonpb.Marshaler {
	EmitDefaults: true,
	OrigName: true,
}

var unMarshaler = &jsonpb.Unmarshaler{
	AllowUnknownFields: true,
}

func ProtoToJson(protoMessage proto.Message) (string,error)  {
	return marshaler.MarshalToString(protoMessage)
}

func JsonToProto(r io.Reader, protoMessage proto.Message) error  {
	return unMarshaler.Unmarshal(r,protoMessage)
}
