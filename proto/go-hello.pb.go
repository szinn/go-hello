// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.25.3
// source: go-hello.proto

package proto

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IsHealthyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *IsHealthyRequest) Reset() {
	*x = IsHealthyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_hello_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsHealthyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsHealthyRequest) ProtoMessage() {}

func (x *IsHealthyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_hello_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsHealthyRequest.ProtoReflect.Descriptor instead.
func (*IsHealthyRequest) Descriptor() ([]byte, []int) {
	return file_go_hello_proto_rawDescGZIP(), []int{0}
}

type IsHealthyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Healthy bool `protobuf:"varint,1,opt,name=healthy,proto3" json:"healthy,omitempty"`
}

func (x *IsHealthyResponse) Reset() {
	*x = IsHealthyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_hello_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsHealthyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsHealthyResponse) ProtoMessage() {}

func (x *IsHealthyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_go_hello_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsHealthyResponse.ProtoReflect.Descriptor instead.
func (*IsHealthyResponse) Descriptor() ([]byte, []int) {
	return file_go_hello_proto_rawDescGZIP(), []int{1}
}

func (x *IsHealthyResponse) GetHealthy() bool {
	if x != nil {
		return x.Healthy
	}
	return false
}

var File_go_hello_proto protoreflect.FileDescriptor

var file_go_hello_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x67, 0x6f, 0x2d, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x12, 0x0a, 0x10, 0x49, 0x73, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2d, 0x0a, 0x11, 0x49,
	0x73, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x32, 0x4b, 0x0a, 0x07, 0x47, 0x6f,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x40, 0x0a, 0x09, 0x49, 0x73, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x79, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x73, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x73, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x7a, 0x69, 0x6e, 0x6e, 0x2f, 0x67, 0x6f, 0x2d, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_go_hello_proto_rawDescOnce sync.Once
	file_go_hello_proto_rawDescData = file_go_hello_proto_rawDesc
)

func file_go_hello_proto_rawDescGZIP() []byte {
	file_go_hello_proto_rawDescOnce.Do(func() {
		file_go_hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_hello_proto_rawDescData)
	})
	return file_go_hello_proto_rawDescData
}

var (
	file_go_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
	file_go_hello_proto_goTypes  = []interface{}{
		(*IsHealthyRequest)(nil),  // 0: proto.IsHealthyRequest
		(*IsHealthyResponse)(nil), // 1: proto.IsHealthyResponse
	}
)
var file_go_hello_proto_depIdxs = []int32{
	0, // 0: proto.GoHello.IsHealthy:input_type -> proto.IsHealthyRequest
	1, // 1: proto.GoHello.IsHealthy:output_type -> proto.IsHealthyResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_go_hello_proto_init() }
func file_go_hello_proto_init() {
	if File_go_hello_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_hello_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsHealthyRequest); i {
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
		file_go_hello_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsHealthyResponse); i {
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
			RawDescriptor: file_go_hello_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_go_hello_proto_goTypes,
		DependencyIndexes: file_go_hello_proto_depIdxs,
		MessageInfos:      file_go_hello_proto_msgTypes,
	}.Build()
	File_go_hello_proto = out.File
	file_go_hello_proto_rawDesc = nil
	file_go_hello_proto_goTypes = nil
	file_go_hello_proto_depIdxs = nil
}
