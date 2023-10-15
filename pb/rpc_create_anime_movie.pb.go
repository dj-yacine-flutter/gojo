// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: rpc_create_anime_movie.proto

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

type CreateAnimeMovieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnimeMovie   *AnimeMovieRequest   `protobuf:"bytes,1,opt,name=animeMovie,proto3" json:"animeMovie,omitempty"`
	AnimeGenres  *AnimeGenresRequest  `protobuf:"bytes,2,opt,name=animeGenres,proto3,oneof" json:"animeGenres,omitempty"`
	AnimeStudios *AnimeStudiosRequest `protobuf:"bytes,3,opt,name=animeStudios,proto3,oneof" json:"animeStudios,omitempty"`
	AnimeMetas   []*AnimeMetaRequest  `protobuf:"bytes,4,rep,name=animeMetas,proto3" json:"animeMetas,omitempty"`
}

func (x *CreateAnimeMovieRequest) Reset() {
	*x = CreateAnimeMovieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_create_anime_movie_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAnimeMovieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAnimeMovieRequest) ProtoMessage() {}

func (x *CreateAnimeMovieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_create_anime_movie_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAnimeMovieRequest.ProtoReflect.Descriptor instead.
func (*CreateAnimeMovieRequest) Descriptor() ([]byte, []int) {
	return file_rpc_create_anime_movie_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAnimeMovieRequest) GetAnimeMovie() *AnimeMovieRequest {
	if x != nil {
		return x.AnimeMovie
	}
	return nil
}

func (x *CreateAnimeMovieRequest) GetAnimeGenres() *AnimeGenresRequest {
	if x != nil {
		return x.AnimeGenres
	}
	return nil
}

func (x *CreateAnimeMovieRequest) GetAnimeStudios() *AnimeStudiosRequest {
	if x != nil {
		return x.AnimeStudios
	}
	return nil
}

func (x *CreateAnimeMovieRequest) GetAnimeMetas() []*AnimeMetaRequest {
	if x != nil {
		return x.AnimeMetas
	}
	return nil
}

type CreateAnimeMovieResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnimeMovie *AnimeMovieResponse `protobuf:"bytes,1,opt,name=animeMovie,proto3" json:"animeMovie,omitempty"`
}

func (x *CreateAnimeMovieResponse) Reset() {
	*x = CreateAnimeMovieResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_create_anime_movie_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAnimeMovieResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAnimeMovieResponse) ProtoMessage() {}

func (x *CreateAnimeMovieResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_create_anime_movie_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAnimeMovieResponse.ProtoReflect.Descriptor instead.
func (*CreateAnimeMovieResponse) Descriptor() ([]byte, []int) {
	return file_rpc_create_anime_movie_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAnimeMovieResponse) GetAnimeMovie() *AnimeMovieResponse {
	if x != nil {
		return x.AnimeMovie
	}
	return nil
}

var File_rpc_create_anime_movie_proto protoreflect.FileDescriptor

var file_rpc_create_anime_movie_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x6e, 0x69,
	0x6d, 0x65, 0x5f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x1a, 0x11, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0c, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0a, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x02, 0x0a,
	0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x0a, 0x61, 0x6e, 0x69, 0x6d,
	0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70,
	0x62, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x0a, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12,
	0x3d, 0x0a, 0x0b, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x47,
	0x65, 0x6e, 0x72, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0b,
	0x61, 0x6e, 0x69, 0x6d, 0x65, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x88, 0x01, 0x01, 0x12, 0x40,
	0x0a, 0x0c, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x53,
	0x74, 0x75, 0x64, 0x69, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x01, 0x52,
	0x0c, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x73, 0x88, 0x01, 0x01,
	0x12, 0x34, 0x0a, 0x0a, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d,
	0x65, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0a, 0x61, 0x6e, 0x69, 0x6d,
	0x65, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x61, 0x6e, 0x69, 0x6d, 0x65,
	0x47, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x61, 0x6e, 0x69, 0x6d, 0x65,
	0x53, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x73, 0x22, 0x52, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0a, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x6e, 0x69,
	0x6d, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x0a, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x42, 0x26, 0x5a, 0x24, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6a, 0x2d, 0x79, 0x61, 0x63,
	0x69, 0x6e, 0x65, 0x2d, 0x66, 0x6c, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x6f, 0x6a, 0x6f,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_create_anime_movie_proto_rawDescOnce sync.Once
	file_rpc_create_anime_movie_proto_rawDescData = file_rpc_create_anime_movie_proto_rawDesc
)

func file_rpc_create_anime_movie_proto_rawDescGZIP() []byte {
	file_rpc_create_anime_movie_proto_rawDescOnce.Do(func() {
		file_rpc_create_anime_movie_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_create_anime_movie_proto_rawDescData)
	})
	return file_rpc_create_anime_movie_proto_rawDescData
}

var file_rpc_create_anime_movie_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_create_anime_movie_proto_goTypes = []interface{}{
	(*CreateAnimeMovieRequest)(nil),  // 0: pb.CreateAnimeMovieRequest
	(*CreateAnimeMovieResponse)(nil), // 1: pb.CreateAnimeMovieResponse
	(*AnimeMovieRequest)(nil),        // 2: pb.AnimeMovieRequest
	(*AnimeGenresRequest)(nil),       // 3: pb.AnimeGenresRequest
	(*AnimeStudiosRequest)(nil),      // 4: pb.AnimeStudiosRequest
	(*AnimeMetaRequest)(nil),         // 5: pb.AnimeMetaRequest
	(*AnimeMovieResponse)(nil),       // 6: pb.AnimeMovieResponse
}
var file_rpc_create_anime_movie_proto_depIdxs = []int32{
	2, // 0: pb.CreateAnimeMovieRequest.animeMovie:type_name -> pb.AnimeMovieRequest
	3, // 1: pb.CreateAnimeMovieRequest.animeGenres:type_name -> pb.AnimeGenresRequest
	4, // 2: pb.CreateAnimeMovieRequest.animeStudios:type_name -> pb.AnimeStudiosRequest
	5, // 3: pb.CreateAnimeMovieRequest.animeMetas:type_name -> pb.AnimeMetaRequest
	6, // 4: pb.CreateAnimeMovieResponse.animeMovie:type_name -> pb.AnimeMovieResponse
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_rpc_create_anime_movie_proto_init() }
func file_rpc_create_anime_movie_proto_init() {
	if File_rpc_create_anime_movie_proto != nil {
		return
	}
	file_anime_movie_proto_init()
	file_genre_proto_init()
	file_studio_proto_init()
	file_meta_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_create_anime_movie_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAnimeMovieRequest); i {
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
		file_rpc_create_anime_movie_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAnimeMovieResponse); i {
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
	file_rpc_create_anime_movie_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_create_anime_movie_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_create_anime_movie_proto_goTypes,
		DependencyIndexes: file_rpc_create_anime_movie_proto_depIdxs,
		MessageInfos:      file_rpc_create_anime_movie_proto_msgTypes,
	}.Build()
	File_rpc_create_anime_movie_proto = out.File
	file_rpc_create_anime_movie_proto_rawDesc = nil
	file_rpc_create_anime_movie_proto_goTypes = nil
	file_rpc_create_anime_movie_proto_depIdxs = nil
}
