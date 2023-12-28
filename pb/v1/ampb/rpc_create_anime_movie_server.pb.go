// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: v1/ampb/rpc_create_anime_movie_server.proto

package ampbv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateAnimeMovieServerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnimeID int64 `protobuf:"varint,1,opt,name=animeID,proto3" json:"animeID,omitempty"`
}

func (x *CreateAnimeMovieServerRequest) Reset() {
	*x = CreateAnimeMovieServerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_ampb_rpc_create_anime_movie_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAnimeMovieServerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAnimeMovieServerRequest) ProtoMessage() {}

func (x *CreateAnimeMovieServerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_ampb_rpc_create_anime_movie_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAnimeMovieServerRequest.ProtoReflect.Descriptor instead.
func (*CreateAnimeMovieServerRequest) Descriptor() ([]byte, []int) {
	return file_v1_ampb_rpc_create_anime_movie_server_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAnimeMovieServerRequest) GetAnimeID() int64 {
	if x != nil {
		return x.AnimeID
	}
	return 0
}

type CreateAnimeMovieServerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnimeID   int64                  `protobuf:"varint,1,opt,name=animeID,proto3" json:"animeID,omitempty"`
	ServerID  int64                  `protobuf:"varint,2,opt,name=serverID,proto3" json:"serverID,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *CreateAnimeMovieServerResponse) Reset() {
	*x = CreateAnimeMovieServerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_ampb_rpc_create_anime_movie_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAnimeMovieServerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAnimeMovieServerResponse) ProtoMessage() {}

func (x *CreateAnimeMovieServerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_ampb_rpc_create_anime_movie_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAnimeMovieServerResponse.ProtoReflect.Descriptor instead.
func (*CreateAnimeMovieServerResponse) Descriptor() ([]byte, []int) {
	return file_v1_ampb_rpc_create_anime_movie_server_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAnimeMovieServerResponse) GetAnimeID() int64 {
	if x != nil {
		return x.AnimeID
	}
	return 0
}

func (x *CreateAnimeMovieServerResponse) GetServerID() int64 {
	if x != nil {
		return x.ServerID
	}
	return 0
}

func (x *CreateAnimeMovieServerResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_v1_ampb_rpc_create_anime_movie_server_proto protoreflect.FileDescriptor

var file_v1_ampb_rpc_create_anime_movie_server_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x76, 0x31, 0x2f, 0x61, 0x6d, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x6f, 0x76, 0x69, 0x65,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x76,
	0x31, 0x2e, 0x61, 0x6d, 0x70, 0x62, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a, 0x1d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x6e,
	0x69, 0x6d, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x6e, 0x69,
	0x6d, 0x65, 0x49, 0x44, 0x22, 0x90, 0x01, 0x0a, 0x1e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x6e, 0x69, 0x6d, 0x65,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x49,
	0x44, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x44, 0x12, 0x38, 0x0a,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6a, 0x2d, 0x79, 0x61, 0x63, 0x69, 0x6e, 0x65, 0x2d,
	0x66, 0x6c, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x6f, 0x6a, 0x6f, 0x2f, 0x70, 0x62, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x6d, 0x70, 0x62, 0x3b, 0x61, 0x6d, 0x70, 0x62, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_ampb_rpc_create_anime_movie_server_proto_rawDescOnce sync.Once
	file_v1_ampb_rpc_create_anime_movie_server_proto_rawDescData = file_v1_ampb_rpc_create_anime_movie_server_proto_rawDesc
)

func file_v1_ampb_rpc_create_anime_movie_server_proto_rawDescGZIP() []byte {
	file_v1_ampb_rpc_create_anime_movie_server_proto_rawDescOnce.Do(func() {
		file_v1_ampb_rpc_create_anime_movie_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_ampb_rpc_create_anime_movie_server_proto_rawDescData)
	})
	return file_v1_ampb_rpc_create_anime_movie_server_proto_rawDescData
}

var file_v1_ampb_rpc_create_anime_movie_server_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_ampb_rpc_create_anime_movie_server_proto_goTypes = []interface{}{
	(*CreateAnimeMovieServerRequest)(nil),  // 0: v1.ampbv1.CreateAnimeMovieServerRequest
	(*CreateAnimeMovieServerResponse)(nil), // 1: v1.ampbv1.CreateAnimeMovieServerResponse
	(*timestamppb.Timestamp)(nil),          // 2: google.protobuf.Timestamp
}
var file_v1_ampb_rpc_create_anime_movie_server_proto_depIdxs = []int32{
	2, // 0: v1.ampbv1.CreateAnimeMovieServerResponse.createdAt:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v1_ampb_rpc_create_anime_movie_server_proto_init() }
func file_v1_ampb_rpc_create_anime_movie_server_proto_init() {
	if File_v1_ampb_rpc_create_anime_movie_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_ampb_rpc_create_anime_movie_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAnimeMovieServerRequest); i {
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
		file_v1_ampb_rpc_create_anime_movie_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAnimeMovieServerResponse); i {
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
			RawDescriptor: file_v1_ampb_rpc_create_anime_movie_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_ampb_rpc_create_anime_movie_server_proto_goTypes,
		DependencyIndexes: file_v1_ampb_rpc_create_anime_movie_server_proto_depIdxs,
		MessageInfos:      file_v1_ampb_rpc_create_anime_movie_server_proto_msgTypes,
	}.Build()
	File_v1_ampb_rpc_create_anime_movie_server_proto = out.File
	file_v1_ampb_rpc_create_anime_movie_server_proto_rawDesc = nil
	file_v1_ampb_rpc_create_anime_movie_server_proto_goTypes = nil
	file_v1_ampb_rpc_create_anime_movie_server_proto_depIdxs = nil
}
