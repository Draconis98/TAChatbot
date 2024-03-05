package api

import (
	"github.com/gin-gonic/gin"
	"go-backend/internal/model"
	"go-backend/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

func NewComment(r *gin.Context) {
	commentCollection := repository.SelectCollection(repository.DB, "comments")
	c := repository.NewCommentRepository(commentCollection)

	questionCollection := repository.SelectCollection(repository.DB, "questions")
	q := repository.NewQuestionRepository(questionCollection)

	userID := r.Query("userID")
	questionID := r.Query("questionID")

	log.Println("userID: ", userID, " is creating a new comment for questionID: ", questionID)

	_userID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	_questionID, err := primitive.ObjectIDFromHex(questionID)
	if err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	newComment := &model.Comment{
		QuestionID: _questionID,
		Content:    r.Query("comment"),
		CommentBy:  _userID,
	}

	comment, err := c.NewComment(r, newComment)
	if err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	log.Println("commentID: ", comment.ID, " is created")

	err = q.AddComment(r, _questionID, comment.ID)
	if err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	r.JSON(http.StatusOK, gin.H{
		"error": nil,
	})
}

func GetComment(r *gin.Context) {
	commentCollection := repository.SelectCollection(repository.DB, "comments")
	c := repository.NewCommentRepository(commentCollection)

	commentID := r.Query("commentID")
	_commentID, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	comment, err := c.GetComment(r, _commentID)
	if err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	r.JSON(http.StatusOK, gin.H{
		"error":       nil,
		"id":          comment.ID,
		"question_id": comment.QuestionID,
		"content":     comment.Content,
		"comment_by":  comment.CommentBy,
	})
}
