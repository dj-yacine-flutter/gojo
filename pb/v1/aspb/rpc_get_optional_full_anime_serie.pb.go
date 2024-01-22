// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: v1/aspb/rpc_get_optional_full_anime_serie.proto

package aspbv1

import (
	apb "github.com/dj-yacine-flutter/gojo/pb/v1/apb"
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

type GetOptionalFullAnimeSerieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnimeID     int64 `protobuf:"varint,1,opt,name=animeID,proto3" json:"animeID,omitempty"`
	LanguageID  int32 `protobuf:"varint,2,opt,name=languageID,proto3" json:"languageID,omitempty"`
	WithLinks   *bool `protobuf:"varint,3,opt,name=withLinks,proto3,oneof" json:"withLinks,omitempty"`
	WithImages  *bool `protobuf:"varint,4,opt,name=withImages,proto3,oneof" json:"withImages,omitempty"`
	WithTrailer *bool `protobuf:"varint,5,opt,name=withTrailer,proto3,oneof" json:"withTrailer,omitempty"`
}

func (x *GetOptionalFullAnimeSerieRequest) Reset() {
	*x = GetOptionalFullAnimeSerieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_aspb_rpc_get_optional_full_anime_serie_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOptionalFullAnimeSerieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOptionalFullAnimeSerieRequest) ProtoMessage() {}

func (x *GetOptionalFullAnimeSerieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_aspb_rpc_get_optional_full_anime_serie_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOptionalFullAnimeSerieRequest.ProtoReflect.Descriptor instead.
func (*GetOptionalFullAnimeSerieRequest) Descriptor() ([]byte, []int) {
	return file_v1_aspb_rpc_get_optional_full_anime_serie_proto_rawDescGZIP(), []int{0}
}

func (x *GetOptionalFullAnimeSerieRequest) GetAnimeID() int64 {
	if x != nil {
		return x.AnimeID
	}
	return 0
}

func (x *GetOptionalFullAnimeSerieRequest) GetLanguageID() int32 {
	if x != nil {
		return x.LanguageID
	}
	return 0
}

func (x *GetOptionalFullAnimeSerieRequest) GetWithLinks() bool {
	if x != nil && x.WithLinks != nil {
		return *x.WithLinks
	}
	return false
}

func (x *GetOptionalFullAnimeSerieRequest) GetWithImages() bool {
	if x != nil && x.WithImages != nil {
		return *x.WithImages
	}
	return false
}

func (x *GetOptionalFullAnimeSerieRequest) GetWithTrailer() bool {
	if x != nil && x.WithTrailer != nil {
		return *x.WithTrailer
	}
	return false
}

type GetOptionalFullAnimeSerieResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnimeSerie    *AnimeSerieResponse         `protobuf:"bytes,1,opt,name=animeSerie,proto3" json:"animeSerie,omitempty"`
	AnimeMeta     *nfpb.AnimeMetaResponse     `protobuf:"bytes,2,opt,name=animeMeta,proto3,oneof" json:"animeMeta,omitempty"`
	AnimeLinks    *apb.AnimeLinkResponse      `protobuf:"bytes,3,opt,name=animeLinks,proto3,oneof" json:"animeLinks,omitempty"`
	AnimeImages   *apb.AnimeImageResponse     `protobuf:"bytes,4,opt,name=animeImages,proto3,oneof" json:"animeImages,omitempty"`
	AnimeTrailers []*apb.AnimeTrailerResponse `protobuf:"bytes,5,rep,name=animeTrailers,proto3" json:"animeTrailers,omitempty"`
}

func (x *GetOptionalFullAnimeSerieResponse) Reset() {
	*x = GetOptionalFullAnimeSerieResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_aspb_rpc_get_optional_full_anime_serie_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOptionalFullAnimeSerieResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOptionalFullAnimeSerieResponse) ProtoMessage() {}

func (x *GetOptionalFullAnimeSerieResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_aspb_rpc_get_optional_full_anime_serie_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOptionalFullAnimeSerieResponse.ProtoReflect.Descriptor instead.
func (*GetOptionalFullAnimeSerieResponse) Descriptor() ([]byte, []int) {
	return file_v1_aspb_rpc_get_optional_full_anime_serie_proto_rawDescGZIP(), []int{1}
}

func (x *GetOptionalFullAnimeSerieResponse) GetAnimeSerie() *AnimeSerieResponse {
	if x != nil {
		return x.AnimeSerie
	}
	return nil
}

func (x *GetOptionalFullAnimeSerieResponse) GetAnimeMeta() *nfpb.AnimeMetaResponse {
	if x != nil {
		return x.AnimeMeta
	}
	return nil
}

func (x *GetOptionalFullAnimeSerieResponse) GetAnimeLinks() *apb.AnimeLinkResponse {
	if x != nil {
		return x.AnimeLinks
	}
	return nil
}

func (x *GetOptionalFullAnimeSerieResponse) GetAnimeImages() *apb.AnimeImageResponse {
	if x != nil {
		return x.AnimeImages
	}
	return nil
}

func (x *GetOptionalFullAnimeSerieResponse) GetAnimeTrailers() []*apb.AnimeTrailerResponse {
	if x != nil {
		return x.AnimeTrailers
	}
	return nil
}

var File_v1_aspb_rpc_get_optional_full_anime_serie_proto protoreflect.FileDescriptor

var file_v1_aspb_rpc_get_optional_full_anime_serie_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x73, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65,
	0x74, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x66, 0x75, 0x6c, 0x6c, 0x5f,
	0x61, 0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x76, 0x31, 0x2e, 0x61, 0x73, 0x70, 0x62, 0x76, 0x31, 0x1a, 0x1d, 0x76, 0x31,
	0x2f, 0x61, 0x73, 0x70, 0x62, 0x2f, 0x6d, 0x73, 0x67, 0x5f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x69, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x76, 0x31, 0x2f,
	0x6e, 0x66, 0x70, 0x62, 0x2f, 0x6d, 0x73, 0x67, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x62, 0x2f, 0x61, 0x6e, 0x69, 0x6d,
	0x65, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x76, 0x31,
	0x2f, 0x61, 0x70, 0x62, 0x2f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x62, 0x2f, 0x61,
	0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xf8, 0x01, 0x0a, 0x20, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x46, 0x75, 0x6c, 0x6c, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x6e, 0x69, 0x6d, 0x65,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x49,
	0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x49,
	0x44, 0x12, 0x21, 0x0a, 0x09, 0x77, 0x69, 0x74, 0x68, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x09, 0x77, 0x69, 0x74, 0x68, 0x4c, 0x69, 0x6e, 0x6b,
	0x73, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x77, 0x69, 0x74, 0x68, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x48, 0x01, 0x52, 0x0a, 0x77, 0x69, 0x74, 0x68,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x77, 0x69, 0x74,
	0x68, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x48, 0x02,
	0x52, 0x0b, 0x77, 0x69, 0x74, 0x68, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x88, 0x01, 0x01,
	0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x42, 0x0d,
	0x0a, 0x0b, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x42, 0x0e, 0x0a,
	0x0c, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x22, 0x9d, 0x03,
	0x0a, 0x21, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x46, 0x75, 0x6c,
	0x6c, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0a, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x73, 0x70,
	0x62, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72,
	0x69, 0x65, 0x12, 0x3f, 0x0a, 0x09, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x76, 0x31, 0x2e, 0x6e, 0x66, 0x70, 0x62, 0x76,
	0x31, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x09, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x65, 0x74, 0x61,
	0x88, 0x01, 0x01, 0x12, 0x40, 0x0a, 0x0a, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x4c, 0x69, 0x6e, 0x6b,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x70, 0x62,
	0x76, 0x31, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x48, 0x01, 0x52, 0x0a, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x4c, 0x69, 0x6e,
	0x6b, 0x73, 0x88, 0x01, 0x01, 0x12, 0x43, 0x0a, 0x0b, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x76, 0x31, 0x2e,
	0x61, 0x70, 0x62, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x02, 0x52, 0x0b, 0x61, 0x6e, 0x69, 0x6d,
	0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x88, 0x01, 0x01, 0x12, 0x44, 0x0a, 0x0d, 0x61, 0x6e,
	0x69, 0x6d, 0x65, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x70, 0x62, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x69,
	0x6d, 0x65, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x0d, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x73,
	0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x42, 0x0d,
	0x0a, 0x0b, 0x5f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x42, 0x0e, 0x0a,
	0x0c, 0x5f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x42, 0x35, 0x5a,
	0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6a, 0x2d, 0x79,
	0x61, 0x63, 0x69, 0x6e, 0x65, 0x2d, 0x66, 0x6c, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x6f,
	0x6a, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x73, 0x70, 0x62, 0x3b, 0x61, 0x73,
	0x70, 0x62, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_aspb_rpc_get_optional_full_anime_serie_proto_rawDescOnce sync.Once
	file_v1_aspb_rpc_get_optional_full_anime_serie_proto_rawDescData = file_v1_aspb_rpc_get_optional_full_anime_serie_proto_rawDesc
)

func file_v1_aspb_rpc_get_optional_full_anime_serie_proto_rawDescGZIP() []byte {
	file_v1_aspb_rpc_get_optional_full_anime_serie_proto_rawDescOnce.Do(func() {
		file_v1_aspb_rpc_get_optional_full_anime_serie_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_aspb_rpc_get_optional_full_anime_serie_proto_rawDescData)
	})
	return file_v1_aspb_rpc_get_optional_full_anime_serie_proto_rawDescData
}

var file_v1_aspb_rpc_get_optional_full_anime_serie_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_aspb_rpc_get_optional_full_anime_serie_proto_goTypes = []interface{}{
	(*GetOptionalFullAnimeSerieRequest)(nil),  // 0: v1.aspbv1.GetOptionalFullAnimeSerieRequest
	(*GetOptionalFullAnimeSerieResponse)(nil), // 1: v1.aspbv1.GetOptionalFullAnimeSerieResponse
	(*AnimeSerieResponse)(nil),                // 2: v1.aspbv1.AnimeSerieResponse
	(*nfpb.AnimeMetaResponse)(nil),            // 3: v1.nfpbv1.AnimeMetaResponse
	(*apb.AnimeLinkResponse)(nil),             // 4: v1.apbv1.AnimeLinkResponse
	(*apb.AnimeImageResponse)(nil),            // 5: v1.apbv1.AnimeImageResponse
	(*apb.AnimeTrailerResponse)(nil),          // 6: v1.apbv1.AnimeTrailerResponse
}
var file_v1_aspb_rpc_get_optional_full_anime_serie_proto_depIdxs = []int32{
	2, // 0: v1.aspbv1.GetOptionalFullAnimeSerieResponse.animeSerie:type_name -> v1.aspbv1.AnimeSerieResponse
	3, // 1: v1.aspbv1.GetOptionalFullAnimeSerieResponse.animeMeta:type_name -> v1.nfpbv1.AnimeMetaResponse
	4, // 2: v1.aspbv1.GetOptionalFullAnimeSerieResponse.animeLinks:type_name -> v1.apbv1.AnimeLinkResponse
	5, // 3: v1.aspbv1.GetOptionalFullAnimeSerieResponse.animeImages:type_name -> v1.apbv1.AnimeImageResponse
	6, // 4: v1.aspbv1.GetOptionalFullAnimeSerieResponse.animeTrailers:type_name -> v1.apbv1.AnimeTrailerResponse
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_v1_aspb_rpc_get_optional_full_anime_serie_proto_init() }
func file_v1_aspb_rpc_get_optional_full_anime_serie_proto_init() {
	if File_v1_aspb_rpc_get_optional_full_anime_serie_proto != nil {
		return
	}
	file_v1_aspb_msg_anime_serie_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_aspb_rpc_get_optional_full_anime_serie_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOptionalFullAnimeSerieRequest); i {
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
		file_v1_aspb_rpc_get_optional_full_anime_serie_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOptionalFullAnimeSerieResponse); i {
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
	file_v1_aspb_rpc_get_optional_full_anime_serie_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_v1_aspb_rpc_get_optional_full_anime_serie_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_aspb_rpc_get_optional_full_anime_serie_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_aspb_rpc_get_optional_full_anime_serie_proto_goTypes,
		DependencyIndexes: file_v1_aspb_rpc_get_optional_full_anime_serie_proto_depIdxs,
		MessageInfos:      file_v1_aspb_rpc_get_optional_full_anime_serie_proto_msgTypes,
	}.Build()
	File_v1_aspb_rpc_get_optional_full_anime_serie_proto = out.File
	file_v1_aspb_rpc_get_optional_full_anime_serie_proto_rawDesc = nil
	file_v1_aspb_rpc_get_optional_full_anime_serie_proto_goTypes = nil
	file_v1_aspb_rpc_get_optional_full_anime_serie_proto_depIdxs = nil
}
