// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: rpc_get_all_anime_metas.proto

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

type AllAnimeMetas struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Language *Language     `protobuf:"bytes,1,opt,name=language,proto3" json:"language,omitempty"`
	Meta     *MetaResponse `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *AllAnimeMetas) Reset() {
	*x = AllAnimeMetas{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_get_all_anime_metas_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllAnimeMetas) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllAnimeMetas) ProtoMessage() {}

func (x *AllAnimeMetas) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_get_all_anime_metas_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllAnimeMetas.ProtoReflect.Descriptor instead.
func (*AllAnimeMetas) Descriptor() ([]byte, []int) {
	return file_rpc_get_all_anime_metas_proto_rawDescGZIP(), []int{0}
}

func (x *AllAnimeMetas) GetLanguage() *Language {
	if x != nil {
		return x.Language
	}
	return nil
}

func (x *AllAnimeMetas) GetMeta() *MetaResponse {
	if x != nil {
		return x.Meta
	}
	return nil
}

type GetAllAnimeMetasRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnimeID int64 `protobuf:"varint,1,opt,name=animeID,proto3" json:"animeID,omitempty"`
}

func (x *GetAllAnimeMetasRequest) Reset() {
	*x = GetAllAnimeMetasRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_get_all_anime_metas_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllAnimeMetasRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllAnimeMetasRequest) ProtoMessage() {}

func (x *GetAllAnimeMetasRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_get_all_anime_metas_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllAnimeMetasRequest.ProtoReflect.Descriptor instead.
func (*GetAllAnimeMetasRequest) Descriptor() ([]byte, []int) {
	return file_rpc_get_all_anime_metas_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllAnimeMetasRequest) GetAnimeID() int64 {
	if x != nil {
		return x.AnimeID
	}
	return 0
}

type GetAllAnimeMetasResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metas []*AllAnimeMetas `protobuf:"bytes,1,rep,name=metas,proto3" json:"metas,omitempty"`
}

func (x *GetAllAnimeMetasResponse) Reset() {
	*x = GetAllAnimeMetasResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_get_all_anime_metas_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllAnimeMetasResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllAnimeMetasResponse) ProtoMessage() {}

func (x *GetAllAnimeMetasResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_get_all_anime_metas_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllAnimeMetasResponse.ProtoReflect.Descriptor instead.
func (*GetAllAnimeMetasResponse) Descriptor() ([]byte, []int) {
	return file_rpc_get_all_anime_metas_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllAnimeMetasResponse) GetMetas() []*AllAnimeMetas {
	if x != nil {
		return x.Metas
	}
	return nil
}

var File_rpc_get_all_anime_metas_proto protoreflect.FileDescriptor

var file_rpc_get_all_anime_metas_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x6c, 0x6c, 0x5f, 0x61, 0x6e,
	0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x1a, 0x0a, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x5f, 0x0a, 0x0d, 0x41, 0x6c, 0x6c, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x73,
	0x12, 0x28, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x6d, 0x65,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61,
	0x22, 0x33, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d,
	0x65, 0x74, 0x61, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x6e, 0x69, 0x6d, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x6e,
	0x69, 0x6d, 0x65, 0x49, 0x44, 0x22, 0x43, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41,
	0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x27, 0x0a, 0x05, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x6c, 0x6c, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x65,
	0x74, 0x61, 0x73, 0x52, 0x05, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6a, 0x2d, 0x79, 0x61, 0x63, 0x69,
	0x6e, 0x65, 0x2d, 0x66, 0x6c, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x6f, 0x6a, 0x6f, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_get_all_anime_metas_proto_rawDescOnce sync.Once
	file_rpc_get_all_anime_metas_proto_rawDescData = file_rpc_get_all_anime_metas_proto_rawDesc
)

func file_rpc_get_all_anime_metas_proto_rawDescGZIP() []byte {
	file_rpc_get_all_anime_metas_proto_rawDescOnce.Do(func() {
		file_rpc_get_all_anime_metas_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_get_all_anime_metas_proto_rawDescData)
	})
	return file_rpc_get_all_anime_metas_proto_rawDescData
}

var file_rpc_get_all_anime_metas_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_rpc_get_all_anime_metas_proto_goTypes = []interface{}{
	(*AllAnimeMetas)(nil),            // 0: pb.AllAnimeMetas
	(*GetAllAnimeMetasRequest)(nil),  // 1: pb.GetAllAnimeMetasRequest
	(*GetAllAnimeMetasResponse)(nil), // 2: pb.GetAllAnimeMetasResponse
	(*Language)(nil),                 // 3: pb.Language
	(*MetaResponse)(nil),             // 4: pb.MetaResponse
}
var file_rpc_get_all_anime_metas_proto_depIdxs = []int32{
	3, // 0: pb.AllAnimeMetas.language:type_name -> pb.Language
	4, // 1: pb.AllAnimeMetas.meta:type_name -> pb.MetaResponse
	0, // 2: pb.GetAllAnimeMetasResponse.metas:type_name -> pb.AllAnimeMetas
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_rpc_get_all_anime_metas_proto_init() }
func file_rpc_get_all_anime_metas_proto_init() {
	if File_rpc_get_all_anime_metas_proto != nil {
		return
	}
	file_meta_proto_init()
	file_language_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_get_all_anime_metas_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllAnimeMetas); i {
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
		file_rpc_get_all_anime_metas_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllAnimeMetasRequest); i {
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
		file_rpc_get_all_anime_metas_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllAnimeMetasResponse); i {
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
			RawDescriptor: file_rpc_get_all_anime_metas_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_get_all_anime_metas_proto_goTypes,
		DependencyIndexes: file_rpc_get_all_anime_metas_proto_depIdxs,
		MessageInfos:      file_rpc_get_all_anime_metas_proto_msgTypes,
	}.Build()
	File_rpc_get_all_anime_metas_proto = out.File
	file_rpc_get_all_anime_metas_proto_rawDesc = nil
	file_rpc_get_all_anime_metas_proto_goTypes = nil
	file_rpc_get_all_anime_metas_proto_depIdxs = nil
}
