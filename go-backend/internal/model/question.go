package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type QA struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Question string             `json:"question" bson:"question"`
	Answer   string             `json:"answer" bson:"answer"`
	Like     int                `json:"like" bson:"like"` // 初始化为0， 1表示喜欢， 2表示不喜欢
	CreateAt string             `json:"create_at" bson:"create_at"`
}
