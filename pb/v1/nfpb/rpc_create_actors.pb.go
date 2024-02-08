// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: v1/nfpb/rpc_create_actors.proto

package nfpbv1

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

type CreateActorsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Actors []*ActorRequest `protobuf:"bytes,1,rep,name=actors,proto3" json:"actors,omitempty"`
}

func (x *CreateActorsRequest) Reset() {
	*x = CreateActorsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_nfpb_rpc_create_actors_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateActorsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateActorsRequest) ProtoMessage() {}

func (x *CreateActorsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_nfpb_rpc_create_actors_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateActorsRequest.ProtoReflect.Descriptor instead.
func (*CreateActorsRequest) Descriptor() ([]byte, []int) {
	return file_v1_nfpb_rpc_create_actors_proto_rawDescGZIP(), []int{0}
}

func (x *CreateActorsRequest) GetActors() []*ActorRequest {
	if x != nil {
		return x.Actors
	}
	return nil
}

type CreateActorsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Actors []*ActorResponse `protobuf:"bytes,1,rep,name=actors,proto3" json:"actors,omitempty"`
}

func (x *CreateActorsResponse) Reset() {
	*x = CreateActorsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_nfpb_rpc_create_actors_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateActorsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateActorsResponse) ProtoMessage() {}

func (x *CreateActorsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_nfpb_rpc_create_actors_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateActorsResponse.ProtoReflect.Descriptor instead.
func (*CreateActorsResponse) Descriptor() ([]byte, []int) {
	return file_v1_nfpb_rpc_create_actors_proto_rawDescGZIP(), []int{1}
}

func (x *CreateActorsResponse) GetActors() []*ActorResponse {
	if x != nil {
		return x.Actors
	}
	return nil
}

var File_v1_nfpb_rpc_create_actors_proto protoreflect.FileDescriptor

var file_v1_nfpb_rpc_create_actors_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x76, 0x31, 0x2f, 0x6e, 0x66, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x76, 0x31, 0x2e, 0x6e, 0x66, 0x70, 0x62, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x31,
	0x2f, 0x6e, 0x66, 0x70, 0x62, 0x2f, 0x6d, 0x73, 0x67, 0x5f, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x63, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x06,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x76,
	0x31, 0x2e, 0x6e, 0x66, 0x70, 0x62, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x06, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x22, 0x48, 0x0a,
	0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x76, 0x31, 0x2e, 0x6e, 0x66, 0x70, 0x62, 0x76,
	0x31, 0x2e, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x06, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6a, 0x2d, 0x79, 0x61, 0x63, 0x69, 0x6e, 0x65, 0x2d,
	0x66, 0x6c, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x6f, 0x6a, 0x6f, 0x2f, 0x70, 0x62, 0x2f,
	0x76, 0x31, 0x2f, 0x6e, 0x66, 0x70, 0x62, 0x3b, 0x6e, 0x66, 0x70, 0x62, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_nfpb_rpc_create_actors_proto_rawDescOnce sync.Once
	file_v1_nfpb_rpc_create_actors_proto_rawDescData = file_v1_nfpb_rpc_create_actors_proto_rawDesc
)

func file_v1_nfpb_rpc_create_actors_proto_rawDescGZIP() []byte {
	file_v1_nfpb_rpc_create_actors_proto_rawDescOnce.Do(func() {
		file_v1_nfpb_rpc_create_actors_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_nfpb_rpc_create_actors_proto_rawDescData)
	})
	return file_v1_nfpb_rpc_create_actors_proto_rawDescData
}

var file_v1_nfpb_rpc_create_actors_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_nfpb_rpc_create_actors_proto_goTypes = []interface{}{
	(*CreateActorsRequest)(nil),  // 0: v1.nfpbv1.CreateActorsRequest
	(*CreateActorsResponse)(nil), // 1: v1.nfpbv1.CreateActorsResponse
	(*ActorRequest)(nil),         // 2: v1.nfpbv1.ActorRequest
	(*ActorResponse)(nil),        // 3: v1.nfpbv1.ActorResponse
}
var file_v1_nfpb_rpc_create_actors_proto_depIdxs = []int32{
	2, // 0: v1.nfpbv1.CreateActorsRequest.actors:type_name -> v1.nfpbv1.ActorRequest
	3, // 1: v1.nfpbv1.CreateActorsResponse.actors:type_name -> v1.nfpbv1.ActorResponse
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_v1_nfpb_rpc_create_actors_proto_init() }
func file_v1_nfpb_rpc_create_actors_proto_init() {
	if File_v1_nfpb_rpc_create_actors_proto != nil {
		return
	}
	file_v1_nfpb_msg_actor_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_nfpb_rpc_create_actors_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateActorsRequest); i {
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
		file_v1_nfpb_rpc_create_actors_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateActorsResponse); i {
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
			RawDescriptor: file_v1_nfpb_rpc_create_actors_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_nfpb_rpc_create_actors_proto_goTypes,
		DependencyIndexes: file_v1_nfpb_rpc_create_actors_proto_depIdxs,
		MessageInfos:      file_v1_nfpb_rpc_create_actors_proto_msgTypes,
	}.Build()
	File_v1_nfpb_rpc_create_actors_proto = out.File
	file_v1_nfpb_rpc_create_actors_proto_rawDesc = nil
	file_v1_nfpb_rpc_create_actors_proto_goTypes = nil
	file_v1_nfpb_rpc_create_actors_proto_depIdxs = nil
}
