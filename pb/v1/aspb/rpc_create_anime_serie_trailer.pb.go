// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: v1/aspb/rpc_create_anime_serie_trailer.proto

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

type CreateAnimeSerieTrailerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnimeID       int64                      `protobuf:"varint,1,opt,name=animeID,proto3" json:"animeID,omitempty"`
	AnimeTrailers []*apb.AnimeTrailerRequest `protobuf:"bytes,2,rep,name=animeTrailers,proto3" json:"animeTrailers,omitempty"`
}

func (x *CreateAnimeSerieTrailerRequest) Reset() {
	*x = CreateAnimeSerieTrailerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_aspb_rpc_create_anime_serie_trailer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAnimeSerieTrailerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAnimeSerieTrailerRequest) ProtoMessage() {}

func (x *CreateAnimeSerieTrailerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_aspb_rpc_create_anime_serie_trailer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAnimeSerieTrailerRequest.ProtoReflect.Descriptor instead.
func (*CreateAnimeSerieTrailerRequest) Descriptor() ([]byte, []int) {
	return file_v1_aspb_rpc_create_anime_serie_trailer_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAnimeSerieTrailerRequest) GetAnimeID() int64 {
	if x != nil {
		return x.AnimeID
	}
	return 0
}

func (x *CreateAnimeSerieTrailerRequest) GetAnimeTrailers() []*apb.AnimeTrailerRequest {
	if x != nil {
		return x.AnimeTrailers
	}
	return nil
}

type CreateAnimeSerieTrailerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnimeSerie    *AnimeSerieResponse         `protobuf:"bytes,1,opt,name=animeSerie,proto3" json:"animeSerie,omitempty"`
	AnimeTrailers []*apb.AnimeTrailerResponse `protobuf:"bytes,2,rep,name=animeTrailers,proto3" json:"animeTrailers,omitempty"`
}

func (x *CreateAnimeSerieTrailerResponse) Reset() {
	*x = CreateAnimeSerieTrailerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_aspb_rpc_create_anime_serie_trailer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAnimeSerieTrailerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAnimeSerieTrailerResponse) ProtoMessage() {}

func (x *CreateAnimeSerieTrailerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_aspb_rpc_create_anime_serie_trailer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAnimeSerieTrailerResponse.ProtoReflect.Descriptor instead.
func (*CreateAnimeSerieTrailerResponse) Descriptor() ([]byte, []int) {
	return file_v1_aspb_rpc_create_anime_serie_trailer_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAnimeSerieTrailerResponse) GetAnimeSerie() *AnimeSerieResponse {
	if x != nil {
		return x.AnimeSerie
	}
	return nil
}

func (x *CreateAnimeSerieTrailerResponse) GetAnimeTrailers() []*apb.AnimeTrailerResponse {
	if x != nil {
		return x.AnimeTrailers
	}
	return nil
}

var File_v1_aspb_rpc_create_anime_serie_trailer_proto protoreflect.FileDescriptor

var file_v1_aspb_rpc_create_anime_serie_trailer_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x76, 0x31, 0x2f, 0x61, 0x73, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65,
	0x5f, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x76, 0x31, 0x2e, 0x61, 0x73, 0x70, 0x62, 0x76, 0x31, 0x1a, 0x1a, 0x76, 0x31, 0x2f, 0x61, 0x70,
	0x62, 0x2f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x76, 0x31, 0x2f, 0x61, 0x73, 0x70, 0x62, 0x2f, 0x6d,
	0x73, 0x67, 0x5f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7f, 0x0a, 0x1e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6e,
	0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x49, 0x44,
	0x12, 0x43, 0x0a, 0x0d, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x70, 0x62,
	0x76, 0x31, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0d, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x54, 0x72, 0x61,
	0x69, 0x6c, 0x65, 0x72, 0x73, 0x22, 0xa6, 0x01, 0x0a, 0x1f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x6e, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0a, 0x61, 0x6e, 0x69,
	0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x76, 0x31, 0x2e, 0x61, 0x73, 0x70, 0x62, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x53,
	0x65, 0x72, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x61, 0x6e,
	0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x12, 0x44, 0x0a, 0x0d, 0x61, 0x6e, 0x69, 0x6d,
	0x65, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x70, 0x62, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65,
	0x54, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x0d, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x73, 0x42, 0x35,
	0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6a, 0x2d,
	0x79, 0x61, 0x63, 0x69, 0x6e, 0x65, 0x2d, 0x66, 0x6c, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67,
	0x6f, 0x6a, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x73, 0x70, 0x62, 0x3b, 0x61,
	0x73, 0x70, 0x62, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_aspb_rpc_create_anime_serie_trailer_proto_rawDescOnce sync.Once
	file_v1_aspb_rpc_create_anime_serie_trailer_proto_rawDescData = file_v1_aspb_rpc_create_anime_serie_trailer_proto_rawDesc
)

func file_v1_aspb_rpc_create_anime_serie_trailer_proto_rawDescGZIP() []byte {
	file_v1_aspb_rpc_create_anime_serie_trailer_proto_rawDescOnce.Do(func() {
		file_v1_aspb_rpc_create_anime_serie_trailer_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_aspb_rpc_create_anime_serie_trailer_proto_rawDescData)
	})
	return file_v1_aspb_rpc_create_anime_serie_trailer_proto_rawDescData
}

var file_v1_aspb_rpc_create_anime_serie_trailer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_aspb_rpc_create_anime_serie_trailer_proto_goTypes = []interface{}{
	(*CreateAnimeSerieTrailerRequest)(nil),  // 0: v1.aspbv1.CreateAnimeSerieTrailerRequest
	(*CreateAnimeSerieTrailerResponse)(nil), // 1: v1.aspbv1.CreateAnimeSerieTrailerResponse
	(*apb.AnimeTrailerRequest)(nil),         // 2: v1.apbv1.AnimeTrailerRequest
	(*AnimeSerieResponse)(nil),              // 3: v1.aspbv1.AnimeSerieResponse
	(*apb.AnimeTrailerResponse)(nil),        // 4: v1.apbv1.AnimeTrailerResponse
}
var file_v1_aspb_rpc_create_anime_serie_trailer_proto_depIdxs = []int32{
	2, // 0: v1.aspbv1.CreateAnimeSerieTrailerRequest.animeTrailers:type_name -> v1.apbv1.AnimeTrailerRequest
	3, // 1: v1.aspbv1.CreateAnimeSerieTrailerResponse.animeSerie:type_name -> v1.aspbv1.AnimeSerieResponse
	4, // 2: v1.aspbv1.CreateAnimeSerieTrailerResponse.animeTrailers:type_name -> v1.apbv1.AnimeTrailerResponse
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_v1_aspb_rpc_create_anime_serie_trailer_proto_init() }
func file_v1_aspb_rpc_create_anime_serie_trailer_proto_init() {
	if File_v1_aspb_rpc_create_anime_serie_trailer_proto != nil {
		return
	}
	file_v1_aspb_msg_anime_serie_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_aspb_rpc_create_anime_serie_trailer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAnimeSerieTrailerRequest); i {
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
		file_v1_aspb_rpc_create_anime_serie_trailer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAnimeSerieTrailerResponse); i {
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
			RawDescriptor: file_v1_aspb_rpc_create_anime_serie_trailer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_aspb_rpc_create_anime_serie_trailer_proto_goTypes,
		DependencyIndexes: file_v1_aspb_rpc_create_anime_serie_trailer_proto_depIdxs,
		MessageInfos:      file_v1_aspb_rpc_create_anime_serie_trailer_proto_msgTypes,
	}.Build()
	File_v1_aspb_rpc_create_anime_serie_trailer_proto = out.File
	file_v1_aspb_rpc_create_anime_serie_trailer_proto_rawDesc = nil
	file_v1_aspb_rpc_create_anime_serie_trailer_proto_goTypes = nil
	file_v1_aspb_rpc_create_anime_serie_trailer_proto_depIdxs = nil
}
