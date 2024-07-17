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
// source: google/cloud/securitycenter/v1/asset.proto

package securitycenterpb

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

// Security Command Center representation of a Google Cloud
// resource.
//
// The Asset is a Security Command Center resource that captures information
// about a single Google Cloud resource. All modifications to an Asset are only
// within the context of Security Command Center and don't affect the referenced
// Google Cloud resource.
type Asset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The relative resource name of this asset. See:
	// https://cloud.google.com/apis/design/resource_names#relative_resource_name
	// Example:
	// "organizations/{organization_id}/assets/{asset_id}".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Security Command Center managed properties. These properties are managed by
	// Security Command Center and cannot be modified by the user.
	SecurityCenterProperties *Asset_SecurityCenterProperties `protobuf:"bytes,2,opt,name=security_center_properties,json=securityCenterProperties,proto3" json:"security_center_properties,omitempty"`
	// Resource managed properties. These properties are managed and defined by
	// the Google Cloud resource and cannot be modified by the user.
	ResourceProperties map[string]*structpb.Value `protobuf:"bytes,7,rep,name=resource_properties,json=resourceProperties,proto3" json:"resource_properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// User specified security marks. These marks are entirely managed by the user
	// and come from the SecurityMarks resource that belongs to the asset.
	SecurityMarks *SecurityMarks `protobuf:"bytes,8,opt,name=security_marks,json=securityMarks,proto3" json:"security_marks,omitempty"`
	// The time at which the asset was created in Security Command Center.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The time at which the asset was last updated or added in Cloud SCC.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Cloud IAM Policy information associated with the Google Cloud resource
	// described by the Security Command Center asset. This information is managed
	// and defined by the Google Cloud resource and cannot be modified by the
	// user.
	IamPolicy *Asset_IamPolicy `protobuf:"bytes,11,opt,name=iam_policy,json=iamPolicy,proto3" json:"iam_policy,omitempty"`
	// The canonical name of the resource. It's either
	// "organizations/{organization_id}/assets/{asset_id}",
	// "folders/{folder_id}/assets/{asset_id}" or
	// "projects/{project_number}/assets/{asset_id}", depending on the closest CRM
	// ancestor of the resource.
	CanonicalName string `protobuf:"bytes,13,opt,name=canonical_name,json=canonicalName,proto3" json:"canonical_name,omitempty"`
}

func (x *Asset) Reset() {
	*x = Asset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_securitycenter_v1_asset_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Asset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Asset) ProtoMessage() {}

func (x *Asset) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v1_asset_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Asset.ProtoReflect.Descriptor instead.
func (*Asset) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v1_asset_proto_rawDescGZIP(), []int{0}
}

func (x *Asset) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Asset) GetSecurityCenterProperties() *Asset_SecurityCenterProperties {
	if x != nil {
		return x.SecurityCenterProperties
	}
	return nil
}

func (x *Asset) GetResourceProperties() map[string]*structpb.Value {
	if x != nil {
		return x.ResourceProperties
	}
	return nil
}

func (x *Asset) GetSecurityMarks() *SecurityMarks {
	if x != nil {
		return x.SecurityMarks
	}
	return nil
}

func (x *Asset) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Asset) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Asset) GetIamPolicy() *Asset_IamPolicy {
	if x != nil {
		return x.IamPolicy
	}
	return nil
}

func (x *Asset) GetCanonicalName() string {
	if x != nil {
		return x.CanonicalName
	}
	return ""
}

// Security Command Center managed properties. These properties are managed by
// Security Command Center and cannot be modified by the user.
type Asset_SecurityCenterProperties struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The full resource name of the Google Cloud resource this asset
	// represents. This field is immutable after create time. See:
	// https://cloud.google.com/apis/design/resource_names#full_resource_name
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The type of the Google Cloud resource. Examples include: APPLICATION,
	// PROJECT, and ORGANIZATION. This is a case insensitive field defined by
	// Security Command Center and/or the producer of the resource and is
	// immutable after create time.
	ResourceType string `protobuf:"bytes,2,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	// The full resource name of the immediate parent of the resource. See:
	// https://cloud.google.com/apis/design/resource_names#full_resource_name
	ResourceParent string `protobuf:"bytes,3,opt,name=resource_parent,json=resourceParent,proto3" json:"resource_parent,omitempty"`
	// The full resource name of the project the resource belongs to. See:
	// https://cloud.google.com/apis/design/resource_names#full_resource_name
	ResourceProject string `protobuf:"bytes,4,opt,name=resource_project,json=resourceProject,proto3" json:"resource_project,omitempty"`
	// Owners of the Google Cloud resource.
	ResourceOwners []string `protobuf:"bytes,5,rep,name=resource_owners,json=resourceOwners,proto3" json:"resource_owners,omitempty"`
	// The user defined display name for this resource.
	ResourceDisplayName string `protobuf:"bytes,6,opt,name=resource_display_name,json=resourceDisplayName,proto3" json:"resource_display_name,omitempty"`
	// The user defined display name for the parent of this resource.
	ResourceParentDisplayName string `protobuf:"bytes,7,opt,name=resource_parent_display_name,json=resourceParentDisplayName,proto3" json:"resource_parent_display_name,omitempty"`
	// The user defined display name for the project of this resource.
	ResourceProjectDisplayName string `protobuf:"bytes,8,opt,name=resource_project_display_name,json=resourceProjectDisplayName,proto3" json:"resource_project_display_name,omitempty"`
	// Contains a Folder message for each folder in the assets ancestry.
	// The first folder is the deepest nested folder, and the last folder is the
	// folder directly under the Organization.
	Folders []*Folder `protobuf:"bytes,10,rep,name=folders,proto3" json:"folders,omitempty"`
}

func (x *Asset_SecurityCenterProperties) Reset() {
	*x = Asset_SecurityCenterProperties{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_securitycenter_v1_asset_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Asset_SecurityCenterProperties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Asset_SecurityCenterProperties) ProtoMessage() {}

func (x *Asset_SecurityCenterProperties) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v1_asset_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Asset_SecurityCenterProperties.ProtoReflect.Descriptor instead.
func (*Asset_SecurityCenterProperties) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v1_asset_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Asset_SecurityCenterProperties) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *Asset_SecurityCenterProperties) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *Asset_SecurityCenterProperties) GetResourceParent() string {
	if x != nil {
		return x.ResourceParent
	}
	return ""
}

func (x *Asset_SecurityCenterProperties) GetResourceProject() string {
	if x != nil {
		return x.ResourceProject
	}
	return ""
}

func (x *Asset_SecurityCenterProperties) GetResourceOwners() []string {
	if x != nil {
		return x.ResourceOwners
	}
	return nil
}

func (x *Asset_SecurityCenterProperties) GetResourceDisplayName() string {
	if x != nil {
		return x.ResourceDisplayName
	}
	return ""
}

func (x *Asset_SecurityCenterProperties) GetResourceParentDisplayName() string {
	if x != nil {
		return x.ResourceParentDisplayName
	}
	return ""
}

func (x *Asset_SecurityCenterProperties) GetResourceProjectDisplayName() string {
	if x != nil {
		return x.ResourceProjectDisplayName
	}
	return ""
}

func (x *Asset_SecurityCenterProperties) GetFolders() []*Folder {
	if x != nil {
		return x.Folders
	}
	return nil
}

// Cloud IAM Policy information associated with the Google Cloud resource
// described by the Security Command Center asset. This information is managed
// and defined by the Google Cloud resource and cannot be modified by the
// user.
type Asset_IamPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The JSON representation of the Policy associated with the asset.
	// See https://cloud.google.com/iam/reference/rest/v1/Policy for format
	// details.
	PolicyBlob string `protobuf:"bytes,1,opt,name=policy_blob,json=policyBlob,proto3" json:"policy_blob,omitempty"`
}

func (x *Asset_IamPolicy) Reset() {
	*x = Asset_IamPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_securitycenter_v1_asset_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Asset_IamPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Asset_IamPolicy) ProtoMessage() {}

func (x *Asset_IamPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v1_asset_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Asset_IamPolicy.ProtoReflect.Descriptor instead.
func (*Asset_IamPolicy) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v1_asset_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Asset_IamPolicy) GetPolicyBlob() string {
	if x != nil {
		return x.PolicyBlob
	}
	return ""
}

var File_google_cloud_securitycenter_v1_asset_proto protoreflect.FileDescriptor

var file_google_cloud_securitycenter_v1_asset_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x6d, 0x61,
	0x72, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe7, 0x0a, 0x0a, 0x05, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x7c, 0x0a, 0x1a, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x73, 0x73, 0x65,
	0x74, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x52, 0x18, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x69, 0x65, 0x73, 0x12, 0x6e, 0x0a, 0x13, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x3d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x12, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x69, 0x65, 0x73, 0x12, 0x54, 0x0a, 0x0e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x4d, 0x61, 0x72, 0x6b, 0x73, 0x52, 0x0d, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x4d, 0x61, 0x72, 0x6b, 0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x4e, 0x0a, 0x0a, 0x69, 0x61, 0x6d, 0x5f, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x2e,
	0x49, 0x61, 0x6d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x09, 0x69, 0x61, 0x6d, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x61, 0x6e, 0x6f, 0x6e, 0x69, 0x63, 0x61,
	0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x61,
	0x6e, 0x6f, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0xdb, 0x03, 0x0a, 0x18,
	0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a,
	0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x12,
	0x32, 0x0a, 0x15, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x3f, 0x0a, 0x1c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x19, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x1d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1a, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x07, 0x66, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x52, 0x07, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x1a, 0x2c, 0x0a, 0x09, 0x49, 0x61, 0x6d,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x5f, 0x62, 0x6c, 0x6f, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x42, 0x6c, 0x6f, 0x62, 0x1a, 0x5d, 0x0a, 0x17, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0xa9, 0x01, 0xea, 0x41, 0xa5, 0x01, 0x0a, 0x23, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x12, 0x2b, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d,
	0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x7b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x7d, 0x12,
	0x1f, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x7d, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x7b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x7d,
	0x12, 0x21, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x7d, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x7b, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x7d, 0x2a, 0x06, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x32, 0x05, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x42, 0xd8, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x4a, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f,
	0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x62, 0x3b, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x62, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x53, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_securitycenter_v1_asset_proto_rawDescOnce sync.Once
	file_google_cloud_securitycenter_v1_asset_proto_rawDescData = file_google_cloud_securitycenter_v1_asset_proto_rawDesc
)

func file_google_cloud_securitycenter_v1_asset_proto_rawDescGZIP() []byte {
	file_google_cloud_securitycenter_v1_asset_proto_rawDescOnce.Do(func() {
		file_google_cloud_securitycenter_v1_asset_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_securitycenter_v1_asset_proto_rawDescData)
	})
	return file_google_cloud_securitycenter_v1_asset_proto_rawDescData
}

var file_google_cloud_securitycenter_v1_asset_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_cloud_securitycenter_v1_asset_proto_goTypes = []any{
	(*Asset)(nil),                          // 0: google.cloud.securitycenter.v1.Asset
	(*Asset_SecurityCenterProperties)(nil), // 1: google.cloud.securitycenter.v1.Asset.SecurityCenterProperties
	(*Asset_IamPolicy)(nil),                // 2: google.cloud.securitycenter.v1.Asset.IamPolicy
	nil,                                    // 3: google.cloud.securitycenter.v1.Asset.ResourcePropertiesEntry
	(*SecurityMarks)(nil),                  // 4: google.cloud.securitycenter.v1.SecurityMarks
	(*timestamppb.Timestamp)(nil),          // 5: google.protobuf.Timestamp
	(*Folder)(nil),                         // 6: google.cloud.securitycenter.v1.Folder
	(*structpb.Value)(nil),                 // 7: google.protobuf.Value
}
var file_google_cloud_securitycenter_v1_asset_proto_depIdxs = []int32{
	1, // 0: google.cloud.securitycenter.v1.Asset.security_center_properties:type_name -> google.cloud.securitycenter.v1.Asset.SecurityCenterProperties
	3, // 1: google.cloud.securitycenter.v1.Asset.resource_properties:type_name -> google.cloud.securitycenter.v1.Asset.ResourcePropertiesEntry
	4, // 2: google.cloud.securitycenter.v1.Asset.security_marks:type_name -> google.cloud.securitycenter.v1.SecurityMarks
	5, // 3: google.cloud.securitycenter.v1.Asset.create_time:type_name -> google.protobuf.Timestamp
	5, // 4: google.cloud.securitycenter.v1.Asset.update_time:type_name -> google.protobuf.Timestamp
	2, // 5: google.cloud.securitycenter.v1.Asset.iam_policy:type_name -> google.cloud.securitycenter.v1.Asset.IamPolicy
	6, // 6: google.cloud.securitycenter.v1.Asset.SecurityCenterProperties.folders:type_name -> google.cloud.securitycenter.v1.Folder
	7, // 7: google.cloud.securitycenter.v1.Asset.ResourcePropertiesEntry.value:type_name -> google.protobuf.Value
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_google_cloud_securitycenter_v1_asset_proto_init() }
func file_google_cloud_securitycenter_v1_asset_proto_init() {
	if File_google_cloud_securitycenter_v1_asset_proto != nil {
		return
	}
	file_google_cloud_securitycenter_v1_folder_proto_init()
	file_google_cloud_securitycenter_v1_security_marks_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_securitycenter_v1_asset_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Asset); i {
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
		file_google_cloud_securitycenter_v1_asset_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Asset_SecurityCenterProperties); i {
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
		file_google_cloud_securitycenter_v1_asset_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Asset_IamPolicy); i {
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
			RawDescriptor: file_google_cloud_securitycenter_v1_asset_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_securitycenter_v1_asset_proto_goTypes,
		DependencyIndexes: file_google_cloud_securitycenter_v1_asset_proto_depIdxs,
		MessageInfos:      file_google_cloud_securitycenter_v1_asset_proto_msgTypes,
	}.Build()
	File_google_cloud_securitycenter_v1_asset_proto = out.File
	file_google_cloud_securitycenter_v1_asset_proto_rawDesc = nil
	file_google_cloud_securitycenter_v1_asset_proto_goTypes = nil
	file_google_cloud_securitycenter_v1_asset_proto_depIdxs = nil
}
