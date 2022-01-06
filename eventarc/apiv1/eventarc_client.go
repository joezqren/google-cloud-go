// Copyright 2022 Google LLC
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

package eventarc

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	eventarcpb "google.golang.org/genproto/googleapis/cloud/eventarc/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newClientHook clientHook

// CallOptions contains the retry settings for each method of Client.
type CallOptions struct {
	GetTrigger    []gax.CallOption
	ListTriggers  []gax.CallOption
	CreateTrigger []gax.CallOption
	UpdateTrigger []gax.CallOption
	DeleteTrigger []gax.CallOption
}

func defaultGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("eventarc.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("eventarc.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://eventarc.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCallOptions() *CallOptions {
	return &CallOptions{
		GetTrigger:    []gax.CallOption{},
		ListTriggers:  []gax.CallOption{},
		CreateTrigger: []gax.CallOption{},
		UpdateTrigger: []gax.CallOption{},
		DeleteTrigger: []gax.CallOption{},
	}
}

// internalClient is an interface that defines the methods availaible from Eventarc API.
type internalClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetTrigger(context.Context, *eventarcpb.GetTriggerRequest, ...gax.CallOption) (*eventarcpb.Trigger, error)
	ListTriggers(context.Context, *eventarcpb.ListTriggersRequest, ...gax.CallOption) *TriggerIterator
	CreateTrigger(context.Context, *eventarcpb.CreateTriggerRequest, ...gax.CallOption) (*CreateTriggerOperation, error)
	CreateTriggerOperation(name string) *CreateTriggerOperation
	UpdateTrigger(context.Context, *eventarcpb.UpdateTriggerRequest, ...gax.CallOption) (*UpdateTriggerOperation, error)
	UpdateTriggerOperation(name string) *UpdateTriggerOperation
	DeleteTrigger(context.Context, *eventarcpb.DeleteTriggerRequest, ...gax.CallOption) (*DeleteTriggerOperation, error)
	DeleteTriggerOperation(name string) *DeleteTriggerOperation
}

// Client is a client for interacting with Eventarc API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Eventarc allows users to subscribe to various events that are provided by
// Google Cloud services and forward them to supported destinations.
type Client struct {
	// The internal transport-dependent client.
	internalClient internalClient

	// The call options for this service.
	CallOptions *CallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *Client) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *Client) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *Client) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetTrigger get a single trigger.
func (c *Client) GetTrigger(ctx context.Context, req *eventarcpb.GetTriggerRequest, opts ...gax.CallOption) (*eventarcpb.Trigger, error) {
	return c.internalClient.GetTrigger(ctx, req, opts...)
}

// ListTriggers list triggers.
func (c *Client) ListTriggers(ctx context.Context, req *eventarcpb.ListTriggersRequest, opts ...gax.CallOption) *TriggerIterator {
	return c.internalClient.ListTriggers(ctx, req, opts...)
}

// CreateTrigger create a new trigger in a particular project and location.
func (c *Client) CreateTrigger(ctx context.Context, req *eventarcpb.CreateTriggerRequest, opts ...gax.CallOption) (*CreateTriggerOperation, error) {
	return c.internalClient.CreateTrigger(ctx, req, opts...)
}

// CreateTriggerOperation returns a new CreateTriggerOperation from a given name.
// The name must be that of a previously created CreateTriggerOperation, possibly from a different process.
func (c *Client) CreateTriggerOperation(name string) *CreateTriggerOperation {
	return c.internalClient.CreateTriggerOperation(name)
}

// UpdateTrigger update a single trigger.
func (c *Client) UpdateTrigger(ctx context.Context, req *eventarcpb.UpdateTriggerRequest, opts ...gax.CallOption) (*UpdateTriggerOperation, error) {
	return c.internalClient.UpdateTrigger(ctx, req, opts...)
}

// UpdateTriggerOperation returns a new UpdateTriggerOperation from a given name.
// The name must be that of a previously created UpdateTriggerOperation, possibly from a different process.
func (c *Client) UpdateTriggerOperation(name string) *UpdateTriggerOperation {
	return c.internalClient.UpdateTriggerOperation(name)
}

// DeleteTrigger delete a single trigger.
func (c *Client) DeleteTrigger(ctx context.Context, req *eventarcpb.DeleteTriggerRequest, opts ...gax.CallOption) (*DeleteTriggerOperation, error) {
	return c.internalClient.DeleteTrigger(ctx, req, opts...)
}

// DeleteTriggerOperation returns a new DeleteTriggerOperation from a given name.
// The name must be that of a previously created DeleteTriggerOperation, possibly from a different process.
func (c *Client) DeleteTriggerOperation(name string) *DeleteTriggerOperation {
	return c.internalClient.DeleteTriggerOperation(name)
}

// gRPCClient is a client for interacting with Eventarc API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type gRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing Client
	CallOptions **CallOptions

	// The gRPC API client.
	client eventarcpb.EventarcClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewClient creates a new eventarc client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Eventarc allows users to subscribe to various events that are provided by
// Google Cloud services and forward them to supported destinations.
func NewClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	clientOpts := defaultGRPCClientOptions()
	if newClientHook != nil {
		hookOpts, err := newClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := Client{CallOptions: defaultCallOptions()}

	c := &gRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		client:           eventarcpb.NewEventarcClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	client.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}
	c.LROClient = &client.LROClient
	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *gRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *gRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *gRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *gRPCClient) GetTrigger(ctx context.Context, req *eventarcpb.GetTriggerRequest, opts ...gax.CallOption) (*eventarcpb.Trigger, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetTrigger[0:len((*c.CallOptions).GetTrigger):len((*c.CallOptions).GetTrigger)], opts...)
	var resp *eventarcpb.Trigger
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetTrigger(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) ListTriggers(ctx context.Context, req *eventarcpb.ListTriggersRequest, opts ...gax.CallOption) *TriggerIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListTriggers[0:len((*c.CallOptions).ListTriggers):len((*c.CallOptions).ListTriggers)], opts...)
	it := &TriggerIterator{}
	req = proto.Clone(req).(*eventarcpb.ListTriggersRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*eventarcpb.Trigger, string, error) {
		resp := &eventarcpb.ListTriggersResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.client.ListTriggers(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetTriggers(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *gRPCClient) CreateTrigger(ctx context.Context, req *eventarcpb.CreateTriggerRequest, opts ...gax.CallOption) (*CreateTriggerOperation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateTrigger[0:len((*c.CallOptions).CreateTrigger):len((*c.CallOptions).CreateTrigger)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.CreateTrigger(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &CreateTriggerOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *gRPCClient) UpdateTrigger(ctx context.Context, req *eventarcpb.UpdateTriggerRequest, opts ...gax.CallOption) (*UpdateTriggerOperation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "trigger.name", url.QueryEscape(req.GetTrigger().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateTrigger[0:len((*c.CallOptions).UpdateTrigger):len((*c.CallOptions).UpdateTrigger)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.UpdateTrigger(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &UpdateTriggerOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *gRPCClient) DeleteTrigger(ctx context.Context, req *eventarcpb.DeleteTriggerRequest, opts ...gax.CallOption) (*DeleteTriggerOperation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteTrigger[0:len((*c.CallOptions).DeleteTrigger):len((*c.CallOptions).DeleteTrigger)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.DeleteTrigger(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &DeleteTriggerOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

// CreateTriggerOperation manages a long-running operation from CreateTrigger.
type CreateTriggerOperation struct {
	lro *longrunning.Operation
}

// CreateTriggerOperation returns a new CreateTriggerOperation from a given name.
// The name must be that of a previously created CreateTriggerOperation, possibly from a different process.
func (c *gRPCClient) CreateTriggerOperation(name string) *CreateTriggerOperation {
	return &CreateTriggerOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *CreateTriggerOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*eventarcpb.Trigger, error) {
	var resp eventarcpb.Trigger
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
func (op *CreateTriggerOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*eventarcpb.Trigger, error) {
	var resp eventarcpb.Trigger
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
func (op *CreateTriggerOperation) Metadata() (*eventarcpb.OperationMetadata, error) {
	var meta eventarcpb.OperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *CreateTriggerOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *CreateTriggerOperation) Name() string {
	return op.lro.Name()
}

// DeleteTriggerOperation manages a long-running operation from DeleteTrigger.
type DeleteTriggerOperation struct {
	lro *longrunning.Operation
}

// DeleteTriggerOperation returns a new DeleteTriggerOperation from a given name.
// The name must be that of a previously created DeleteTriggerOperation, possibly from a different process.
func (c *gRPCClient) DeleteTriggerOperation(name string) *DeleteTriggerOperation {
	return &DeleteTriggerOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *DeleteTriggerOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*eventarcpb.Trigger, error) {
	var resp eventarcpb.Trigger
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
func (op *DeleteTriggerOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*eventarcpb.Trigger, error) {
	var resp eventarcpb.Trigger
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
func (op *DeleteTriggerOperation) Metadata() (*eventarcpb.OperationMetadata, error) {
	var meta eventarcpb.OperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *DeleteTriggerOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *DeleteTriggerOperation) Name() string {
	return op.lro.Name()
}

// UpdateTriggerOperation manages a long-running operation from UpdateTrigger.
type UpdateTriggerOperation struct {
	lro *longrunning.Operation
}

// UpdateTriggerOperation returns a new UpdateTriggerOperation from a given name.
// The name must be that of a previously created UpdateTriggerOperation, possibly from a different process.
func (c *gRPCClient) UpdateTriggerOperation(name string) *UpdateTriggerOperation {
	return &UpdateTriggerOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *UpdateTriggerOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*eventarcpb.Trigger, error) {
	var resp eventarcpb.Trigger
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
func (op *UpdateTriggerOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*eventarcpb.Trigger, error) {
	var resp eventarcpb.Trigger
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
func (op *UpdateTriggerOperation) Metadata() (*eventarcpb.OperationMetadata, error) {
	var meta eventarcpb.OperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *UpdateTriggerOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *UpdateTriggerOperation) Name() string {
	return op.lro.Name()
}

// TriggerIterator manages a stream of *eventarcpb.Trigger.
type TriggerIterator struct {
	items    []*eventarcpb.Trigger
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
	InternalFetch func(pageSize int, pageToken string) (results []*eventarcpb.Trigger, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *TriggerIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *TriggerIterator) Next() (*eventarcpb.Trigger, error) {
	var item *eventarcpb.Trigger
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *TriggerIterator) bufLen() int {
	return len(it.items)
}

func (it *TriggerIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
