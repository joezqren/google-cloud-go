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

package retail

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"math"
	"net/http"
	"net/url"
	"time"

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	retailpb "cloud.google.com/go/retail/apiv2alpha/retailpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newCompletionClientHook clientHook

// CompletionCallOptions contains the retry settings for each method of CompletionClient.
type CompletionCallOptions struct {
	CompleteQuery        []gax.CallOption
	ImportCompletionData []gax.CallOption
	GetOperation         []gax.CallOption
	ListOperations       []gax.CallOption
}

func defaultCompletionGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("retail.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("retail.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("retail.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://retail.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		internaloption.EnableNewAuthLibrary(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCompletionCallOptions() *CompletionCallOptions {
	return &CompletionCallOptions{
		CompleteQuery: []gax.CallOption{
			gax.WithTimeout(5000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        5000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ImportCompletionData: []gax.CallOption{
			gax.WithTimeout(5000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        5000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetOperation: []gax.CallOption{},
		ListOperations: []gax.CallOption{
			gax.WithTimeout(300000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        300000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

func defaultCompletionRESTCallOptions() *CompletionCallOptions {
	return &CompletionCallOptions{
		CompleteQuery: []gax.CallOption{
			gax.WithTimeout(5000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        5000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable,
					http.StatusGatewayTimeout)
			}),
		},
		ImportCompletionData: []gax.CallOption{
			gax.WithTimeout(5000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        5000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable,
					http.StatusGatewayTimeout)
			}),
		},
		GetOperation: []gax.CallOption{},
		ListOperations: []gax.CallOption{
			gax.WithTimeout(300000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        300000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable,
					http.StatusGatewayTimeout)
			}),
		},
	}
}

// internalCompletionClient is an interface that defines the methods available from Vertex AI Search for Retail API.
type internalCompletionClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CompleteQuery(context.Context, *retailpb.CompleteQueryRequest, ...gax.CallOption) (*retailpb.CompleteQueryResponse, error)
	ImportCompletionData(context.Context, *retailpb.ImportCompletionDataRequest, ...gax.CallOption) (*ImportCompletionDataOperation, error)
	ImportCompletionDataOperation(name string) *ImportCompletionDataOperation
	GetOperation(context.Context, *longrunningpb.GetOperationRequest, ...gax.CallOption) (*longrunningpb.Operation, error)
	ListOperations(context.Context, *longrunningpb.ListOperationsRequest, ...gax.CallOption) *OperationIterator
}

// CompletionClient is a client for interacting with Vertex AI Search for Retail API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Autocomplete service for retail.
//
// This feature is only available for users who have Retail Search enabled.
// Enable Retail Search on Cloud Console before using this feature.
type CompletionClient struct {
	// The internal transport-dependent client.
	internalClient internalCompletionClient

	// The call options for this service.
	CallOptions *CompletionCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *CompletionClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *CompletionClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *CompletionClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CompleteQuery completes the specified prefix with keyword suggestions.
//
// This feature is only available for users who have Retail Search enabled.
// Enable Retail Search on Cloud Console before using this feature.
func (c *CompletionClient) CompleteQuery(ctx context.Context, req *retailpb.CompleteQueryRequest, opts ...gax.CallOption) (*retailpb.CompleteQueryResponse, error) {
	return c.internalClient.CompleteQuery(ctx, req, opts...)
}

// ImportCompletionData bulk import of processed completion dataset.
//
// Request processing is asynchronous. Partial updating is not supported.
//
// The operation is successfully finished only after the imported suggestions
// are indexed successfully and ready for serving. The process takes hours.
//
// This feature is only available for users who have Retail Search enabled.
// Enable Retail Search on Cloud Console before using this feature.
func (c *CompletionClient) ImportCompletionData(ctx context.Context, req *retailpb.ImportCompletionDataRequest, opts ...gax.CallOption) (*ImportCompletionDataOperation, error) {
	return c.internalClient.ImportCompletionData(ctx, req, opts...)
}

// ImportCompletionDataOperation returns a new ImportCompletionDataOperation from a given name.
// The name must be that of a previously created ImportCompletionDataOperation, possibly from a different process.
func (c *CompletionClient) ImportCompletionDataOperation(name string) *ImportCompletionDataOperation {
	return c.internalClient.ImportCompletionDataOperation(name)
}

// GetOperation is a utility method from google.longrunning.Operations.
func (c *CompletionClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	return c.internalClient.GetOperation(ctx, req, opts...)
}

// ListOperations is a utility method from google.longrunning.Operations.
func (c *CompletionClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	return c.internalClient.ListOperations(ctx, req, opts...)
}

// completionGRPCClient is a client for interacting with Vertex AI Search for Retail API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type completionGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing CompletionClient
	CallOptions **CompletionCallOptions

	// The gRPC API client.
	completionClient retailpb.CompletionServiceClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	operationsClient longrunningpb.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string

	logger *slog.Logger
}

// NewCompletionClient creates a new completion service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Autocomplete service for retail.
//
// This feature is only available for users who have Retail Search enabled.
// Enable Retail Search on Cloud Console before using this feature.
func NewCompletionClient(ctx context.Context, opts ...option.ClientOption) (*CompletionClient, error) {
	clientOpts := defaultCompletionGRPCClientOptions()
	if newCompletionClientHook != nil {
		hookOpts, err := newCompletionClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := CompletionClient{CallOptions: defaultCompletionCallOptions()}

	c := &completionGRPCClient{
		connPool:         connPool,
		completionClient: retailpb.NewCompletionServiceClient(connPool),
		CallOptions:      &client.CallOptions,
		logger:           internaloption.GetLogger(opts),
		operationsClient: longrunningpb.NewOperationsClient(connPool),
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
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *completionGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *completionGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *completionGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type completionRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing CompletionClient
	CallOptions **CompletionCallOptions

	logger *slog.Logger
}

// NewCompletionRESTClient creates a new completion service rest client.
//
// Autocomplete service for retail.
//
// This feature is only available for users who have Retail Search enabled.
// Enable Retail Search on Cloud Console before using this feature.
func NewCompletionRESTClient(ctx context.Context, opts ...option.ClientOption) (*CompletionClient, error) {
	clientOpts := append(defaultCompletionRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultCompletionRESTCallOptions()
	c := &completionRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
		logger:      internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	lroOpts := []option.ClientOption{
		option.WithHTTPClient(httpClient),
		option.WithEndpoint(endpoint),
	}
	opClient, err := lroauto.NewOperationsRESTClient(ctx, lroOpts...)
	if err != nil {
		return nil, err
	}
	c.LROClient = &opClient

	return &CompletionClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultCompletionRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://retail.googleapis.com"),
		internaloption.WithDefaultEndpointTemplate("https://retail.UNIVERSE_DOMAIN"),
		internaloption.WithDefaultMTLSEndpoint("https://retail.mtls.googleapis.com"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://retail.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableNewAuthLibrary(),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *completionRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *completionRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *completionRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *completionGRPCClient) CompleteQuery(ctx context.Context, req *retailpb.CompleteQueryRequest, opts ...gax.CallOption) (*retailpb.CompleteQueryResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "catalog", url.QueryEscape(req.GetCatalog()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CompleteQuery[0:len((*c.CallOptions).CompleteQuery):len((*c.CallOptions).CompleteQuery)], opts...)
	var resp *retailpb.CompleteQueryResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.completionClient.CompleteQuery, req, settings.GRPC, c.logger, "CompleteQuery")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *completionGRPCClient) ImportCompletionData(ctx context.Context, req *retailpb.ImportCompletionDataRequest, opts ...gax.CallOption) (*ImportCompletionDataOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ImportCompletionData[0:len((*c.CallOptions).ImportCompletionData):len((*c.CallOptions).ImportCompletionData)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.completionClient.ImportCompletionData, req, settings.GRPC, c.logger, "ImportCompletionData")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &ImportCompletionDataOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *completionGRPCClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.operationsClient.GetOperation, req, settings.GRPC, c.logger, "GetOperation")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *completionGRPCClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListOperations[0:len((*c.CallOptions).ListOperations):len((*c.CallOptions).ListOperations)], opts...)
	it := &OperationIterator{}
	req = proto.Clone(req).(*longrunningpb.ListOperationsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*longrunningpb.Operation, string, error) {
		resp := &longrunningpb.ListOperationsResponse{}
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
			resp, err = executeRPC(ctx, c.operationsClient.ListOperations, req, settings.GRPC, c.logger, "ListOperations")
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetOperations(), resp.GetNextPageToken(), nil
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

// CompleteQuery completes the specified prefix with keyword suggestions.
//
// This feature is only available for users who have Retail Search enabled.
// Enable Retail Search on Cloud Console before using this feature.
func (c *completionRESTClient) CompleteQuery(ctx context.Context, req *retailpb.CompleteQueryRequest, opts ...gax.CallOption) (*retailpb.CompleteQueryResponse, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v2alpha/%v:completeQuery", req.GetCatalog())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")
	if req.GetDataset() != "" {
		params.Add("dataset", fmt.Sprintf("%v", req.GetDataset()))
	}
	if req.GetDeviceType() != "" {
		params.Add("deviceType", fmt.Sprintf("%v", req.GetDeviceType()))
	}
	if req.GetEnableAttributeSuggestions() {
		params.Add("enableAttributeSuggestions", fmt.Sprintf("%v", req.GetEnableAttributeSuggestions()))
	}
	if req.GetEntity() != "" {
		params.Add("entity", fmt.Sprintf("%v", req.GetEntity()))
	}
	if items := req.GetLanguageCodes(); len(items) > 0 {
		for _, item := range items {
			params.Add("languageCodes", fmt.Sprintf("%v", item))
		}
	}
	if req.GetMaxSuggestions() != 0 {
		params.Add("maxSuggestions", fmt.Sprintf("%v", req.GetMaxSuggestions()))
	}
	params.Add("query", fmt.Sprintf("%v", req.GetQuery()))
	if req.GetVisitorId() != "" {
		params.Add("visitorId", fmt.Sprintf("%v", req.GetVisitorId()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "catalog", url.QueryEscape(req.GetCatalog()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).CompleteQuery[0:len((*c.CallOptions).CompleteQuery):len((*c.CallOptions).CompleteQuery)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &retailpb.CompleteQueryResponse{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "CompleteQuery")
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// ImportCompletionData bulk import of processed completion dataset.
//
// Request processing is asynchronous. Partial updating is not supported.
//
// The operation is successfully finished only after the imported suggestions
// are indexed successfully and ready for serving. The process takes hours.
//
// This feature is only available for users who have Retail Search enabled.
// Enable Retail Search on Cloud Console before using this feature.
func (c *completionRESTClient) ImportCompletionData(ctx context.Context, req *retailpb.ImportCompletionDataRequest, opts ...gax.CallOption) (*ImportCompletionDataOperation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v2alpha/%v/completionData:import", req.GetParent())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &longrunningpb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "ImportCompletionData")
		if err != nil {
			return err
		}
		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}

	override := fmt.Sprintf("/v2alpha/%s", resp.GetName())
	return &ImportCompletionDataOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, resp),
		pollPath: override,
	}, nil
}

// GetOperation is a utility method from google.longrunning.Operations.
func (c *completionRESTClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v2alpha/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &longrunningpb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "GetOperation")
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// ListOperations is a utility method from google.longrunning.Operations.
func (c *completionRESTClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	it := &OperationIterator{}
	req = proto.Clone(req).(*longrunningpb.ListOperationsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*longrunningpb.Operation, string, error) {
		resp := &longrunningpb.ListOperationsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		baseUrl, err := url.Parse(c.endpoint)
		if err != nil {
			return nil, "", err
		}
		baseUrl.Path += fmt.Sprintf("/v2alpha/%v/operations", req.GetName())

		params := url.Values{}
		params.Add("$alt", "json;enum-encoding=int")
		if req.GetFilter() != "" {
			params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
		}
		if req.GetPageSize() != 0 {
			params.Add("pageSize", fmt.Sprintf("%v", req.GetPageSize()))
		}
		if req.GetPageToken() != "" {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}

		baseUrl.RawQuery = params.Encode()

		// Build HTTP headers from client and context metadata.
		hds := append(c.xGoogHeaders, "Content-Type", "application/json")
		headers := gax.BuildHeaders(ctx, hds...)
		e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			if settings.Path != "" {
				baseUrl.Path = settings.Path
			}
			httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
			if err != nil {
				return err
			}
			httpReq.Header = headers

			buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "ListOperations")
			if err != nil {
				return err
			}
			if err := unm.Unmarshal(buf, resp); err != nil {
				return err
			}

			return nil
		}, opts...)
		if e != nil {
			return nil, "", e
		}
		it.Response = resp
		return resp.GetOperations(), resp.GetNextPageToken(), nil
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

// ImportCompletionDataOperation returns a new ImportCompletionDataOperation from a given name.
// The name must be that of a previously created ImportCompletionDataOperation, possibly from a different process.
func (c *completionGRPCClient) ImportCompletionDataOperation(name string) *ImportCompletionDataOperation {
	return &ImportCompletionDataOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// ImportCompletionDataOperation returns a new ImportCompletionDataOperation from a given name.
// The name must be that of a previously created ImportCompletionDataOperation, possibly from a different process.
func (c *completionRESTClient) ImportCompletionDataOperation(name string) *ImportCompletionDataOperation {
	override := fmt.Sprintf("/v2alpha/%s", name)
	return &ImportCompletionDataOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
		pollPath: override,
	}
}
