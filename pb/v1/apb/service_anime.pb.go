// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: v1/apb/service_anime.proto

package apbv1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_v1_apb_service_anime_proto protoreflect.FileDescriptor

var file_v1_apb_service_anime_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x76, 0x31,
	0x2e, 0x61, 0x70, 0x62, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63,
	0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x63, 0x68,
	0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x76,
	0x31, 0x2f, 0x61, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x26, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x76, 0x31, 0x2f, 0x61, 0x70,
	0x62, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x6e, 0x69,
	0x6d, 0x65, 0x5f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25,
	0x76, 0x31, 0x2f, 0x61, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x62, 0x2f, 0x72, 0x70,
	0x63, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x76, 0x31, 0x2f, 0x61,
	0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x6e,
	0x69, 0x6d, 0x65, 0x5f, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x76,
	0x31, 0x2f, 0x61, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63,
	0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xe5, 0x0c, 0x0a, 0x0c, 0x41, 0x6e,
	0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xc4, 0x01, 0x0a, 0x14, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63,
	0x74, 0x65, 0x72, 0x12, 0x25, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x70, 0x62, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x76, 0x31, 0x2e,
	0x61, 0x70, 0x62, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d,
	0x65, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x5d, 0x92, 0x41, 0x2a, 0x1a, 0x28, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69,
	0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20,
	0x61, 0x20, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x20, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65,
	0x72, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x3a, 0x01, 0x2a, 0x32, 0x25, 0x2f, 0x61, 0x6e, 0x69,
	0x6d, 0x65, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x2f, 0x7b, 0x63, 0x68,
	0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x49, 0x44, 0x7d, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0xa6, 0x01, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d,
	0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x20, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x70, 0x62, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4c, 0x69, 0x6e, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x70, 0x62,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4c, 0x69,
	0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4e, 0x92, 0x41, 0x25, 0x1a,
	0x23, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f,
	0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x20,
	0x6c, 0x69, 0x6e, 0x6b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a, 0x01, 0x2a, 0x32, 0x1b, 0x2f,
	0x61, 0x6e, 0x69, 0x6d, 0x65, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x7b, 0x6c, 0x69, 0x6e, 0x6b,
	0x49, 0x44, 0x7d, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0xbe, 0x01, 0x0a, 0x13, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x24, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x70, 0x62, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x70,
	0x62, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x5a, 0x92, 0x41, 0x29, 0x1a, 0x27, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41,
	0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x61,
	0x6e, 0x69, 0x6d, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x28, 0x3a, 0x01, 0x2a, 0x32, 0x23, 0x2f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x7b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x49, 0x44, 0x7d, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0xac, 0x01, 0x0a, 0x10,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x12, 0x21, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x70, 0x62, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x70, 0x62, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x51, 0x92, 0x41, 0x26, 0x1a, 0x24, 0x55, 0x73,
	0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x20, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x3a, 0x01, 0x2a, 0x32, 0x1d, 0x2f, 0x61, 0x6e,
	0x69, 0x6d, 0x65, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2f, 0x7b, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x49, 0x44, 0x7d, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0xb8, 0x01, 0x0a, 0x12, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x12, 0x23, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x70, 0x62, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x70, 0x62, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x54, 0x6f, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x57, 0x92, 0x41,
	0x28, 0x1a, 0x26, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20,
	0x74, 0x6f, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x61, 0x6e, 0x69, 0x6d,
	0x65, 0x20, 0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x3a,
	0x01, 0x2a, 0x32, 0x21, 0x2f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x2f, 0x74, 0x6f, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x2f, 0x7b, 0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x7d, 0x2f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0xac, 0x01, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x6e, 0x69, 0x6d, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x21, 0x2e, 0x76, 0x31, 0x2e,
	0x61, 0x70, 0x62, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d,
	0x65, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x76, 0x31, 0x2e, 0x61, 0x70, 0x62, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x6e, 0x69, 0x6d, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x51, 0x92, 0x41, 0x26, 0x1a, 0x24, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73,
	0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x61,
	0x20, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x20, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x22, 0x3a, 0x01, 0x2a, 0x32, 0x1d, 0x2f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x2f, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x2f, 0x7b, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x49, 0x44, 0x7d, 0x2f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0xa0, 0x01, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x6e, 0x69, 0x6d, 0x65, 0x54, 0x61, 0x67, 0x12, 0x1f, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x70, 0x62,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x54, 0x61,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x70,
	0x62, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x54,
	0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4b, 0x92, 0x41, 0x24, 0x1a,
	0x22, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f,
	0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x20,
	0x74, 0x61, 0x67, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01, 0x2a, 0x32, 0x19, 0x2f, 0x61,
	0x6e, 0x69, 0x6d, 0x65, 0x2f, 0x74, 0x61, 0x67, 0x2f, 0x7b, 0x74, 0x61, 0x67, 0x49, 0x44, 0x7d,
	0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0xb8, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x12, 0x23,
	0x2e, 0x76, 0x31, 0x2e, 0x61, 0x70, 0x62, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x6e, 0x69, 0x6d, 0x65, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x70, 0x62, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x57, 0x92, 0x41, 0x28, 0x1a, 0x26,
	0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x20, 0x74,
	0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x3a, 0x01, 0x2a, 0x32,
	0x21, 0x2f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x2f, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2f,
	0x7b, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x49, 0x44, 0x7d, 0x2f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0xac, 0x01, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x69,
	0x6d, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x21, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x70, 0x62,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x76, 0x31, 0x2e,
	0x61, 0x70, 0x62, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d,
	0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x51,
	0x92, 0x41, 0x26, 0x1a, 0x24, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50,
	0x49, 0x20, 0x74, 0x6f, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x61, 0x6e,
	0x69, 0x6d, 0x65, 0x20, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x3a,
	0x01, 0x2a, 0x32, 0x1d, 0x2f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x2f, 0x7b, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x44, 0x7d, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x64, 0x6a, 0x2d, 0x79, 0x61, 0x63, 0x69, 0x6e, 0x65, 0x2d, 0x66, 0x6c, 0x75, 0x74, 0x74, 0x65,
	0x72, 0x2f, 0x67, 0x6f, 0x6a, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x62,
	0x3b, 0x61, 0x70, 0x62, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_v1_apb_service_anime_proto_goTypes = []interface{}{
	(*UpdateAnimeCharacterRequest)(nil),  // 0: v1.apbv1.UpdateAnimeCharacterRequest
	(*UpdateAnimeLinkRequest)(nil),       // 1: v1.apbv1.UpdateAnimeLinkRequest
	(*UpdateAnimeResourceRequest)(nil),   // 2: v1.apbv1.UpdateAnimeResourceRequest
	(*UpdateAnimeVideoRequest)(nil),      // 3: v1.apbv1.UpdateAnimeVideoRequest
	(*UpdateAnimeTorrentRequest)(nil),    // 4: v1.apbv1.UpdateAnimeTorrentRequest
	(*UpdateAnimeTitleRequest)(nil),      // 5: v1.apbv1.UpdateAnimeTitleRequest
	(*UpdateAnimeTagRequest)(nil),        // 6: v1.apbv1.UpdateAnimeTagRequest
	(*UpdateAnimeTrailerRequest)(nil),    // 7: v1.apbv1.UpdateAnimeTrailerRequest
	(*UpdateAnimeImageRequest)(nil),      // 8: v1.apbv1.UpdateAnimeImageRequest
	(*UpdateAnimeCharacterResponse)(nil), // 9: v1.apbv1.UpdateAnimeCharacterResponse
	(*UpdateAnimeLinkResponse)(nil),      // 10: v1.apbv1.UpdateAnimeLinkResponse
	(*UpdateAnimeResourceResponse)(nil),  // 11: v1.apbv1.UpdateAnimeResourceResponse
	(*UpdateAnimeVideoResponse)(nil),     // 12: v1.apbv1.UpdateAnimeVideoResponse
	(*UpdateAnimeTorrentResponse)(nil),   // 13: v1.apbv1.UpdateAnimeTorrentResponse
	(*UpdateAnimeTitleResponse)(nil),     // 14: v1.apbv1.UpdateAnimeTitleResponse
	(*UpdateAnimeTagResponse)(nil),       // 15: v1.apbv1.UpdateAnimeTagResponse
	(*UpdateAnimeTrailerResponse)(nil),   // 16: v1.apbv1.UpdateAnimeTrailerResponse
	(*UpdateAnimeImageResponse)(nil),     // 17: v1.apbv1.UpdateAnimeImageResponse
}
var file_v1_apb_service_anime_proto_depIdxs = []int32{
	0,  // 0: v1.apbv1.AnimeService.UpdateAnimeCharacter:input_type -> v1.apbv1.UpdateAnimeCharacterRequest
	1,  // 1: v1.apbv1.AnimeService.UpdateAnimeLink:input_type -> v1.apbv1.UpdateAnimeLinkRequest
	2,  // 2: v1.apbv1.AnimeService.UpdateAnimeResource:input_type -> v1.apbv1.UpdateAnimeResourceRequest
	3,  // 3: v1.apbv1.AnimeService.UpdateAnimeVideo:input_type -> v1.apbv1.UpdateAnimeVideoRequest
	4,  // 4: v1.apbv1.AnimeService.UpdateAnimeTorrent:input_type -> v1.apbv1.UpdateAnimeTorrentRequest
	5,  // 5: v1.apbv1.AnimeService.UpdateAnimeTitle:input_type -> v1.apbv1.UpdateAnimeTitleRequest
	6,  // 6: v1.apbv1.AnimeService.UpdateAnimeTag:input_type -> v1.apbv1.UpdateAnimeTagRequest
	7,  // 7: v1.apbv1.AnimeService.UpdateAnimeTrailer:input_type -> v1.apbv1.UpdateAnimeTrailerRequest
	8,  // 8: v1.apbv1.AnimeService.UpdateAnimeImage:input_type -> v1.apbv1.UpdateAnimeImageRequest
	9,  // 9: v1.apbv1.AnimeService.UpdateAnimeCharacter:output_type -> v1.apbv1.UpdateAnimeCharacterResponse
	10, // 10: v1.apbv1.AnimeService.UpdateAnimeLink:output_type -> v1.apbv1.UpdateAnimeLinkResponse
	11, // 11: v1.apbv1.AnimeService.UpdateAnimeResource:output_type -> v1.apbv1.UpdateAnimeResourceResponse
	12, // 12: v1.apbv1.AnimeService.UpdateAnimeVideo:output_type -> v1.apbv1.UpdateAnimeVideoResponse
	13, // 13: v1.apbv1.AnimeService.UpdateAnimeTorrent:output_type -> v1.apbv1.UpdateAnimeTorrentResponse
	14, // 14: v1.apbv1.AnimeService.UpdateAnimeTitle:output_type -> v1.apbv1.UpdateAnimeTitleResponse
	15, // 15: v1.apbv1.AnimeService.UpdateAnimeTag:output_type -> v1.apbv1.UpdateAnimeTagResponse
	16, // 16: v1.apbv1.AnimeService.UpdateAnimeTrailer:output_type -> v1.apbv1.UpdateAnimeTrailerResponse
	17, // 17: v1.apbv1.AnimeService.UpdateAnimeImage:output_type -> v1.apbv1.UpdateAnimeImageResponse
	9,  // [9:18] is the sub-list for method output_type
	0,  // [0:9] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_v1_apb_service_anime_proto_init() }
func file_v1_apb_service_anime_proto_init() {
	if File_v1_apb_service_anime_proto != nil {
		return
	}
	file_v1_apb_rpc_update_anime_character_proto_init()
	file_v1_apb_rpc_update_anime_link_proto_init()
	file_v1_apb_rpc_update_anime_resource_proto_init()
	file_v1_apb_rpc_update_anime_video_proto_init()
	file_v1_apb_rpc_update_anime_torrent_proto_init()
	file_v1_apb_rpc_update_anime_title_proto_init()
	file_v1_apb_rpc_update_anime_tag_proto_init()
	file_v1_apb_rpc_update_anime_trailer_proto_init()
	file_v1_apb_rpc_update_anime_image_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_apb_service_anime_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_apb_service_anime_proto_goTypes,
		DependencyIndexes: file_v1_apb_service_anime_proto_depIdxs,
	}.Build()
	File_v1_apb_service_anime_proto = out.File
	file_v1_apb_service_anime_proto_rawDesc = nil
	file_v1_apb_service_anime_proto_goTypes = nil
	file_v1_apb_service_anime_proto_depIdxs = nil
}
