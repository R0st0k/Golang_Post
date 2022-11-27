// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// PostcodesBySettlementGet implements GET /postcodes_by_settlement operation.
	//
	// Get information about postcodes in cities. Return map with `settlement` key and `postcode` array
	// value.
	//
	// GET /postcodes_by_settlement
	PostcodesBySettlementGet(ctx context.Context) (PostcodesBySettlementGetResponse, error)
	// SendingFilterGet implements GET /sending_filter operation.
	//
	// Get sendings that fit the filter. Require `page` and `elems_on_page`. Return amount of sendings
	// that fit the filter and sendings on the selected page.
	//
	// GET /sending_filter
	SendingFilterGet(ctx context.Context, params SendingFilterGetParams) (SendingFilterGetRes, error)
	// SendingGet implements GET /sending operation.
	//
	// Get information about a sending by `order_id`. Require a complete match of `order_id`. Return
	// `type`, `status` and `stages`.
	//
	// GET /sending
	SendingGet(ctx context.Context, params SendingGetParams) (SendingGetRes, error)
	// SendingPost implements POST /sending operation.
	//
	// Registration of a new sending. Require `type`, `sender`, `receiver`, `size`, `weight`. Return
	// `order_id` of new sending.
	//
	// POST /sending
	SendingPost(ctx context.Context, req SendingPostReq) (SendingPostRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...Option) (*Server, error) {
	s, err := newConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}
