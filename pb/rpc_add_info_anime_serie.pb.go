// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: rpc_add_info_anime_serie.proto

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

type AddInfoAnimeSerieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnimeID int64                `protobuf:"varint,1,opt,name=animeID,proto3" json:"animeID,omitempty"`
	Genres  *AnimeGenresRequest  `protobuf:"bytes,2,opt,name=genres,proto3" json:"genres,omitempty"`
	Studios *AnimeStudiosRequest `protobuf:"bytes,3,opt,name=studios,proto3" json:"studios,omitempty"`
}

func (x *AddInfoAnimeSerieRequest) Reset() {
	*x = AddInfoAnimeSerieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_add_info_anime_serie_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddInfoAnimeSerieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddInfoAnimeSerieRequest) ProtoMessage() {}

func (x *AddInfoAnimeSerieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_add_info_anime_serie_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddInfoAnimeSerieRequest.ProtoReflect.Descriptor instead.
func (*AddInfoAnimeSerieRequest) Descriptor() ([]byte, []int) {
	return file_rpc_add_info_anime_serie_proto_rawDescGZIP(), []int{0}
}

func (x *AddInfoAnimeSerieRequest) GetAnimeID() int64 {
	if x != nil {
		return x.AnimeID
	}
	return 0
}

func (x *AddInfoAnimeSerieRequest) GetGenres() *AnimeGenresRequest {
	if x != nil {
		return x.Genres
	}
	return nil
}

func (x *AddInfoAnimeSerieRequest) GetStudios() *AnimeStudiosRequest {
	if x != nil {
		return x.Studios
	}
	return nil
}

type AddInfoAnimeSerieResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnimeSerie *AnimeSerieResponse   `protobuf:"bytes,1,opt,name=animeSerie,proto3" json:"animeSerie,omitempty"`
	Genres     *AnimeGenresResponse  `protobuf:"bytes,2,opt,name=genres,proto3" json:"genres,omitempty"`
	Studios    *AnimeStudiosResponse `protobuf:"bytes,3,opt,name=studios,proto3" json:"studios,omitempty"`
}

func (x *AddInfoAnimeSerieResponse) Reset() {
	*x = AddInfoAnimeSerieResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_add_info_anime_serie_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddInfoAnimeSerieResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddInfoAnimeSerieResponse) ProtoMessage() {}

func (x *AddInfoAnimeSerieResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_add_info_anime_serie_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddInfoAnimeSerieResponse.ProtoReflect.Descriptor instead.
func (*AddInfoAnimeSerieResponse) Descriptor() ([]byte, []int) {
	return file_rpc_add_info_anime_serie_proto_rawDescGZIP(), []int{1}
}

func (x *AddInfoAnimeSerieResponse) GetAnimeSerie() *AnimeSerieResponse {
	if x != nil {
		return x.AnimeSerie
	}
	return nil
}

func (x *AddInfoAnimeSerieResponse) GetGenres() *AnimeGenresResponse {
	if x != nil {
		return x.Genres
	}
	return nil
}

func (x *AddInfoAnimeSerieResponse) GetStudios() *AnimeStudiosResponse {
	if x != nil {
		return x.Studios
	}
	return nil
}

var File_rpc_add_info_anime_serie_proto protoreflect.FileDescriptor

var file_rpc_add_info_anime_serie_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x61,
	0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x1a, 0x11, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x69,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x97, 0x01, 0x0a, 0x18, 0x41, 0x64, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x41, 0x6e,
	0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x49, 0x44, 0x12, 0x2e, 0x0a, 0x06, 0x67, 0x65, 0x6e,
	0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x41,
	0x6e, 0x69, 0x6d, 0x65, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x12, 0x31, 0x0a, 0x07, 0x73, 0x74, 0x75,
	0x64, 0x69, 0x6f, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x62, 0x2e,
	0x41, 0x6e, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x07, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x73, 0x22, 0xb8, 0x01, 0x0a,
	0x19, 0x41, 0x64, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72,
	0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0a, 0x61, 0x6e,
	0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x70, 0x62, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72,
	0x69, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x47, 0x65, 0x6e,
	0x72, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06, 0x67, 0x65, 0x6e,
	0x72, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x07, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x53,
	0x74, 0x75, 0x64, 0x69, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x07,
	0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x73, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6a, 0x2d, 0x79, 0x61, 0x63, 0x69, 0x6e, 0x65, 0x2d,
	0x66, 0x6c, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x6f, 0x6a, 0x6f, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_add_info_anime_serie_proto_rawDescOnce sync.Once
	file_rpc_add_info_anime_serie_proto_rawDescData = file_rpc_add_info_anime_serie_proto_rawDesc
)

func file_rpc_add_info_anime_serie_proto_rawDescGZIP() []byte {
	file_rpc_add_info_anime_serie_proto_rawDescOnce.Do(func() {
		file_rpc_add_info_anime_serie_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_add_info_anime_serie_proto_rawDescData)
	})
	return file_rpc_add_info_anime_serie_proto_rawDescData
}

var file_rpc_add_info_anime_serie_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_add_info_anime_serie_proto_goTypes = []interface{}{
	(*AddInfoAnimeSerieRequest)(nil),  // 0: pb.AddInfoAnimeSerieRequest
	(*AddInfoAnimeSerieResponse)(nil), // 1: pb.AddInfoAnimeSerieResponse
	(*AnimeGenresRequest)(nil),        // 2: pb.AnimeGenresRequest
	(*AnimeStudiosRequest)(nil),       // 3: pb.AnimeStudiosRequest
	(*AnimeSerieResponse)(nil),        // 4: pb.AnimeSerieResponse
	(*AnimeGenresResponse)(nil),       // 5: pb.AnimeGenresResponse
	(*AnimeStudiosResponse)(nil),      // 6: pb.AnimeStudiosResponse
}
var file_rpc_add_info_anime_serie_proto_depIdxs = []int32{
	2, // 0: pb.AddInfoAnimeSerieRequest.genres:type_name -> pb.AnimeGenresRequest
	3, // 1: pb.AddInfoAnimeSerieRequest.studios:type_name -> pb.AnimeStudiosRequest
	4, // 2: pb.AddInfoAnimeSerieResponse.animeSerie:type_name -> pb.AnimeSerieResponse
	5, // 3: pb.AddInfoAnimeSerieResponse.genres:type_name -> pb.AnimeGenresResponse
	6, // 4: pb.AddInfoAnimeSerieResponse.studios:type_name -> pb.AnimeStudiosResponse
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_rpc_add_info_anime_serie_proto_init() }
func file_rpc_add_info_anime_serie_proto_init() {
	if File_rpc_add_info_anime_serie_proto != nil {
		return
	}
	file_anime_serie_proto_init()
	file_genre_proto_init()
	file_studio_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_add_info_anime_serie_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddInfoAnimeSerieRequest); i {
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
		file_rpc_add_info_anime_serie_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddInfoAnimeSerieResponse); i {
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
			RawDescriptor: file_rpc_add_info_anime_serie_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_add_info_anime_serie_proto_goTypes,
		DependencyIndexes: file_rpc_add_info_anime_serie_proto_depIdxs,
		MessageInfos:      file_rpc_add_info_anime_serie_proto_msgTypes,
	}.Build()
	File_rpc_add_info_anime_serie_proto = out.File
	file_rpc_add_info_anime_serie_proto_rawDesc = nil
	file_rpc_add_info_anime_serie_proto_goTypes = nil
	file_rpc_add_info_anime_serie_proto_depIdxs = nil
}
