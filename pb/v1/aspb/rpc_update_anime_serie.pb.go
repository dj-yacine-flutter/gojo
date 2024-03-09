// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: v1/aspb/rpc_update_anime_serie.proto

package aspbv1

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

type UpdateAnimeSerieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnimeID           int64   `protobuf:"varint,1,opt,name=animeID,proto3" json:"animeID,omitempty"`
	OriginalTitle     *string `protobuf:"bytes,2,opt,name=originalTitle,proto3,oneof" json:"originalTitle,omitempty"`
	FirstYear         *int32  `protobuf:"varint,3,opt,name=firstYear,proto3,oneof" json:"firstYear,omitempty"`
	LastYear          *int32  `protobuf:"varint,4,opt,name=lastYear,proto3,oneof" json:"lastYear,omitempty"`
	MalID             *int32  `protobuf:"varint,5,opt,name=malID,proto3,oneof" json:"malID,omitempty"`
	TvdbID            *int32  `protobuf:"varint,6,opt,name=tvdbID,proto3,oneof" json:"tvdbID,omitempty"`
	TmdbID            *int32  `protobuf:"varint,7,opt,name=tmdbID,proto3,oneof" json:"tmdbID,omitempty"`
	PortraitPoster    *string `protobuf:"bytes,8,opt,name=portraitPoster,proto3,oneof" json:"portraitPoster,omitempty"`
	PortraitBlurHash  *string `protobuf:"bytes,9,opt,name=portraitBlurHash,proto3,oneof" json:"portraitBlurHash,omitempty"`
	LandscapePoster   *string `protobuf:"bytes,10,opt,name=landscapePoster,proto3,oneof" json:"landscapePoster,omitempty"`
	LandscapeBlurHash *string `protobuf:"bytes,11,opt,name=landscapeBlurHash,proto3,oneof" json:"landscapeBlurHash,omitempty"`
}

func (x *UpdateAnimeSerieRequest) Reset() {
	*x = UpdateAnimeSerieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_aspb_rpc_update_anime_serie_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAnimeSerieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAnimeSerieRequest) ProtoMessage() {}

func (x *UpdateAnimeSerieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_aspb_rpc_update_anime_serie_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAnimeSerieRequest.ProtoReflect.Descriptor instead.
func (*UpdateAnimeSerieRequest) Descriptor() ([]byte, []int) {
	return file_v1_aspb_rpc_update_anime_serie_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateAnimeSerieRequest) GetAnimeID() int64 {
	if x != nil {
		return x.AnimeID
	}
	return 0
}

func (x *UpdateAnimeSerieRequest) GetOriginalTitle() string {
	if x != nil && x.OriginalTitle != nil {
		return *x.OriginalTitle
	}
	return ""
}

func (x *UpdateAnimeSerieRequest) GetFirstYear() int32 {
	if x != nil && x.FirstYear != nil {
		return *x.FirstYear
	}
	return 0
}

func (x *UpdateAnimeSerieRequest) GetLastYear() int32 {
	if x != nil && x.LastYear != nil {
		return *x.LastYear
	}
	return 0
}

func (x *UpdateAnimeSerieRequest) GetMalID() int32 {
	if x != nil && x.MalID != nil {
		return *x.MalID
	}
	return 0
}

func (x *UpdateAnimeSerieRequest) GetTvdbID() int32 {
	if x != nil && x.TvdbID != nil {
		return *x.TvdbID
	}
	return 0
}

func (x *UpdateAnimeSerieRequest) GetTmdbID() int32 {
	if x != nil && x.TmdbID != nil {
		return *x.TmdbID
	}
	return 0
}

func (x *UpdateAnimeSerieRequest) GetPortraitPoster() string {
	if x != nil && x.PortraitPoster != nil {
		return *x.PortraitPoster
	}
	return ""
}

func (x *UpdateAnimeSerieRequest) GetPortraitBlurHash() string {
	if x != nil && x.PortraitBlurHash != nil {
		return *x.PortraitBlurHash
	}
	return ""
}

func (x *UpdateAnimeSerieRequest) GetLandscapePoster() string {
	if x != nil && x.LandscapePoster != nil {
		return *x.LandscapePoster
	}
	return ""
}

func (x *UpdateAnimeSerieRequest) GetLandscapeBlurHash() string {
	if x != nil && x.LandscapeBlurHash != nil {
		return *x.LandscapeBlurHash
	}
	return ""
}

type UpdateAnimeSerieResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnimeSerie *AnimeSerieResponse `protobuf:"bytes,1,opt,name=animeSerie,proto3" json:"animeSerie,omitempty"`
}

func (x *UpdateAnimeSerieResponse) Reset() {
	*x = UpdateAnimeSerieResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_aspb_rpc_update_anime_serie_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAnimeSerieResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAnimeSerieResponse) ProtoMessage() {}

func (x *UpdateAnimeSerieResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_aspb_rpc_update_anime_serie_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAnimeSerieResponse.ProtoReflect.Descriptor instead.
func (*UpdateAnimeSerieResponse) Descriptor() ([]byte, []int) {
	return file_v1_aspb_rpc_update_anime_serie_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateAnimeSerieResponse) GetAnimeSerie() *AnimeSerieResponse {
	if x != nil {
		return x.AnimeSerie
	}
	return nil
}

var File_v1_aspb_rpc_update_anime_serie_proto protoreflect.FileDescriptor

var file_v1_aspb_rpc_update_anime_serie_proto_rawDesc = []byte{
	0x0a, 0x24, 0x76, 0x31, 0x2f, 0x61, 0x73, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x76, 0x31, 0x2e, 0x61, 0x73, 0x70, 0x62, 0x76,
	0x31, 0x1a, 0x1d, 0x76, 0x31, 0x2f, 0x61, 0x73, 0x70, 0x62, 0x2f, 0x6d, 0x73, 0x67, 0x5f, 0x61,
	0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xd6, 0x04, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65,
	0x53, 0x65, 0x72, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x6e, 0x69, 0x6d, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61,
	0x6e, 0x69, 0x6d, 0x65, 0x49, 0x44, 0x12, 0x29, 0x0a, 0x0d, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x61, 0x6c, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x0d, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x21, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x59, 0x65, 0x61, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x59, 0x65, 0x61,
	0x72, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x59, 0x65, 0x61, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x02, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x59, 0x65,
	0x61, 0x72, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x6d, 0x61, 0x6c, 0x49, 0x44, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x48, 0x03, 0x52, 0x05, 0x6d, 0x61, 0x6c, 0x49, 0x44, 0x88, 0x01, 0x01,
	0x12, 0x1b, 0x0a, 0x06, 0x74, 0x76, 0x64, 0x62, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x48, 0x04, 0x52, 0x06, 0x74, 0x76, 0x64, 0x62, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a,
	0x06, 0x74, 0x6d, 0x64, 0x62, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x48, 0x05, 0x52,
	0x06, 0x74, 0x6d, 0x64, 0x62, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x0e, 0x70, 0x6f,
	0x72, 0x74, 0x72, 0x61, 0x69, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x06, 0x52, 0x0e, 0x70, 0x6f, 0x72, 0x74, 0x72, 0x61, 0x69, 0x74, 0x50, 0x6f,
	0x73, 0x74, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a, 0x10, 0x70, 0x6f, 0x72, 0x74, 0x72,
	0x61, 0x69, 0x74, 0x42, 0x6c, 0x75, 0x72, 0x48, 0x61, 0x73, 0x68, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x07, 0x52, 0x10, 0x70, 0x6f, 0x72, 0x74, 0x72, 0x61, 0x69, 0x74, 0x42, 0x6c, 0x75,
	0x72, 0x48, 0x61, 0x73, 0x68, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x0f, 0x6c, 0x61, 0x6e, 0x64,
	0x73, 0x63, 0x61, 0x70, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x08, 0x52, 0x0f, 0x6c, 0x61, 0x6e, 0x64, 0x73, 0x63, 0x61, 0x70, 0x65, 0x50, 0x6f,
	0x73, 0x74, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x31, 0x0a, 0x11, 0x6c, 0x61, 0x6e, 0x64, 0x73,
	0x63, 0x61, 0x70, 0x65, 0x42, 0x6c, 0x75, 0x72, 0x48, 0x61, 0x73, 0x68, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x09, 0x52, 0x11, 0x6c, 0x61, 0x6e, 0x64, 0x73, 0x63, 0x61, 0x70, 0x65, 0x42,
	0x6c, 0x75, 0x72, 0x48, 0x61, 0x73, 0x68, 0x88, 0x01, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x42, 0x0c, 0x0a, 0x0a,
	0x5f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x59, 0x65, 0x61, 0x72, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6c,
	0x61, 0x73, 0x74, 0x59, 0x65, 0x61, 0x72, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6d, 0x61, 0x6c, 0x49,
	0x44, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x74, 0x76, 0x64, 0x62, 0x49, 0x44, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x74, 0x6d, 0x64, 0x62, 0x49, 0x44, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x70, 0x6f, 0x72, 0x74,
	0x72, 0x61, 0x69, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x70,
	0x6f, 0x72, 0x74, 0x72, 0x61, 0x69, 0x74, 0x42, 0x6c, 0x75, 0x72, 0x48, 0x61, 0x73, 0x68, 0x42,
	0x12, 0x0a, 0x10, 0x5f, 0x6c, 0x61, 0x6e, 0x64, 0x73, 0x63, 0x61, 0x70, 0x65, 0x50, 0x6f, 0x73,
	0x74, 0x65, 0x72, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x6c, 0x61, 0x6e, 0x64, 0x73, 0x63, 0x61, 0x70,
	0x65, 0x42, 0x6c, 0x75, 0x72, 0x48, 0x61, 0x73, 0x68, 0x22, 0x59, 0x0a, 0x18, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0a, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x53, 0x65,
	0x72, 0x69, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x76, 0x31, 0x2e, 0x61,
	0x73, 0x70, 0x62, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x53,
	0x65, 0x72, 0x69, 0x65, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x64, 0x6a, 0x2d, 0x79, 0x61, 0x63, 0x69, 0x6e, 0x65, 0x2d, 0x66, 0x6c, 0x75,
	0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x6f, 0x6a, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x73, 0x70, 0x62, 0x3b, 0x61, 0x73, 0x70, 0x62, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_v1_aspb_rpc_update_anime_serie_proto_rawDescOnce sync.Once
	file_v1_aspb_rpc_update_anime_serie_proto_rawDescData = file_v1_aspb_rpc_update_anime_serie_proto_rawDesc
)

func file_v1_aspb_rpc_update_anime_serie_proto_rawDescGZIP() []byte {
	file_v1_aspb_rpc_update_anime_serie_proto_rawDescOnce.Do(func() {
		file_v1_aspb_rpc_update_anime_serie_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_aspb_rpc_update_anime_serie_proto_rawDescData)
	})
	return file_v1_aspb_rpc_update_anime_serie_proto_rawDescData
}

var file_v1_aspb_rpc_update_anime_serie_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_aspb_rpc_update_anime_serie_proto_goTypes = []interface{}{
	(*UpdateAnimeSerieRequest)(nil),  // 0: v1.aspbv1.UpdateAnimeSerieRequest
	(*UpdateAnimeSerieResponse)(nil), // 1: v1.aspbv1.UpdateAnimeSerieResponse
	(*AnimeSerieResponse)(nil),       // 2: v1.aspbv1.AnimeSerieResponse
}
var file_v1_aspb_rpc_update_anime_serie_proto_depIdxs = []int32{
	2, // 0: v1.aspbv1.UpdateAnimeSerieResponse.animeSerie:type_name -> v1.aspbv1.AnimeSerieResponse
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v1_aspb_rpc_update_anime_serie_proto_init() }
func file_v1_aspb_rpc_update_anime_serie_proto_init() {
	if File_v1_aspb_rpc_update_anime_serie_proto != nil {
		return
	}
	file_v1_aspb_msg_anime_serie_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_aspb_rpc_update_anime_serie_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAnimeSerieRequest); i {
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
		file_v1_aspb_rpc_update_anime_serie_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAnimeSerieResponse); i {
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
	file_v1_aspb_rpc_update_anime_serie_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_aspb_rpc_update_anime_serie_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_aspb_rpc_update_anime_serie_proto_goTypes,
		DependencyIndexes: file_v1_aspb_rpc_update_anime_serie_proto_depIdxs,
		MessageInfos:      file_v1_aspb_rpc_update_anime_serie_proto_msgTypes,
	}.Build()
	File_v1_aspb_rpc_update_anime_serie_proto = out.File
	file_v1_aspb_rpc_update_anime_serie_proto_rawDesc = nil
	file_v1_aspb_rpc_update_anime_serie_proto_goTypes = nil
	file_v1_aspb_rpc_update_anime_serie_proto_depIdxs = nil
}
