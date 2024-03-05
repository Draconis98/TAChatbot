package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Comment struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	QuestionID primitive.ObjectID `json:"question_id" bson:"question_id"`
	Content    string             `json:"content" bson:"content"`
	CommentBy  primitive.ObjectID `json:"comment_by" bson:"comment_by"`
}
