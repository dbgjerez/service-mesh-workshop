package repository

import (
	"context"
	"film-api/utils"
	"log"
	"time"

	c "github.com/dbgjerez/workshop-golang-grpc/comment/grpc"
	"google.golang.org/grpc"
)

const (
	COMMENT_SERVICE_URL string = "COMMENT_SERVICE_URL"
	COMMENT_OBJECT_TYPE        = "film"
)

type CommentRepository struct {
	client c.CommentServiceClient
}

func NewCommentRepository() *CommentRepository {
	url := utils.GetEnv(COMMENT_SERVICE_URL, "")
	if url == "" {
		log.Fatalf("env var COMMENT_SERVICE_URL not defined")
	}
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client := c.NewCommentServiceClient(conn)
	return &CommentRepository{
		client: client,
	}
}

func (repo *CommentRepository) FindComments(ids []int32) (*c.Comments, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	defer cancel()
	stream, err := repo.client.Retrieve(ctx)
	if err != nil {
		log.Printf("error creating the stream %v:", err)
		return &c.Comments{}, err
	}
	for _, id := range ids {
		if err := stream.Send(&c.RetrieveRequest{IdObject: id, TypeObject: COMMENT_OBJECT_TYPE}); err != nil {
			return &c.Comments{}, err
		}
	}
	return stream.CloseAndRecv()
}
