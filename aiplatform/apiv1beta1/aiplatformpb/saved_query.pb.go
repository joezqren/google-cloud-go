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
// source: google/cloud/aiplatform/v1beta1/saved_query.proto

package aiplatformpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A SavedQuery is a view of the dataset. It references a subset of annotations
// by problem type and filters.
type SavedQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Resource name of the SavedQuery.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The user-defined name of the SavedQuery.
	// The name can be up to 128 characters long and can consist of any UTF-8
	// characters.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Some additional information about the SavedQuery.
	Metadata *structpb.Value `protobuf:"bytes,12,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Output only. Timestamp when this SavedQuery was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. Timestamp when SavedQuery was last updated.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Output only. Filters on the Annotations in the dataset.
	AnnotationFilter string `protobuf:"bytes,5,opt,name=annotation_filter,json=annotationFilter,proto3" json:"annotation_filter,omitempty"`
	// Required. Problem type of the SavedQuery.
	// Allowed values:
	//
	// * IMAGE_CLASSIFICATION_SINGLE_LABEL
	// * IMAGE_CLASSIFICATION_MULTI_LABEL
	// * IMAGE_BOUNDING_POLY
	// * IMAGE_BOUNDING_BOX
	// * TEXT_CLASSIFICATION_SINGLE_LABEL
	// * TEXT_CLASSIFICATION_MULTI_LABEL
	// * TEXT_EXTRACTION
	// * TEXT_SENTIMENT
	// * VIDEO_CLASSIFICATION
	// * VIDEO_OBJECT_TRACKING
	ProblemType string `protobuf:"bytes,6,opt,name=problem_type,json=problemType,proto3" json:"problem_type,omitempty"`
	// Output only. Number of AnnotationSpecs in the context of the SavedQuery.
	AnnotationSpecCount int32 `protobuf:"varint,10,opt,name=annotation_spec_count,json=annotationSpecCount,proto3" json:"annotation_spec_count,omitempty"`
	// Used to perform a consistent read-modify-write update. If not set, a blind
	// "overwrite" update happens.
	Etag string `protobuf:"bytes,8,opt,name=etag,proto3" json:"etag,omitempty"`
	// Output only. If the Annotations belonging to the SavedQuery can be used for
	// AutoML training.
	SupportAutomlTraining bool `protobuf:"varint,9,opt,name=support_automl_training,json=supportAutomlTraining,proto3" json:"support_automl_training,omitempty"`
}

func (x *SavedQuery) Reset() {
	*x = SavedQuery{}
	mi := &file_google_cloud_aiplatform_v1beta1_saved_query_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SavedQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SavedQuery) ProtoMessage() {}

func (x *SavedQuery) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_saved_query_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SavedQuery.ProtoReflect.Descriptor instead.
func (*SavedQuery) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_saved_query_proto_rawDescGZIP(), []int{0}
}

func (x *SavedQuery) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SavedQuery) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *SavedQuery) GetMetadata() *structpb.Value {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *SavedQuery) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *SavedQuery) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *SavedQuery) GetAnnotationFilter() string {
	if x != nil {
		return x.AnnotationFilter
	}
	return ""
}

func (x *SavedQuery) GetProblemType() string {
	if x != nil {
		return x.ProblemType
	}
	return ""
}

func (x *SavedQuery) GetAnnotationSpecCount() int32 {
	if x != nil {
		return x.AnnotationSpecCount
	}
	return 0
}

func (x *SavedQuery) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

func (x *SavedQuery) GetSupportAutomlTraining() bool {
	if x != nil {
		return x.SupportAutomlTraining
	}
	return false
}

var File_google_cloud_aiplatform_v1beta1_saved_query_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1beta1_saved_query_proto_rawDesc = []byte{
	0x0a, 0x31, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x73, 0x61, 0x76, 0x65, 0x64, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xec, 0x04, 0x0a, 0x0a, 0x53, 0x61, 0x76, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x17,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x32, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x11, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x10, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0c, 0x70, 0x72, 0x6f,
	0x62, 0x6c, 0x65, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x37, 0x0a, 0x15, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x73, 0x70, 0x65, 0x63, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x13, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x70, 0x65, 0x63, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x74,
	0x61, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x74, 0x61, 0x67, 0x12, 0x3b,
	0x0a, 0x17, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c,
	0x5f, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x15, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x41, 0x75, 0x74,
	0x6f, 0x6d, 0x6c, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x3a, 0x80, 0x01, 0xea, 0x41,
	0x7d, 0x0a, 0x24, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x61, 0x76,
	0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x55, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x7b, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x7d, 0x2f, 0x73, 0x61, 0x76, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73,
	0x2f, 0x7b, 0x73, 0x61, 0x76, 0x65, 0x64, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x7d, 0x42, 0xe6,
	0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0f, 0x53, 0x61, 0x76, 0x65, 0x64, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70,
	0x62, 0x3b, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62, 0xaa, 0x02,
	0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x49,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0x31,
	0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c,
	0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x3a, 0x3a, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1beta1_saved_query_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1beta1_saved_query_proto_rawDescData = file_google_cloud_aiplatform_v1beta1_saved_query_proto_rawDesc
)

func file_google_cloud_aiplatform_v1beta1_saved_query_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1beta1_saved_query_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1beta1_saved_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1beta1_saved_query_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1beta1_saved_query_proto_rawDescData
}

var file_google_cloud_aiplatform_v1beta1_saved_query_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_aiplatform_v1beta1_saved_query_proto_goTypes = []any{
	(*SavedQuery)(nil),            // 0: google.cloud.aiplatform.v1beta1.SavedQuery
	(*structpb.Value)(nil),        // 1: google.protobuf.Value
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_google_cloud_aiplatform_v1beta1_saved_query_proto_depIdxs = []int32{
	1, // 0: google.cloud.aiplatform.v1beta1.SavedQuery.metadata:type_name -> google.protobuf.Value
	2, // 1: google.cloud.aiplatform.v1beta1.SavedQuery.create_time:type_name -> google.protobuf.Timestamp
	2, // 2: google.cloud.aiplatform.v1beta1.SavedQuery.update_time:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_cloud_aiplatform_v1beta1_saved_query_proto_init() }
func file_google_cloud_aiplatform_v1beta1_saved_query_proto_init() {
	if File_google_cloud_aiplatform_v1beta1_saved_query_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_aiplatform_v1beta1_saved_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_aiplatform_v1beta1_saved_query_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1beta1_saved_query_proto_depIdxs,
		MessageInfos:      file_google_cloud_aiplatform_v1beta1_saved_query_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1beta1_saved_query_proto = out.File
	file_google_cloud_aiplatform_v1beta1_saved_query_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1beta1_saved_query_proto_goTypes = nil
	file_google_cloud_aiplatform_v1beta1_saved_query_proto_depIdxs = nil
}
