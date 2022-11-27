// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"net/http"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"

	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
)

// handlePostcodesBySettlementGetRequest handles GET /postcodes_by_settlement operation.
//
// Get information about postcodes in cities. Return map with `settlement` key and `postcode` array
// value.
//
// GET /postcodes_by_settlement
func (s *Server) handlePostcodesBySettlementGetRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	var otelAttrs []attribute.KeyValue

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PostcodesBySettlementGet",
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, otelAttrs...)

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, otelAttrs...)
		}
		err error
	)

	var response PostcodesBySettlementGetResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "PostcodesBySettlementGet",
			OperationID:   "",
			Body:          nil,
			Params:        map[string]any{},
			Raw:           r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = PostcodesBySettlementGetResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (Response, error) {
				return s.h.PostcodesBySettlementGet(ctx)
			},
		)
	} else {
		response, err = s.h.PostcodesBySettlementGet(ctx)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodePostcodesBySettlementGetResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handleSendingFilterGetRequest handles GET /sending_filter operation.
//
// Get sendings that fit the filter. Require `page` and `elems_on_page`. Return amount of sendings
// that fit the filter and sendings on the selected page.
//
// GET /sending_filter
func (s *Server) handleSendingFilterGetRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	var otelAttrs []attribute.KeyValue

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "SendingFilterGet",
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, otelAttrs...)

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, otelAttrs...)
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "SendingFilterGet",
			ID:   "",
		}
	)
	params, err := decodeSendingFilterGetParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response SendingFilterGetRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "SendingFilterGet",
			OperationID:   "",
			Body:          nil,
			Params: map[string]any{
				"page":          params.Page,
				"elems_on_page": params.ElemsOnPage,
				"filter":        params.Filter,
				"sort":          params.Sort,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = SendingFilterGetParams
			Response = SendingFilterGetRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackSendingFilterGetParams,
			func(ctx context.Context, request Request, params Params) (Response, error) {
				return s.h.SendingFilterGet(ctx, params)
			},
		)
	} else {
		response, err = s.h.SendingFilterGet(ctx, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeSendingFilterGetResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handleSendingGetRequest handles GET /sending operation.
//
// Get information about a sending by `order_id`. Require a complete match of `order_id`. Return
// `type`, `status` and `stages`.
//
// GET /sending
func (s *Server) handleSendingGetRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	var otelAttrs []attribute.KeyValue

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "SendingGet",
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, otelAttrs...)

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, otelAttrs...)
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "SendingGet",
			ID:   "",
		}
	)
	params, err := decodeSendingGetParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response SendingGetRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "SendingGet",
			OperationID:   "",
			Body:          nil,
			Params: map[string]any{
				"order_id": params.OrderID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = SendingGetParams
			Response = SendingGetRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackSendingGetParams,
			func(ctx context.Context, request Request, params Params) (Response, error) {
				return s.h.SendingGet(ctx, params)
			},
		)
	} else {
		response, err = s.h.SendingGet(ctx, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeSendingGetResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handleSendingPostRequest handles POST /sending operation.
//
// Registration of a new sending. Require `type`, `sender`, `receiver`, `size`, `weight`. Return
// `order_id` of new sending.
//
// POST /sending
func (s *Server) handleSendingPostRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	var otelAttrs []attribute.KeyValue

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "SendingPost",
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, otelAttrs...)

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, otelAttrs...)
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "SendingPost",
			ID:   "",
		}
	)
	request, close, err := s.decodeSendingPostRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response SendingPostRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "SendingPost",
			OperationID:   "",
			Body:          request,
			Params:        map[string]any{},
			Raw:           r,
		}

		type (
			Request  = SendingPostReq
			Params   = struct{}
			Response = SendingPostRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (Response, error) {
				return s.h.SendingPost(ctx, request)
			},
		)
	} else {
		response, err = s.h.SendingPost(ctx, request)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeSendingPostResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}
