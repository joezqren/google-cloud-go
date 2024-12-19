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
// source: google/cloud/securitycenter/v1/contact_details.proto

package securitycenterpb

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

// Details about specific contacts
type ContactDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of contacts
	Contacts []*Contact `protobuf:"bytes,1,rep,name=contacts,proto3" json:"contacts,omitempty"`
}

func (x *ContactDetails) Reset() {
	*x = ContactDetails{}
	mi := &file_google_cloud_securitycenter_v1_contact_details_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ContactDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactDetails) ProtoMessage() {}

func (x *ContactDetails) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v1_contact_details_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactDetails.ProtoReflect.Descriptor instead.
func (*ContactDetails) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v1_contact_details_proto_rawDescGZIP(), []int{0}
}

func (x *ContactDetails) GetContacts() []*Contact {
	if x != nil {
		return x.Contacts
	}
	return nil
}

// The email address of a contact.
type Contact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An email address. For example, "`person123@company.com`".
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *Contact) Reset() {
	*x = Contact{}
	mi := &file_google_cloud_securitycenter_v1_contact_details_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Contact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contact) ProtoMessage() {}

func (x *Contact) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v1_contact_details_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contact.ProtoReflect.Descriptor instead.
func (*Contact) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v1_contact_details_proto_rawDescGZIP(), []int{1}
}

func (x *Contact) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_google_cloud_securitycenter_v1_contact_details_proto protoreflect.FileDescriptor

var file_google_cloud_securitycenter_v1_contact_details_proto_rawDesc = []byte{
	0x0a, 0x34, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x55, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x43, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x22, 0x1f, 0x0a,
	0x07, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0xed,
	0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x13, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x6f, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x62, 0x3b, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x62, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x21, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x53, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_securitycenter_v1_contact_details_proto_rawDescOnce sync.Once
	file_google_cloud_securitycenter_v1_contact_details_proto_rawDescData = file_google_cloud_securitycenter_v1_contact_details_proto_rawDesc
)

func file_google_cloud_securitycenter_v1_contact_details_proto_rawDescGZIP() []byte {
	file_google_cloud_securitycenter_v1_contact_details_proto_rawDescOnce.Do(func() {
		file_google_cloud_securitycenter_v1_contact_details_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_securitycenter_v1_contact_details_proto_rawDescData)
	})
	return file_google_cloud_securitycenter_v1_contact_details_proto_rawDescData
}

var file_google_cloud_securitycenter_v1_contact_details_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_securitycenter_v1_contact_details_proto_goTypes = []any{
	(*ContactDetails)(nil), // 0: google.cloud.securitycenter.v1.ContactDetails
	(*Contact)(nil),        // 1: google.cloud.securitycenter.v1.Contact
}
var file_google_cloud_securitycenter_v1_contact_details_proto_depIdxs = []int32{
	1, // 0: google.cloud.securitycenter.v1.ContactDetails.contacts:type_name -> google.cloud.securitycenter.v1.Contact
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_securitycenter_v1_contact_details_proto_init() }
func file_google_cloud_securitycenter_v1_contact_details_proto_init() {
	if File_google_cloud_securitycenter_v1_contact_details_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_securitycenter_v1_contact_details_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_securitycenter_v1_contact_details_proto_goTypes,
		DependencyIndexes: file_google_cloud_securitycenter_v1_contact_details_proto_depIdxs,
		MessageInfos:      file_google_cloud_securitycenter_v1_contact_details_proto_msgTypes,
	}.Build()
	File_google_cloud_securitycenter_v1_contact_details_proto = out.File
	file_google_cloud_securitycenter_v1_contact_details_proto_rawDesc = nil
	file_google_cloud_securitycenter_v1_contact_details_proto_goTypes = nil
	file_google_cloud_securitycenter_v1_contact_details_proto_depIdxs = nil
}
