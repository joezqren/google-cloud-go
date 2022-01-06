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

package compute

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	httptransport "google.golang.org/api/transport/http"
	computepb "google.golang.org/genproto/googleapis/cloud/compute/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newLicensesClientHook clientHook

// LicensesCallOptions contains the retry settings for each method of LicensesClient.
type LicensesCallOptions struct {
	Delete             []gax.CallOption
	Get                []gax.CallOption
	GetIamPolicy       []gax.CallOption
	Insert             []gax.CallOption
	List               []gax.CallOption
	SetIamPolicy       []gax.CallOption
	TestIamPermissions []gax.CallOption
}

// internalLicensesClient is an interface that defines the methods availaible from Google Compute Engine API.
type internalLicensesClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	Delete(context.Context, *computepb.DeleteLicenseRequest, ...gax.CallOption) (*Operation, error)
	Get(context.Context, *computepb.GetLicenseRequest, ...gax.CallOption) (*computepb.License, error)
	GetIamPolicy(context.Context, *computepb.GetIamPolicyLicenseRequest, ...gax.CallOption) (*computepb.Policy, error)
	Insert(context.Context, *computepb.InsertLicenseRequest, ...gax.CallOption) (*Operation, error)
	List(context.Context, *computepb.ListLicensesRequest, ...gax.CallOption) *LicenseIterator
	SetIamPolicy(context.Context, *computepb.SetIamPolicyLicenseRequest, ...gax.CallOption) (*computepb.Policy, error)
	TestIamPermissions(context.Context, *computepb.TestIamPermissionsLicenseRequest, ...gax.CallOption) (*computepb.TestPermissionsResponse, error)
}

// LicensesClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The Licenses API.
type LicensesClient struct {
	// The internal transport-dependent client.
	internalClient internalLicensesClient

	// The call options for this service.
	CallOptions *LicensesCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *LicensesClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *LicensesClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *LicensesClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// Delete deletes the specified license. Caution This resource is intended for use only by third-party partners who are creating Cloud Marketplace images.
func (c *LicensesClient) Delete(ctx context.Context, req *computepb.DeleteLicenseRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// Get returns the specified License resource. Caution This resource is intended for use only by third-party partners who are creating Cloud Marketplace images.
func (c *LicensesClient) Get(ctx context.Context, req *computepb.GetLicenseRequest, opts ...gax.CallOption) (*computepb.License, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// GetIamPolicy gets the access control policy for a resource. May be empty if no such policy or resource exists. Caution This resource is intended for use only by third-party partners who are creating Cloud Marketplace images.
func (c *LicensesClient) GetIamPolicy(ctx context.Context, req *computepb.GetIamPolicyLicenseRequest, opts ...gax.CallOption) (*computepb.Policy, error) {
	return c.internalClient.GetIamPolicy(ctx, req, opts...)
}

// Insert create a License resource in the specified project. Caution This resource is intended for use only by third-party partners who are creating Cloud Marketplace images.
func (c *LicensesClient) Insert(ctx context.Context, req *computepb.InsertLicenseRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List retrieves the list of licenses available in the specified project. This method does not get any licenses that belong to other projects, including licenses attached to publicly-available images, like Debian 9. If you want to get a list of publicly-available licenses, use this method to make a request to the respective image project, such as debian-cloud or windows-cloud. Caution This resource is intended for use only by third-party partners who are creating Cloud Marketplace images.
func (c *LicensesClient) List(ctx context.Context, req *computepb.ListLicensesRequest, opts ...gax.CallOption) *LicenseIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// SetIamPolicy sets the access control policy on the specified resource. Replaces any existing policy. Caution This resource is intended for use only by third-party partners who are creating Cloud Marketplace images.
func (c *LicensesClient) SetIamPolicy(ctx context.Context, req *computepb.SetIamPolicyLicenseRequest, opts ...gax.CallOption) (*computepb.Policy, error) {
	return c.internalClient.SetIamPolicy(ctx, req, opts...)
}

// TestIamPermissions returns permissions that a caller has on the specified resource. Caution This resource is intended for use only by third-party partners who are creating Cloud Marketplace images.
func (c *LicensesClient) TestIamPermissions(ctx context.Context, req *computepb.TestIamPermissionsLicenseRequest, opts ...gax.CallOption) (*computepb.TestPermissionsResponse, error) {
	return c.internalClient.TestIamPermissions(ctx, req, opts...)
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type licensesRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewLicensesRESTClient creates a new licenses rest client.
//
// The Licenses API.
func NewLicensesRESTClient(ctx context.Context, opts ...option.ClientOption) (*LicensesClient, error) {
	clientOpts := append(defaultLicensesRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	c := &licensesRESTClient{
		endpoint:   endpoint,
		httpClient: httpClient,
	}
	c.setGoogleClientInfo()

	return &LicensesClient{internalClient: c, CallOptions: &LicensesCallOptions{}}, nil
}

func defaultLicensesRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://compute.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://compute.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://compute.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *licensesRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *licensesRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *licensesRESTClient) Connection() *grpc.ClientConn {
	return nil
}

// Delete deletes the specified license. Caution This resource is intended for use only by third-party partners who are creating Cloud Marketplace images.
func (c *licensesRESTClient) Delete(ctx context.Context, req *computepb.DeleteLicenseRequest, opts ...gax.CallOption) (*Operation, error) {
	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/licenses/%v", req.GetProject(), req.GetLicense())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("DELETE", baseUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if err = googleapi.CheckResponse(httpRsp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, maybeUnknownEnum(err)
	}
	op := &Operation{proto: rsp}
	return op, err
}

// Get returns the specified License resource. Caution This resource is intended for use only by third-party partners who are creating Cloud Marketplace images.
func (c *licensesRESTClient) Get(ctx context.Context, req *computepb.GetLicenseRequest, opts ...gax.CallOption) (*computepb.License, error) {
	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/licenses/%v", req.GetProject(), req.GetLicense())

	httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if err = googleapi.CheckResponse(httpRsp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.License{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, maybeUnknownEnum(err)
	}
	return rsp, nil
}

// GetIamPolicy gets the access control policy for a resource. May be empty if no such policy or resource exists. Caution This resource is intended for use only by third-party partners who are creating Cloud Marketplace images.
func (c *licensesRESTClient) GetIamPolicy(ctx context.Context, req *computepb.GetIamPolicyLicenseRequest, opts ...gax.CallOption) (*computepb.Policy, error) {
	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/licenses/%v/getIamPolicy", req.GetProject(), req.GetResource())

	params := url.Values{}
	if req != nil && req.OptionsRequestedPolicyVersion != nil {
		params.Add("optionsRequestedPolicyVersion", fmt.Sprintf("%v", req.GetOptionsRequestedPolicyVersion()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if err = googleapi.CheckResponse(httpRsp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Policy{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, maybeUnknownEnum(err)
	}
	return rsp, nil
}

// Insert create a License resource in the specified project. Caution This resource is intended for use only by third-party partners who are creating Cloud Marketplace images.
func (c *licensesRESTClient) Insert(ctx context.Context, req *computepb.InsertLicenseRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetLicenseResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/licenses", req.GetProject())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if err = googleapi.CheckResponse(httpRsp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, maybeUnknownEnum(err)
	}
	op := &Operation{proto: rsp}
	return op, err
}

// List retrieves the list of licenses available in the specified project. This method does not get any licenses that belong to other projects, including licenses attached to publicly-available images, like Debian 9. If you want to get a list of publicly-available licenses, use this method to make a request to the respective image project, such as debian-cloud or windows-cloud. Caution This resource is intended for use only by third-party partners who are creating Cloud Marketplace images.
func (c *licensesRESTClient) List(ctx context.Context, req *computepb.ListLicensesRequest, opts ...gax.CallOption) *LicenseIterator {
	it := &LicenseIterator{}
	req = proto.Clone(req).(*computepb.ListLicensesRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.License, string, error) {
		resp := &computepb.LicensesListResponse{}
		if pageToken != "" {
			req.PageToken = proto.String(pageToken)
		}
		if pageSize > math.MaxInt32 {
			req.MaxResults = proto.Uint32(math.MaxInt32)
		} else if pageSize != 0 {
			req.MaxResults = proto.Uint32(uint32(pageSize))
		}
		baseUrl, _ := url.Parse(c.endpoint)
		baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/licenses", req.GetProject())

		params := url.Values{}
		if req != nil && req.Filter != nil {
			params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
		}
		if req != nil && req.MaxResults != nil {
			params.Add("maxResults", fmt.Sprintf("%v", req.GetMaxResults()))
		}
		if req != nil && req.OrderBy != nil {
			params.Add("orderBy", fmt.Sprintf("%v", req.GetOrderBy()))
		}
		if req != nil && req.PageToken != nil {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}
		if req != nil && req.ReturnPartialSuccess != nil {
			params.Add("returnPartialSuccess", fmt.Sprintf("%v", req.GetReturnPartialSuccess()))
		}

		baseUrl.RawQuery = params.Encode()

		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return nil, "", err
		}

		// Set the headers
		for k, v := range c.xGoogMetadata {
			httpReq.Header[k] = v
		}

		httpReq.Header["Content-Type"] = []string{"application/json"}
		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return nil, "", err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return nil, "", err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return nil, "", err
		}

		unm.Unmarshal(buf, resp)
		it.Response = resp
		return resp.GetItems(), resp.GetNextPageToken(), nil
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
	it.pageInfo.MaxSize = int(req.GetMaxResults())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// SetIamPolicy sets the access control policy on the specified resource. Replaces any existing policy. Caution This resource is intended for use only by third-party partners who are creating Cloud Marketplace images.
func (c *licensesRESTClient) SetIamPolicy(ctx context.Context, req *computepb.SetIamPolicyLicenseRequest, opts ...gax.CallOption) (*computepb.Policy, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetGlobalSetPolicyRequestResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/licenses/%v/setIamPolicy", req.GetProject(), req.GetResource())

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if err = googleapi.CheckResponse(httpRsp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Policy{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, maybeUnknownEnum(err)
	}
	return rsp, nil
}

// TestIamPermissions returns permissions that a caller has on the specified resource. Caution This resource is intended for use only by third-party partners who are creating Cloud Marketplace images.
func (c *licensesRESTClient) TestIamPermissions(ctx context.Context, req *computepb.TestIamPermissionsLicenseRequest, opts ...gax.CallOption) (*computepb.TestPermissionsResponse, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetTestPermissionsRequestResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/licenses/%v/testIamPermissions", req.GetProject(), req.GetResource())

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if err = googleapi.CheckResponse(httpRsp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.TestPermissionsResponse{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, maybeUnknownEnum(err)
	}
	return rsp, nil
}

// LicenseIterator manages a stream of *computepb.License.
type LicenseIterator struct {
	items    []*computepb.License
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
	InternalFetch func(pageSize int, pageToken string) (results []*computepb.License, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *LicenseIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *LicenseIterator) Next() (*computepb.License, error) {
	var item *computepb.License
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *LicenseIterator) bufLen() int {
	return len(it.items)
}

func (it *LicenseIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
