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

package dialogflow

import (
	"iter"

	dialogflowpb "cloud.google.com/go/dialogflow/apiv2/dialogflowpb"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/googleapis/gax-go/v2/iterator"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
)

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *AgentIterator) All() iter.Seq2[*dialogflowpb.Agent, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *AnswerRecordIterator) All() iter.Seq2[*dialogflowpb.AnswerRecord, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *ContextIterator) All() iter.Seq2[*dialogflowpb.Context, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *ConversationDatasetIterator) All() iter.Seq2[*dialogflowpb.ConversationDataset, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *ConversationIterator) All() iter.Seq2[*dialogflowpb.Conversation, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *ConversationModelEvaluationIterator) All() iter.Seq2[*dialogflowpb.ConversationModelEvaluation, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *ConversationModelIterator) All() iter.Seq2[*dialogflowpb.ConversationModel, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *ConversationProfileIterator) All() iter.Seq2[*dialogflowpb.ConversationProfile, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *DocumentIterator) All() iter.Seq2[*dialogflowpb.Document, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *EntityTypeIterator) All() iter.Seq2[*dialogflowpb.EntityType, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *EnvironmentHistory_EntryIterator) All() iter.Seq2[*dialogflowpb.EnvironmentHistory_Entry, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *EnvironmentIterator) All() iter.Seq2[*dialogflowpb.Environment, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *GeneratorIterator) All() iter.Seq2[*dialogflowpb.Generator, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *IntentIterator) All() iter.Seq2[*dialogflowpb.Intent, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *KnowledgeBaseIterator) All() iter.Seq2[*dialogflowpb.KnowledgeBase, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *LocationIterator) All() iter.Seq2[*locationpb.Location, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *MessageIterator) All() iter.Seq2[*dialogflowpb.Message, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *OperationIterator) All() iter.Seq2[*longrunningpb.Operation, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *ParticipantIterator) All() iter.Seq2[*dialogflowpb.Participant, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *SessionEntityTypeIterator) All() iter.Seq2[*dialogflowpb.SessionEntityType, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *VersionIterator) All() iter.Seq2[*dialogflowpb.Version, error] {
	return iterator.RangeAdapter(it.Next)
}
