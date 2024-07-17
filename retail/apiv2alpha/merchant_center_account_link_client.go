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
	"io"
	"math"
	"net/http"
	"net/url"
	"time"

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	retailpb "cloud.google.com/go/retail/apiv2alpha/retailpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
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

var newMerchantCenterAccountLinkClientHook clientHook

// MerchantCenterAccountLinkCallOptions contains the retry settings for each method of MerchantCenterAccountLinkClient.
type MerchantCenterAccountLinkCallOptions struct {
	ListMerchantCenterAccountLinks  []gax.CallOption
	CreateMerchantCenterAccountLink []gax.CallOption
	DeleteMerchantCenterAccountLink []gax.CallOption
	GetOperation                    []gax.CallOption
	ListOperations                  []gax.CallOption
}

func defaultMerchantCenterAccountLinkGRPCClientOptions() []option.ClientOption {
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

func defaultMerchantCenterAccountLinkCallOptions() *MerchantCenterAccountLinkCallOptions {
	return &MerchantCenterAccountLinkCallOptions{
		ListMerchantCenterAccountLinks:  []gax.CallOption{},
		CreateMerchantCenterAccountLink: []gax.CallOption{},
		DeleteMerchantCenterAccountLink: []gax.CallOption{},
		GetOperation:                    []gax.CallOption{},
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

func defaultMerchantCenterAccountLinkRESTCallOptions() *MerchantCenterAccountLinkCallOptions {
	return &MerchantCenterAccountLinkCallOptions{
		ListMerchantCenterAccountLinks:  []gax.CallOption{},
		CreateMerchantCenterAccountLink: []gax.CallOption{},
		DeleteMerchantCenterAccountLink: []gax.CallOption{},
		GetOperation:                    []gax.CallOption{},
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

// internalMerchantCenterAccountLinkClient is an interface that defines the methods available from Vertex AI Search for Retail API.
type internalMerchantCenterAccountLinkClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListMerchantCenterAccountLinks(context.Context, *retailpb.ListMerchantCenterAccountLinksRequest, ...gax.CallOption) (*retailpb.ListMerchantCenterAccountLinksResponse, error)
	CreateMerchantCenterAccountLink(context.Context, *retailpb.CreateMerchantCenterAccountLinkRequest, ...gax.CallOption) (*CreateMerchantCenterAccountLinkOperation, error)
	CreateMerchantCenterAccountLinkOperation(name string) *CreateMerchantCenterAccountLinkOperation
	DeleteMerchantCenterAccountLink(context.Context, *retailpb.DeleteMerchantCenterAccountLinkRequest, ...gax.CallOption) error
	GetOperation(context.Context, *longrunningpb.GetOperationRequest, ...gax.CallOption) (*longrunningpb.Operation, error)
	ListOperations(context.Context, *longrunningpb.ListOperationsRequest, ...gax.CallOption) *OperationIterator
}

// MerchantCenterAccountLinkClient is a client for interacting with Vertex AI Search for Retail API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Merchant Center Link service to link a Branch to a Merchant Center Account.
type MerchantCenterAccountLinkClient struct {
	// The internal transport-dependent client.
	internalClient internalMerchantCenterAccountLinkClient

	// The call options for this service.
	CallOptions *MerchantCenterAccountLinkCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *MerchantCenterAccountLinkClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *MerchantCenterAccountLinkClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *MerchantCenterAccountLinkClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListMerchantCenterAccountLinks lists all
// MerchantCenterAccountLinks
// under the specified parent Catalog.
func (c *MerchantCenterAccountLinkClient) ListMerchantCenterAccountLinks(ctx context.Context, req *retailpb.ListMerchantCenterAccountLinksRequest, opts ...gax.CallOption) (*retailpb.ListMerchantCenterAccountLinksResponse, error) {
	return c.internalClient.ListMerchantCenterAccountLinks(ctx, req, opts...)
}

// CreateMerchantCenterAccountLink creates a
// MerchantCenterAccountLink.
func (c *MerchantCenterAccountLinkClient) CreateMerchantCenterAccountLink(ctx context.Context, req *retailpb.CreateMerchantCenterAccountLinkRequest, opts ...gax.CallOption) (*CreateMerchantCenterAccountLinkOperation, error) {
	return c.internalClient.CreateMerchantCenterAccountLink(ctx, req, opts...)
}

// CreateMerchantCenterAccountLinkOperation returns a new CreateMerchantCenterAccountLinkOperation from a given name.
// The name must be that of a previously created CreateMerchantCenterAccountLinkOperation, possibly from a different process.
func (c *MerchantCenterAccountLinkClient) CreateMerchantCenterAccountLinkOperation(name string) *CreateMerchantCenterAccountLinkOperation {
	return c.internalClient.CreateMerchantCenterAccountLinkOperation(name)
}

// DeleteMerchantCenterAccountLink deletes a
// MerchantCenterAccountLink.
// If the
// MerchantCenterAccountLink
// to delete does not exist, a NOT_FOUND error is returned.
func (c *MerchantCenterAccountLinkClient) DeleteMerchantCenterAccountLink(ctx context.Context, req *retailpb.DeleteMerchantCenterAccountLinkRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteMerchantCenterAccountLink(ctx, req, opts...)
}

// GetOperation is a utility method from google.longrunning.Operations.
func (c *MerchantCenterAccountLinkClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	return c.internalClient.GetOperation(ctx, req, opts...)
}

// ListOperations is a utility method from google.longrunning.Operations.
func (c *MerchantCenterAccountLinkClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	return c.internalClient.ListOperations(ctx, req, opts...)
}

// merchantCenterAccountLinkGRPCClient is a client for interacting with Vertex AI Search for Retail API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type merchantCenterAccountLinkGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing MerchantCenterAccountLinkClient
	CallOptions **MerchantCenterAccountLinkCallOptions

	// The gRPC API client.
	merchantCenterAccountLinkClient retailpb.MerchantCenterAccountLinkServiceClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	operationsClient longrunningpb.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewMerchantCenterAccountLinkClient creates a new merchant center account link service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Merchant Center Link service to link a Branch to a Merchant Center Account.
func NewMerchantCenterAccountLinkClient(ctx context.Context, opts ...option.ClientOption) (*MerchantCenterAccountLinkClient, error) {
	clientOpts := defaultMerchantCenterAccountLinkGRPCClientOptions()
	if newMerchantCenterAccountLinkClientHook != nil {
		hookOpts, err := newMerchantCenterAccountLinkClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := MerchantCenterAccountLinkClient{CallOptions: defaultMerchantCenterAccountLinkCallOptions()}

	c := &merchantCenterAccountLinkGRPCClient{
		connPool:                        connPool,
		merchantCenterAccountLinkClient: retailpb.NewMerchantCenterAccountLinkServiceClient(connPool),
		CallOptions:                     &client.CallOptions,
		operationsClient:                longrunningpb.NewOperationsClient(connPool),
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
func (c *merchantCenterAccountLinkGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *merchantCenterAccountLinkGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *merchantCenterAccountLinkGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type merchantCenterAccountLinkRESTClient struct {
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

	// Points back to the CallOptions field of the containing MerchantCenterAccountLinkClient
	CallOptions **MerchantCenterAccountLinkCallOptions
}

// NewMerchantCenterAccountLinkRESTClient creates a new merchant center account link service rest client.
//
// Merchant Center Link service to link a Branch to a Merchant Center Account.
func NewMerchantCenterAccountLinkRESTClient(ctx context.Context, opts ...option.ClientOption) (*MerchantCenterAccountLinkClient, error) {
	clientOpts := append(defaultMerchantCenterAccountLinkRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultMerchantCenterAccountLinkRESTCallOptions()
	c := &merchantCenterAccountLinkRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
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

	return &MerchantCenterAccountLinkClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultMerchantCenterAccountLinkRESTClientOptions() []option.ClientOption {
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
func (c *merchantCenterAccountLinkRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *merchantCenterAccountLinkRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *merchantCenterAccountLinkRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *merchantCenterAccountLinkGRPCClient) ListMerchantCenterAccountLinks(ctx context.Context, req *retailpb.ListMerchantCenterAccountLinksRequest, opts ...gax.CallOption) (*retailpb.ListMerchantCenterAccountLinksResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListMerchantCenterAccountLinks[0:len((*c.CallOptions).ListMerchantCenterAccountLinks):len((*c.CallOptions).ListMerchantCenterAccountLinks)], opts...)
	var resp *retailpb.ListMerchantCenterAccountLinksResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.merchantCenterAccountLinkClient.ListMerchantCenterAccountLinks(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *merchantCenterAccountLinkGRPCClient) CreateMerchantCenterAccountLink(ctx context.Context, req *retailpb.CreateMerchantCenterAccountLinkRequest, opts ...gax.CallOption) (*CreateMerchantCenterAccountLinkOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CreateMerchantCenterAccountLink[0:len((*c.CallOptions).CreateMerchantCenterAccountLink):len((*c.CallOptions).CreateMerchantCenterAccountLink)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.merchantCenterAccountLinkClient.CreateMerchantCenterAccountLink(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &CreateMerchantCenterAccountLinkOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *merchantCenterAccountLinkGRPCClient) DeleteMerchantCenterAccountLink(ctx context.Context, req *retailpb.DeleteMerchantCenterAccountLinkRequest, opts ...gax.CallOption) error {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeleteMerchantCenterAccountLink[0:len((*c.CallOptions).DeleteMerchantCenterAccountLink):len((*c.CallOptions).DeleteMerchantCenterAccountLink)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.merchantCenterAccountLinkClient.DeleteMerchantCenterAccountLink(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *merchantCenterAccountLinkGRPCClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.operationsClient.GetOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *merchantCenterAccountLinkGRPCClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
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
			resp, err = c.operationsClient.ListOperations(ctx, req, settings.GRPC...)
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

// ListMerchantCenterAccountLinks lists all
// MerchantCenterAccountLinks
// under the specified parent Catalog.
func (c *merchantCenterAccountLinkRESTClient) ListMerchantCenterAccountLinks(ctx context.Context, req *retailpb.ListMerchantCenterAccountLinksRequest, opts ...gax.CallOption) (*retailpb.ListMerchantCenterAccountLinksResponse, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v2alpha/%v/merchantCenterAccountLinks", req.GetParent())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).ListMerchantCenterAccountLinks[0:len((*c.CallOptions).ListMerchantCenterAccountLinks):len((*c.CallOptions).ListMerchantCenterAccountLinks)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &retailpb.ListMerchantCenterAccountLinksResponse{}
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

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
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

// CreateMerchantCenterAccountLink creates a
// MerchantCenterAccountLink.
func (c *merchantCenterAccountLinkRESTClient) CreateMerchantCenterAccountLink(ctx context.Context, req *retailpb.CreateMerchantCenterAccountLinkRequest, opts ...gax.CallOption) (*CreateMerchantCenterAccountLinkOperation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetMerchantCenterAccountLink()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v2alpha/%v/merchantCenterAccountLinks", req.GetParent())

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

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
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
	return &CreateMerchantCenterAccountLinkOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, resp),
		pollPath: override,
	}, nil
}

// DeleteMerchantCenterAccountLink deletes a
// MerchantCenterAccountLink.
// If the
// MerchantCenterAccountLink
// to delete does not exist, a NOT_FOUND error is returned.
func (c *merchantCenterAccountLinkRESTClient) DeleteMerchantCenterAccountLink(ctx context.Context, req *retailpb.DeleteMerchantCenterAccountLinkRequest, opts ...gax.CallOption) error {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return err
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
	return gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("DELETE", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		// Returns nil if there is no error, otherwise wraps
		// the response code and body into a non-nil error
		return googleapi.CheckResponse(httpRsp)
	}, opts...)
}

// GetOperation is a utility method from google.longrunning.Operations.
func (c *merchantCenterAccountLinkRESTClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
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

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
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
func (c *merchantCenterAccountLinkRESTClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
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

			httpRsp, err := c.httpClient.Do(httpReq)
			if err != nil {
				return err
			}
			defer httpRsp.Body.Close()

			if err = googleapi.CheckResponse(httpRsp); err != nil {
				return err
			}

			buf, err := io.ReadAll(httpRsp.Body)
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

// CreateMerchantCenterAccountLinkOperation returns a new CreateMerchantCenterAccountLinkOperation from a given name.
// The name must be that of a previously created CreateMerchantCenterAccountLinkOperation, possibly from a different process.
func (c *merchantCenterAccountLinkGRPCClient) CreateMerchantCenterAccountLinkOperation(name string) *CreateMerchantCenterAccountLinkOperation {
	return &CreateMerchantCenterAccountLinkOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// CreateMerchantCenterAccountLinkOperation returns a new CreateMerchantCenterAccountLinkOperation from a given name.
// The name must be that of a previously created CreateMerchantCenterAccountLinkOperation, possibly from a different process.
func (c *merchantCenterAccountLinkRESTClient) CreateMerchantCenterAccountLinkOperation(name string) *CreateMerchantCenterAccountLinkOperation {
	override := fmt.Sprintf("/v2alpha/%s", name)
	return &CreateMerchantCenterAccountLinkOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
		pollPath: override,
	}
}
