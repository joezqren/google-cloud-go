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

package binaryauthorization

import (
	"context"
	"fmt"
	"log/slog"
	"math"
	"net/http"
	"net/url"

	binaryauthorizationpb "cloud.google.com/go/binaryauthorization/apiv1beta1/binaryauthorizationpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

var newSystemPolicyV1Beta1ClientHook clientHook

// SystemPolicyV1Beta1CallOptions contains the retry settings for each method of SystemPolicyV1Beta1Client.
type SystemPolicyV1Beta1CallOptions struct {
	GetSystemPolicy []gax.CallOption
}

func defaultSystemPolicyV1Beta1GRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("binaryauthorization.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("binaryauthorization.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("binaryauthorization.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://binaryauthorization.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		internaloption.EnableNewAuthLibrary(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultSystemPolicyV1Beta1CallOptions() *SystemPolicyV1Beta1CallOptions {
	return &SystemPolicyV1Beta1CallOptions{
		GetSystemPolicy: []gax.CallOption{},
	}
}

func defaultSystemPolicyV1Beta1RESTCallOptions() *SystemPolicyV1Beta1CallOptions {
	return &SystemPolicyV1Beta1CallOptions{
		GetSystemPolicy: []gax.CallOption{},
	}
}

// internalSystemPolicyV1Beta1Client is an interface that defines the methods available from Binary Authorization API.
type internalSystemPolicyV1Beta1Client interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetSystemPolicy(context.Context, *binaryauthorizationpb.GetSystemPolicyRequest, ...gax.CallOption) (*binaryauthorizationpb.Policy, error)
}

// SystemPolicyV1Beta1Client is a client for interacting with Binary Authorization API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// API for working with the system policy.
type SystemPolicyV1Beta1Client struct {
	// The internal transport-dependent client.
	internalClient internalSystemPolicyV1Beta1Client

	// The call options for this service.
	CallOptions *SystemPolicyV1Beta1CallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *SystemPolicyV1Beta1Client) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *SystemPolicyV1Beta1Client) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *SystemPolicyV1Beta1Client) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetSystemPolicy gets the current system policy in the specified location.
func (c *SystemPolicyV1Beta1Client) GetSystemPolicy(ctx context.Context, req *binaryauthorizationpb.GetSystemPolicyRequest, opts ...gax.CallOption) (*binaryauthorizationpb.Policy, error) {
	return c.internalClient.GetSystemPolicy(ctx, req, opts...)
}

// systemPolicyV1Beta1GRPCClient is a client for interacting with Binary Authorization API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type systemPolicyV1Beta1GRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing SystemPolicyV1Beta1Client
	CallOptions **SystemPolicyV1Beta1CallOptions

	// The gRPC API client.
	systemPolicyV1Beta1Client binaryauthorizationpb.SystemPolicyV1Beta1Client

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string

	logger *slog.Logger
}

// NewSystemPolicyV1Beta1Client creates a new system policy v1 beta1 client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// API for working with the system policy.
func NewSystemPolicyV1Beta1Client(ctx context.Context, opts ...option.ClientOption) (*SystemPolicyV1Beta1Client, error) {
	clientOpts := defaultSystemPolicyV1Beta1GRPCClientOptions()
	if newSystemPolicyV1Beta1ClientHook != nil {
		hookOpts, err := newSystemPolicyV1Beta1ClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := SystemPolicyV1Beta1Client{CallOptions: defaultSystemPolicyV1Beta1CallOptions()}

	c := &systemPolicyV1Beta1GRPCClient{
		connPool:                  connPool,
		systemPolicyV1Beta1Client: binaryauthorizationpb.NewSystemPolicyV1Beta1Client(connPool),
		CallOptions:               &client.CallOptions,
		logger:                    internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *systemPolicyV1Beta1GRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *systemPolicyV1Beta1GRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *systemPolicyV1Beta1GRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type systemPolicyV1Beta1RESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing SystemPolicyV1Beta1Client
	CallOptions **SystemPolicyV1Beta1CallOptions

	logger *slog.Logger
}

// NewSystemPolicyV1Beta1RESTClient creates a new system policy v1 beta1 rest client.
//
// API for working with the system policy.
func NewSystemPolicyV1Beta1RESTClient(ctx context.Context, opts ...option.ClientOption) (*SystemPolicyV1Beta1Client, error) {
	clientOpts := append(defaultSystemPolicyV1Beta1RESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultSystemPolicyV1Beta1RESTCallOptions()
	c := &systemPolicyV1Beta1RESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
		logger:      internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	return &SystemPolicyV1Beta1Client{internalClient: c, CallOptions: callOpts}, nil
}

func defaultSystemPolicyV1Beta1RESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://binaryauthorization.googleapis.com"),
		internaloption.WithDefaultEndpointTemplate("https://binaryauthorization.UNIVERSE_DOMAIN"),
		internaloption.WithDefaultMTLSEndpoint("https://binaryauthorization.mtls.googleapis.com"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://binaryauthorization.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableNewAuthLibrary(),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *systemPolicyV1Beta1RESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *systemPolicyV1Beta1RESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *systemPolicyV1Beta1RESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *systemPolicyV1Beta1GRPCClient) GetSystemPolicy(ctx context.Context, req *binaryauthorizationpb.GetSystemPolicyRequest, opts ...gax.CallOption) (*binaryauthorizationpb.Policy, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetSystemPolicy[0:len((*c.CallOptions).GetSystemPolicy):len((*c.CallOptions).GetSystemPolicy)], opts...)
	var resp *binaryauthorizationpb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.systemPolicyV1Beta1Client.GetSystemPolicy, req, settings.GRPC, c.logger, "GetSystemPolicy")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetSystemPolicy gets the current system policy in the specified location.
func (c *systemPolicyV1Beta1RESTClient) GetSystemPolicy(ctx context.Context, req *binaryauthorizationpb.GetSystemPolicyRequest, opts ...gax.CallOption) (*binaryauthorizationpb.Policy, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta1/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetSystemPolicy[0:len((*c.CallOptions).GetSystemPolicy):len((*c.CallOptions).GetSystemPolicy)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &binaryauthorizationpb.Policy{}
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

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "GetSystemPolicy")
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
