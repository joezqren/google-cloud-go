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

package appengine_test

import (
	"context"

	appengine "cloud.google.com/go/appengine/apiv1"
	"google.golang.org/api/iterator"
	appenginepb "google.golang.org/genproto/googleapis/appengine/v1"
)

func ExampleNewServicesClient() {
	ctx := context.Background()
	c, err := appengine.NewServicesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleServicesClient_ListServices() {
	ctx := context.Background()
	c, err := appengine.NewServicesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &appenginepb.ListServicesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/appengine/v1#ListServicesRequest.
	}
	it := c.ListServices(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleServicesClient_GetService() {
	ctx := context.Background()
	c, err := appengine.NewServicesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &appenginepb.GetServiceRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/appengine/v1#GetServiceRequest.
	}
	resp, err := c.GetService(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleServicesClient_UpdateService() {
	ctx := context.Background()
	c, err := appengine.NewServicesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &appenginepb.UpdateServiceRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/appengine/v1#UpdateServiceRequest.
	}
	op, err := c.UpdateService(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleServicesClient_DeleteService() {
	ctx := context.Background()
	c, err := appengine.NewServicesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &appenginepb.DeleteServiceRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/appengine/v1#DeleteServiceRequest.
	}
	op, err := c.DeleteService(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}
