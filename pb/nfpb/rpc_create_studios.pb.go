// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: nfpb/rpc_create_studios.proto

package nfpb

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

type CreateStudiosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Names []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
}

func (x *CreateStudiosRequest) Reset() {
	*x = CreateStudiosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nfpb_rpc_create_studios_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStudiosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStudiosRequest) ProtoMessage() {}

func (x *CreateStudiosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nfpb_rpc_create_studios_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStudiosRequest.ProtoReflect.Descriptor instead.
func (*CreateStudiosRequest) Descriptor() ([]byte, []int) {
	return file_nfpb_rpc_create_studios_proto_rawDescGZIP(), []int{0}
}

func (x *CreateStudiosRequest) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

type CreateStudiosResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Studios []*Studio `protobuf:"bytes,1,rep,name=studios,proto3" json:"studios,omitempty"`
}

func (x *CreateStudiosResponse) Reset() {
	*x = CreateStudiosResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nfpb_rpc_create_studios_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStudiosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStudiosResponse) ProtoMessage() {}

func (x *CreateStudiosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nfpb_rpc_create_studios_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStudiosResponse.ProtoReflect.Descriptor instead.
func (*CreateStudiosResponse) Descriptor() ([]byte, []int) {
	return file_nfpb_rpc_create_studios_proto_rawDescGZIP(), []int{1}
}

func (x *CreateStudiosResponse) GetStudios() []*Studio {
	if x != nil {
		return x.Studios
	}
	return nil
}

var File_nfpb_rpc_create_studios_proto protoreflect.FileDescriptor

var file_nfpb_rpc_create_studios_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6e, 0x66, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x6e, 0x66, 0x70, 0x62, 0x1a, 0x11, 0x6e, 0x66, 0x70, 0x62, 0x2f, 0x73, 0x74, 0x75, 0x64,
	0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x3f, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x26, 0x0a, 0x07, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x6e, 0x66, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x52, 0x07,
	0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x73, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6a, 0x2d, 0x79, 0x61, 0x63, 0x69, 0x6e, 0x65, 0x2d,
	0x66, 0x6c, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x6f, 0x6a, 0x6f, 0x2f, 0x70, 0x62, 0x2f,
	0x6e, 0x66, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nfpb_rpc_create_studios_proto_rawDescOnce sync.Once
	file_nfpb_rpc_create_studios_proto_rawDescData = file_nfpb_rpc_create_studios_proto_rawDesc
)

func file_nfpb_rpc_create_studios_proto_rawDescGZIP() []byte {
	file_nfpb_rpc_create_studios_proto_rawDescOnce.Do(func() {
		file_nfpb_rpc_create_studios_proto_rawDescData = protoimpl.X.CompressGZIP(file_nfpb_rpc_create_studios_proto_rawDescData)
	})
	return file_nfpb_rpc_create_studios_proto_rawDescData
}

var file_nfpb_rpc_create_studios_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_nfpb_rpc_create_studios_proto_goTypes = []interface{}{
	(*CreateStudiosRequest)(nil),  // 0: nfpb.CreateStudiosRequest
	(*CreateStudiosResponse)(nil), // 1: nfpb.CreateStudiosResponse
	(*Studio)(nil),                // 2: nfpb.Studio
}
var file_nfpb_rpc_create_studios_proto_depIdxs = []int32{
	2, // 0: nfpb.CreateStudiosResponse.studios:type_name -> nfpb.Studio
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_nfpb_rpc_create_studios_proto_init() }
func file_nfpb_rpc_create_studios_proto_init() {
	if File_nfpb_rpc_create_studios_proto != nil {
		return
	}
	file_nfpb_studio_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_nfpb_rpc_create_studios_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStudiosRequest); i {
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
		file_nfpb_rpc_create_studios_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStudiosResponse); i {
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
			RawDescriptor: file_nfpb_rpc_create_studios_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nfpb_rpc_create_studios_proto_goTypes,
		DependencyIndexes: file_nfpb_rpc_create_studios_proto_depIdxs,
		MessageInfos:      file_nfpb_rpc_create_studios_proto_msgTypes,
	}.Build()
	File_nfpb_rpc_create_studios_proto = out.File
	file_nfpb_rpc_create_studios_proto_rawDesc = nil
	file_nfpb_rpc_create_studios_proto_goTypes = nil
	file_nfpb_rpc_create_studios_proto_depIdxs = nil
}
