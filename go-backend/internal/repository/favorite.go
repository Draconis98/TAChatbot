package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FavoriteRepo interface {
	NewFavoriteRepository(collection *mongo.Collection) *FavoriteRepository
	UserFavoriteACard(userID primitive.ObjectID, cardID primitive.ObjectID) error
	UserUnfavoriteACard(userID primitive.ObjectID, cardID primitive.ObjectID) error
}

type FavoriteRepository struct {
	collection *mongo.Collection
}

func (r *FavoriteRepository) NewFavoriteRepository(collection *mongo.Collection) *FavoriteRepository {
	return &FavoriteRepository{collection}
}

// User perspective
func (r *FavoriteRepository) UserFavoriteACard(ctx context.Context, userID primitive.ObjectID, cardID primitive.ObjectID) error {
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": userID}, bson.M{"$push": bson.M{"favorites": cardID}})
	if err != nil {
		return err
	}
	return nil
}

func (r *FavoriteRepository) UserUnfavoriteACard(ctx context.Context, userID primitive.ObjectID, cardID primitive.ObjectID) error {
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": userID}, bson.M{"$pull": bson.M{"favorites": cardID}})
	if err != nil {
		return err
	}
	return nil
}

func (r *FavoriteRepository) GetUserFavorites(ctx context.Context, userID primitive.ObjectID) ([]primitive.ObjectID, error) {
	var user struct {
		Favorites []primitive.ObjectID `bson:"favorites"`
	}
	err := r.collection.FindOne(ctx, bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user.Favorites, nil
}

// Card perspective
func (r *FavoriteRepository) UpdateCardFavoritesCount(ctx context.Context, cardID primitive.ObjectID, count int) error {
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": cardID}, bson.M{"$set": bson.M{"favoritesCount": count}})
	if err != nil {
		return err
	}
	return nil
}

func (r *FavoriteRepository) GetCardFavoritesCount(ctx context.Context, cardID primitive.ObjectID) (int, error) {
	var card struct {
		FavoritesCount int `bson:"favoritesCount"`
	}
	err := r.collection.FindOne(ctx, bson.M{"_id": cardID}).Decode(&card)
	if err != nil {
		return 0, err
	}
	return card.FavoritesCount, nil
}
