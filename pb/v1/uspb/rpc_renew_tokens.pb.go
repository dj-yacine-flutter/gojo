// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: v1/uspb/rpc_renew_tokens.proto

package uspbv1

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

type RenewTokensRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefreshToken string `protobuf:"bytes,1,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
}

func (x *RenewTokensRequest) Reset() {
	*x = RenewTokensRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_uspb_rpc_renew_tokens_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenewTokensRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenewTokensRequest) ProtoMessage() {}

func (x *RenewTokensRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_uspb_rpc_renew_tokens_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenewTokensRequest.ProtoReflect.Descriptor instead.
func (*RenewTokensRequest) Descriptor() ([]byte, []int) {
	return file_v1_uspb_rpc_renew_tokens_proto_rawDescGZIP(), []int{0}
}

func (x *RenewTokensRequest) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type RenewTokensResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionID             string                 `protobuf:"bytes,1,opt,name=sessionID,proto3" json:"sessionID,omitempty"`
	AccessToken           string                 `protobuf:"bytes,2,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	AccessTokenExpiresAt  *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=accessTokenExpiresAt,proto3" json:"accessTokenExpiresAt,omitempty"`
	RefreshToken          string                 `protobuf:"bytes,4,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
	RefreshTokenExpiresAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=refreshTokenExpiresAt,proto3" json:"refreshTokenExpiresAt,omitempty"`
}

func (x *RenewTokensResponse) Reset() {
	*x = RenewTokensResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_uspb_rpc_renew_tokens_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenewTokensResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenewTokensResponse) ProtoMessage() {}

func (x *RenewTokensResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_uspb_rpc_renew_tokens_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenewTokensResponse.ProtoReflect.Descriptor instead.
func (*RenewTokensResponse) Descriptor() ([]byte, []int) {
	return file_v1_uspb_rpc_renew_tokens_proto_rawDescGZIP(), []int{1}
}

func (x *RenewTokensResponse) GetSessionID() string {
	if x != nil {
		return x.SessionID
	}
	return ""
}

func (x *RenewTokensResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *RenewTokensResponse) GetAccessTokenExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.AccessTokenExpiresAt
	}
	return nil
}

func (x *RenewTokensResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *RenewTokensResponse) GetRefreshTokenExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.RefreshTokenExpiresAt
	}
	return nil
}

var File_v1_uspb_rpc_renew_tokens_proto protoreflect.FileDescriptor

var file_v1_uspb_rpc_renew_tokens_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x72, 0x65,
	0x6e, 0x65, 0x77, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x09, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x70, 0x62, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a, 0x12,
	0x52, 0x65, 0x6e, 0x65, 0x77, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x9b, 0x02, 0x0a, 0x13, 0x52, 0x65, 0x6e, 0x65, 0x77,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x4e,
	0x0a, 0x14, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x14, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x12, 0x22,
	0x0a, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x50, 0x0a, 0x15, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x15, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x73, 0x41, 0x74, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x64, 0x6a, 0x2d, 0x79, 0x61, 0x63, 0x69, 0x6e, 0x65, 0x2d, 0x66, 0x6c, 0x75,
	0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x6f, 0x6a, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x2f,
	0x75, 0x73, 0x70, 0x62, 0x3b, 0x75, 0x73, 0x70, 0x62, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_v1_uspb_rpc_renew_tokens_proto_rawDescOnce sync.Once
	file_v1_uspb_rpc_renew_tokens_proto_rawDescData = file_v1_uspb_rpc_renew_tokens_proto_rawDesc
)

func file_v1_uspb_rpc_renew_tokens_proto_rawDescGZIP() []byte {
	file_v1_uspb_rpc_renew_tokens_proto_rawDescOnce.Do(func() {
		file_v1_uspb_rpc_renew_tokens_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_uspb_rpc_renew_tokens_proto_rawDescData)
	})
	return file_v1_uspb_rpc_renew_tokens_proto_rawDescData
}

var file_v1_uspb_rpc_renew_tokens_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_uspb_rpc_renew_tokens_proto_goTypes = []interface{}{
	(*RenewTokensRequest)(nil),    // 0: v1.uspbv1.RenewTokensRequest
	(*RenewTokensResponse)(nil),   // 1: v1.uspbv1.RenewTokensResponse
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_v1_uspb_rpc_renew_tokens_proto_depIdxs = []int32{
	2, // 0: v1.uspbv1.RenewTokensResponse.accessTokenExpiresAt:type_name -> google.protobuf.Timestamp
	2, // 1: v1.uspbv1.RenewTokensResponse.refreshTokenExpiresAt:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_v1_uspb_rpc_renew_tokens_proto_init() }
func file_v1_uspb_rpc_renew_tokens_proto_init() {
	if File_v1_uspb_rpc_renew_tokens_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_uspb_rpc_renew_tokens_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenewTokensRequest); i {
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
		file_v1_uspb_rpc_renew_tokens_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenewTokensResponse); i {
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
			RawDescriptor: file_v1_uspb_rpc_renew_tokens_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_uspb_rpc_renew_tokens_proto_goTypes,
		DependencyIndexes: file_v1_uspb_rpc_renew_tokens_proto_depIdxs,
		MessageInfos:      file_v1_uspb_rpc_renew_tokens_proto_msgTypes,
	}.Build()
	File_v1_uspb_rpc_renew_tokens_proto = out.File
	file_v1_uspb_rpc_renew_tokens_proto_rawDesc = nil
	file_v1_uspb_rpc_renew_tokens_proto_goTypes = nil
	file_v1_uspb_rpc_renew_tokens_proto_depIdxs = nil
}
