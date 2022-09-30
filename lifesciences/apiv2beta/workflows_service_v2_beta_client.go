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

package lifesciences

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"time"

	lifesciencespb "cloud.google.com/go/lifesciences/apiv2beta/lifesciencespb"
	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
)

var newWorkflowsServiceV2BetaClientHook clientHook

// WorkflowsServiceV2BetaCallOptions contains the retry settings for each method of WorkflowsServiceV2BetaClient.
type WorkflowsServiceV2BetaCallOptions struct {
	RunPipeline []gax.CallOption
}

func defaultWorkflowsServiceV2BetaGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("lifesciences.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("lifesciences.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://lifesciences.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultWorkflowsServiceV2BetaCallOptions() *WorkflowsServiceV2BetaCallOptions {
	return &WorkflowsServiceV2BetaCallOptions{
		RunPipeline: []gax.CallOption{},
	}
}

func defaultWorkflowsServiceV2BetaRESTCallOptions() *WorkflowsServiceV2BetaCallOptions {
	return &WorkflowsServiceV2BetaCallOptions{
		RunPipeline: []gax.CallOption{},
	}
}

// internalWorkflowsServiceV2BetaClient is an interface that defines the methods available from Cloud Life Sciences API.
type internalWorkflowsServiceV2BetaClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	RunPipeline(context.Context, *lifesciencespb.RunPipelineRequest, ...gax.CallOption) (*RunPipelineOperation, error)
	RunPipelineOperation(name string) *RunPipelineOperation
}

// WorkflowsServiceV2BetaClient is a client for interacting with Cloud Life Sciences API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// A service for running workflows, such as pipelines consisting of Docker
// containers.
type WorkflowsServiceV2BetaClient struct {
	// The internal transport-dependent client.
	internalClient internalWorkflowsServiceV2BetaClient

	// The call options for this service.
	CallOptions *WorkflowsServiceV2BetaCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *WorkflowsServiceV2BetaClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *WorkflowsServiceV2BetaClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *WorkflowsServiceV2BetaClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// RunPipeline runs a pipeline.  The returned Operation’s [metadata]
// [google.longrunning.Operation.metadata] field will contain a
// google.cloud.lifesciences.v2beta.Metadata object describing the status
// of the pipeline execution. The
// response field will contain a
// google.cloud.lifesciences.v2beta.RunPipelineResponse object if the
// pipeline completes successfully.
//
// Note: Before you can use this method, the Life Sciences Service Agent
// must have access to your project. This is done automatically when the
// Cloud Life Sciences API is first enabled, but if you delete this permission
// you must disable and re-enable the API to grant the Life Sciences
// Service Agent the required permissions.
// Authorization requires the following Google
// IAM (at https://cloud.google.com/iam/) permission:
//
//	lifesciences.workflows.run
func (c *WorkflowsServiceV2BetaClient) RunPipeline(ctx context.Context, req *lifesciencespb.RunPipelineRequest, opts ...gax.CallOption) (*RunPipelineOperation, error) {
	return c.internalClient.RunPipeline(ctx, req, opts...)
}

// RunPipelineOperation returns a new RunPipelineOperation from a given name.
// The name must be that of a previously created RunPipelineOperation, possibly from a different process.
func (c *WorkflowsServiceV2BetaClient) RunPipelineOperation(name string) *RunPipelineOperation {
	return c.internalClient.RunPipelineOperation(name)
}

// workflowsServiceV2BetaGRPCClient is a client for interacting with Cloud Life Sciences API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type workflowsServiceV2BetaGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing WorkflowsServiceV2BetaClient
	CallOptions **WorkflowsServiceV2BetaCallOptions

	// The gRPC API client.
	workflowsServiceV2BetaClient lifesciencespb.WorkflowsServiceV2BetaClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewWorkflowsServiceV2BetaClient creates a new workflows service v2 beta client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// A service for running workflows, such as pipelines consisting of Docker
// containers.
func NewWorkflowsServiceV2BetaClient(ctx context.Context, opts ...option.ClientOption) (*WorkflowsServiceV2BetaClient, error) {
	clientOpts := defaultWorkflowsServiceV2BetaGRPCClientOptions()
	if newWorkflowsServiceV2BetaClientHook != nil {
		hookOpts, err := newWorkflowsServiceV2BetaClientHook(ctx, clientHookParams{})
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
	client := WorkflowsServiceV2BetaClient{CallOptions: defaultWorkflowsServiceV2BetaCallOptions()}

	c := &workflowsServiceV2BetaGRPCClient{
		connPool:                     connPool,
		disableDeadlines:             disableDeadlines,
		workflowsServiceV2BetaClient: lifesciencespb.NewWorkflowsServiceV2BetaClient(connPool),
		CallOptions:                  &client.CallOptions,
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
func (c *workflowsServiceV2BetaGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *workflowsServiceV2BetaGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *workflowsServiceV2BetaGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type workflowsServiceV2BetaRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD

	// Points back to the CallOptions field of the containing WorkflowsServiceV2BetaClient
	CallOptions **WorkflowsServiceV2BetaCallOptions
}

// NewWorkflowsServiceV2BetaRESTClient creates a new workflows service v2 beta rest client.
//
// A service for running workflows, such as pipelines consisting of Docker
// containers.
func NewWorkflowsServiceV2BetaRESTClient(ctx context.Context, opts ...option.ClientOption) (*WorkflowsServiceV2BetaClient, error) {
	clientOpts := append(defaultWorkflowsServiceV2BetaRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultWorkflowsServiceV2BetaRESTCallOptions()
	c := &workflowsServiceV2BetaRESTClient{
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

	return &WorkflowsServiceV2BetaClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultWorkflowsServiceV2BetaRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://lifesciences.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://lifesciences.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://lifesciences.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *workflowsServiceV2BetaRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *workflowsServiceV2BetaRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *workflowsServiceV2BetaRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *workflowsServiceV2BetaGRPCClient) RunPipeline(ctx context.Context, req *lifesciencespb.RunPipelineRequest, opts ...gax.CallOption) (*RunPipelineOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).RunPipeline[0:len((*c.CallOptions).RunPipeline):len((*c.CallOptions).RunPipeline)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.workflowsServiceV2BetaClient.RunPipeline(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &RunPipelineOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

// RunPipeline runs a pipeline.  The returned Operation’s [metadata]
// [google.longrunning.Operation.metadata] field will contain a
// google.cloud.lifesciences.v2beta.Metadata object describing the status
// of the pipeline execution. The
// response field will contain a
// google.cloud.lifesciences.v2beta.RunPipelineResponse object if the
// pipeline completes successfully.
//
// Note: Before you can use this method, the Life Sciences Service Agent
// must have access to your project. This is done automatically when the
// Cloud Life Sciences API is first enabled, but if you delete this permission
// you must disable and re-enable the API to grant the Life Sciences
// Service Agent the required permissions.
// Authorization requires the following Google
// IAM (at https://cloud.google.com/iam/) permission:
//
//	lifesciences.workflows.run
func (c *workflowsServiceV2BetaRESTClient) RunPipeline(ctx context.Context, req *lifesciencespb.RunPipelineRequest, opts ...gax.CallOption) (*RunPipelineOperation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v2beta/%v/pipelines:run", req.GetParent())

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
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

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}

	override := fmt.Sprintf("/v2beta/%s", resp.GetName())
	return &RunPipelineOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, resp),
		pollPath: override,
	}, nil
}

// RunPipelineOperation manages a long-running operation from RunPipeline.
type RunPipelineOperation struct {
	lro      *longrunning.Operation
	pollPath string
}

// RunPipelineOperation returns a new RunPipelineOperation from a given name.
// The name must be that of a previously created RunPipelineOperation, possibly from a different process.
func (c *workflowsServiceV2BetaGRPCClient) RunPipelineOperation(name string) *RunPipelineOperation {
	return &RunPipelineOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// RunPipelineOperation returns a new RunPipelineOperation from a given name.
// The name must be that of a previously created RunPipelineOperation, possibly from a different process.
func (c *workflowsServiceV2BetaRESTClient) RunPipelineOperation(name string) *RunPipelineOperation {
	override := fmt.Sprintf("/v2beta/%s", name)
	return &RunPipelineOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
		pollPath: override,
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *RunPipelineOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*lifesciencespb.RunPipelineResponse, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp lifesciencespb.RunPipelineResponse
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
func (op *RunPipelineOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*lifesciencespb.RunPipelineResponse, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp lifesciencespb.RunPipelineResponse
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
func (op *RunPipelineOperation) Metadata() (*lifesciencespb.Metadata, error) {
	var meta lifesciencespb.Metadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *RunPipelineOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *RunPipelineOperation) Name() string {
	return op.lro.Name()
}
