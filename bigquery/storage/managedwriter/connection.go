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

package managedwriter

import (
	"context"
	"errors"
	"fmt"
	"io"
	"sync"

	"cloud.google.com/go/bigquery/storage/apiv1/storagepb"
	"github.com/googleapis/gax-go/v2"
	"go.opencensus.io/tag"
	"google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"
)

const (
	poolIDPrefix   string = "connectionpool"
	connIDPrefix   string = "connection"
	writerIDPrefix string = "writer"
)

var (
	errNoRouterForPool = errors.New("no router for connection pool")
)

// connectionPool represents a pooled set of connections.
//
// The pool retains references to connections, and maintains the mapping between writers
// and connections.
//
// TODO: connection and writer mappings will be added in a subsequent PR.
type connectionPool struct {
	id                   string
	allowMultipleWriters bool // whether this pool can be used by multiple writers.

	// the pool retains the long-lived context responsible for opening/maintaining bidi connections.
	ctx    context.Context
	cancel context.CancelFunc

	baseFlowController *flowController // template flow controller used for building connections.

	// We centralize the open function on the pool, rather than having an instance of the open func on every
	// connection.  Opening the connection is a stateless operation.
	open func(opts ...gax.CallOption) (storagepb.BigQueryWrite_AppendRowsClient, error)

	// We specify one set of calloptions for the pool.
	// All connections in the pool open with the same call options.
	callOptions []gax.CallOption

	router poolRouter // poolManager makes the decisions about connections and routing.

	retry *statelessRetryer // default retryer for the pool.
}

// activateRouter handles wiring up a connection pool and it's router.
func (pool *connectionPool) activateRouter(rtr poolRouter) error {
	if pool.router != nil {
		return fmt.Errorf("router already activated")
	}
	if err := rtr.poolAttach(pool); err != nil {
		return fmt.Errorf("router rejected attach: %w", err)
	}
	pool.router = rtr
	return nil
}

func (pool *connectionPool) Close() error {
	// Signal router and cancel context, which should propagate to all writers.
	var err error
	if pool.router != nil {
		err = pool.router.poolDetach()
	}
	if cancel := pool.cancel; cancel != nil {
		cancel()
	}
	return err
}

// pickConnection is used by writers to select a connection.
func (pool *connectionPool) selectConn(pw *pendingWrite) (*connection, error) {
	if pool.router == nil {
		return nil, errNoRouterForPool
	}
	return pool.router.pickConnection(pw)
}

func (pool *connectionPool) addWriter(writer *ManagedStream) error {
	if p := writer.pool; p != nil {
		return fmt.Errorf("writer already attached to pool %q", p.id)
	}
	if pool.router == nil {
		return errNoRouterForPool
	}
	if err := pool.router.writerAttach(writer); err != nil {
		return err
	}
	writer.pool = pool
	return nil
}

func (pool *connectionPool) removeWriter(writer *ManagedStream) error {
	if pool.router == nil {
		return errNoRouterForPool
	}
	detachErr := pool.router.writerDetach(writer)
	// trigger single-writer pool closure regardless of detach errors
	if !pool.allowMultipleWriters {
		if err := pool.Close(); detachErr == nil {
			detachErr = err
		}
	}
	return detachErr
}

// openWithRetry establishes a new bidi stream and channel pair.  It is used by connection objects
// when (re)opening the network connection to the backend.
//
// The connection.getStream() func should be the only consumer of this.
func (cp *connectionPool) openWithRetry(co *connection) (storagepb.BigQueryWrite_AppendRowsClient, chan *pendingWrite, error) {
	r := &unaryRetryer{}
	for {
		recordStat(cp.ctx, AppendClientOpenCount, 1)
		arc, err := cp.open(cp.callOptions...)
		if err != nil {
			bo, shouldRetry := r.Retry(err)
			if shouldRetry {
				recordStat(cp.ctx, AppendClientOpenRetryCount, 1)
				if err := gax.Sleep(cp.ctx, bo); err != nil {
					return nil, nil, err
				}
				continue
			} else {
				// non-retriable error while opening
				return nil, nil, err
			}
		}
		// The channel relationship with its ARC is 1:1.  If we get a new ARC, create a new pending
		// write channel and fire up the associated receive processor.  The channel ensures that
		// responses for a connection are processed in the same order that appends were sent.
		depth := 1000 // default backend queue limit
		if d := co.fc.maxInsertCount; d > 0 {
			depth = d
		}
		ch := make(chan *pendingWrite, depth)
		go connRecvProcessor(co, arc, ch)
		return arc, ch, nil
	}
}

// returns the stateless default retryer for the pool.  If one's not set (re-enqueue retries disabled),
// it returns a retryer that only permits single attempts.
func (cp *connectionPool) defaultRetryer() *statelessRetryer {
	if cp.retry != nil {
		return cp.retry
	}
	return &statelessRetryer{
		maxAttempts: 1,
	}
}

// connection models the underlying AppendRows grpc bidi connection used for writing
// data and receiving acknowledgements.  It is responsible for enqueing writes and processing
// responses from the backend.
type connection struct {
	id   string
	pool *connectionPool // each connection retains a reference to its owning pool.

	fc     *flowController // each connection has it's own flow controller.
	ctx    context.Context // retained context for maintaining the connection, derived from the owning pool.
	cancel context.CancelFunc

	retry     *statelessRetryer
	optimizer sendOptimizer

	mu        sync.Mutex
	arc       *storagepb.BigQueryWrite_AppendRowsClient // reference to the grpc connection (send, recv, close)
	reconnect bool                                      //
	err       error                                     // terminal connection error
	pending   chan *pendingWrite
}

func newConnection(pool *connectionPool, mode string) *connection {
	if pool == nil {
		return nil
	}
	// create and retain a cancellable context.
	connCtx, cancel := context.WithCancel(pool.ctx)
	fc := newFlowController(0, 0)
	if pool != nil {
		fc = copyFlowController(pool.baseFlowController)
	}
	return &connection{
		id:        newUUID(connIDPrefix),
		pool:      pool,
		fc:        fc,
		ctx:       connCtx,
		cancel:    cancel,
		optimizer: optimizer(mode),
	}
}

func optimizer(mode string) sendOptimizer {
	switch mode {
	case "MULTIPLEX":
		return &multiplexOptimizer{}
	case "VERBOSE":
		return &verboseOptimizer{}
	default:
		return &simplexOptimizer{}
	}
}

// release is used to signal flow control release when a write is no longer in flight.
func (co *connection) release(pw *pendingWrite) {
	co.fc.release(pw.reqSize)
}

// close closes a connection.
func (co *connection) close() {
	co.mu.Lock()
	defer co.mu.Unlock()
	// first, cancel the retained context.
	if co.cancel != nil {
		co.cancel()
		co.cancel = nil
	}
	// close sending if we have a real ARC.
	if co.arc != nil && (*co.arc) != (storagepb.BigQueryWrite_AppendRowsClient)(nil) {
		(*co.arc).CloseSend()
		co.arc = nil
	}
	// mark terminal error if not already set.
	if co.err != nil {
		co.err = io.EOF
	}
	// signal pending channel close.
	if co.pending != nil {
		close(co.pending)
	}
}

// lockingAppend handles a single append request on a given connection.
func (co *connection) lockingAppend(pw *pendingWrite) error {
	// Don't both calling/retrying if this append's context is already expired.
	if err := pw.reqCtx.Err(); err != nil {
		return err
	}

	if err := co.fc.acquire(pw.reqCtx, pw.reqSize); err != nil {
		// We've failed to acquire.  This may get retried on a different connection, so marking the write done is incorrect.
		return err
	}

	var statsOnExit func()

	// critical section:  Things that need to happen inside the critical section:
	//
	// * get/open conenction
	// * issue the append
	// * add the pending write to the channel for the connection (ordering for the response)
	co.mu.Lock()
	defer func() {
		co.mu.Unlock()
		if statsOnExit != nil {
			statsOnExit()
		}
	}()

	var arc *storagepb.BigQueryWrite_AppendRowsClient
	var ch chan *pendingWrite
	var err error

	// We still need to reconnect if we need to signal a new schema for explicit streams.
	// Rather than adding more state to the connection, we just look at the request as we
	// do not allow multiplexing to include explicit streams.
	forceReconnect := false
	if !canMultiplex(pw.writeStreamID) {
		if pw.writer != nil && pw.descVersion != nil && pw.descVersion.isNewer(pw.writer.curDescVersion) {
			forceReconnect = true
			pw.writer.curDescVersion = pw.descVersion
		}
	}

	arc, ch, err = co.getStream(arc, forceReconnect)
	if err != nil {
		return err
	}

	pw.attemptCount = pw.attemptCount + 1
	if co.optimizer != nil {
		err = co.optimizer.optimizeSend((*arc), pw)
		if err != nil {
			// Reset optimizer state on error.
			co.optimizer.signalReset()
		}
	} else {
		// No optimizer present, send a fully populated request.
		err = (*arc).Send(pw.constructFullRequest(true))
	}
	if err != nil {
		if shouldReconnect(err) {
			// if we think this connection is unhealthy, force a reconnect on the next send.
			co.reconnect = true
		}
		return err
	}

	// Compute numRows, once we pass ownership to the channel the request may be
	// cleared.
	var numRows int64
	if r := pw.req.GetProtoRows(); r != nil {
		if pr := r.GetRows(); pr != nil {
			numRows = int64(len(pr.GetSerializedRows()))
		}
	}
	statsOnExit = func() {
		// these will get recorded once we exit the critical section.
		// TODO: resolve open questions around what labels should be attached (connection, streamID, etc)
		recordStat(co.ctx, AppendRequestRows, numRows)
		recordStat(co.ctx, AppendRequests, 1)
		recordStat(co.ctx, AppendRequestBytes, int64(pw.reqSize))
	}
	ch <- pw
	return nil
}

// getStream returns either a valid ARC client stream or permanent error.
//
// Any calls to getStream should do so in possesion of the critical section lock.
func (co *connection) getStream(arc *storagepb.BigQueryWrite_AppendRowsClient, forceReconnect bool) (*storagepb.BigQueryWrite_AppendRowsClient, chan *pendingWrite, error) {
	if co.err != nil {
		return nil, nil, co.err
	}
	co.err = co.ctx.Err()
	if co.err != nil {
		return nil, nil, co.err
	}

	// Previous activity on the stream indicated it is not healthy, so propagate that as a reconnect.
	if co.reconnect {
		forceReconnect = true
		co.reconnect = false
	}
	// Always return the retained ARC if the arg differs.
	if arc != co.arc && !forceReconnect {
		return co.arc, co.pending, nil
	}
	// We need to (re)open a connection.  Cleanup previous connection and channel if they are present.
	if co.arc != nil && (*co.arc) != (storagepb.BigQueryWrite_AppendRowsClient)(nil) {
		(*co.arc).CloseSend()
	}
	if co.pending != nil {
		close(co.pending)
	}

	co.arc = new(storagepb.BigQueryWrite_AppendRowsClient)
	// We're going to (re)open the connection, so clear any optimizer state.
	if co.optimizer != nil {
		co.optimizer.signalReset()
	}
	*co.arc, co.pending, co.err = co.pool.openWithRetry(co)
	return co.arc, co.pending, co.err
}

// enables testing
type streamClientFunc func(context.Context, ...gax.CallOption) (storagepb.BigQueryWrite_AppendRowsClient, error)

// connRecvProcessor is used to propagate append responses back up with the originating write requests.  It
// It runs as a goroutine.  A connection object allows for reconnection, and each reconnection establishes a new
// processing gorouting and backing channel.
func connRecvProcessor(co *connection, arc storagepb.BigQueryWrite_AppendRowsClient, ch <-chan *pendingWrite) {
	for {
		select {
		case <-co.ctx.Done():
			// Context is done, so we're not going to get further updates.  Mark all work left in the channel
			// with the context error.  We don't attempt to re-enqueue in this case.
			for {
				pw, ok := <-ch
				if !ok {
					return
				}
				// It's unlikely this connection will recover here, but for correctness keep the flow controller
				// state correct by releasing.
				co.release(pw)
				pw.markDone(nil, co.ctx.Err())
			}
		case nextWrite, ok := <-ch:
			if !ok {
				// Channel closed, all elements processed.
				return
			}
			// block until we get a corresponding response or err from stream.
			resp, err := arc.Recv()
			co.release(nextWrite)
			if err != nil {
				nextWrite.writer.processRetry(nextWrite, co, nil, err)
				continue
			}
			// Record that we did in fact get a response from the backend.
			recordStat(co.ctx, AppendResponses, 1)

			if status := resp.GetError(); status != nil {
				// The response from the backend embedded a status error.  We record that the error
				// occurred, and tag it based on the response code of the status.
				if tagCtx, tagErr := tag.New(co.ctx, tag.Insert(keyError, codes.Code(status.GetCode()).String())); tagErr == nil {
					recordStat(tagCtx, AppendResponseErrors, 1)
				}
				respErr := grpcstatus.ErrorProto(status)

				nextWrite.writer.processRetry(nextWrite, co, resp, respErr)

				continue
			}
			// We had no error in the receive or in the response.  Mark the write done.
			nextWrite.markDone(resp, nil)
		}
	}
}

type poolRouter interface {

	// poolAttach is called once to signal a router that it is responsible for a given pool.
	poolAttach(pool *connectionPool) error

	// poolDetach is called as part of clean connectionPool shutdown.
	// It provides an opportunity for the router to shut down internal state.
	poolDetach() error

	// writerAttach is a hook to notify the router that a new writer is being attached to the pool.
	// It provides an opportunity for the router to allocate resources and update internal state.
	writerAttach(writer *ManagedStream) error

	// writerAttach signals the router that a given writer is being removed from the pool.  The router
	// does not have responsibility for closing the writer, but this is called as part of writer close.
	writerDetach(writer *ManagedStream) error

	// pickConnection is used to select a connection for a given pending write.
	pickConnection(pw *pendingWrite) (*connection, error)
}

// simpleRouter is a primitive traffic router that routes all traffic to its single connection instance.
//
// This router is designed for our migration case, where an single ManagedStream writer has as 1:1 relationship
// with a connectionPool.  You can multiplex with this router, but it will never scale beyond a single connection.
type simpleRouter struct {
	mode string
	pool *connectionPool

	mu      sync.RWMutex
	conn    *connection
	writers map[string]struct{}
}

func (rtr *simpleRouter) poolAttach(pool *connectionPool) error {
	if rtr.pool == nil {
		rtr.pool = pool
		return nil
	}
	return fmt.Errorf("router already attached to pool %q", rtr.pool.id)
}

func (rtr *simpleRouter) poolDetach() error {
	rtr.mu.Lock()
	defer rtr.mu.Unlock()
	if rtr.conn != nil {
		rtr.conn.close()
		rtr.conn = nil
	}
	return nil
}

func (rtr *simpleRouter) writerAttach(writer *ManagedStream) error {
	if writer.id == "" {
		return fmt.Errorf("writer has no ID")
	}
	rtr.mu.Lock()
	defer rtr.mu.Unlock()
	rtr.writers[writer.id] = struct{}{}
	if rtr.conn == nil {
		rtr.conn = newConnection(rtr.pool, rtr.mode)
	}
	return nil
}

func (rtr *simpleRouter) writerDetach(writer *ManagedStream) error {
	if writer.id == "" {
		return fmt.Errorf("writer has no ID")
	}
	rtr.mu.Lock()
	defer rtr.mu.Unlock()
	delete(rtr.writers, writer.id)
	if len(rtr.writers) == 0 && rtr.conn != nil {
		// no attached writers, cleanup and remove connection.
		defer rtr.conn.close()
		rtr.conn = nil
	}
	return nil
}

// Picking a connection is easy; there's only one.
func (rtr *simpleRouter) pickConnection(pw *pendingWrite) (*connection, error) {
	rtr.mu.RLock()
	defer rtr.mu.RUnlock()
	if rtr.conn != nil {
		return rtr.conn, nil
	}
	return nil, fmt.Errorf("no connection available")
}

func newSimpleRouter(mode string) *simpleRouter {
	return &simpleRouter{
		// We don't add a connection until writers attach.
		mode:    mode,
		writers: make(map[string]struct{}),
	}
}
