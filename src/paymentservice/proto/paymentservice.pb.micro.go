// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/paymentservice.proto

package retrostation

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for PaymentService service

func NewPaymentServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for PaymentService service

type PaymentService interface {
	Charge(ctx context.Context, in *ChargeRequest, opts ...client.CallOption) (*ChargeResponse, error)
}

type paymentService struct {
	c    client.Client
	name string
}

func NewPaymentService(name string, c client.Client) PaymentService {
	return &paymentService{
		c:    c,
		name: name,
	}
}

func (c *paymentService) Charge(ctx context.Context, in *ChargeRequest, opts ...client.CallOption) (*ChargeResponse, error) {
	req := c.c.NewRequest(c.name, "PaymentService.Charge", in)
	out := new(ChargeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PaymentService service

type PaymentServiceHandler interface {
	Charge(context.Context, *ChargeRequest, *ChargeResponse) error
}

func RegisterPaymentServiceHandler(s server.Server, hdlr PaymentServiceHandler, opts ...server.HandlerOption) error {
	type paymentService interface {
		Charge(ctx context.Context, in *ChargeRequest, out *ChargeResponse) error
	}
	type PaymentService struct {
		paymentService
	}
	h := &paymentServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&PaymentService{h}, opts...))
}

type paymentServiceHandler struct {
	PaymentServiceHandler
}

func (h *paymentServiceHandler) Charge(ctx context.Context, in *ChargeRequest, out *ChargeResponse) error {
	return h.PaymentServiceHandler.Charge(ctx, in, out)
}
