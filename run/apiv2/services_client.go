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

package run

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"regexp"
	"strings"
	"time"

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	runpb "google.golang.org/genproto/googleapis/cloud/run/v2"
	iampb "google.golang.org/genproto/googleapis/iam/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newServicesClientHook clientHook

// ServicesCallOptions contains the retry settings for each method of ServicesClient.
type ServicesCallOptions struct {
	CreateService      []gax.CallOption
	GetService         []gax.CallOption
	ListServices       []gax.CallOption
	UpdateService      []gax.CallOption
	DeleteService      []gax.CallOption
	GetIamPolicy       []gax.CallOption
	SetIamPolicy       []gax.CallOption
	TestIamPermissions []gax.CallOption
}

func defaultServicesGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("run.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("run.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://run.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultServicesCallOptions() *ServicesCallOptions {
	return &ServicesCallOptions{
		CreateService: []gax.CallOption{},
		GetService: []gax.CallOption{
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
		ListServices: []gax.CallOption{
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
		UpdateService:      []gax.CallOption{},
		DeleteService:      []gax.CallOption{},
		GetIamPolicy:       []gax.CallOption{},
		SetIamPolicy:       []gax.CallOption{},
		TestIamPermissions: []gax.CallOption{},
	}
}

// internalServicesClient is an interface that defines the methods available from Cloud Run Admin API.
type internalServicesClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateService(context.Context, *runpb.CreateServiceRequest, ...gax.CallOption) (*CreateServiceOperation, error)
	CreateServiceOperation(name string) *CreateServiceOperation
	GetService(context.Context, *runpb.GetServiceRequest, ...gax.CallOption) (*runpb.Service, error)
	ListServices(context.Context, *runpb.ListServicesRequest, ...gax.CallOption) *ServiceIterator
	UpdateService(context.Context, *runpb.UpdateServiceRequest, ...gax.CallOption) (*UpdateServiceOperation, error)
	UpdateServiceOperation(name string) *UpdateServiceOperation
	DeleteService(context.Context, *runpb.DeleteServiceRequest, ...gax.CallOption) (*DeleteServiceOperation, error)
	DeleteServiceOperation(name string) *DeleteServiceOperation
	GetIamPolicy(context.Context, *iampb.GetIamPolicyRequest, ...gax.CallOption) (*iampb.Policy, error)
	SetIamPolicy(context.Context, *iampb.SetIamPolicyRequest, ...gax.CallOption) (*iampb.Policy, error)
	TestIamPermissions(context.Context, *iampb.TestIamPermissionsRequest, ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error)
}

// ServicesClient is a client for interacting with Cloud Run Admin API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Cloud Run Service Control Plane API
type ServicesClient struct {
	// The internal transport-dependent client.
	internalClient internalServicesClient

	// The call options for this service.
	CallOptions *ServicesCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ServicesClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ServicesClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *ServicesClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateService creates a new Service in a given project and location.
func (c *ServicesClient) CreateService(ctx context.Context, req *runpb.CreateServiceRequest, opts ...gax.CallOption) (*CreateServiceOperation, error) {
	return c.internalClient.CreateService(ctx, req, opts...)
}

// CreateServiceOperation returns a new CreateServiceOperation from a given name.
// The name must be that of a previously created CreateServiceOperation, possibly from a different process.
func (c *ServicesClient) CreateServiceOperation(name string) *CreateServiceOperation {
	return c.internalClient.CreateServiceOperation(name)
}

// GetService gets information about a Service.
func (c *ServicesClient) GetService(ctx context.Context, req *runpb.GetServiceRequest, opts ...gax.CallOption) (*runpb.Service, error) {
	return c.internalClient.GetService(ctx, req, opts...)
}

// ListServices list Services.
func (c *ServicesClient) ListServices(ctx context.Context, req *runpb.ListServicesRequest, opts ...gax.CallOption) *ServiceIterator {
	return c.internalClient.ListServices(ctx, req, opts...)
}

// UpdateService updates a Service.
func (c *ServicesClient) UpdateService(ctx context.Context, req *runpb.UpdateServiceRequest, opts ...gax.CallOption) (*UpdateServiceOperation, error) {
	return c.internalClient.UpdateService(ctx, req, opts...)
}

// UpdateServiceOperation returns a new UpdateServiceOperation from a given name.
// The name must be that of a previously created UpdateServiceOperation, possibly from a different process.
func (c *ServicesClient) UpdateServiceOperation(name string) *UpdateServiceOperation {
	return c.internalClient.UpdateServiceOperation(name)
}

// DeleteService deletes a Service.
// This will cause the Service to stop serving traffic and will delete all
// revisions.
func (c *ServicesClient) DeleteService(ctx context.Context, req *runpb.DeleteServiceRequest, opts ...gax.CallOption) (*DeleteServiceOperation, error) {
	return c.internalClient.DeleteService(ctx, req, opts...)
}

// DeleteServiceOperation returns a new DeleteServiceOperation from a given name.
// The name must be that of a previously created DeleteServiceOperation, possibly from a different process.
func (c *ServicesClient) DeleteServiceOperation(name string) *DeleteServiceOperation {
	return c.internalClient.DeleteServiceOperation(name)
}

// GetIamPolicy get the IAM Access Control policy currently in effect for the given
// Cloud Run Service. This result does not include any inherited policies.
func (c *ServicesClient) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	return c.internalClient.GetIamPolicy(ctx, req, opts...)
}

// SetIamPolicy sets the IAM Access control policy for the specified Service. Overwrites
// any existing policy.
func (c *ServicesClient) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	return c.internalClient.SetIamPolicy(ctx, req, opts...)
}

// TestIamPermissions returns permissions that a caller has on the specified Project.
//
// There are no permissions required for making this API call.
func (c *ServicesClient) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest, opts ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error) {
	return c.internalClient.TestIamPermissions(ctx, req, opts...)
}

// servicesGRPCClient is a client for interacting with Cloud Run Admin API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type servicesGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing ServicesClient
	CallOptions **ServicesCallOptions

	// The gRPC API client.
	servicesClient runpb.ServicesClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewServicesClient creates a new services client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Cloud Run Service Control Plane API
func NewServicesClient(ctx context.Context, opts ...option.ClientOption) (*ServicesClient, error) {
	clientOpts := defaultServicesGRPCClientOptions()
	if newServicesClientHook != nil {
		hookOpts, err := newServicesClientHook(ctx, clientHookParams{})
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
	client := ServicesClient{CallOptions: defaultServicesCallOptions()}

	c := &servicesGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		servicesClient:   runpb.NewServicesClient(connPool),
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
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *servicesGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *servicesGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *servicesGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *servicesGRPCClient) CreateService(ctx context.Context, req *runpb.CreateServiceRequest, opts ...gax.CallOption) (*CreateServiceOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 15000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	routingHeaders := ""
	routingHeadersMap := make(map[string]string)
	if reg := regexp.MustCompile("projects/[^/]+/locations/(?P<location>[^/]+)"); reg.MatchString(req.GetParent()) && len(url.QueryEscape(reg.FindStringSubmatch(req.GetParent())[1])) > 0 {
		routingHeadersMap["location"] = url.QueryEscape(reg.FindStringSubmatch(req.GetParent())[1])
	}
	for headerName, headerValue := range routingHeadersMap {
		routingHeaders = fmt.Sprintf("%s%s=%s&", routingHeaders, headerName, headerValue)
	}
	routingHeaders = strings.TrimSuffix(routingHeaders, "&")
	md := metadata.Pairs("x-goog-request-params", routingHeaders)

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateService[0:len((*c.CallOptions).CreateService):len((*c.CallOptions).CreateService)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.servicesClient.CreateService(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &CreateServiceOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *servicesGRPCClient) GetService(ctx context.Context, req *runpb.GetServiceRequest, opts ...gax.CallOption) (*runpb.Service, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 10000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	routingHeaders := ""
	routingHeadersMap := make(map[string]string)
	if reg := regexp.MustCompile("projects/[^/]+/locations/(?P<location>[^/]+)(?:/.*)?"); reg.MatchString(req.GetName()) && len(url.QueryEscape(reg.FindStringSubmatch(req.GetName())[1])) > 0 {
		routingHeadersMap["location"] = url.QueryEscape(reg.FindStringSubmatch(req.GetName())[1])
	}
	for headerName, headerValue := range routingHeadersMap {
		routingHeaders = fmt.Sprintf("%s%s=%s&", routingHeaders, headerName, headerValue)
	}
	routingHeaders = strings.TrimSuffix(routingHeaders, "&")
	md := metadata.Pairs("x-goog-request-params", routingHeaders)

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetService[0:len((*c.CallOptions).GetService):len((*c.CallOptions).GetService)], opts...)
	var resp *runpb.Service
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.servicesClient.GetService(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *servicesGRPCClient) ListServices(ctx context.Context, req *runpb.ListServicesRequest, opts ...gax.CallOption) *ServiceIterator {
	routingHeaders := ""
	routingHeadersMap := make(map[string]string)
	if reg := regexp.MustCompile("projects/[^/]+/locations/(?P<location>[^/]+)"); reg.MatchString(req.GetParent()) && len(url.QueryEscape(reg.FindStringSubmatch(req.GetParent())[1])) > 0 {
		routingHeadersMap["location"] = url.QueryEscape(reg.FindStringSubmatch(req.GetParent())[1])
	}
	for headerName, headerValue := range routingHeadersMap {
		routingHeaders = fmt.Sprintf("%s%s=%s&", routingHeaders, headerName, headerValue)
	}
	routingHeaders = strings.TrimSuffix(routingHeaders, "&")
	md := metadata.Pairs("x-goog-request-params", routingHeaders)

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListServices[0:len((*c.CallOptions).ListServices):len((*c.CallOptions).ListServices)], opts...)
	it := &ServiceIterator{}
	req = proto.Clone(req).(*runpb.ListServicesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*runpb.Service, string, error) {
		resp := &runpb.ListServicesResponse{}
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
			resp, err = c.servicesClient.ListServices(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetServices(), resp.GetNextPageToken(), nil
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

func (c *servicesGRPCClient) UpdateService(ctx context.Context, req *runpb.UpdateServiceRequest, opts ...gax.CallOption) (*UpdateServiceOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 15000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	routingHeaders := ""
	routingHeadersMap := make(map[string]string)
	if reg := regexp.MustCompile("projects/[^/]+/locations/(?P<location>[^/]+)(?:/.*)?"); reg.MatchString(req.GetService().GetName()) && len(url.QueryEscape(reg.FindStringSubmatch(req.GetService().GetName())[1])) > 0 {
		routingHeadersMap["location"] = url.QueryEscape(reg.FindStringSubmatch(req.GetService().GetName())[1])
	}
	for headerName, headerValue := range routingHeadersMap {
		routingHeaders = fmt.Sprintf("%s%s=%s&", routingHeaders, headerName, headerValue)
	}
	routingHeaders = strings.TrimSuffix(routingHeaders, "&")
	md := metadata.Pairs("x-goog-request-params", routingHeaders)

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateService[0:len((*c.CallOptions).UpdateService):len((*c.CallOptions).UpdateService)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.servicesClient.UpdateService(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &UpdateServiceOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *servicesGRPCClient) DeleteService(ctx context.Context, req *runpb.DeleteServiceRequest, opts ...gax.CallOption) (*DeleteServiceOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 10000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	routingHeaders := ""
	routingHeadersMap := make(map[string]string)
	if reg := regexp.MustCompile("projects/[^/]+/locations/(?P<location>[^/]+)(?:/.*)?"); reg.MatchString(req.GetName()) && len(url.QueryEscape(reg.FindStringSubmatch(req.GetName())[1])) > 0 {
		routingHeadersMap["location"] = url.QueryEscape(reg.FindStringSubmatch(req.GetName())[1])
	}
	for headerName, headerValue := range routingHeadersMap {
		routingHeaders = fmt.Sprintf("%s%s=%s&", routingHeaders, headerName, headerValue)
	}
	routingHeaders = strings.TrimSuffix(routingHeaders, "&")
	md := metadata.Pairs("x-goog-request-params", routingHeaders)

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteService[0:len((*c.CallOptions).DeleteService):len((*c.CallOptions).DeleteService)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.servicesClient.DeleteService(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &DeleteServiceOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *servicesGRPCClient) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetIamPolicy[0:len((*c.CallOptions).GetIamPolicy):len((*c.CallOptions).GetIamPolicy)], opts...)
	var resp *iampb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.servicesClient.GetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *servicesGRPCClient) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).SetIamPolicy[0:len((*c.CallOptions).SetIamPolicy):len((*c.CallOptions).SetIamPolicy)], opts...)
	var resp *iampb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.servicesClient.SetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *servicesGRPCClient) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest, opts ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).TestIamPermissions[0:len((*c.CallOptions).TestIamPermissions):len((*c.CallOptions).TestIamPermissions)], opts...)
	var resp *iampb.TestIamPermissionsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.servicesClient.TestIamPermissions(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateServiceOperation manages a long-running operation from CreateService.
type CreateServiceOperation struct {
	lro *longrunning.Operation
}

// CreateServiceOperation returns a new CreateServiceOperation from a given name.
// The name must be that of a previously created CreateServiceOperation, possibly from a different process.
func (c *servicesGRPCClient) CreateServiceOperation(name string) *CreateServiceOperation {
	return &CreateServiceOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *CreateServiceOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*runpb.Service, error) {
	var resp runpb.Service
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
func (op *CreateServiceOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*runpb.Service, error) {
	var resp runpb.Service
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
func (op *CreateServiceOperation) Metadata() (*runpb.Service, error) {
	var meta runpb.Service
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *CreateServiceOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *CreateServiceOperation) Name() string {
	return op.lro.Name()
}

// DeleteServiceOperation manages a long-running operation from DeleteService.
type DeleteServiceOperation struct {
	lro *longrunning.Operation
}

// DeleteServiceOperation returns a new DeleteServiceOperation from a given name.
// The name must be that of a previously created DeleteServiceOperation, possibly from a different process.
func (c *servicesGRPCClient) DeleteServiceOperation(name string) *DeleteServiceOperation {
	return &DeleteServiceOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *DeleteServiceOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*runpb.Service, error) {
	var resp runpb.Service
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
func (op *DeleteServiceOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*runpb.Service, error) {
	var resp runpb.Service
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
func (op *DeleteServiceOperation) Metadata() (*runpb.Service, error) {
	var meta runpb.Service
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *DeleteServiceOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *DeleteServiceOperation) Name() string {
	return op.lro.Name()
}

// UpdateServiceOperation manages a long-running operation from UpdateService.
type UpdateServiceOperation struct {
	lro *longrunning.Operation
}

// UpdateServiceOperation returns a new UpdateServiceOperation from a given name.
// The name must be that of a previously created UpdateServiceOperation, possibly from a different process.
func (c *servicesGRPCClient) UpdateServiceOperation(name string) *UpdateServiceOperation {
	return &UpdateServiceOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *UpdateServiceOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*runpb.Service, error) {
	var resp runpb.Service
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
func (op *UpdateServiceOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*runpb.Service, error) {
	var resp runpb.Service
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
func (op *UpdateServiceOperation) Metadata() (*runpb.Service, error) {
	var meta runpb.Service
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *UpdateServiceOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *UpdateServiceOperation) Name() string {
	return op.lro.Name()
}

// ServiceIterator manages a stream of *runpb.Service.
type ServiceIterator struct {
	items    []*runpb.Service
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
	InternalFetch func(pageSize int, pageToken string) (results []*runpb.Service, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *ServiceIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ServiceIterator) Next() (*runpb.Service, error) {
	var item *runpb.Service
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ServiceIterator) bufLen() int {
	return len(it.items)
}

func (it *ServiceIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
