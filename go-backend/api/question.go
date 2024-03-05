package api

import (
	"github.com/gin-gonic/gin"
	"go-backend/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

func NewQuestion(c *gin.Context) {
	userID := c.Query("userID")
	cardID := c.Query("cardID")
	log.Println("userID: ", userID, " is creating a new question for cardID: ", cardID)
}

func GetQuestion(r *gin.Context) {
	questionCollection := repository.SelectCollection(repository.DB, "questions")
	q := repository.NewQuestionRepository(questionCollection)

	questionID := r.Query("questionID")
	_questionID, err := primitive.ObjectIDFromHex(questionID)
	if err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{
			"error":            err.Error(),
			"id":               nil,
			"question":         nil,
			"answer":           nil,
			"likes":            nil,
			"create_at":        nil,
			"comments":         nil,
			"comments_content": nil,
		})
		return
	}

	qa, err := q.GetQA(r, _questionID)
	if err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{
			"error":            err.Error(),
			"id":               nil,
			"question":         nil,
			"answer":           nil,
			"likes":            nil,
			"create_at":        nil,
			"comments":         nil,
			"comments_content": nil,
		})
		return
	}

	r.JSON(http.StatusOK, gin.H{
		"error":            nil,
		"id":               qa.ID,
		"question":         qa.Question,
		"answer":           qa.Answer,
		"likes":            qa.Like,
		"create_at":        qa.CreateAt,
		"comments":         qa.Comments,
		"comments_content": qa.CommentsContent,
	})
}
