// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: v1/ampb/rpc_get_anime_movie_images.proto

package ampbv1

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

type GetAnimeMovieImagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnimeID int64 `protobuf:"varint,1,opt,name=animeID,proto3" json:"animeID,omitempty"`
}

func (x *GetAnimeMovieImagesRequest) Reset() {
	*x = GetAnimeMovieImagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_ampb_rpc_get_anime_movie_images_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAnimeMovieImagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAnimeMovieImagesRequest) ProtoMessage() {}

func (x *GetAnimeMovieImagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_ampb_rpc_get_anime_movie_images_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAnimeMovieImagesRequest.ProtoReflect.Descriptor instead.
func (*GetAnimeMovieImagesRequest) Descriptor() ([]byte, []int) {
	return file_v1_ampb_rpc_get_anime_movie_images_proto_rawDescGZIP(), []int{0}
}

func (x *GetAnimeMovieImagesRequest) GetAnimeID() int64 {
	if x != nil {
		return x.AnimeID
	}
	return 0
}

type GetAnimeMovieImagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnimeImages *shpb.AnimeImageResponse `protobuf:"bytes,1,opt,name=animeImages,proto3" json:"animeImages,omitempty"`
}

func (x *GetAnimeMovieImagesResponse) Reset() {
	*x = GetAnimeMovieImagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_ampb_rpc_get_anime_movie_images_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAnimeMovieImagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAnimeMovieImagesResponse) ProtoMessage() {}

func (x *GetAnimeMovieImagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_ampb_rpc_get_anime_movie_images_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAnimeMovieImagesResponse.ProtoReflect.Descriptor instead.
func (*GetAnimeMovieImagesResponse) Descriptor() ([]byte, []int) {
	return file_v1_ampb_rpc_get_anime_movie_images_proto_rawDescGZIP(), []int{1}
}

func (x *GetAnimeMovieImagesResponse) GetAnimeImages() *shpb.AnimeImageResponse {
	if x != nil {
		return x.AnimeImages
	}
	return nil
}

var File_v1_ampb_rpc_get_anime_movie_images_proto protoreflect.FileDescriptor

var file_v1_ampb_rpc_get_anime_movie_images_proto_rawDesc = []byte{
	0x0a, 0x28, 0x76, 0x31, 0x2f, 0x61, 0x6d, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65,
	0x74, 0x5f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x5f, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x76, 0x31, 0x2e, 0x61,
	0x6d, 0x70, 0x62, 0x76, 0x31, 0x1a, 0x16, 0x73, 0x68, 0x70, 0x62, 0x2f, 0x61, 0x6e, 0x69, 0x6d,
	0x65, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a,
	0x1a, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x6e, 0x69, 0x6d, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x6e,
	0x69, 0x6d, 0x65, 0x49, 0x44, 0x22, 0x59, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x69, 0x6d,
	0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0b, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x68, 0x70, 0x62,
	0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x0b, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73,
	0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x6a, 0x2d, 0x79, 0x61, 0x63, 0x69, 0x6e, 0x65, 0x2d, 0x66, 0x6c, 0x75, 0x74, 0x74, 0x65, 0x72,
	0x2f, 0x67, 0x6f, 0x6a, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6d, 0x70, 0x62,
	0x3b, 0x61, 0x6d, 0x70, 0x62, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_ampb_rpc_get_anime_movie_images_proto_rawDescOnce sync.Once
	file_v1_ampb_rpc_get_anime_movie_images_proto_rawDescData = file_v1_ampb_rpc_get_anime_movie_images_proto_rawDesc
)

func file_v1_ampb_rpc_get_anime_movie_images_proto_rawDescGZIP() []byte {
	file_v1_ampb_rpc_get_anime_movie_images_proto_rawDescOnce.Do(func() {
		file_v1_ampb_rpc_get_anime_movie_images_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_ampb_rpc_get_anime_movie_images_proto_rawDescData)
	})
	return file_v1_ampb_rpc_get_anime_movie_images_proto_rawDescData
}

var file_v1_ampb_rpc_get_anime_movie_images_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_ampb_rpc_get_anime_movie_images_proto_goTypes = []interface{}{
	(*GetAnimeMovieImagesRequest)(nil),  // 0: v1.ampbv1.GetAnimeMovieImagesRequest
	(*GetAnimeMovieImagesResponse)(nil), // 1: v1.ampbv1.GetAnimeMovieImagesResponse
	(*shpb.AnimeImageResponse)(nil),     // 2: shpb.AnimeImageResponse
}
var file_v1_ampb_rpc_get_anime_movie_images_proto_depIdxs = []int32{
	2, // 0: v1.ampbv1.GetAnimeMovieImagesResponse.animeImages:type_name -> shpb.AnimeImageResponse
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v1_ampb_rpc_get_anime_movie_images_proto_init() }
func file_v1_ampb_rpc_get_anime_movie_images_proto_init() {
	if File_v1_ampb_rpc_get_anime_movie_images_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_ampb_rpc_get_anime_movie_images_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAnimeMovieImagesRequest); i {
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
		file_v1_ampb_rpc_get_anime_movie_images_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAnimeMovieImagesResponse); i {
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
			RawDescriptor: file_v1_ampb_rpc_get_anime_movie_images_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_ampb_rpc_get_anime_movie_images_proto_goTypes,
		DependencyIndexes: file_v1_ampb_rpc_get_anime_movie_images_proto_depIdxs,
		MessageInfos:      file_v1_ampb_rpc_get_anime_movie_images_proto_msgTypes,
	}.Build()
	File_v1_ampb_rpc_get_anime_movie_images_proto = out.File
	file_v1_ampb_rpc_get_anime_movie_images_proto_rawDesc = nil
	file_v1_ampb_rpc_get_anime_movie_images_proto_goTypes = nil
	file_v1_ampb_rpc_get_anime_movie_images_proto_depIdxs = nil
}
