// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: aspb/video.proto

package aspb

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

type AnimeSerieVideoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LanguageID int32  `protobuf:"varint,1,opt,name=languageID,proto3" json:"languageID,omitempty"`
	Autority   string `protobuf:"bytes,2,opt,name=autority,proto3" json:"autority,omitempty"`
	Referer    string `protobuf:"bytes,3,opt,name=referer,proto3" json:"referer,omitempty"`
	Link       string `protobuf:"bytes,4,opt,name=link,proto3" json:"link,omitempty"`
	Quality    string `protobuf:"bytes,5,opt,name=quality,proto3" json:"quality,omitempty"`
}

func (x *AnimeSerieVideoRequest) Reset() {
	*x = AnimeSerieVideoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aspb_video_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnimeSerieVideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnimeSerieVideoRequest) ProtoMessage() {}

func (x *AnimeSerieVideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aspb_video_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnimeSerieVideoRequest.ProtoReflect.Descriptor instead.
func (*AnimeSerieVideoRequest) Descriptor() ([]byte, []int) {
	return file_aspb_video_proto_rawDescGZIP(), []int{0}
}

func (x *AnimeSerieVideoRequest) GetLanguageID() int32 {
	if x != nil {
		return x.LanguageID
	}
	return 0
}

func (x *AnimeSerieVideoRequest) GetAutority() string {
	if x != nil {
		return x.Autority
	}
	return ""
}

func (x *AnimeSerieVideoRequest) GetReferer() string {
	if x != nil {
		return x.Referer
	}
	return ""
}

func (x *AnimeSerieVideoRequest) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *AnimeSerieVideoRequest) GetQuality() string {
	if x != nil {
		return x.Quality
	}
	return ""
}

type AnimeSerieVideoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         int64                  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	LanguageID int32                  `protobuf:"varint,2,opt,name=languageID,proto3" json:"languageID,omitempty"`
	Autority   string                 `protobuf:"bytes,3,opt,name=autority,proto3" json:"autority,omitempty"`
	Referer    string                 `protobuf:"bytes,4,opt,name=referer,proto3" json:"referer,omitempty"`
	Link       string                 `protobuf:"bytes,5,opt,name=link,proto3" json:"link,omitempty"`
	Quality    string                 `protobuf:"bytes,6,opt,name=quality,proto3" json:"quality,omitempty"`
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *AnimeSerieVideoResponse) Reset() {
	*x = AnimeSerieVideoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aspb_video_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnimeSerieVideoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnimeSerieVideoResponse) ProtoMessage() {}

func (x *AnimeSerieVideoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aspb_video_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnimeSerieVideoResponse.ProtoReflect.Descriptor instead.
func (*AnimeSerieVideoResponse) Descriptor() ([]byte, []int) {
	return file_aspb_video_proto_rawDescGZIP(), []int{1}
}

func (x *AnimeSerieVideoResponse) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *AnimeSerieVideoResponse) GetLanguageID() int32 {
	if x != nil {
		return x.LanguageID
	}
	return 0
}

func (x *AnimeSerieVideoResponse) GetAutority() string {
	if x != nil {
		return x.Autority
	}
	return ""
}

func (x *AnimeSerieVideoResponse) GetReferer() string {
	if x != nil {
		return x.Referer
	}
	return ""
}

func (x *AnimeSerieVideoResponse) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *AnimeSerieVideoResponse) GetQuality() string {
	if x != nil {
		return x.Quality
	}
	return ""
}

func (x *AnimeSerieVideoResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_aspb_video_proto protoreflect.FileDescriptor

var file_aspb_video_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x73, 0x70, 0x62, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x61, 0x73, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x01, 0x0a, 0x16, 0x41, 0x6e,
	0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69,
	0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x18,
	0x0a, 0x07, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x22, 0xe7, 0x01, 0x0a, 0x17, 0x41, 0x6e, 0x69,
	0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69,
	0x6e, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x18,
	0x0a, 0x07, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x64, 0x6a, 0x2d, 0x79, 0x61, 0x63, 0x69, 0x6e, 0x65, 0x2d, 0x66, 0x6c, 0x75, 0x74, 0x74,
	0x65, 0x72, 0x2f, 0x67, 0x6f, 0x6a, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x61, 0x73, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aspb_video_proto_rawDescOnce sync.Once
	file_aspb_video_proto_rawDescData = file_aspb_video_proto_rawDesc
)

func file_aspb_video_proto_rawDescGZIP() []byte {
	file_aspb_video_proto_rawDescOnce.Do(func() {
		file_aspb_video_proto_rawDescData = protoimpl.X.CompressGZIP(file_aspb_video_proto_rawDescData)
	})
	return file_aspb_video_proto_rawDescData
}

var file_aspb_video_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_aspb_video_proto_goTypes = []interface{}{
	(*AnimeSerieVideoRequest)(nil),  // 0: aspb.AnimeSerieVideoRequest
	(*AnimeSerieVideoResponse)(nil), // 1: aspb.AnimeSerieVideoResponse
	(*timestamppb.Timestamp)(nil),   // 2: google.protobuf.Timestamp
}
var file_aspb_video_proto_depIdxs = []int32{
	2, // 0: aspb.AnimeSerieVideoResponse.createdAt:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_aspb_video_proto_init() }
func file_aspb_video_proto_init() {
	if File_aspb_video_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aspb_video_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnimeSerieVideoRequest); i {
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
		file_aspb_video_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnimeSerieVideoResponse); i {
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
			RawDescriptor: file_aspb_video_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aspb_video_proto_goTypes,
		DependencyIndexes: file_aspb_video_proto_depIdxs,
		MessageInfos:      file_aspb_video_proto_msgTypes,
	}.Build()
	File_aspb_video_proto = out.File
	file_aspb_video_proto_rawDesc = nil
	file_aspb_video_proto_goTypes = nil
	file_aspb_video_proto_depIdxs = nil
}
