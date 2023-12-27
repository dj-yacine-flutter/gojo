// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: shpb/anime_link.proto

package shpb

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

type AnimeLinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OfficialWebsite *string  `protobuf:"bytes,1,opt,name=officialWebsite,proto3,oneof" json:"officialWebsite,omitempty"`
	WikipediaUrl    *string  `protobuf:"bytes,2,opt,name=wikipediaUrl,proto3,oneof" json:"wikipediaUrl,omitempty"`
	CrunchyrollUrl  *string  `protobuf:"bytes,3,opt,name=crunchyrollUrl,proto3,oneof" json:"crunchyrollUrl,omitempty"`
	SocialMedia     []string `protobuf:"bytes,4,rep,name=socialMedia,proto3" json:"socialMedia,omitempty"`
}

func (x *AnimeLinkRequest) Reset() {
	*x = AnimeLinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shpb_anime_link_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnimeLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnimeLinkRequest) ProtoMessage() {}

func (x *AnimeLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shpb_anime_link_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnimeLinkRequest.ProtoReflect.Descriptor instead.
func (*AnimeLinkRequest) Descriptor() ([]byte, []int) {
	return file_shpb_anime_link_proto_rawDescGZIP(), []int{0}
}

func (x *AnimeLinkRequest) GetOfficialWebsite() string {
	if x != nil && x.OfficialWebsite != nil {
		return *x.OfficialWebsite
	}
	return ""
}

func (x *AnimeLinkRequest) GetWikipediaUrl() string {
	if x != nil && x.WikipediaUrl != nil {
		return *x.WikipediaUrl
	}
	return ""
}

func (x *AnimeLinkRequest) GetCrunchyrollUrl() string {
	if x != nil && x.CrunchyrollUrl != nil {
		return *x.CrunchyrollUrl
	}
	return ""
}

func (x *AnimeLinkRequest) GetSocialMedia() []string {
	if x != nil {
		return x.SocialMedia
	}
	return nil
}

type AnimeLinkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID              int64                  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	OfficialWebsite string                 `protobuf:"bytes,2,opt,name=officialWebsite,proto3" json:"officialWebsite,omitempty"`
	WikipediaUrl    string                 `protobuf:"bytes,3,opt,name=wikipediaUrl,proto3" json:"wikipediaUrl,omitempty"`
	CrunchyrollUrl  string                 `protobuf:"bytes,4,opt,name=crunchyrollUrl,proto3" json:"crunchyrollUrl,omitempty"`
	SocialMedia     []string               `protobuf:"bytes,5,rep,name=socialMedia,proto3" json:"socialMedia,omitempty"`
	CreatedAt       *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *AnimeLinkResponse) Reset() {
	*x = AnimeLinkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shpb_anime_link_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnimeLinkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnimeLinkResponse) ProtoMessage() {}

func (x *AnimeLinkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shpb_anime_link_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnimeLinkResponse.ProtoReflect.Descriptor instead.
func (*AnimeLinkResponse) Descriptor() ([]byte, []int) {
	return file_shpb_anime_link_proto_rawDescGZIP(), []int{1}
}

func (x *AnimeLinkResponse) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *AnimeLinkResponse) GetOfficialWebsite() string {
	if x != nil {
		return x.OfficialWebsite
	}
	return ""
}

func (x *AnimeLinkResponse) GetWikipediaUrl() string {
	if x != nil {
		return x.WikipediaUrl
	}
	return ""
}

func (x *AnimeLinkResponse) GetCrunchyrollUrl() string {
	if x != nil {
		return x.CrunchyrollUrl
	}
	return ""
}

func (x *AnimeLinkResponse) GetSocialMedia() []string {
	if x != nil {
		return x.SocialMedia
	}
	return nil
}

func (x *AnimeLinkResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_shpb_anime_link_proto protoreflect.FileDescriptor

var file_shpb_anime_link_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x68, 0x70, 0x62, 0x2f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x6c, 0x69, 0x6e,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x73, 0x68, 0x70, 0x62, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf1,
	0x01, 0x0a, 0x10, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0f, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x57,
	0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0f,
	0x6f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x27, 0x0a, 0x0c, 0x77, 0x69, 0x6b, 0x69, 0x70, 0x65, 0x64, 0x69, 0x61, 0x55,
	0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0c, 0x77, 0x69, 0x6b, 0x69,
	0x70, 0x65, 0x64, 0x69, 0x61, 0x55, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x0e, 0x63,
	0x72, 0x75, 0x6e, 0x63, 0x68, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x55, 0x72, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0e, 0x63, 0x72, 0x75, 0x6e, 0x63, 0x68, 0x79, 0x72, 0x6f,
	0x6c, 0x6c, 0x55, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x6f, 0x63, 0x69,
	0x61, 0x6c, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x73,
	0x6f, 0x63, 0x69, 0x61, 0x6c, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x6f,
	0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x42, 0x0f,
	0x0a, 0x0d, 0x5f, 0x77, 0x69, 0x6b, 0x69, 0x70, 0x65, 0x64, 0x69, 0x61, 0x55, 0x72, 0x6c, 0x42,
	0x11, 0x0a, 0x0f, 0x5f, 0x63, 0x72, 0x75, 0x6e, 0x63, 0x68, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x55,
	0x72, 0x6c, 0x22, 0xf5, 0x01, 0x0a, 0x11, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4c, 0x69, 0x6e, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x28, 0x0a, 0x0f, 0x6f, 0x66, 0x66, 0x69,
	0x63, 0x69, 0x61, 0x6c, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x57, 0x65, 0x62, 0x73, 0x69,
	0x74, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x77, 0x69, 0x6b, 0x69, 0x70, 0x65, 0x64, 0x69, 0x61, 0x55,
	0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x77, 0x69, 0x6b, 0x69, 0x70, 0x65,
	0x64, 0x69, 0x61, 0x55, 0x72, 0x6c, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x72, 0x75, 0x6e, 0x63, 0x68,
	0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x55, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x63, 0x72, 0x75, 0x6e, 0x63, 0x68, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x55, 0x72, 0x6c, 0x12, 0x20,
	0x0a, 0x0b, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x4d, 0x65, 0x64, 0x69, 0x61,
	0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6a, 0x2d, 0x79, 0x61, 0x63, 0x69,
	0x6e, 0x65, 0x2d, 0x66, 0x6c, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x6f, 0x6a, 0x6f, 0x2f,
	0x70, 0x62, 0x2f, 0x73, 0x68, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shpb_anime_link_proto_rawDescOnce sync.Once
	file_shpb_anime_link_proto_rawDescData = file_shpb_anime_link_proto_rawDesc
)

func file_shpb_anime_link_proto_rawDescGZIP() []byte {
	file_shpb_anime_link_proto_rawDescOnce.Do(func() {
		file_shpb_anime_link_proto_rawDescData = protoimpl.X.CompressGZIP(file_shpb_anime_link_proto_rawDescData)
	})
	return file_shpb_anime_link_proto_rawDescData
}

var file_shpb_anime_link_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_shpb_anime_link_proto_goTypes = []interface{}{
	(*AnimeLinkRequest)(nil),      // 0: shpb.AnimeLinkRequest
	(*AnimeLinkResponse)(nil),     // 1: shpb.AnimeLinkResponse
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_shpb_anime_link_proto_depIdxs = []int32{
	2, // 0: shpb.AnimeLinkResponse.createdAt:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_shpb_anime_link_proto_init() }
func file_shpb_anime_link_proto_init() {
	if File_shpb_anime_link_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shpb_anime_link_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnimeLinkRequest); i {
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
		file_shpb_anime_link_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnimeLinkResponse); i {
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
	file_shpb_anime_link_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_shpb_anime_link_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shpb_anime_link_proto_goTypes,
		DependencyIndexes: file_shpb_anime_link_proto_depIdxs,
		MessageInfos:      file_shpb_anime_link_proto_msgTypes,
	}.Build()
	File_shpb_anime_link_proto = out.File
	file_shpb_anime_link_proto_rawDesc = nil
	file_shpb_anime_link_proto_goTypes = nil
	file_shpb_anime_link_proto_depIdxs = nil
}
