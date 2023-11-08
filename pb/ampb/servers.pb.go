// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: ampb/servers.proto

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

type AnimeMovieSubDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Videos   []*AnimeMovieVideoRequest   `protobuf:"bytes,1,rep,name=videos,proto3" json:"videos,omitempty"`
	Torrents []*AnimeMovieTorrentRequest `protobuf:"bytes,2,rep,name=torrents,proto3" json:"torrents,omitempty"`
}

func (x *AnimeMovieSubDataRequest) Reset() {
	*x = AnimeMovieSubDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ampb_servers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnimeMovieSubDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnimeMovieSubDataRequest) ProtoMessage() {}

func (x *AnimeMovieSubDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ampb_servers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnimeMovieSubDataRequest.ProtoReflect.Descriptor instead.
func (*AnimeMovieSubDataRequest) Descriptor() ([]byte, []int) {
	return file_ampb_servers_proto_rawDescGZIP(), []int{0}
}

func (x *AnimeMovieSubDataRequest) GetVideos() []*AnimeMovieVideoRequest {
	if x != nil {
		return x.Videos
	}
	return nil
}

func (x *AnimeMovieSubDataRequest) GetTorrents() []*AnimeMovieTorrentRequest {
	if x != nil {
		return x.Torrents
	}
	return nil
}

type AnimeMovieDubDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Videos   []*AnimeMovieVideoRequest   `protobuf:"bytes,1,rep,name=videos,proto3" json:"videos,omitempty"`
	Torrents []*AnimeMovieTorrentRequest `protobuf:"bytes,2,rep,name=torrents,proto3" json:"torrents,omitempty"`
}

func (x *AnimeMovieDubDataRequest) Reset() {
	*x = AnimeMovieDubDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ampb_servers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnimeMovieDubDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnimeMovieDubDataRequest) ProtoMessage() {}

func (x *AnimeMovieDubDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ampb_servers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnimeMovieDubDataRequest.ProtoReflect.Descriptor instead.
func (*AnimeMovieDubDataRequest) Descriptor() ([]byte, []int) {
	return file_ampb_servers_proto_rawDescGZIP(), []int{1}
}

func (x *AnimeMovieDubDataRequest) GetVideos() []*AnimeMovieVideoRequest {
	if x != nil {
		return x.Videos
	}
	return nil
}

func (x *AnimeMovieDubDataRequest) GetTorrents() []*AnimeMovieTorrentRequest {
	if x != nil {
		return x.Torrents
	}
	return nil
}

type AnimeMovieSubDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Videos   []*AnimeMovieVideoResponse   `protobuf:"bytes,1,rep,name=videos,proto3" json:"videos,omitempty"`
	Torrents []*AnimeMovieTorrentResponse `protobuf:"bytes,2,rep,name=torrents,proto3" json:"torrents,omitempty"`
}

func (x *AnimeMovieSubDataResponse) Reset() {
	*x = AnimeMovieSubDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ampb_servers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnimeMovieSubDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnimeMovieSubDataResponse) ProtoMessage() {}

func (x *AnimeMovieSubDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ampb_servers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnimeMovieSubDataResponse.ProtoReflect.Descriptor instead.
func (*AnimeMovieSubDataResponse) Descriptor() ([]byte, []int) {
	return file_ampb_servers_proto_rawDescGZIP(), []int{2}
}

func (x *AnimeMovieSubDataResponse) GetVideos() []*AnimeMovieVideoResponse {
	if x != nil {
		return x.Videos
	}
	return nil
}

func (x *AnimeMovieSubDataResponse) GetTorrents() []*AnimeMovieTorrentResponse {
	if x != nil {
		return x.Torrents
	}
	return nil
}

type AnimeMovieDubDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Videos   []*AnimeMovieVideoResponse   `protobuf:"bytes,1,rep,name=videos,proto3" json:"videos,omitempty"`
	Torrents []*AnimeMovieTorrentResponse `protobuf:"bytes,2,rep,name=torrents,proto3" json:"torrents,omitempty"`
}

func (x *AnimeMovieDubDataResponse) Reset() {
	*x = AnimeMovieDubDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ampb_servers_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnimeMovieDubDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnimeMovieDubDataResponse) ProtoMessage() {}

func (x *AnimeMovieDubDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ampb_servers_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnimeMovieDubDataResponse.ProtoReflect.Descriptor instead.
func (*AnimeMovieDubDataResponse) Descriptor() ([]byte, []int) {
	return file_ampb_servers_proto_rawDescGZIP(), []int{3}
}

func (x *AnimeMovieDubDataResponse) GetVideos() []*AnimeMovieVideoResponse {
	if x != nil {
		return x.Videos
	}
	return nil
}

func (x *AnimeMovieDubDataResponse) GetTorrents() []*AnimeMovieTorrentResponse {
	if x != nil {
		return x.Torrents
	}
	return nil
}

var File_ampb_servers_proto protoreflect.FileDescriptor

var file_ampb_servers_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x6d, 0x70, 0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x6d, 0x70, 0x62, 0x1a, 0x10, 0x61, 0x6d, 0x70, 0x62,
	0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x61, 0x6d,
	0x70, 0x62, 0x2f, 0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x8c, 0x01, 0x0a, 0x18, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x53,
	0x75, 0x62, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a,
	0x06, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x61, 0x6d, 0x70, 0x62, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x06, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x73, 0x12, 0x3a, 0x0a, 0x08, 0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x6d, 0x70, 0x62, 0x2e, 0x41, 0x6e, 0x69,
	0x6d, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x08, 0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x73, 0x22,
	0x8c, 0x01, 0x0a, 0x18, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44, 0x75,
	0x62, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x06,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61,
	0x6d, 0x70, 0x62, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x06, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x73, 0x12, 0x3a, 0x0a, 0x08, 0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x6d, 0x70, 0x62, 0x2e, 0x41, 0x6e, 0x69, 0x6d,
	0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x08, 0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x8f,
	0x01, 0x0a, 0x19, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x53, 0x75, 0x62,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x06,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61,
	0x6d, 0x70, 0x62, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x73, 0x12, 0x3b, 0x0a, 0x08, 0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x6d, 0x70, 0x62, 0x2e, 0x41, 0x6e, 0x69,
	0x6d, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x73,
	0x22, 0x8f, 0x01, 0x0a, 0x19, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44,
	0x75, 0x62, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35,
	0x0a, 0x06, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x61, 0x6d, 0x70, 0x62, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x73, 0x12, 0x3b, 0x0a, 0x08, 0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x6d, 0x70, 0x62, 0x2e, 0x41,
	0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x73, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x64, 0x6a, 0x2d, 0x79, 0x61, 0x63, 0x69, 0x6e, 0x65, 0x2d, 0x66, 0x6c, 0x75, 0x74, 0x74,
	0x65, 0x72, 0x2f, 0x67, 0x6f, 0x6a, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x61, 0x6d, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ampb_servers_proto_rawDescOnce sync.Once
	file_ampb_servers_proto_rawDescData = file_ampb_servers_proto_rawDesc
)

func file_ampb_servers_proto_rawDescGZIP() []byte {
	file_ampb_servers_proto_rawDescOnce.Do(func() {
		file_ampb_servers_proto_rawDescData = protoimpl.X.CompressGZIP(file_ampb_servers_proto_rawDescData)
	})
	return file_ampb_servers_proto_rawDescData
}

var file_ampb_servers_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_ampb_servers_proto_goTypes = []interface{}{
	(*AnimeMovieSubDataRequest)(nil),  // 0: ampb.AnimeMovieSubDataRequest
	(*AnimeMovieDubDataRequest)(nil),  // 1: ampb.AnimeMovieDubDataRequest
	(*AnimeMovieSubDataResponse)(nil), // 2: ampb.AnimeMovieSubDataResponse
	(*AnimeMovieDubDataResponse)(nil), // 3: ampb.AnimeMovieDubDataResponse
	(*AnimeMovieVideoRequest)(nil),    // 4: ampb.AnimeMovieVideoRequest
	(*AnimeMovieTorrentRequest)(nil),  // 5: ampb.AnimeMovieTorrentRequest
	(*AnimeMovieVideoResponse)(nil),   // 6: ampb.AnimeMovieVideoResponse
	(*AnimeMovieTorrentResponse)(nil), // 7: ampb.AnimeMovieTorrentResponse
}
var file_ampb_servers_proto_depIdxs = []int32{
	4, // 0: ampb.AnimeMovieSubDataRequest.videos:type_name -> ampb.AnimeMovieVideoRequest
	5, // 1: ampb.AnimeMovieSubDataRequest.torrents:type_name -> ampb.AnimeMovieTorrentRequest
	4, // 2: ampb.AnimeMovieDubDataRequest.videos:type_name -> ampb.AnimeMovieVideoRequest
	5, // 3: ampb.AnimeMovieDubDataRequest.torrents:type_name -> ampb.AnimeMovieTorrentRequest
	6, // 4: ampb.AnimeMovieSubDataResponse.videos:type_name -> ampb.AnimeMovieVideoResponse
	7, // 5: ampb.AnimeMovieSubDataResponse.torrents:type_name -> ampb.AnimeMovieTorrentResponse
	6, // 6: ampb.AnimeMovieDubDataResponse.videos:type_name -> ampb.AnimeMovieVideoResponse
	7, // 7: ampb.AnimeMovieDubDataResponse.torrents:type_name -> ampb.AnimeMovieTorrentResponse
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_ampb_servers_proto_init() }
func file_ampb_servers_proto_init() {
	if File_ampb_servers_proto != nil {
		return
	}
	file_ampb_video_proto_init()
	file_ampb_torrent_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ampb_servers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnimeMovieSubDataRequest); i {
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
		file_ampb_servers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnimeMovieDubDataRequest); i {
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
		file_ampb_servers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnimeMovieSubDataResponse); i {
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
		file_ampb_servers_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnimeMovieDubDataResponse); i {
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
			RawDescriptor: file_ampb_servers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ampb_servers_proto_goTypes,
		DependencyIndexes: file_ampb_servers_proto_depIdxs,
		MessageInfos:      file_ampb_servers_proto_msgTypes,
	}.Build()
	File_ampb_servers_proto = out.File
	file_ampb_servers_proto_rawDesc = nil
	file_ampb_servers_proto_goTypes = nil
	file_ampb_servers_proto_depIdxs = nil
}
