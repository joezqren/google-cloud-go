// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.3
// source: google/cloud/cloudcontrolspartner/v1beta/completion_state.proto

package cloudcontrolspartnerpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Enum for possible completion states.
type CompletionState int32

const (
	// Unspecified completion state.
	CompletionState_COMPLETION_STATE_UNSPECIFIED CompletionState = 0
	// Task started (has start date) but not yet completed.
	CompletionState_PENDING CompletionState = 1
	// Succeeded state.
	CompletionState_SUCCEEDED CompletionState = 2
	// Failed state.
	CompletionState_FAILED CompletionState = 3
	// Not applicable state.
	CompletionState_NOT_APPLICABLE CompletionState = 4
)

// Enum value maps for CompletionState.
var (
	CompletionState_name = map[int32]string{
		0: "COMPLETION_STATE_UNSPECIFIED",
		1: "PENDING",
		2: "SUCCEEDED",
		3: "FAILED",
		4: "NOT_APPLICABLE",
	}
	CompletionState_value = map[string]int32{
		"COMPLETION_STATE_UNSPECIFIED": 0,
		"PENDING":                      1,
		"SUCCEEDED":                    2,
		"FAILED":                       3,
		"NOT_APPLICABLE":               4,
	}
)

func (x CompletionState) Enum() *CompletionState {
	p := new(CompletionState)
	*p = x
	return p
}

func (x CompletionState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CompletionState) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_cloudcontrolspartner_v1beta_completion_state_proto_enumTypes[0].Descriptor()
}

func (CompletionState) Type() protoreflect.EnumType {
	return &file_google_cloud_cloudcontrolspartner_v1beta_completion_state_proto_enumTypes[0]
}

func (x CompletionState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CompletionState.Descriptor instead.
func (CompletionState) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_cloudcontrolspartner_v1beta_completion_state_proto_rawDescGZIP(), []int{0}
}

var File_google_cloud_cloudcontrolspartner_v1beta_completion_state_proto protoreflect.FileDescriptor

var file_google_cloud_cloudcontrolspartner_v1beta_completion_state_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x70, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x70, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2a, 0x6f, 0x0a, 0x0f, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x20,
	0x0a, 0x1c, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0d, 0x0a,
	0x09, 0x53, 0x55, 0x43, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06,
	0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x4e, 0x4f, 0x54, 0x5f,
	0x41, 0x50, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x04, 0x42, 0xac, 0x02, 0x0a,
	0x2c, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x70,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x42, 0x14, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x60, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f,
	0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x70, 0x62,
	0x3b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x70, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x70, 0x62, 0xaa, 0x02, 0x28, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x73, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x56, 0x31, 0x42, 0x65,
	0x74, 0x61, 0xca, 0x02, 0x28, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x50,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0xea, 0x02, 0x2b,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x50, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_cloudcontrolspartner_v1beta_completion_state_proto_rawDescOnce sync.Once
	file_google_cloud_cloudcontrolspartner_v1beta_completion_state_proto_rawDescData = file_google_cloud_cloudcontrolspartner_v1beta_completion_state_proto_rawDesc
)

func file_google_cloud_cloudcontrolspartner_v1beta_completion_state_proto_rawDescGZIP() []byte {
	file_google_cloud_cloudcontrolspartner_v1beta_completion_state_proto_rawDescOnce.Do(func() {
		file_google_cloud_cloudcontrolspartner_v1beta_completion_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_cloudcontrolspartner_v1beta_completion_state_proto_rawDescData)
	})
	return file_google_cloud_cloudcontrolspartner_v1beta_completion_state_proto_rawDescData
}

var file_google_cloud_cloudcontrolspartner_v1beta_completion_state_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_cloudcontrolspartner_v1beta_completion_state_proto_goTypes = []any{
	(CompletionState)(0), // 0: google.cloud.cloudcontrolspartner.v1beta.CompletionState
}
var file_google_cloud_cloudcontrolspartner_v1beta_completion_state_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_cloudcontrolspartner_v1beta_completion_state_proto_init() }
func file_google_cloud_cloudcontrolspartner_v1beta_completion_state_proto_init() {
	if File_google_cloud_cloudcontrolspartner_v1beta_completion_state_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_cloudcontrolspartner_v1beta_completion_state_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_cloudcontrolspartner_v1beta_completion_state_proto_goTypes,
		DependencyIndexes: file_google_cloud_cloudcontrolspartner_v1beta_completion_state_proto_depIdxs,
		EnumInfos:         file_google_cloud_cloudcontrolspartner_v1beta_completion_state_proto_enumTypes,
	}.Build()
	File_google_cloud_cloudcontrolspartner_v1beta_completion_state_proto = out.File
	file_google_cloud_cloudcontrolspartner_v1beta_completion_state_proto_rawDesc = nil
	file_google_cloud_cloudcontrolspartner_v1beta_completion_state_proto_goTypes = nil
	file_google_cloud_cloudcontrolspartner_v1beta_completion_state_proto_depIdxs = nil
}
