// Copyright 2023 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: idl/demo.proto

package multiservice

import (
	context "context"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_demo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_idl_demo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_idl_demo_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_demo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_idl_demo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_idl_demo_proto_rawDescGZIP(), []int{1}
}

func (x *Reply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_idl_demo_proto protoreflect.FileDescriptor

var file_idl_demo_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x69, 0x64, 0x6c, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0c, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x1d,
	0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x21, 0x0a,
	0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x32, 0x45, 0x0a, 0x08, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x12, 0x39, 0x0a, 0x05,
	0x43, 0x68, 0x61, 0x74, 0x41, 0x12, 0x15, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x6d,
	0x75, 0x6c, 0x74, 0x69, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x32, 0x43, 0x0a, 0x08, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x42, 0x12, 0x37, 0x0a, 0x05, 0x43, 0x68, 0x61, 0x74, 0x42, 0x12, 0x15, 0x2e, 0x6d,
	0x75, 0x6c, 0x74, 0x69, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x28, 0x01, 0x42, 0x4d, 0x5a, 0x4b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x77, 0x65, 0x67, 0x6f, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x6d,
	0x75, 0x6c, 0x74, 0x69, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_idl_demo_proto_rawDescOnce sync.Once
	file_idl_demo_proto_rawDescData = file_idl_demo_proto_rawDesc
)

func file_idl_demo_proto_rawDescGZIP() []byte {
	file_idl_demo_proto_rawDescOnce.Do(func() {
		file_idl_demo_proto_rawDescData = protoimpl.X.CompressGZIP(file_idl_demo_proto_rawDescData)
	})
	return file_idl_demo_proto_rawDescData
}

var file_idl_demo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_idl_demo_proto_goTypes = []interface{}{
	(*Request)(nil), // 0: multiservice.Request
	(*Reply)(nil),   // 1: multiservice.Reply
}
var file_idl_demo_proto_depIdxs = []int32{
	0, // 0: multiservice.ServiceA.ChatA:input_type -> multiservice.Request
	0, // 1: multiservice.ServiceB.ChatB:input_type -> multiservice.Request
	1, // 2: multiservice.ServiceA.ChatA:output_type -> multiservice.Reply
	1, // 3: multiservice.ServiceB.ChatB:output_type -> multiservice.Reply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_idl_demo_proto_init() }
func file_idl_demo_proto_init() {
	if File_idl_demo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_idl_demo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_idl_demo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply); i {
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
			RawDescriptor: file_idl_demo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_idl_demo_proto_goTypes,
		DependencyIndexes: file_idl_demo_proto_depIdxs,
		MessageInfos:      file_idl_demo_proto_msgTypes,
	}.Build()
	File_idl_demo_proto = out.File
	file_idl_demo_proto_rawDesc = nil
	file_idl_demo_proto_goTypes = nil
	file_idl_demo_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.6.2. DO NOT EDIT.

type ServiceA interface {
	ChatA(stream ServiceA_ChatAServer) (err error)
}

type ServiceA_ChatAServer interface {
	streaming.Stream
	Recv() (*Request, error)
	Send(*Reply) error
}

type ServiceB interface {
	ChatB(stream ServiceB_ChatBServer) (err error)
}

type ServiceB_ChatBServer interface {
	streaming.Stream
	Recv() (*Request, error)
	SendAndClose(*Reply) error
}
