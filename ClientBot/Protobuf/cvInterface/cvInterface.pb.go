// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.7
// source: cvInterface.proto

package cvInterface

import (
	context "context"
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

type ResultClass int32

const (
	ResultClass_NONE    ResultClass = 0 //Placeholder
	ResultClass_OK      ResultClass = 1 //Request processed successfuly
	ResultClass_PREREQ  ResultClass = 2 //Request failed due to prereqs
	ResultClass_TIME    ResultClass = 3 //Request failed due to timing constraints
	ResultClass_INVALID ResultClass = 5 //Request failed due to invalid index
	ResultClass_ERROR   ResultClass = 6 //Request failed due to Server side error
	ResultClass_SQS     ResultClass = 7 //Request failed due to error sending message to SQS
)

// Enum value maps for ResultClass.
var (
	ResultClass_name = map[int32]string{
		0: "NONE",
		1: "OK",
		2: "PREREQ",
		3: "TIME",
		5: "INVALID",
		6: "ERROR",
		7: "SQS",
	}
	ResultClass_value = map[string]int32{
		"NONE":    0,
		"OK":      1,
		"PREREQ":  2,
		"TIME":    3,
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

type ReqOp int32

const (
	ReqOp_NONEN ReqOp = 0 //Placeholder
	ReqOp_ADD   ReqOp = 1
	ReqOp_DROP  ReqOp = 2
	ReqOp_SPN   ReqOp = 3
)

// Enum value maps for ReqOp.
var (
	ReqOp_name = map[int32]string{
		0: "NONEN",
		1: "ADD",
		2: "DROP",
		3: "SPN",
	}
	ReqOp_value = map[string]int32{
		"NONEN": 0,
		"ADD":   1,
		"DROP":  2,
		"SPN":   3,
	}
)

func (x ReqOp) Enum() *ReqOp {
	p := new(ReqOp)
	*p = x
	return p
}

func (x ReqOp) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReqOp) Descriptor() protoreflect.EnumDescriptor {
	return file_cvInterface_proto_enumTypes[1].Descriptor()
}

func (ReqOp) Type() protoreflect.EnumType {
	return &file_cvInterface_proto_enumTypes[1]
}

func (x ReqOp) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReqOp.Descriptor instead.
func (ReqOp) EnumDescriptor() ([]byte, []int) {
	return file_cvInterface_proto_rawDescGZIP(), []int{1}
}

type ClassOperations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index string `protobuf:"bytes,1,opt,name=Index,proto3" json:"Index,omitempty"`
	Op    ReqOp  `protobuf:"varint,2,opt,name=Op,proto3,enum=cvInterface.ReqOp" json:"Op,omitempty"`
}

func (x *ClassOperations) Reset() {
	*x = ClassOperations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cvInterface_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClassOperations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassOperations) ProtoMessage() {}

func (x *ClassOperations) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ClassOperations.ProtoReflect.Descriptor instead.
func (*ClassOperations) Descriptor() ([]byte, []int) {
	return file_cvInterface_proto_rawDescGZIP(), []int{0}
}

func (x *ClassOperations) GetIndex() string {
	if x != nil {
		return x.Index
	}
	return ""
}

func (x *ClassOperations) GetOp() ReqOp {
	if x != nil {
		return x.Op
	}
	return ReqOp_NONEN
}

type RegistrationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token   string             `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	Classes []*ClassOperations `protobuf:"bytes,2,rep,name=Classes,proto3" json:"Classes,omitempty"`
}

func (x *RegistrationRequest) Reset() {
	*x = RegistrationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cvInterface_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistrationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistrationRequest) ProtoMessage() {}

func (x *RegistrationRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use RegistrationRequest.ProtoReflect.Descriptor instead.
func (*RegistrationRequest) Descriptor() ([]byte, []int) {
	return file_cvInterface_proto_rawDescGZIP(), []int{1}
}

func (x *RegistrationRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *RegistrationRequest) GetClasses() []*ClassOperations {
	if x != nil {
		return x.Classes
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
		mi := &file_cvInterface_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistrationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistrationResponse) ProtoMessage() {}

func (x *RegistrationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cvInterface_proto_msgTypes[2]
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
	return file_cvInterface_proto_rawDescGZIP(), []int{2}
}

func (x *RegistrationResponse) GetResults() map[string]ResultClass {
	if x != nil {
		return x.Results
	}
	return nil
}

type SPNRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	Spn   string `protobuf:"bytes,2,opt,name=spn,proto3" json:"spn,omitempty"`
	Index string `protobuf:"bytes,3,opt,name=index,proto3" json:"index,omitempty"`
}

func (x *SPNRequest) Reset() {
	*x = SPNRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cvInterface_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SPNRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SPNRequest) ProtoMessage() {}

func (x *SPNRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cvInterface_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SPNRequest.ProtoReflect.Descriptor instead.
func (*SPNRequest) Descriptor() ([]byte, []int) {
	return file_cvInterface_proto_rawDescGZIP(), []int{3}
}

func (x *SPNRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SPNRequest) GetSpn() string {
	if x != nil {
		return x.Spn
	}
	return ""
}

func (x *SPNRequest) GetIndex() string {
	if x != nil {
		return x.Index
	}
	return ""
}

type SPNResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valid  bool        `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
	Result ResultClass `protobuf:"varint,2,opt,name=result,proto3,enum=cvInterface.ResultClass" json:"result,omitempty"`
}

func (x *SPNResponse) Reset() {
	*x = SPNResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cvInterface_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SPNResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SPNResponse) ProtoMessage() {}

func (x *SPNResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cvInterface_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SPNResponse.ProtoReflect.Descriptor instead.
func (*SPNResponse) Descriptor() ([]byte, []int) {
	return file_cvInterface_proto_rawDescGZIP(), []int{4}
}

func (x *SPNResponse) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

func (x *SPNResponse) GetResult() ResultClass {
	if x != nil {
		return x.Result
	}
	return ResultClass_NONE
}

var File_cvInterface_proto protoreflect.FileDescriptor

var file_cvInterface_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x76, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x76, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x22, 0x4b, 0x0a, 0x0f, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x22, 0x0a, 0x02, 0x4f, 0x70, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x63, 0x76, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x4f, 0x70, 0x52, 0x02, 0x4f, 0x70, 0x22, 0x63, 0x0a,
	0x13, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x36, 0x0a, 0x07, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x76,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x65, 0x73, 0x22, 0xb6, 0x01, 0x0a, 0x14, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x07, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x63,
	0x76, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x1a, 0x54, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x63, 0x76, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x4a, 0x0a, 0x0a, 0x53,
	0x50, 0x4e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x10, 0x0a, 0x03, 0x73, 0x70, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x70,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x55, 0x0a, 0x0b, 0x53, 0x50, 0x4e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x63,
	0x76, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2a, 0x56,
	0x0a, 0x0b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x08, 0x0a,
	0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x01, 0x12,
	0x0a, 0x0a, 0x06, 0x50, 0x52, 0x45, 0x52, 0x45, 0x51, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x54,
	0x49, 0x4d, 0x45, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44,
	0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x06, 0x12, 0x07, 0x0a,
	0x03, 0x53, 0x51, 0x53, 0x10, 0x07, 0x2a, 0x2e, 0x0a, 0x05, 0x52, 0x65, 0x71, 0x4f, 0x70, 0x12,
	0x09, 0x0a, 0x05, 0x4e, 0x4f, 0x4e, 0x45, 0x4e, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x44,
	0x44, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x52, 0x4f, 0x50, 0x10, 0x02, 0x12, 0x07, 0x0a,
	0x03, 0x53, 0x50, 0x4e, 0x10, 0x03, 0x32, 0xae, 0x01, 0x0a, 0x10, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5b, 0x0a, 0x12, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x20, 0x2e, 0x63, 0x76, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x63, 0x76, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x53,
	0x50, 0x4e, 0x12, 0x17, 0x2e, 0x63, 0x76, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x2e, 0x53, 0x50, 0x4e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x76,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x53, 0x50, 0x4e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x63, 0x76, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_cvInterface_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_cvInterface_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_cvInterface_proto_goTypes = []interface{}{
	(ResultClass)(0),             // 0: cvInterface.ResultClass
	(ReqOp)(0),                   // 1: cvInterface.ReqOp
	(*ClassOperations)(nil),      // 2: cvInterface.ClassOperations
	(*RegistrationRequest)(nil),  // 3: cvInterface.RegistrationRequest
	(*RegistrationResponse)(nil), // 4: cvInterface.RegistrationResponse
	(*SPNRequest)(nil),           // 5: cvInterface.SPNRequest
	(*SPNResponse)(nil),          // 6: cvInterface.SPNResponse
	nil,                          // 7: cvInterface.RegistrationResponse.ResultsEntry
}
var file_cvInterface_proto_depIdxs = []int32{
	1, // 0: cvInterface.ClassOperations.Op:type_name -> cvInterface.ReqOp
	2, // 1: cvInterface.RegistrationRequest.Classes:type_name -> cvInterface.ClassOperations
	7, // 2: cvInterface.RegistrationResponse.Results:type_name -> cvInterface.RegistrationResponse.ResultsEntry
	0, // 3: cvInterface.SPNResponse.result:type_name -> cvInterface.ResultClass
	0, // 4: cvInterface.RegistrationResponse.ResultsEntry.value:type_name -> cvInterface.ResultClass
	3, // 5: cvInterface.CourseValidation.ChangeRegistration:input_type -> cvInterface.RegistrationRequest
	5, // 6: cvInterface.CourseValidation.AddSPN:input_type -> cvInterface.SPNRequest
	4, // 7: cvInterface.CourseValidation.ChangeRegistration:output_type -> cvInterface.RegistrationResponse
	6, // 8: cvInterface.CourseValidation.AddSPN:output_type -> cvInterface.SPNResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_cvInterface_proto_init() }
func file_cvInterface_proto_init() {
	if File_cvInterface_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cvInterface_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClassOperations); i {
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
		file_cvInterface_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_cvInterface_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SPNRequest); i {
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
		file_cvInterface_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SPNResponse); i {
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
			NumEnums:      2,
			NumMessages:   6,
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
	AddSPN(ctx context.Context, in *SPNRequest, opts ...grpc.CallOption) (*SPNResponse, error)
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

func (c *courseValidationClient) AddSPN(ctx context.Context, in *SPNRequest, opts ...grpc.CallOption) (*SPNResponse, error) {
	out := new(SPNResponse)
	err := c.cc.Invoke(ctx, "/cvInterface.CourseValidation/AddSPN", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CourseValidationServer is the server API for CourseValidation service.
type CourseValidationServer interface {
	ChangeRegistration(context.Context, *RegistrationRequest) (*RegistrationResponse, error)
	AddSPN(context.Context, *SPNRequest) (*SPNResponse, error)
}

// UnimplementedCourseValidationServer can be embedded to have forward compatible implementations.
type UnimplementedCourseValidationServer struct {
}

func (*UnimplementedCourseValidationServer) ChangeRegistration(context.Context, *RegistrationRequest) (*RegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeRegistration not implemented")
}
func (*UnimplementedCourseValidationServer) AddSPN(context.Context, *SPNRequest) (*SPNResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSPN not implemented")
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

func _CourseValidation_AddSPN_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SPNRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseValidationServer).AddSPN(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cvInterface.CourseValidation/AddSPN",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseValidationServer).AddSPN(ctx, req.(*SPNRequest))
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
		{
			MethodName: "AddSPN",
			Handler:    _CourseValidation_AddSPN_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cvInterface.proto",
}
