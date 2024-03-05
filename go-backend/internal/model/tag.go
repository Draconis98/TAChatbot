package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Tag is a model for tag
type Tag struct {
	ID    primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Name  string               `json:"name" bson:"name"`
	Cards []primitive.ObjectID `json:"cards" bson:"cards"`
}
