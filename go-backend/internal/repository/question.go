package repository

import (
	"context"
	"go-backend/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type QuestionRepo interface {
	NewQuestionRepository(collection *mongo.Collection) *QuestionRepository
}

type QuestionRepository struct {
	collection *mongo.Collection
}

func NewQuestionRepository(collection *mongo.Collection) *QuestionRepository {
	return &QuestionRepository{collection}
}

func (r *QuestionRepository) GetQA(ctx context.Context, questionid primitive.ObjectID) (*model.QA, error) {
	var qa *model.QA
	err := r.collection.FindOne(ctx, bson.M{"_id": questionid}).Decode(&qa)
	if err != nil {
		return nil, err
	}

	return qa, nil
}
