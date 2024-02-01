// Copyright 2023 Google LLC
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
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: google/cloud/baremetalsolution/v2/volume_snapshot.proto

package baremetalsolutionpb

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

// Represents the type of a snapshot.
type VolumeSnapshot_SnapshotType int32

const (
	// Type is not specified.
	VolumeSnapshot_SNAPSHOT_TYPE_UNSPECIFIED VolumeSnapshot_SnapshotType = 0
	// Snapshot was taken manually by user.
	VolumeSnapshot_AD_HOC VolumeSnapshot_SnapshotType = 1
	// Snapshot was taken automatically as a part of a snapshot schedule.
	VolumeSnapshot_SCHEDULED VolumeSnapshot_SnapshotType = 2
)

// Enum value maps for VolumeSnapshot_SnapshotType.
var (
	VolumeSnapshot_SnapshotType_name = map[int32]string{
		0: "SNAPSHOT_TYPE_UNSPECIFIED",
		1: "AD_HOC",
		2: "SCHEDULED",
	}
	VolumeSnapshot_SnapshotType_value = map[string]int32{
		"SNAPSHOT_TYPE_UNSPECIFIED": 0,
		"AD_HOC":                    1,
		"SCHEDULED":                 2,
	}
)

func (x VolumeSnapshot_SnapshotType) Enum() *VolumeSnapshot_SnapshotType {
	p := new(VolumeSnapshot_SnapshotType)
	*p = x
	return p
}

func (x VolumeSnapshot_SnapshotType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VolumeSnapshot_SnapshotType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_enumTypes[0].Descriptor()
}

func (VolumeSnapshot_SnapshotType) Type() protoreflect.EnumType {
	return &file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_enumTypes[0]
}

func (x VolumeSnapshot_SnapshotType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VolumeSnapshot_SnapshotType.Descriptor instead.
func (VolumeSnapshot_SnapshotType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_rawDescGZIP(), []int{0, 0}
}

// A snapshot of a volume. Only boot volumes can have snapshots.
type VolumeSnapshot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the snapshot.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. An identifier for the snapshot, generated by the backend.
	Id string `protobuf:"bytes,6,opt,name=id,proto3" json:"id,omitempty"`
	// The description of the snapshot.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Output only. The creation time of the snapshot.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The name of the volume which this snapshot belongs to.
	StorageVolume string `protobuf:"bytes,5,opt,name=storage_volume,json=storageVolume,proto3" json:"storage_volume,omitempty"`
	// Output only. The type of the snapshot which indicates whether it was
	// scheduled or manual/ad-hoc.
	Type VolumeSnapshot_SnapshotType `protobuf:"varint,7,opt,name=type,proto3,enum=google.cloud.baremetalsolution.v2.VolumeSnapshot_SnapshotType" json:"type,omitempty"`
}

func (x *VolumeSnapshot) Reset() {
	*x = VolumeSnapshot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VolumeSnapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VolumeSnapshot) ProtoMessage() {}

func (x *VolumeSnapshot) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VolumeSnapshot.ProtoReflect.Descriptor instead.
func (*VolumeSnapshot) Descriptor() ([]byte, []int) {
	return file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_rawDescGZIP(), []int{0}
}

func (x *VolumeSnapshot) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *VolumeSnapshot) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *VolumeSnapshot) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *VolumeSnapshot) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *VolumeSnapshot) GetStorageVolume() string {
	if x != nil {
		return x.StorageVolume
	}
	return ""
}

func (x *VolumeSnapshot) GetType() VolumeSnapshot_SnapshotType {
	if x != nil {
		return x.Type
	}
	return VolumeSnapshot_SNAPSHOT_TYPE_UNSPECIFIED
}

// Message for requesting volume snapshot information.
type GetVolumeSnapshotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the snapshot.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetVolumeSnapshotRequest) Reset() {
	*x = GetVolumeSnapshotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVolumeSnapshotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVolumeSnapshotRequest) ProtoMessage() {}

func (x *GetVolumeSnapshotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVolumeSnapshotRequest.ProtoReflect.Descriptor instead.
func (*GetVolumeSnapshotRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_rawDescGZIP(), []int{1}
}

func (x *GetVolumeSnapshotRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Message for requesting a list of volume snapshots.
type ListVolumeSnapshotsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Parent value for ListVolumesRequest.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Requested page size. The server might return fewer items than requested.
	// If unspecified, server will pick an appropriate default.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// A token identifying a page of results from the server.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListVolumeSnapshotsRequest) Reset() {
	*x = ListVolumeSnapshotsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListVolumeSnapshotsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVolumeSnapshotsRequest) ProtoMessage() {}

func (x *ListVolumeSnapshotsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVolumeSnapshotsRequest.ProtoReflect.Descriptor instead.
func (*ListVolumeSnapshotsRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_rawDescGZIP(), []int{2}
}

func (x *ListVolumeSnapshotsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListVolumeSnapshotsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListVolumeSnapshotsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// Response message containing the list of volume snapshots.
type ListVolumeSnapshotsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of snapshots.
	VolumeSnapshots []*VolumeSnapshot `protobuf:"bytes,1,rep,name=volume_snapshots,json=volumeSnapshots,proto3" json:"volume_snapshots,omitempty"`
	// A token identifying a page of results from the server.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// Locations that could not be reached.
	Unreachable []string `protobuf:"bytes,3,rep,name=unreachable,proto3" json:"unreachable,omitempty"`
}

func (x *ListVolumeSnapshotsResponse) Reset() {
	*x = ListVolumeSnapshotsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListVolumeSnapshotsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVolumeSnapshotsResponse) ProtoMessage() {}

func (x *ListVolumeSnapshotsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVolumeSnapshotsResponse.ProtoReflect.Descriptor instead.
func (*ListVolumeSnapshotsResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_rawDescGZIP(), []int{3}
}

func (x *ListVolumeSnapshotsResponse) GetVolumeSnapshots() []*VolumeSnapshot {
	if x != nil {
		return x.VolumeSnapshots
	}
	return nil
}

func (x *ListVolumeSnapshotsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *ListVolumeSnapshotsResponse) GetUnreachable() []string {
	if x != nil {
		return x.Unreachable
	}
	return nil
}

// Message for deleting named Volume snapshot.
type DeleteVolumeSnapshotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the snapshot to delete.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteVolumeSnapshotRequest) Reset() {
	*x = DeleteVolumeSnapshotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteVolumeSnapshotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteVolumeSnapshotRequest) ProtoMessage() {}

func (x *DeleteVolumeSnapshotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteVolumeSnapshotRequest.ProtoReflect.Descriptor instead.
func (*DeleteVolumeSnapshotRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteVolumeSnapshotRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Message for creating a volume snapshot.
type CreateVolumeSnapshotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The volume to snapshot.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. The snapshot to create.
	VolumeSnapshot *VolumeSnapshot `protobuf:"bytes,2,opt,name=volume_snapshot,json=volumeSnapshot,proto3" json:"volume_snapshot,omitempty"`
}

func (x *CreateVolumeSnapshotRequest) Reset() {
	*x = CreateVolumeSnapshotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateVolumeSnapshotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVolumeSnapshotRequest) ProtoMessage() {}

func (x *CreateVolumeSnapshotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVolumeSnapshotRequest.ProtoReflect.Descriptor instead.
func (*CreateVolumeSnapshotRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_rawDescGZIP(), []int{5}
}

func (x *CreateVolumeSnapshotRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreateVolumeSnapshotRequest) GetVolumeSnapshot() *VolumeSnapshot {
	if x != nil {
		return x.VolumeSnapshot
	}
	return nil
}

// Message for restoring a volume snapshot.
type RestoreVolumeSnapshotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Name of the snapshot which will be used to restore its parent
	// volume.
	VolumeSnapshot string `protobuf:"bytes,1,opt,name=volume_snapshot,json=volumeSnapshot,proto3" json:"volume_snapshot,omitempty"`
}

func (x *RestoreVolumeSnapshotRequest) Reset() {
	*x = RestoreVolumeSnapshotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestoreVolumeSnapshotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestoreVolumeSnapshotRequest) ProtoMessage() {}

func (x *RestoreVolumeSnapshotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestoreVolumeSnapshotRequest.ProtoReflect.Descriptor instead.
func (*RestoreVolumeSnapshotRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_rawDescGZIP(), []int{6}
}

func (x *RestoreVolumeSnapshotRequest) GetVolumeSnapshot() string {
	if x != nil {
		return x.VolumeSnapshot
	}
	return ""
}

var File_google_cloud_baremetalsolution_v2_volume_snapshot_proto protoreflect.FileDescriptor

var file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_rawDesc = []byte{
	0x0a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62,
	0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x76, 0x32, 0x2f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x73, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61,
	0x6c, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x04, 0x0a, 0x0e, 0x56, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x56, 0x0a, 0x0e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x2f, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x29, 0x0a, 0x27, 0x62, 0x61, 0x72, 0x65, 0x6d,
	0x65, 0x74, 0x61, 0x6c, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x56, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x52, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x12, 0x57, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x3e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62,
	0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x32, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68,
	0x6f, 0x74, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x48, 0x0a, 0x0c, 0x53, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x53, 0x4e,
	0x41, 0x50, 0x53, 0x48, 0x4f, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x44, 0x5f,
	0x48, 0x4f, 0x43, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c,
	0x45, 0x44, 0x10, 0x02, 0x3a, 0x84, 0x01, 0xea, 0x41, 0x80, 0x01, 0x0a, 0x2f, 0x62, 0x61, 0x72,
	0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x56, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x4d, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d,
	0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x2f, 0x7b, 0x76,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x7d, 0x2f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73,
	0x2f, 0x7b, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x7d, 0x22, 0x67, 0x0a, 0x18, 0x47,
	0x65, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x37, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x31, 0x0a, 0x2f, 0x62,
	0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0xa1, 0x01, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x47, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x2f, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x29, 0x0a, 0x27, 0x62, 0x61, 0x72,
	0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x56, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xc5, 0x01, 0x0a, 0x1b, 0x4c, 0x69, 0x73,
	0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x10, 0x76, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x5f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x6f, 0x6c, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x53, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x0f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x53, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20,
	0x0a, 0x0b, 0x75, 0x6e, 0x72, 0x65, 0x61, 0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x6e, 0x72, 0x65, 0x61, 0x63, 0x68, 0x61, 0x62, 0x6c, 0x65,
	0x22, 0x6a, 0x0a, 0x1b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x4b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x37, 0xe0,
	0x41, 0x02, 0xfa, 0x41, 0x31, 0x0a, 0x2f, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c,
	0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x53, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xc7, 0x01, 0x0a,
	0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x53, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x47, 0x0a, 0x06,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2f, 0xe0, 0x41,
	0x02, 0xfa, 0x41, 0x29, 0x0a, 0x27, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73,
	0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x06, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x5f, 0x0a, 0x0f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f,
	0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x61,
	0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x32, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f,
	0x74, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0e, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x53, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x22, 0x80, 0x01, 0x0a, 0x1c, 0x52, 0x65, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x60, 0x0a, 0x0f, 0x76, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x5f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x37, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x31, 0x0a, 0x2f, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65,
	0x74, 0x61, 0x6c, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x56, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x0e, 0x76, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x42, 0x82, 0x02, 0x0a, 0x25, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62,
	0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x32, 0x42, 0x13, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x53, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x53, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f,
	0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61,
	0x6c, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0x3b, 0x62, 0x61, 0x72, 0x65,
	0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0xaa,
	0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x42,
	0x61, 0x72, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x6c, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x56, 0x32, 0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x5c, 0x42, 0x61, 0x72, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x6c, 0x53, 0x6f, 0x6c, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x32, 0xea, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x42, 0x61, 0x72, 0x65, 0x4d, 0x65, 0x74,
	0x61, 0x6c, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_rawDescOnce sync.Once
	file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_rawDescData = file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_rawDesc
)

func file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_rawDescGZIP() []byte {
	file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_rawDescOnce.Do(func() {
		file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_rawDescData)
	})
	return file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_rawDescData
}

var file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_goTypes = []interface{}{
	(VolumeSnapshot_SnapshotType)(0),     // 0: google.cloud.baremetalsolution.v2.VolumeSnapshot.SnapshotType
	(*VolumeSnapshot)(nil),               // 1: google.cloud.baremetalsolution.v2.VolumeSnapshot
	(*GetVolumeSnapshotRequest)(nil),     // 2: google.cloud.baremetalsolution.v2.GetVolumeSnapshotRequest
	(*ListVolumeSnapshotsRequest)(nil),   // 3: google.cloud.baremetalsolution.v2.ListVolumeSnapshotsRequest
	(*ListVolumeSnapshotsResponse)(nil),  // 4: google.cloud.baremetalsolution.v2.ListVolumeSnapshotsResponse
	(*DeleteVolumeSnapshotRequest)(nil),  // 5: google.cloud.baremetalsolution.v2.DeleteVolumeSnapshotRequest
	(*CreateVolumeSnapshotRequest)(nil),  // 6: google.cloud.baremetalsolution.v2.CreateVolumeSnapshotRequest
	(*RestoreVolumeSnapshotRequest)(nil), // 7: google.cloud.baremetalsolution.v2.RestoreVolumeSnapshotRequest
	(*timestamppb.Timestamp)(nil),        // 8: google.protobuf.Timestamp
}
var file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_depIdxs = []int32{
	8, // 0: google.cloud.baremetalsolution.v2.VolumeSnapshot.create_time:type_name -> google.protobuf.Timestamp
	0, // 1: google.cloud.baremetalsolution.v2.VolumeSnapshot.type:type_name -> google.cloud.baremetalsolution.v2.VolumeSnapshot.SnapshotType
	1, // 2: google.cloud.baremetalsolution.v2.ListVolumeSnapshotsResponse.volume_snapshots:type_name -> google.cloud.baremetalsolution.v2.VolumeSnapshot
	1, // 3: google.cloud.baremetalsolution.v2.CreateVolumeSnapshotRequest.volume_snapshot:type_name -> google.cloud.baremetalsolution.v2.VolumeSnapshot
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_init() }
func file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_init() {
	if File_google_cloud_baremetalsolution_v2_volume_snapshot_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VolumeSnapshot); i {
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
		file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVolumeSnapshotRequest); i {
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
		file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListVolumeSnapshotsRequest); i {
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
		file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListVolumeSnapshotsResponse); i {
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
		file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteVolumeSnapshotRequest); i {
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
		file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateVolumeSnapshotRequest); i {
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
		file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestoreVolumeSnapshotRequest); i {
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
			RawDescriptor: file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_goTypes,
		DependencyIndexes: file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_depIdxs,
		EnumInfos:         file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_enumTypes,
		MessageInfos:      file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_msgTypes,
	}.Build()
	File_google_cloud_baremetalsolution_v2_volume_snapshot_proto = out.File
	file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_rawDesc = nil
	file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_goTypes = nil
	file_google_cloud_baremetalsolution_v2_volume_snapshot_proto_depIdxs = nil
}
