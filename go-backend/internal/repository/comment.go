package repository

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-backend/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CommentRepo interface {
	NewCommentRepository(collection *mongo.Collection) *CommentRepository
	NewComment(c *gin.Context) (model.Comment, error)
	GetComment(c *gin.Context) (*model.Comment, error)
}

type CommentRepository struct {
	collection *mongo.Collection
}

func NewCommentRepository(collection *mongo.Collection) *CommentRepository {
	return &CommentRepository{collection}
}

func (r *CommentRepository) NewComment(ctx context.Context, comment *model.Comment) (model.Comment, error) {
	result, err := r.collection.InsertOne(ctx, comment)
	if err != nil {
		return model.Comment{}, err
	}
	comment.ID = result.InsertedID.(primitive.ObjectID)
	return *comment, nil
}

func (r *CommentRepository) GetComment(ctx context.Context, commentID primitive.ObjectID) (*model.Comment, error) {
	var comment *model.Comment
	err := r.collection.FindOne(ctx, bson.M{"_id": commentID}).Decode(&comment)
	if err != nil {
		return nil, err
	}
	return comment, nil
}
