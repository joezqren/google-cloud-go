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

package aiplatform

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
	aiplatformpb "google.golang.org/genproto/googleapis/cloud/aiplatform/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newModelClientHook clientHook

// ModelCallOptions contains the retry settings for each method of ModelClient.
type ModelCallOptions struct {
	UploadModel               []gax.CallOption
	GetModel                  []gax.CallOption
	ListModels                []gax.CallOption
	UpdateModel               []gax.CallOption
	DeleteModel               []gax.CallOption
	ExportModel               []gax.CallOption
	GetModelEvaluation        []gax.CallOption
	ListModelEvaluations      []gax.CallOption
	GetModelEvaluationSlice   []gax.CallOption
	ListModelEvaluationSlices []gax.CallOption
}

func defaultModelGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("aiplatform.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("aiplatform.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://aiplatform.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultModelCallOptions() *ModelCallOptions {
	return &ModelCallOptions{
		UploadModel:               []gax.CallOption{},
		GetModel:                  []gax.CallOption{},
		ListModels:                []gax.CallOption{},
		UpdateModel:               []gax.CallOption{},
		DeleteModel:               []gax.CallOption{},
		ExportModel:               []gax.CallOption{},
		GetModelEvaluation:        []gax.CallOption{},
		ListModelEvaluations:      []gax.CallOption{},
		GetModelEvaluationSlice:   []gax.CallOption{},
		ListModelEvaluationSlices: []gax.CallOption{},
	}
}

// internalModelClient is an interface that defines the methods availaible from Vertex AI API.
type internalModelClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	UploadModel(context.Context, *aiplatformpb.UploadModelRequest, ...gax.CallOption) (*UploadModelOperation, error)
	UploadModelOperation(name string) *UploadModelOperation
	GetModel(context.Context, *aiplatformpb.GetModelRequest, ...gax.CallOption) (*aiplatformpb.Model, error)
	ListModels(context.Context, *aiplatformpb.ListModelsRequest, ...gax.CallOption) *ModelIterator
	UpdateModel(context.Context, *aiplatformpb.UpdateModelRequest, ...gax.CallOption) (*aiplatformpb.Model, error)
	DeleteModel(context.Context, *aiplatformpb.DeleteModelRequest, ...gax.CallOption) (*DeleteModelOperation, error)
	DeleteModelOperation(name string) *DeleteModelOperation
	ExportModel(context.Context, *aiplatformpb.ExportModelRequest, ...gax.CallOption) (*ExportModelOperation, error)
	ExportModelOperation(name string) *ExportModelOperation
	GetModelEvaluation(context.Context, *aiplatformpb.GetModelEvaluationRequest, ...gax.CallOption) (*aiplatformpb.ModelEvaluation, error)
	ListModelEvaluations(context.Context, *aiplatformpb.ListModelEvaluationsRequest, ...gax.CallOption) *ModelEvaluationIterator
	GetModelEvaluationSlice(context.Context, *aiplatformpb.GetModelEvaluationSliceRequest, ...gax.CallOption) (*aiplatformpb.ModelEvaluationSlice, error)
	ListModelEvaluationSlices(context.Context, *aiplatformpb.ListModelEvaluationSlicesRequest, ...gax.CallOption) *ModelEvaluationSliceIterator
}

// ModelClient is a client for interacting with Vertex AI API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// A service for managing Vertex AI’s machine learning Models.
type ModelClient struct {
	// The internal transport-dependent client.
	internalClient internalModelClient

	// The call options for this service.
	CallOptions *ModelCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ModelClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ModelClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *ModelClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// UploadModel uploads a Model artifact into Vertex AI.
func (c *ModelClient) UploadModel(ctx context.Context, req *aiplatformpb.UploadModelRequest, opts ...gax.CallOption) (*UploadModelOperation, error) {
	return c.internalClient.UploadModel(ctx, req, opts...)
}

// UploadModelOperation returns a new UploadModelOperation from a given name.
// The name must be that of a previously created UploadModelOperation, possibly from a different process.
func (c *ModelClient) UploadModelOperation(name string) *UploadModelOperation {
	return c.internalClient.UploadModelOperation(name)
}

// GetModel gets a Model.
func (c *ModelClient) GetModel(ctx context.Context, req *aiplatformpb.GetModelRequest, opts ...gax.CallOption) (*aiplatformpb.Model, error) {
	return c.internalClient.GetModel(ctx, req, opts...)
}

// ListModels lists Models in a Location.
func (c *ModelClient) ListModels(ctx context.Context, req *aiplatformpb.ListModelsRequest, opts ...gax.CallOption) *ModelIterator {
	return c.internalClient.ListModels(ctx, req, opts...)
}

// UpdateModel updates a Model.
func (c *ModelClient) UpdateModel(ctx context.Context, req *aiplatformpb.UpdateModelRequest, opts ...gax.CallOption) (*aiplatformpb.Model, error) {
	return c.internalClient.UpdateModel(ctx, req, opts...)
}

// DeleteModel deletes a Model.
//
// A model cannot be deleted if any Endpoint resource has a
// DeployedModel based on the model in its
// deployed_models field.
func (c *ModelClient) DeleteModel(ctx context.Context, req *aiplatformpb.DeleteModelRequest, opts ...gax.CallOption) (*DeleteModelOperation, error) {
	return c.internalClient.DeleteModel(ctx, req, opts...)
}

// DeleteModelOperation returns a new DeleteModelOperation from a given name.
// The name must be that of a previously created DeleteModelOperation, possibly from a different process.
func (c *ModelClient) DeleteModelOperation(name string) *DeleteModelOperation {
	return c.internalClient.DeleteModelOperation(name)
}

// ExportModel exports a trained, exportable Model to a location specified by the
// user. A Model is considered to be exportable if it has at least one
// [supported export format][google.cloud.aiplatform.v1.Model.supported_export_formats].
func (c *ModelClient) ExportModel(ctx context.Context, req *aiplatformpb.ExportModelRequest, opts ...gax.CallOption) (*ExportModelOperation, error) {
	return c.internalClient.ExportModel(ctx, req, opts...)
}

// ExportModelOperation returns a new ExportModelOperation from a given name.
// The name must be that of a previously created ExportModelOperation, possibly from a different process.
func (c *ModelClient) ExportModelOperation(name string) *ExportModelOperation {
	return c.internalClient.ExportModelOperation(name)
}

// GetModelEvaluation gets a ModelEvaluation.
func (c *ModelClient) GetModelEvaluation(ctx context.Context, req *aiplatformpb.GetModelEvaluationRequest, opts ...gax.CallOption) (*aiplatformpb.ModelEvaluation, error) {
	return c.internalClient.GetModelEvaluation(ctx, req, opts...)
}

// ListModelEvaluations lists ModelEvaluations in a Model.
func (c *ModelClient) ListModelEvaluations(ctx context.Context, req *aiplatformpb.ListModelEvaluationsRequest, opts ...gax.CallOption) *ModelEvaluationIterator {
	return c.internalClient.ListModelEvaluations(ctx, req, opts...)
}

// GetModelEvaluationSlice gets a ModelEvaluationSlice.
func (c *ModelClient) GetModelEvaluationSlice(ctx context.Context, req *aiplatformpb.GetModelEvaluationSliceRequest, opts ...gax.CallOption) (*aiplatformpb.ModelEvaluationSlice, error) {
	return c.internalClient.GetModelEvaluationSlice(ctx, req, opts...)
}

// ListModelEvaluationSlices lists ModelEvaluationSlices in a ModelEvaluation.
func (c *ModelClient) ListModelEvaluationSlices(ctx context.Context, req *aiplatformpb.ListModelEvaluationSlicesRequest, opts ...gax.CallOption) *ModelEvaluationSliceIterator {
	return c.internalClient.ListModelEvaluationSlices(ctx, req, opts...)
}

// modelGRPCClient is a client for interacting with Vertex AI API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type modelGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing ModelClient
	CallOptions **ModelCallOptions

	// The gRPC API client.
	modelClient aiplatformpb.ModelServiceClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewModelClient creates a new model service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// A service for managing Vertex AI’s machine learning Models.
func NewModelClient(ctx context.Context, opts ...option.ClientOption) (*ModelClient, error) {
	clientOpts := defaultModelGRPCClientOptions()
	if newModelClientHook != nil {
		hookOpts, err := newModelClientHook(ctx, clientHookParams{})
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
	client := ModelClient{CallOptions: defaultModelCallOptions()}

	c := &modelGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		modelClient:      aiplatformpb.NewModelServiceClient(connPool),
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
func (c *modelGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *modelGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *modelGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *modelGRPCClient) UploadModel(ctx context.Context, req *aiplatformpb.UploadModelRequest, opts ...gax.CallOption) (*UploadModelOperation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UploadModel[0:len((*c.CallOptions).UploadModel):len((*c.CallOptions).UploadModel)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.modelClient.UploadModel(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &UploadModelOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *modelGRPCClient) GetModel(ctx context.Context, req *aiplatformpb.GetModelRequest, opts ...gax.CallOption) (*aiplatformpb.Model, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetModel[0:len((*c.CallOptions).GetModel):len((*c.CallOptions).GetModel)], opts...)
	var resp *aiplatformpb.Model
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.modelClient.GetModel(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *modelGRPCClient) ListModels(ctx context.Context, req *aiplatformpb.ListModelsRequest, opts ...gax.CallOption) *ModelIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListModels[0:len((*c.CallOptions).ListModels):len((*c.CallOptions).ListModels)], opts...)
	it := &ModelIterator{}
	req = proto.Clone(req).(*aiplatformpb.ListModelsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*aiplatformpb.Model, string, error) {
		resp := &aiplatformpb.ListModelsResponse{}
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
			resp, err = c.modelClient.ListModels(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetModels(), resp.GetNextPageToken(), nil
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

func (c *modelGRPCClient) UpdateModel(ctx context.Context, req *aiplatformpb.UpdateModelRequest, opts ...gax.CallOption) (*aiplatformpb.Model, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "model.name", url.QueryEscape(req.GetModel().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateModel[0:len((*c.CallOptions).UpdateModel):len((*c.CallOptions).UpdateModel)], opts...)
	var resp *aiplatformpb.Model
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.modelClient.UpdateModel(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *modelGRPCClient) DeleteModel(ctx context.Context, req *aiplatformpb.DeleteModelRequest, opts ...gax.CallOption) (*DeleteModelOperation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteModel[0:len((*c.CallOptions).DeleteModel):len((*c.CallOptions).DeleteModel)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.modelClient.DeleteModel(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &DeleteModelOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *modelGRPCClient) ExportModel(ctx context.Context, req *aiplatformpb.ExportModelRequest, opts ...gax.CallOption) (*ExportModelOperation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ExportModel[0:len((*c.CallOptions).ExportModel):len((*c.CallOptions).ExportModel)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.modelClient.ExportModel(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &ExportModelOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *modelGRPCClient) GetModelEvaluation(ctx context.Context, req *aiplatformpb.GetModelEvaluationRequest, opts ...gax.CallOption) (*aiplatformpb.ModelEvaluation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetModelEvaluation[0:len((*c.CallOptions).GetModelEvaluation):len((*c.CallOptions).GetModelEvaluation)], opts...)
	var resp *aiplatformpb.ModelEvaluation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.modelClient.GetModelEvaluation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *modelGRPCClient) ListModelEvaluations(ctx context.Context, req *aiplatformpb.ListModelEvaluationsRequest, opts ...gax.CallOption) *ModelEvaluationIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListModelEvaluations[0:len((*c.CallOptions).ListModelEvaluations):len((*c.CallOptions).ListModelEvaluations)], opts...)
	it := &ModelEvaluationIterator{}
	req = proto.Clone(req).(*aiplatformpb.ListModelEvaluationsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*aiplatformpb.ModelEvaluation, string, error) {
		resp := &aiplatformpb.ListModelEvaluationsResponse{}
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
			resp, err = c.modelClient.ListModelEvaluations(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetModelEvaluations(), resp.GetNextPageToken(), nil
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

func (c *modelGRPCClient) GetModelEvaluationSlice(ctx context.Context, req *aiplatformpb.GetModelEvaluationSliceRequest, opts ...gax.CallOption) (*aiplatformpb.ModelEvaluationSlice, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetModelEvaluationSlice[0:len((*c.CallOptions).GetModelEvaluationSlice):len((*c.CallOptions).GetModelEvaluationSlice)], opts...)
	var resp *aiplatformpb.ModelEvaluationSlice
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.modelClient.GetModelEvaluationSlice(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *modelGRPCClient) ListModelEvaluationSlices(ctx context.Context, req *aiplatformpb.ListModelEvaluationSlicesRequest, opts ...gax.CallOption) *ModelEvaluationSliceIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListModelEvaluationSlices[0:len((*c.CallOptions).ListModelEvaluationSlices):len((*c.CallOptions).ListModelEvaluationSlices)], opts...)
	it := &ModelEvaluationSliceIterator{}
	req = proto.Clone(req).(*aiplatformpb.ListModelEvaluationSlicesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*aiplatformpb.ModelEvaluationSlice, string, error) {
		resp := &aiplatformpb.ListModelEvaluationSlicesResponse{}
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
			resp, err = c.modelClient.ListModelEvaluationSlices(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetModelEvaluationSlices(), resp.GetNextPageToken(), nil
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

// DeleteModelOperation manages a long-running operation from DeleteModel.
type DeleteModelOperation struct {
	lro *longrunning.Operation
}

// DeleteModelOperation returns a new DeleteModelOperation from a given name.
// The name must be that of a previously created DeleteModelOperation, possibly from a different process.
func (c *modelGRPCClient) DeleteModelOperation(name string) *DeleteModelOperation {
	return &DeleteModelOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *DeleteModelOperation) Wait(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.WaitWithInterval(ctx, nil, time.Minute, opts...)
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
func (op *DeleteModelOperation) Poll(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.Poll(ctx, nil, opts...)
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *DeleteModelOperation) Metadata() (*aiplatformpb.DeleteOperationMetadata, error) {
	var meta aiplatformpb.DeleteOperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *DeleteModelOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *DeleteModelOperation) Name() string {
	return op.lro.Name()
}

// ExportModelOperation manages a long-running operation from ExportModel.
type ExportModelOperation struct {
	lro *longrunning.Operation
}

// ExportModelOperation returns a new ExportModelOperation from a given name.
// The name must be that of a previously created ExportModelOperation, possibly from a different process.
func (c *modelGRPCClient) ExportModelOperation(name string) *ExportModelOperation {
	return &ExportModelOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *ExportModelOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*aiplatformpb.ExportModelResponse, error) {
	var resp aiplatformpb.ExportModelResponse
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
func (op *ExportModelOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*aiplatformpb.ExportModelResponse, error) {
	var resp aiplatformpb.ExportModelResponse
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
func (op *ExportModelOperation) Metadata() (*aiplatformpb.ExportModelOperationMetadata, error) {
	var meta aiplatformpb.ExportModelOperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *ExportModelOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *ExportModelOperation) Name() string {
	return op.lro.Name()
}

// UploadModelOperation manages a long-running operation from UploadModel.
type UploadModelOperation struct {
	lro *longrunning.Operation
}

// UploadModelOperation returns a new UploadModelOperation from a given name.
// The name must be that of a previously created UploadModelOperation, possibly from a different process.
func (c *modelGRPCClient) UploadModelOperation(name string) *UploadModelOperation {
	return &UploadModelOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *UploadModelOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*aiplatformpb.UploadModelResponse, error) {
	var resp aiplatformpb.UploadModelResponse
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
func (op *UploadModelOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*aiplatformpb.UploadModelResponse, error) {
	var resp aiplatformpb.UploadModelResponse
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
func (op *UploadModelOperation) Metadata() (*aiplatformpb.UploadModelOperationMetadata, error) {
	var meta aiplatformpb.UploadModelOperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *UploadModelOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *UploadModelOperation) Name() string {
	return op.lro.Name()
}

// ModelEvaluationIterator manages a stream of *aiplatformpb.ModelEvaluation.
type ModelEvaluationIterator struct {
	items    []*aiplatformpb.ModelEvaluation
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
	InternalFetch func(pageSize int, pageToken string) (results []*aiplatformpb.ModelEvaluation, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *ModelEvaluationIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ModelEvaluationIterator) Next() (*aiplatformpb.ModelEvaluation, error) {
	var item *aiplatformpb.ModelEvaluation
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ModelEvaluationIterator) bufLen() int {
	return len(it.items)
}

func (it *ModelEvaluationIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// ModelEvaluationSliceIterator manages a stream of *aiplatformpb.ModelEvaluationSlice.
type ModelEvaluationSliceIterator struct {
	items    []*aiplatformpb.ModelEvaluationSlice
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
	InternalFetch func(pageSize int, pageToken string) (results []*aiplatformpb.ModelEvaluationSlice, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *ModelEvaluationSliceIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ModelEvaluationSliceIterator) Next() (*aiplatformpb.ModelEvaluationSlice, error) {
	var item *aiplatformpb.ModelEvaluationSlice
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ModelEvaluationSliceIterator) bufLen() int {
	return len(it.items)
}

func (it *ModelEvaluationSliceIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// ModelIterator manages a stream of *aiplatformpb.Model.
type ModelIterator struct {
	items    []*aiplatformpb.Model
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
	InternalFetch func(pageSize int, pageToken string) (results []*aiplatformpb.Model, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *ModelIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ModelIterator) Next() (*aiplatformpb.Model, error) {
	var item *aiplatformpb.Model
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ModelIterator) bufLen() int {
	return len(it.items)
}

func (it *ModelIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
