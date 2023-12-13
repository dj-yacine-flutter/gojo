// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: aspb/rpc_create_anime_season_image.proto

package aspb

import (
	shpb "github.com/dj-yacine-flutter/gojo/pb/shpb"
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

type CreateAnimeSerieSeasonImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeasonID int64                `protobuf:"varint,1,opt,name=seasonID,proto3" json:"seasonID,omitempty"`
	Posters  []*shpb.ImageRequest `protobuf:"bytes,2,rep,name=posters,proto3" json:"posters,omitempty"`
}

func (x *CreateAnimeSerieSeasonImageRequest) Reset() {
	*x = CreateAnimeSerieSeasonImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aspb_rpc_create_anime_season_image_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAnimeSerieSeasonImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAnimeSerieSeasonImageRequest) ProtoMessage() {}

func (x *CreateAnimeSerieSeasonImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aspb_rpc_create_anime_season_image_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAnimeSerieSeasonImageRequest.ProtoReflect.Descriptor instead.
func (*CreateAnimeSerieSeasonImageRequest) Descriptor() ([]byte, []int) {
	return file_aspb_rpc_create_anime_season_image_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAnimeSerieSeasonImageRequest) GetSeasonID() int64 {
	if x != nil {
		return x.SeasonID
	}
	return 0
}

func (x *CreateAnimeSerieSeasonImageRequest) GetPosters() []*shpb.ImageRequest {
	if x != nil {
		return x.Posters
	}
	return nil
}

type CreateAnimeSerieSeasonImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnimeSeason *AnimeSerieSeasonResponse `protobuf:"bytes,1,opt,name=animeSeason,proto3" json:"animeSeason,omitempty"`
	Posters     []*shpb.ImageResponse     `protobuf:"bytes,2,rep,name=posters,proto3" json:"posters,omitempty"`
}

func (x *CreateAnimeSerieSeasonImageResponse) Reset() {
	*x = CreateAnimeSerieSeasonImageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aspb_rpc_create_anime_season_image_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAnimeSerieSeasonImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAnimeSerieSeasonImageResponse) ProtoMessage() {}

func (x *CreateAnimeSerieSeasonImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aspb_rpc_create_anime_season_image_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAnimeSerieSeasonImageResponse.ProtoReflect.Descriptor instead.
func (*CreateAnimeSerieSeasonImageResponse) Descriptor() ([]byte, []int) {
	return file_aspb_rpc_create_anime_season_image_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAnimeSerieSeasonImageResponse) GetAnimeSeason() *AnimeSerieSeasonResponse {
	if x != nil {
		return x.AnimeSeason
	}
	return nil
}

func (x *CreateAnimeSerieSeasonImageResponse) GetPosters() []*shpb.ImageResponse {
	if x != nil {
		return x.Posters
	}
	return nil
}

var File_aspb_rpc_create_anime_season_image_proto protoreflect.FileDescriptor

var file_aspb_rpc_create_anime_season_image_proto_rawDesc = []byte{
	0x0a, 0x28, 0x61, 0x73, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x5f, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x73, 0x70, 0x62,
	0x1a, 0x16, 0x73, 0x68, 0x70, 0x62, 0x2f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x61, 0x73, 0x70, 0x62, 0x2f, 0x61,
	0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x5f, 0x73, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a, 0x22, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x53, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x2c, 0x0a, 0x07, 0x70, 0x6f, 0x73,
	0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x68, 0x70,
	0x62, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07,
	0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x73, 0x22, 0x96, 0x01, 0x0a, 0x23, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x53, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x40, 0x0a, 0x0b, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x73, 0x70, 0x62, 0x2e, 0x41, 0x6e, 0x69, 0x6d,
	0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0b, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x12, 0x2d, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x68, 0x70, 0x62, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x73,
	0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x6a, 0x2d, 0x79, 0x61, 0x63, 0x69, 0x6e, 0x65, 0x2d, 0x66, 0x6c, 0x75, 0x74, 0x74, 0x65, 0x72,
	0x2f, 0x67, 0x6f, 0x6a, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x61, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aspb_rpc_create_anime_season_image_proto_rawDescOnce sync.Once
	file_aspb_rpc_create_anime_season_image_proto_rawDescData = file_aspb_rpc_create_anime_season_image_proto_rawDesc
)

func file_aspb_rpc_create_anime_season_image_proto_rawDescGZIP() []byte {
	file_aspb_rpc_create_anime_season_image_proto_rawDescOnce.Do(func() {
		file_aspb_rpc_create_anime_season_image_proto_rawDescData = protoimpl.X.CompressGZIP(file_aspb_rpc_create_anime_season_image_proto_rawDescData)
	})
	return file_aspb_rpc_create_anime_season_image_proto_rawDescData
}

var file_aspb_rpc_create_anime_season_image_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_aspb_rpc_create_anime_season_image_proto_goTypes = []interface{}{
	(*CreateAnimeSerieSeasonImageRequest)(nil),  // 0: aspb.CreateAnimeSerieSeasonImageRequest
	(*CreateAnimeSerieSeasonImageResponse)(nil), // 1: aspb.CreateAnimeSerieSeasonImageResponse
	(*shpb.ImageRequest)(nil),                   // 2: shpb.ImageRequest
	(*AnimeSerieSeasonResponse)(nil),            // 3: aspb.AnimeSerieSeasonResponse
	(*shpb.ImageResponse)(nil),                  // 4: shpb.ImageResponse
}
var file_aspb_rpc_create_anime_season_image_proto_depIdxs = []int32{
	2, // 0: aspb.CreateAnimeSerieSeasonImageRequest.posters:type_name -> shpb.ImageRequest
	3, // 1: aspb.CreateAnimeSerieSeasonImageResponse.animeSeason:type_name -> aspb.AnimeSerieSeasonResponse
	4, // 2: aspb.CreateAnimeSerieSeasonImageResponse.posters:type_name -> shpb.ImageResponse
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_aspb_rpc_create_anime_season_image_proto_init() }
func file_aspb_rpc_create_anime_season_image_proto_init() {
	if File_aspb_rpc_create_anime_season_image_proto != nil {
		return
	}
	file_aspb_anime_serie_season_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_aspb_rpc_create_anime_season_image_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAnimeSerieSeasonImageRequest); i {
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
		file_aspb_rpc_create_anime_season_image_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAnimeSerieSeasonImageResponse); i {
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
			RawDescriptor: file_aspb_rpc_create_anime_season_image_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aspb_rpc_create_anime_season_image_proto_goTypes,
		DependencyIndexes: file_aspb_rpc_create_anime_season_image_proto_depIdxs,
		MessageInfos:      file_aspb_rpc_create_anime_season_image_proto_msgTypes,
	}.Build()
	File_aspb_rpc_create_anime_season_image_proto = out.File
	file_aspb_rpc_create_anime_season_image_proto_rawDesc = nil
	file_aspb_rpc_create_anime_season_image_proto_goTypes = nil
	file_aspb_rpc_create_anime_season_image_proto_depIdxs = nil
}