package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Username  string               `json:"username" bson:"username"`
	Role      string               `json:"role" bson:"role"`
	Email     string               `json:"email" bson:"email"`
	Cards     []primitive.ObjectID `json:"cards" bson:"cards"`
	Favorites []primitive.ObjectID `json:"favorites" bson:"favorites"`
}
