// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: proto/cities.proto

package cities

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

type EmptyInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyInput) Reset() {
	*x = EmptyInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cities_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyInput) ProtoMessage() {}

func (x *EmptyInput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cities_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyInput.ProtoReflect.Descriptor instead.
func (*EmptyInput) Descriptor() ([]byte, []int) {
	return file_proto_cities_proto_rawDescGZIP(), []int{0}
}

type GetCityInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCityInput) Reset() {
	*x = GetCityInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cities_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCityInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCityInput) ProtoMessage() {}

func (x *GetCityInput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cities_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCityInput.ProtoReflect.Descriptor instead.
func (*GetCityInput) Descriptor() ([]byte, []int) {
	return file_proto_cities_proto_rawDescGZIP(), []int{1}
}

func (x *GetCityInput) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type NewCityInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *NewCityInput) Reset() {
	*x = NewCityInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cities_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewCityInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewCityInput) ProtoMessage() {}

func (x *NewCityInput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cities_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewCityInput.ProtoReflect.Descriptor instead.
func (*NewCityInput) Descriptor() ([]byte, []int) {
	return file_proto_cities_proto_rawDescGZIP(), []int{2}
}

func (x *NewCityInput) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Cities struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cities []*City `protobuf:"bytes,1,rep,name=cities,proto3" json:"cities,omitempty"`
}

func (x *Cities) Reset() {
	*x = Cities{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cities_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cities) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cities) ProtoMessage() {}

func (x *Cities) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cities_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cities.ProtoReflect.Descriptor instead.
func (*Cities) Descriptor() ([]byte, []int) {
	return file_proto_cities_proto_rawDescGZIP(), []int{3}
}

func (x *Cities) GetCities() []*City {
	if x != nil {
		return x.Cities
	}
	return nil
}

type City struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *City) Reset() {
	*x = City{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cities_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *City) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*City) ProtoMessage() {}

func (x *City) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cities_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use City.ProtoReflect.Descriptor instead.
func (*City) Descriptor() ([]byte, []int) {
	return file_proto_cities_proto_rawDescGZIP(), []int{4}
}

func (x *City) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *City) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_proto_cities_proto protoreflect.FileDescriptor

var file_proto_cities_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0c, 0x0a, 0x0a, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x22, 0x1e, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x22, 0x0a, 0x0c, 0x4e, 0x65, 0x77, 0x43, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x27, 0x0a, 0x06, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x12, 0x1d, 0x0a, 0x06, 0x63, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x05, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x52, 0x06, 0x63, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22,
	0x2a, 0x0a, 0x04, 0x43, 0x69, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x7d, 0x0a, 0x0d, 0x43,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x43, 0x69, 0x74, 0x79, 0x12, 0x0d, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x69, 0x74,
	0x79, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x05, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x22, 0x00, 0x12,
	0x23, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x0b, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x07, 0x2e, 0x43, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x22, 0x00, 0x12, 0x24, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x69,
	0x74, 0x79, 0x12, 0x0d, 0x2e, 0x4e, 0x65, 0x77, 0x43, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x1a, 0x05, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x22, 0x00, 0x42, 0x14, 0x5a, 0x12, 0x77, 0x61,
	0x72, 0x65, 0x73, 0x69, 0x78, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_cities_proto_rawDescOnce sync.Once
	file_proto_cities_proto_rawDescData = file_proto_cities_proto_rawDesc
)

func file_proto_cities_proto_rawDescGZIP() []byte {
	file_proto_cities_proto_rawDescOnce.Do(func() {
		file_proto_cities_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_cities_proto_rawDescData)
	})
	return file_proto_cities_proto_rawDescData
}

var file_proto_cities_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_cities_proto_goTypes = []interface{}{
	(*EmptyInput)(nil),   // 0: EmptyInput
	(*GetCityInput)(nil), // 1: GetCityInput
	(*NewCityInput)(nil), // 2: NewCityInput
	(*Cities)(nil),       // 3: Cities
	(*City)(nil),         // 4: City
}
var file_proto_cities_proto_depIdxs = []int32{
	4, // 0: Cities.cities:type_name -> City
	1, // 1: CitiesService.GetCity:input_type -> GetCityInput
	0, // 2: CitiesService.GetCities:input_type -> EmptyInput
	2, // 3: CitiesService.CreateCity:input_type -> NewCityInput
	4, // 4: CitiesService.GetCity:output_type -> City
	3, // 5: CitiesService.GetCities:output_type -> Cities
	4, // 6: CitiesService.CreateCity:output_type -> City
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_cities_proto_init() }
func file_proto_cities_proto_init() {
	if File_proto_cities_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_cities_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyInput); i {
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
		file_proto_cities_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCityInput); i {
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
		file_proto_cities_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewCityInput); i {
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
		file_proto_cities_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cities); i {
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
		file_proto_cities_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*City); i {
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
			RawDescriptor: file_proto_cities_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_cities_proto_goTypes,
		DependencyIndexes: file_proto_cities_proto_depIdxs,
		MessageInfos:      file_proto_cities_proto_msgTypes,
	}.Build()
	File_proto_cities_proto = out.File
	file_proto_cities_proto_rawDesc = nil
	file_proto_cities_proto_goTypes = nil
	file_proto_cities_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CitiesServiceClient is the client API for CitiesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CitiesServiceClient interface {
	GetCity(ctx context.Context, in *GetCityInput, opts ...grpc.CallOption) (*City, error)
	GetCities(ctx context.Context, in *EmptyInput, opts ...grpc.CallOption) (*Cities, error)
	CreateCity(ctx context.Context, in *NewCityInput, opts ...grpc.CallOption) (*City, error)
}

type citiesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCitiesServiceClient(cc grpc.ClientConnInterface) CitiesServiceClient {
	return &citiesServiceClient{cc}
}

func (c *citiesServiceClient) GetCity(ctx context.Context, in *GetCityInput, opts ...grpc.CallOption) (*City, error) {
	out := new(City)
	err := c.cc.Invoke(ctx, "/CitiesService/GetCity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *citiesServiceClient) GetCities(ctx context.Context, in *EmptyInput, opts ...grpc.CallOption) (*Cities, error) {
	out := new(Cities)
	err := c.cc.Invoke(ctx, "/CitiesService/GetCities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *citiesServiceClient) CreateCity(ctx context.Context, in *NewCityInput, opts ...grpc.CallOption) (*City, error) {
	out := new(City)
	err := c.cc.Invoke(ctx, "/CitiesService/CreateCity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CitiesServiceServer is the server API for CitiesService service.
type CitiesServiceServer interface {
	GetCity(context.Context, *GetCityInput) (*City, error)
	GetCities(context.Context, *EmptyInput) (*Cities, error)
	CreateCity(context.Context, *NewCityInput) (*City, error)
}

// UnimplementedCitiesServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCitiesServiceServer struct {
}

func (*UnimplementedCitiesServiceServer) GetCity(context.Context, *GetCityInput) (*City, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCity not implemented")
}
func (*UnimplementedCitiesServiceServer) GetCities(context.Context, *EmptyInput) (*Cities, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCities not implemented")
}
func (*UnimplementedCitiesServiceServer) CreateCity(context.Context, *NewCityInput) (*City, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCity not implemented")
}

func RegisterCitiesServiceServer(s *grpc.Server, srv CitiesServiceServer) {
	s.RegisterService(&_CitiesService_serviceDesc, srv)
}

func _CitiesService_GetCity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCityInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CitiesServiceServer).GetCity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CitiesService/GetCity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CitiesServiceServer).GetCity(ctx, req.(*GetCityInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _CitiesService_GetCities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CitiesServiceServer).GetCities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CitiesService/GetCities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CitiesServiceServer).GetCities(ctx, req.(*EmptyInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _CitiesService_CreateCity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewCityInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CitiesServiceServer).CreateCity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CitiesService/CreateCity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CitiesServiceServer).CreateCity(ctx, req.(*NewCityInput))
	}
	return interceptor(ctx, in, info, handler)
}

var _CitiesService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CitiesService",
	HandlerType: (*CitiesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCity",
			Handler:    _CitiesService_GetCity_Handler,
		},
		{
			MethodName: "GetCities",
			Handler:    _CitiesService_GetCities_Handler,
		},
		{
			MethodName: "CreateCity",
			Handler:    _CitiesService_CreateCity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/cities.proto",
}