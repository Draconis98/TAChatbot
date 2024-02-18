package repository

import (
	"context"
	"go-backend/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserRepo is a repository interface for user
type UserRepo interface {
	NewUserRepository(collection *mongo.Collection) *UserRepository
	NewUser(user *model.User) (*model.User, error)
	GetUser(username, email string) (*model.User, error)
	GetUsername(id primitive.ObjectID) (string, error)
	CheckUserExistence(id primitive.ObjectID) (bool, error)
}

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) *UserRepository {
	return &UserRepository{collection}
}

func (r *UserRepository) NewUser(ctx context.Context, user *model.User) (model.User, error) {
	result, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return model.User{}, err
	}
	user.ID = result.InsertedID.(primitive.ObjectID)
	return *user, nil
}

func (r *UserRepository) GetUser(ctx context.Context, username, email string) (*model.User, error) {
	var user model.User
	err := r.collection.FindOne(ctx, bson.M{"username": username, "email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUsername(ctx context.Context, id primitive.ObjectID) (string, error) {
	var user model.User
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return "", err
	}
	return user.Username, nil
}

func (r *UserRepository) CheckUserExistence(ctx context.Context, id primitive.ObjectID) (bool, error) {
	var user model.User
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return false, nil
	}
	return true, nil
}

func (r *UserRepository) AddCard(ctx context.Context, userID, cardID primitive.ObjectID) error {
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": userID}, bson.M{"$push": bson.M{"cards": cardID}})
	if err != nil {
		return err
	}
	return nil
}
