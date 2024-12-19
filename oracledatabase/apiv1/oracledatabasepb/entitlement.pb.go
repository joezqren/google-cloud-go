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
// source: google/cloud/oracledatabase/v1/entitlement.proto

package oracledatabasepb

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

// The various lifecycle states of the subscription.
type Entitlement_State int32

const (
	// Default unspecified value.
	Entitlement_STATE_UNSPECIFIED Entitlement_State = 0
	// Account not linked.
	Entitlement_ACCOUNT_NOT_LINKED Entitlement_State = 1
	// Account is linked but not active.
	Entitlement_ACCOUNT_NOT_ACTIVE Entitlement_State = 2
	// Entitlement and Account are active.
	Entitlement_ACTIVE Entitlement_State = 3
)

// Enum value maps for Entitlement_State.
var (
	Entitlement_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "ACCOUNT_NOT_LINKED",
		2: "ACCOUNT_NOT_ACTIVE",
		3: "ACTIVE",
	}
	Entitlement_State_value = map[string]int32{
		"STATE_UNSPECIFIED":  0,
		"ACCOUNT_NOT_LINKED": 1,
		"ACCOUNT_NOT_ACTIVE": 2,
		"ACTIVE":             3,
	}
)

func (x Entitlement_State) Enum() *Entitlement_State {
	p := new(Entitlement_State)
	*p = x
	return p
}

func (x Entitlement_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Entitlement_State) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_oracledatabase_v1_entitlement_proto_enumTypes[0].Descriptor()
}

func (Entitlement_State) Type() protoreflect.EnumType {
	return &file_google_cloud_oracledatabase_v1_entitlement_proto_enumTypes[0]
}

func (x Entitlement_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Entitlement_State.Descriptor instead.
func (Entitlement_State) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_oracledatabase_v1_entitlement_proto_rawDescGZIP(), []int{0, 0}
}

// Details of the Entitlement resource.
type Entitlement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier. The name of the Entitlement resource with the format:
	// projects/{project}/locations/{region}/entitlements/{entitlement}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Details of the OCI Cloud Account.
	CloudAccountDetails *CloudAccountDetails `protobuf:"bytes,2,opt,name=cloud_account_details,json=cloudAccountDetails,proto3" json:"cloud_account_details,omitempty"`
	// Output only. Google Cloud Marketplace order ID (aka entitlement ID)
	EntitlementId string `protobuf:"bytes,3,opt,name=entitlement_id,json=entitlementId,proto3" json:"entitlement_id,omitempty"`
	// Output only. Entitlement State.
	State Entitlement_State `protobuf:"varint,4,opt,name=state,proto3,enum=google.cloud.oracledatabase.v1.Entitlement_State" json:"state,omitempty"`
}

func (x *Entitlement) Reset() {
	*x = Entitlement{}
	mi := &file_google_cloud_oracledatabase_v1_entitlement_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Entitlement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entitlement) ProtoMessage() {}

func (x *Entitlement) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_oracledatabase_v1_entitlement_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entitlement.ProtoReflect.Descriptor instead.
func (*Entitlement) Descriptor() ([]byte, []int) {
	return file_google_cloud_oracledatabase_v1_entitlement_proto_rawDescGZIP(), []int{0}
}

func (x *Entitlement) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Entitlement) GetCloudAccountDetails() *CloudAccountDetails {
	if x != nil {
		return x.CloudAccountDetails
	}
	return nil
}

func (x *Entitlement) GetEntitlementId() string {
	if x != nil {
		return x.EntitlementId
	}
	return ""
}

func (x *Entitlement) GetState() Entitlement_State {
	if x != nil {
		return x.State
	}
	return Entitlement_STATE_UNSPECIFIED
}

// Details of the OCI Cloud Account.
type CloudAccountDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. OCI account name.
	CloudAccount string `protobuf:"bytes,1,opt,name=cloud_account,json=cloudAccount,proto3" json:"cloud_account,omitempty"`
	// Output only. OCI account home region.
	CloudAccountHomeRegion string `protobuf:"bytes,2,opt,name=cloud_account_home_region,json=cloudAccountHomeRegion,proto3" json:"cloud_account_home_region,omitempty"`
	// Output only. URL to link an existing account.
	LinkExistingAccountUri *string `protobuf:"bytes,3,opt,name=link_existing_account_uri,json=linkExistingAccountUri,proto3,oneof" json:"link_existing_account_uri,omitempty"`
	// Output only. URL to create a new account and link.
	AccountCreationUri *string `protobuf:"bytes,4,opt,name=account_creation_uri,json=accountCreationUri,proto3,oneof" json:"account_creation_uri,omitempty"`
}

func (x *CloudAccountDetails) Reset() {
	*x = CloudAccountDetails{}
	mi := &file_google_cloud_oracledatabase_v1_entitlement_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CloudAccountDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudAccountDetails) ProtoMessage() {}

func (x *CloudAccountDetails) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_oracledatabase_v1_entitlement_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudAccountDetails.ProtoReflect.Descriptor instead.
func (*CloudAccountDetails) Descriptor() ([]byte, []int) {
	return file_google_cloud_oracledatabase_v1_entitlement_proto_rawDescGZIP(), []int{1}
}

func (x *CloudAccountDetails) GetCloudAccount() string {
	if x != nil {
		return x.CloudAccount
	}
	return ""
}

func (x *CloudAccountDetails) GetCloudAccountHomeRegion() string {
	if x != nil {
		return x.CloudAccountHomeRegion
	}
	return ""
}

func (x *CloudAccountDetails) GetLinkExistingAccountUri() string {
	if x != nil && x.LinkExistingAccountUri != nil {
		return *x.LinkExistingAccountUri
	}
	return ""
}

func (x *CloudAccountDetails) GetAccountCreationUri() string {
	if x != nil && x.AccountCreationUri != nil {
		return *x.AccountCreationUri
	}
	return ""
}

var File_google_cloud_oracledatabase_v1_entitlement_proto protoreflect.FileDescriptor

var file_google_cloud_oracledatabase_v1_entitlement_proto_rawDesc = []byte{
	0x0a, 0x30, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f,
	0x72, 0x61, 0x63, 0x6c, 0x65, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf6,
	0x03, 0x0a, 0x0b, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x17,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x08, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x67, 0x0a, 0x15, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x13, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x12, 0x2a, 0x0a, 0x0e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0d, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x4c, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x31, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x72, 0x61, 0x63, 0x6c,
	0x65, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x5a, 0x0a, 0x05, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x41, 0x43,
	0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x4c, 0x49, 0x4e, 0x4b, 0x45, 0x44,
	0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x4e, 0x4f,
	0x54, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43,
	0x54, 0x49, 0x56, 0x45, 0x10, 0x03, 0x3a, 0x8e, 0x01, 0xea, 0x41, 0x8a, 0x01, 0x0a, 0x29, 0x6f,
	0x72, 0x61, 0x63, 0x6c, 0x65, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x42, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x7d, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x7d, 0x2a, 0x0c, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x0b, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0xb7, 0x02, 0x0a, 0x13, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12,
	0x28, 0x0a, 0x0d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0c, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3e, 0x0a, 0x19, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x68, 0x6f, 0x6d, 0x65, 0x5f,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x16, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x48,
	0x6f, 0x6d, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x43, 0x0a, 0x19, 0x6c, 0x69, 0x6e,
	0x6b, 0x5f, 0x65, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x48, 0x00, 0x52, 0x16, 0x6c, 0x69, 0x6e, 0x6b, 0x45, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x72, 0x69, 0x88, 0x01, 0x01, 0x12, 0x3a,
	0x0a, 0x14, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x48, 0x01, 0x52, 0x12, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x55, 0x72, 0x69, 0x88, 0x01, 0x01, 0x42, 0x1c, 0x0a, 0x1a, 0x5f, 0x6c,
	0x69, 0x6e, 0x6b, 0x5f, 0x65, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x75, 0x72, 0x69, 0x42, 0x17, 0x0a, 0x15, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x72,
	0x69, 0x42, 0xea, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x6f, 0x2f, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x70, 0x62, 0x3b, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x70, 0x62, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x21, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x4f, 0x72, 0x61, 0x63,
	0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_oracledatabase_v1_entitlement_proto_rawDescOnce sync.Once
	file_google_cloud_oracledatabase_v1_entitlement_proto_rawDescData = file_google_cloud_oracledatabase_v1_entitlement_proto_rawDesc
)

func file_google_cloud_oracledatabase_v1_entitlement_proto_rawDescGZIP() []byte {
	file_google_cloud_oracledatabase_v1_entitlement_proto_rawDescOnce.Do(func() {
		file_google_cloud_oracledatabase_v1_entitlement_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_oracledatabase_v1_entitlement_proto_rawDescData)
	})
	return file_google_cloud_oracledatabase_v1_entitlement_proto_rawDescData
}

var file_google_cloud_oracledatabase_v1_entitlement_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_oracledatabase_v1_entitlement_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_oracledatabase_v1_entitlement_proto_goTypes = []any{
	(Entitlement_State)(0),      // 0: google.cloud.oracledatabase.v1.Entitlement.State
	(*Entitlement)(nil),         // 1: google.cloud.oracledatabase.v1.Entitlement
	(*CloudAccountDetails)(nil), // 2: google.cloud.oracledatabase.v1.CloudAccountDetails
}
var file_google_cloud_oracledatabase_v1_entitlement_proto_depIdxs = []int32{
	2, // 0: google.cloud.oracledatabase.v1.Entitlement.cloud_account_details:type_name -> google.cloud.oracledatabase.v1.CloudAccountDetails
	0, // 1: google.cloud.oracledatabase.v1.Entitlement.state:type_name -> google.cloud.oracledatabase.v1.Entitlement.State
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_cloud_oracledatabase_v1_entitlement_proto_init() }
func file_google_cloud_oracledatabase_v1_entitlement_proto_init() {
	if File_google_cloud_oracledatabase_v1_entitlement_proto != nil {
		return
	}
	file_google_cloud_oracledatabase_v1_entitlement_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_oracledatabase_v1_entitlement_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_oracledatabase_v1_entitlement_proto_goTypes,
		DependencyIndexes: file_google_cloud_oracledatabase_v1_entitlement_proto_depIdxs,
		EnumInfos:         file_google_cloud_oracledatabase_v1_entitlement_proto_enumTypes,
		MessageInfos:      file_google_cloud_oracledatabase_v1_entitlement_proto_msgTypes,
	}.Build()
	File_google_cloud_oracledatabase_v1_entitlement_proto = out.File
	file_google_cloud_oracledatabase_v1_entitlement_proto_rawDesc = nil
	file_google_cloud_oracledatabase_v1_entitlement_proto_goTypes = nil
	file_google_cloud_oracledatabase_v1_entitlement_proto_depIdxs = nil
}
