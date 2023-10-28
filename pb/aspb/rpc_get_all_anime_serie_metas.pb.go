// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: aspb/rpc_get_all_anime_serie_metas.proto

package aspb

import (
	nfpb "github.com/dj-yacine-flutter/gojo/pb/nfpb"
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

type GetAllAnimeSerieMetasRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnimeID int64 `protobuf:"varint,1,opt,name=animeID,proto3" json:"animeID,omitempty"`
}

func (x *GetAllAnimeSerieMetasRequest) Reset() {
	*x = GetAllAnimeSerieMetasRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aspb_rpc_get_all_anime_serie_metas_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllAnimeSerieMetasRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllAnimeSerieMetasRequest) ProtoMessage() {}

func (x *GetAllAnimeSerieMetasRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aspb_rpc_get_all_anime_serie_metas_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllAnimeSerieMetasRequest.ProtoReflect.Descriptor instead.
func (*GetAllAnimeSerieMetasRequest) Descriptor() ([]byte, []int) {
	return file_aspb_rpc_get_all_anime_serie_metas_proto_rawDescGZIP(), []int{0}
}

func (x *GetAllAnimeSerieMetasRequest) GetAnimeID() int64 {
	if x != nil {
		return x.AnimeID
	}
	return 0
}

type GetAllAnimeSerieMetasResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metas []*nfpb.AnimeMetaResponse `protobuf:"bytes,1,rep,name=metas,proto3" json:"metas,omitempty"`
}

func (x *GetAllAnimeSerieMetasResponse) Reset() {
	*x = GetAllAnimeSerieMetasResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aspb_rpc_get_all_anime_serie_metas_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllAnimeSerieMetasResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllAnimeSerieMetasResponse) ProtoMessage() {}

func (x *GetAllAnimeSerieMetasResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aspb_rpc_get_all_anime_serie_metas_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllAnimeSerieMetasResponse.ProtoReflect.Descriptor instead.
func (*GetAllAnimeSerieMetasResponse) Descriptor() ([]byte, []int) {
	return file_aspb_rpc_get_all_anime_serie_metas_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllAnimeSerieMetasResponse) GetMetas() []*nfpb.AnimeMetaResponse {
	if x != nil {
		return x.Metas
	}
	return nil
}

var File_aspb_rpc_get_all_anime_serie_metas_proto protoreflect.FileDescriptor

var file_aspb_rpc_get_all_anime_serie_metas_proto_rawDesc = []byte{
	0x0a, 0x28, 0x61, 0x73, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x61,
	0x6c, 0x6c, 0x5f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x5f, 0x6d,
	0x65, 0x74, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x73, 0x70, 0x62,
	0x1a, 0x0f, 0x6e, 0x66, 0x70, 0x62, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x38, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x6e, 0x69, 0x6d, 0x65,
	0x53, 0x65, 0x72, 0x69, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x49, 0x44, 0x22, 0x4e, 0x0a, 0x1d, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x4d,
	0x65, 0x74, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x05,
	0x6d, 0x65, 0x74, 0x61, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6e, 0x66,
	0x70, 0x62, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x42, 0x2b, 0x5a, 0x29, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6a, 0x2d, 0x79, 0x61, 0x63,
	0x69, 0x6e, 0x65, 0x2d, 0x66, 0x6c, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x6f, 0x6a, 0x6f,
	0x2f, 0x70, 0x62, 0x2f, 0x61, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aspb_rpc_get_all_anime_serie_metas_proto_rawDescOnce sync.Once
	file_aspb_rpc_get_all_anime_serie_metas_proto_rawDescData = file_aspb_rpc_get_all_anime_serie_metas_proto_rawDesc
)

func file_aspb_rpc_get_all_anime_serie_metas_proto_rawDescGZIP() []byte {
	file_aspb_rpc_get_all_anime_serie_metas_proto_rawDescOnce.Do(func() {
		file_aspb_rpc_get_all_anime_serie_metas_proto_rawDescData = protoimpl.X.CompressGZIP(file_aspb_rpc_get_all_anime_serie_metas_proto_rawDescData)
	})
	return file_aspb_rpc_get_all_anime_serie_metas_proto_rawDescData
}

var file_aspb_rpc_get_all_anime_serie_metas_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_aspb_rpc_get_all_anime_serie_metas_proto_goTypes = []interface{}{
	(*GetAllAnimeSerieMetasRequest)(nil),  // 0: aspb.GetAllAnimeSerieMetasRequest
	(*GetAllAnimeSerieMetasResponse)(nil), // 1: aspb.GetAllAnimeSerieMetasResponse
	(*nfpb.AnimeMetaResponse)(nil),        // 2: nfpb.AnimeMetaResponse
}
var file_aspb_rpc_get_all_anime_serie_metas_proto_depIdxs = []int32{
	2, // 0: aspb.GetAllAnimeSerieMetasResponse.metas:type_name -> nfpb.AnimeMetaResponse
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_aspb_rpc_get_all_anime_serie_metas_proto_init() }
func file_aspb_rpc_get_all_anime_serie_metas_proto_init() {
	if File_aspb_rpc_get_all_anime_serie_metas_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aspb_rpc_get_all_anime_serie_metas_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllAnimeSerieMetasRequest); i {
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
		file_aspb_rpc_get_all_anime_serie_metas_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllAnimeSerieMetasResponse); i {
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
			RawDescriptor: file_aspb_rpc_get_all_anime_serie_metas_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aspb_rpc_get_all_anime_serie_metas_proto_goTypes,
		DependencyIndexes: file_aspb_rpc_get_all_anime_serie_metas_proto_depIdxs,
		MessageInfos:      file_aspb_rpc_get_all_anime_serie_metas_proto_msgTypes,
	}.Build()
	File_aspb_rpc_get_all_anime_serie_metas_proto = out.File
	file_aspb_rpc_get_all_anime_serie_metas_proto_rawDesc = nil
	file_aspb_rpc_get_all_anime_serie_metas_proto_goTypes = nil
	file_aspb_rpc_get_all_anime_serie_metas_proto_depIdxs = nil
}
