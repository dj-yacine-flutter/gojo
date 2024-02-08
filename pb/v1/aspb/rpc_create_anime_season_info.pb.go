// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: v1/aspb/rpc_create_anime_season_info.proto

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

type CreateAnimeSeasonInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeasonID  int64   `protobuf:"varint,1,opt,name=seasonID,proto3" json:"seasonID,omitempty"`
	GenreIDs  []int32 `protobuf:"varint,2,rep,packed,name=genreIDs,proto3" json:"genreIDs,omitempty"`
	StudioIDs []int32 `protobuf:"varint,3,rep,packed,name=studioIDs,proto3" json:"studioIDs,omitempty"`
}

func (x *CreateAnimeSeasonInfoRequest) Reset() {
	*x = CreateAnimeSeasonInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_aspb_rpc_create_anime_season_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAnimeSeasonInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAnimeSeasonInfoRequest) ProtoMessage() {}

func (x *CreateAnimeSeasonInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_aspb_rpc_create_anime_season_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAnimeSeasonInfoRequest.ProtoReflect.Descriptor instead.
func (*CreateAnimeSeasonInfoRequest) Descriptor() ([]byte, []int) {
	return file_v1_aspb_rpc_create_anime_season_info_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAnimeSeasonInfoRequest) GetSeasonID() int64 {
	if x != nil {
		return x.SeasonID
	}
	return 0
}

func (x *CreateAnimeSeasonInfoRequest) GetGenreIDs() []int32 {
	if x != nil {
		return x.GenreIDs
	}
	return nil
}

func (x *CreateAnimeSeasonInfoRequest) GetStudioIDs() []int32 {
	if x != nil {
		return x.StudioIDs
	}
	return nil
}

type CreateAnimeSeasonInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnimeSeason *AnimeSeasonResponse `protobuf:"bytes,1,opt,name=animeSeason,proto3" json:"animeSeason,omitempty"`
	Genres      []*nfpb.Genre        `protobuf:"bytes,2,rep,name=genres,proto3" json:"genres,omitempty"`
	Studios     []*nfpb.Studio       `protobuf:"bytes,3,rep,name=studios,proto3" json:"studios,omitempty"`
}

func (x *CreateAnimeSeasonInfoResponse) Reset() {
	*x = CreateAnimeSeasonInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_aspb_rpc_create_anime_season_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAnimeSeasonInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAnimeSeasonInfoResponse) ProtoMessage() {}

func (x *CreateAnimeSeasonInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_aspb_rpc_create_anime_season_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAnimeSeasonInfoResponse.ProtoReflect.Descriptor instead.
func (*CreateAnimeSeasonInfoResponse) Descriptor() ([]byte, []int) {
	return file_v1_aspb_rpc_create_anime_season_info_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAnimeSeasonInfoResponse) GetAnimeSeason() *AnimeSeasonResponse {
	if x != nil {
		return x.AnimeSeason
	}
	return nil
}

func (x *CreateAnimeSeasonInfoResponse) GetGenres() []*nfpb.Genre {
	if x != nil {
		return x.Genres
	}
	return nil
}

func (x *CreateAnimeSeasonInfoResponse) GetStudios() []*nfpb.Studio {
	if x != nil {
		return x.Studios
	}
	return nil
}

var File_v1_aspb_rpc_create_anime_season_info_proto protoreflect.FileDescriptor

var file_v1_aspb_rpc_create_anime_season_info_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x76, 0x31, 0x2f, 0x61, 0x73, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x76, 0x31,
	0x2e, 0x61, 0x73, 0x70, 0x62, 0x76, 0x31, 0x1a, 0x1e, 0x76, 0x31, 0x2f, 0x61, 0x73, 0x70, 0x62,
	0x2f, 0x6d, 0x73, 0x67, 0x5f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x31, 0x2f, 0x6e, 0x66, 0x70, 0x62,
	0x2f, 0x6d, 0x73, 0x67, 0x5f, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x18, 0x76, 0x31, 0x2f, 0x6e, 0x66, 0x70, 0x62, 0x2f, 0x6d, 0x73, 0x67, 0x5f, 0x73, 0x74,
	0x75, 0x64, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x74, 0x0a, 0x1c, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x49,
	0x44, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x08, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x49,
	0x44, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x49, 0x44, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x49, 0x44, 0x73,
	0x22, 0xb8, 0x01, 0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65,
	0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x73, 0x70,
	0x62, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0b, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x53, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x6e, 0x66, 0x70, 0x62, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x12, 0x2b,
	0x0a, 0x07, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x76, 0x31, 0x2e, 0x6e, 0x66, 0x70, 0x62, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x75, 0x64,
	0x69, 0x6f, 0x52, 0x07, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x73, 0x42, 0x35, 0x5a, 0x33, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6a, 0x2d, 0x79, 0x61, 0x63,
	0x69, 0x6e, 0x65, 0x2d, 0x66, 0x6c, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x6f, 0x6a, 0x6f,
	0x2f, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x73, 0x70, 0x62, 0x3b, 0x61, 0x73, 0x70, 0x62,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_aspb_rpc_create_anime_season_info_proto_rawDescOnce sync.Once
	file_v1_aspb_rpc_create_anime_season_info_proto_rawDescData = file_v1_aspb_rpc_create_anime_season_info_proto_rawDesc
)

func file_v1_aspb_rpc_create_anime_season_info_proto_rawDescGZIP() []byte {
	file_v1_aspb_rpc_create_anime_season_info_proto_rawDescOnce.Do(func() {
		file_v1_aspb_rpc_create_anime_season_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_aspb_rpc_create_anime_season_info_proto_rawDescData)
	})
	return file_v1_aspb_rpc_create_anime_season_info_proto_rawDescData
}

var file_v1_aspb_rpc_create_anime_season_info_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_aspb_rpc_create_anime_season_info_proto_goTypes = []interface{}{
	(*CreateAnimeSeasonInfoRequest)(nil),  // 0: v1.aspbv1.CreateAnimeSeasonInfoRequest
	(*CreateAnimeSeasonInfoResponse)(nil), // 1: v1.aspbv1.CreateAnimeSeasonInfoResponse
	(*AnimeSeasonResponse)(nil),           // 2: v1.aspbv1.AnimeSeasonResponse
	(*nfpb.Genre)(nil),                    // 3: v1.nfpbv1.Genre
	(*nfpb.Studio)(nil),                   // 4: v1.nfpbv1.Studio
}
var file_v1_aspb_rpc_create_anime_season_info_proto_depIdxs = []int32{
	2, // 0: v1.aspbv1.CreateAnimeSeasonInfoResponse.animeSeason:type_name -> v1.aspbv1.AnimeSeasonResponse
	3, // 1: v1.aspbv1.CreateAnimeSeasonInfoResponse.genres:type_name -> v1.nfpbv1.Genre
	4, // 2: v1.aspbv1.CreateAnimeSeasonInfoResponse.studios:type_name -> v1.nfpbv1.Studio
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_v1_aspb_rpc_create_anime_season_info_proto_init() }
func file_v1_aspb_rpc_create_anime_season_info_proto_init() {
	if File_v1_aspb_rpc_create_anime_season_info_proto != nil {
		return
	}
	file_v1_aspb_msg_anime_season_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_aspb_rpc_create_anime_season_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAnimeSeasonInfoRequest); i {
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
		file_v1_aspb_rpc_create_anime_season_info_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAnimeSeasonInfoResponse); i {
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
			RawDescriptor: file_v1_aspb_rpc_create_anime_season_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_aspb_rpc_create_anime_season_info_proto_goTypes,
		DependencyIndexes: file_v1_aspb_rpc_create_anime_season_info_proto_depIdxs,
		MessageInfos:      file_v1_aspb_rpc_create_anime_season_info_proto_msgTypes,
	}.Build()
	File_v1_aspb_rpc_create_anime_season_info_proto = out.File
	file_v1_aspb_rpc_create_anime_season_info_proto_rawDesc = nil
	file_v1_aspb_rpc_create_anime_season_info_proto_goTypes = nil
	file_v1_aspb_rpc_create_anime_season_info_proto_depIdxs = nil
}
