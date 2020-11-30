package handler

import (
	"context"
	"time"

	pb "github.com/abvarun226/comment/proto"
	"github.com/google/uuid"
)

// Comment struct.
type Comment struct{}

// NewComment creates a new comment handler.
func NewComment() *Comment {
	return &Comment{}
}

// New handler to create a new comment.
func (c *Comment) New(ctx context.Context, req *pb.NewRequest, rsp *pb.NewResponse) error {
	rsp.Comment = &pb.Item{
		Id:      uuid.New().String(),
		Post:    req.Post,
		Author:  req.Author,
		Message: req.Message,
		Created: time.Now().UTC().Unix(),
	}
	return nil
}

// List handler to list comments in a post.
func (c *Comment) List(ctx context.Context, req *pb.ListRequest, rsp *pb.ListResponse) error {
	rsp.Comments = []*pb.Item{
		{
			Id:      uuid.New().String(),
			Post:    req.Post,
			Author:  "greycell",
			Message: "this is a comment",
			Created: 1604650806,
		},
	}
	return nil
}
