// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: v1/nfpb/rpc_get_all_studios.proto

package nfpbv1

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

type GetAllStudiosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNumber int32 `protobuf:"varint,1,opt,name=pageNumber,proto3" json:"pageNumber,omitempty"`
	PageSize   int32 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *GetAllStudiosRequest) Reset() {
	*x = GetAllStudiosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_nfpb_rpc_get_all_studios_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllStudiosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllStudiosRequest) ProtoMessage() {}

func (x *GetAllStudiosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_nfpb_rpc_get_all_studios_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllStudiosRequest.ProtoReflect.Descriptor instead.
func (*GetAllStudiosRequest) Descriptor() ([]byte, []int) {
	return file_v1_nfpb_rpc_get_all_studios_proto_rawDescGZIP(), []int{0}
}

func (x *GetAllStudiosRequest) GetPageNumber() int32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *GetAllStudiosRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type GetAllStudiosResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Studios []*Studio `protobuf:"bytes,1,rep,name=studios,proto3" json:"studios,omitempty"`
}

func (x *GetAllStudiosResponse) Reset() {
	*x = GetAllStudiosResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_nfpb_rpc_get_all_studios_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllStudiosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllStudiosResponse) ProtoMessage() {}

func (x *GetAllStudiosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_nfpb_rpc_get_all_studios_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllStudiosResponse.ProtoReflect.Descriptor instead.
func (*GetAllStudiosResponse) Descriptor() ([]byte, []int) {
	return file_v1_nfpb_rpc_get_all_studios_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllStudiosResponse) GetStudios() []*Studio {
	if x != nil {
		return x.Studios
	}
	return nil
}

var File_v1_nfpb_rpc_get_all_studios_proto protoreflect.FileDescriptor

var file_v1_nfpb_rpc_get_all_studios_proto_rawDesc = []byte{
	0x0a, 0x21, 0x76, 0x31, 0x2f, 0x6e, 0x66, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65,
	0x74, 0x5f, 0x61, 0x6c, 0x6c, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x76, 0x31, 0x2e, 0x6e, 0x66, 0x70, 0x62, 0x1a, 0x18, 0x76, 0x31,
	0x2f, 0x6e, 0x66, 0x70, 0x62, 0x2f, 0x6d, 0x73, 0x67, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x53, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x42, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x53, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x07, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x76, 0x31, 0x2e, 0x6e, 0x66, 0x70, 0x62, 0x2e, 0x53,
	0x74, 0x75, 0x64, 0x69, 0x6f, 0x52, 0x07, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x73, 0x42, 0x35,
	0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6a, 0x2d,
	0x79, 0x61, 0x63, 0x69, 0x6e, 0x65, 0x2d, 0x66, 0x6c, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67,
	0x6f, 0x6a, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x66, 0x70, 0x62, 0x3b, 0x6e,
	0x66, 0x70, 0x62, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_nfpb_rpc_get_all_studios_proto_rawDescOnce sync.Once
	file_v1_nfpb_rpc_get_all_studios_proto_rawDescData = file_v1_nfpb_rpc_get_all_studios_proto_rawDesc
)

func file_v1_nfpb_rpc_get_all_studios_proto_rawDescGZIP() []byte {
	file_v1_nfpb_rpc_get_all_studios_proto_rawDescOnce.Do(func() {
		file_v1_nfpb_rpc_get_all_studios_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_nfpb_rpc_get_all_studios_proto_rawDescData)
	})
	return file_v1_nfpb_rpc_get_all_studios_proto_rawDescData
}

var file_v1_nfpb_rpc_get_all_studios_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_nfpb_rpc_get_all_studios_proto_goTypes = []interface{}{
	(*GetAllStudiosRequest)(nil),  // 0: v1.nfpb.GetAllStudiosRequest
	(*GetAllStudiosResponse)(nil), // 1: v1.nfpb.GetAllStudiosResponse
	(*Studio)(nil),                // 2: v1.nfpb.Studio
}
var file_v1_nfpb_rpc_get_all_studios_proto_depIdxs = []int32{
	2, // 0: v1.nfpb.GetAllStudiosResponse.studios:type_name -> v1.nfpb.Studio
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v1_nfpb_rpc_get_all_studios_proto_init() }
func file_v1_nfpb_rpc_get_all_studios_proto_init() {
	if File_v1_nfpb_rpc_get_all_studios_proto != nil {
		return
	}
	file_v1_nfpb_msg_studio_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_nfpb_rpc_get_all_studios_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllStudiosRequest); i {
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
		file_v1_nfpb_rpc_get_all_studios_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllStudiosResponse); i {
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
			RawDescriptor: file_v1_nfpb_rpc_get_all_studios_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_nfpb_rpc_get_all_studios_proto_goTypes,
		DependencyIndexes: file_v1_nfpb_rpc_get_all_studios_proto_depIdxs,
		MessageInfos:      file_v1_nfpb_rpc_get_all_studios_proto_msgTypes,
	}.Build()
	File_v1_nfpb_rpc_get_all_studios_proto = out.File
	file_v1_nfpb_rpc_get_all_studios_proto_rawDesc = nil
	file_v1_nfpb_rpc_get_all_studios_proto_goTypes = nil
	file_v1_nfpb_rpc_get_all_studios_proto_depIdxs = nil
}
