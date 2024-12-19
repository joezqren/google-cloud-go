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

package cloudquotas

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"math"
	"net/http"
	"net/url"
	"time"

	cloudquotaspb "cloud.google.com/go/cloudquotas/apiv1/cloudquotaspb"
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

var newClientHook clientHook

// CallOptions contains the retry settings for each method of Client.
type CallOptions struct {
	ListQuotaInfos        []gax.CallOption
	GetQuotaInfo          []gax.CallOption
	ListQuotaPreferences  []gax.CallOption
	GetQuotaPreference    []gax.CallOption
	CreateQuotaPreference []gax.CallOption
	UpdateQuotaPreference []gax.CallOption
}

func defaultGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("cloudquotas.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("cloudquotas.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("cloudquotas.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://cloudquotas.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		internaloption.EnableNewAuthLibrary(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCallOptions() *CallOptions {
	return &CallOptions{
		ListQuotaInfos: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetQuotaInfo: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListQuotaPreferences: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetQuotaPreference: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		CreateQuotaPreference: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		UpdateQuotaPreference: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

func defaultRESTCallOptions() *CallOptions {
	return &CallOptions{
		ListQuotaInfos: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		GetQuotaInfo: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		ListQuotaPreferences: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		GetQuotaPreference: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		CreateQuotaPreference: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		UpdateQuotaPreference: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
	}
}

// internalClient is an interface that defines the methods available from Cloud Quotas API.
type internalClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListQuotaInfos(context.Context, *cloudquotaspb.ListQuotaInfosRequest, ...gax.CallOption) *QuotaInfoIterator
	GetQuotaInfo(context.Context, *cloudquotaspb.GetQuotaInfoRequest, ...gax.CallOption) (*cloudquotaspb.QuotaInfo, error)
	ListQuotaPreferences(context.Context, *cloudquotaspb.ListQuotaPreferencesRequest, ...gax.CallOption) *QuotaPreferenceIterator
	GetQuotaPreference(context.Context, *cloudquotaspb.GetQuotaPreferenceRequest, ...gax.CallOption) (*cloudquotaspb.QuotaPreference, error)
	CreateQuotaPreference(context.Context, *cloudquotaspb.CreateQuotaPreferenceRequest, ...gax.CallOption) (*cloudquotaspb.QuotaPreference, error)
	UpdateQuotaPreference(context.Context, *cloudquotaspb.UpdateQuotaPreferenceRequest, ...gax.CallOption) (*cloudquotaspb.QuotaPreference, error)
}

// Client is a client for interacting with Cloud Quotas API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The Cloud Quotas API is an infrastructure service for Google Cloud that lets
// service consumers list and manage their resource usage limits.
//
//	List/Get the metadata and current status of the quotas for a service.
//
//	Create/Update quota preferencess that declare the preferred quota values.
//
//	Check the status of a quota preference request.
//
//	List/Get pending and historical quota preference.
type Client struct {
	// The internal transport-dependent client.
	internalClient internalClient

	// The call options for this service.
	CallOptions *CallOptions
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
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *Client) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListQuotaInfos lists QuotaInfos of all quotas for a given project, folder or organization.
func (c *Client) ListQuotaInfos(ctx context.Context, req *cloudquotaspb.ListQuotaInfosRequest, opts ...gax.CallOption) *QuotaInfoIterator {
	return c.internalClient.ListQuotaInfos(ctx, req, opts...)
}

// GetQuotaInfo retrieve the QuotaInfo of a quota for a project, folder or organization.
func (c *Client) GetQuotaInfo(ctx context.Context, req *cloudquotaspb.GetQuotaInfoRequest, opts ...gax.CallOption) (*cloudquotaspb.QuotaInfo, error) {
	return c.internalClient.GetQuotaInfo(ctx, req, opts...)
}

// ListQuotaPreferences lists QuotaPreferences in a given project, folder or organization.
func (c *Client) ListQuotaPreferences(ctx context.Context, req *cloudquotaspb.ListQuotaPreferencesRequest, opts ...gax.CallOption) *QuotaPreferenceIterator {
	return c.internalClient.ListQuotaPreferences(ctx, req, opts...)
}

// GetQuotaPreference gets details of a single QuotaPreference.
func (c *Client) GetQuotaPreference(ctx context.Context, req *cloudquotaspb.GetQuotaPreferenceRequest, opts ...gax.CallOption) (*cloudquotaspb.QuotaPreference, error) {
	return c.internalClient.GetQuotaPreference(ctx, req, opts...)
}

// CreateQuotaPreference creates a new QuotaPreference that declares the desired value for a quota.
func (c *Client) CreateQuotaPreference(ctx context.Context, req *cloudquotaspb.CreateQuotaPreferenceRequest, opts ...gax.CallOption) (*cloudquotaspb.QuotaPreference, error) {
	return c.internalClient.CreateQuotaPreference(ctx, req, opts...)
}

// UpdateQuotaPreference updates the parameters of a single QuotaPreference. It can updates the
// config in any states, not just the ones pending approval.
func (c *Client) UpdateQuotaPreference(ctx context.Context, req *cloudquotaspb.UpdateQuotaPreferenceRequest, opts ...gax.CallOption) (*cloudquotaspb.QuotaPreference, error) {
	return c.internalClient.UpdateQuotaPreference(ctx, req, opts...)
}

// gRPCClient is a client for interacting with Cloud Quotas API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type gRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing Client
	CallOptions **CallOptions

	// The gRPC API client.
	client cloudquotaspb.CloudQuotasClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string

	logger *slog.Logger
}

// NewClient creates a new cloud quotas client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// The Cloud Quotas API is an infrastructure service for Google Cloud that lets
// service consumers list and manage their resource usage limits.
//
//	List/Get the metadata and current status of the quotas for a service.
//
//	Create/Update quota preferencess that declare the preferred quota values.
//
//	Check the status of a quota preference request.
//
//	List/Get pending and historical quota preference.
func NewClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	clientOpts := defaultGRPCClientOptions()
	if newClientHook != nil {
		hookOpts, err := newClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := Client{CallOptions: defaultCallOptions()}

	c := &gRPCClient{
		connPool:    connPool,
		client:      cloudquotaspb.NewCloudQuotasClient(connPool),
		CallOptions: &client.CallOptions,
		logger:      internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *gRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *gRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *gRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type restClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing Client
	CallOptions **CallOptions

	logger *slog.Logger
}

// NewRESTClient creates a new cloud quotas rest client.
//
// The Cloud Quotas API is an infrastructure service for Google Cloud that lets
// service consumers list and manage their resource usage limits.
//
//	List/Get the metadata and current status of the quotas for a service.
//
//	Create/Update quota preferencess that declare the preferred quota values.
//
//	Check the status of a quota preference request.
//
//	List/Get pending and historical quota preference.
func NewRESTClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	clientOpts := append(defaultRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultRESTCallOptions()
	c := &restClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
		logger:      internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	return &Client{internalClient: c, CallOptions: callOpts}, nil
}

func defaultRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://cloudquotas.googleapis.com"),
		internaloption.WithDefaultEndpointTemplate("https://cloudquotas.UNIVERSE_DOMAIN"),
		internaloption.WithDefaultMTLSEndpoint("https://cloudquotas.mtls.googleapis.com"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://cloudquotas.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableNewAuthLibrary(),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *restClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *restClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *restClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *gRPCClient) ListQuotaInfos(ctx context.Context, req *cloudquotaspb.ListQuotaInfosRequest, opts ...gax.CallOption) *QuotaInfoIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListQuotaInfos[0:len((*c.CallOptions).ListQuotaInfos):len((*c.CallOptions).ListQuotaInfos)], opts...)
	it := &QuotaInfoIterator{}
	req = proto.Clone(req).(*cloudquotaspb.ListQuotaInfosRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*cloudquotaspb.QuotaInfo, string, error) {
		resp := &cloudquotaspb.ListQuotaInfosResponse{}
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
			resp, err = executeRPC(ctx, c.client.ListQuotaInfos, req, settings.GRPC, c.logger, "ListQuotaInfos")
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetQuotaInfos(), resp.GetNextPageToken(), nil
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

func (c *gRPCClient) GetQuotaInfo(ctx context.Context, req *cloudquotaspb.GetQuotaInfoRequest, opts ...gax.CallOption) (*cloudquotaspb.QuotaInfo, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetQuotaInfo[0:len((*c.CallOptions).GetQuotaInfo):len((*c.CallOptions).GetQuotaInfo)], opts...)
	var resp *cloudquotaspb.QuotaInfo
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.client.GetQuotaInfo, req, settings.GRPC, c.logger, "GetQuotaInfo")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) ListQuotaPreferences(ctx context.Context, req *cloudquotaspb.ListQuotaPreferencesRequest, opts ...gax.CallOption) *QuotaPreferenceIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListQuotaPreferences[0:len((*c.CallOptions).ListQuotaPreferences):len((*c.CallOptions).ListQuotaPreferences)], opts...)
	it := &QuotaPreferenceIterator{}
	req = proto.Clone(req).(*cloudquotaspb.ListQuotaPreferencesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*cloudquotaspb.QuotaPreference, string, error) {
		resp := &cloudquotaspb.ListQuotaPreferencesResponse{}
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
			resp, err = executeRPC(ctx, c.client.ListQuotaPreferences, req, settings.GRPC, c.logger, "ListQuotaPreferences")
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetQuotaPreferences(), resp.GetNextPageToken(), nil
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

func (c *gRPCClient) GetQuotaPreference(ctx context.Context, req *cloudquotaspb.GetQuotaPreferenceRequest, opts ...gax.CallOption) (*cloudquotaspb.QuotaPreference, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetQuotaPreference[0:len((*c.CallOptions).GetQuotaPreference):len((*c.CallOptions).GetQuotaPreference)], opts...)
	var resp *cloudquotaspb.QuotaPreference
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.client.GetQuotaPreference, req, settings.GRPC, c.logger, "GetQuotaPreference")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) CreateQuotaPreference(ctx context.Context, req *cloudquotaspb.CreateQuotaPreferenceRequest, opts ...gax.CallOption) (*cloudquotaspb.QuotaPreference, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CreateQuotaPreference[0:len((*c.CallOptions).CreateQuotaPreference):len((*c.CallOptions).CreateQuotaPreference)], opts...)
	var resp *cloudquotaspb.QuotaPreference
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.client.CreateQuotaPreference, req, settings.GRPC, c.logger, "CreateQuotaPreference")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) UpdateQuotaPreference(ctx context.Context, req *cloudquotaspb.UpdateQuotaPreferenceRequest, opts ...gax.CallOption) (*cloudquotaspb.QuotaPreference, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "quota_preference.name", url.QueryEscape(req.GetQuotaPreference().GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).UpdateQuotaPreference[0:len((*c.CallOptions).UpdateQuotaPreference):len((*c.CallOptions).UpdateQuotaPreference)], opts...)
	var resp *cloudquotaspb.QuotaPreference
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.client.UpdateQuotaPreference, req, settings.GRPC, c.logger, "UpdateQuotaPreference")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ListQuotaInfos lists QuotaInfos of all quotas for a given project, folder or organization.
func (c *restClient) ListQuotaInfos(ctx context.Context, req *cloudquotaspb.ListQuotaInfosRequest, opts ...gax.CallOption) *QuotaInfoIterator {
	it := &QuotaInfoIterator{}
	req = proto.Clone(req).(*cloudquotaspb.ListQuotaInfosRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*cloudquotaspb.QuotaInfo, string, error) {
		resp := &cloudquotaspb.ListQuotaInfosResponse{}
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
		baseUrl.Path += fmt.Sprintf("/v1/%v/quotaInfos", req.GetParent())

		params := url.Values{}
		params.Add("$alt", "json;enum-encoding=int")
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

			buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "ListQuotaInfos")
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
		return resp.GetQuotaInfos(), resp.GetNextPageToken(), nil
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

// GetQuotaInfo retrieve the QuotaInfo of a quota for a project, folder or organization.
func (c *restClient) GetQuotaInfo(ctx context.Context, req *cloudquotaspb.GetQuotaInfoRequest, opts ...gax.CallOption) (*cloudquotaspb.QuotaInfo, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetQuotaInfo[0:len((*c.CallOptions).GetQuotaInfo):len((*c.CallOptions).GetQuotaInfo)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &cloudquotaspb.QuotaInfo{}
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

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "GetQuotaInfo")
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

// ListQuotaPreferences lists QuotaPreferences in a given project, folder or organization.
func (c *restClient) ListQuotaPreferences(ctx context.Context, req *cloudquotaspb.ListQuotaPreferencesRequest, opts ...gax.CallOption) *QuotaPreferenceIterator {
	it := &QuotaPreferenceIterator{}
	req = proto.Clone(req).(*cloudquotaspb.ListQuotaPreferencesRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*cloudquotaspb.QuotaPreference, string, error) {
		resp := &cloudquotaspb.ListQuotaPreferencesResponse{}
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
		baseUrl.Path += fmt.Sprintf("/v1/%v/quotaPreferences", req.GetParent())

		params := url.Values{}
		params.Add("$alt", "json;enum-encoding=int")
		if req.GetFilter() != "" {
			params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
		}
		if req.GetOrderBy() != "" {
			params.Add("orderBy", fmt.Sprintf("%v", req.GetOrderBy()))
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

			buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "ListQuotaPreferences")
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
		return resp.GetQuotaPreferences(), resp.GetNextPageToken(), nil
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

// GetQuotaPreference gets details of a single QuotaPreference.
func (c *restClient) GetQuotaPreference(ctx context.Context, req *cloudquotaspb.GetQuotaPreferenceRequest, opts ...gax.CallOption) (*cloudquotaspb.QuotaPreference, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetQuotaPreference[0:len((*c.CallOptions).GetQuotaPreference):len((*c.CallOptions).GetQuotaPreference)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &cloudquotaspb.QuotaPreference{}
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

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "GetQuotaPreference")
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

// CreateQuotaPreference creates a new QuotaPreference that declares the desired value for a quota.
func (c *restClient) CreateQuotaPreference(ctx context.Context, req *cloudquotaspb.CreateQuotaPreferenceRequest, opts ...gax.CallOption) (*cloudquotaspb.QuotaPreference, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetQuotaPreference()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v/quotaPreferences", req.GetParent())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")
	if items := req.GetIgnoreSafetyChecks(); len(items) > 0 {
		for _, item := range items {
			params.Add("ignoreSafetyChecks", fmt.Sprintf("%v", item))
		}
	}
	if req.GetQuotaPreferenceId() != "" {
		params.Add("quotaPreferenceId", fmt.Sprintf("%v", req.GetQuotaPreferenceId()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).CreateQuotaPreference[0:len((*c.CallOptions).CreateQuotaPreference):len((*c.CallOptions).CreateQuotaPreference)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &cloudquotaspb.QuotaPreference{}
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

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "CreateQuotaPreference")
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

// UpdateQuotaPreference updates the parameters of a single QuotaPreference. It can updates the
// config in any states, not just the ones pending approval.
func (c *restClient) UpdateQuotaPreference(ctx context.Context, req *cloudquotaspb.UpdateQuotaPreferenceRequest, opts ...gax.CallOption) (*cloudquotaspb.QuotaPreference, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetQuotaPreference()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v", req.GetQuotaPreference().GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")
	if req.GetAllowMissing() {
		params.Add("allowMissing", fmt.Sprintf("%v", req.GetAllowMissing()))
	}
	if items := req.GetIgnoreSafetyChecks(); len(items) > 0 {
		for _, item := range items {
			params.Add("ignoreSafetyChecks", fmt.Sprintf("%v", item))
		}
	}
	if req.GetUpdateMask() != nil {
		field, err := protojson.Marshal(req.GetUpdateMask())
		if err != nil {
			return nil, err
		}
		params.Add("updateMask", string(field[1:len(field)-1]))
	}
	if req.GetValidateOnly() {
		params.Add("validateOnly", fmt.Sprintf("%v", req.GetValidateOnly()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "quota_preference.name", url.QueryEscape(req.GetQuotaPreference().GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).UpdateQuotaPreference[0:len((*c.CallOptions).UpdateQuotaPreference):len((*c.CallOptions).UpdateQuotaPreference)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &cloudquotaspb.QuotaPreference{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("PATCH", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "UpdateQuotaPreference")
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
