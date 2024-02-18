package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type QA struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Question string             `json:"question" bson:"question"`
	Answer   string             `json:"answer" bson:"answer"`
}
