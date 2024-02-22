package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Card is a model for card
type Card struct {
	ID             primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	UserID         primitive.ObjectID   `json:"userID" bson:"userID"`
	Title          string               `json:"title" bson:"title"`
	Tags           []primitive.ObjectID `json:"tags" bson:"tags"`
	Content        []string             `json:"content" bson:"content"`
	Questions      []primitive.ObjectID `json:"questions" bson:"questions"`
	FavoritesCount int                  `json:"favoritesCount" bson:"favoritesCount"`
	Summary        string               `json:"summary" bson:"summary"`
	CreateAt       string               `json:"create_at" bson:"create_at"`
}
