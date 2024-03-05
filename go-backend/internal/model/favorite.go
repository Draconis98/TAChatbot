package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type FavoriteUser struct {
	UserID primitive.ObjectID   `json:"user_id" bson:"user_id"`
	CardID []primitive.ObjectID `json:"card_id" bson:"card_id"`
}

type FavoriteCard struct {
	CardID         primitive.ObjectID   `json:"card_id" bson:"card_id"`
	UserID         []primitive.ObjectID `json:"user_id" bson:"user_id"`
	FavoritesCount int                  `json:"favoritesCount" bson:"favoritesCount"`
}
