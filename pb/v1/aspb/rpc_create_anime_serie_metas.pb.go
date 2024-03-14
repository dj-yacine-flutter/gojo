// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: v1/aspb/rpc_create_anime_serie_metas.proto

package aspbv1

import (
	nfpb "github.com/dj-yacine-flutter/gojo/pb/v1/nfpb"
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

type CreateAnimeSerieMetasRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnimeID    int64                    `protobuf:"varint,1,opt,name=animeID,proto3" json:"animeID,omitempty"`
	AnimeMetas []*nfpb.AnimeMetaRequest `protobuf:"bytes,2,rep,name=animeMetas,proto3" json:"animeMetas,omitempty"`
}

func (x *CreateAnimeSerieMetasRequest) Reset() {
	*x = CreateAnimeSerieMetasRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_aspb_rpc_create_anime_serie_metas_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAnimeSerieMetasRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAnimeSerieMetasRequest) ProtoMessage() {}

func (x *CreateAnimeSerieMetasRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_aspb_rpc_create_anime_serie_metas_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAnimeSerieMetasRequest.ProtoReflect.Descriptor instead.
func (*CreateAnimeSerieMetasRequest) Descriptor() ([]byte, []int) {
	return file_v1_aspb_rpc_create_anime_serie_metas_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAnimeSerieMetasRequest) GetAnimeID() int64 {
	if x != nil {
		return x.AnimeID
	}
	return 0
}

func (x *CreateAnimeSerieMetasRequest) GetAnimeMetas() []*nfpb.AnimeMetaRequest {
	if x != nil {
		return x.AnimeMetas
	}
	return nil
}

type CreateAnimeSerieMetasResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnimeID    int64                     `protobuf:"varint,1,opt,name=animeID,proto3" json:"animeID,omitempty"`
	AnimeMetas []*nfpb.AnimeMetaResponse `protobuf:"bytes,2,rep,name=animeMetas,proto3" json:"animeMetas,omitempty"`
}

func (x *CreateAnimeSerieMetasResponse) Reset() {
	*x = CreateAnimeSerieMetasResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_aspb_rpc_create_anime_serie_metas_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAnimeSerieMetasResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAnimeSerieMetasResponse) ProtoMessage() {}

func (x *CreateAnimeSerieMetasResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_aspb_rpc_create_anime_serie_metas_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAnimeSerieMetasResponse.ProtoReflect.Descriptor instead.
func (*CreateAnimeSerieMetasResponse) Descriptor() ([]byte, []int) {
	return file_v1_aspb_rpc_create_anime_serie_metas_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAnimeSerieMetasResponse) GetAnimeID() int64 {
	if x != nil {
		return x.AnimeID
	}
	return 0
}

func (x *CreateAnimeSerieMetasResponse) GetAnimeMetas() []*nfpb.AnimeMetaResponse {
	if x != nil {
		return x.AnimeMetas
	}
	return nil
}

var File_v1_aspb_rpc_create_anime_serie_metas_proto protoreflect.FileDescriptor

var file_v1_aspb_rpc_create_anime_serie_metas_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x76, 0x31, 0x2f, 0x61, 0x73, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65,
	0x5f, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x76, 0x31,
	0x2e, 0x61, 0x73, 0x70, 0x62, 0x76, 0x31, 0x1a, 0x16, 0x76, 0x31, 0x2f, 0x6e, 0x66, 0x70, 0x62,
	0x2f, 0x6d, 0x73, 0x67, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x75, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x53, 0x65,
	0x72, 0x69, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x49, 0x44, 0x12, 0x3b, 0x0a, 0x0a, 0x61, 0x6e, 0x69,
	0x6d, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x76, 0x31, 0x2e, 0x6e, 0x66, 0x70, 0x62, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d,
	0x65, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0a, 0x61, 0x6e, 0x69, 0x6d,
	0x65, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x22, 0x77, 0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x6e, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x6e, 0x69, 0x6d, 0x65,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x49,
	0x44, 0x12, 0x3c, 0x0a, 0x0a, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x76, 0x31, 0x2e, 0x6e, 0x66, 0x70, 0x62, 0x76,
	0x31, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x0a, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x42,
	0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6a,
	0x2d, 0x79, 0x61, 0x63, 0x69, 0x6e, 0x65, 0x2d, 0x66, 0x6c, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f,
	0x67, 0x6f, 0x6a, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x73, 0x70, 0x62, 0x3b,
	0x61, 0x73, 0x70, 0x62, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_aspb_rpc_create_anime_serie_metas_proto_rawDescOnce sync.Once
	file_v1_aspb_rpc_create_anime_serie_metas_proto_rawDescData = file_v1_aspb_rpc_create_anime_serie_metas_proto_rawDesc
)

func file_v1_aspb_rpc_create_anime_serie_metas_proto_rawDescGZIP() []byte {
	file_v1_aspb_rpc_create_anime_serie_metas_proto_rawDescOnce.Do(func() {
		file_v1_aspb_rpc_create_anime_serie_metas_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_aspb_rpc_create_anime_serie_metas_proto_rawDescData)
	})
	return file_v1_aspb_rpc_create_anime_serie_metas_proto_rawDescData
}

var file_v1_aspb_rpc_create_anime_serie_metas_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_aspb_rpc_create_anime_serie_metas_proto_goTypes = []interface{}{
	(*CreateAnimeSerieMetasRequest)(nil),  // 0: v1.aspbv1.CreateAnimeSerieMetasRequest
	(*CreateAnimeSerieMetasResponse)(nil), // 1: v1.aspbv1.CreateAnimeSerieMetasResponse
	(*nfpb.AnimeMetaRequest)(nil),         // 2: v1.nfpbv1.AnimeMetaRequest
	(*nfpb.AnimeMetaResponse)(nil),        // 3: v1.nfpbv1.AnimeMetaResponse
}
var file_v1_aspb_rpc_create_anime_serie_metas_proto_depIdxs = []int32{
	2, // 0: v1.aspbv1.CreateAnimeSerieMetasRequest.animeMetas:type_name -> v1.nfpbv1.AnimeMetaRequest
	3, // 1: v1.aspbv1.CreateAnimeSerieMetasResponse.animeMetas:type_name -> v1.nfpbv1.AnimeMetaResponse
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_v1_aspb_rpc_create_anime_serie_metas_proto_init() }
func file_v1_aspb_rpc_create_anime_serie_metas_proto_init() {
	if File_v1_aspb_rpc_create_anime_serie_metas_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_aspb_rpc_create_anime_serie_metas_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAnimeSerieMetasRequest); i {
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
		file_v1_aspb_rpc_create_anime_serie_metas_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAnimeSerieMetasResponse); i {
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
			RawDescriptor: file_v1_aspb_rpc_create_anime_serie_metas_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_aspb_rpc_create_anime_serie_metas_proto_goTypes,
		DependencyIndexes: file_v1_aspb_rpc_create_anime_serie_metas_proto_depIdxs,
		MessageInfos:      file_v1_aspb_rpc_create_anime_serie_metas_proto_msgTypes,
	}.Build()
	File_v1_aspb_rpc_create_anime_serie_metas_proto = out.File
	file_v1_aspb_rpc_create_anime_serie_metas_proto_rawDesc = nil
	file_v1_aspb_rpc_create_anime_serie_metas_proto_goTypes = nil
	file_v1_aspb_rpc_create_anime_serie_metas_proto_depIdxs = nil
}
