// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: ampb/anime_movie.proto

package ampb

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

type AnimeMovieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OriginalTitle     string                 `protobuf:"bytes,1,opt,name=originalTitle,proto3" json:"originalTitle,omitempty"`
	Aired             *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=aired,proto3" json:"aired,omitempty"`
	ReleaseYear       int32                  `protobuf:"varint,3,opt,name=releaseYear,proto3" json:"releaseYear,omitempty"`
	Rating            string                 `protobuf:"bytes,4,opt,name=rating,proto3" json:"rating,omitempty"`
	Duration          *durationpb.Duration   `protobuf:"bytes,5,opt,name=duration,proto3" json:"duration,omitempty"`
	PortriatPoster    string                 `protobuf:"bytes,6,opt,name=portriatPoster,proto3" json:"portriatPoster,omitempty"`
	PortriatBlurHash  string                 `protobuf:"bytes,7,opt,name=portriatBlurHash,proto3" json:"portriatBlurHash,omitempty"`
	LandscapePoster   string                 `protobuf:"bytes,8,opt,name=landscapePoster,proto3" json:"landscapePoster,omitempty"`
	LandscapeBlurHash string                 `protobuf:"bytes,9,opt,name=landscapeBlurHash,proto3" json:"landscapeBlurHash,omitempty"`
}

func (x *AnimeMovieRequest) Reset() {
	*x = AnimeMovieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ampb_anime_movie_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnimeMovieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnimeMovieRequest) ProtoMessage() {}

func (x *AnimeMovieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ampb_anime_movie_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnimeMovieRequest.ProtoReflect.Descriptor instead.
func (*AnimeMovieRequest) Descriptor() ([]byte, []int) {
	return file_ampb_anime_movie_proto_rawDescGZIP(), []int{0}
}

func (x *AnimeMovieRequest) GetOriginalTitle() string {
	if x != nil {
		return x.OriginalTitle
	}
	return ""
}

func (x *AnimeMovieRequest) GetAired() *timestamppb.Timestamp {
	if x != nil {
		return x.Aired
	}
	return nil
}

func (x *AnimeMovieRequest) GetReleaseYear() int32 {
	if x != nil {
		return x.ReleaseYear
	}
	return 0
}

func (x *AnimeMovieRequest) GetRating() string {
	if x != nil {
		return x.Rating
	}
	return ""
}

func (x *AnimeMovieRequest) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *AnimeMovieRequest) GetPortriatPoster() string {
	if x != nil {
		return x.PortriatPoster
	}
	return ""
}

func (x *AnimeMovieRequest) GetPortriatBlurHash() string {
	if x != nil {
		return x.PortriatBlurHash
	}
	return ""
}

func (x *AnimeMovieRequest) GetLandscapePoster() string {
	if x != nil {
		return x.LandscapePoster
	}
	return ""
}

func (x *AnimeMovieRequest) GetLandscapeBlurHash() string {
	if x != nil {
		return x.LandscapeBlurHash
	}
	return ""
}

type AnimeMovieResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID                int64                  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	OriginalTitle     string                 `protobuf:"bytes,2,opt,name=originalTitle,proto3" json:"originalTitle,omitempty"`
	Aired             *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=aired,proto3" json:"aired,omitempty"`
	ReleaseYear       int32                  `protobuf:"varint,4,opt,name=releaseYear,proto3" json:"releaseYear,omitempty"`
	Rating            string                 `protobuf:"bytes,5,opt,name=rating,proto3" json:"rating,omitempty"`
	Duration          *durationpb.Duration   `protobuf:"bytes,6,opt,name=duration,proto3" json:"duration,omitempty"`
	PortriatPoster    string                 `protobuf:"bytes,7,opt,name=portriatPoster,proto3" json:"portriatPoster,omitempty"`
	PortriatBlurHash  string                 `protobuf:"bytes,8,opt,name=portriatBlurHash,proto3" json:"portriatBlurHash,omitempty"`
	LandscapePoster   string                 `protobuf:"bytes,9,opt,name=landscapePoster,proto3" json:"landscapePoster,omitempty"`
	LandscapeBlurHash string                 `protobuf:"bytes,10,opt,name=landscapeBlurHash,proto3" json:"landscapeBlurHash,omitempty"`
	CreatedAt         *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *AnimeMovieResponse) Reset() {
	*x = AnimeMovieResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ampb_anime_movie_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnimeMovieResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnimeMovieResponse) ProtoMessage() {}

func (x *AnimeMovieResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ampb_anime_movie_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnimeMovieResponse.ProtoReflect.Descriptor instead.
func (*AnimeMovieResponse) Descriptor() ([]byte, []int) {
	return file_ampb_anime_movie_proto_rawDescGZIP(), []int{1}
}

func (x *AnimeMovieResponse) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *AnimeMovieResponse) GetOriginalTitle() string {
	if x != nil {
		return x.OriginalTitle
	}
	return ""
}

func (x *AnimeMovieResponse) GetAired() *timestamppb.Timestamp {
	if x != nil {
		return x.Aired
	}
	return nil
}

func (x *AnimeMovieResponse) GetReleaseYear() int32 {
	if x != nil {
		return x.ReleaseYear
	}
	return 0
}

func (x *AnimeMovieResponse) GetRating() string {
	if x != nil {
		return x.Rating
	}
	return ""
}

func (x *AnimeMovieResponse) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *AnimeMovieResponse) GetPortriatPoster() string {
	if x != nil {
		return x.PortriatPoster
	}
	return ""
}

func (x *AnimeMovieResponse) GetPortriatBlurHash() string {
	if x != nil {
		return x.PortriatBlurHash
	}
	return ""
}

func (x *AnimeMovieResponse) GetLandscapePoster() string {
	if x != nil {
		return x.LandscapePoster
	}
	return ""
}

func (x *AnimeMovieResponse) GetLandscapeBlurHash() string {
	if x != nil {
		return x.LandscapeBlurHash
	}
	return ""
}

func (x *AnimeMovieResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_ampb_anime_movie_proto protoreflect.FileDescriptor

var file_ampb_anime_movie_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x6d, 0x70, 0x62, 0x2f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x6f, 0x76,
	0x69, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x6d, 0x70, 0x62, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x88, 0x03, 0x0a, 0x11, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61,
	0x6c, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x61,
	0x69, 0x72, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x61, 0x69, 0x72, 0x65, 0x64, 0x12, 0x20, 0x0a,
	0x0b, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x59, 0x65, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x59, 0x65, 0x61, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26,
	0x0a, 0x0e, 0x70, 0x6f, 0x72, 0x74, 0x72, 0x69, 0x61, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x6f, 0x72, 0x74, 0x72, 0x69, 0x61, 0x74,
	0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x10, 0x70, 0x6f, 0x72, 0x74, 0x72, 0x69,
	0x61, 0x74, 0x42, 0x6c, 0x75, 0x72, 0x48, 0x61, 0x73, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x70, 0x6f, 0x72, 0x74, 0x72, 0x69, 0x61, 0x74, 0x42, 0x6c, 0x75, 0x72, 0x48, 0x61,
	0x73, 0x68, 0x12, 0x28, 0x0a, 0x0f, 0x6c, 0x61, 0x6e, 0x64, 0x73, 0x63, 0x61, 0x70, 0x65, 0x50,
	0x6f, 0x73, 0x74, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6c, 0x61, 0x6e,
	0x64, 0x73, 0x63, 0x61, 0x70, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x11,
	0x6c, 0x61, 0x6e, 0x64, 0x73, 0x63, 0x61, 0x70, 0x65, 0x42, 0x6c, 0x75, 0x72, 0x48, 0x61, 0x73,
	0x68, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6c, 0x61, 0x6e, 0x64, 0x73, 0x63, 0x61,
	0x70, 0x65, 0x42, 0x6c, 0x75, 0x72, 0x48, 0x61, 0x73, 0x68, 0x22, 0xd3, 0x03, 0x0a, 0x12, 0x41,
	0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49,
	0x44, 0x12, 0x24, 0x0a, 0x0d, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x61, 0x6c, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x61, 0x69, 0x72, 0x65, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x05, 0x61, 0x69, 0x72, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x59, 0x65, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b,
	0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x59, 0x65, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x6f,
	0x72, 0x74, 0x72, 0x69, 0x61, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x70, 0x6f, 0x72, 0x74, 0x72, 0x69, 0x61, 0x74, 0x50, 0x6f, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x2a, 0x0a, 0x10, 0x70, 0x6f, 0x72, 0x74, 0x72, 0x69, 0x61, 0x74, 0x42, 0x6c,
	0x75, 0x72, 0x48, 0x61, 0x73, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x6f,
	0x72, 0x74, 0x72, 0x69, 0x61, 0x74, 0x42, 0x6c, 0x75, 0x72, 0x48, 0x61, 0x73, 0x68, 0x12, 0x28,
	0x0a, 0x0f, 0x6c, 0x61, 0x6e, 0x64, 0x73, 0x63, 0x61, 0x70, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x65,
	0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6c, 0x61, 0x6e, 0x64, 0x73, 0x63, 0x61,
	0x70, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x11, 0x6c, 0x61, 0x6e, 0x64,
	0x73, 0x63, 0x61, 0x70, 0x65, 0x42, 0x6c, 0x75, 0x72, 0x48, 0x61, 0x73, 0x68, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x11, 0x6c, 0x61, 0x6e, 0x64, 0x73, 0x63, 0x61, 0x70, 0x65, 0x42, 0x6c,
	0x75, 0x72, 0x48, 0x61, 0x73, 0x68, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x6a, 0x2d, 0x79, 0x61, 0x63, 0x69, 0x6e, 0x65, 0x2d, 0x66, 0x6c, 0x75, 0x74, 0x74, 0x65, 0x72,
	0x2f, 0x67, 0x6f, 0x6a, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x61, 0x6d, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ampb_anime_movie_proto_rawDescOnce sync.Once
	file_ampb_anime_movie_proto_rawDescData = file_ampb_anime_movie_proto_rawDesc
)

func file_ampb_anime_movie_proto_rawDescGZIP() []byte {
	file_ampb_anime_movie_proto_rawDescOnce.Do(func() {
		file_ampb_anime_movie_proto_rawDescData = protoimpl.X.CompressGZIP(file_ampb_anime_movie_proto_rawDescData)
	})
	return file_ampb_anime_movie_proto_rawDescData
}

var file_ampb_anime_movie_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ampb_anime_movie_proto_goTypes = []interface{}{
	(*AnimeMovieRequest)(nil),     // 0: ampb.AnimeMovieRequest
	(*AnimeMovieResponse)(nil),    // 1: ampb.AnimeMovieResponse
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),   // 3: google.protobuf.Duration
}
var file_ampb_anime_movie_proto_depIdxs = []int32{
	2, // 0: ampb.AnimeMovieRequest.aired:type_name -> google.protobuf.Timestamp
	3, // 1: ampb.AnimeMovieRequest.duration:type_name -> google.protobuf.Duration
	2, // 2: ampb.AnimeMovieResponse.aired:type_name -> google.protobuf.Timestamp
	3, // 3: ampb.AnimeMovieResponse.duration:type_name -> google.protobuf.Duration
	2, // 4: ampb.AnimeMovieResponse.createdAt:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_ampb_anime_movie_proto_init() }
func file_ampb_anime_movie_proto_init() {
	if File_ampb_anime_movie_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ampb_anime_movie_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnimeMovieRequest); i {
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
		file_ampb_anime_movie_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnimeMovieResponse); i {
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
			RawDescriptor: file_ampb_anime_movie_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ampb_anime_movie_proto_goTypes,
		DependencyIndexes: file_ampb_anime_movie_proto_depIdxs,
		MessageInfos:      file_ampb_anime_movie_proto_msgTypes,
	}.Build()
	File_ampb_anime_movie_proto = out.File
	file_ampb_anime_movie_proto_rawDesc = nil
	file_ampb_anime_movie_proto_goTypes = nil
	file_ampb_anime_movie_proto_depIdxs = nil
}
