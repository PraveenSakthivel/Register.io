// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: cvInterface.proto

package cvInterface

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

type ResultClass int32

const (
	ResultClass_OK      ResultClass = 0 //Request processed successfuly
	ResultClass_PREREQ  ResultClass = 1 //Request failed due to prereqs
	ResultClass_TIME    ResultClass = 3 //Request failed due to timing constraints
	ResultClass_SPN     ResultClass = 4 //Request failed due to invalid SPN
	ResultClass_INVALID ResultClass = 5 //Request failed due to invalid index
	ResultClass_ERROR   ResultClass = 6 //Request failed due to Server side error
	ResultClass_SQS     ResultClass = 7 //Request failed due to error sending message to SQS
)

// Enum value maps for ResultClass.
var (
	ResultClass_name = map[int32]string{
		0: "OK",
		1: "PREREQ",
		3: "TIME",
		4: "SPN",
		5: "INVALID",
		6: "ERROR",
		7: "SQS",
	}
	ResultClass_value = map[string]int32{
		"OK":      0,
		"PREREQ":  1,
		"TIME":    3,
		"SPN":     4,
		"INVALID": 5,
		"ERROR":   6,
		"SQS":     7,
	}
)

func (x ResultClass) Enum() *ResultClass {
	p := new(ResultClass)
	*p = x
	return p
}

func (x ResultClass) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResultClass) Descriptor() protoreflect.EnumDescriptor {
	return file_cvInterface_proto_enumTypes[0].Descriptor()
}

func (ResultClass) Type() protoreflect.EnumType {
	return &file_cvInterface_proto_enumTypes[0]
}

func (x ResultClass) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResultClass.Descriptor instead.
func (ResultClass) EnumDescriptor() ([]byte, []int) {
	return file_cvInterface_proto_rawDescGZIP(), []int{0}
}

type RegistrationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token   string   `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	Indices []string `protobuf:"bytes,2,rep,name=Indices,proto3" json:"Indices,omitempty"`
}

func (x *RegistrationRequest) Reset() {
	*x = RegistrationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cvInterface_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistrationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistrationRequest) ProtoMessage() {}

func (x *RegistrationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cvInterface_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistrationRequest.ProtoReflect.Descriptor instead.
func (*RegistrationRequest) Descriptor() ([]byte, []int) {
	return file_cvInterface_proto_rawDescGZIP(), []int{0}
}

func (x *RegistrationRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *RegistrationRequest) GetIndices() []string {
	if x != nil {
		return x.Indices
	}
	return nil
}

type RegistrationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results map[string]ResultClass `protobuf:"bytes,1,rep,name=Results,proto3" json:"Results,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=cvInterface.ResultClass"`
}

func (x *RegistrationResponse) Reset() {
	*x = RegistrationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cvInterface_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistrationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistrationResponse) ProtoMessage() {}

func (x *RegistrationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cvInterface_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistrationResponse.ProtoReflect.Descriptor instead.
func (*RegistrationResponse) Descriptor() ([]byte, []int) {
	return file_cvInterface_proto_rawDescGZIP(), []int{1}
}

func (x *RegistrationResponse) GetResults() map[string]ResultClass {
	if x != nil {
		return x.Results
	}
	return nil
}

var File_cvInterface_proto protoreflect.FileDescriptor

var file_cvInterface_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x76, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x76, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x22, 0x45, 0x0a, 0x13, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07,
	0x49, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x22, 0xb6, 0x01, 0x0a, 0x14, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x48, 0x0a, 0x07, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x76, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x07, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x1a, 0x54, 0x0a, 0x0c, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2e, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x63, 0x76,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x2a, 0x55, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12,
	0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x52, 0x45, 0x52, 0x45,
	0x51, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x03, 0x12, 0x07, 0x0a,
	0x03, 0x53, 0x50, 0x4e, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x06, 0x12, 0x07,
	0x0a, 0x03, 0x53, 0x51, 0x53, 0x10, 0x07, 0x32, 0x6f, 0x0a, 0x10, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5b, 0x0a, 0x12, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x20, 0x2e, 0x63, 0x76, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x63, 0x76, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cvInterface_proto_rawDescOnce sync.Once
	file_cvInterface_proto_rawDescData = file_cvInterface_proto_rawDesc
)

func file_cvInterface_proto_rawDescGZIP() []byte {
	file_cvInterface_proto_rawDescOnce.Do(func() {
		file_cvInterface_proto_rawDescData = protoimpl.X.CompressGZIP(file_cvInterface_proto_rawDescData)
	})
	return file_cvInterface_proto_rawDescData
}

var file_cvInterface_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cvInterface_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cvInterface_proto_goTypes = []interface{}{
	(ResultClass)(0),             // 0: cvInterface.ResultClass
	(*RegistrationRequest)(nil),  // 1: cvInterface.RegistrationRequest
	(*RegistrationResponse)(nil), // 2: cvInterface.RegistrationResponse
	nil,                          // 3: cvInterface.RegistrationResponse.ResultsEntry
}
var file_cvInterface_proto_depIdxs = []int32{
	3, // 0: cvInterface.RegistrationResponse.Results:type_name -> cvInterface.RegistrationResponse.ResultsEntry
	0, // 1: cvInterface.RegistrationResponse.ResultsEntry.value:type_name -> cvInterface.ResultClass
	1, // 2: cvInterface.CourseValidation.ChangeRegistration:input_type -> cvInterface.RegistrationRequest
	2, // 3: cvInterface.CourseValidation.ChangeRegistration:output_type -> cvInterface.RegistrationResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cvInterface_proto_init() }
func file_cvInterface_proto_init() {
	if File_cvInterface_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cvInterface_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistrationRequest); i {
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
		file_cvInterface_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistrationResponse); i {
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
			RawDescriptor: file_cvInterface_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cvInterface_proto_goTypes,
		DependencyIndexes: file_cvInterface_proto_depIdxs,
		EnumInfos:         file_cvInterface_proto_enumTypes,
		MessageInfos:      file_cvInterface_proto_msgTypes,
	}.Build()
	File_cvInterface_proto = out.File
	file_cvInterface_proto_rawDesc = nil
	file_cvInterface_proto_goTypes = nil
	file_cvInterface_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CourseValidationClient is the client API for CourseValidation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CourseValidationClient interface {
	ChangeRegistration(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*RegistrationResponse, error)
}

type courseValidationClient struct {
	cc grpc.ClientConnInterface
}

func NewCourseValidationClient(cc grpc.ClientConnInterface) CourseValidationClient {
	return &courseValidationClient{cc}
}

func (c *courseValidationClient) ChangeRegistration(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*RegistrationResponse, error) {
	out := new(RegistrationResponse)
	err := c.cc.Invoke(ctx, "/cvInterface.CourseValidation/ChangeRegistration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CourseValidationServer is the server API for CourseValidation service.
type CourseValidationServer interface {
	ChangeRegistration(context.Context, *RegistrationRequest) (*RegistrationResponse, error)
}

// UnimplementedCourseValidationServer can be embedded to have forward compatible implementations.
type UnimplementedCourseValidationServer struct {
}

func (*UnimplementedCourseValidationServer) ChangeRegistration(context.Context, *RegistrationRequest) (*RegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeRegistration not implemented")
}

func RegisterCourseValidationServer(s *grpc.Server, srv CourseValidationServer) {
	s.RegisterService(&_CourseValidation_serviceDesc, srv)
}

func _CourseValidation_ChangeRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseValidationServer).ChangeRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cvInterface.CourseValidation/ChangeRegistration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseValidationServer).ChangeRegistration(ctx, req.(*RegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CourseValidation_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cvInterface.CourseValidation",
	HandlerType: (*CourseValidationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ChangeRegistration",
			Handler:    _CourseValidation_ChangeRegistration_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cvInterface.proto",
}
