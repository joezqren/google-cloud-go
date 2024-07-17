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
// source: google/cloud/documentai/v1/document_schema.proto

package documentaipb

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

// Types of occurrences of the entity type in the document.  This
// represents the number of instances, not mentions, of an entity.
// For example, a bank statement might only have one
// `account_number`, but this account number can be mentioned in several
// places on the document.  In this case, the `account_number` is
// considered a `REQUIRED_ONCE` entity type. If, on the other hand, we
// expect a bank statement to contain the status of multiple different
// accounts for the customers, the occurrence type is set to
// `REQUIRED_MULTIPLE`.
type DocumentSchema_EntityType_Property_OccurrenceType int32

const (
	// Unspecified occurrence type.
	DocumentSchema_EntityType_Property_OCCURRENCE_TYPE_UNSPECIFIED DocumentSchema_EntityType_Property_OccurrenceType = 0
	// There will be zero or one instance of this entity type.  The same
	// entity instance may be mentioned multiple times.
	DocumentSchema_EntityType_Property_OPTIONAL_ONCE DocumentSchema_EntityType_Property_OccurrenceType = 1
	// The entity type will appear zero or multiple times.
	DocumentSchema_EntityType_Property_OPTIONAL_MULTIPLE DocumentSchema_EntityType_Property_OccurrenceType = 2
	// The entity type will only appear exactly once.  The same
	// entity instance may be mentioned multiple times.
	DocumentSchema_EntityType_Property_REQUIRED_ONCE DocumentSchema_EntityType_Property_OccurrenceType = 3
	// The entity type will appear once or more times.
	DocumentSchema_EntityType_Property_REQUIRED_MULTIPLE DocumentSchema_EntityType_Property_OccurrenceType = 4
)

// Enum value maps for DocumentSchema_EntityType_Property_OccurrenceType.
var (
	DocumentSchema_EntityType_Property_OccurrenceType_name = map[int32]string{
		0: "OCCURRENCE_TYPE_UNSPECIFIED",
		1: "OPTIONAL_ONCE",
		2: "OPTIONAL_MULTIPLE",
		3: "REQUIRED_ONCE",
		4: "REQUIRED_MULTIPLE",
	}
	DocumentSchema_EntityType_Property_OccurrenceType_value = map[string]int32{
		"OCCURRENCE_TYPE_UNSPECIFIED": 0,
		"OPTIONAL_ONCE":               1,
		"OPTIONAL_MULTIPLE":           2,
		"REQUIRED_ONCE":               3,
		"REQUIRED_MULTIPLE":           4,
	}
)

func (x DocumentSchema_EntityType_Property_OccurrenceType) Enum() *DocumentSchema_EntityType_Property_OccurrenceType {
	p := new(DocumentSchema_EntityType_Property_OccurrenceType)
	*p = x
	return p
}

func (x DocumentSchema_EntityType_Property_OccurrenceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DocumentSchema_EntityType_Property_OccurrenceType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_documentai_v1_document_schema_proto_enumTypes[0].Descriptor()
}

func (DocumentSchema_EntityType_Property_OccurrenceType) Type() protoreflect.EnumType {
	return &file_google_cloud_documentai_v1_document_schema_proto_enumTypes[0]
}

func (x DocumentSchema_EntityType_Property_OccurrenceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DocumentSchema_EntityType_Property_OccurrenceType.Descriptor instead.
func (DocumentSchema_EntityType_Property_OccurrenceType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_documentai_v1_document_schema_proto_rawDescGZIP(), []int{0, 0, 1, 0}
}

// The schema defines the output of the processed document by a processor.
type DocumentSchema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Display name to show to users.
	DisplayName string `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Description of the schema.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Entity types of the schema.
	EntityTypes []*DocumentSchema_EntityType `protobuf:"bytes,3,rep,name=entity_types,json=entityTypes,proto3" json:"entity_types,omitempty"`
	// Metadata of the schema.
	Metadata *DocumentSchema_Metadata `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *DocumentSchema) Reset() {
	*x = DocumentSchema{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_documentai_v1_document_schema_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocumentSchema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocumentSchema) ProtoMessage() {}

func (x *DocumentSchema) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_documentai_v1_document_schema_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocumentSchema.ProtoReflect.Descriptor instead.
func (*DocumentSchema) Descriptor() ([]byte, []int) {
	return file_google_cloud_documentai_v1_document_schema_proto_rawDescGZIP(), []int{0}
}

func (x *DocumentSchema) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *DocumentSchema) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *DocumentSchema) GetEntityTypes() []*DocumentSchema_EntityType {
	if x != nil {
		return x.EntityTypes
	}
	return nil
}

func (x *DocumentSchema) GetMetadata() *DocumentSchema_Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// EntityType is the wrapper of a label of the corresponding model with
// detailed attributes and limitations for entity-based processors. Multiple
// types can also compose a dependency tree to represent nested types.
type DocumentSchema_EntityType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ValueSource:
	//
	//	*DocumentSchema_EntityType_EnumValues_
	ValueSource isDocumentSchema_EntityType_ValueSource `protobuf_oneof:"value_source"`
	// User defined name for the type.
	DisplayName string `protobuf:"bytes,13,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Name of the type. It must be unique within the schema file and
	// cannot be a "Common Type".  The following naming conventions are used:
	//
	//   - Use `snake_casing`.
	//   - Name matching is case-sensitive.
	//   - Maximum 64 characters.
	//   - Must start with a letter.
	//   - Allowed characters: ASCII letters `[a-z0-9_-]`.  (For backward
	//     compatibility internal infrastructure and tooling can handle any ascii
	//     character.)
	//   - The `/` is sometimes used to denote a property of a type.  For example
	//     `line_item/amount`.  This convention is deprecated, but will still be
	//     honored for backward compatibility.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The entity type that this type is derived from.  For now, one and only
	// one should be set.
	BaseTypes []string `protobuf:"bytes,2,rep,name=base_types,json=baseTypes,proto3" json:"base_types,omitempty"`
	// Description the nested structure, or composition of an entity.
	Properties []*DocumentSchema_EntityType_Property `protobuf:"bytes,6,rep,name=properties,proto3" json:"properties,omitempty"`
}

func (x *DocumentSchema_EntityType) Reset() {
	*x = DocumentSchema_EntityType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_documentai_v1_document_schema_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocumentSchema_EntityType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocumentSchema_EntityType) ProtoMessage() {}

func (x *DocumentSchema_EntityType) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_documentai_v1_document_schema_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocumentSchema_EntityType.ProtoReflect.Descriptor instead.
func (*DocumentSchema_EntityType) Descriptor() ([]byte, []int) {
	return file_google_cloud_documentai_v1_document_schema_proto_rawDescGZIP(), []int{0, 0}
}

func (m *DocumentSchema_EntityType) GetValueSource() isDocumentSchema_EntityType_ValueSource {
	if m != nil {
		return m.ValueSource
	}
	return nil
}

func (x *DocumentSchema_EntityType) GetEnumValues() *DocumentSchema_EntityType_EnumValues {
	if x, ok := x.GetValueSource().(*DocumentSchema_EntityType_EnumValues_); ok {
		return x.EnumValues
	}
	return nil
}

func (x *DocumentSchema_EntityType) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *DocumentSchema_EntityType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DocumentSchema_EntityType) GetBaseTypes() []string {
	if x != nil {
		return x.BaseTypes
	}
	return nil
}

func (x *DocumentSchema_EntityType) GetProperties() []*DocumentSchema_EntityType_Property {
	if x != nil {
		return x.Properties
	}
	return nil
}

type isDocumentSchema_EntityType_ValueSource interface {
	isDocumentSchema_EntityType_ValueSource()
}

type DocumentSchema_EntityType_EnumValues_ struct {
	// If specified, lists all the possible values for this entity.  This
	// should not be more than a handful of values.  If the number of values
	// is >10 or could change frequently use the `EntityType.value_ontology`
	// field and specify a list of all possible values in a value ontology
	// file.
	EnumValues *DocumentSchema_EntityType_EnumValues `protobuf:"bytes,14,opt,name=enum_values,json=enumValues,proto3,oneof"`
}

func (*DocumentSchema_EntityType_EnumValues_) isDocumentSchema_EntityType_ValueSource() {}

// Metadata for global schema behavior.
type DocumentSchema_Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If true, a `document` entity type can be applied to subdocument
	// (splitting). Otherwise, it can only be applied to the entire document
	// (classification).
	DocumentSplitter bool `protobuf:"varint,1,opt,name=document_splitter,json=documentSplitter,proto3" json:"document_splitter,omitempty"`
	// If true, on a given page, there can be multiple `document` annotations
	// covering it.
	DocumentAllowMultipleLabels bool `protobuf:"varint,2,opt,name=document_allow_multiple_labels,json=documentAllowMultipleLabels,proto3" json:"document_allow_multiple_labels,omitempty"`
	// If set, all the nested entities must be prefixed with the parents.
	PrefixedNamingOnProperties bool `protobuf:"varint,6,opt,name=prefixed_naming_on_properties,json=prefixedNamingOnProperties,proto3" json:"prefixed_naming_on_properties,omitempty"`
	// If set, we will skip the naming format validation in the schema. So the
	// string values in `DocumentSchema.EntityType.name` and
	// `DocumentSchema.EntityType.Property.name` will not be checked.
	SkipNamingValidation bool `protobuf:"varint,7,opt,name=skip_naming_validation,json=skipNamingValidation,proto3" json:"skip_naming_validation,omitempty"`
}

func (x *DocumentSchema_Metadata) Reset() {
	*x = DocumentSchema_Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_documentai_v1_document_schema_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocumentSchema_Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocumentSchema_Metadata) ProtoMessage() {}

func (x *DocumentSchema_Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_documentai_v1_document_schema_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocumentSchema_Metadata.ProtoReflect.Descriptor instead.
func (*DocumentSchema_Metadata) Descriptor() ([]byte, []int) {
	return file_google_cloud_documentai_v1_document_schema_proto_rawDescGZIP(), []int{0, 1}
}

func (x *DocumentSchema_Metadata) GetDocumentSplitter() bool {
	if x != nil {
		return x.DocumentSplitter
	}
	return false
}

func (x *DocumentSchema_Metadata) GetDocumentAllowMultipleLabels() bool {
	if x != nil {
		return x.DocumentAllowMultipleLabels
	}
	return false
}

func (x *DocumentSchema_Metadata) GetPrefixedNamingOnProperties() bool {
	if x != nil {
		return x.PrefixedNamingOnProperties
	}
	return false
}

func (x *DocumentSchema_Metadata) GetSkipNamingValidation() bool {
	if x != nil {
		return x.SkipNamingValidation
	}
	return false
}

// Defines the a list of enum values.
type DocumentSchema_EntityType_EnumValues struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The individual values that this enum values type can include.
	Values []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *DocumentSchema_EntityType_EnumValues) Reset() {
	*x = DocumentSchema_EntityType_EnumValues{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_documentai_v1_document_schema_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocumentSchema_EntityType_EnumValues) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocumentSchema_EntityType_EnumValues) ProtoMessage() {}

func (x *DocumentSchema_EntityType_EnumValues) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_documentai_v1_document_schema_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocumentSchema_EntityType_EnumValues.ProtoReflect.Descriptor instead.
func (*DocumentSchema_EntityType_EnumValues) Descriptor() ([]byte, []int) {
	return file_google_cloud_documentai_v1_document_schema_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *DocumentSchema_EntityType_EnumValues) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

// Defines properties that can be part of the entity type.
type DocumentSchema_EntityType_Property struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the property.  Follows the same guidelines as the
	// EntityType name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// User defined name for the property.
	DisplayName string `protobuf:"bytes,6,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// A reference to the value type of the property.  This type is subject
	// to the same conventions as the `Entity.base_types` field.
	ValueType string `protobuf:"bytes,2,opt,name=value_type,json=valueType,proto3" json:"value_type,omitempty"`
	// Occurrence type limits the number of instances an entity type appears
	// in the document.
	OccurrenceType DocumentSchema_EntityType_Property_OccurrenceType `protobuf:"varint,3,opt,name=occurrence_type,json=occurrenceType,proto3,enum=google.cloud.documentai.v1.DocumentSchema_EntityType_Property_OccurrenceType" json:"occurrence_type,omitempty"`
}

func (x *DocumentSchema_EntityType_Property) Reset() {
	*x = DocumentSchema_EntityType_Property{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_documentai_v1_document_schema_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocumentSchema_EntityType_Property) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocumentSchema_EntityType_Property) ProtoMessage() {}

func (x *DocumentSchema_EntityType_Property) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_documentai_v1_document_schema_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocumentSchema_EntityType_Property.ProtoReflect.Descriptor instead.
func (*DocumentSchema_EntityType_Property) Descriptor() ([]byte, []int) {
	return file_google_cloud_documentai_v1_document_schema_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *DocumentSchema_EntityType_Property) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DocumentSchema_EntityType_Property) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *DocumentSchema_EntityType_Property) GetValueType() string {
	if x != nil {
		return x.ValueType
	}
	return ""
}

func (x *DocumentSchema_EntityType_Property) GetOccurrenceType() DocumentSchema_EntityType_Property_OccurrenceType {
	if x != nil {
		return x.OccurrenceType
	}
	return DocumentSchema_EntityType_Property_OCCURRENCE_TYPE_UNSPECIFIED
}

var File_google_cloud_documentai_v1_document_schema_proto protoreflect.FileDescriptor

var file_google_cloud_documentai_v1_document_schema_proto_rawDesc = []byte{
	0x0a, 0x30, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x22, 0xbb,
	0x09, 0x0a, 0x0e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x58, 0x0a, 0x0c, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x12, 0x4f, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x1a, 0xc0, 0x05, 0x0a, 0x0a, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x63, 0x0a, 0x0b, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x45, 0x6e, 0x75,
	0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x48, 0x00, 0x52, 0x0a, 0x65, 0x6e, 0x75, 0x6d, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x62, 0x61, 0x73, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x09, 0x62, 0x61, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x5e, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x3e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52,
	0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x1a, 0x24, 0x0a, 0x0a, 0x45,
	0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x1a, 0xe0, 0x02, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x76, 0x0a, 0x0f, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x2e, 0x4f, 0x63,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0e, 0x6f, 0x63,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x85, 0x01, 0x0a,
	0x0e, 0x4f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1f, 0x0a, 0x1b, 0x4f, 0x43, 0x43, 0x55, 0x52, 0x52, 0x45, 0x4e, 0x43, 0x45, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x11, 0x0a, 0x0d, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x5f, 0x4f, 0x4e, 0x43,
	0x45, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x5f,
	0x4d, 0x55, 0x4c, 0x54, 0x49, 0x50, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x45,
	0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x5f, 0x4f, 0x4e, 0x43, 0x45, 0x10, 0x03, 0x12, 0x15, 0x0a,
	0x11, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x5f, 0x4d, 0x55, 0x4c, 0x54, 0x49, 0x50,
	0x4c, 0x45, 0x10, 0x04, 0x42, 0x0e, 0x0a, 0x0c, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x1a, 0xf5, 0x01, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x70,
	0x6c, 0x69, 0x74, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x64, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x74, 0x65, 0x72, 0x12, 0x43,
	0x0a, 0x1e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1b, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x12, 0x41, 0x0a, 0x1d, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x64, 0x5f,
	0x6e, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x69, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1a, 0x70, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x4f, 0x6e, 0x50, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x16, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x6e,
	0x61, 0x6d, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x73, 0x6b, 0x69, 0x70, 0x4e, 0x61, 0x6d, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0xd6, 0x01, 0x0a,
	0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x42,
	0x18, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x69, 0x44, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x50, 0x01, 0x5a, 0x3e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x76,
	0x31, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x69, 0x70, 0x62, 0x3b, 0x64,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x69, 0x70, 0x62, 0xaa, 0x02, 0x1a, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x41, 0x49, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x41, 0x49, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x41,
	0x49, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_documentai_v1_document_schema_proto_rawDescOnce sync.Once
	file_google_cloud_documentai_v1_document_schema_proto_rawDescData = file_google_cloud_documentai_v1_document_schema_proto_rawDesc
)

func file_google_cloud_documentai_v1_document_schema_proto_rawDescGZIP() []byte {
	file_google_cloud_documentai_v1_document_schema_proto_rawDescOnce.Do(func() {
		file_google_cloud_documentai_v1_document_schema_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_documentai_v1_document_schema_proto_rawDescData)
	})
	return file_google_cloud_documentai_v1_document_schema_proto_rawDescData
}

var file_google_cloud_documentai_v1_document_schema_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_documentai_v1_document_schema_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_cloud_documentai_v1_document_schema_proto_goTypes = []any{
	(DocumentSchema_EntityType_Property_OccurrenceType)(0), // 0: google.cloud.documentai.v1.DocumentSchema.EntityType.Property.OccurrenceType
	(*DocumentSchema)(nil),                                 // 1: google.cloud.documentai.v1.DocumentSchema
	(*DocumentSchema_EntityType)(nil),                      // 2: google.cloud.documentai.v1.DocumentSchema.EntityType
	(*DocumentSchema_Metadata)(nil),                        // 3: google.cloud.documentai.v1.DocumentSchema.Metadata
	(*DocumentSchema_EntityType_EnumValues)(nil),           // 4: google.cloud.documentai.v1.DocumentSchema.EntityType.EnumValues
	(*DocumentSchema_EntityType_Property)(nil),             // 5: google.cloud.documentai.v1.DocumentSchema.EntityType.Property
}
var file_google_cloud_documentai_v1_document_schema_proto_depIdxs = []int32{
	2, // 0: google.cloud.documentai.v1.DocumentSchema.entity_types:type_name -> google.cloud.documentai.v1.DocumentSchema.EntityType
	3, // 1: google.cloud.documentai.v1.DocumentSchema.metadata:type_name -> google.cloud.documentai.v1.DocumentSchema.Metadata
	4, // 2: google.cloud.documentai.v1.DocumentSchema.EntityType.enum_values:type_name -> google.cloud.documentai.v1.DocumentSchema.EntityType.EnumValues
	5, // 3: google.cloud.documentai.v1.DocumentSchema.EntityType.properties:type_name -> google.cloud.documentai.v1.DocumentSchema.EntityType.Property
	0, // 4: google.cloud.documentai.v1.DocumentSchema.EntityType.Property.occurrence_type:type_name -> google.cloud.documentai.v1.DocumentSchema.EntityType.Property.OccurrenceType
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_google_cloud_documentai_v1_document_schema_proto_init() }
func file_google_cloud_documentai_v1_document_schema_proto_init() {
	if File_google_cloud_documentai_v1_document_schema_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_documentai_v1_document_schema_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*DocumentSchema); i {
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
		file_google_cloud_documentai_v1_document_schema_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*DocumentSchema_EntityType); i {
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
		file_google_cloud_documentai_v1_document_schema_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*DocumentSchema_Metadata); i {
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
		file_google_cloud_documentai_v1_document_schema_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*DocumentSchema_EntityType_EnumValues); i {
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
		file_google_cloud_documentai_v1_document_schema_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*DocumentSchema_EntityType_Property); i {
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
	file_google_cloud_documentai_v1_document_schema_proto_msgTypes[1].OneofWrappers = []any{
		(*DocumentSchema_EntityType_EnumValues_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_documentai_v1_document_schema_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_documentai_v1_document_schema_proto_goTypes,
		DependencyIndexes: file_google_cloud_documentai_v1_document_schema_proto_depIdxs,
		EnumInfos:         file_google_cloud_documentai_v1_document_schema_proto_enumTypes,
		MessageInfos:      file_google_cloud_documentai_v1_document_schema_proto_msgTypes,
	}.Build()
	File_google_cloud_documentai_v1_document_schema_proto = out.File
	file_google_cloud_documentai_v1_document_schema_proto_rawDesc = nil
	file_google_cloud_documentai_v1_document_schema_proto_goTypes = nil
	file_google_cloud_documentai_v1_document_schema_proto_depIdxs = nil
}
