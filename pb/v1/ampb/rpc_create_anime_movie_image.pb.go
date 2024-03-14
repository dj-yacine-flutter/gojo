// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: v1/ampb/rpc_create_anime_movie_image.proto

package ampbv1

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

type CreateAnimeMovieImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnimeID     int64                  `protobuf:"varint,1,opt,name=animeID,proto3" json:"animeID,omitempty"`
	AnimeImages *apb.AnimeImageRequest `protobuf:"bytes,2,opt,name=animeImages,proto3" json:"animeImages,omitempty"`
}

func (x *CreateAnimeMovieImageRequest) Reset() {
	*x = CreateAnimeMovieImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_ampb_rpc_create_anime_movie_image_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAnimeMovieImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAnimeMovieImageRequest) ProtoMessage() {}

func (x *CreateAnimeMovieImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_ampb_rpc_create_anime_movie_image_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAnimeMovieImageRequest.ProtoReflect.Descriptor instead.
func (*CreateAnimeMovieImageRequest) Descriptor() ([]byte, []int) {
	return file_v1_ampb_rpc_create_anime_movie_image_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAnimeMovieImageRequest) GetAnimeID() int64 {
	if x != nil {
		return x.AnimeID
	}
	return 0
}

func (x *CreateAnimeMovieImageRequest) GetAnimeImages() *apb.AnimeImageRequest {
	if x != nil {
		return x.AnimeImages
	}
	return nil
}

type CreateAnimeMovieImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnimeMovie  *AnimeMovieResponse     `protobuf:"bytes,1,opt,name=animeMovie,proto3" json:"animeMovie,omitempty"`
	AnimeImages *apb.AnimeImageResponse `protobuf:"bytes,2,opt,name=animeImages,proto3" json:"animeImages,omitempty"`
}

func (x *CreateAnimeMovieImageResponse) Reset() {
	*x = CreateAnimeMovieImageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_ampb_rpc_create_anime_movie_image_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAnimeMovieImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAnimeMovieImageResponse) ProtoMessage() {}

func (x *CreateAnimeMovieImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_ampb_rpc_create_anime_movie_image_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAnimeMovieImageResponse.ProtoReflect.Descriptor instead.
func (*CreateAnimeMovieImageResponse) Descriptor() ([]byte, []int) {
	return file_v1_ampb_rpc_create_anime_movie_image_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAnimeMovieImageResponse) GetAnimeMovie() *AnimeMovieResponse {
	if x != nil {
		return x.AnimeMovie
	}
	return nil
}

func (x *CreateAnimeMovieImageResponse) GetAnimeImages() *apb.AnimeImageResponse {
	if x != nil {
		return x.AnimeImages
	}
	return nil
}

var File_v1_ampb_rpc_create_anime_movie_image_proto protoreflect.FileDescriptor

var file_v1_ampb_rpc_create_anime_movie_image_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x76, 0x31, 0x2f, 0x61, 0x6d, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x6f, 0x76, 0x69, 0x65,
	0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x76, 0x31,
	0x2e, 0x61, 0x6d, 0x70, 0x62, 0x76, 0x31, 0x1a, 0x18, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x62, 0x2f,
	0x61, 0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1d, 0x76, 0x31, 0x2f, 0x61, 0x6d, 0x70, 0x62, 0x2f, 0x6d, 0x73, 0x67, 0x5f, 0x61,
	0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x77, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x49, 0x44, 0x12, 0x3d, 0x0a, 0x0b, 0x61, 0x6e,
	0x69, 0x6d, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x70, 0x62, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x61, 0x6e,
	0x69, 0x6d, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x22, 0x9e, 0x01, 0x0a, 0x1d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0a, 0x61,
	0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x6d, 0x70, 0x62, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x69, 0x6d,
	0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a,
	0x61, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x61, 0x6e,
	0x69, 0x6d, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x70, 0x62, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0b, 0x61,
	0x6e, 0x69, 0x6d, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6a, 0x2d, 0x79, 0x61, 0x63, 0x69,
	0x6e, 0x65, 0x2d, 0x66, 0x6c, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x6f, 0x6a, 0x6f, 0x2f,
	0x70, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6d, 0x70, 0x62, 0x3b, 0x61, 0x6d, 0x70, 0x62, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_ampb_rpc_create_anime_movie_image_proto_rawDescOnce sync.Once
	file_v1_ampb_rpc_create_anime_movie_image_proto_rawDescData = file_v1_ampb_rpc_create_anime_movie_image_proto_rawDesc
)

func file_v1_ampb_rpc_create_anime_movie_image_proto_rawDescGZIP() []byte {
	file_v1_ampb_rpc_create_anime_movie_image_proto_rawDescOnce.Do(func() {
		file_v1_ampb_rpc_create_anime_movie_image_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_ampb_rpc_create_anime_movie_image_proto_rawDescData)
	})
	return file_v1_ampb_rpc_create_anime_movie_image_proto_rawDescData
}

var file_v1_ampb_rpc_create_anime_movie_image_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_ampb_rpc_create_anime_movie_image_proto_goTypes = []interface{}{
	(*CreateAnimeMovieImageRequest)(nil),  // 0: v1.ampbv1.CreateAnimeMovieImageRequest
	(*CreateAnimeMovieImageResponse)(nil), // 1: v1.ampbv1.CreateAnimeMovieImageResponse
	(*apb.AnimeImageRequest)(nil),         // 2: v1.apbv1.AnimeImageRequest
	(*AnimeMovieResponse)(nil),            // 3: v1.ampbv1.AnimeMovieResponse
	(*apb.AnimeImageResponse)(nil),        // 4: v1.apbv1.AnimeImageResponse
}
var file_v1_ampb_rpc_create_anime_movie_image_proto_depIdxs = []int32{
	2, // 0: v1.ampbv1.CreateAnimeMovieImageRequest.animeImages:type_name -> v1.apbv1.AnimeImageRequest
	3, // 1: v1.ampbv1.CreateAnimeMovieImageResponse.animeMovie:type_name -> v1.ampbv1.AnimeMovieResponse
	4, // 2: v1.ampbv1.CreateAnimeMovieImageResponse.animeImages:type_name -> v1.apbv1.AnimeImageResponse
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_v1_ampb_rpc_create_anime_movie_image_proto_init() }
func file_v1_ampb_rpc_create_anime_movie_image_proto_init() {
	if File_v1_ampb_rpc_create_anime_movie_image_proto != nil {
		return
	}
	file_v1_ampb_msg_anime_movie_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_ampb_rpc_create_anime_movie_image_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAnimeMovieImageRequest); i {
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
		file_v1_ampb_rpc_create_anime_movie_image_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAnimeMovieImageResponse); i {
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
			RawDescriptor: file_v1_ampb_rpc_create_anime_movie_image_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_ampb_rpc_create_anime_movie_image_proto_goTypes,
		DependencyIndexes: file_v1_ampb_rpc_create_anime_movie_image_proto_depIdxs,
		MessageInfos:      file_v1_ampb_rpc_create_anime_movie_image_proto_msgTypes,
	}.Build()
	File_v1_ampb_rpc_create_anime_movie_image_proto = out.File
	file_v1_ampb_rpc_create_anime_movie_image_proto_rawDesc = nil
	file_v1_ampb_rpc_create_anime_movie_image_proto_goTypes = nil
	file_v1_ampb_rpc_create_anime_movie_image_proto_depIdxs = nil
}
