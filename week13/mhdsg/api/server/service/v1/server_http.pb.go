// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.3

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	res "mhdsg/pkg/res"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type ServerHTTPServer interface {
	GetServer(context.Context, *GetServerRequestViewModel) (*res.ResultObj, error)
}

func RegisterServerHTTPServer(s *http.Server, srv ServerHTTPServer) {
	r := s.Route("/")
	r.GET("/api/server/user", _Server_GetServer0_HTTP_Handler(srv))
}

func _Server_GetServer0_HTTP_Handler(srv ServerHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetServerRequestViewModel
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.server.service.v1.Server/GetServer")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetServer(ctx, req.(*GetServerRequestViewModel))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*res.ResultObj)
		return ctx.Result(200, reply)
	}
}

type ServerHTTPClient interface {
	GetServer(ctx context.Context, req *GetServerRequestViewModel, opts ...http.CallOption) (rsp *res.ResultObj, err error)
}

type ServerHTTPClientImpl struct {
	cc *http.Client
}

func NewServerHTTPClient(client *http.Client) ServerHTTPClient {
	return &ServerHTTPClientImpl{client}
}

func (c *ServerHTTPClientImpl) GetServer(ctx context.Context, in *GetServerRequestViewModel, opts ...http.CallOption) (*res.ResultObj, error) {
	var out res.ResultObj
	pattern := "/api/server/user"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.server.service.v1.Server/GetServer"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
