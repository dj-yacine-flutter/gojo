// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: v1/apb/anime_resource.proto

package apbv1

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

type AnimeResourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TvdbID        *int32  `protobuf:"varint,1,opt,name=tvdbID,proto3,oneof" json:"tvdbID,omitempty"`
	TmdbID        *int32  `protobuf:"varint,2,opt,name=tmdbID,proto3,oneof" json:"tmdbID,omitempty"`
	ImdbID        *string `protobuf:"bytes,3,opt,name=imdbID,proto3,oneof" json:"imdbID,omitempty"`
	LivechartID   *int32  `protobuf:"varint,4,opt,name=livechartID,proto3,oneof" json:"livechartID,omitempty"`
	AnimePlanetID *string `protobuf:"bytes,5,opt,name=animePlanetID,proto3,oneof" json:"animePlanetID,omitempty"`
	AnisearchID   *int32  `protobuf:"varint,6,opt,name=anisearchID,proto3,oneof" json:"anisearchID,omitempty"`
	AnidbID       *int32  `protobuf:"varint,7,opt,name=anidbID,proto3,oneof" json:"anidbID,omitempty"`
	KitsuID       *int32  `protobuf:"varint,8,opt,name=kitsuID,proto3,oneof" json:"kitsuID,omitempty"`
	MalID         *int32  `protobuf:"varint,9,opt,name=malID,proto3,oneof" json:"malID,omitempty"`
	NotifyMoeID   *string `protobuf:"bytes,10,opt,name=notifyMoeID,proto3,oneof" json:"notifyMoeID,omitempty"`
	AnilistID     *int32  `protobuf:"varint,11,opt,name=anilistID,proto3,oneof" json:"anilistID,omitempty"`
}

func (x *AnimeResourceRequest) Reset() {
	*x = AnimeResourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_apb_anime_resource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnimeResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnimeResourceRequest) ProtoMessage() {}

func (x *AnimeResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_apb_anime_resource_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnimeResourceRequest.ProtoReflect.Descriptor instead.
func (*AnimeResourceRequest) Descriptor() ([]byte, []int) {
	return file_v1_apb_anime_resource_proto_rawDescGZIP(), []int{0}
}

func (x *AnimeResourceRequest) GetTvdbID() int32 {
	if x != nil && x.TvdbID != nil {
		return *x.TvdbID
	}
	return 0
}

func (x *AnimeResourceRequest) GetTmdbID() int32 {
	if x != nil && x.TmdbID != nil {
		return *x.TmdbID
	}
	return 0
}

func (x *AnimeResourceRequest) GetImdbID() string {
	if x != nil && x.ImdbID != nil {
		return *x.ImdbID
	}
	return ""
}

func (x *AnimeResourceRequest) GetLivechartID() int32 {
	if x != nil && x.LivechartID != nil {
		return *x.LivechartID
	}
	return 0
}

func (x *AnimeResourceRequest) GetAnimePlanetID() string {
	if x != nil && x.AnimePlanetID != nil {
		return *x.AnimePlanetID
	}
	return ""
}

func (x *AnimeResourceRequest) GetAnisearchID() int32 {
	if x != nil && x.AnisearchID != nil {
		return *x.AnisearchID
	}
	return 0
}

func (x *AnimeResourceRequest) GetAnidbID() int32 {
	if x != nil && x.AnidbID != nil {
		return *x.AnidbID
	}
	return 0
}

func (x *AnimeResourceRequest) GetKitsuID() int32 {
	if x != nil && x.KitsuID != nil {
		return *x.KitsuID
	}
	return 0
}

func (x *AnimeResourceRequest) GetMalID() int32 {
	if x != nil && x.MalID != nil {
		return *x.MalID
	}
	return 0
}

func (x *AnimeResourceRequest) GetNotifyMoeID() string {
	if x != nil && x.NotifyMoeID != nil {
		return *x.NotifyMoeID
	}
	return ""
}

func (x *AnimeResourceRequest) GetAnilistID() int32 {
	if x != nil && x.AnilistID != nil {
		return *x.AnilistID
	}
	return 0
}

type AnimeResourceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID            int64                  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	TvdbID        int32                  `protobuf:"varint,2,opt,name=tvdbID,proto3" json:"tvdbID,omitempty"`
	TmdbID        int32                  `protobuf:"varint,3,opt,name=tmdbID,proto3" json:"tmdbID,omitempty"`
	ImdbID        string                 `protobuf:"bytes,4,opt,name=imdbID,proto3" json:"imdbID,omitempty"`
	LivechartID   int32                  `protobuf:"varint,5,opt,name=livechartID,proto3" json:"livechartID,omitempty"`
	AnimePlanetID string                 `protobuf:"bytes,6,opt,name=animePlanetID,proto3" json:"animePlanetID,omitempty"`
	AnisearchID   int32                  `protobuf:"varint,7,opt,name=anisearchID,proto3" json:"anisearchID,omitempty"`
	AnidbID       int32                  `protobuf:"varint,8,opt,name=anidbID,proto3" json:"anidbID,omitempty"`
	KitsuID       int32                  `protobuf:"varint,9,opt,name=kitsuID,proto3" json:"kitsuID,omitempty"`
	MalID         int32                  `protobuf:"varint,10,opt,name=malID,proto3" json:"malID,omitempty"`
	NotifyMoeID   string                 `protobuf:"bytes,11,opt,name=notifyMoeID,proto3" json:"notifyMoeID,omitempty"`
	AnilistID     int32                  `protobuf:"varint,12,opt,name=anilistID,proto3" json:"anilistID,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *AnimeResourceResponse) Reset() {
	*x = AnimeResourceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_apb_anime_resource_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnimeResourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnimeResourceResponse) ProtoMessage() {}

func (x *AnimeResourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_apb_anime_resource_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnimeResourceResponse.ProtoReflect.Descriptor instead.
func (*AnimeResourceResponse) Descriptor() ([]byte, []int) {
	return file_v1_apb_anime_resource_proto_rawDescGZIP(), []int{1}
}

func (x *AnimeResourceResponse) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *AnimeResourceResponse) GetTvdbID() int32 {
	if x != nil {
		return x.TvdbID
	}
	return 0
}

func (x *AnimeResourceResponse) GetTmdbID() int32 {
	if x != nil {
		return x.TmdbID
	}
	return 0
}

func (x *AnimeResourceResponse) GetImdbID() string {
	if x != nil {
		return x.ImdbID
	}
	return ""
}

func (x *AnimeResourceResponse) GetLivechartID() int32 {
	if x != nil {
		return x.LivechartID
	}
	return 0
}

func (x *AnimeResourceResponse) GetAnimePlanetID() string {
	if x != nil {
		return x.AnimePlanetID
	}
	return ""
}

func (x *AnimeResourceResponse) GetAnisearchID() int32 {
	if x != nil {
		return x.AnisearchID
	}
	return 0
}

func (x *AnimeResourceResponse) GetAnidbID() int32 {
	if x != nil {
		return x.AnidbID
	}
	return 0
}

func (x *AnimeResourceResponse) GetKitsuID() int32 {
	if x != nil {
		return x.KitsuID
	}
	return 0
}

func (x *AnimeResourceResponse) GetMalID() int32 {
	if x != nil {
		return x.MalID
	}
	return 0
}

func (x *AnimeResourceResponse) GetNotifyMoeID() string {
	if x != nil {
		return x.NotifyMoeID
	}
	return ""
}

func (x *AnimeResourceResponse) GetAnilistID() int32 {
	if x != nil {
		return x.AnilistID
	}
	return 0
}

func (x *AnimeResourceResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_v1_apb_anime_resource_proto protoreflect.FileDescriptor

var file_v1_apb_anime_resource_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x62, 0x2f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x76,
	0x31, 0x2e, 0x61, 0x70, 0x62, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x04, 0x0a, 0x14, 0x41, 0x6e, 0x69,
	0x6d, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x06, 0x74, 0x76, 0x64, 0x62, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x00, 0x52, 0x06, 0x74, 0x76, 0x64, 0x62, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x1b,
	0x0a, 0x06, 0x74, 0x6d, 0x64, 0x62, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01,
	0x52, 0x06, 0x74, 0x6d, 0x64, 0x62, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x69,
	0x6d, 0x64, 0x62, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x06, 0x69,
	0x6d, 0x64, 0x62, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x6c, 0x69, 0x76, 0x65,
	0x63, 0x68, 0x61, 0x72, 0x74, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x03, 0x52,
	0x0b, 0x6c, 0x69, 0x76, 0x65, 0x63, 0x68, 0x61, 0x72, 0x74, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12,
	0x29, 0x0a, 0x0d, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x49, 0x44,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x0d, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x50,
	0x6c, 0x61, 0x6e, 0x65, 0x74, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x61, 0x6e,
	0x69, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x48,
	0x05, 0x52, 0x0b, 0x61, 0x6e, 0x69, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x44, 0x88, 0x01,
	0x01, 0x12, 0x1d, 0x0a, 0x07, 0x61, 0x6e, 0x69, 0x64, 0x62, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x05, 0x48, 0x06, 0x52, 0x07, 0x61, 0x6e, 0x69, 0x64, 0x62, 0x49, 0x44, 0x88, 0x01, 0x01,
	0x12, 0x1d, 0x0a, 0x07, 0x6b, 0x69, 0x74, 0x73, 0x75, 0x49, 0x44, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x07, 0x52, 0x07, 0x6b, 0x69, 0x74, 0x73, 0x75, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12,
	0x19, 0x0a, 0x05, 0x6d, 0x61, 0x6c, 0x49, 0x44, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x48, 0x08,
	0x52, 0x05, 0x6d, 0x61, 0x6c, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x4d, 0x6f, 0x65, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x09, 0x52, 0x0b, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4d, 0x6f, 0x65, 0x49, 0x44, 0x88, 0x01,
	0x01, 0x12, 0x21, 0x0a, 0x09, 0x61, 0x6e, 0x69, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x44, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x05, 0x48, 0x0a, 0x52, 0x09, 0x61, 0x6e, 0x69, 0x6c, 0x69, 0x73, 0x74, 0x49,
	0x44, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x74, 0x76, 0x64, 0x62, 0x49, 0x44, 0x42,
	0x09, 0x0a, 0x07, 0x5f, 0x74, 0x6d, 0x64, 0x62, 0x49, 0x44, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x69,
	0x6d, 0x64, 0x62, 0x49, 0x44, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6c, 0x69, 0x76, 0x65, 0x63, 0x68,
	0x61, 0x72, 0x74, 0x49, 0x44, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x50,
	0x6c, 0x61, 0x6e, 0x65, 0x74, 0x49, 0x44, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x61, 0x6e, 0x69, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x44, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x61, 0x6e, 0x69, 0x64,
	0x62, 0x49, 0x44, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6b, 0x69, 0x74, 0x73, 0x75, 0x49, 0x44, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x6d, 0x61, 0x6c, 0x49, 0x44, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x4d, 0x6f, 0x65, 0x49, 0x44, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x61, 0x6e,
	0x69, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x44, 0x22, 0x9d, 0x03, 0x0a, 0x15, 0x41, 0x6e, 0x69, 0x6d,
	0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49,
	0x44, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x76, 0x64, 0x62, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x74, 0x76, 0x64, 0x62, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6d, 0x64,
	0x62, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x74, 0x6d, 0x64, 0x62, 0x49,
	0x44, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d, 0x64, 0x62, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x69, 0x6d, 0x64, 0x62, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x6c, 0x69, 0x76,
	0x65, 0x63, 0x68, 0x61, 0x72, 0x74, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b,
	0x6c, 0x69, 0x76, 0x65, 0x63, 0x68, 0x61, 0x72, 0x74, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x61,
	0x6e, 0x69, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x49,
	0x44, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x6e, 0x69, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x44,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x61, 0x6e, 0x69, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x6e, 0x69, 0x64, 0x62, 0x49, 0x44, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x61, 0x6e, 0x69, 0x64, 0x62, 0x49, 0x44, 0x12, 0x18, 0x0a,
	0x07, 0x6b, 0x69, 0x74, 0x73, 0x75, 0x49, 0x44, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x6b, 0x69, 0x74, 0x73, 0x75, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x61, 0x6c, 0x49, 0x44,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x61, 0x6c, 0x49, 0x44, 0x12, 0x20, 0x0a,
	0x0b, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4d, 0x6f, 0x65, 0x49, 0x44, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4d, 0x6f, 0x65, 0x49, 0x44, 0x12,
	0x1c, 0x0a, 0x09, 0x61, 0x6e, 0x69, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x44, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x61, 0x6e, 0x69, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x44, 0x12, 0x38, 0x0a,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6a, 0x2d, 0x79, 0x61, 0x63, 0x69, 0x6e, 0x65, 0x2d,
	0x66, 0x6c, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x6f, 0x6a, 0x6f, 0x2f, 0x70, 0x62, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x70, 0x62, 0x3b, 0x61, 0x70, 0x62, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_apb_anime_resource_proto_rawDescOnce sync.Once
	file_v1_apb_anime_resource_proto_rawDescData = file_v1_apb_anime_resource_proto_rawDesc
)

func file_v1_apb_anime_resource_proto_rawDescGZIP() []byte {
	file_v1_apb_anime_resource_proto_rawDescOnce.Do(func() {
		file_v1_apb_anime_resource_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_apb_anime_resource_proto_rawDescData)
	})
	return file_v1_apb_anime_resource_proto_rawDescData
}

var file_v1_apb_anime_resource_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_apb_anime_resource_proto_goTypes = []interface{}{
	(*AnimeResourceRequest)(nil),  // 0: v1.apbv1.AnimeResourceRequest
	(*AnimeResourceResponse)(nil), // 1: v1.apbv1.AnimeResourceResponse
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_v1_apb_anime_resource_proto_depIdxs = []int32{
	2, // 0: v1.apbv1.AnimeResourceResponse.createdAt:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v1_apb_anime_resource_proto_init() }
func file_v1_apb_anime_resource_proto_init() {
	if File_v1_apb_anime_resource_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_apb_anime_resource_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnimeResourceRequest); i {
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
		file_v1_apb_anime_resource_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnimeResourceResponse); i {
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
	file_v1_apb_anime_resource_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_apb_anime_resource_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_apb_anime_resource_proto_goTypes,
		DependencyIndexes: file_v1_apb_anime_resource_proto_depIdxs,
		MessageInfos:      file_v1_apb_anime_resource_proto_msgTypes,
	}.Build()
	File_v1_apb_anime_resource_proto = out.File
	file_v1_apb_anime_resource_proto_rawDesc = nil
	file_v1_apb_anime_resource_proto_goTypes = nil
	file_v1_apb_anime_resource_proto_depIdxs = nil
}
