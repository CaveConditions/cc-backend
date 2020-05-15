// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.4
// source: pkg/api/api.proto

package api

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

type Cave struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string  `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	CountryName string  `protobuf:"bytes,2,opt,name=countryName,proto3" json:"countryName,omitempty"`
	RegionName  string  `protobuf:"bytes,3,opt,name=regionName,proto3" json:"regionName,omitempty"`
	Longitude   float32 `protobuf:"fixed32,4,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Latitude    float32 `protobuf:"fixed32,5,opt,name=latitude,proto3" json:"latitude,omitempty"`
}

func (x *Cave) Reset() {
	*x = Cave{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cave) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cave) ProtoMessage() {}

func (x *Cave) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cave.ProtoReflect.Descriptor instead.
func (*Cave) Descriptor() ([]byte, []int) {
	return file_pkg_api_api_proto_rawDescGZIP(), []int{0}
}

func (x *Cave) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Cave) GetCountryName() string {
	if x != nil {
		return x.CountryName
	}
	return ""
}

func (x *Cave) GetRegionName() string {
	if x != nil {
		return x.RegionName
	}
	return ""
}

func (x *Cave) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *Cave) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

type AddCaveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cave *Cave `protobuf:"bytes,1,opt,name=cave,proto3" json:"cave,omitempty"`
}

func (x *AddCaveRequest) Reset() {
	*x = AddCaveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCaveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCaveRequest) ProtoMessage() {}

func (x *AddCaveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCaveRequest.ProtoReflect.Descriptor instead.
func (*AddCaveRequest) Descriptor() ([]byte, []int) {
	return file_pkg_api_api_proto_rawDescGZIP(), []int{1}
}

func (x *AddCaveRequest) GetCave() *Cave {
	if x != nil {
		return x.Cave
	}
	return nil
}

type AddCaveReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddCaveReply) Reset() {
	*x = AddCaveReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCaveReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCaveReply) ProtoMessage() {}

func (x *AddCaveReply) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCaveReply.ProtoReflect.Descriptor instead.
func (*AddCaveReply) Descriptor() ([]byte, []int) {
	return file_pkg_api_api_proto_rawDescGZIP(), []int{2}
}

type GetCaveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *GetCaveRequest) Reset() {
	*x = GetCaveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCaveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCaveRequest) ProtoMessage() {}

func (x *GetCaveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCaveRequest.ProtoReflect.Descriptor instead.
func (*GetCaveRequest) Descriptor() ([]byte, []int) {
	return file_pkg_api_api_proto_rawDescGZIP(), []int{3}
}

func (x *GetCaveRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type GetCaveReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cave *Cave `protobuf:"bytes,1,opt,name=cave,proto3" json:"cave,omitempty"`
}

func (x *GetCaveReply) Reset() {
	*x = GetCaveReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCaveReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCaveReply) ProtoMessage() {}

func (x *GetCaveReply) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCaveReply.ProtoReflect.Descriptor instead.
func (*GetCaveReply) Descriptor() ([]byte, []int) {
	return file_pkg_api_api_proto_rawDescGZIP(), []int{4}
}

func (x *GetCaveReply) GetCave() *Cave {
	if x != nil {
		return x.Cave
	}
	return nil
}

var File_pkg_api_api_proto protoreflect.FileDescriptor

var file_pkg_api_api_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x98, 0x01, 0x0a, 0x04, 0x43, 0x61, 0x76,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e,
	0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x6c, 0x6f,
	0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x22, 0x2f, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x43, 0x61, 0x76, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x04, 0x63, 0x61, 0x76, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x61, 0x76, 0x65, 0x52, 0x04,
	0x63, 0x61, 0x76, 0x65, 0x22, 0x0e, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x43, 0x61, 0x76, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x26, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x76, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x2d, 0x0a, 0x0c,
	0x47, 0x65, 0x74, 0x43, 0x61, 0x76, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1d, 0x0a, 0x04,
	0x63, 0x61, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x43, 0x61, 0x76, 0x65, 0x52, 0x04, 0x63, 0x61, 0x76, 0x65, 0x32, 0x7a, 0x0a, 0x0e, 0x43,
	0x61, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x33, 0x0a,
	0x07, 0x41, 0x64, 0x64, 0x43, 0x61, 0x76, 0x65, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41,
	0x64, 0x64, 0x43, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x61, 0x76, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x33, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x43, 0x61, 0x76, 0x65, 0x12, 0x13, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x76, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_api_api_proto_rawDescOnce sync.Once
	file_pkg_api_api_proto_rawDescData = file_pkg_api_api_proto_rawDesc
)

func file_pkg_api_api_proto_rawDescGZIP() []byte {
	file_pkg_api_api_proto_rawDescOnce.Do(func() {
		file_pkg_api_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_api_api_proto_rawDescData)
	})
	return file_pkg_api_api_proto_rawDescData
}

var file_pkg_api_api_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_api_api_proto_goTypes = []interface{}{
	(*Cave)(nil),           // 0: api.Cave
	(*AddCaveRequest)(nil), // 1: api.AddCaveRequest
	(*AddCaveReply)(nil),   // 2: api.AddCaveReply
	(*GetCaveRequest)(nil), // 3: api.GetCaveRequest
	(*GetCaveReply)(nil),   // 4: api.GetCaveReply
}
var file_pkg_api_api_proto_depIdxs = []int32{
	0, // 0: api.AddCaveRequest.cave:type_name -> api.Cave
	0, // 1: api.GetCaveReply.cave:type_name -> api.Cave
	1, // 2: api.CaveConditions.AddCave:input_type -> api.AddCaveRequest
	3, // 3: api.CaveConditions.GetCave:input_type -> api.GetCaveRequest
	2, // 4: api.CaveConditions.AddCave:output_type -> api.AddCaveReply
	4, // 5: api.CaveConditions.GetCave:output_type -> api.GetCaveReply
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_api_api_proto_init() }
func file_pkg_api_api_proto_init() {
	if File_pkg_api_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_api_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cave); i {
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
		file_pkg_api_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCaveRequest); i {
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
		file_pkg_api_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCaveReply); i {
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
		file_pkg_api_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCaveRequest); i {
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
		file_pkg_api_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCaveReply); i {
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
			RawDescriptor: file_pkg_api_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_api_api_proto_goTypes,
		DependencyIndexes: file_pkg_api_api_proto_depIdxs,
		MessageInfos:      file_pkg_api_api_proto_msgTypes,
	}.Build()
	File_pkg_api_api_proto = out.File
	file_pkg_api_api_proto_rawDesc = nil
	file_pkg_api_api_proto_goTypes = nil
	file_pkg_api_api_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CaveConditionsClient is the client API for CaveConditions service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CaveConditionsClient interface {
	AddCave(ctx context.Context, in *AddCaveRequest, opts ...grpc.CallOption) (*AddCaveReply, error)
	GetCave(ctx context.Context, in *GetCaveRequest, opts ...grpc.CallOption) (*GetCaveReply, error)
}

type caveConditionsClient struct {
	cc grpc.ClientConnInterface
}

func NewCaveConditionsClient(cc grpc.ClientConnInterface) CaveConditionsClient {
	return &caveConditionsClient{cc}
}

func (c *caveConditionsClient) AddCave(ctx context.Context, in *AddCaveRequest, opts ...grpc.CallOption) (*AddCaveReply, error) {
	out := new(AddCaveReply)
	err := c.cc.Invoke(ctx, "/api.CaveConditions/AddCave", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *caveConditionsClient) GetCave(ctx context.Context, in *GetCaveRequest, opts ...grpc.CallOption) (*GetCaveReply, error) {
	out := new(GetCaveReply)
	err := c.cc.Invoke(ctx, "/api.CaveConditions/GetCave", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CaveConditionsServer is the server API for CaveConditions service.
type CaveConditionsServer interface {
	AddCave(context.Context, *AddCaveRequest) (*AddCaveReply, error)
	GetCave(context.Context, *GetCaveRequest) (*GetCaveReply, error)
}

// UnimplementedCaveConditionsServer can be embedded to have forward compatible implementations.
type UnimplementedCaveConditionsServer struct {
}

func (*UnimplementedCaveConditionsServer) AddCave(context.Context, *AddCaveRequest) (*AddCaveReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCave not implemented")
}
func (*UnimplementedCaveConditionsServer) GetCave(context.Context, *GetCaveRequest) (*GetCaveReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCave not implemented")
}

func RegisterCaveConditionsServer(s *grpc.Server, srv CaveConditionsServer) {
	s.RegisterService(&_CaveConditions_serviceDesc, srv)
}

func _CaveConditions_AddCave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaveConditionsServer).AddCave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CaveConditions/AddCave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaveConditionsServer).AddCave(ctx, req.(*AddCaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaveConditions_GetCave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaveConditionsServer).GetCave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CaveConditions/GetCave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaveConditionsServer).GetCave(ctx, req.(*GetCaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CaveConditions_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.CaveConditions",
	HandlerType: (*CaveConditionsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCave",
			Handler:    _CaveConditions_AddCave_Handler,
		},
		{
			MethodName: "GetCave",
			Handler:    _CaveConditions_GetCave_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/api/api.proto",
}
