// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: proto/sample-server.proto

package sample

import (
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sample_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sample_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_sample_server_proto_rawDescGZIP(), []int{0}
}

type Array struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Array) Reset() {
	*x = Array{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sample_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Array) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Array) ProtoMessage() {}

func (x *Array) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sample_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Array.ProtoReflect.Descriptor instead.
func (*Array) Descriptor() ([]byte, []int) {
	return file_proto_sample_server_proto_rawDescGZIP(), []int{1}
}

func (x *Array) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Array) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateVMRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateVMRequest) Reset() {
	*x = CreateVMRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sample_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateVMRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVMRequest) ProtoMessage() {}

func (x *CreateVMRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sample_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVMRequest.ProtoReflect.Descriptor instead.
func (*CreateVMRequest) Descriptor() ([]byte, []int) {
	return file_proto_sample_server_proto_rawDescGZIP(), []int{2}
}

func (x *CreateVMRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateVMResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateVMResponse) Reset() {
	*x = CreateVMResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sample_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateVMResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVMResponse) ProtoMessage() {}

func (x *CreateVMResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sample_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVMResponse.ProtoReflect.Descriptor instead.
func (*CreateVMResponse) Descriptor() ([]byte, []int) {
	return file_proto_sample_server_proto_rawDescGZIP(), []int{3}
}

func (x *CreateVMResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetVMListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Arrays []*Array `protobuf:"bytes,1,rep,name=arrays,proto3" json:"arrays,omitempty"`
}

func (x *GetVMListResponse) Reset() {
	*x = GetVMListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sample_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVMListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVMListResponse) ProtoMessage() {}

func (x *GetVMListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sample_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVMListResponse.ProtoReflect.Descriptor instead.
func (*GetVMListResponse) Descriptor() ([]byte, []int) {
	return file_proto_sample_server_proto_rawDescGZIP(), []int{4}
}

func (x *GetVMListResponse) GetArrays() []*Array {
	if x != nil {
		return x.Arrays
	}
	return nil
}

type GetVMRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *GetVMRequest) Reset() {
	*x = GetVMRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sample_server_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVMRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVMRequest) ProtoMessage() {}

func (x *GetVMRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sample_server_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVMRequest.ProtoReflect.Descriptor instead.
func (*GetVMRequest) Descriptor() ([]byte, []int) {
	return file_proto_sample_server_proto_rawDescGZIP(), []int{5}
}

func (x *GetVMRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type GetVMResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Array *Array `protobuf:"bytes,1,opt,name=array,proto3" json:"array,omitempty"`
}

func (x *GetVMResponse) Reset() {
	*x = GetVMResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sample_server_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVMResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVMResponse) ProtoMessage() {}

func (x *GetVMResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sample_server_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVMResponse.ProtoReflect.Descriptor instead.
func (*GetVMResponse) Descriptor() ([]byte, []int) {
	return file_proto_sample_server_proto_rawDescGZIP(), []int{6}
}

func (x *GetVMResponse) GetArray() *Array {
	if x != nil {
		return x.Array
	}
	return nil
}

type GetVMStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *GetVMStatusRequest) Reset() {
	*x = GetVMStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sample_server_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVMStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVMStatusRequest) ProtoMessage() {}

func (x *GetVMStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sample_server_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVMStatusRequest.ProtoReflect.Descriptor instead.
func (*GetVMStatusRequest) Descriptor() ([]byte, []int) {
	return file_proto_sample_server_proto_rawDescGZIP(), []int{7}
}

func (x *GetVMStatusRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type GetVMStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CpuUtilization int32 `protobuf:"varint,1,opt,name=cpu_utilization,json=cpuUtilization,proto3" json:"cpu_utilization,omitempty"`
}

func (x *GetVMStatusResponse) Reset() {
	*x = GetVMStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sample_server_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVMStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVMStatusResponse) ProtoMessage() {}

func (x *GetVMStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sample_server_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVMStatusResponse.ProtoReflect.Descriptor instead.
func (*GetVMStatusResponse) Descriptor() ([]byte, []int) {
	return file_proto_sample_server_proto_rawDescGZIP(), []int{8}
}

func (x *GetVMStatusResponse) GetCpuUtilization() int32 {
	if x != nil {
		return x.CpuUtilization
	}
	return 0
}

type DeleteVMRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *DeleteVMRequest) Reset() {
	*x = DeleteVMRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sample_server_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteVMRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteVMRequest) ProtoMessage() {}

func (x *DeleteVMRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sample_server_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteVMRequest.ProtoReflect.Descriptor instead.
func (*DeleteVMRequest) Descriptor() ([]byte, []int) {
	return file_proto_sample_server_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteVMRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type CheckNameAvailabilityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CheckNameAvailabilityRequest) Reset() {
	*x = CheckNameAvailabilityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sample_server_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckNameAvailabilityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckNameAvailabilityRequest) ProtoMessage() {}

func (x *CheckNameAvailabilityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sample_server_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckNameAvailabilityRequest.ProtoReflect.Descriptor instead.
func (*CheckNameAvailabilityRequest) Descriptor() ([]byte, []int) {
	return file_proto_sample_server_proto_rawDescGZIP(), []int{10}
}

func (x *CheckNameAvailabilityRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_proto_sample_server_proto protoreflect.FileDescriptor

var file_proto_sample_server_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2d, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2b, 0x0a, 0x05, 0x61, 0x72,
	0x72, 0x61, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x25, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x56, 0x4d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x22,
	0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x4d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x3a, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x56, 0x4d, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x61, 0x72, 0x72, 0x61, 0x79,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x61, 0x72, 0x72, 0x61, 0x79, 0x52, 0x06, 0x61, 0x72, 0x72, 0x61, 0x79, 0x73, 0x22, 0x22,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x56, 0x4d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x22, 0x34, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x56, 0x4d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x61, 0x72, 0x72, 0x61, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x61, 0x72, 0x72, 0x61,
	0x79, 0x52, 0x05, 0x61, 0x72, 0x72, 0x61, 0x79, 0x22, 0x28, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x56,
	0x4d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x22, 0x3e, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x56, 0x4d, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x70, 0x75,
	0x5f, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0e, 0x63, 0x70, 0x75, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x25, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x4d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x32, 0x0a, 0x1c, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x94, 0x04,
	0x0a, 0x0c, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x4f,
	0x0a, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x4d, 0x12, 0x17, 0x2e, 0x73, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x4d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x56, 0x4d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x22, 0x08, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x12,
	0x47, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x56, 0x4d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0d, 0x2e, 0x73,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x19, 0x2e, 0x73, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x4d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x12, 0x4e, 0x0a, 0x05, 0x47, 0x65, 0x74, 0x56,
	0x4d, 0x12, 0x14, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x4d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x47, 0x65, 0x74, 0x56, 0x4d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73,
	0x2f, 0x7b, 0x76, 0x6d, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x67, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x56,
	0x4d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x47, 0x65, 0x74, 0x56, 0x4d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x56, 0x4d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x73, 0x2f, 0x7b, 0x76, 0x6d, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x4c, 0x0a, 0x08, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x4d, 0x12, 0x17, 0x2e,
	0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x4d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x2a, 0x10, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x76, 0x6d, 0x5f, 0x69, 0x64, 0x7d, 0x12,
	0x63, 0x0a, 0x15, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x41, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x24, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x41, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d,
	0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x15, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2f, 0x7b, 0x6e,
	0x61, 0x6d, 0x65, 0x7d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_sample_server_proto_rawDescOnce sync.Once
	file_proto_sample_server_proto_rawDescData = file_proto_sample_server_proto_rawDesc
)

func file_proto_sample_server_proto_rawDescGZIP() []byte {
	file_proto_sample_server_proto_rawDescOnce.Do(func() {
		file_proto_sample_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_sample_server_proto_rawDescData)
	})
	return file_proto_sample_server_proto_rawDescData
}

var file_proto_sample_server_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_sample_server_proto_goTypes = []interface{}{
	(*Empty)(nil),                        // 0: sample.Empty
	(*Array)(nil),                        // 1: sample.array
	(*CreateVMRequest)(nil),              // 2: sample.CreateVMRequest
	(*CreateVMResponse)(nil),             // 3: sample.CreateVMResponse
	(*GetVMListResponse)(nil),            // 4: sample.GetVMListResponse
	(*GetVMRequest)(nil),                 // 5: sample.GetVMRequest
	(*GetVMResponse)(nil),                // 6: sample.GetVMResponse
	(*GetVMStatusRequest)(nil),           // 7: sample.GetVMStatusRequest
	(*GetVMStatusResponse)(nil),          // 8: sample.GetVMStatusResponse
	(*DeleteVMRequest)(nil),              // 9: sample.DeleteVMRequest
	(*CheckNameAvailabilityRequest)(nil), // 10: sample.CheckNameAvailabilityRequest
}
var file_proto_sample_server_proto_depIdxs = []int32{
	1,  // 0: sample.GetVMListResponse.arrays:type_name -> sample.array
	1,  // 1: sample.GetVMResponse.array:type_name -> sample.array
	2,  // 2: sample.SampleServer.CreateVM:input_type -> sample.CreateVMRequest
	0,  // 3: sample.SampleServer.GetVMList:input_type -> sample.Empty
	5,  // 4: sample.SampleServer.GetVM:input_type -> sample.GetVMRequest
	7,  // 5: sample.SampleServer.GetVMStatus:input_type -> sample.GetVMStatusRequest
	9,  // 6: sample.SampleServer.DeleteVM:input_type -> sample.DeleteVMRequest
	10, // 7: sample.SampleServer.CheckNameAvailability:input_type -> sample.CheckNameAvailabilityRequest
	3,  // 8: sample.SampleServer.CreateVM:output_type -> sample.CreateVMResponse
	4,  // 9: sample.SampleServer.GetVMList:output_type -> sample.GetVMListResponse
	6,  // 10: sample.SampleServer.GetVM:output_type -> sample.GetVMResponse
	8,  // 11: sample.SampleServer.GetVMStatus:output_type -> sample.GetVMStatusResponse
	0,  // 12: sample.SampleServer.DeleteVM:output_type -> sample.Empty
	0,  // 13: sample.SampleServer.CheckNameAvailability:output_type -> sample.Empty
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_proto_sample_server_proto_init() }
func file_proto_sample_server_proto_init() {
	if File_proto_sample_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_sample_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_proto_sample_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Array); i {
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
		file_proto_sample_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateVMRequest); i {
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
		file_proto_sample_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateVMResponse); i {
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
		file_proto_sample_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVMListResponse); i {
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
		file_proto_sample_server_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVMRequest); i {
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
		file_proto_sample_server_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVMResponse); i {
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
		file_proto_sample_server_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVMStatusRequest); i {
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
		file_proto_sample_server_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVMStatusResponse); i {
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
		file_proto_sample_server_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteVMRequest); i {
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
		file_proto_sample_server_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckNameAvailabilityRequest); i {
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
			RawDescriptor: file_proto_sample_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_sample_server_proto_goTypes,
		DependencyIndexes: file_proto_sample_server_proto_depIdxs,
		MessageInfos:      file_proto_sample_server_proto_msgTypes,
	}.Build()
	File_proto_sample_server_proto = out.File
	file_proto_sample_server_proto_rawDesc = nil
	file_proto_sample_server_proto_goTypes = nil
	file_proto_sample_server_proto_depIdxs = nil
}
