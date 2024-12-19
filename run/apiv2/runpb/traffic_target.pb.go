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
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.3
// source: google/cloud/run/v2/traffic_target.proto

package runpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The type of instance allocation.
type TrafficTargetAllocationType int32

const (
	// Unspecified instance allocation type.
	TrafficTargetAllocationType_TRAFFIC_TARGET_ALLOCATION_TYPE_UNSPECIFIED TrafficTargetAllocationType = 0
	// Allocates instances to the Service's latest ready Revision.
	TrafficTargetAllocationType_TRAFFIC_TARGET_ALLOCATION_TYPE_LATEST TrafficTargetAllocationType = 1
	// Allocates instances to a Revision by name.
	TrafficTargetAllocationType_TRAFFIC_TARGET_ALLOCATION_TYPE_REVISION TrafficTargetAllocationType = 2
)

// Enum value maps for TrafficTargetAllocationType.
var (
	TrafficTargetAllocationType_name = map[int32]string{
		0: "TRAFFIC_TARGET_ALLOCATION_TYPE_UNSPECIFIED",
		1: "TRAFFIC_TARGET_ALLOCATION_TYPE_LATEST",
		2: "TRAFFIC_TARGET_ALLOCATION_TYPE_REVISION",
	}
	TrafficTargetAllocationType_value = map[string]int32{
		"TRAFFIC_TARGET_ALLOCATION_TYPE_UNSPECIFIED": 0,
		"TRAFFIC_TARGET_ALLOCATION_TYPE_LATEST":      1,
		"TRAFFIC_TARGET_ALLOCATION_TYPE_REVISION":    2,
	}
)

func (x TrafficTargetAllocationType) Enum() *TrafficTargetAllocationType {
	p := new(TrafficTargetAllocationType)
	*p = x
	return p
}

func (x TrafficTargetAllocationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TrafficTargetAllocationType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_run_v2_traffic_target_proto_enumTypes[0].Descriptor()
}

func (TrafficTargetAllocationType) Type() protoreflect.EnumType {
	return &file_google_cloud_run_v2_traffic_target_proto_enumTypes[0]
}

func (x TrafficTargetAllocationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TrafficTargetAllocationType.Descriptor instead.
func (TrafficTargetAllocationType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_run_v2_traffic_target_proto_rawDescGZIP(), []int{0}
}

// Holds a single traffic routing entry for the Service. Allocations can be done
// to a specific Revision name, or pointing to the latest Ready Revision.
type TrafficTarget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The allocation type for this traffic target.
	Type TrafficTargetAllocationType `protobuf:"varint,1,opt,name=type,proto3,enum=google.cloud.run.v2.TrafficTargetAllocationType" json:"type,omitempty"`
	// Revision to which to send this portion of traffic, if traffic allocation is
	// by revision.
	Revision string `protobuf:"bytes,2,opt,name=revision,proto3" json:"revision,omitempty"`
	// Specifies percent of the traffic to this Revision.
	// This defaults to zero if unspecified.
	Percent int32 `protobuf:"varint,3,opt,name=percent,proto3" json:"percent,omitempty"`
	// Indicates a string to be part of the URI to exclusively reference this
	// target.
	Tag string `protobuf:"bytes,4,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *TrafficTarget) Reset() {
	*x = TrafficTarget{}
	mi := &file_google_cloud_run_v2_traffic_target_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TrafficTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrafficTarget) ProtoMessage() {}

func (x *TrafficTarget) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_run_v2_traffic_target_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrafficTarget.ProtoReflect.Descriptor instead.
func (*TrafficTarget) Descriptor() ([]byte, []int) {
	return file_google_cloud_run_v2_traffic_target_proto_rawDescGZIP(), []int{0}
}

func (x *TrafficTarget) GetType() TrafficTargetAllocationType {
	if x != nil {
		return x.Type
	}
	return TrafficTargetAllocationType_TRAFFIC_TARGET_ALLOCATION_TYPE_UNSPECIFIED
}

func (x *TrafficTarget) GetRevision() string {
	if x != nil {
		return x.Revision
	}
	return ""
}

func (x *TrafficTarget) GetPercent() int32 {
	if x != nil {
		return x.Percent
	}
	return 0
}

func (x *TrafficTarget) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

// Represents the observed state of a single `TrafficTarget` entry.
type TrafficTargetStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The allocation type for this traffic target.
	Type TrafficTargetAllocationType `protobuf:"varint,1,opt,name=type,proto3,enum=google.cloud.run.v2.TrafficTargetAllocationType" json:"type,omitempty"`
	// Revision to which this traffic is sent.
	Revision string `protobuf:"bytes,2,opt,name=revision,proto3" json:"revision,omitempty"`
	// Specifies percent of the traffic to this Revision.
	Percent int32 `protobuf:"varint,3,opt,name=percent,proto3" json:"percent,omitempty"`
	// Indicates the string used in the URI to exclusively reference this target.
	Tag string `protobuf:"bytes,4,opt,name=tag,proto3" json:"tag,omitempty"`
	// Displays the target URI.
	Uri string `protobuf:"bytes,5,opt,name=uri,proto3" json:"uri,omitempty"`
}

func (x *TrafficTargetStatus) Reset() {
	*x = TrafficTargetStatus{}
	mi := &file_google_cloud_run_v2_traffic_target_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TrafficTargetStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrafficTargetStatus) ProtoMessage() {}

func (x *TrafficTargetStatus) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_run_v2_traffic_target_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrafficTargetStatus.ProtoReflect.Descriptor instead.
func (*TrafficTargetStatus) Descriptor() ([]byte, []int) {
	return file_google_cloud_run_v2_traffic_target_proto_rawDescGZIP(), []int{1}
}

func (x *TrafficTargetStatus) GetType() TrafficTargetAllocationType {
	if x != nil {
		return x.Type
	}
	return TrafficTargetAllocationType_TRAFFIC_TARGET_ALLOCATION_TYPE_UNSPECIFIED
}

func (x *TrafficTargetStatus) GetRevision() string {
	if x != nil {
		return x.Revision
	}
	return ""
}

func (x *TrafficTargetStatus) GetPercent() int32 {
	if x != nil {
		return x.Percent
	}
	return 0
}

func (x *TrafficTargetStatus) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *TrafficTargetStatus) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

var File_google_cloud_run_v2_traffic_target_proto protoreflect.FileDescriptor

var file_google_cloud_run_v2_traffic_target_proto_rawDesc = []byte{
	0x0a, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72,
	0x75, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x5f, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x32, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x01, 0x0a, 0x0d, 0x54,
	0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x44, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x32,
	0x2e, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x3c, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x20, 0xfa, 0x41, 0x1d, 0x0a, 0x1b, 0x72, 0x75, 0x6e, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x65,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61,
	0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x22, 0xd7, 0x01, 0x0a,
	0x13, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x44, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3c, 0x0a, 0x08, 0x72, 0x65,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x20, 0xfa, 0x41,
	0x1d, 0x0a, 0x1b, 0x72, 0x75, 0x6e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x08,
	0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x74, 0x61, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x2a, 0xa5, 0x01, 0x0a, 0x1b, 0x54, 0x72, 0x61, 0x66, 0x66,
	0x69, 0x63, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2e, 0x0a, 0x2a, 0x54, 0x52, 0x41, 0x46, 0x46, 0x49,
	0x43, 0x5f, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x43, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x29, 0x0a, 0x25, 0x54, 0x52, 0x41, 0x46, 0x46, 0x49,
	0x43, 0x5f, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x43, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4c, 0x41, 0x54, 0x45, 0x53, 0x54, 0x10,
	0x01, 0x12, 0x2b, 0x0a, 0x27, 0x54, 0x52, 0x41, 0x46, 0x46, 0x49, 0x43, 0x5f, 0x54, 0x41, 0x52,
	0x47, 0x45, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x56, 0x49, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x42, 0x5a,
	0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x32, 0x42, 0x12, 0x54, 0x72, 0x61, 0x66, 0x66,
	0x69, 0x63, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x29, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x72, 0x75, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x72,
	0x75, 0x6e, 0x70, 0x62, 0x3b, 0x72, 0x75, 0x6e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_google_cloud_run_v2_traffic_target_proto_rawDescOnce sync.Once
	file_google_cloud_run_v2_traffic_target_proto_rawDescData = file_google_cloud_run_v2_traffic_target_proto_rawDesc
)

func file_google_cloud_run_v2_traffic_target_proto_rawDescGZIP() []byte {
	file_google_cloud_run_v2_traffic_target_proto_rawDescOnce.Do(func() {
		file_google_cloud_run_v2_traffic_target_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_run_v2_traffic_target_proto_rawDescData)
	})
	return file_google_cloud_run_v2_traffic_target_proto_rawDescData
}

var file_google_cloud_run_v2_traffic_target_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_run_v2_traffic_target_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_run_v2_traffic_target_proto_goTypes = []any{
	(TrafficTargetAllocationType)(0), // 0: google.cloud.run.v2.TrafficTargetAllocationType
	(*TrafficTarget)(nil),            // 1: google.cloud.run.v2.TrafficTarget
	(*TrafficTargetStatus)(nil),      // 2: google.cloud.run.v2.TrafficTargetStatus
}
var file_google_cloud_run_v2_traffic_target_proto_depIdxs = []int32{
	0, // 0: google.cloud.run.v2.TrafficTarget.type:type_name -> google.cloud.run.v2.TrafficTargetAllocationType
	0, // 1: google.cloud.run.v2.TrafficTargetStatus.type:type_name -> google.cloud.run.v2.TrafficTargetAllocationType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_cloud_run_v2_traffic_target_proto_init() }
func file_google_cloud_run_v2_traffic_target_proto_init() {
	if File_google_cloud_run_v2_traffic_target_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_run_v2_traffic_target_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_run_v2_traffic_target_proto_goTypes,
		DependencyIndexes: file_google_cloud_run_v2_traffic_target_proto_depIdxs,
		EnumInfos:         file_google_cloud_run_v2_traffic_target_proto_enumTypes,
		MessageInfos:      file_google_cloud_run_v2_traffic_target_proto_msgTypes,
	}.Build()
	File_google_cloud_run_v2_traffic_target_proto = out.File
	file_google_cloud_run_v2_traffic_target_proto_rawDesc = nil
	file_google_cloud_run_v2_traffic_target_proto_goTypes = nil
	file_google_cloud_run_v2_traffic_target_proto_depIdxs = nil
}
