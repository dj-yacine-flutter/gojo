// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: aspb/title.proto

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

type AnimeSeasonTitle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        int64                  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	TitleText string                 `protobuf:"bytes,2,opt,name=titleText,proto3" json:"titleText,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *AnimeSeasonTitle) Reset() {
	*x = AnimeSeasonTitle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aspb_title_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnimeSeasonTitle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnimeSeasonTitle) ProtoMessage() {}

func (x *AnimeSeasonTitle) ProtoReflect() protoreflect.Message {
	mi := &file_aspb_title_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnimeSeasonTitle.ProtoReflect.Descriptor instead.
func (*AnimeSeasonTitle) Descriptor() ([]byte, []int) {
	return file_aspb_title_proto_rawDescGZIP(), []int{0}
}

func (x *AnimeSeasonTitle) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *AnimeSeasonTitle) GetTitleText() string {
	if x != nil {
		return x.TitleText
	}
	return ""
}

func (x *AnimeSeasonTitle) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type AnimeSeasonTitleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Official []string `protobuf:"bytes,1,rep,name=official,proto3" json:"official,omitempty"`
	Short    []string `protobuf:"bytes,2,rep,name=short,proto3" json:"short,omitempty"`
	Other    []string `protobuf:"bytes,3,rep,name=other,proto3" json:"other,omitempty"`
}

func (x *AnimeSeasonTitleRequest) Reset() {
	*x = AnimeSeasonTitleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aspb_title_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnimeSeasonTitleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnimeSeasonTitleRequest) ProtoMessage() {}

func (x *AnimeSeasonTitleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aspb_title_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnimeSeasonTitleRequest.ProtoReflect.Descriptor instead.
func (*AnimeSeasonTitleRequest) Descriptor() ([]byte, []int) {
	return file_aspb_title_proto_rawDescGZIP(), []int{1}
}

func (x *AnimeSeasonTitleRequest) GetOfficial() []string {
	if x != nil {
		return x.Official
	}
	return nil
}

func (x *AnimeSeasonTitleRequest) GetShort() []string {
	if x != nil {
		return x.Short
	}
	return nil
}

func (x *AnimeSeasonTitleRequest) GetOther() []string {
	if x != nil {
		return x.Other
	}
	return nil
}

type AnimeSeasonTitleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Official []*AnimeSeasonTitle `protobuf:"bytes,1,rep,name=official,proto3" json:"official,omitempty"`
	Short    []*AnimeSeasonTitle `protobuf:"bytes,2,rep,name=short,proto3" json:"short,omitempty"`
	Other    []*AnimeSeasonTitle `protobuf:"bytes,3,rep,name=other,proto3" json:"other,omitempty"`
}

func (x *AnimeSeasonTitleResponse) Reset() {
	*x = AnimeSeasonTitleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aspb_title_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnimeSeasonTitleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnimeSeasonTitleResponse) ProtoMessage() {}

func (x *AnimeSeasonTitleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aspb_title_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnimeSeasonTitleResponse.ProtoReflect.Descriptor instead.
func (*AnimeSeasonTitleResponse) Descriptor() ([]byte, []int) {
	return file_aspb_title_proto_rawDescGZIP(), []int{2}
}

func (x *AnimeSeasonTitleResponse) GetOfficial() []*AnimeSeasonTitle {
	if x != nil {
		return x.Official
	}
	return nil
}

func (x *AnimeSeasonTitleResponse) GetShort() []*AnimeSeasonTitle {
	if x != nil {
		return x.Short
	}
	return nil
}

func (x *AnimeSeasonTitleResponse) GetOther() []*AnimeSeasonTitle {
	if x != nil {
		return x.Other
	}
	return nil
}

var File_aspb_title_proto protoreflect.FileDescriptor

var file_aspb_title_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x73, 0x70, 0x62, 0x2f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x61, 0x73, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7a, 0x0a, 0x10, 0x41, 0x6e, 0x69,
	0x6d, 0x65, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x54, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x54, 0x65, 0x78, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x61, 0x0a, 0x17, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x53, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x08, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x05, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x22, 0xaa, 0x01, 0x0a, 0x18, 0x41, 0x6e, 0x69,
	0x6d, 0x65, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x08, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x73, 0x70, 0x62, 0x2e, 0x41,
	0x6e, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x52,
	0x08, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x12, 0x2c, 0x0a, 0x05, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x73, 0x70, 0x62, 0x2e,
	0x41, 0x6e, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x52, 0x05, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x12, 0x2c, 0x0a, 0x05, 0x6f, 0x74, 0x68, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x73, 0x70, 0x62, 0x2e, 0x41, 0x6e,
	0x69, 0x6d, 0x65, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x52, 0x05,
	0x6f, 0x74, 0x68, 0x65, 0x72, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6a, 0x2d, 0x79, 0x61, 0x63, 0x69, 0x6e, 0x65, 0x2d, 0x66, 0x6c,
	0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x6f, 0x6a, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x61, 0x73,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aspb_title_proto_rawDescOnce sync.Once
	file_aspb_title_proto_rawDescData = file_aspb_title_proto_rawDesc
)

func file_aspb_title_proto_rawDescGZIP() []byte {
	file_aspb_title_proto_rawDescOnce.Do(func() {
		file_aspb_title_proto_rawDescData = protoimpl.X.CompressGZIP(file_aspb_title_proto_rawDescData)
	})
	return file_aspb_title_proto_rawDescData
}

var file_aspb_title_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_aspb_title_proto_goTypes = []interface{}{
	(*AnimeSeasonTitle)(nil),         // 0: aspb.AnimeSeasonTitle
	(*AnimeSeasonTitleRequest)(nil),  // 1: aspb.AnimeSeasonTitleRequest
	(*AnimeSeasonTitleResponse)(nil), // 2: aspb.AnimeSeasonTitleResponse
	(*timestamppb.Timestamp)(nil),    // 3: google.protobuf.Timestamp
}
var file_aspb_title_proto_depIdxs = []int32{
	3, // 0: aspb.AnimeSeasonTitle.createdAt:type_name -> google.protobuf.Timestamp
	0, // 1: aspb.AnimeSeasonTitleResponse.official:type_name -> aspb.AnimeSeasonTitle
	0, // 2: aspb.AnimeSeasonTitleResponse.short:type_name -> aspb.AnimeSeasonTitle
	0, // 3: aspb.AnimeSeasonTitleResponse.other:type_name -> aspb.AnimeSeasonTitle
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_aspb_title_proto_init() }
func file_aspb_title_proto_init() {
	if File_aspb_title_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aspb_title_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnimeSeasonTitle); i {
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
		file_aspb_title_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnimeSeasonTitleRequest); i {
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
		file_aspb_title_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnimeSeasonTitleResponse); i {
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
			RawDescriptor: file_aspb_title_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aspb_title_proto_goTypes,
		DependencyIndexes: file_aspb_title_proto_depIdxs,
		MessageInfos:      file_aspb_title_proto_msgTypes,
	}.Build()
	File_aspb_title_proto = out.File
	file_aspb_title_proto_rawDesc = nil
	file_aspb_title_proto_goTypes = nil
	file_aspb_title_proto_depIdxs = nil
}
