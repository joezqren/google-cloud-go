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

//go:build go1.23

package networkconnectivity

import (
	"iter"

	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	networkconnectivitypb "cloud.google.com/go/networkconnectivity/apiv1/networkconnectivitypb"
	"github.com/googleapis/gax-go/v2/iterator"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
)

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *GroupIterator) All() iter.Seq2[*networkconnectivitypb.Group, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *HubIterator) All() iter.Seq2[*networkconnectivitypb.Hub, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *HubStatusEntryIterator) All() iter.Seq2[*networkconnectivitypb.HubStatusEntry, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *LocationIterator) All() iter.Seq2[*locationpb.Location, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *OperationIterator) All() iter.Seq2[*longrunningpb.Operation, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *PolicyBasedRouteIterator) All() iter.Seq2[*networkconnectivitypb.PolicyBasedRoute, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *RouteIterator) All() iter.Seq2[*networkconnectivitypb.Route, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *RouteTableIterator) All() iter.Seq2[*networkconnectivitypb.RouteTable, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *SpokeIterator) All() iter.Seq2[*networkconnectivitypb.Spoke, error] {
	return iterator.RangeAdapter(it.Next)
}
