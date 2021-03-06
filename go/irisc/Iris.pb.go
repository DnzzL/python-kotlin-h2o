// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kotlin/src/main/proto/Iris.proto

/*
Package h2o is a generated protocol buffer package.

It is generated from these files:
	kotlin/src/main/proto/Iris.proto

It has these top-level messages:
	IrisRequest
	IrisReply
*/
package h2o

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

type IrisRequest struct {
	SepalLength float64 `protobuf:"fixed64,1,opt,name=SepalLength,json=sepalLength" json:"SepalLength,omitempty"`
	SepalWidth  float64 `protobuf:"fixed64,2,opt,name=SepalWidth,json=sepalWidth" json:"SepalWidth,omitempty"`
	PetalLength float64 `protobuf:"fixed64,3,opt,name=PetalLength,json=petalLength" json:"PetalLength,omitempty"`
	PetalWidth  float64 `protobuf:"fixed64,4,opt,name=PetalWidth,json=petalWidth" json:"PetalWidth,omitempty"`
}

func (m *IrisRequest) Reset()                    { *m = IrisRequest{} }
func (m *IrisRequest) String() string            { return proto.CompactTextString(m) }
func (*IrisRequest) ProtoMessage()               {}
func (*IrisRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IrisRequest) GetSepalLength() float64 {
	if m != nil {
		return m.SepalLength
	}
	return 0
}

func (m *IrisRequest) GetSepalWidth() float64 {
	if m != nil {
		return m.SepalWidth
	}
	return 0
}

func (m *IrisRequest) GetPetalLength() float64 {
	if m != nil {
		return m.PetalLength
	}
	return 0
}

func (m *IrisRequest) GetPetalWidth() float64 {
	if m != nil {
		return m.PetalWidth
	}
	return 0
}

type IrisReply struct {
	Species string `protobuf:"bytes,1,opt,name=species" json:"species,omitempty"`
}

func (m *IrisReply) Reset()                    { *m = IrisReply{} }
func (m *IrisReply) String() string            { return proto.CompactTextString(m) }
func (*IrisReply) ProtoMessage()               {}
func (*IrisReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *IrisReply) GetSpecies() string {
	if m != nil {
		return m.Species
	}
	return ""
}

func init() {
	proto.RegisterType((*IrisRequest)(nil), "h2o.IrisRequest")
	proto.RegisterType((*IrisReply)(nil), "h2o.IrisReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Predictor service

type PredictorClient interface {
	Predict(ctx context.Context, in *IrisRequest, opts ...grpc.CallOption) (*IrisReply, error)
}

type predictorClient struct {
	cc *grpc.ClientConn
}

func NewPredictorClient(cc *grpc.ClientConn) PredictorClient {
	return &predictorClient{cc}
}

func (c *predictorClient) Predict(ctx context.Context, in *IrisRequest, opts ...grpc.CallOption) (*IrisReply, error) {
	out := new(IrisReply)
	err := grpc.Invoke(ctx, "/h2o.Predictor/Predict", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Predictor service

type PredictorServer interface {
	Predict(context.Context, *IrisRequest) (*IrisReply, error)
}

func RegisterPredictorServer(s *grpc.Server, srv PredictorServer) {
	s.RegisterService(&_Predictor_serviceDesc, srv)
}

func _Predictor_Predict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IrisRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictorServer).Predict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/h2o.Predictor/Predict",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictorServer).Predict(ctx, req.(*IrisRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Predictor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "h2o.Predictor",
	HandlerType: (*PredictorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Predict",
			Handler:    _Predictor_Predict_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kotlin/src/main/proto/Iris.proto",
}

func init() { proto.RegisterFile("kotlin/src/main/proto/Iris.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xcd, 0x56, 0x5c, 0x3a, 0x05, 0x91, 0x1c, 0xa4, 0x78, 0x90, 0x52, 0x10, 0x04, 0xb1,
	0x85, 0x7a, 0xf3, 0xb8, 0x27, 0x85, 0x3d, 0x94, 0x7a, 0xd8, 0x73, 0x6d, 0x87, 0x4d, 0x30, 0xdb,
	0xc4, 0x64, 0x3c, 0xec, 0x63, 0xf8, 0x0a, 0x3e, 0xa9, 0x24, 0xd1, 0x6e, 0x6f, 0x99, 0x8f, 0x9f,
	0x2f, 0x33, 0x3f, 0x14, 0x1f, 0x9a, 0x94, 0x9c, 0x6a, 0x67, 0x87, 0xfa, 0xd0, 0xcb, 0xa9, 0x36,
	0x56, 0x93, 0xae, 0x5f, 0xad, 0x74, 0x55, 0x78, 0xf2, 0x44, 0x34, 0xba, 0xfc, 0x66, 0x90, 0x79,
	0xd6, 0xe1, 0xe7, 0x17, 0x3a, 0xe2, 0x05, 0x64, 0x6f, 0x68, 0x7a, 0xb5, 0xc5, 0x69, 0x4f, 0x22,
	0x67, 0x05, 0xbb, 0x67, 0x5d, 0xe6, 0x4e, 0x88, 0xdf, 0x02, 0x84, 0xc4, 0x4e, 0x8e, 0x24, 0xf2,
	0x55, 0x08, 0x80, 0x9b, 0x89, 0x37, 0xb4, 0x48, 0xb3, 0x21, 0x89, 0x06, 0x73, 0x42, 0xde, 0x10,
	0x12, 0xd1, 0x70, 0x1e, 0x0d, 0x66, 0x26, 0xe5, 0x1d, 0xa4, 0x71, 0x25, 0xa3, 0x8e, 0x3c, 0x87,
	0xb5, 0x33, 0x38, 0x48, 0x74, 0x61, 0x99, 0xb4, 0xfb, 0x1f, 0x9b, 0x67, 0x48, 0x5b, 0x8b, 0xa3,
	0x1c, 0x48, 0x5b, 0xfe, 0x08, 0xeb, 0xbf, 0x81, 0x5f, 0x55, 0xa2, 0xd1, 0xd5, 0xe2, 0xa8, 0x9b,
	0xcb, 0x05, 0x31, 0xea, 0x58, 0x9e, 0x6d, 0x1e, 0xe0, 0x9a, 0x70, 0x10, 0x15, 0x09, 0x7d, 0xe8,
	0x9d, 0xc2, 0xbd, 0xed, 0xa7, 0xd1, 0xa7, 0x36, 0xe1, 0xeb, 0xd6, 0x17, 0xd4, 0xb2, 0x9f, 0x55,
	0xf2, 0xb2, 0xdd, 0xbd, 0x5f, 0x84, 0xbe, 0x9e, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff, 0xa0, 0xe2,
	0xa4, 0x19, 0x53, 0x01, 0x00, 0x00,
}
