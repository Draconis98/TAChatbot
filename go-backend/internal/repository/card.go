package repository

import (
	"context"
	"go-backend/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type CardRepo interface {
	NewCardRepository(collection *mongo.Collection) *CardRepository
	NewCard(card *model.Card) (*model.Card, error)
	GetCard(id primitive.ObjectID) (*model.Card, error)
}

type CardRepository struct {
	collection *mongo.Collection
}

func NewCardRepository(collection *mongo.Collection) *CardRepository {
	return &CardRepository{collection}
}

func (r *CardRepository) NewCard(ctx context.Context, card *model.Card) (*model.Card, error) {
	result, err := r.collection.InsertOne(ctx, card)
	if err != nil {
		return nil, err
	}
	card.ID = result.InsertedID.(primitive.ObjectID)
	return card, nil
}

func (r *CardRepository) GetCard(ctx context.Context, cardid primitive.ObjectID) (*model.Card, error) {
	var card *model.Card
	err := r.collection.FindOne(ctx, bson.M{"_id": cardid}).Decode(&card)
	if err != nil {
		return nil, err
	}

	return card, nil
}

func (r *CardRepository) CheckCardExistence(ctx context.Context, userid, cardid primitive.ObjectID) (bool, error) {
	var card *model.Card
	err := r.collection.FindOne(ctx, bson.M{"userID": userid, "_id": cardid}).Decode(&card)
	if err != nil {
		return false, nil
	}
	return true, nil
}

func (r *CardRepository) SortCard(ctx context.Context, method string) ([]model.Card, error) {
	var cards []model.Card

	// 获得指向collection中的指针
	cur, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	// 遍历指针获得具体的每一项
	for cur.Next(ctx) {
		var card model.Card
		err = cur.Decode(&card)
		if err != nil {
			return nil, err
		}
		log.Println(card)
		if method == "latest" { // 如果是要求按最新的顺序排列，则从切片前向插入
			cards = append([]model.Card{card}, cards...)
		} else if method == "hottest" {
			// TODO: 按照热度排序
		} else { // 否则从切片尾部插入
			cards = append(cards, card)
		}

	}
	log.Println(cards)
	return cards, nil
}
