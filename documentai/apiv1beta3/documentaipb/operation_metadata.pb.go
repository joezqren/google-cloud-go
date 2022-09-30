// Copyright 2022 Google LLC
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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.2
// source: google/cloud/documentai/v1beta3/operation_metadata.proto

package documentaipb

import (
	reflect "reflect"
	sync "sync"

	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// State of the longrunning operation.
type CommonOperationMetadata_State int32

const (
	// Unspecified state.
	CommonOperationMetadata_STATE_UNSPECIFIED CommonOperationMetadata_State = 0
	// Operation is still running.
	CommonOperationMetadata_RUNNING CommonOperationMetadata_State = 1
	// Operation is being cancelled.
	CommonOperationMetadata_CANCELLING CommonOperationMetadata_State = 2
	// Operation succeeded.
	CommonOperationMetadata_SUCCEEDED CommonOperationMetadata_State = 3
	// Operation failed.
	CommonOperationMetadata_FAILED CommonOperationMetadata_State = 4
	// Operation is cancelled.
	CommonOperationMetadata_CANCELLED CommonOperationMetadata_State = 5
)

// Enum value maps for CommonOperationMetadata_State.
var (
	CommonOperationMetadata_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "RUNNING",
		2: "CANCELLING",
		3: "SUCCEEDED",
		4: "FAILED",
		5: "CANCELLED",
	}
	CommonOperationMetadata_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"RUNNING":           1,
		"CANCELLING":        2,
		"SUCCEEDED":         3,
		"FAILED":            4,
		"CANCELLED":         5,
	}
)

func (x CommonOperationMetadata_State) Enum() *CommonOperationMetadata_State {
	p := new(CommonOperationMetadata_State)
	*p = x
	return p
}

func (x CommonOperationMetadata_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommonOperationMetadata_State) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_documentai_v1beta3_operation_metadata_proto_enumTypes[0].Descriptor()
}

func (CommonOperationMetadata_State) Type() protoreflect.EnumType {
	return &file_google_cloud_documentai_v1beta3_operation_metadata_proto_enumTypes[0]
}

func (x CommonOperationMetadata_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommonOperationMetadata_State.Descriptor instead.
func (CommonOperationMetadata_State) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_documentai_v1beta3_operation_metadata_proto_rawDescGZIP(), []int{0, 0}
}

// The common metadata for long running operations.
type CommonOperationMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The state of the operation.
	State CommonOperationMetadata_State `protobuf:"varint,1,opt,name=state,proto3,enum=google.cloud.documentai.v1beta3.CommonOperationMetadata_State" json:"state,omitempty"`
	// A message providing more details about the current state of processing.
	StateMessage string `protobuf:"bytes,2,opt,name=state_message,json=stateMessage,proto3" json:"state_message,omitempty"`
	// A related resource to this operation.
	Resource string `protobuf:"bytes,5,opt,name=resource,proto3" json:"resource,omitempty"`
	// The creation time of the operation.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The last update time of the operation.
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *CommonOperationMetadata) Reset() {
	*x = CommonOperationMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_documentai_v1beta3_operation_metadata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonOperationMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonOperationMetadata) ProtoMessage() {}

func (x *CommonOperationMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_documentai_v1beta3_operation_metadata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonOperationMetadata.ProtoReflect.Descriptor instead.
func (*CommonOperationMetadata) Descriptor() ([]byte, []int) {
	return file_google_cloud_documentai_v1beta3_operation_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *CommonOperationMetadata) GetState() CommonOperationMetadata_State {
	if x != nil {
		return x.State
	}
	return CommonOperationMetadata_STATE_UNSPECIFIED
}

func (x *CommonOperationMetadata) GetStateMessage() string {
	if x != nil {
		return x.StateMessage
	}
	return ""
}

func (x *CommonOperationMetadata) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *CommonOperationMetadata) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *CommonOperationMetadata) GetUpdateTime() *timestamp.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

var File_google_cloud_documentai_v1beta3_operation_metadata_proto protoreflect.FileDescriptor

var file_google_cloud_documentai_v1beta3_operation_metadata_proto_rawDesc = []byte{
	0x0a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x33, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x33, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x03, 0x0a,
	0x17, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x54, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61,
	0x69, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x33, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x65, 0x0a, 0x05, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e,
	0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c,
	0x4c, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x55, 0x43, 0x43, 0x45, 0x45,
	0x44, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10,
	0x04, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x05,
	0x42, 0xf3, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x69,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x33, 0x42, 0x16, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x49, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x33, 0x3b, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x69, 0xaa, 0x02, 0x1f,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x49, 0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0x33, 0xca,
	0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x49, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x33, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x3a, 0x3a, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x49, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_documentai_v1beta3_operation_metadata_proto_rawDescOnce sync.Once
	file_google_cloud_documentai_v1beta3_operation_metadata_proto_rawDescData = file_google_cloud_documentai_v1beta3_operation_metadata_proto_rawDesc
)

func file_google_cloud_documentai_v1beta3_operation_metadata_proto_rawDescGZIP() []byte {
	file_google_cloud_documentai_v1beta3_operation_metadata_proto_rawDescOnce.Do(func() {
		file_google_cloud_documentai_v1beta3_operation_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_documentai_v1beta3_operation_metadata_proto_rawDescData)
	})
	return file_google_cloud_documentai_v1beta3_operation_metadata_proto_rawDescData
}

var file_google_cloud_documentai_v1beta3_operation_metadata_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_documentai_v1beta3_operation_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_documentai_v1beta3_operation_metadata_proto_goTypes = []interface{}{
	(CommonOperationMetadata_State)(0), // 0: google.cloud.documentai.v1beta3.CommonOperationMetadata.State
	(*CommonOperationMetadata)(nil),    // 1: google.cloud.documentai.v1beta3.CommonOperationMetadata
	(*timestamp.Timestamp)(nil),        // 2: google.protobuf.Timestamp
}
var file_google_cloud_documentai_v1beta3_operation_metadata_proto_depIdxs = []int32{
	0, // 0: google.cloud.documentai.v1beta3.CommonOperationMetadata.state:type_name -> google.cloud.documentai.v1beta3.CommonOperationMetadata.State
	2, // 1: google.cloud.documentai.v1beta3.CommonOperationMetadata.create_time:type_name -> google.protobuf.Timestamp
	2, // 2: google.cloud.documentai.v1beta3.CommonOperationMetadata.update_time:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_cloud_documentai_v1beta3_operation_metadata_proto_init() }
func file_google_cloud_documentai_v1beta3_operation_metadata_proto_init() {
	if File_google_cloud_documentai_v1beta3_operation_metadata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_documentai_v1beta3_operation_metadata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonOperationMetadata); i {
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
			RawDescriptor: file_google_cloud_documentai_v1beta3_operation_metadata_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_documentai_v1beta3_operation_metadata_proto_goTypes,
		DependencyIndexes: file_google_cloud_documentai_v1beta3_operation_metadata_proto_depIdxs,
		EnumInfos:         file_google_cloud_documentai_v1beta3_operation_metadata_proto_enumTypes,
		MessageInfos:      file_google_cloud_documentai_v1beta3_operation_metadata_proto_msgTypes,
	}.Build()
	File_google_cloud_documentai_v1beta3_operation_metadata_proto = out.File
	file_google_cloud_documentai_v1beta3_operation_metadata_proto_rawDesc = nil
	file_google_cloud_documentai_v1beta3_operation_metadata_proto_goTypes = nil
	file_google_cloud_documentai_v1beta3_operation_metadata_proto_depIdxs = nil
}
