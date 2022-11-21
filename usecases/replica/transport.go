//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

package replica

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/semi-technologies/weaviate/entities/storobj"
	"github.com/semi-technologies/weaviate/usecases/objects"
)

type SimpleResponse struct {
	Errors []string `json:"errors,omitempty"`
}

func (r *SimpleResponse) FirstError() error {
	for _, msg := range r.Errors {
		if msg != "" {
			return &Error{Msg: msg}
		}
	}
	return nil
}

type UUID2Error struct {
	UUID  string `json:"uuid,omitempty"`
	Error string `json:"error,omitempty"`
}

// DeleteBatchResponse represents the response returned by DeleteObjects
type DeleteBatchResponse struct {
	Batch []UUID2Error `json:"batch,omitempty"`
}

// FirstError returns the first found error
func (r *DeleteBatchResponse) FirstError() error {
	for _, r := range r.Batch {
		if r.Error != "" {
			return &Error{Msg: r.Error}
		}
	}
	return nil
}

type client interface {
	PutObject(ctx context.Context, host, index, shard, requestID string,
		obj *storobj.Object) (SimpleResponse, error)
	DeleteObject(ctx context.Context, host, index, shard, requestID string,
		id strfmt.UUID) (SimpleResponse, error)
	PutObjects(ctx context.Context, host, index, shard, requestID string,
		objs []*storobj.Object) (SimpleResponse, error)
	MergeObject(ctx context.Context, host, index, shard, requestID string,
		mergeDoc *objects.MergeDocument) (SimpleResponse, error)
	DeleteObjects(ctx context.Context, host, index, shard, requestID string,
		docIDs []uint64, dryRun bool) (SimpleResponse, error)

	AddReferences(ctx context.Context, host, index, shard, requestID string,
		refs objects.BatchReferences) (SimpleResponse, error)

	Commit(ctx context.Context, host, index, shard, requestID string, resp interface{}) error

	Abort(ctx context.Context, host, index, shard, requestID string) (SimpleResponse, error)
}