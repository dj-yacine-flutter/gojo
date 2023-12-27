// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: v1/ampb/title.proto

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

type AnimeMovieTitle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        int64                  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	TitleText string                 `protobuf:"bytes,2,opt,name=titleText,proto3" json:"titleText,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *AnimeMovieTitle) Reset() {
	*x = AnimeMovieTitle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_ampb_title_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnimeMovieTitle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnimeMovieTitle) ProtoMessage() {}

func (x *AnimeMovieTitle) ProtoReflect() protoreflect.Message {
	mi := &file_v1_ampb_title_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnimeMovieTitle.ProtoReflect.Descriptor instead.
func (*AnimeMovieTitle) Descriptor() ([]byte, []int) {
	return file_v1_ampb_title_proto_rawDescGZIP(), []int{0}
}

func (x *AnimeMovieTitle) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *AnimeMovieTitle) GetTitleText() string {
	if x != nil {
		return x.TitleText
	}
	return ""
}

func (x *AnimeMovieTitle) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type AnimeMovieTitleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Official []string `protobuf:"bytes,1,rep,name=official,proto3" json:"official,omitempty"`
	Short    []string `protobuf:"bytes,2,rep,name=short,proto3" json:"short,omitempty"`
	Other    []string `protobuf:"bytes,3,rep,name=other,proto3" json:"other,omitempty"`
}

func (x *AnimeMovieTitleRequest) Reset() {
	*x = AnimeMovieTitleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_ampb_title_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnimeMovieTitleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnimeMovieTitleRequest) ProtoMessage() {}

func (x *AnimeMovieTitleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_ampb_title_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnimeMovieTitleRequest.ProtoReflect.Descriptor instead.
func (*AnimeMovieTitleRequest) Descriptor() ([]byte, []int) {
	return file_v1_ampb_title_proto_rawDescGZIP(), []int{1}
}

func (x *AnimeMovieTitleRequest) GetOfficial() []string {
	if x != nil {
		return x.Official
	}
	return nil
}

func (x *AnimeMovieTitleRequest) GetShort() []string {
	if x != nil {
		return x.Short
	}
	return nil
}

func (x *AnimeMovieTitleRequest) GetOther() []string {
	if x != nil {
		return x.Other
	}
	return nil
}

type AnimeMovieTitleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Official []*AnimeMovieTitle `protobuf:"bytes,1,rep,name=official,proto3" json:"official,omitempty"`
	Short    []*AnimeMovieTitle `protobuf:"bytes,2,rep,name=short,proto3" json:"short,omitempty"`
	Other    []*AnimeMovieTitle `protobuf:"bytes,3,rep,name=other,proto3" json:"other,omitempty"`
}

func (x *AnimeMovieTitleResponse) Reset() {
	*x = AnimeMovieTitleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_ampb_title_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnimeMovieTitleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnimeMovieTitleResponse) ProtoMessage() {}

func (x *AnimeMovieTitleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_ampb_title_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnimeMovieTitleResponse.ProtoReflect.Descriptor instead.
func (*AnimeMovieTitleResponse) Descriptor() ([]byte, []int) {
	return file_v1_ampb_title_proto_rawDescGZIP(), []int{2}
}

func (x *AnimeMovieTitleResponse) GetOfficial() []*AnimeMovieTitle {
	if x != nil {
		return x.Official
	}
	return nil
}

func (x *AnimeMovieTitleResponse) GetShort() []*AnimeMovieTitle {
	if x != nil {
		return x.Short
	}
	return nil
}

func (x *AnimeMovieTitleResponse) GetOther() []*AnimeMovieTitle {
	if x != nil {
		return x.Other
	}
	return nil
}

var File_v1_ampb_title_proto protoreflect.FileDescriptor

var file_v1_ampb_title_proto_rawDesc = []byte{
	0x0a, 0x13, 0x76, 0x31, 0x2f, 0x61, 0x6d, 0x70, 0x62, 0x2f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x76, 0x31, 0x2e, 0x61, 0x6d, 0x70, 0x62, 0x76, 0x31,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x79, 0x0a, 0x0f, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x54, 0x65, 0x78,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x54, 0x65,
	0x78, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x60, 0x0a, 0x16,
	0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x69,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x69,
	0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x05, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x74, 0x68, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x22, 0xb5,
	0x01, 0x0a, 0x17, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x08, 0x6f, 0x66,
	0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x76,
	0x31, 0x2e, 0x61, 0x6d, 0x70, 0x62, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x6f,
	0x76, 0x69, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x52, 0x08, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x69,
	0x61, 0x6c, 0x12, 0x30, 0x0a, 0x05, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x6d, 0x70, 0x62, 0x76, 0x31, 0x2e, 0x41, 0x6e,
	0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x52, 0x05, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x12, 0x30, 0x0a, 0x05, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x6d, 0x70, 0x62, 0x76, 0x31, 0x2e,
	0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x52,
	0x05, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6a, 0x2d, 0x79, 0x61, 0x63, 0x69, 0x6e, 0x65, 0x2d, 0x66,
	0x6c, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x6f, 0x6a, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x6d, 0x70, 0x62, 0x3b, 0x61, 0x6d, 0x70, 0x62, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_ampb_title_proto_rawDescOnce sync.Once
	file_v1_ampb_title_proto_rawDescData = file_v1_ampb_title_proto_rawDesc
)

func file_v1_ampb_title_proto_rawDescGZIP() []byte {
	file_v1_ampb_title_proto_rawDescOnce.Do(func() {
		file_v1_ampb_title_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_ampb_title_proto_rawDescData)
	})
	return file_v1_ampb_title_proto_rawDescData
}

var file_v1_ampb_title_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_v1_ampb_title_proto_goTypes = []interface{}{
	(*AnimeMovieTitle)(nil),         // 0: v1.ampbv1.AnimeMovieTitle
	(*AnimeMovieTitleRequest)(nil),  // 1: v1.ampbv1.AnimeMovieTitleRequest
	(*AnimeMovieTitleResponse)(nil), // 2: v1.ampbv1.AnimeMovieTitleResponse
	(*timestamppb.Timestamp)(nil),   // 3: google.protobuf.Timestamp
}
var file_v1_ampb_title_proto_depIdxs = []int32{
	3, // 0: v1.ampbv1.AnimeMovieTitle.createdAt:type_name -> google.protobuf.Timestamp
	0, // 1: v1.ampbv1.AnimeMovieTitleResponse.official:type_name -> v1.ampbv1.AnimeMovieTitle
	0, // 2: v1.ampbv1.AnimeMovieTitleResponse.short:type_name -> v1.ampbv1.AnimeMovieTitle
	0, // 3: v1.ampbv1.AnimeMovieTitleResponse.other:type_name -> v1.ampbv1.AnimeMovieTitle
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_v1_ampb_title_proto_init() }
func file_v1_ampb_title_proto_init() {
	if File_v1_ampb_title_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_ampb_title_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnimeMovieTitle); i {
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
		file_v1_ampb_title_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnimeMovieTitleRequest); i {
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
		file_v1_ampb_title_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnimeMovieTitleResponse); i {
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
			RawDescriptor: file_v1_ampb_title_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_ampb_title_proto_goTypes,
		DependencyIndexes: file_v1_ampb_title_proto_depIdxs,
		MessageInfos:      file_v1_ampb_title_proto_msgTypes,
	}.Build()
	File_v1_ampb_title_proto = out.File
	file_v1_ampb_title_proto_rawDesc = nil
	file_v1_ampb_title_proto_goTypes = nil
	file_v1_ampb_title_proto_depIdxs = nil
}
