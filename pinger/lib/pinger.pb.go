// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.11.2
// source: pinger.proto

package pinger

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Reqest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Reqest) Reset() {
	*x = Reqest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pinger_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reqest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reqest) ProtoMessage() {}

func (x *Reqest) ProtoReflect() protoreflect.Message {
	mi := &file_pinger_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reqest.ProtoReflect.Descriptor instead.
func (*Reqest) Descriptor() ([]byte, []int) {
	return file_pinger_proto_rawDescGZIP(), []int{0}
}

func (x *Reqest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text  string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Count int32  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pinger_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_pinger_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_pinger_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Response) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_pinger_proto protoreflect.FileDescriptor

var file_pinger_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x70, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x22, 0x1c, 0x0a, 0x06, 0x52, 0x65, 0x71, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x22, 0x34, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x34, 0x0a, 0x06, 0x50, 0x69,
	0x6e, 0x67, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x2e, 0x70,
	0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70,
	0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pinger_proto_rawDescOnce sync.Once
	file_pinger_proto_rawDescData = file_pinger_proto_rawDesc
)

func file_pinger_proto_rawDescGZIP() []byte {
	file_pinger_proto_rawDescOnce.Do(func() {
		file_pinger_proto_rawDescData = protoimpl.X.CompressGZIP(file_pinger_proto_rawDescData)
	})
	return file_pinger_proto_rawDescData
}

var file_pinger_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pinger_proto_goTypes = []interface{}{
	(*Reqest)(nil),   // 0: pinger.Reqest
	(*Response)(nil), // 1: pinger.Response
}
var file_pinger_proto_depIdxs = []int32{
	0, // 0: pinger.Pinger.Ping:input_type -> pinger.Reqest
	1, // 1: pinger.Pinger.Ping:output_type -> pinger.Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pinger_proto_init() }
func file_pinger_proto_init() {
	if File_pinger_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pinger_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reqest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pinger_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pinger_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pinger_proto_goTypes,
		DependencyIndexes: file_pinger_proto_depIdxs,
		MessageInfos:      file_pinger_proto_msgTypes,
	}.Build()
	File_pinger_proto = out.File
	file_pinger_proto_rawDesc = nil
	file_pinger_proto_goTypes = nil
	file_pinger_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PingerClient is the client API for Pinger service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PingerClient interface {
	Ping(ctx context.Context, in *Reqest, opts ...grpc.CallOption) (*Response, error)
}

type pingerClient struct {
	cc grpc.ClientConnInterface
}

func NewPingerClient(cc grpc.ClientConnInterface) PingerClient {
	return &pingerClient{cc}
}

func (c *pingerClient) Ping(ctx context.Context, in *Reqest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pinger.Pinger/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PingerServer is the server API for Pinger service.
type PingerServer interface {
	Ping(context.Context, *Reqest) (*Response, error)
}

// UnimplementedPingerServer can be embedded to have forward compatible implementations.
type UnimplementedPingerServer struct {
}

func (*UnimplementedPingerServer) Ping(context.Context, *Reqest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

func RegisterPingerServer(s *grpc.Server, srv PingerServer) {
	s.RegisterService(&_Pinger_serviceDesc, srv)
}

func _Pinger_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Reqest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingerServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pinger.Pinger/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingerServer).Ping(ctx, req.(*Reqest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Pinger_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pinger.Pinger",
	HandlerType: (*PingerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Pinger_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pinger.proto",
}