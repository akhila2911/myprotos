// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: myservice.proto

package myservice

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StatusCode int32

const (
	StatusCode_UNSPECIFIED StatusCode = 0
	StatusCode_SUCCESS     StatusCode = 1
	StatusCode_FAILURE     StatusCode = 2
)

// Enum value maps for StatusCode.
var (
	StatusCode_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "SUCCESS",
		2: "FAILURE",
	}
	StatusCode_value = map[string]int32{
		"UNSPECIFIED": 0,
		"SUCCESS":     1,
		"FAILURE":     2,
	}
)

func (x StatusCode) Enum() *StatusCode {
	p := new(StatusCode)
	*p = x
	return p
}

func (x StatusCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusCode) Descriptor() protoreflect.EnumDescriptor {
	return file_myservice_proto_enumTypes[0].Descriptor()
}

func (StatusCode) Type() protoreflect.EnumType {
	return &file_myservice_proto_enumTypes[0]
}

func (x StatusCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusCode.Descriptor instead.
func (StatusCode) EnumDescriptor() ([]byte, []int) {
	return file_myservice_proto_rawDescGZIP(), []int{0}
}

type DescriptionCode int32

const (
	DescriptionCode_NOT_SPECIFIED         DescriptionCode = 0
	DescriptionCode_INVALID_ARGUMENT      DescriptionCode = 1
	DescriptionCode_OK                    DescriptionCode = 200
	DescriptionCode_CREATED               DescriptionCode = 201
	DescriptionCode_ACCEPTED              DescriptionCode = 202
	DescriptionCode_NO_CONTENT            DescriptionCode = 204
	DescriptionCode_BAD_REQUEST           DescriptionCode = 400
	DescriptionCode_UNAUTHORIZED          DescriptionCode = 401
	DescriptionCode_FORBIDDEN             DescriptionCode = 403
	DescriptionCode_NOT_FOUND             DescriptionCode = 404
	DescriptionCode_INTERNAL_SERVER_ERROR DescriptionCode = 500
	DescriptionCode_NOT_IMPLEMENTED       DescriptionCode = 501
	DescriptionCode_BAD_GATEWAY           DescriptionCode = 502
	DescriptionCode_SERVICE_UNAVAILABLE   DescriptionCode = 503
	DescriptionCode_GATEWAY_TIMEOUT       DescriptionCode = 504
)

// Enum value maps for DescriptionCode.
var (
	DescriptionCode_name = map[int32]string{
		0:   "NOT_SPECIFIED",
		1:   "INVALID_ARGUMENT",
		200: "OK",
		201: "CREATED",
		202: "ACCEPTED",
		204: "NO_CONTENT",
		400: "BAD_REQUEST",
		401: "UNAUTHORIZED",
		403: "FORBIDDEN",
		404: "NOT_FOUND",
		500: "INTERNAL_SERVER_ERROR",
		501: "NOT_IMPLEMENTED",
		502: "BAD_GATEWAY",
		503: "SERVICE_UNAVAILABLE",
		504: "GATEWAY_TIMEOUT",
	}
	DescriptionCode_value = map[string]int32{
		"NOT_SPECIFIED":         0,
		"INVALID_ARGUMENT":      1,
		"OK":                    200,
		"CREATED":               201,
		"ACCEPTED":              202,
		"NO_CONTENT":            204,
		"BAD_REQUEST":           400,
		"UNAUTHORIZED":          401,
		"FORBIDDEN":             403,
		"NOT_FOUND":             404,
		"INTERNAL_SERVER_ERROR": 500,
		"NOT_IMPLEMENTED":       501,
		"BAD_GATEWAY":           502,
		"SERVICE_UNAVAILABLE":   503,
		"GATEWAY_TIMEOUT":       504,
	}
)

func (x DescriptionCode) Enum() *DescriptionCode {
	p := new(DescriptionCode)
	*p = x
	return p
}

func (x DescriptionCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DescriptionCode) Descriptor() protoreflect.EnumDescriptor {
	return file_myservice_proto_enumTypes[1].Descriptor()
}

func (DescriptionCode) Type() protoreflect.EnumType {
	return &file_myservice_proto_enumTypes[1]
}

func (x DescriptionCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DescriptionCode.Descriptor instead.
func (DescriptionCode) EnumDescriptor() ([]byte, []int) {
	return file_myservice_proto_rawDescGZIP(), []int{1}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_myservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_myservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_myservice_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []int64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *GetUsersRequest) Reset() {
	*x = GetUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_myservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersRequest) ProtoMessage() {}

func (x *GetUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_myservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersRequest.ProtoReflect.Descriptor instead.
func (*GetUsersRequest) Descriptor() ([]byte, []int) {
	return file_myservice_proto_rawDescGZIP(), []int{1}
}

func (x *GetUsersRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type GetUsersResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users  []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	Status *Status `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GetUsersResp) Reset() {
	*x = GetUsersResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_myservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersResp) ProtoMessage() {}

func (x *GetUsersResp) ProtoReflect() protoreflect.Message {
	mi := &file_myservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersResp.ProtoReflect.Descriptor instead.
func (*GetUsersResp) Descriptor() ([]byte, []int) {
	return file_myservice_proto_rawDescGZIP(), []int{2}
}

func (x *GetUsersResp) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *GetUsersResp) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

type HealthReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestInfo string `protobuf:"bytes,1,opt,name=request_info,json=requestInfo,proto3" json:"request_info,omitempty"`
}

func (x *HealthReq) Reset() {
	*x = HealthReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_myservice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthReq) ProtoMessage() {}

func (x *HealthReq) ProtoReflect() protoreflect.Message {
	mi := &file_myservice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthReq.ProtoReflect.Descriptor instead.
func (*HealthReq) Descriptor() ([]byte, []int) {
	return file_myservice_proto_rawDescGZIP(), []int{3}
}

func (x *HealthReq) GetRequestInfo() string {
	if x != nil {
		return x.RequestInfo
	}
	return ""
}

type HealthResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string  `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Status  *Status `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *HealthResp) Reset() {
	*x = HealthResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_myservice_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthResp) ProtoMessage() {}

func (x *HealthResp) ProtoReflect() protoreflect.Message {
	mi := &file_myservice_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthResp.ProtoReflect.Descriptor instead.
func (*HealthResp) Descriptor() ([]byte, []int) {
	return file_myservice_proto_rawDescGZIP(), []int{4}
}

func (x *HealthResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *HealthResp) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode        StatusCode         `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3,enum=myservice.StatusCode" json:"status_code,omitempty"`
	StatusDescription *StatusDescription `protobuf:"bytes,2,opt,name=status_description,json=statusDescription,proto3" json:"status_description,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_myservice_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_myservice_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_myservice_proto_rawDescGZIP(), []int{5}
}

func (x *Status) GetStatusCode() StatusCode {
	if x != nil {
		return x.StatusCode
	}
	return StatusCode_UNSPECIFIED
}

func (x *Status) GetStatusDescription() *StatusDescription {
	if x != nil {
		return x.StatusDescription
	}
	return nil
}

type StatusDescription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DescriptionCode DescriptionCode `protobuf:"varint,1,opt,name=description_code,json=descriptionCode,proto3,enum=myservice.DescriptionCode" json:"description_code,omitempty"`
	Description     string          `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *StatusDescription) Reset() {
	*x = StatusDescription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_myservice_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusDescription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusDescription) ProtoMessage() {}

func (x *StatusDescription) ProtoReflect() protoreflect.Message {
	mi := &file_myservice_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusDescription.ProtoReflect.Descriptor instead.
func (*StatusDescription) Descriptor() ([]byte, []int) {
	return file_myservice_proto_rawDescGZIP(), []int{6}
}

func (x *StatusDescription) GetDescriptionCode() DescriptionCode {
	if x != nil {
		return x.DescriptionCode
	}
	return DescriptionCode_NOT_SPECIFIED
}

func (x *StatusDescription) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_myservice_proto protoreflect.FileDescriptor

var file_myservice_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6d, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x6d, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x04, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x23, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x60, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x25, 0x0a, 0x05, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x79, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x12, 0x29, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2e, 0x0a, 0x09,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x71, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x51, 0x0a, 0x0a,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x8d, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x36, 0x0a, 0x0b, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x15, 0x2e, 0x6d, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x4b, 0x0a, 0x12, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x6d, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x7c, 0x0a, 0x11, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x45, 0x0a, 0x10, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a,
	0x2e, 0x6d, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x0f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x37, 0x0a,
	0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x41, 0x49,
	0x4c, 0x55, 0x52, 0x45, 0x10, 0x02, 0x2a, 0xaa, 0x02, 0x0a, 0x0f, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x4e, 0x4f,
	0x54, 0x5f, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a,
	0x10, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e,
	0x54, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0xc8, 0x01, 0x12, 0x0c, 0x0a, 0x07,
	0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0xc9, 0x01, 0x12, 0x0d, 0x0a, 0x08, 0x41, 0x43,
	0x43, 0x45, 0x50, 0x54, 0x45, 0x44, 0x10, 0xca, 0x01, 0x12, 0x0f, 0x0a, 0x0a, 0x4e, 0x4f, 0x5f,
	0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x10, 0xcc, 0x01, 0x12, 0x10, 0x0a, 0x0b, 0x42, 0x41,
	0x44, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x90, 0x03, 0x12, 0x11, 0x0a, 0x0c,
	0x55, 0x4e, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x45, 0x44, 0x10, 0x91, 0x03, 0x12,
	0x0e, 0x0a, 0x09, 0x46, 0x4f, 0x52, 0x42, 0x49, 0x44, 0x44, 0x45, 0x4e, 0x10, 0x93, 0x03, 0x12,
	0x0e, 0x0a, 0x09, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x94, 0x03, 0x12,
	0x1a, 0x0a, 0x15, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x53, 0x45, 0x52, 0x56,
	0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xf4, 0x03, 0x12, 0x14, 0x0a, 0x0f, 0x4e,
	0x4f, 0x54, 0x5f, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x45, 0x44, 0x10, 0xf5,
	0x03, 0x12, 0x10, 0x0a, 0x0b, 0x42, 0x41, 0x44, 0x5f, 0x47, 0x41, 0x54, 0x45, 0x57, 0x41, 0x59,
	0x10, 0xf6, 0x03, 0x12, 0x18, 0x0a, 0x13, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x55,
	0x4e, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0xf7, 0x03, 0x12, 0x14, 0x0a,
	0x0f, 0x47, 0x41, 0x54, 0x45, 0x57, 0x41, 0x59, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54,
	0x10, 0xf8, 0x03, 0x32, 0x50, 0x0a, 0x12, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x0b, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x14, 0x2e, 0x6d, 0x79, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x15,
	0x2e, 0x6d, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x32, 0x4c, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17, 0x2e, 0x6d, 0x79, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x00, 0x42, 0x25, 0x5a, 0x23, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x2f, 0x6d, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_myservice_proto_rawDescOnce sync.Once
	file_myservice_proto_rawDescData = file_myservice_proto_rawDesc
)

func file_myservice_proto_rawDescGZIP() []byte {
	file_myservice_proto_rawDescOnce.Do(func() {
		file_myservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_myservice_proto_rawDescData)
	})
	return file_myservice_proto_rawDescData
}

var file_myservice_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_myservice_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_myservice_proto_goTypes = []interface{}{
	(StatusCode)(0),           // 0: myservice.StatusCode
	(DescriptionCode)(0),      // 1: myservice.DescriptionCode
	(*User)(nil),              // 2: myservice.User
	(*GetUsersRequest)(nil),   // 3: myservice.GetUsersRequest
	(*GetUsersResp)(nil),      // 4: myservice.GetUsersResp
	(*HealthReq)(nil),         // 5: myservice.HealthReq
	(*HealthResp)(nil),        // 6: myservice.HealthResp
	(*Status)(nil),            // 7: myservice.Status
	(*StatusDescription)(nil), // 8: myservice.StatusDescription
	(*emptypb.Empty)(nil),     // 9: google.protobuf.Empty
}
var file_myservice_proto_depIdxs = []int32{
	2, // 0: myservice.GetUsersResp.users:type_name -> myservice.User
	7, // 1: myservice.GetUsersResp.status:type_name -> myservice.Status
	7, // 2: myservice.HealthResp.status:type_name -> myservice.Status
	0, // 3: myservice.Status.status_code:type_name -> myservice.StatusCode
	8, // 4: myservice.Status.status_description:type_name -> myservice.StatusDescription
	1, // 5: myservice.StatusDescription.description_code:type_name -> myservice.DescriptionCode
	5, // 6: myservice.HealthCheckService.CheckHealth:input_type -> myservice.HealthReq
	9, // 7: myservice.UserService.GetUsers:input_type -> google.protobuf.Empty
	6, // 8: myservice.HealthCheckService.CheckHealth:output_type -> myservice.HealthResp
	4, // 9: myservice.UserService.GetUsers:output_type -> myservice.GetUsersResp
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_myservice_proto_init() }
func file_myservice_proto_init() {
	if File_myservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_myservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_myservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersRequest); i {
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
		file_myservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersResp); i {
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
		file_myservice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthReq); i {
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
		file_myservice_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthResp); i {
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
		file_myservice_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_myservice_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusDescription); i {
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
			RawDescriptor: file_myservice_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_myservice_proto_goTypes,
		DependencyIndexes: file_myservice_proto_depIdxs,
		EnumInfos:         file_myservice_proto_enumTypes,
		MessageInfos:      file_myservice_proto_msgTypes,
	}.Build()
	File_myservice_proto = out.File
	file_myservice_proto_rawDesc = nil
	file_myservice_proto_goTypes = nil
	file_myservice_proto_depIdxs = nil
}
