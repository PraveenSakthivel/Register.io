// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.7
// source: dbRequests.proto

//protoc --go_out=. --go-grpc_out=. dbRequests.proto

package dbRequests

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

type AddStatus int32

const (
	AddStatus_PENDING AddStatus = 0 // class add is still pending (not in db, class isn't full)
	AddStatus_ADDED   AddStatus = 1 // able to add class (in db)
	AddStatus_FAILED  AddStatus = 2 // unable to add class (not in db, class is full)
)

// Enum value maps for AddStatus.
var (
	AddStatus_name = map[int32]string{
		0: "PENDING",
		1: "ADDED",
		2: "FAILED",
	}
	AddStatus_value = map[string]int32{
		"PENDING": 0,
		"ADDED":   1,
		"FAILED":  2,
	}
)

func (x AddStatus) Enum() *AddStatus {
	p := new(AddStatus)
	*p = x
	return p
}

func (x AddStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AddStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_dbRequests_proto_enumTypes[0].Descriptor()
}

func (AddStatus) Type() protoreflect.EnumType {
	return &file_dbRequests_proto_enumTypes[0]
}

func (x AddStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AddStatus.Descriptor instead.
func (AddStatus) EnumDescriptor() ([]byte, []int) {
	return file_dbRequests_proto_rawDescGZIP(), []int{0}
}

type ClassAddStatusParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Index []string `protobuf:"bytes,2,rep,name=index,proto3" json:"index,omitempty"`
}

func (x *ClassAddStatusParams) Reset() {
	*x = ClassAddStatusParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbRequests_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClassAddStatusParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassAddStatusParams) ProtoMessage() {}

func (x *ClassAddStatusParams) ProtoReflect() protoreflect.Message {
	mi := &file_dbRequests_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClassAddStatusParams.ProtoReflect.Descriptor instead.
func (*ClassAddStatusParams) Descriptor() ([]byte, []int) {
	return file_dbRequests_proto_rawDescGZIP(), []int{0}
}

func (x *ClassAddStatusParams) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ClassAddStatusParams) GetIndex() []string {
	if x != nil {
		return x.Index
	}
	return nil
}

type AddStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Statuses map[string]AddStatus `protobuf:"bytes,1,rep,name=statuses,proto3" json:"statuses,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=dbRequests.AddStatus"`
}

func (x *AddStatusResponse) Reset() {
	*x = AddStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbRequests_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddStatusResponse) ProtoMessage() {}

func (x *AddStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dbRequests_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddStatusResponse.ProtoReflect.Descriptor instead.
func (*AddStatusResponse) Descriptor() ([]byte, []int) {
	return file_dbRequests_proto_rawDescGZIP(), []int{1}
}

func (x *AddStatusResponse) GetStatuses() map[string]AddStatus {
	if x != nil {
		return x.Statuses
	}
	return nil
}

type ReceiveClassesParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReceiveClassesParams) Reset() {
	*x = ReceiveClassesParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbRequests_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveClassesParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveClassesParams) ProtoMessage() {}

func (x *ReceiveClassesParams) ProtoReflect() protoreflect.Message {
	mi := &file_dbRequests_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveClassesParams.ProtoReflect.Descriptor instead.
func (*ReceiveClassesParams) Descriptor() ([]byte, []int) {
	return file_dbRequests_proto_rawDescGZIP(), []int{2}
}

type ClassesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Classes []*Class `protobuf:"bytes,1,rep,name=classes,proto3" json:"classes,omitempty"`
}

func (x *ClassesResponse) Reset() {
	*x = ClassesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbRequests_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClassesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassesResponse) ProtoMessage() {}

func (x *ClassesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dbRequests_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClassesResponse.ProtoReflect.Descriptor instead.
func (*ClassesResponse) Descriptor() ([]byte, []int) {
	return file_dbRequests_proto_rawDescGZIP(), []int{3}
}

func (x *ClassesResponse) GetClasses() []*Class {
	if x != nil {
		return x.Classes
	}
	return nil
}

type Class struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level      string     `protobuf:"bytes,1,opt,name=level,proto3" json:"level,omitempty"`
	School     int32      `protobuf:"varint,2,opt,name=school,proto3" json:"school,omitempty"`
	Department int32      `protobuf:"varint,3,opt,name=department,proto3" json:"department,omitempty"`
	ClassNum   int32      `protobuf:"varint,4,opt,name=classNum,proto3" json:"classNum,omitempty"`
	Name       string     `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Codes      []string   `protobuf:"bytes,6,rep,name=codes,proto3" json:"codes,omitempty"`
	Synopsis   string     `protobuf:"bytes,7,opt,name=synopsis,proto3" json:"synopsis,omitempty"`
	Books      []string   `protobuf:"bytes,8,rep,name=books,proto3" json:"books,omitempty"`
	Sections   []*Section `protobuf:"bytes,9,rep,name=sections,proto3" json:"sections,omitempty"`
}

func (x *Class) Reset() {
	*x = Class{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbRequests_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Class) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Class) ProtoMessage() {}

func (x *Class) ProtoReflect() protoreflect.Message {
	mi := &file_dbRequests_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Class.ProtoReflect.Descriptor instead.
func (*Class) Descriptor() ([]byte, []int) {
	return file_dbRequests_proto_rawDescGZIP(), []int{4}
}

func (x *Class) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *Class) GetSchool() int32 {
	if x != nil {
		return x.School
	}
	return 0
}

func (x *Class) GetDepartment() int32 {
	if x != nil {
		return x.Department
	}
	return 0
}

func (x *Class) GetClassNum() int32 {
	if x != nil {
		return x.ClassNum
	}
	return 0
}

func (x *Class) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Class) GetCodes() []string {
	if x != nil {
		return x.Codes
	}
	return nil
}

func (x *Class) GetSynopsis() string {
	if x != nil {
		return x.Synopsis
	}
	return ""
}

func (x *Class) GetBooks() []string {
	if x != nil {
		return x.Books
	}
	return nil
}

func (x *Class) GetSections() []*Section {
	if x != nil {
		return x.Sections
	}
	return nil
}

type Section struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index       string     `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Section     string     `protobuf:"bytes,2,opt,name=section,proto3" json:"section,omitempty"`
	Meetings    []*Meeting `protobuf:"bytes,3,rep,name=meetings,proto3" json:"meetings,omitempty"`
	Instructors []string   `protobuf:"bytes,4,rep,name=instructors,proto3" json:"instructors,omitempty"`
	Available   bool       `protobuf:"varint,5,opt,name=available,proto3" json:"available,omitempty"`
	Exam        string     `protobuf:"bytes,6,opt,name=exam,proto3" json:"exam,omitempty"`
}

func (x *Section) Reset() {
	*x = Section{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbRequests_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Section) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Section) ProtoMessage() {}

func (x *Section) ProtoReflect() protoreflect.Message {
	mi := &file_dbRequests_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Section.ProtoReflect.Descriptor instead.
func (*Section) Descriptor() ([]byte, []int) {
	return file_dbRequests_proto_rawDescGZIP(), []int{5}
}

func (x *Section) GetIndex() string {
	if x != nil {
		return x.Index
	}
	return ""
}

func (x *Section) GetSection() string {
	if x != nil {
		return x.Section
	}
	return ""
}

func (x *Section) GetMeetings() []*Meeting {
	if x != nil {
		return x.Meetings
	}
	return nil
}

func (x *Section) GetInstructors() []string {
	if x != nil {
		return x.Instructors
	}
	return nil
}

func (x *Section) GetAvailable() bool {
	if x != nil {
		return x.Available
	}
	return false
}

func (x *Section) GetExam() string {
	if x != nil {
		return x.Exam
	}
	return ""
}

type Meeting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MeetingTime     string `protobuf:"bytes,1,opt,name=meeting_time,json=meetingTime,proto3" json:"meeting_time,omitempty"`
	MeetingLocation string `protobuf:"bytes,2,opt,name=meeting_location,json=meetingLocation,proto3" json:"meeting_location,omitempty"`
	Campus          string `protobuf:"bytes,3,opt,name=campus,proto3" json:"campus,omitempty"`
}

func (x *Meeting) Reset() {
	*x = Meeting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbRequests_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meeting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meeting) ProtoMessage() {}

func (x *Meeting) ProtoReflect() protoreflect.Message {
	mi := &file_dbRequests_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meeting.ProtoReflect.Descriptor instead.
func (*Meeting) Descriptor() ([]byte, []int) {
	return file_dbRequests_proto_rawDescGZIP(), []int{6}
}

func (x *Meeting) GetMeetingTime() string {
	if x != nil {
		return x.MeetingTime
	}
	return ""
}

func (x *Meeting) GetMeetingLocation() string {
	if x != nil {
		return x.MeetingLocation
	}
	return ""
}

func (x *Meeting) GetCampus() string {
	if x != nil {
		return x.Campus
	}
	return ""
}

var File_dbRequests_proto protoreflect.FileDescriptor

var file_dbRequests_proto_rawDesc = []byte{
	0x0a, 0x10, 0x64, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x64, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x22, 0x42,
	0x0a, 0x14, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x41, 0x64, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x22, 0xb0, 0x01, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x64, 0x62, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65,
	0x73, 0x1a, 0x52, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x2b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x64, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73,
	0x2e, 0x41, 0x64, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x16, 0x0a, 0x14, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x3e, 0x0a,
	0x0f, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2b, 0x0a, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x64, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x52, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x22, 0xfe, 0x01,
	0x0a, 0x05, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73,
	0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x4e, 0x75,
	0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x4e, 0x75,
	0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x79, 0x6e, 0x6f, 0x70, 0x73, 0x69, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x79, 0x6e, 0x6f, 0x70, 0x73, 0x69, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73,
	0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x2f, 0x0a,
	0x08, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x64, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x53, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xbe,
	0x01, 0x0a, 0x07, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x08, 0x6d, 0x65,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64,
	0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x08, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x69,
	0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x1c, 0x0a,
	0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x65,
	0x78, 0x61, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x78, 0x61, 0x6d, 0x22,
	0x6f, 0x0a, 0x07, 0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x29, 0x0a,
	0x10, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x6d, 0x70,
	0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73,
	0x2a, 0x2f, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a,
	0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x44,
	0x44, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10,
	0x02, 0x32, 0xba, 0x01, 0x0a, 0x0f, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x57, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x12, 0x52, 0x0a, 0x0f, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76,
	0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x12, 0x20, 0x2e, 0x64, 0x62, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x65, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x1b, 0x2e, 0x64, 0x62, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0e, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x41, 0x64, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x2e, 0x64, 0x62,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x41, 0x64,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x1d, 0x2e,
	0x64, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0e,
	0x5a, 0x0c, 0x2e, 0x2f, 0x64, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dbRequests_proto_rawDescOnce sync.Once
	file_dbRequests_proto_rawDescData = file_dbRequests_proto_rawDesc
)

func file_dbRequests_proto_rawDescGZIP() []byte {
	file_dbRequests_proto_rawDescOnce.Do(func() {
		file_dbRequests_proto_rawDescData = protoimpl.X.CompressGZIP(file_dbRequests_proto_rawDescData)
	})
	return file_dbRequests_proto_rawDescData
}

var file_dbRequests_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_dbRequests_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_dbRequests_proto_goTypes = []interface{}{
	(AddStatus)(0),               // 0: dbRequests.AddStatus
	(*ClassAddStatusParams)(nil), // 1: dbRequests.ClassAddStatusParams
	(*AddStatusResponse)(nil),    // 2: dbRequests.AddStatusResponse
	(*ReceiveClassesParams)(nil), // 3: dbRequests.ReceiveClassesParams
	(*ClassesResponse)(nil),      // 4: dbRequests.ClassesResponse
	(*Class)(nil),                // 5: dbRequests.Class
	(*Section)(nil),              // 6: dbRequests.Section
	(*Meeting)(nil),              // 7: dbRequests.Meeting
	nil,                          // 8: dbRequests.AddStatusResponse.StatusesEntry
}
var file_dbRequests_proto_depIdxs = []int32{
	8, // 0: dbRequests.AddStatusResponse.statuses:type_name -> dbRequests.AddStatusResponse.StatusesEntry
	5, // 1: dbRequests.ClassesResponse.classes:type_name -> dbRequests.Class
	6, // 2: dbRequests.Class.sections:type_name -> dbRequests.Section
	7, // 3: dbRequests.Section.meetings:type_name -> dbRequests.Meeting
	0, // 4: dbRequests.AddStatusResponse.StatusesEntry.value:type_name -> dbRequests.AddStatus
	3, // 5: dbRequests.DatabaseWrapper.RetrieveClasses:input_type -> dbRequests.ReceiveClassesParams
	1, // 6: dbRequests.DatabaseWrapper.ClassAddStatus:input_type -> dbRequests.ClassAddStatusParams
	4, // 7: dbRequests.DatabaseWrapper.RetrieveClasses:output_type -> dbRequests.ClassesResponse
	2, // 8: dbRequests.DatabaseWrapper.ClassAddStatus:output_type -> dbRequests.AddStatusResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_dbRequests_proto_init() }
func file_dbRequests_proto_init() {
	if File_dbRequests_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dbRequests_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClassAddStatusParams); i {
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
		file_dbRequests_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddStatusResponse); i {
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
		file_dbRequests_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveClassesParams); i {
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
		file_dbRequests_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClassesResponse); i {
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
		file_dbRequests_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Class); i {
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
		file_dbRequests_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Section); i {
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
		file_dbRequests_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Meeting); i {
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
			RawDescriptor: file_dbRequests_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dbRequests_proto_goTypes,
		DependencyIndexes: file_dbRequests_proto_depIdxs,
		EnumInfos:         file_dbRequests_proto_enumTypes,
		MessageInfos:      file_dbRequests_proto_msgTypes,
	}.Build()
	File_dbRequests_proto = out.File
	file_dbRequests_proto_rawDesc = nil
	file_dbRequests_proto_goTypes = nil
	file_dbRequests_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DatabaseWrapperClient is the client API for DatabaseWrapper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DatabaseWrapperClient interface {
	RetrieveClasses(ctx context.Context, in *ReceiveClassesParams, opts ...grpc.CallOption) (*ClassesResponse, error)
	ClassAddStatus(ctx context.Context, in *ClassAddStatusParams, opts ...grpc.CallOption) (*AddStatusResponse, error)
}

type databaseWrapperClient struct {
	cc grpc.ClientConnInterface
}

func NewDatabaseWrapperClient(cc grpc.ClientConnInterface) DatabaseWrapperClient {
	return &databaseWrapperClient{cc}
}

func (c *databaseWrapperClient) RetrieveClasses(ctx context.Context, in *ReceiveClassesParams, opts ...grpc.CallOption) (*ClassesResponse, error) {
	out := new(ClassesResponse)
	err := c.cc.Invoke(ctx, "/dbRequests.DatabaseWrapper/RetrieveClasses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseWrapperClient) ClassAddStatus(ctx context.Context, in *ClassAddStatusParams, opts ...grpc.CallOption) (*AddStatusResponse, error) {
	out := new(AddStatusResponse)
	err := c.cc.Invoke(ctx, "/dbRequests.DatabaseWrapper/ClassAddStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatabaseWrapperServer is the server API for DatabaseWrapper service.
type DatabaseWrapperServer interface {
	RetrieveClasses(context.Context, *ReceiveClassesParams) (*ClassesResponse, error)
	ClassAddStatus(context.Context, *ClassAddStatusParams) (*AddStatusResponse, error)
}

// UnimplementedDatabaseWrapperServer can be embedded to have forward compatible implementations.
type UnimplementedDatabaseWrapperServer struct {
}

func (*UnimplementedDatabaseWrapperServer) RetrieveClasses(context.Context, *ReceiveClassesParams) (*ClassesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveClasses not implemented")
}
func (*UnimplementedDatabaseWrapperServer) ClassAddStatus(context.Context, *ClassAddStatusParams) (*AddStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClassAddStatus not implemented")
}

func RegisterDatabaseWrapperServer(s *grpc.Server, srv DatabaseWrapperServer) {
	s.RegisterService(&_DatabaseWrapper_serviceDesc, srv)
}

func _DatabaseWrapper_RetrieveClasses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReceiveClassesParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseWrapperServer).RetrieveClasses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbRequests.DatabaseWrapper/RetrieveClasses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseWrapperServer).RetrieveClasses(ctx, req.(*ReceiveClassesParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseWrapper_ClassAddStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClassAddStatusParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseWrapperServer).ClassAddStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbRequests.DatabaseWrapper/ClassAddStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseWrapperServer).ClassAddStatus(ctx, req.(*ClassAddStatusParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _DatabaseWrapper_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dbRequests.DatabaseWrapper",
	HandlerType: (*DatabaseWrapperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RetrieveClasses",
			Handler:    _DatabaseWrapper_RetrieveClasses_Handler,
		},
		{
			MethodName: "ClassAddStatus",
			Handler:    _DatabaseWrapper_ClassAddStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dbRequests.proto",
}
