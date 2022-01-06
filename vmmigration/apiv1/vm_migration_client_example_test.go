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

package vmmigration_test

import (
	"context"

	vmmigration "cloud.google.com/go/vmmigration/apiv1"
	"google.golang.org/api/iterator"
	vmmigrationpb "google.golang.org/genproto/googleapis/cloud/vmmigration/v1"
)

func ExampleNewClient() {
	ctx := context.Background()
	c, err := vmmigration.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleClient_ListSources() {
	ctx := context.Background()
	c, err := vmmigration.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmmigrationpb.ListSourcesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vmmigration/v1#ListSourcesRequest.
	}
	it := c.ListSources(ctx, req)
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

func ExampleClient_GetSource() {
	ctx := context.Background()
	c, err := vmmigration.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmmigrationpb.GetSourceRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vmmigration/v1#GetSourceRequest.
	}
	resp, err := c.GetSource(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_CreateSource() {
	ctx := context.Background()
	c, err := vmmigration.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmmigrationpb.CreateSourceRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vmmigration/v1#CreateSourceRequest.
	}
	op, err := c.CreateSource(ctx, req)
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

func ExampleClient_UpdateSource() {
	ctx := context.Background()
	c, err := vmmigration.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmmigrationpb.UpdateSourceRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vmmigration/v1#UpdateSourceRequest.
	}
	op, err := c.UpdateSource(ctx, req)
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

func ExampleClient_DeleteSource() {
	ctx := context.Background()
	c, err := vmmigration.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmmigrationpb.DeleteSourceRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vmmigration/v1#DeleteSourceRequest.
	}
	op, err := c.DeleteSource(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_FetchInventory() {
	ctx := context.Background()
	c, err := vmmigration.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmmigrationpb.FetchInventoryRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vmmigration/v1#FetchInventoryRequest.
	}
	resp, err := c.FetchInventory(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListUtilizationReports() {
	ctx := context.Background()
	c, err := vmmigration.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmmigrationpb.ListUtilizationReportsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vmmigration/v1#ListUtilizationReportsRequest.
	}
	it := c.ListUtilizationReports(ctx, req)
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

func ExampleClient_GetUtilizationReport() {
	ctx := context.Background()
	c, err := vmmigration.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmmigrationpb.GetUtilizationReportRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vmmigration/v1#GetUtilizationReportRequest.
	}
	resp, err := c.GetUtilizationReport(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_CreateUtilizationReport() {
	ctx := context.Background()
	c, err := vmmigration.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmmigrationpb.CreateUtilizationReportRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vmmigration/v1#CreateUtilizationReportRequest.
	}
	op, err := c.CreateUtilizationReport(ctx, req)
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

func ExampleClient_DeleteUtilizationReport() {
	ctx := context.Background()
	c, err := vmmigration.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmmigrationpb.DeleteUtilizationReportRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vmmigration/v1#DeleteUtilizationReportRequest.
	}
	op, err := c.DeleteUtilizationReport(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_ListDatacenterConnectors() {
	ctx := context.Background()
	c, err := vmmigration.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmmigrationpb.ListDatacenterConnectorsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vmmigration/v1#ListDatacenterConnectorsRequest.
	}
	it := c.ListDatacenterConnectors(ctx, req)
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

func ExampleClient_GetDatacenterConnector() {
	ctx := context.Background()
	c, err := vmmigration.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmmigrationpb.GetDatacenterConnectorRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vmmigration/v1#GetDatacenterConnectorRequest.
	}
	resp, err := c.GetDatacenterConnector(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_CreateDatacenterConnector() {
	ctx := context.Background()
	c, err := vmmigration.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmmigrationpb.CreateDatacenterConnectorRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vmmigration/v1#CreateDatacenterConnectorRequest.
	}
	op, err := c.CreateDatacenterConnector(ctx, req)
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

func ExampleClient_DeleteDatacenterConnector() {
	ctx := context.Background()
	c, err := vmmigration.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmmigrationpb.DeleteDatacenterConnectorRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vmmigration/v1#DeleteDatacenterConnectorRequest.
	}
	op, err := c.DeleteDatacenterConnector(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_CreateMigratingVm() {
	ctx := context.Background()
	c, err := vmmigration.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmmigrationpb.CreateMigratingVmRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vmmigration/v1#CreateMigratingVmRequest.
	}
	op, err := c.CreateMigratingVm(ctx, req)
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

func ExampleClient_ListMigratingVms() {
	ctx := context.Background()
	c, err := vmmigration.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmmigrationpb.ListMigratingVmsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vmmigration/v1#ListMigratingVmsRequest.
	}
	it := c.ListMigratingVms(ctx, req)
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

func ExampleClient_GetMigratingVm() {
	ctx := context.Background()
	c, err := vmmigration.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmmigrationpb.GetMigratingVmRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vmmigration/v1#GetMigratingVmRequest.
	}
	resp, err := c.GetMigratingVm(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_UpdateMigratingVm() {
	ctx := context.Background()
	c, err := vmmigration.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmmigrationpb.UpdateMigratingVmRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vmmigration/v1#UpdateMigratingVmRequest.
	}
	op, err := c.UpdateMigratingVm(ctx, req)
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

func ExampleClient_DeleteMigratingVm() {
	ctx := context.Background()
	c, err := vmmigration.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmmigrationpb.DeleteMigratingVmRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vmmigration/v1#DeleteMigratingVmRequest.
	}
	op, err := c.DeleteMigratingVm(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_StartMigration() {
	ctx := context.Background()
	c, err := vmmigration.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmmigrationpb.StartMigrationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vmmigration/v1#StartMigrationRequest.
	}
	op, err := c.StartMigration(ctx, req)
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

func ExampleClient_ResumeMigration() {
	ctx := context.Background()
	c, err := vmmigration.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmmigrationpb.ResumeMigrationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vmmigration/v1#ResumeMigrationRequest.
	}
	op, err := c.ResumeMigration(ctx, req)
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

func ExampleClient_PauseMigration() {
	ctx := context.Background()
	c, err := vmmigration.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmmigrationpb.PauseMigrationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vmmigration/v1#PauseMigrationRequest.
	}
	op, err := c.PauseMigration(ctx, req)
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

func ExampleClient_FinalizeMigration() {
	ctx := context.Background()
	c, err := vmmigration.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmmigrationpb.FinalizeMigrationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vmmigration/v1#FinalizeMigrationRequest.
	}
	op, err := c.FinalizeMigration(ctx, req)
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

func ExampleClient_CreateCloneJob() {
	ctx := context.Background()
	c, err := vmmigration.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmmigrationpb.CreateCloneJobRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vmmigration/v1#CreateCloneJobRequest.
	}
	op, err := c.CreateCloneJob(ctx, req)
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

func ExampleClient_CancelCloneJob() {
	ctx := context.Background()
	c, err := vmmigration.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmmigrationpb.CancelCloneJobRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vmmigration/v1#CancelCloneJobRequest.
	}
	op, err := c.CancelCloneJob(ctx, req)
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

func ExampleClient_ListCloneJobs() {
	ctx := context.Background()
	c, err := vmmigration.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmmigrationpb.ListCloneJobsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vmmigration/v1#ListCloneJobsRequest.
	}
	it := c.ListCloneJobs(ctx, req)
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

func ExampleClient_GetCloneJob() {
	ctx := context.Background()
	c, err := vmmigration.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmmigrationpb.GetCloneJobRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vmmigration/v1#GetCloneJobRequest.
	}
	resp, err := c.GetCloneJob(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_CreateCutoverJob() {
	ctx := context.Background()
	c, err := vmmigration.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmmigrationpb.CreateCutoverJobRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vmmigration/v1#CreateCutoverJobRequest.
	}
	op, err := c.CreateCutoverJob(ctx, req)
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

func ExampleClient_CancelCutoverJob() {
	ctx := context.Background()
	c, err := vmmigration.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmmigrationpb.CancelCutoverJobRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vmmigration/v1#CancelCutoverJobRequest.
	}
	op, err := c.CancelCutoverJob(ctx, req)
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

func ExampleClient_ListCutoverJobs() {
	ctx := context.Background()
	c, err := vmmigration.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmmigrationpb.ListCutoverJobsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vmmigration/v1#ListCutoverJobsRequest.
	}
	it := c.ListCutoverJobs(ctx, req)
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

func ExampleClient_GetCutoverJob() {
	ctx := context.Background()
	c, err := vmmigration.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmmigrationpb.GetCutoverJobRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vmmigration/v1#GetCutoverJobRequest.
	}
	resp, err := c.GetCutoverJob(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListGroups() {
	ctx := context.Background()
	c, err := vmmigration.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmmigrationpb.ListGroupsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vmmigration/v1#ListGroupsRequest.
	}
	it := c.ListGroups(ctx, req)
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

func ExampleClient_GetGroup() {
	ctx := context.Background()
	c, err := vmmigration.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmmigrationpb.GetGroupRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vmmigration/v1#GetGroupRequest.
	}
	resp, err := c.GetGroup(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_CreateGroup() {
	ctx := context.Background()
	c, err := vmmigration.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmmigrationpb.CreateGroupRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vmmigration/v1#CreateGroupRequest.
	}
	op, err := c.CreateGroup(ctx, req)
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

func ExampleClient_UpdateGroup() {
	ctx := context.Background()
	c, err := vmmigration.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmmigrationpb.UpdateGroupRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vmmigration/v1#UpdateGroupRequest.
	}
	op, err := c.UpdateGroup(ctx, req)
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

func ExampleClient_DeleteGroup() {
	ctx := context.Background()
	c, err := vmmigration.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmmigrationpb.DeleteGroupRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vmmigration/v1#DeleteGroupRequest.
	}
	op, err := c.DeleteGroup(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_AddGroupMigration() {
	ctx := context.Background()
	c, err := vmmigration.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmmigrationpb.AddGroupMigrationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vmmigration/v1#AddGroupMigrationRequest.
	}
	op, err := c.AddGroupMigration(ctx, req)
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

func ExampleClient_RemoveGroupMigration() {
	ctx := context.Background()
	c, err := vmmigration.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmmigrationpb.RemoveGroupMigrationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vmmigration/v1#RemoveGroupMigrationRequest.
	}
	op, err := c.RemoveGroupMigration(ctx, req)
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

func ExampleClient_ListTargetProjects() {
	ctx := context.Background()
	c, err := vmmigration.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmmigrationpb.ListTargetProjectsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vmmigration/v1#ListTargetProjectsRequest.
	}
	it := c.ListTargetProjects(ctx, req)
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

func ExampleClient_GetTargetProject() {
	ctx := context.Background()
	c, err := vmmigration.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmmigrationpb.GetTargetProjectRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vmmigration/v1#GetTargetProjectRequest.
	}
	resp, err := c.GetTargetProject(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_CreateTargetProject() {
	ctx := context.Background()
	c, err := vmmigration.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmmigrationpb.CreateTargetProjectRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vmmigration/v1#CreateTargetProjectRequest.
	}
	op, err := c.CreateTargetProject(ctx, req)
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

func ExampleClient_UpdateTargetProject() {
	ctx := context.Background()
	c, err := vmmigration.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmmigrationpb.UpdateTargetProjectRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vmmigration/v1#UpdateTargetProjectRequest.
	}
	op, err := c.UpdateTargetProject(ctx, req)
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

func ExampleClient_DeleteTargetProject() {
	ctx := context.Background()
	c, err := vmmigration.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmmigrationpb.DeleteTargetProjectRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vmmigration/v1#DeleteTargetProjectRequest.
	}
	op, err := c.DeleteTargetProject(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}
