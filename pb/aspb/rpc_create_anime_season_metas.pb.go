// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: aspb/rpc_create_anime_season_metas.proto

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

type CreateAnimeSeasonMetasRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeasonID    int32                    `protobuf:"varint,1,opt,name=seasonID,proto3" json:"seasonID,omitempty"`
	SeasonMetas []*nfpb.AnimeMetaRequest `protobuf:"bytes,2,rep,name=seasonMetas,proto3" json:"seasonMetas,omitempty"`
}

func (x *CreateAnimeSeasonMetasRequest) Reset() {
	*x = CreateAnimeSeasonMetasRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aspb_rpc_create_anime_season_metas_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAnimeSeasonMetasRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAnimeSeasonMetasRequest) ProtoMessage() {}

func (x *CreateAnimeSeasonMetasRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aspb_rpc_create_anime_season_metas_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAnimeSeasonMetasRequest.ProtoReflect.Descriptor instead.
func (*CreateAnimeSeasonMetasRequest) Descriptor() ([]byte, []int) {
	return file_aspb_rpc_create_anime_season_metas_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAnimeSeasonMetasRequest) GetSeasonID() int32 {
	if x != nil {
		return x.SeasonID
	}
	return 0
}

func (x *CreateAnimeSeasonMetasRequest) GetSeasonMetas() []*nfpb.AnimeMetaRequest {
	if x != nil {
		return x.SeasonMetas
	}
	return nil
}

type CreateAnimeSeasonMetasResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Season      *AnimeSeasonResponse      `protobuf:"bytes,1,opt,name=season,proto3" json:"season,omitempty"`
	SeasonMetas []*nfpb.AnimeMetaResponse `protobuf:"bytes,2,rep,name=seasonMetas,proto3" json:"seasonMetas,omitempty"`
}

func (x *CreateAnimeSeasonMetasResponse) Reset() {
	*x = CreateAnimeSeasonMetasResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aspb_rpc_create_anime_season_metas_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAnimeSeasonMetasResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAnimeSeasonMetasResponse) ProtoMessage() {}

func (x *CreateAnimeSeasonMetasResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aspb_rpc_create_anime_season_metas_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAnimeSeasonMetasResponse.ProtoReflect.Descriptor instead.
func (*CreateAnimeSeasonMetasResponse) Descriptor() ([]byte, []int) {
	return file_aspb_rpc_create_anime_season_metas_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAnimeSeasonMetasResponse) GetSeason() *AnimeSeasonResponse {
	if x != nil {
		return x.Season
	}
	return nil
}

func (x *CreateAnimeSeasonMetasResponse) GetSeasonMetas() []*nfpb.AnimeMetaResponse {
	if x != nil {
		return x.SeasonMetas
	}
	return nil
}

var File_aspb_rpc_create_anime_season_metas_proto protoreflect.FileDescriptor

var file_aspb_rpc_create_anime_season_metas_proto_rawDesc = []byte{
	0x0a, 0x28, 0x61, 0x73, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x5f, 0x6d,
	0x65, 0x74, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x73, 0x70, 0x62,
	0x1a, 0x17, 0x61, 0x73, 0x70, 0x62, 0x2f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x6e, 0x66, 0x70, 0x62, 0x2f,
	0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x75, 0x0a, 0x1d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x4d,
	0x65, 0x74, 0x61, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x38, 0x0a, 0x0b, 0x73, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6e,
	0x66, 0x70, 0x62, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61,
	0x73, 0x22, 0x8e, 0x01, 0x0a, 0x1e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d,
	0x65, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x73, 0x70, 0x62, 0x2e, 0x41, 0x6e, 0x69, 0x6d,
	0x65, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x06, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0b, 0x73, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6e,
	0x66, 0x70, 0x62, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0b, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x4d, 0x65, 0x74,
	0x61, 0x73, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x64, 0x6a, 0x2d, 0x79, 0x61, 0x63, 0x69, 0x6e, 0x65, 0x2d, 0x66, 0x6c, 0x75, 0x74, 0x74,
	0x65, 0x72, 0x2f, 0x67, 0x6f, 0x6a, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x61, 0x73, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aspb_rpc_create_anime_season_metas_proto_rawDescOnce sync.Once
	file_aspb_rpc_create_anime_season_metas_proto_rawDescData = file_aspb_rpc_create_anime_season_metas_proto_rawDesc
)

func file_aspb_rpc_create_anime_season_metas_proto_rawDescGZIP() []byte {
	file_aspb_rpc_create_anime_season_metas_proto_rawDescOnce.Do(func() {
		file_aspb_rpc_create_anime_season_metas_proto_rawDescData = protoimpl.X.CompressGZIP(file_aspb_rpc_create_anime_season_metas_proto_rawDescData)
	})
	return file_aspb_rpc_create_anime_season_metas_proto_rawDescData
}

var file_aspb_rpc_create_anime_season_metas_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_aspb_rpc_create_anime_season_metas_proto_goTypes = []interface{}{
	(*CreateAnimeSeasonMetasRequest)(nil),  // 0: aspb.CreateAnimeSeasonMetasRequest
	(*CreateAnimeSeasonMetasResponse)(nil), // 1: aspb.CreateAnimeSeasonMetasResponse
	(*nfpb.AnimeMetaRequest)(nil),          // 2: nfpb.AnimeMetaRequest
	(*AnimeSeasonResponse)(nil),            // 3: aspb.AnimeSeasonResponse
	(*nfpb.AnimeMetaResponse)(nil),         // 4: nfpb.AnimeMetaResponse
}
var file_aspb_rpc_create_anime_season_metas_proto_depIdxs = []int32{
	2, // 0: aspb.CreateAnimeSeasonMetasRequest.seasonMetas:type_name -> nfpb.AnimeMetaRequest
	3, // 1: aspb.CreateAnimeSeasonMetasResponse.season:type_name -> aspb.AnimeSeasonResponse
	4, // 2: aspb.CreateAnimeSeasonMetasResponse.seasonMetas:type_name -> nfpb.AnimeMetaResponse
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_aspb_rpc_create_anime_season_metas_proto_init() }
func file_aspb_rpc_create_anime_season_metas_proto_init() {
	if File_aspb_rpc_create_anime_season_metas_proto != nil {
		return
	}
	file_aspb_anime_season_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_aspb_rpc_create_anime_season_metas_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAnimeSeasonMetasRequest); i {
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
		file_aspb_rpc_create_anime_season_metas_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAnimeSeasonMetasResponse); i {
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
			RawDescriptor: file_aspb_rpc_create_anime_season_metas_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aspb_rpc_create_anime_season_metas_proto_goTypes,
		DependencyIndexes: file_aspb_rpc_create_anime_season_metas_proto_depIdxs,
		MessageInfos:      file_aspb_rpc_create_anime_season_metas_proto_msgTypes,
	}.Build()
	File_aspb_rpc_create_anime_season_metas_proto = out.File
	file_aspb_rpc_create_anime_season_metas_proto_rawDesc = nil
	file_aspb_rpc_create_anime_season_metas_proto_goTypes = nil
	file_aspb_rpc_create_anime_season_metas_proto_depIdxs = nil
}
