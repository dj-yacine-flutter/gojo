// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: v1/nfpb/service_info.proto

package nfpbv1

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

var File_v1_nfpb_service_info_proto protoreflect.FileDescriptor

var file_v1_nfpb_service_info_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x76, 0x31, 0x2f, 0x6e, 0x66, 0x70, 0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x76, 0x31,
	0x2e, 0x6e, 0x66, 0x70, 0x62, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x76, 0x31, 0x2f, 0x6e, 0x66, 0x70, 0x62, 0x2f, 0x72,
	0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x76, 0x31, 0x2f, 0x6e, 0x66, 0x70, 0x62, 0x2f,
	0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x69,
	0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x76, 0x31, 0x2f, 0x6e, 0x66, 0x70,
	0x62, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x76, 0x31,
	0x2f, 0x6e, 0x66, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x6c,
	0x6c, 0x5f, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21,
	0x76, 0x31, 0x2f, 0x6e, 0x66, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65, 0x74, 0x5f,
	0x61, 0x6c, 0x6c, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x23, 0x76, 0x31, 0x2f, 0x6e, 0x66, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x67,
	0x65, 0x74, 0x5f, 0x61, 0x6c, 0x6c, 0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xa8, 0x07, 0x0a, 0x0b, 0x49, 0x6e, 0x66, 0x6f, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x97, 0x01, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x12, 0x1e, 0x2e, 0x76, 0x31, 0x2e, 0x6e, 0x66, 0x70,
	0x62, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x76, 0x31, 0x2e, 0x6e, 0x66, 0x70,
	0x62, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x46, 0x92, 0x41, 0x25, 0x1a, 0x23, 0x55,
	0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x67, 0x65, 0x6e, 0x72,
	0x65, 0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f, 0x61, 0x6e,
	0x69, 0x6d, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x9c, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x69,
	0x6f, 0x73, 0x12, 0x1f, 0x2e, 0x76, 0x31, 0x2e, 0x6e, 0x66, 0x70, 0x62, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x76, 0x31, 0x2e, 0x6e, 0x66, 0x70, 0x62, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x48, 0x92, 0x41, 0x26, 0x1a, 0x24, 0x55, 0x73, 0x65, 0x20,
	0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x20, 0x61, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x73,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f, 0x61, 0x6e, 0x69, 0x6d,
	0x65, 0x2f, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0xa0, 0x01, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x73, 0x12, 0x21, 0x2e, 0x76, 0x31, 0x2e, 0x6e, 0x66, 0x70, 0x62, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x76, 0x31, 0x2e, 0x6e, 0x66, 0x70, 0x62,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x46, 0x92, 0x41, 0x28, 0x1a,
	0x26, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f,
	0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a,
	0x22, 0x10, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x8e, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x47, 0x65, 0x6e,
	0x72, 0x65, 0x73, 0x12, 0x1e, 0x2e, 0x76, 0x31, 0x2e, 0x6e, 0x66, 0x70, 0x62, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x76, 0x31, 0x2e, 0x6e, 0x66, 0x70, 0x62, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3d, 0x92, 0x41, 0x26, 0x1a, 0x24, 0x55, 0x73, 0x65, 0x20, 0x74,
	0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x67, 0x65, 0x74, 0x20, 0x61,
	0x6c, 0x6c, 0x20, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x20, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x2f, 0x67, 0x65,
	0x6e, 0x72, 0x65, 0x12, 0x93, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x74,
	0x75, 0x64, 0x69, 0x6f, 0x73, 0x12, 0x1f, 0x2e, 0x76, 0x31, 0x2e, 0x6e, 0x66, 0x70, 0x62, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x76, 0x31, 0x2e, 0x6e, 0x66, 0x70, 0x62,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3f, 0x92, 0x41, 0x27, 0x1a, 0x25, 0x55,
	0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x67,
	0x65, 0x74, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x20, 0x73, 0x74, 0x75,
	0x64, 0x69, 0x6f, 0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x61, 0x6e, 0x69,
	0x6d, 0x65, 0x2f, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x12, 0x95, 0x01, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x12, 0x21, 0x2e,
	0x76, 0x31, 0x2e, 0x6e, 0x66, 0x70, 0x62, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x76, 0x31, 0x2e, 0x6e, 0x66, 0x70, 0x62, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3b, 0x92, 0x41, 0x27, 0x1a, 0x25, 0x55, 0x73, 0x65, 0x20, 0x74,
	0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x67, 0x65, 0x74, 0x20, 0x61,
	0x6c, 0x6c, 0x20, 0x61, 0x70, 0x70, 0x20, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x64, 0x6a, 0x2d, 0x79, 0x61, 0x63, 0x69, 0x6e, 0x65, 0x2d, 0x66, 0x6c, 0x75, 0x74, 0x74, 0x65,
	0x72, 0x2f, 0x67, 0x6f, 0x6a, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x66, 0x70,
	0x62, 0x3b, 0x6e, 0x66, 0x70, 0x62, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_v1_nfpb_service_info_proto_goTypes = []interface{}{
	(*CreateGenresRequest)(nil),     // 0: v1.nfpbv1.CreateGenresRequest
	(*CreateStudiosRequest)(nil),    // 1: v1.nfpbv1.CreateStudiosRequest
	(*CreateLanguagesRequest)(nil),  // 2: v1.nfpbv1.CreateLanguagesRequest
	(*GetAllGenresRequest)(nil),     // 3: v1.nfpbv1.GetAllGenresRequest
	(*GetAllStudiosRequest)(nil),    // 4: v1.nfpbv1.GetAllStudiosRequest
	(*GetAllLanguagesRequest)(nil),  // 5: v1.nfpbv1.GetAllLanguagesRequest
	(*CreateGenresResponse)(nil),    // 6: v1.nfpbv1.CreateGenresResponse
	(*CreateStudiosResponse)(nil),   // 7: v1.nfpbv1.CreateStudiosResponse
	(*CreateLanguagesResponse)(nil), // 8: v1.nfpbv1.CreateLanguagesResponse
	(*GetAllGenresResponse)(nil),    // 9: v1.nfpbv1.GetAllGenresResponse
	(*GetAllStudiosResponse)(nil),   // 10: v1.nfpbv1.GetAllStudiosResponse
	(*GetAllLanguagesResponse)(nil), // 11: v1.nfpbv1.GetAllLanguagesResponse
}
var file_v1_nfpb_service_info_proto_depIdxs = []int32{
	0,  // 0: v1.nfpbv1.InfoService.CreateGenres:input_type -> v1.nfpbv1.CreateGenresRequest
	1,  // 1: v1.nfpbv1.InfoService.CreateStudios:input_type -> v1.nfpbv1.CreateStudiosRequest
	2,  // 2: v1.nfpbv1.InfoService.CreateLanguages:input_type -> v1.nfpbv1.CreateLanguagesRequest
	3,  // 3: v1.nfpbv1.InfoService.GetAllGenres:input_type -> v1.nfpbv1.GetAllGenresRequest
	4,  // 4: v1.nfpbv1.InfoService.GetAllStudios:input_type -> v1.nfpbv1.GetAllStudiosRequest
	5,  // 5: v1.nfpbv1.InfoService.GetAllLanguages:input_type -> v1.nfpbv1.GetAllLanguagesRequest
	6,  // 6: v1.nfpbv1.InfoService.CreateGenres:output_type -> v1.nfpbv1.CreateGenresResponse
	7,  // 7: v1.nfpbv1.InfoService.CreateStudios:output_type -> v1.nfpbv1.CreateStudiosResponse
	8,  // 8: v1.nfpbv1.InfoService.CreateLanguages:output_type -> v1.nfpbv1.CreateLanguagesResponse
	9,  // 9: v1.nfpbv1.InfoService.GetAllGenres:output_type -> v1.nfpbv1.GetAllGenresResponse
	10, // 10: v1.nfpbv1.InfoService.GetAllStudios:output_type -> v1.nfpbv1.GetAllStudiosResponse
	11, // 11: v1.nfpbv1.InfoService.GetAllLanguages:output_type -> v1.nfpbv1.GetAllLanguagesResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_v1_nfpb_service_info_proto_init() }
func file_v1_nfpb_service_info_proto_init() {
	if File_v1_nfpb_service_info_proto != nil {
		return
	}
	file_v1_nfpb_rpc_create_genres_proto_init()
	file_v1_nfpb_rpc_create_studios_proto_init()
	file_v1_nfpb_rpc_create_languages_proto_init()
	file_v1_nfpb_rpc_get_all_genres_proto_init()
	file_v1_nfpb_rpc_get_all_studios_proto_init()
	file_v1_nfpb_rpc_get_all_languages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_nfpb_service_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_nfpb_service_info_proto_goTypes,
		DependencyIndexes: file_v1_nfpb_service_info_proto_depIdxs,
	}.Build()
	File_v1_nfpb_service_info_proto = out.File
	file_v1_nfpb_service_info_proto_rawDesc = nil
	file_v1_nfpb_service_info_proto_goTypes = nil
	file_v1_nfpb_service_info_proto_depIdxs = nil
}
