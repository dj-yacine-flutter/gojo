// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: aspb/servers.proto

package aspb

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

type AnimeSerieSubVideoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Videos []*AnimeSerieVideoRequest `protobuf:"bytes,1,rep,name=videos,proto3" json:"videos,omitempty"`
}

func (x *AnimeSerieSubVideoRequest) Reset() {
	*x = AnimeSerieSubVideoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aspb_servers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnimeSerieSubVideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnimeSerieSubVideoRequest) ProtoMessage() {}

func (x *AnimeSerieSubVideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aspb_servers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnimeSerieSubVideoRequest.ProtoReflect.Descriptor instead.
func (*AnimeSerieSubVideoRequest) Descriptor() ([]byte, []int) {
	return file_aspb_servers_proto_rawDescGZIP(), []int{0}
}

func (x *AnimeSerieSubVideoRequest) GetVideos() []*AnimeSerieVideoRequest {
	if x != nil {
		return x.Videos
	}
	return nil
}

type AnimeSerieDubVideoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Videos []*AnimeSerieVideoRequest `protobuf:"bytes,1,rep,name=videos,proto3" json:"videos,omitempty"`
}

func (x *AnimeSerieDubVideoRequest) Reset() {
	*x = AnimeSerieDubVideoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aspb_servers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnimeSerieDubVideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnimeSerieDubVideoRequest) ProtoMessage() {}

func (x *AnimeSerieDubVideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aspb_servers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnimeSerieDubVideoRequest.ProtoReflect.Descriptor instead.
func (*AnimeSerieDubVideoRequest) Descriptor() ([]byte, []int) {
	return file_aspb_servers_proto_rawDescGZIP(), []int{1}
}

func (x *AnimeSerieDubVideoRequest) GetVideos() []*AnimeSerieVideoRequest {
	if x != nil {
		return x.Videos
	}
	return nil
}

type AnimeSerieSubVideoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Videos []*AnimeSerieVideoResponse `protobuf:"bytes,1,rep,name=videos,proto3" json:"videos,omitempty"`
}

func (x *AnimeSerieSubVideoResponse) Reset() {
	*x = AnimeSerieSubVideoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aspb_servers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnimeSerieSubVideoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnimeSerieSubVideoResponse) ProtoMessage() {}

func (x *AnimeSerieSubVideoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aspb_servers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnimeSerieSubVideoResponse.ProtoReflect.Descriptor instead.
func (*AnimeSerieSubVideoResponse) Descriptor() ([]byte, []int) {
	return file_aspb_servers_proto_rawDescGZIP(), []int{2}
}

func (x *AnimeSerieSubVideoResponse) GetVideos() []*AnimeSerieVideoResponse {
	if x != nil {
		return x.Videos
	}
	return nil
}

type AnimeSerieDubVideoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Videos []*AnimeSerieVideoResponse `protobuf:"bytes,1,rep,name=videos,proto3" json:"videos,omitempty"`
}

func (x *AnimeSerieDubVideoResponse) Reset() {
	*x = AnimeSerieDubVideoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aspb_servers_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnimeSerieDubVideoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnimeSerieDubVideoResponse) ProtoMessage() {}

func (x *AnimeSerieDubVideoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aspb_servers_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnimeSerieDubVideoResponse.ProtoReflect.Descriptor instead.
func (*AnimeSerieDubVideoResponse) Descriptor() ([]byte, []int) {
	return file_aspb_servers_proto_rawDescGZIP(), []int{3}
}

func (x *AnimeSerieDubVideoResponse) GetVideos() []*AnimeSerieVideoResponse {
	if x != nil {
		return x.Videos
	}
	return nil
}

var File_aspb_servers_proto protoreflect.FileDescriptor

var file_aspb_servers_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x73, 0x70, 0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x73, 0x70, 0x62, 0x1a, 0x10, 0x61, 0x73, 0x70, 0x62,
	0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51, 0x0a, 0x19,
	0x41, 0x6e, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x53, 0x75, 0x62, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x06, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x73, 0x70, 0x62,
	0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x06, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x22,
	0x51, 0x0a, 0x19, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x44, 0x75, 0x62,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x06,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61,
	0x73, 0x70, 0x62, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x06, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x73, 0x22, 0x53, 0x0a, 0x1a, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65,
	0x53, 0x75, 0x62, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x35, 0x0a, 0x06, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x61, 0x73, 0x70, 0x62, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72,
	0x69, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x06, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x22, 0x53, 0x0a, 0x1a, 0x41, 0x6e, 0x69, 0x6d, 0x65,
	0x53, 0x65, 0x72, 0x69, 0x65, 0x44, 0x75, 0x62, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x73, 0x70, 0x62, 0x2e, 0x41, 0x6e, 0x69,
	0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x42, 0x2b, 0x5a, 0x29,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6a, 0x2d, 0x79, 0x61,
	0x63, 0x69, 0x6e, 0x65, 0x2d, 0x66, 0x6c, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x6f, 0x6a,
	0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x61, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_aspb_servers_proto_rawDescOnce sync.Once
	file_aspb_servers_proto_rawDescData = file_aspb_servers_proto_rawDesc
)

func file_aspb_servers_proto_rawDescGZIP() []byte {
	file_aspb_servers_proto_rawDescOnce.Do(func() {
		file_aspb_servers_proto_rawDescData = protoimpl.X.CompressGZIP(file_aspb_servers_proto_rawDescData)
	})
	return file_aspb_servers_proto_rawDescData
}

var file_aspb_servers_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_aspb_servers_proto_goTypes = []interface{}{
	(*AnimeSerieSubVideoRequest)(nil),  // 0: aspb.AnimeSerieSubVideoRequest
	(*AnimeSerieDubVideoRequest)(nil),  // 1: aspb.AnimeSerieDubVideoRequest
	(*AnimeSerieSubVideoResponse)(nil), // 2: aspb.AnimeSerieSubVideoResponse
	(*AnimeSerieDubVideoResponse)(nil), // 3: aspb.AnimeSerieDubVideoResponse
	(*AnimeSerieVideoRequest)(nil),     // 4: aspb.AnimeSerieVideoRequest
	(*AnimeSerieVideoResponse)(nil),    // 5: aspb.AnimeSerieVideoResponse
}
var file_aspb_servers_proto_depIdxs = []int32{
	4, // 0: aspb.AnimeSerieSubVideoRequest.videos:type_name -> aspb.AnimeSerieVideoRequest
	4, // 1: aspb.AnimeSerieDubVideoRequest.videos:type_name -> aspb.AnimeSerieVideoRequest
	5, // 2: aspb.AnimeSerieSubVideoResponse.videos:type_name -> aspb.AnimeSerieVideoResponse
	5, // 3: aspb.AnimeSerieDubVideoResponse.videos:type_name -> aspb.AnimeSerieVideoResponse
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_aspb_servers_proto_init() }
func file_aspb_servers_proto_init() {
	if File_aspb_servers_proto != nil {
		return
	}
	file_aspb_video_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_aspb_servers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnimeSerieSubVideoRequest); i {
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
		file_aspb_servers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnimeSerieDubVideoRequest); i {
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
		file_aspb_servers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnimeSerieSubVideoResponse); i {
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
		file_aspb_servers_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnimeSerieDubVideoResponse); i {
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
			RawDescriptor: file_aspb_servers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aspb_servers_proto_goTypes,
		DependencyIndexes: file_aspb_servers_proto_depIdxs,
		MessageInfos:      file_aspb_servers_proto_msgTypes,
	}.Build()
	File_aspb_servers_proto = out.File
	file_aspb_servers_proto_rawDesc = nil
	file_aspb_servers_proto_goTypes = nil
	file_aspb_servers_proto_depIdxs = nil
}
