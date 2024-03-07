package api

import (
	"github.com/gin-gonic/gin"
	"go-backend/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

func Email(r *gin.Context) {
	cardCollection := repository.SelectCollection(repository.DB, "cards")
	c := repository.NewCardRepository(cardCollection)

	cardID := r.Query("cardID")
	_cardID, err := primitive.ObjectIDFromHex(cardID)
	log.Println("cardID: ", cardID)
	if err != nil {
		r.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	card, err := c.GetCard(r, _cardID)
	if err != nil {
		r.JSON(501, gin.H{
			"error": err.Error(),
		})
		return
	}

	userCollection := repository.SelectCollection(repository.DB, "users")
	u := repository.NewUserRepository(userCollection)

	email, err := u.GetEmail(r, card.UserID)
	if err != nil {
		r.JSON(502, gin.H{
			"error": err.Error(),
		})
		return
	}
	username, err := u.GetUsername(r, card.UserID)
	if err != nil {
		r.JSON(503, gin.H{
			"error": err.Error(),
		})
		return
	}

	questionCollection := repository.SelectCollection(repository.DB, "questions")
	q := repository.NewQuestionRepository(questionCollection)

	questionID := r.Query("questionID")
	_questionID, err := primitive.ObjectIDFromHex(questionID)
	if err != nil {
		r.JSON(504, gin.H{
			"error": err.Error(),
		})
		return
	}
	qa, err := q.GetQA(r, _questionID)
	if err != nil {
		r.JSON(505, gin.H{
			"error": err.Error(),
		})
		return
	}

	question := qa.Question
	answer := qa.Answer

	time.Sleep(30 * time.Second)

	comment := qa.CommentsContent[len(qa.Comments)-1]

	// 发送邮件

	repository.SendEmail(email, username, question, answer, comment)

	r.JSON(200, gin.H{
		"error": nil,
	})
}
