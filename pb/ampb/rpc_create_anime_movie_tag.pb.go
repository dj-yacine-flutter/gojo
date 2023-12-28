// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: ampb/rpc_create_anime_movie_tag.proto

package ampb

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

type CreateAnimeMovieTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnimeID   int64    `protobuf:"varint,1,opt,name=animeID,proto3" json:"animeID,omitempty"`
	AnimeTags []string `protobuf:"bytes,2,rep,name=animeTags,proto3" json:"animeTags,omitempty"`
}

func (x *CreateAnimeMovieTagRequest) Reset() {
	*x = CreateAnimeMovieTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ampb_rpc_create_anime_movie_tag_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAnimeMovieTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAnimeMovieTagRequest) ProtoMessage() {}

func (x *CreateAnimeMovieTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ampb_rpc_create_anime_movie_tag_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAnimeMovieTagRequest.ProtoReflect.Descriptor instead.
func (*CreateAnimeMovieTagRequest) Descriptor() ([]byte, []int) {
	return file_ampb_rpc_create_anime_movie_tag_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAnimeMovieTagRequest) GetAnimeID() int64 {
	if x != nil {
		return x.AnimeID
	}
	return 0
}

func (x *CreateAnimeMovieTagRequest) GetAnimeTags() []string {
	if x != nil {
		return x.AnimeTags
	}
	return nil
}

type CreateAnimeMovieTagResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnimeMovie *AnimeMovieResponse `protobuf:"bytes,1,opt,name=animeMovie,proto3" json:"animeMovie,omitempty"`
	AnimeTags  []*AnimeMovieTag    `protobuf:"bytes,2,rep,name=animeTags,proto3" json:"animeTags,omitempty"`
}

func (x *CreateAnimeMovieTagResponse) Reset() {
	*x = CreateAnimeMovieTagResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ampb_rpc_create_anime_movie_tag_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAnimeMovieTagResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAnimeMovieTagResponse) ProtoMessage() {}

func (x *CreateAnimeMovieTagResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ampb_rpc_create_anime_movie_tag_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAnimeMovieTagResponse.ProtoReflect.Descriptor instead.
func (*CreateAnimeMovieTagResponse) Descriptor() ([]byte, []int) {
	return file_ampb_rpc_create_anime_movie_tag_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAnimeMovieTagResponse) GetAnimeMovie() *AnimeMovieResponse {
	if x != nil {
		return x.AnimeMovie
	}
	return nil
}

func (x *CreateAnimeMovieTagResponse) GetAnimeTags() []*AnimeMovieTag {
	if x != nil {
		return x.AnimeTags
	}
	return nil
}

var File_ampb_rpc_create_anime_movie_tag_proto protoreflect.FileDescriptor

var file_ampb_rpc_create_anime_movie_tag_proto_rawDesc = []byte{
	0x0a, 0x25, 0x61, 0x6d, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x5f, 0x74, 0x61,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x6d, 0x70, 0x62, 0x1a, 0x16, 0x61,
	0x6d, 0x70, 0x62, 0x2f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x61, 0x6d, 0x70, 0x62, 0x2f, 0x74, 0x61, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x49, 0x44, 0x12, 0x1c, 0x0a,
	0x09, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x54, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x09, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x54, 0x61, 0x67, 0x73, 0x22, 0x8a, 0x01, 0x0a, 0x1b,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0a, 0x61,
	0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x61, 0x6d, 0x70, 0x62, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x61, 0x6e, 0x69, 0x6d, 0x65,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x31, 0x0a, 0x09, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x54, 0x61,
	0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x6d, 0x70, 0x62, 0x2e,
	0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x54, 0x61, 0x67, 0x52, 0x09, 0x61,
	0x6e, 0x69, 0x6d, 0x65, 0x54, 0x61, 0x67, 0x73, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6a, 0x2d, 0x79, 0x61, 0x63, 0x69, 0x6e, 0x65,
	0x2d, 0x66, 0x6c, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x6f, 0x6a, 0x6f, 0x2f, 0x70, 0x62,
	0x2f, 0x61, 0x6d, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ampb_rpc_create_anime_movie_tag_proto_rawDescOnce sync.Once
	file_ampb_rpc_create_anime_movie_tag_proto_rawDescData = file_ampb_rpc_create_anime_movie_tag_proto_rawDesc
)

func file_ampb_rpc_create_anime_movie_tag_proto_rawDescGZIP() []byte {
	file_ampb_rpc_create_anime_movie_tag_proto_rawDescOnce.Do(func() {
		file_ampb_rpc_create_anime_movie_tag_proto_rawDescData = protoimpl.X.CompressGZIP(file_ampb_rpc_create_anime_movie_tag_proto_rawDescData)
	})
	return file_ampb_rpc_create_anime_movie_tag_proto_rawDescData
}

var file_ampb_rpc_create_anime_movie_tag_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ampb_rpc_create_anime_movie_tag_proto_goTypes = []interface{}{
	(*CreateAnimeMovieTagRequest)(nil),  // 0: ampb.CreateAnimeMovieTagRequest
	(*CreateAnimeMovieTagResponse)(nil), // 1: ampb.CreateAnimeMovieTagResponse
	(*AnimeMovieResponse)(nil),          // 2: ampb.AnimeMovieResponse
	(*AnimeMovieTag)(nil),               // 3: ampb.AnimeMovieTag
}
var file_ampb_rpc_create_anime_movie_tag_proto_depIdxs = []int32{
	2, // 0: ampb.CreateAnimeMovieTagResponse.animeMovie:type_name -> ampb.AnimeMovieResponse
	3, // 1: ampb.CreateAnimeMovieTagResponse.animeTags:type_name -> ampb.AnimeMovieTag
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ampb_rpc_create_anime_movie_tag_proto_init() }
func file_ampb_rpc_create_anime_movie_tag_proto_init() {
	if File_ampb_rpc_create_anime_movie_tag_proto != nil {
		return
	}
	file_ampb_anime_movie_proto_init()
	file_ampb_tag_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ampb_rpc_create_anime_movie_tag_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAnimeMovieTagRequest); i {
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
		file_ampb_rpc_create_anime_movie_tag_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAnimeMovieTagResponse); i {
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
			RawDescriptor: file_ampb_rpc_create_anime_movie_tag_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ampb_rpc_create_anime_movie_tag_proto_goTypes,
		DependencyIndexes: file_ampb_rpc_create_anime_movie_tag_proto_depIdxs,
		MessageInfos:      file_ampb_rpc_create_anime_movie_tag_proto_msgTypes,
	}.Build()
	File_ampb_rpc_create_anime_movie_tag_proto = out.File
	file_ampb_rpc_create_anime_movie_tag_proto_rawDesc = nil
	file_ampb_rpc_create_anime_movie_tag_proto_goTypes = nil
	file_ampb_rpc_create_anime_movie_tag_proto_depIdxs = nil
}
