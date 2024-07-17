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
// source: google/cloud/channel/v1/channel_partner_links.proto

package channelpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The level of granularity the
// [ChannelPartnerLink][google.cloud.channel.v1.ChannelPartnerLink] will
// display.
type ChannelPartnerLinkView int32

const (
	// The default / unset value.
	// The API will default to the BASIC view.
	ChannelPartnerLinkView_UNSPECIFIED ChannelPartnerLinkView = 0
	// Includes all fields except the
	// [ChannelPartnerLink.channel_partner_cloud_identity_info][google.cloud.channel.v1.ChannelPartnerLink.channel_partner_cloud_identity_info].
	ChannelPartnerLinkView_BASIC ChannelPartnerLinkView = 1
	// Includes all fields.
	ChannelPartnerLinkView_FULL ChannelPartnerLinkView = 2
)

// Enum value maps for ChannelPartnerLinkView.
var (
	ChannelPartnerLinkView_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "BASIC",
		2: "FULL",
	}
	ChannelPartnerLinkView_value = map[string]int32{
		"UNSPECIFIED": 0,
		"BASIC":       1,
		"FULL":        2,
	}
)

func (x ChannelPartnerLinkView) Enum() *ChannelPartnerLinkView {
	p := new(ChannelPartnerLinkView)
	*p = x
	return p
}

func (x ChannelPartnerLinkView) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChannelPartnerLinkView) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_channel_v1_channel_partner_links_proto_enumTypes[0].Descriptor()
}

func (ChannelPartnerLinkView) Type() protoreflect.EnumType {
	return &file_google_cloud_channel_v1_channel_partner_links_proto_enumTypes[0]
}

func (x ChannelPartnerLinkView) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChannelPartnerLinkView.Descriptor instead.
func (ChannelPartnerLinkView) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_channel_v1_channel_partner_links_proto_rawDescGZIP(), []int{0}
}

// ChannelPartnerLinkState represents state of a channel partner link.
type ChannelPartnerLinkState int32

const (
	// Not used.
	ChannelPartnerLinkState_CHANNEL_PARTNER_LINK_STATE_UNSPECIFIED ChannelPartnerLinkState = 0
	// An invitation has been sent to the reseller to create a channel partner
	// link.
	ChannelPartnerLinkState_INVITED ChannelPartnerLinkState = 1
	// Status when the reseller is active.
	ChannelPartnerLinkState_ACTIVE ChannelPartnerLinkState = 2
	// Status when the reseller has been revoked by the distributor.
	ChannelPartnerLinkState_REVOKED ChannelPartnerLinkState = 3
	// Status when the reseller is suspended by Google or distributor.
	ChannelPartnerLinkState_SUSPENDED ChannelPartnerLinkState = 4
)

// Enum value maps for ChannelPartnerLinkState.
var (
	ChannelPartnerLinkState_name = map[int32]string{
		0: "CHANNEL_PARTNER_LINK_STATE_UNSPECIFIED",
		1: "INVITED",
		2: "ACTIVE",
		3: "REVOKED",
		4: "SUSPENDED",
	}
	ChannelPartnerLinkState_value = map[string]int32{
		"CHANNEL_PARTNER_LINK_STATE_UNSPECIFIED": 0,
		"INVITED":                                1,
		"ACTIVE":                                 2,
		"REVOKED":                                3,
		"SUSPENDED":                              4,
	}
)

func (x ChannelPartnerLinkState) Enum() *ChannelPartnerLinkState {
	p := new(ChannelPartnerLinkState)
	*p = x
	return p
}

func (x ChannelPartnerLinkState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChannelPartnerLinkState) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_channel_v1_channel_partner_links_proto_enumTypes[1].Descriptor()
}

func (ChannelPartnerLinkState) Type() protoreflect.EnumType {
	return &file_google_cloud_channel_v1_channel_partner_links_proto_enumTypes[1]
}

func (x ChannelPartnerLinkState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChannelPartnerLinkState.Descriptor instead.
func (ChannelPartnerLinkState) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_channel_v1_channel_partner_links_proto_rawDescGZIP(), []int{1}
}

// Entity representing a link between distributors and their indirect
// resellers in an n-tier resale channel.
type ChannelPartnerLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Resource name for the channel partner link, in the format
	// accounts/{account_id}/channelPartnerLinks/{id}.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. Cloud Identity ID of the linked reseller.
	ResellerCloudIdentityId string `protobuf:"bytes,2,opt,name=reseller_cloud_identity_id,json=resellerCloudIdentityId,proto3" json:"reseller_cloud_identity_id,omitempty"`
	// Required. State of the channel partner link.
	LinkState ChannelPartnerLinkState `protobuf:"varint,3,opt,name=link_state,json=linkState,proto3,enum=google.cloud.channel.v1.ChannelPartnerLinkState" json:"link_state,omitempty"`
	// Output only. URI of the web page where partner accepts the link invitation.
	InviteLinkUri string `protobuf:"bytes,4,opt,name=invite_link_uri,json=inviteLinkUri,proto3" json:"invite_link_uri,omitempty"`
	// Output only. Timestamp of when the channel partner link is created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. Timestamp of when the channel partner link is updated.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Output only. Public identifier that a customer must use to generate a
	// transfer token to move to this distributor-reseller combination.
	PublicId string `protobuf:"bytes,7,opt,name=public_id,json=publicId,proto3" json:"public_id,omitempty"`
	// Output only. Cloud Identity info of the channel partner (IR).
	ChannelPartnerCloudIdentityInfo *CloudIdentityInfo `protobuf:"bytes,8,opt,name=channel_partner_cloud_identity_info,json=channelPartnerCloudIdentityInfo,proto3" json:"channel_partner_cloud_identity_info,omitempty"`
}

func (x *ChannelPartnerLink) Reset() {
	*x = ChannelPartnerLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_channel_v1_channel_partner_links_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelPartnerLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelPartnerLink) ProtoMessage() {}

func (x *ChannelPartnerLink) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_channel_v1_channel_partner_links_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelPartnerLink.ProtoReflect.Descriptor instead.
func (*ChannelPartnerLink) Descriptor() ([]byte, []int) {
	return file_google_cloud_channel_v1_channel_partner_links_proto_rawDescGZIP(), []int{0}
}

func (x *ChannelPartnerLink) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChannelPartnerLink) GetResellerCloudIdentityId() string {
	if x != nil {
		return x.ResellerCloudIdentityId
	}
	return ""
}

func (x *ChannelPartnerLink) GetLinkState() ChannelPartnerLinkState {
	if x != nil {
		return x.LinkState
	}
	return ChannelPartnerLinkState_CHANNEL_PARTNER_LINK_STATE_UNSPECIFIED
}

func (x *ChannelPartnerLink) GetInviteLinkUri() string {
	if x != nil {
		return x.InviteLinkUri
	}
	return ""
}

func (x *ChannelPartnerLink) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *ChannelPartnerLink) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *ChannelPartnerLink) GetPublicId() string {
	if x != nil {
		return x.PublicId
	}
	return ""
}

func (x *ChannelPartnerLink) GetChannelPartnerCloudIdentityInfo() *CloudIdentityInfo {
	if x != nil {
		return x.ChannelPartnerCloudIdentityInfo
	}
	return nil
}

var File_google_cloud_channel_v1_channel_partner_links_proto protoreflect.FileDescriptor

var file_google_cloud_channel_v1_channel_partner_links_proto_rawDesc = []byte{
	0x0a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x8b, 0x05, 0x0a, 0x12, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x50, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x40, 0x0a, 0x1a, 0x72, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x17, 0x72, 0x65, 0x73, 0x65,
	0x6c, 0x6c, 0x65, 0x72, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x49, 0x64, 0x12, 0x54, 0x0a, 0x0a, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72,
	0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x09,
	0x6c, 0x69, 0x6e, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x0f, 0x69, 0x6e, 0x76,
	0x69, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0d, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x4c,
	0x69, 0x6e, 0x6b, 0x55, 0x72, 0x69, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x09, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x08, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x64, 0x12, 0x7d, 0x0a, 0x23,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x1f, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x3a, 0x72, 0xea, 0x41, 0x6f,
	0x0a, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b,
	0x12, 0x3d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x7d, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x50, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x2f, 0x7b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x7d, 0x2a,
	0x3e, 0x0a, 0x16, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x56, 0x69, 0x65, 0x77, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x42, 0x41,
	0x53, 0x49, 0x43, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x55, 0x4c, 0x4c, 0x10, 0x02, 0x2a,
	0x7a, 0x0a, 0x17, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2a, 0x0a, 0x26, 0x43, 0x48,
	0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x50, 0x41, 0x52, 0x54, 0x4e, 0x45, 0x52, 0x5f, 0x4c, 0x49,
	0x4e, 0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56, 0x49, 0x54, 0x45,
	0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x02, 0x12,
	0x0b, 0x0a, 0x07, 0x52, 0x45, 0x56, 0x4f, 0x4b, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09,
	0x53, 0x55, 0x53, 0x50, 0x45, 0x4e, 0x44, 0x45, 0x44, 0x10, 0x04, 0x42, 0x70, 0x0a, 0x1b, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x42, 0x18, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x70, 0x62, 0x3b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_channel_v1_channel_partner_links_proto_rawDescOnce sync.Once
	file_google_cloud_channel_v1_channel_partner_links_proto_rawDescData = file_google_cloud_channel_v1_channel_partner_links_proto_rawDesc
)

func file_google_cloud_channel_v1_channel_partner_links_proto_rawDescGZIP() []byte {
	file_google_cloud_channel_v1_channel_partner_links_proto_rawDescOnce.Do(func() {
		file_google_cloud_channel_v1_channel_partner_links_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_channel_v1_channel_partner_links_proto_rawDescData)
	})
	return file_google_cloud_channel_v1_channel_partner_links_proto_rawDescData
}

var file_google_cloud_channel_v1_channel_partner_links_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_cloud_channel_v1_channel_partner_links_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_channel_v1_channel_partner_links_proto_goTypes = []any{
	(ChannelPartnerLinkView)(0),   // 0: google.cloud.channel.v1.ChannelPartnerLinkView
	(ChannelPartnerLinkState)(0),  // 1: google.cloud.channel.v1.ChannelPartnerLinkState
	(*ChannelPartnerLink)(nil),    // 2: google.cloud.channel.v1.ChannelPartnerLink
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*CloudIdentityInfo)(nil),     // 4: google.cloud.channel.v1.CloudIdentityInfo
}
var file_google_cloud_channel_v1_channel_partner_links_proto_depIdxs = []int32{
	1, // 0: google.cloud.channel.v1.ChannelPartnerLink.link_state:type_name -> google.cloud.channel.v1.ChannelPartnerLinkState
	3, // 1: google.cloud.channel.v1.ChannelPartnerLink.create_time:type_name -> google.protobuf.Timestamp
	3, // 2: google.cloud.channel.v1.ChannelPartnerLink.update_time:type_name -> google.protobuf.Timestamp
	4, // 3: google.cloud.channel.v1.ChannelPartnerLink.channel_partner_cloud_identity_info:type_name -> google.cloud.channel.v1.CloudIdentityInfo
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_cloud_channel_v1_channel_partner_links_proto_init() }
func file_google_cloud_channel_v1_channel_partner_links_proto_init() {
	if File_google_cloud_channel_v1_channel_partner_links_proto != nil {
		return
	}
	file_google_cloud_channel_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_channel_v1_channel_partner_links_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ChannelPartnerLink); i {
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
			RawDescriptor: file_google_cloud_channel_v1_channel_partner_links_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_channel_v1_channel_partner_links_proto_goTypes,
		DependencyIndexes: file_google_cloud_channel_v1_channel_partner_links_proto_depIdxs,
		EnumInfos:         file_google_cloud_channel_v1_channel_partner_links_proto_enumTypes,
		MessageInfos:      file_google_cloud_channel_v1_channel_partner_links_proto_msgTypes,
	}.Build()
	File_google_cloud_channel_v1_channel_partner_links_proto = out.File
	file_google_cloud_channel_v1_channel_partner_links_proto_rawDesc = nil
	file_google_cloud_channel_v1_channel_partner_links_proto_goTypes = nil
	file_google_cloud_channel_v1_channel_partner_links_proto_depIdxs = nil
}
