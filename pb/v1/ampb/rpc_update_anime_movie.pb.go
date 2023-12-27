// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: v1/ampb/rpc_update_anime_movie.proto

package ampbv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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

type UpdateAnimeMovieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnimeID           int64                  `protobuf:"varint,1,opt,name=animeID,proto3" json:"animeID,omitempty"`
	OriginalTitle     *string                `protobuf:"bytes,2,opt,name=originalTitle,proto3,oneof" json:"originalTitle,omitempty"`
	Aired             *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=aired,proto3,oneof" json:"aired,omitempty"`
	ReleaseYear       *int32                 `protobuf:"varint,4,opt,name=releaseYear,proto3,oneof" json:"releaseYear,omitempty"`
	Rating            *string                `protobuf:"bytes,5,opt,name=rating,proto3,oneof" json:"rating,omitempty"`
	Duration          *durationpb.Duration   `protobuf:"bytes,6,opt,name=duration,proto3,oneof" json:"duration,omitempty"`
	PortraitPoster    *string                `protobuf:"bytes,7,opt,name=portraitPoster,proto3,oneof" json:"portraitPoster,omitempty"`
	PortraitBlurHash  *string                `protobuf:"bytes,8,opt,name=portraitBlurHash,proto3,oneof" json:"portraitBlurHash,omitempty"`
	LandscapePoster   *string                `protobuf:"bytes,9,opt,name=landscapePoster,proto3,oneof" json:"landscapePoster,omitempty"`
	LandscapeBlurHash *string                `protobuf:"bytes,10,opt,name=landscapeBlurHash,proto3,oneof" json:"landscapeBlurHash,omitempty"`
}

func (x *UpdateAnimeMovieRequest) Reset() {
	*x = UpdateAnimeMovieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_ampb_rpc_update_anime_movie_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAnimeMovieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAnimeMovieRequest) ProtoMessage() {}

func (x *UpdateAnimeMovieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_ampb_rpc_update_anime_movie_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAnimeMovieRequest.ProtoReflect.Descriptor instead.
func (*UpdateAnimeMovieRequest) Descriptor() ([]byte, []int) {
	return file_v1_ampb_rpc_update_anime_movie_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateAnimeMovieRequest) GetAnimeID() int64 {
	if x != nil {
		return x.AnimeID
	}
	return 0
}

func (x *UpdateAnimeMovieRequest) GetOriginalTitle() string {
	if x != nil && x.OriginalTitle != nil {
		return *x.OriginalTitle
	}
	return ""
}

func (x *UpdateAnimeMovieRequest) GetAired() *timestamppb.Timestamp {
	if x != nil {
		return x.Aired
	}
	return nil
}

func (x *UpdateAnimeMovieRequest) GetReleaseYear() int32 {
	if x != nil && x.ReleaseYear != nil {
		return *x.ReleaseYear
	}
	return 0
}

func (x *UpdateAnimeMovieRequest) GetRating() string {
	if x != nil && x.Rating != nil {
		return *x.Rating
	}
	return ""
}

func (x *UpdateAnimeMovieRequest) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *UpdateAnimeMovieRequest) GetPortraitPoster() string {
	if x != nil && x.PortraitPoster != nil {
		return *x.PortraitPoster
	}
	return ""
}

func (x *UpdateAnimeMovieRequest) GetPortraitBlurHash() string {
	if x != nil && x.PortraitBlurHash != nil {
		return *x.PortraitBlurHash
	}
	return ""
}

func (x *UpdateAnimeMovieRequest) GetLandscapePoster() string {
	if x != nil && x.LandscapePoster != nil {
		return *x.LandscapePoster
	}
	return ""
}

func (x *UpdateAnimeMovieRequest) GetLandscapeBlurHash() string {
	if x != nil && x.LandscapeBlurHash != nil {
		return *x.LandscapeBlurHash
	}
	return ""
}

type UpdateAnimeMovieResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnimeMovie *AnimeMovieResponse `protobuf:"bytes,1,opt,name=animeMovie,proto3" json:"animeMovie,omitempty"`
}

func (x *UpdateAnimeMovieResponse) Reset() {
	*x = UpdateAnimeMovieResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_ampb_rpc_update_anime_movie_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAnimeMovieResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAnimeMovieResponse) ProtoMessage() {}

func (x *UpdateAnimeMovieResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_ampb_rpc_update_anime_movie_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAnimeMovieResponse.ProtoReflect.Descriptor instead.
func (*UpdateAnimeMovieResponse) Descriptor() ([]byte, []int) {
	return file_v1_ampb_rpc_update_anime_movie_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateAnimeMovieResponse) GetAnimeMovie() *AnimeMovieResponse {
	if x != nil {
		return x.AnimeMovie
	}
	return nil
}

var File_v1_ampb_rpc_update_anime_movie_proto protoreflect.FileDescriptor

var file_v1_ampb_rpc_update_anime_movie_proto_rawDesc = []byte{
	0x0a, 0x24, 0x76, 0x31, 0x2f, 0x61, 0x6d, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x6f, 0x76, 0x69, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x76, 0x31, 0x2e, 0x61, 0x6d, 0x70, 0x62, 0x76,
	0x31, 0x1a, 0x19, 0x76, 0x31, 0x2f, 0x61, 0x6d, 0x70, 0x62, 0x2f, 0x61, 0x6e, 0x69, 0x6d, 0x65,
	0x5f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xeb, 0x04,
	0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x6e, 0x69,
	0x6d, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x6e, 0x69, 0x6d,
	0x65, 0x49, 0x44, 0x12, 0x29, 0x0a, 0x0d, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0d, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x35,
	0x0a, 0x05, 0x61, 0x69, 0x72, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x01, 0x52, 0x05, 0x61, 0x69, 0x72,
	0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x59, 0x65, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x02, 0x52, 0x0b, 0x72, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x59, 0x65, 0x61, 0x72, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06,
	0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x06,
	0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x88, 0x01, 0x01, 0x12, 0x3a, 0x0a, 0x08, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x04, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x0e, 0x70, 0x6f, 0x72, 0x74, 0x72, 0x61, 0x69,
	0x74, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52,
	0x0e, 0x70, 0x6f, 0x72, 0x74, 0x72, 0x61, 0x69, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x88,
	0x01, 0x01, 0x12, 0x2f, 0x0a, 0x10, 0x70, 0x6f, 0x72, 0x74, 0x72, 0x61, 0x69, 0x74, 0x42, 0x6c,
	0x75, 0x72, 0x48, 0x61, 0x73, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x06, 0x52, 0x10,
	0x70, 0x6f, 0x72, 0x74, 0x72, 0x61, 0x69, 0x74, 0x42, 0x6c, 0x75, 0x72, 0x48, 0x61, 0x73, 0x68,
	0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x0f, 0x6c, 0x61, 0x6e, 0x64, 0x73, 0x63, 0x61, 0x70, 0x65,
	0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x07, 0x52, 0x0f,
	0x6c, 0x61, 0x6e, 0x64, 0x73, 0x63, 0x61, 0x70, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x88,
	0x01, 0x01, 0x12, 0x31, 0x0a, 0x11, 0x6c, 0x61, 0x6e, 0x64, 0x73, 0x63, 0x61, 0x70, 0x65, 0x42,
	0x6c, 0x75, 0x72, 0x48, 0x61, 0x73, 0x68, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x08, 0x52,
	0x11, 0x6c, 0x61, 0x6e, 0x64, 0x73, 0x63, 0x61, 0x70, 0x65, 0x42, 0x6c, 0x75, 0x72, 0x48, 0x61,
	0x73, 0x68, 0x88, 0x01, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x61, 0x6c, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x61, 0x69, 0x72, 0x65,
	0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x59, 0x65, 0x61,
	0x72, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x0b, 0x0a, 0x09,
	0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x70, 0x6f,
	0x72, 0x74, 0x72, 0x61, 0x69, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x42, 0x13, 0x0a, 0x11,
	0x5f, 0x70, 0x6f, 0x72, 0x74, 0x72, 0x61, 0x69, 0x74, 0x42, 0x6c, 0x75, 0x72, 0x48, 0x61, 0x73,
	0x68, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x6c, 0x61, 0x6e, 0x64, 0x73, 0x63, 0x61, 0x70, 0x65, 0x50,
	0x6f, 0x73, 0x74, 0x65, 0x72, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x6c, 0x61, 0x6e, 0x64, 0x73, 0x63,
	0x61, 0x70, 0x65, 0x42, 0x6c, 0x75, 0x72, 0x48, 0x61, 0x73, 0x68, 0x22, 0x59, 0x0a, 0x18, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0a, 0x61, 0x6e, 0x69, 0x6d, 0x65,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x76, 0x31,
	0x2e, 0x61, 0x6d, 0x70, 0x62, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x61, 0x6e, 0x69, 0x6d,
	0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6a, 0x2d, 0x79, 0x61, 0x63, 0x69, 0x6e, 0x65, 0x2d, 0x66,
	0x6c, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x6f, 0x6a, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x6d, 0x70, 0x62, 0x3b, 0x61, 0x6d, 0x70, 0x62, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_ampb_rpc_update_anime_movie_proto_rawDescOnce sync.Once
	file_v1_ampb_rpc_update_anime_movie_proto_rawDescData = file_v1_ampb_rpc_update_anime_movie_proto_rawDesc
)

func file_v1_ampb_rpc_update_anime_movie_proto_rawDescGZIP() []byte {
	file_v1_ampb_rpc_update_anime_movie_proto_rawDescOnce.Do(func() {
		file_v1_ampb_rpc_update_anime_movie_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_ampb_rpc_update_anime_movie_proto_rawDescData)
	})
	return file_v1_ampb_rpc_update_anime_movie_proto_rawDescData
}

var file_v1_ampb_rpc_update_anime_movie_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_ampb_rpc_update_anime_movie_proto_goTypes = []interface{}{
	(*UpdateAnimeMovieRequest)(nil),  // 0: v1.ampbv1.UpdateAnimeMovieRequest
	(*UpdateAnimeMovieResponse)(nil), // 1: v1.ampbv1.UpdateAnimeMovieResponse
	(*timestamppb.Timestamp)(nil),    // 2: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),      // 3: google.protobuf.Duration
	(*AnimeMovieResponse)(nil),       // 4: v1.ampbv1.AnimeMovieResponse
}
var file_v1_ampb_rpc_update_anime_movie_proto_depIdxs = []int32{
	2, // 0: v1.ampbv1.UpdateAnimeMovieRequest.aired:type_name -> google.protobuf.Timestamp
	3, // 1: v1.ampbv1.UpdateAnimeMovieRequest.duration:type_name -> google.protobuf.Duration
	4, // 2: v1.ampbv1.UpdateAnimeMovieResponse.animeMovie:type_name -> v1.ampbv1.AnimeMovieResponse
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_v1_ampb_rpc_update_anime_movie_proto_init() }
func file_v1_ampb_rpc_update_anime_movie_proto_init() {
	if File_v1_ampb_rpc_update_anime_movie_proto != nil {
		return
	}
	file_v1_ampb_anime_movie_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_ampb_rpc_update_anime_movie_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAnimeMovieRequest); i {
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
		file_v1_ampb_rpc_update_anime_movie_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAnimeMovieResponse); i {
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
	file_v1_ampb_rpc_update_anime_movie_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_ampb_rpc_update_anime_movie_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_ampb_rpc_update_anime_movie_proto_goTypes,
		DependencyIndexes: file_v1_ampb_rpc_update_anime_movie_proto_depIdxs,
		MessageInfos:      file_v1_ampb_rpc_update_anime_movie_proto_msgTypes,
	}.Build()
	File_v1_ampb_rpc_update_anime_movie_proto = out.File
	file_v1_ampb_rpc_update_anime_movie_proto_rawDesc = nil
	file_v1_ampb_rpc_update_anime_movie_proto_goTypes = nil
	file_v1_ampb_rpc_update_anime_movie_proto_depIdxs = nil
}
