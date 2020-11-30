// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/comment.proto

package comment

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
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

// Api Endpoints for Comment service

func NewCommentEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Comment service

type CommentService interface {
	New(ctx context.Context, in *NewRequest, opts ...client.CallOption) (*NewResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
}

type commentService struct {
	c    client.Client
	name string
}

func NewCommentService(name string, c client.Client) CommentService {
	return &commentService{
		c:    c,
		name: name,
	}
}

func (c *commentService) New(ctx context.Context, in *NewRequest, opts ...client.CallOption) (*NewResponse, error) {
	req := c.c.NewRequest(c.name, "Comment.New", in)
	out := new(NewResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentService) List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "Comment.List", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Comment service

type CommentHandler interface {
	New(context.Context, *NewRequest, *NewResponse) error
	List(context.Context, *ListRequest, *ListResponse) error
}

func RegisterCommentHandler(s server.Server, hdlr CommentHandler, opts ...server.HandlerOption) error {
	type comment interface {
		New(ctx context.Context, in *NewRequest, out *NewResponse) error
		List(ctx context.Context, in *ListRequest, out *ListResponse) error
	}
	type Comment struct {
		comment
	}
	h := &commentHandler{hdlr}
	return s.Handle(s.NewHandler(&Comment{h}, opts...))
}

type commentHandler struct {
	CommentHandler
}

func (h *commentHandler) New(ctx context.Context, in *NewRequest, out *NewResponse) error {
	return h.CommentHandler.New(ctx, in, out)
}

func (h *commentHandler) List(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.CommentHandler.List(ctx, in, out)
}
