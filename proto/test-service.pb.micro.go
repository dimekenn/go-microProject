// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/test-service.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

import (
	context "context"
	api "github.com/unistack-org/micro/v3/api"
	client "github.com/unistack-org/micro/v3/client"
	server "github.com/unistack-org/micro/v3/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for PantoCustomerService service

func NewPantoCustomerServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		&api.Endpoint{
			Name:    "PantoCustomerService.PanById",
			Path:    []string{"/api/v1/messages"},
			Method:  []string{"POST"},
			Body:    "",
			Handler: "rpc",
		},
	}
}

// Client API for PantoCustomerService service

type PantoCustomerService interface {
	PanById(ctx context.Context, req *PanRequest, opts ...client.CallOption) (*PanResponse, error)
}

type pantoCustomerService struct {
	c    client.Client
	name string
}

func NewPantoCustomerService(name string, c client.Client) PantoCustomerService {
	return &pantoCustomerService{
		c:    c,
		name: name,
	}
}

func (c *pantoCustomerService) PanById(ctx context.Context, req *PanRequest, opts ...client.CallOption) (*PanResponse, error) {
	rsp := &PanResponse{}
	err := c.c.Call(ctx, c.c.NewRequest(c.name, "PantoCustomerService.PanById", req), rsp, opts...)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

// Server API for PantoCustomerService service

type PantoCustomerServiceHandler interface {
	PanById(context.Context, *PanRequest, *PanResponse) error
}

func RegisterPantoCustomerServiceHandler(s server.Server, hdlr PantoCustomerServiceHandler, opts ...server.HandlerOption) error {
	type pantoCustomerService interface {
		PanById(ctx context.Context, req *PanRequest, rsp *PanResponse) error
	}
	type PantoCustomerService struct {
		pantoCustomerService
	}
	h := &pantoCustomerServiceHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "PantoCustomerService.PanById",
		Path:    []string{"/api/v1/messages"},
		Method:  []string{"POST"},
		Body:    "",
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&PantoCustomerService{h}, opts...))
}

type pantoCustomerServiceHandler struct {
	PantoCustomerServiceHandler
}

func (h *pantoCustomerServiceHandler) PanById(ctx context.Context, req *PanRequest, rsp *PanResponse) error {
	return h.PantoCustomerServiceHandler.PanById(ctx, req, rsp)
}
