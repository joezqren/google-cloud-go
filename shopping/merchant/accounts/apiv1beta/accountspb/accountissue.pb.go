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
// source: google/shopping/merchant/accounts/v1beta/accountissue.proto

package accountspb

import (
	context "context"
	reflect "reflect"
	sync "sync"

	typepb "cloud.google.com/go/shopping/type/typepb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// All possible issue severities.
type AccountIssue_Severity int32

const (
	// The severity is unknown.
	AccountIssue_SEVERITY_UNSPECIFIED AccountIssue_Severity = 0
	// The issue causes offers to not serve.
	AccountIssue_CRITICAL AccountIssue_Severity = 1
	// The issue might affect offers (in the future) or might be an
	// indicator of issues with offers.
	AccountIssue_ERROR AccountIssue_Severity = 2
	// The issue is a suggestion for improvement.
	AccountIssue_SUGGESTION AccountIssue_Severity = 3
)

// Enum value maps for AccountIssue_Severity.
var (
	AccountIssue_Severity_name = map[int32]string{
		0: "SEVERITY_UNSPECIFIED",
		1: "CRITICAL",
		2: "ERROR",
		3: "SUGGESTION",
	}
	AccountIssue_Severity_value = map[string]int32{
		"SEVERITY_UNSPECIFIED": 0,
		"CRITICAL":             1,
		"ERROR":                2,
		"SUGGESTION":           3,
	}
)

func (x AccountIssue_Severity) Enum() *AccountIssue_Severity {
	p := new(AccountIssue_Severity)
	*p = x
	return p
}

func (x AccountIssue_Severity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AccountIssue_Severity) Descriptor() protoreflect.EnumDescriptor {
	return file_google_shopping_merchant_accounts_v1beta_accountissue_proto_enumTypes[0].Descriptor()
}

func (AccountIssue_Severity) Type() protoreflect.EnumType {
	return &file_google_shopping_merchant_accounts_v1beta_accountissue_proto_enumTypes[0]
}

func (x AccountIssue_Severity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AccountIssue_Severity.Descriptor instead.
func (AccountIssue_Severity) EnumDescriptor() ([]byte, []int) {
	return file_google_shopping_merchant_accounts_v1beta_accountissue_proto_rawDescGZIP(), []int{0, 0}
}

// An
// [`AccountIssue`](https://support.google.com/merchants/answer/12153802?sjid=17798438912526418908-EU#account).
type AccountIssue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier. The resource name of the account issue.
	// Format: `accounts/{account}/issues/{id}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The localized title of the issue.
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// The overall severity of the issue.
	Severity AccountIssue_Severity `protobuf:"varint,3,opt,name=severity,proto3,enum=google.shopping.merchant.accounts.v1beta.AccountIssue_Severity" json:"severity,omitempty"`
	// The impact this issue has on various destinations.
	ImpactedDestinations []*AccountIssue_ImpactedDestination `protobuf:"bytes,4,rep,name=impacted_destinations,json=impactedDestinations,proto3" json:"impacted_destinations,omitempty"`
	// Further localized details about the issue.
	Detail string `protobuf:"bytes,5,opt,name=detail,proto3" json:"detail,omitempty"`
	// Link to Merchant Center Help Center providing further information about the
	// issue and how to fix it.
	DocumentationUri string `protobuf:"bytes,6,opt,name=documentation_uri,json=documentationUri,proto3" json:"documentation_uri,omitempty"`
}

func (x *AccountIssue) Reset() {
	*x = AccountIssue{}
	mi := &file_google_shopping_merchant_accounts_v1beta_accountissue_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AccountIssue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountIssue) ProtoMessage() {}

func (x *AccountIssue) ProtoReflect() protoreflect.Message {
	mi := &file_google_shopping_merchant_accounts_v1beta_accountissue_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountIssue.ProtoReflect.Descriptor instead.
func (*AccountIssue) Descriptor() ([]byte, []int) {
	return file_google_shopping_merchant_accounts_v1beta_accountissue_proto_rawDescGZIP(), []int{0}
}

func (x *AccountIssue) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AccountIssue) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AccountIssue) GetSeverity() AccountIssue_Severity {
	if x != nil {
		return x.Severity
	}
	return AccountIssue_SEVERITY_UNSPECIFIED
}

func (x *AccountIssue) GetImpactedDestinations() []*AccountIssue_ImpactedDestination {
	if x != nil {
		return x.ImpactedDestinations
	}
	return nil
}

func (x *AccountIssue) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

func (x *AccountIssue) GetDocumentationUri() string {
	if x != nil {
		return x.DocumentationUri
	}
	return ""
}

// Request message for the `ListAccountIssues` method.
type ListAccountIssuesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The parent, which owns this collection of issues.
	// Format: `accounts/{account}`
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Optional. The maximum number of issues to return. The service may return
	// fewer than this value. If unspecified, at most 50 users will be returned.
	// The maximum value is 100; values above 100 will be coerced to 100
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional. A page token, received from a previous `ListAccountIssues` call.
	// Provide this to retrieve the subsequent page.
	//
	// When paginating, all other parameters provided to `ListAccountIssues` must
	// match the call that provided the page token.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Optional. The issues in the response will have human-readable fields in the
	// given language. The format is [BCP-47](https://tools.ietf.org/html/bcp47),
	// such as `en-US` or `sr-Latn`. If not value is provided, `en-US` will be
	// used.
	LanguageCode string `protobuf:"bytes,4,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	// Optional. The [IANA](https://www.iana.org/time-zones) timezone used to
	// localize times in human-readable fields. For example 'America/Los_Angeles'.
	// If not set, 'America/Los_Angeles' will be used.
	TimeZone string `protobuf:"bytes,5,opt,name=time_zone,json=timeZone,proto3" json:"time_zone,omitempty"`
}

func (x *ListAccountIssuesRequest) Reset() {
	*x = ListAccountIssuesRequest{}
	mi := &file_google_shopping_merchant_accounts_v1beta_accountissue_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAccountIssuesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAccountIssuesRequest) ProtoMessage() {}

func (x *ListAccountIssuesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_shopping_merchant_accounts_v1beta_accountissue_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAccountIssuesRequest.ProtoReflect.Descriptor instead.
func (*ListAccountIssuesRequest) Descriptor() ([]byte, []int) {
	return file_google_shopping_merchant_accounts_v1beta_accountissue_proto_rawDescGZIP(), []int{1}
}

func (x *ListAccountIssuesRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListAccountIssuesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListAccountIssuesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListAccountIssuesRequest) GetLanguageCode() string {
	if x != nil {
		return x.LanguageCode
	}
	return ""
}

func (x *ListAccountIssuesRequest) GetTimeZone() string {
	if x != nil {
		return x.TimeZone
	}
	return ""
}

// Response message for the `ListAccountIssues` method.
type ListAccountIssuesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The issues from the specified account.
	AccountIssues []*AccountIssue `protobuf:"bytes,1,rep,name=account_issues,json=accountIssues,proto3" json:"account_issues,omitempty"`
	// A token, which can be sent as `page_token` to retrieve the next page.
	// If this field is omitted, there are no subsequent pages.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListAccountIssuesResponse) Reset() {
	*x = ListAccountIssuesResponse{}
	mi := &file_google_shopping_merchant_accounts_v1beta_accountissue_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAccountIssuesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAccountIssuesResponse) ProtoMessage() {}

func (x *ListAccountIssuesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_shopping_merchant_accounts_v1beta_accountissue_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAccountIssuesResponse.ProtoReflect.Descriptor instead.
func (*ListAccountIssuesResponse) Descriptor() ([]byte, []int) {
	return file_google_shopping_merchant_accounts_v1beta_accountissue_proto_rawDescGZIP(), []int{2}
}

func (x *ListAccountIssuesResponse) GetAccountIssues() []*AccountIssue {
	if x != nil {
		return x.AccountIssues
	}
	return nil
}

func (x *ListAccountIssuesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// The impact of the issue on a destination.
type AccountIssue_ImpactedDestination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The impacted reporting context.
	ReportingContext *typepb.ReportingContext_ReportingContextEnum `protobuf:"varint,1,opt,name=reporting_context,json=reportingContext,proto3,enum=google.shopping.type.ReportingContext_ReportingContextEnum,oneof" json:"reporting_context,omitempty"`
	// The (negative) impact for various regions on the given destination.
	Impacts []*AccountIssue_ImpactedDestination_Impact `protobuf:"bytes,2,rep,name=impacts,proto3" json:"impacts,omitempty"`
}

func (x *AccountIssue_ImpactedDestination) Reset() {
	*x = AccountIssue_ImpactedDestination{}
	mi := &file_google_shopping_merchant_accounts_v1beta_accountissue_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AccountIssue_ImpactedDestination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountIssue_ImpactedDestination) ProtoMessage() {}

func (x *AccountIssue_ImpactedDestination) ProtoReflect() protoreflect.Message {
	mi := &file_google_shopping_merchant_accounts_v1beta_accountissue_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountIssue_ImpactedDestination.ProtoReflect.Descriptor instead.
func (*AccountIssue_ImpactedDestination) Descriptor() ([]byte, []int) {
	return file_google_shopping_merchant_accounts_v1beta_accountissue_proto_rawDescGZIP(), []int{0, 0}
}

func (x *AccountIssue_ImpactedDestination) GetReportingContext() typepb.ReportingContext_ReportingContextEnum {
	if x != nil && x.ReportingContext != nil {
		return *x.ReportingContext
	}
	return typepb.ReportingContext_ReportingContextEnum(0)
}

func (x *AccountIssue_ImpactedDestination) GetImpacts() []*AccountIssue_ImpactedDestination_Impact {
	if x != nil {
		return x.Impacts
	}
	return nil
}

// The impact of the issue on a region.
type AccountIssue_ImpactedDestination_Impact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The [CLDR region code](https://cldr.unicode.org/) where this issue
	// applies.
	RegionCode string `protobuf:"bytes,1,opt,name=region_code,json=regionCode,proto3" json:"region_code,omitempty"`
	// The severity of the issue on the destination and region.
	Severity AccountIssue_Severity `protobuf:"varint,2,opt,name=severity,proto3,enum=google.shopping.merchant.accounts.v1beta.AccountIssue_Severity" json:"severity,omitempty"`
}

func (x *AccountIssue_ImpactedDestination_Impact) Reset() {
	*x = AccountIssue_ImpactedDestination_Impact{}
	mi := &file_google_shopping_merchant_accounts_v1beta_accountissue_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AccountIssue_ImpactedDestination_Impact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountIssue_ImpactedDestination_Impact) ProtoMessage() {}

func (x *AccountIssue_ImpactedDestination_Impact) ProtoReflect() protoreflect.Message {
	mi := &file_google_shopping_merchant_accounts_v1beta_accountissue_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountIssue_ImpactedDestination_Impact.ProtoReflect.Descriptor instead.
func (*AccountIssue_ImpactedDestination_Impact) Descriptor() ([]byte, []int) {
	return file_google_shopping_merchant_accounts_v1beta_accountissue_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *AccountIssue_ImpactedDestination_Impact) GetRegionCode() string {
	if x != nil {
		return x.RegionCode
	}
	return ""
}

func (x *AccountIssue_ImpactedDestination_Impact) GetSeverity() AccountIssue_Severity {
	if x != nil {
		return x.Severity
	}
	return AccountIssue_SEVERITY_UNSPECIFIED
}

var File_google_shopping_merchant_accounts_v1beta_accountissue_proto protoreflect.FileDescriptor

var file_google_shopping_merchant_accounts_v1beta_accountissue_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x2f, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x69, 0x73, 0x73, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x6d,
	0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x07, 0x0a,
	0x0c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x12, 0x17, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x08,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x5b, 0x0a, 0x08,
	0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3f,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x2e, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x52,
	0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x7f, 0x0a, 0x15, 0x69, 0x6d, 0x70,
	0x61, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68,
	0x61, 0x6e, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65,
	0x2e, 0x49, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x65, 0x64, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x14, 0x69, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x65, 0x64, 0x44, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x72, 0x69, 0x1a,
	0x90, 0x03, 0x0a, 0x13, 0x49, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x65, 0x64, 0x44, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x6d, 0x0a, 0x11, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x3b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x45, 0x6e, 0x75, 0x6d, 0x48,
	0x00, 0x52, 0x10, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x88, 0x01, 0x01, 0x12, 0x6b, 0x0a, 0x07, 0x69, 0x6d, 0x70, 0x61, 0x63, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x51, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61,
	0x6e, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x2e,
	0x49, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x65, 0x64, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x52, 0x07, 0x69, 0x6d, 0x70, 0x61,
	0x63, 0x74, 0x73, 0x1a, 0x86, 0x01, 0x0a, 0x06, 0x49, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x5b, 0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x3f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x2e, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69,
	0x74, 0x79, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x42, 0x14, 0x0a, 0x12,
	0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x22, 0x4d, 0x0a, 0x08, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x18,
	0x0a, 0x14, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52, 0x49, 0x54,
	0x49, 0x43, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10,
	0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x55, 0x47, 0x47, 0x45, 0x53, 0x54, 0x49, 0x4f, 0x4e, 0x10,
	0x03, 0x3a, 0x6c, 0xea, 0x41, 0x69, 0x0a, 0x27, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74,
	0x61, 0x70, 0x69, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x12,
	0x21, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x7d, 0x2f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x73, 0x73, 0x75,
	0x65, 0x7d, 0x2a, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65,
	0x73, 0x32, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x22,
	0xf0, 0x01, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x73, 0x73, 0x75, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x06,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2a, 0xe0, 0x41,
	0x02, 0xfa, 0x41, 0x24, 0x0a, 0x22, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x61, 0x70,
	0x69, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x12, 0x20, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x22, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x28, 0x0a, 0x0d, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x01, 0x52, 0x0c, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x20, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x5a, 0x6f,
	0x6e, 0x65, 0x22, 0xa2, 0x01, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x5d, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x73, 0x73, 0x75,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68,
	0x61, 0x6e, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65,
	0x52, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x73, 0x12,
	0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61,
	0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xbb, 0x02, 0x0a, 0x13, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0xda, 0x01, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x73, 0x73, 0x75, 0x65, 0x73, 0x12, 0x42, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73,
	0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74,
	0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x73, 0x73, 0x75,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x43, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x73, 0x73, 0x75, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3c,
	0xda, 0x41, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2d, 0x12,
	0x2b, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x73, 0x1a, 0x47, 0xca, 0x41,
	0x1a, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x27, 0x68, 0x74,
	0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x93, 0x01, 0x0a, 0x2c, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65,
	0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x42, 0x11, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x73, 0x73, 0x75, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x2f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61,
	0x6e, 0x74, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x70, 0x62,
	0x3b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_google_shopping_merchant_accounts_v1beta_accountissue_proto_rawDescOnce sync.Once
	file_google_shopping_merchant_accounts_v1beta_accountissue_proto_rawDescData = file_google_shopping_merchant_accounts_v1beta_accountissue_proto_rawDesc
)

func file_google_shopping_merchant_accounts_v1beta_accountissue_proto_rawDescGZIP() []byte {
	file_google_shopping_merchant_accounts_v1beta_accountissue_proto_rawDescOnce.Do(func() {
		file_google_shopping_merchant_accounts_v1beta_accountissue_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_shopping_merchant_accounts_v1beta_accountissue_proto_rawDescData)
	})
	return file_google_shopping_merchant_accounts_v1beta_accountissue_proto_rawDescData
}

var file_google_shopping_merchant_accounts_v1beta_accountissue_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_shopping_merchant_accounts_v1beta_accountissue_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_shopping_merchant_accounts_v1beta_accountissue_proto_goTypes = []any{
	(AccountIssue_Severity)(0),                        // 0: google.shopping.merchant.accounts.v1beta.AccountIssue.Severity
	(*AccountIssue)(nil),                              // 1: google.shopping.merchant.accounts.v1beta.AccountIssue
	(*ListAccountIssuesRequest)(nil),                  // 2: google.shopping.merchant.accounts.v1beta.ListAccountIssuesRequest
	(*ListAccountIssuesResponse)(nil),                 // 3: google.shopping.merchant.accounts.v1beta.ListAccountIssuesResponse
	(*AccountIssue_ImpactedDestination)(nil),          // 4: google.shopping.merchant.accounts.v1beta.AccountIssue.ImpactedDestination
	(*AccountIssue_ImpactedDestination_Impact)(nil),   // 5: google.shopping.merchant.accounts.v1beta.AccountIssue.ImpactedDestination.Impact
	(typepb.ReportingContext_ReportingContextEnum)(0), // 6: google.shopping.type.ReportingContext.ReportingContextEnum
}
var file_google_shopping_merchant_accounts_v1beta_accountissue_proto_depIdxs = []int32{
	0, // 0: google.shopping.merchant.accounts.v1beta.AccountIssue.severity:type_name -> google.shopping.merchant.accounts.v1beta.AccountIssue.Severity
	4, // 1: google.shopping.merchant.accounts.v1beta.AccountIssue.impacted_destinations:type_name -> google.shopping.merchant.accounts.v1beta.AccountIssue.ImpactedDestination
	1, // 2: google.shopping.merchant.accounts.v1beta.ListAccountIssuesResponse.account_issues:type_name -> google.shopping.merchant.accounts.v1beta.AccountIssue
	6, // 3: google.shopping.merchant.accounts.v1beta.AccountIssue.ImpactedDestination.reporting_context:type_name -> google.shopping.type.ReportingContext.ReportingContextEnum
	5, // 4: google.shopping.merchant.accounts.v1beta.AccountIssue.ImpactedDestination.impacts:type_name -> google.shopping.merchant.accounts.v1beta.AccountIssue.ImpactedDestination.Impact
	0, // 5: google.shopping.merchant.accounts.v1beta.AccountIssue.ImpactedDestination.Impact.severity:type_name -> google.shopping.merchant.accounts.v1beta.AccountIssue.Severity
	2, // 6: google.shopping.merchant.accounts.v1beta.AccountIssueService.ListAccountIssues:input_type -> google.shopping.merchant.accounts.v1beta.ListAccountIssuesRequest
	3, // 7: google.shopping.merchant.accounts.v1beta.AccountIssueService.ListAccountIssues:output_type -> google.shopping.merchant.accounts.v1beta.ListAccountIssuesResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_shopping_merchant_accounts_v1beta_accountissue_proto_init() }
func file_google_shopping_merchant_accounts_v1beta_accountissue_proto_init() {
	if File_google_shopping_merchant_accounts_v1beta_accountissue_proto != nil {
		return
	}
	file_google_shopping_merchant_accounts_v1beta_accountissue_proto_msgTypes[3].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_shopping_merchant_accounts_v1beta_accountissue_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_shopping_merchant_accounts_v1beta_accountissue_proto_goTypes,
		DependencyIndexes: file_google_shopping_merchant_accounts_v1beta_accountissue_proto_depIdxs,
		EnumInfos:         file_google_shopping_merchant_accounts_v1beta_accountissue_proto_enumTypes,
		MessageInfos:      file_google_shopping_merchant_accounts_v1beta_accountissue_proto_msgTypes,
	}.Build()
	File_google_shopping_merchant_accounts_v1beta_accountissue_proto = out.File
	file_google_shopping_merchant_accounts_v1beta_accountissue_proto_rawDesc = nil
	file_google_shopping_merchant_accounts_v1beta_accountissue_proto_goTypes = nil
	file_google_shopping_merchant_accounts_v1beta_accountissue_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AccountIssueServiceClient is the client API for AccountIssueService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountIssueServiceClient interface {
	// Lists all account issues of a Merchant Center account.
	ListAccountIssues(ctx context.Context, in *ListAccountIssuesRequest, opts ...grpc.CallOption) (*ListAccountIssuesResponse, error)
}

type accountIssueServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountIssueServiceClient(cc grpc.ClientConnInterface) AccountIssueServiceClient {
	return &accountIssueServiceClient{cc}
}

func (c *accountIssueServiceClient) ListAccountIssues(ctx context.Context, in *ListAccountIssuesRequest, opts ...grpc.CallOption) (*ListAccountIssuesResponse, error) {
	out := new(ListAccountIssuesResponse)
	err := c.cc.Invoke(ctx, "/google.shopping.merchant.accounts.v1beta.AccountIssueService/ListAccountIssues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountIssueServiceServer is the server API for AccountIssueService service.
type AccountIssueServiceServer interface {
	// Lists all account issues of a Merchant Center account.
	ListAccountIssues(context.Context, *ListAccountIssuesRequest) (*ListAccountIssuesResponse, error)
}

// UnimplementedAccountIssueServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAccountIssueServiceServer struct {
}

func (*UnimplementedAccountIssueServiceServer) ListAccountIssues(context.Context, *ListAccountIssuesRequest) (*ListAccountIssuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccountIssues not implemented")
}

func RegisterAccountIssueServiceServer(s *grpc.Server, srv AccountIssueServiceServer) {
	s.RegisterService(&_AccountIssueService_serviceDesc, srv)
}

func _AccountIssueService_ListAccountIssues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAccountIssuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountIssueServiceServer).ListAccountIssues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.shopping.merchant.accounts.v1beta.AccountIssueService/ListAccountIssues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountIssueServiceServer).ListAccountIssues(ctx, req.(*ListAccountIssuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountIssueService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.shopping.merchant.accounts.v1beta.AccountIssueService",
	HandlerType: (*AccountIssueServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAccountIssues",
			Handler:    _AccountIssueService_ListAccountIssues_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/shopping/merchant/accounts/v1beta/accountissue.proto",
}
