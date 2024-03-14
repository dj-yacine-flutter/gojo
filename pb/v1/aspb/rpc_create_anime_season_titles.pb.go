// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: v1/aspb/rpc_create_anime_season_titles.proto

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

type CreateAnimeSeasonTitlesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeasonID     int64                   `protobuf:"varint,1,opt,name=seasonID,proto3" json:"seasonID,omitempty"`
	SeasonTitles *apb.AnimeTitlesRequest `protobuf:"bytes,2,opt,name=seasonTitles,proto3" json:"seasonTitles,omitempty"`
}

func (x *CreateAnimeSeasonTitlesRequest) Reset() {
	*x = CreateAnimeSeasonTitlesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_aspb_rpc_create_anime_season_titles_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAnimeSeasonTitlesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAnimeSeasonTitlesRequest) ProtoMessage() {}

func (x *CreateAnimeSeasonTitlesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_aspb_rpc_create_anime_season_titles_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAnimeSeasonTitlesRequest.ProtoReflect.Descriptor instead.
func (*CreateAnimeSeasonTitlesRequest) Descriptor() ([]byte, []int) {
	return file_v1_aspb_rpc_create_anime_season_titles_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAnimeSeasonTitlesRequest) GetSeasonID() int64 {
	if x != nil {
		return x.SeasonID
	}
	return 0
}

func (x *CreateAnimeSeasonTitlesRequest) GetSeasonTitles() *apb.AnimeTitlesRequest {
	if x != nil {
		return x.SeasonTitles
	}
	return nil
}

type CreateAnimeSeasonTitlesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeasonTitles *apb.AnimeTitlesResponse `protobuf:"bytes,1,opt,name=seasonTitles,proto3" json:"seasonTitles,omitempty"`
}

func (x *CreateAnimeSeasonTitlesResponse) Reset() {
	*x = CreateAnimeSeasonTitlesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_aspb_rpc_create_anime_season_titles_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAnimeSeasonTitlesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAnimeSeasonTitlesResponse) ProtoMessage() {}

func (x *CreateAnimeSeasonTitlesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_aspb_rpc_create_anime_season_titles_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAnimeSeasonTitlesResponse.ProtoReflect.Descriptor instead.
func (*CreateAnimeSeasonTitlesResponse) Descriptor() ([]byte, []int) {
	return file_v1_aspb_rpc_create_anime_season_titles_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAnimeSeasonTitlesResponse) GetSeasonTitles() *apb.AnimeTitlesResponse {
	if x != nil {
		return x.SeasonTitles
	}
	return nil
}

var File_v1_aspb_rpc_create_anime_season_titles_proto protoreflect.FileDescriptor

var file_v1_aspb_rpc_create_anime_season_titles_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x76, 0x31, 0x2f, 0x61, 0x73, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x76, 0x31, 0x2e, 0x61, 0x73, 0x70, 0x62, 0x76, 0x31, 0x1a, 0x18, 0x76, 0x31, 0x2f, 0x61, 0x70,
	0x62, 0x2f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x7e, 0x0a, 0x1e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x69,
	0x6d, 0x65, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49,
	0x44, 0x12, 0x40, 0x0a, 0x0c, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x70, 0x62,
	0x76, 0x31, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0c, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x73, 0x22, 0x64, 0x0a, 0x1f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x69,
	0x6d, 0x65, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0c, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x76,
	0x31, 0x2e, 0x61, 0x70, 0x62, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0c, 0x73, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x73, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6a, 0x2d, 0x79, 0x61, 0x63, 0x69, 0x6e,
	0x65, 0x2d, 0x66, 0x6c, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x6f, 0x6a, 0x6f, 0x2f, 0x70,
	0x62, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x73, 0x70, 0x62, 0x3b, 0x61, 0x73, 0x70, 0x62, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_aspb_rpc_create_anime_season_titles_proto_rawDescOnce sync.Once
	file_v1_aspb_rpc_create_anime_season_titles_proto_rawDescData = file_v1_aspb_rpc_create_anime_season_titles_proto_rawDesc
)

func file_v1_aspb_rpc_create_anime_season_titles_proto_rawDescGZIP() []byte {
	file_v1_aspb_rpc_create_anime_season_titles_proto_rawDescOnce.Do(func() {
		file_v1_aspb_rpc_create_anime_season_titles_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_aspb_rpc_create_anime_season_titles_proto_rawDescData)
	})
	return file_v1_aspb_rpc_create_anime_season_titles_proto_rawDescData
}

var file_v1_aspb_rpc_create_anime_season_titles_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_aspb_rpc_create_anime_season_titles_proto_goTypes = []interface{}{
	(*CreateAnimeSeasonTitlesRequest)(nil),  // 0: v1.aspbv1.CreateAnimeSeasonTitlesRequest
	(*CreateAnimeSeasonTitlesResponse)(nil), // 1: v1.aspbv1.CreateAnimeSeasonTitlesResponse
	(*apb.AnimeTitlesRequest)(nil),          // 2: v1.apbv1.AnimeTitlesRequest
	(*apb.AnimeTitlesResponse)(nil),         // 3: v1.apbv1.AnimeTitlesResponse
}
var file_v1_aspb_rpc_create_anime_season_titles_proto_depIdxs = []int32{
	2, // 0: v1.aspbv1.CreateAnimeSeasonTitlesRequest.seasonTitles:type_name -> v1.apbv1.AnimeTitlesRequest
	3, // 1: v1.aspbv1.CreateAnimeSeasonTitlesResponse.seasonTitles:type_name -> v1.apbv1.AnimeTitlesResponse
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_v1_aspb_rpc_create_anime_season_titles_proto_init() }
func file_v1_aspb_rpc_create_anime_season_titles_proto_init() {
	if File_v1_aspb_rpc_create_anime_season_titles_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_aspb_rpc_create_anime_season_titles_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAnimeSeasonTitlesRequest); i {
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
		file_v1_aspb_rpc_create_anime_season_titles_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAnimeSeasonTitlesResponse); i {
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
			RawDescriptor: file_v1_aspb_rpc_create_anime_season_titles_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_aspb_rpc_create_anime_season_titles_proto_goTypes,
		DependencyIndexes: file_v1_aspb_rpc_create_anime_season_titles_proto_depIdxs,
		MessageInfos:      file_v1_aspb_rpc_create_anime_season_titles_proto_msgTypes,
	}.Build()
	File_v1_aspb_rpc_create_anime_season_titles_proto = out.File
	file_v1_aspb_rpc_create_anime_season_titles_proto_rawDesc = nil
	file_v1_aspb_rpc_create_anime_season_titles_proto_goTypes = nil
	file_v1_aspb_rpc_create_anime_season_titles_proto_depIdxs = nil
}
