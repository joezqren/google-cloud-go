// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package securitycenter

import (
	"context"
	"time"

	"cloud.google.com/go/longrunning"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	securitycenterpb "cloud.google.com/go/securitycenter/apiv1/securitycenterpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// BulkMuteFindingsOperation manages a long-running operation from BulkMuteFindings.
type BulkMuteFindingsOperation struct {
	lro      *longrunning.Operation
	pollPath string
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *BulkMuteFindingsOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*securitycenterpb.BulkMuteFindingsResponse, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp securitycenterpb.BulkMuteFindingsResponse
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *BulkMuteFindingsOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*securitycenterpb.BulkMuteFindingsResponse, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp securitycenterpb.BulkMuteFindingsResponse
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *BulkMuteFindingsOperation) Metadata() (*emptypb.Empty, error) {
	var meta emptypb.Empty
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *BulkMuteFindingsOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *BulkMuteFindingsOperation) Name() string {
	return op.lro.Name()
}

// RunAssetDiscoveryOperation manages a long-running operation from RunAssetDiscovery.
type RunAssetDiscoveryOperation struct {
	lro      *longrunning.Operation
	pollPath string
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *RunAssetDiscoveryOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*securitycenterpb.RunAssetDiscoveryResponse, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp securitycenterpb.RunAssetDiscoveryResponse
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *RunAssetDiscoveryOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*securitycenterpb.RunAssetDiscoveryResponse, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp securitycenterpb.RunAssetDiscoveryResponse
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *RunAssetDiscoveryOperation) Metadata() (*emptypb.Empty, error) {
	var meta emptypb.Empty
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *RunAssetDiscoveryOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *RunAssetDiscoveryOperation) Name() string {
	return op.lro.Name()
}

// AttackPathIterator manages a stream of *securitycenterpb.AttackPath.
type AttackPathIterator struct {
	items    []*securitycenterpb.AttackPath
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*securitycenterpb.AttackPath, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *AttackPathIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *AttackPathIterator) Next() (*securitycenterpb.AttackPath, error) {
	var item *securitycenterpb.AttackPath
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *AttackPathIterator) bufLen() int {
	return len(it.items)
}

func (it *AttackPathIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// BigQueryExportIterator manages a stream of *securitycenterpb.BigQueryExport.
type BigQueryExportIterator struct {
	items    []*securitycenterpb.BigQueryExport
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*securitycenterpb.BigQueryExport, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *BigQueryExportIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *BigQueryExportIterator) Next() (*securitycenterpb.BigQueryExport, error) {
	var item *securitycenterpb.BigQueryExport
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *BigQueryExportIterator) bufLen() int {
	return len(it.items)
}

func (it *BigQueryExportIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// EffectiveEventThreatDetectionCustomModuleIterator manages a stream of *securitycenterpb.EffectiveEventThreatDetectionCustomModule.
type EffectiveEventThreatDetectionCustomModuleIterator struct {
	items    []*securitycenterpb.EffectiveEventThreatDetectionCustomModule
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*securitycenterpb.EffectiveEventThreatDetectionCustomModule, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *EffectiveEventThreatDetectionCustomModuleIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *EffectiveEventThreatDetectionCustomModuleIterator) Next() (*securitycenterpb.EffectiveEventThreatDetectionCustomModule, error) {
	var item *securitycenterpb.EffectiveEventThreatDetectionCustomModule
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *EffectiveEventThreatDetectionCustomModuleIterator) bufLen() int {
	return len(it.items)
}

func (it *EffectiveEventThreatDetectionCustomModuleIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// EffectiveSecurityHealthAnalyticsCustomModuleIterator manages a stream of *securitycenterpb.EffectiveSecurityHealthAnalyticsCustomModule.
type EffectiveSecurityHealthAnalyticsCustomModuleIterator struct {
	items    []*securitycenterpb.EffectiveSecurityHealthAnalyticsCustomModule
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*securitycenterpb.EffectiveSecurityHealthAnalyticsCustomModule, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *EffectiveSecurityHealthAnalyticsCustomModuleIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *EffectiveSecurityHealthAnalyticsCustomModuleIterator) Next() (*securitycenterpb.EffectiveSecurityHealthAnalyticsCustomModule, error) {
	var item *securitycenterpb.EffectiveSecurityHealthAnalyticsCustomModule
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *EffectiveSecurityHealthAnalyticsCustomModuleIterator) bufLen() int {
	return len(it.items)
}

func (it *EffectiveSecurityHealthAnalyticsCustomModuleIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// EventThreatDetectionCustomModuleIterator manages a stream of *securitycenterpb.EventThreatDetectionCustomModule.
type EventThreatDetectionCustomModuleIterator struct {
	items    []*securitycenterpb.EventThreatDetectionCustomModule
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*securitycenterpb.EventThreatDetectionCustomModule, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *EventThreatDetectionCustomModuleIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *EventThreatDetectionCustomModuleIterator) Next() (*securitycenterpb.EventThreatDetectionCustomModule, error) {
	var item *securitycenterpb.EventThreatDetectionCustomModule
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *EventThreatDetectionCustomModuleIterator) bufLen() int {
	return len(it.items)
}

func (it *EventThreatDetectionCustomModuleIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// GroupResultIterator manages a stream of *securitycenterpb.GroupResult.
type GroupResultIterator struct {
	items    []*securitycenterpb.GroupResult
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*securitycenterpb.GroupResult, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *GroupResultIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *GroupResultIterator) Next() (*securitycenterpb.GroupResult, error) {
	var item *securitycenterpb.GroupResult
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *GroupResultIterator) bufLen() int {
	return len(it.items)
}

func (it *GroupResultIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// ListAssetsResponse_ListAssetsResultIterator manages a stream of *securitycenterpb.ListAssetsResponse_ListAssetsResult.
type ListAssetsResponse_ListAssetsResultIterator struct {
	items    []*securitycenterpb.ListAssetsResponse_ListAssetsResult
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*securitycenterpb.ListAssetsResponse_ListAssetsResult, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *ListAssetsResponse_ListAssetsResultIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ListAssetsResponse_ListAssetsResultIterator) Next() (*securitycenterpb.ListAssetsResponse_ListAssetsResult, error) {
	var item *securitycenterpb.ListAssetsResponse_ListAssetsResult
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ListAssetsResponse_ListAssetsResultIterator) bufLen() int {
	return len(it.items)
}

func (it *ListAssetsResponse_ListAssetsResultIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// ListFindingsResponse_ListFindingsResultIterator manages a stream of *securitycenterpb.ListFindingsResponse_ListFindingsResult.
type ListFindingsResponse_ListFindingsResultIterator struct {
	items    []*securitycenterpb.ListFindingsResponse_ListFindingsResult
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*securitycenterpb.ListFindingsResponse_ListFindingsResult, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *ListFindingsResponse_ListFindingsResultIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ListFindingsResponse_ListFindingsResultIterator) Next() (*securitycenterpb.ListFindingsResponse_ListFindingsResult, error) {
	var item *securitycenterpb.ListFindingsResponse_ListFindingsResult
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ListFindingsResponse_ListFindingsResultIterator) bufLen() int {
	return len(it.items)
}

func (it *ListFindingsResponse_ListFindingsResultIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// MuteConfigIterator manages a stream of *securitycenterpb.MuteConfig.
type MuteConfigIterator struct {
	items    []*securitycenterpb.MuteConfig
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*securitycenterpb.MuteConfig, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *MuteConfigIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *MuteConfigIterator) Next() (*securitycenterpb.MuteConfig, error) {
	var item *securitycenterpb.MuteConfig
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *MuteConfigIterator) bufLen() int {
	return len(it.items)
}

func (it *MuteConfigIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// NotificationConfigIterator manages a stream of *securitycenterpb.NotificationConfig.
type NotificationConfigIterator struct {
	items    []*securitycenterpb.NotificationConfig
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*securitycenterpb.NotificationConfig, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *NotificationConfigIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *NotificationConfigIterator) Next() (*securitycenterpb.NotificationConfig, error) {
	var item *securitycenterpb.NotificationConfig
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *NotificationConfigIterator) bufLen() int {
	return len(it.items)
}

func (it *NotificationConfigIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// OperationIterator manages a stream of *longrunningpb.Operation.
type OperationIterator struct {
	items    []*longrunningpb.Operation
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*longrunningpb.Operation, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *OperationIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *OperationIterator) Next() (*longrunningpb.Operation, error) {
	var item *longrunningpb.Operation
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *OperationIterator) bufLen() int {
	return len(it.items)
}

func (it *OperationIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// ResourceValueConfigIterator manages a stream of *securitycenterpb.ResourceValueConfig.
type ResourceValueConfigIterator struct {
	items    []*securitycenterpb.ResourceValueConfig
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*securitycenterpb.ResourceValueConfig, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *ResourceValueConfigIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ResourceValueConfigIterator) Next() (*securitycenterpb.ResourceValueConfig, error) {
	var item *securitycenterpb.ResourceValueConfig
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ResourceValueConfigIterator) bufLen() int {
	return len(it.items)
}

func (it *ResourceValueConfigIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// SecurityHealthAnalyticsCustomModuleIterator manages a stream of *securitycenterpb.SecurityHealthAnalyticsCustomModule.
type SecurityHealthAnalyticsCustomModuleIterator struct {
	items    []*securitycenterpb.SecurityHealthAnalyticsCustomModule
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*securitycenterpb.SecurityHealthAnalyticsCustomModule, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *SecurityHealthAnalyticsCustomModuleIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *SecurityHealthAnalyticsCustomModuleIterator) Next() (*securitycenterpb.SecurityHealthAnalyticsCustomModule, error) {
	var item *securitycenterpb.SecurityHealthAnalyticsCustomModule
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *SecurityHealthAnalyticsCustomModuleIterator) bufLen() int {
	return len(it.items)
}

func (it *SecurityHealthAnalyticsCustomModuleIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// SourceIterator manages a stream of *securitycenterpb.Source.
type SourceIterator struct {
	items    []*securitycenterpb.Source
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*securitycenterpb.Source, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *SourceIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *SourceIterator) Next() (*securitycenterpb.Source, error) {
	var item *securitycenterpb.Source
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *SourceIterator) bufLen() int {
	return len(it.items)
}

func (it *SourceIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// ValuedResourceIterator manages a stream of *securitycenterpb.ValuedResource.
type ValuedResourceIterator struct {
	items    []*securitycenterpb.ValuedResource
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*securitycenterpb.ValuedResource, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *ValuedResourceIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ValuedResourceIterator) Next() (*securitycenterpb.ValuedResource, error) {
	var item *securitycenterpb.ValuedResource
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ValuedResourceIterator) bufLen() int {
	return len(it.items)
}

func (it *ValuedResourceIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
