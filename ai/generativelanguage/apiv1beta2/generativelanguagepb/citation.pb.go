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
// source: google/ai/generativelanguage/v1beta2/citation.proto

package generativelanguagepb

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

// A collection of source attributions for a piece of content.
type CitationMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Citations to sources for a specific response.
	CitationSources []*CitationSource `protobuf:"bytes,1,rep,name=citation_sources,json=citationSources,proto3" json:"citation_sources,omitempty"`
}

func (x *CitationMetadata) Reset() {
	*x = CitationMetadata{}
	mi := &file_google_ai_generativelanguage_v1beta2_citation_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CitationMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CitationMetadata) ProtoMessage() {}

func (x *CitationMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_ai_generativelanguage_v1beta2_citation_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CitationMetadata.ProtoReflect.Descriptor instead.
func (*CitationMetadata) Descriptor() ([]byte, []int) {
	return file_google_ai_generativelanguage_v1beta2_citation_proto_rawDescGZIP(), []int{0}
}

func (x *CitationMetadata) GetCitationSources() []*CitationSource {
	if x != nil {
		return x.CitationSources
	}
	return nil
}

// A citation to a source for a portion of a specific response.
type CitationSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. Start of segment of the response that is attributed to this
	// source.
	//
	// Index indicates the start of the segment, measured in bytes.
	StartIndex *int32 `protobuf:"varint,1,opt,name=start_index,json=startIndex,proto3,oneof" json:"start_index,omitempty"`
	// Optional. End of the attributed segment, exclusive.
	EndIndex *int32 `protobuf:"varint,2,opt,name=end_index,json=endIndex,proto3,oneof" json:"end_index,omitempty"`
	// Optional. URI that is attributed as a source for a portion of the text.
	Uri *string `protobuf:"bytes,3,opt,name=uri,proto3,oneof" json:"uri,omitempty"`
	// Optional. License for the GitHub project that is attributed as a source for
	// segment.
	//
	// License info is required for code citations.
	License *string `protobuf:"bytes,4,opt,name=license,proto3,oneof" json:"license,omitempty"`
}

func (x *CitationSource) Reset() {
	*x = CitationSource{}
	mi := &file_google_ai_generativelanguage_v1beta2_citation_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CitationSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CitationSource) ProtoMessage() {}

func (x *CitationSource) ProtoReflect() protoreflect.Message {
	mi := &file_google_ai_generativelanguage_v1beta2_citation_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CitationSource.ProtoReflect.Descriptor instead.
func (*CitationSource) Descriptor() ([]byte, []int) {
	return file_google_ai_generativelanguage_v1beta2_citation_proto_rawDescGZIP(), []int{1}
}

func (x *CitationSource) GetStartIndex() int32 {
	if x != nil && x.StartIndex != nil {
		return *x.StartIndex
	}
	return 0
}

func (x *CitationSource) GetEndIndex() int32 {
	if x != nil && x.EndIndex != nil {
		return *x.EndIndex
	}
	return 0
}

func (x *CitationSource) GetUri() string {
	if x != nil && x.Uri != nil {
		return *x.Uri
	}
	return ""
}

func (x *CitationSource) GetLicense() string {
	if x != nil && x.License != nil {
		return *x.License
	}
	return ""
}

var File_google_ai_generativelanguage_v1beta2_citation_proto protoreflect.FileDescriptor

var file_google_ai_generativelanguage_v1beta2_citation_proto_rawDesc = []byte{
	0x0a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2f, 0x63, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x69,
	0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65,
	0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x73, 0x0a, 0x10,
	0x43, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x5f, 0x0a, 0x10, 0x63, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x69, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76,
	0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x32, 0x2e, 0x43, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x0f, 0x63, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x22, 0xd4, 0x01, 0x0a, 0x0e, 0x43, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x29, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x48, 0x00,
	0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x88, 0x01, 0x01, 0x12,
	0x25, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x48, 0x01, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x48, 0x02, 0x52, 0x03, 0x75, 0x72, 0x69, 0x88,
	0x01, 0x01, 0x12, 0x22, 0x0a, 0x07, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x48, 0x03, 0x52, 0x07, 0x6c, 0x69, 0x63, 0x65,
	0x6e, 0x73, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x75, 0x72, 0x69, 0x42, 0x0a, 0x0a, 0x08,
	0x5f, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x42, 0x9b, 0x01, 0x0a, 0x28, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x69, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x32, 0x42, 0x0d, 0x43, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x69, 0x2f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2f, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x70,
	0x62, 0x3b, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ai_generativelanguage_v1beta2_citation_proto_rawDescOnce sync.Once
	file_google_ai_generativelanguage_v1beta2_citation_proto_rawDescData = file_google_ai_generativelanguage_v1beta2_citation_proto_rawDesc
)

func file_google_ai_generativelanguage_v1beta2_citation_proto_rawDescGZIP() []byte {
	file_google_ai_generativelanguage_v1beta2_citation_proto_rawDescOnce.Do(func() {
		file_google_ai_generativelanguage_v1beta2_citation_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ai_generativelanguage_v1beta2_citation_proto_rawDescData)
	})
	return file_google_ai_generativelanguage_v1beta2_citation_proto_rawDescData
}

var file_google_ai_generativelanguage_v1beta2_citation_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_ai_generativelanguage_v1beta2_citation_proto_goTypes = []any{
	(*CitationMetadata)(nil), // 0: google.ai.generativelanguage.v1beta2.CitationMetadata
	(*CitationSource)(nil),   // 1: google.ai.generativelanguage.v1beta2.CitationSource
}
var file_google_ai_generativelanguage_v1beta2_citation_proto_depIdxs = []int32{
	1, // 0: google.ai.generativelanguage.v1beta2.CitationMetadata.citation_sources:type_name -> google.ai.generativelanguage.v1beta2.CitationSource
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_ai_generativelanguage_v1beta2_citation_proto_init() }
func file_google_ai_generativelanguage_v1beta2_citation_proto_init() {
	if File_google_ai_generativelanguage_v1beta2_citation_proto != nil {
		return
	}
	file_google_ai_generativelanguage_v1beta2_citation_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ai_generativelanguage_v1beta2_citation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ai_generativelanguage_v1beta2_citation_proto_goTypes,
		DependencyIndexes: file_google_ai_generativelanguage_v1beta2_citation_proto_depIdxs,
		MessageInfos:      file_google_ai_generativelanguage_v1beta2_citation_proto_msgTypes,
	}.Build()
	File_google_ai_generativelanguage_v1beta2_citation_proto = out.File
	file_google_ai_generativelanguage_v1beta2_citation_proto_rawDesc = nil
	file_google_ai_generativelanguage_v1beta2_citation_proto_goTypes = nil
	file_google_ai_generativelanguage_v1beta2_citation_proto_depIdxs = nil
}
