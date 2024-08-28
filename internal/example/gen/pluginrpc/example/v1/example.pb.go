// Copyright 2024 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: pluginrpc/example/v1/example.proto

package examplev1

import (
	v1 "buf.build/gen/go/pluginrpc/pluginrpc/protocolbuffers/go/pluginrpc/v1"
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

// A request to echo the given message.
type EchoRequestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The message to echo back.
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EchoRequestRequest) Reset() {
	*x = EchoRequestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pluginrpc_example_v1_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoRequestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoRequestRequest) ProtoMessage() {}

func (x *EchoRequestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pluginrpc_example_v1_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoRequestRequest.ProtoReflect.Descriptor instead.
func (*EchoRequestRequest) Descriptor() ([]byte, []int) {
	return file_pluginrpc_example_v1_example_proto_rawDescGZIP(), []int{0}
}

func (x *EchoRequestRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// A response to echo.
type EchoRequestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The echoed message.
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EchoRequestResponse) Reset() {
	*x = EchoRequestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pluginrpc_example_v1_example_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoRequestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoRequestResponse) ProtoMessage() {}

func (x *EchoRequestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pluginrpc_example_v1_example_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoRequestResponse.ProtoReflect.Descriptor instead.
func (*EchoRequestResponse) Descriptor() ([]byte, []int) {
	return file_pluginrpc_example_v1_example_proto_rawDescGZIP(), []int{1}
}

func (x *EchoRequestResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// An error to echo back.
type EchoErrorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The error code to return as part of the error.
	Code v1.Code `protobuf:"varint,1,opt,name=code,proto3,enum=pluginrpc.v1.Code" json:"code,omitempty"`
	// The error message to return as part of the error.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EchoErrorRequest) Reset() {
	*x = EchoErrorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pluginrpc_example_v1_example_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoErrorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoErrorRequest) ProtoMessage() {}

func (x *EchoErrorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pluginrpc_example_v1_example_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoErrorRequest.ProtoReflect.Descriptor instead.
func (*EchoErrorRequest) Descriptor() ([]byte, []int) {
	return file_pluginrpc_example_v1_example_proto_rawDescGZIP(), []int{2}
}

func (x *EchoErrorRequest) GetCode() v1.Code {
	if x != nil {
		return x.Code
	}
	return v1.Code(0)
}

func (x *EchoErrorRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// A blank response.
type EchoErrorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EchoErrorResponse) Reset() {
	*x = EchoErrorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pluginrpc_example_v1_example_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoErrorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoErrorResponse) ProtoMessage() {}

func (x *EchoErrorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pluginrpc_example_v1_example_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoErrorResponse.ProtoReflect.Descriptor instead.
func (*EchoErrorResponse) Descriptor() ([]byte, []int) {
	return file_pluginrpc_example_v1_example_proto_rawDescGZIP(), []int{3}
}

// A request to echo a static list back The request is purposefully
// empty to demonstrate how pluginrpc works with empty requests.
type EchoListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EchoListRequest) Reset() {
	*x = EchoListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pluginrpc_example_v1_example_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoListRequest) ProtoMessage() {}

func (x *EchoListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pluginrpc_example_v1_example_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoListRequest.ProtoReflect.Descriptor instead.
func (*EchoListRequest) Descriptor() ([]byte, []int) {
	return file_pluginrpc_example_v1_example_proto_rawDescGZIP(), []int{4}
}

// A response that will always contain the list ["foo", "bar"].
type EchoListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list that will always be ["foo", "bar"].
	List []string `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *EchoListResponse) Reset() {
	*x = EchoListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pluginrpc_example_v1_example_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoListResponse) ProtoMessage() {}

func (x *EchoListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pluginrpc_example_v1_example_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoListResponse.ProtoReflect.Descriptor instead.
func (*EchoListResponse) Descriptor() ([]byte, []int) {
	return file_pluginrpc_example_v1_example_proto_rawDescGZIP(), []int{5}
}

func (x *EchoListResponse) GetList() []string {
	if x != nil {
		return x.List
	}
	return nil
}

var File_pluginrpc_example_v1_example_proto protoreflect.FileDescriptor

var file_pluginrpc_example_v1_example_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x72, 0x70, 0x63, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x72,
	0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x12, 0x45, 0x63, 0x68, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2f, 0x0a, 0x13, 0x45, 0x63, 0x68, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x54, 0x0a, 0x10, 0x45, 0x63, 0x68,
	0x6f, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x13, 0x0a, 0x11, 0x45, 0x63, 0x68, 0x6f, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x0a, 0x0f, 0x45, 0x63, 0x68, 0x6f, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x26, 0x0a, 0x10, 0x45, 0x63, 0x68, 0x6f, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x32,
	0xaa, 0x02, 0x0a, 0x0b, 0x45, 0x63, 0x68, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x62, 0x0a, 0x0b, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28,
	0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x09, 0x45, 0x63, 0x68, 0x6f, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x26, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x63, 0x68, 0x6f, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x59, 0x0a, 0x08, 0x45, 0x63, 0x68, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x25, 0x2e,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x72, 0x70, 0x63,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x63, 0x68, 0x6f,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xe7, 0x01, 0x0a,
	0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x72, 0x70, 0x63, 0x2e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x45, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x72,
	0x70, 0x63, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x72, 0x70,
	0x63, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x45, 0x58, 0xaa, 0x02, 0x14, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x14, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x72, 0x70, 0x63, 0x5c,
	0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x20, 0x50, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x72, 0x70, 0x63, 0x5c, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x16,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x72, 0x70, 0x63, 0x3a, 0x3a, 0x45, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pluginrpc_example_v1_example_proto_rawDescOnce sync.Once
	file_pluginrpc_example_v1_example_proto_rawDescData = file_pluginrpc_example_v1_example_proto_rawDesc
)

func file_pluginrpc_example_v1_example_proto_rawDescGZIP() []byte {
	file_pluginrpc_example_v1_example_proto_rawDescOnce.Do(func() {
		file_pluginrpc_example_v1_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_pluginrpc_example_v1_example_proto_rawDescData)
	})
	return file_pluginrpc_example_v1_example_proto_rawDescData
}

var file_pluginrpc_example_v1_example_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pluginrpc_example_v1_example_proto_goTypes = []any{
	(*EchoRequestRequest)(nil),  // 0: pluginrpc.example.v1.EchoRequestRequest
	(*EchoRequestResponse)(nil), // 1: pluginrpc.example.v1.EchoRequestResponse
	(*EchoErrorRequest)(nil),    // 2: pluginrpc.example.v1.EchoErrorRequest
	(*EchoErrorResponse)(nil),   // 3: pluginrpc.example.v1.EchoErrorResponse
	(*EchoListRequest)(nil),     // 4: pluginrpc.example.v1.EchoListRequest
	(*EchoListResponse)(nil),    // 5: pluginrpc.example.v1.EchoListResponse
	(v1.Code)(0),                // 6: pluginrpc.v1.Code
}
var file_pluginrpc_example_v1_example_proto_depIdxs = []int32{
	6, // 0: pluginrpc.example.v1.EchoErrorRequest.code:type_name -> pluginrpc.v1.Code
	0, // 1: pluginrpc.example.v1.EchoService.EchoRequest:input_type -> pluginrpc.example.v1.EchoRequestRequest
	2, // 2: pluginrpc.example.v1.EchoService.EchoError:input_type -> pluginrpc.example.v1.EchoErrorRequest
	4, // 3: pluginrpc.example.v1.EchoService.EchoList:input_type -> pluginrpc.example.v1.EchoListRequest
	1, // 4: pluginrpc.example.v1.EchoService.EchoRequest:output_type -> pluginrpc.example.v1.EchoRequestResponse
	3, // 5: pluginrpc.example.v1.EchoService.EchoError:output_type -> pluginrpc.example.v1.EchoErrorResponse
	5, // 6: pluginrpc.example.v1.EchoService.EchoList:output_type -> pluginrpc.example.v1.EchoListResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pluginrpc_example_v1_example_proto_init() }
func file_pluginrpc_example_v1_example_proto_init() {
	if File_pluginrpc_example_v1_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pluginrpc_example_v1_example_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*EchoRequestRequest); i {
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
		file_pluginrpc_example_v1_example_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*EchoRequestResponse); i {
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
		file_pluginrpc_example_v1_example_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*EchoErrorRequest); i {
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
		file_pluginrpc_example_v1_example_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*EchoErrorResponse); i {
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
		file_pluginrpc_example_v1_example_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*EchoListRequest); i {
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
		file_pluginrpc_example_v1_example_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*EchoListResponse); i {
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
			RawDescriptor: file_pluginrpc_example_v1_example_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pluginrpc_example_v1_example_proto_goTypes,
		DependencyIndexes: file_pluginrpc_example_v1_example_proto_depIdxs,
		MessageInfos:      file_pluginrpc_example_v1_example_proto_msgTypes,
	}.Build()
	File_pluginrpc_example_v1_example_proto = out.File
	file_pluginrpc_example_v1_example_proto_rawDesc = nil
	file_pluginrpc_example_v1_example_proto_goTypes = nil
	file_pluginrpc_example_v1_example_proto_depIdxs = nil
}
