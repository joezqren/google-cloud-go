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

package generativelanguage

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"math"
	"net/http"
	"net/url"
	"time"

	generativelanguagepb "cloud.google.com/go/ai/generativelanguage/apiv1beta2/generativelanguagepb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/encoding/protojson"
)

var newTextClientHook clientHook

// TextCallOptions contains the retry settings for each method of TextClient.
type TextCallOptions struct {
	GenerateText []gax.CallOption
	EmbedText    []gax.CallOption
}

func defaultTextGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("generativelanguage.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("generativelanguage.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("generativelanguage.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://generativelanguage.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		internaloption.EnableNewAuthLibrary(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultTextCallOptions() *TextCallOptions {
	return &TextCallOptions{
		GenerateText: []gax.CallOption{
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
		EmbedText: []gax.CallOption{
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

func defaultTextRESTCallOptions() *TextCallOptions {
	return &TextCallOptions{
		GenerateText: []gax.CallOption{
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
		EmbedText: []gax.CallOption{
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

// internalTextClient is an interface that defines the methods available from Generative Language API.
type internalTextClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GenerateText(context.Context, *generativelanguagepb.GenerateTextRequest, ...gax.CallOption) (*generativelanguagepb.GenerateTextResponse, error)
	EmbedText(context.Context, *generativelanguagepb.EmbedTextRequest, ...gax.CallOption) (*generativelanguagepb.EmbedTextResponse, error)
}

// TextClient is a client for interacting with Generative Language API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// API for using Generative Language Models (GLMs) trained to generate text.
//
// Also known as Large Language Models (LLM)s, these generate text given an
// input prompt from the user.
type TextClient struct {
	// The internal transport-dependent client.
	internalClient internalTextClient

	// The call options for this service.
	CallOptions *TextCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *TextClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *TextClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *TextClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GenerateText generates a response from the model given an input message.
func (c *TextClient) GenerateText(ctx context.Context, req *generativelanguagepb.GenerateTextRequest, opts ...gax.CallOption) (*generativelanguagepb.GenerateTextResponse, error) {
	return c.internalClient.GenerateText(ctx, req, opts...)
}

// EmbedText generates an embedding from the model given an input message.
func (c *TextClient) EmbedText(ctx context.Context, req *generativelanguagepb.EmbedTextRequest, opts ...gax.CallOption) (*generativelanguagepb.EmbedTextResponse, error) {
	return c.internalClient.EmbedText(ctx, req, opts...)
}

// textGRPCClient is a client for interacting with Generative Language API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type textGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing TextClient
	CallOptions **TextCallOptions

	// The gRPC API client.
	textClient generativelanguagepb.TextServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string

	logger *slog.Logger
}

// NewTextClient creates a new text service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// API for using Generative Language Models (GLMs) trained to generate text.
//
// Also known as Large Language Models (LLM)s, these generate text given an
// input prompt from the user.
func NewTextClient(ctx context.Context, opts ...option.ClientOption) (*TextClient, error) {
	clientOpts := defaultTextGRPCClientOptions()
	if newTextClientHook != nil {
		hookOpts, err := newTextClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := TextClient{CallOptions: defaultTextCallOptions()}

	c := &textGRPCClient{
		connPool:    connPool,
		textClient:  generativelanguagepb.NewTextServiceClient(connPool),
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
func (c *textGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *textGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *textGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type textRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing TextClient
	CallOptions **TextCallOptions

	logger *slog.Logger
}

// NewTextRESTClient creates a new text service rest client.
//
// API for using Generative Language Models (GLMs) trained to generate text.
//
// Also known as Large Language Models (LLM)s, these generate text given an
// input prompt from the user.
func NewTextRESTClient(ctx context.Context, opts ...option.ClientOption) (*TextClient, error) {
	clientOpts := append(defaultTextRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultTextRESTCallOptions()
	c := &textRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
		logger:      internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	return &TextClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultTextRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://generativelanguage.googleapis.com"),
		internaloption.WithDefaultEndpointTemplate("https://generativelanguage.UNIVERSE_DOMAIN"),
		internaloption.WithDefaultMTLSEndpoint("https://generativelanguage.mtls.googleapis.com"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://generativelanguage.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableNewAuthLibrary(),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *textRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *textRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *textRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *textGRPCClient) GenerateText(ctx context.Context, req *generativelanguagepb.GenerateTextRequest, opts ...gax.CallOption) (*generativelanguagepb.GenerateTextResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "model", url.QueryEscape(req.GetModel()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GenerateText[0:len((*c.CallOptions).GenerateText):len((*c.CallOptions).GenerateText)], opts...)
	var resp *generativelanguagepb.GenerateTextResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.textClient.GenerateText, req, settings.GRPC, c.logger, "GenerateText")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *textGRPCClient) EmbedText(ctx context.Context, req *generativelanguagepb.EmbedTextRequest, opts ...gax.CallOption) (*generativelanguagepb.EmbedTextResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "model", url.QueryEscape(req.GetModel()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).EmbedText[0:len((*c.CallOptions).EmbedText):len((*c.CallOptions).EmbedText)], opts...)
	var resp *generativelanguagepb.EmbedTextResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.textClient.EmbedText, req, settings.GRPC, c.logger, "EmbedText")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GenerateText generates a response from the model given an input message.
func (c *textRESTClient) GenerateText(ctx context.Context, req *generativelanguagepb.GenerateTextRequest, opts ...gax.CallOption) (*generativelanguagepb.GenerateTextResponse, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta2/%v:generateText", req.GetModel())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "model", url.QueryEscape(req.GetModel()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GenerateText[0:len((*c.CallOptions).GenerateText):len((*c.CallOptions).GenerateText)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &generativelanguagepb.GenerateTextResponse{}
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

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "GenerateText")
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

// EmbedText generates an embedding from the model given an input message.
func (c *textRESTClient) EmbedText(ctx context.Context, req *generativelanguagepb.EmbedTextRequest, opts ...gax.CallOption) (*generativelanguagepb.EmbedTextResponse, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta2/%v:embedText", req.GetModel())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "model", url.QueryEscape(req.GetModel()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).EmbedText[0:len((*c.CallOptions).EmbedText):len((*c.CallOptions).EmbedText)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &generativelanguagepb.EmbedTextResponse{}
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

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "EmbedText")
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
