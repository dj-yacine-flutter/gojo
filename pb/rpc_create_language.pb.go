// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: rpc_create_language.proto

package pb

import (
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

type CreateLanguageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Languages []*LanguageRequest `protobuf:"bytes,1,rep,name=languages,proto3" json:"languages,omitempty"`
}

func (x *CreateLanguageRequest) Reset() {
	*x = CreateLanguageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_create_language_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLanguageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLanguageRequest) ProtoMessage() {}

func (x *CreateLanguageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_create_language_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLanguageRequest.ProtoReflect.Descriptor instead.
func (*CreateLanguageRequest) Descriptor() ([]byte, []int) {
	return file_rpc_create_language_proto_rawDescGZIP(), []int{0}
}

func (x *CreateLanguageRequest) GetLanguages() []*LanguageRequest {
	if x != nil {
		return x.Languages
	}
	return nil
}

type CreateLanguageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Languages []*LanguageResponse `protobuf:"bytes,1,rep,name=languages,proto3" json:"languages,omitempty"`
}

func (x *CreateLanguageResponse) Reset() {
	*x = CreateLanguageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_create_language_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLanguageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLanguageResponse) ProtoMessage() {}

func (x *CreateLanguageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_create_language_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLanguageResponse.ProtoReflect.Descriptor instead.
func (*CreateLanguageResponse) Descriptor() ([]byte, []int) {
	return file_rpc_create_language_proto_rawDescGZIP(), []int{1}
}

func (x *CreateLanguageResponse) GetLanguages() []*LanguageResponse {
	if x != nil {
		return x.Languages
	}
	return nil
}

var File_rpc_create_language_proto protoreflect.FileDescriptor

var file_rpc_create_language_proto_rawDesc = []byte{
	0x0a, 0x19, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a,
	0x0e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x4a, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x09, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x62,
	0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x09, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x22, 0x4c, 0x0a, 0x16, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x09,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6a, 0x2d, 0x79, 0x61, 0x63, 0x69, 0x6e,
	0x65, 0x2d, 0x66, 0x6c, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x6f, 0x6a, 0x6f, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_create_language_proto_rawDescOnce sync.Once
	file_rpc_create_language_proto_rawDescData = file_rpc_create_language_proto_rawDesc
)

func file_rpc_create_language_proto_rawDescGZIP() []byte {
	file_rpc_create_language_proto_rawDescOnce.Do(func() {
		file_rpc_create_language_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_create_language_proto_rawDescData)
	})
	return file_rpc_create_language_proto_rawDescData
}

var file_rpc_create_language_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_create_language_proto_goTypes = []interface{}{
	(*CreateLanguageRequest)(nil),  // 0: pb.CreateLanguageRequest
	(*CreateLanguageResponse)(nil), // 1: pb.CreateLanguageResponse
	(*LanguageRequest)(nil),        // 2: pb.LanguageRequest
	(*LanguageResponse)(nil),       // 3: pb.LanguageResponse
}
var file_rpc_create_language_proto_depIdxs = []int32{
	2, // 0: pb.CreateLanguageRequest.languages:type_name -> pb.LanguageRequest
	3, // 1: pb.CreateLanguageResponse.languages:type_name -> pb.LanguageResponse
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rpc_create_language_proto_init() }
func file_rpc_create_language_proto_init() {
	if File_rpc_create_language_proto != nil {
		return
	}
	file_language_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_create_language_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLanguageRequest); i {
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
		file_rpc_create_language_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLanguageResponse); i {
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
			RawDescriptor: file_rpc_create_language_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_create_language_proto_goTypes,
		DependencyIndexes: file_rpc_create_language_proto_depIdxs,
		MessageInfos:      file_rpc_create_language_proto_msgTypes,
	}.Build()
	File_rpc_create_language_proto = out.File
	file_rpc_create_language_proto_rawDesc = nil
	file_rpc_create_language_proto_goTypes = nil
	file_rpc_create_language_proto_depIdxs = nil
}
