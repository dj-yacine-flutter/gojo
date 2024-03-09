// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: v1/aspb/rpc_get_anime_season_characters.proto

package aspbv1

import (
	apb "github.com/dj-yacine-flutter/gojo/pb/v1/apb"
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

type GetAnimeSeasonCharactersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeasonID int64 `protobuf:"varint,1,opt,name=seasonID,proto3" json:"seasonID,omitempty"`
}

func (x *GetAnimeSeasonCharactersRequest) Reset() {
	*x = GetAnimeSeasonCharactersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_aspb_rpc_get_anime_season_characters_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAnimeSeasonCharactersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAnimeSeasonCharactersRequest) ProtoMessage() {}

func (x *GetAnimeSeasonCharactersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_aspb_rpc_get_anime_season_characters_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAnimeSeasonCharactersRequest.ProtoReflect.Descriptor instead.
func (*GetAnimeSeasonCharactersRequest) Descriptor() ([]byte, []int) {
	return file_v1_aspb_rpc_get_anime_season_characters_proto_rawDescGZIP(), []int{0}
}

func (x *GetAnimeSeasonCharactersRequest) GetSeasonID() int64 {
	if x != nil {
		return x.SeasonID
	}
	return 0
}

type GetAnimeSeasonCharactersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeasonCharacters []*apb.AnimeCharacter `protobuf:"bytes,1,rep,name=seasonCharacters,proto3" json:"seasonCharacters,omitempty"`
}

func (x *GetAnimeSeasonCharactersResponse) Reset() {
	*x = GetAnimeSeasonCharactersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_aspb_rpc_get_anime_season_characters_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAnimeSeasonCharactersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAnimeSeasonCharactersResponse) ProtoMessage() {}

func (x *GetAnimeSeasonCharactersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_aspb_rpc_get_anime_season_characters_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAnimeSeasonCharactersResponse.ProtoReflect.Descriptor instead.
func (*GetAnimeSeasonCharactersResponse) Descriptor() ([]byte, []int) {
	return file_v1_aspb_rpc_get_anime_season_characters_proto_rawDescGZIP(), []int{1}
}

func (x *GetAnimeSeasonCharactersResponse) GetSeasonCharacters() []*apb.AnimeCharacter {
	if x != nil {
		return x.SeasonCharacters
	}
	return nil
}

var File_v1_aspb_rpc_get_anime_season_characters_proto protoreflect.FileDescriptor

var file_v1_aspb_rpc_get_anime_season_characters_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x76, 0x31, 0x2f, 0x61, 0x73, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65,
	0x74, 0x5f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x5f, 0x63,
	0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x76, 0x31, 0x2e, 0x61, 0x73, 0x70, 0x62, 0x76, 0x31, 0x1a, 0x1c, 0x76, 0x31, 0x2f, 0x61,
	0x70, 0x62, 0x2f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3d, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x41,
	0x6e, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63,
	0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x44, 0x22, 0x68, 0x0a, 0x20, 0x47, 0x65, 0x74, 0x41, 0x6e,
	0x69, 0x6d, 0x65, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x10, 0x73,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x70, 0x62, 0x76, 0x31,
	0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52,
	0x10, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72,
	0x73, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x64, 0x6a, 0x2d, 0x79, 0x61, 0x63, 0x69, 0x6e, 0x65, 0x2d, 0x66, 0x6c, 0x75, 0x74, 0x74, 0x65,
	0x72, 0x2f, 0x67, 0x6f, 0x6a, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x73, 0x70,
	0x62, 0x3b, 0x61, 0x73, 0x70, 0x62, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_aspb_rpc_get_anime_season_characters_proto_rawDescOnce sync.Once
	file_v1_aspb_rpc_get_anime_season_characters_proto_rawDescData = file_v1_aspb_rpc_get_anime_season_characters_proto_rawDesc
)

func file_v1_aspb_rpc_get_anime_season_characters_proto_rawDescGZIP() []byte {
	file_v1_aspb_rpc_get_anime_season_characters_proto_rawDescOnce.Do(func() {
		file_v1_aspb_rpc_get_anime_season_characters_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_aspb_rpc_get_anime_season_characters_proto_rawDescData)
	})
	return file_v1_aspb_rpc_get_anime_season_characters_proto_rawDescData
}

var file_v1_aspb_rpc_get_anime_season_characters_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_aspb_rpc_get_anime_season_characters_proto_goTypes = []interface{}{
	(*GetAnimeSeasonCharactersRequest)(nil),  // 0: v1.aspbv1.GetAnimeSeasonCharactersRequest
	(*GetAnimeSeasonCharactersResponse)(nil), // 1: v1.aspbv1.GetAnimeSeasonCharactersResponse
	(*apb.AnimeCharacter)(nil),               // 2: v1.apbv1.AnimeCharacter
}
var file_v1_aspb_rpc_get_anime_season_characters_proto_depIdxs = []int32{
	2, // 0: v1.aspbv1.GetAnimeSeasonCharactersResponse.seasonCharacters:type_name -> v1.apbv1.AnimeCharacter
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v1_aspb_rpc_get_anime_season_characters_proto_init() }
func file_v1_aspb_rpc_get_anime_season_characters_proto_init() {
	if File_v1_aspb_rpc_get_anime_season_characters_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_aspb_rpc_get_anime_season_characters_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAnimeSeasonCharactersRequest); i {
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
		file_v1_aspb_rpc_get_anime_season_characters_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAnimeSeasonCharactersResponse); i {
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
			RawDescriptor: file_v1_aspb_rpc_get_anime_season_characters_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_aspb_rpc_get_anime_season_characters_proto_goTypes,
		DependencyIndexes: file_v1_aspb_rpc_get_anime_season_characters_proto_depIdxs,
		MessageInfos:      file_v1_aspb_rpc_get_anime_season_characters_proto_msgTypes,
	}.Build()
	File_v1_aspb_rpc_get_anime_season_characters_proto = out.File
	file_v1_aspb_rpc_get_anime_season_characters_proto_rawDesc = nil
	file_v1_aspb_rpc_get_anime_season_characters_proto_goTypes = nil
	file_v1_aspb_rpc_get_anime_season_characters_proto_depIdxs = nil
}
