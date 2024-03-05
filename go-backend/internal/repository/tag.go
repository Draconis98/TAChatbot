package repository

import (
	"context"
	"go-backend/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TagRepo interface {
	NewTagRepository(collection *mongo.Collection) *TagRepository
	NewTag(tag *model.Tag) (*model.Tag, error)
	GetTag(id primitive.ObjectID) (*model.Tag, error)
	UpdateTagCards(tag *model.Tag, cardID primitive.ObjectID) (*model.Tag, error)
}

type TagRepository struct {
	collection *mongo.Collection
}

func NewTagRepository(collection *mongo.Collection) *TagRepository {
	return &TagRepository{collection}
}

func (r *TagRepository) NewTag(ctx context.Context, tag *model.Tag) (*model.Tag, error) {
	result, err := r.collection.InsertOne(ctx, tag)
	if err != nil {
		return nil, err
	}
	tag.ID = result.InsertedID.(primitive.ObjectID)
	return tag, nil
}

func (r *TagRepository) GetTag(ctx context.Context, id primitive.ObjectID) (*model.Tag, error) {
	var tag *model.Tag
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&tag)
	if err != nil {
		return nil, err
	}
	return tag, nil
}

func (r *TagRepository) UpdateTagCards(ctx context.Context, tag *model.Tag, cardID []primitive.ObjectID) (*model.Tag, error) {
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": tag.ID}, bson.M{"$set": bson.M{"cards": cardID}})
	if err != nil {
		return nil, err
	}
	tag.Cards = cardID
	return tag, nil
}
