package repository

import (
	"context"
	"film-api/utils"
	"log"
	"time"

	c "github.com/dbgjerez/workshop-golang-grpc/comment/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	//defer conn.Close()
	client := c.NewCommentServiceClient(conn)
	return &CommentRepository{
		client: client,
	}
}

func (repo *CommentRepository) FindComments(idObject int32) *c.Comments {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	comments, err := repo.client.Retrieve(ctx, &c.RetrieveRequest{IdObject: idObject, TypeObject: COMMENT_OBJECT_TYPE})
	if err != nil {
		log.Printf("Error calling comment service for %d film: %v", idObject, err)
		return &c.Comments{}
	}
	return comments
}
