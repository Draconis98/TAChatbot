package api

import (
	"github.com/gin-gonic/gin"
	"go-backend/internal/model"
	"go-backend/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"time"
)

func NewCard(r *gin.Context) {
	userCollection := repository.SelectCollection(repository.DB, "users")
	u := repository.NewUserRepository(userCollection)

	cardCollection := repository.SelectCollection(repository.DB, "cards")
	c := repository.NewCardRepository(cardCollection)

	userID := r.Query("userID")
	log.Println("userID: ", userID, " is creating a new card")

	_userID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	newCard := &model.Card{
		UserID:         _userID,
		Title:          "",
		Tags:           []primitive.ObjectID{},
		Content:        []model.QA{},
		Summary:        "",
		FavoritesCount: 0,
		CreateAt:       time.Now(),
	}

	card, err := c.NewCard(r, newCard)
	if err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	log.Println("cardID: ", card.ID, " is created")

	err = u.AddCard(r, _userID, card.ID)
	if err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	r.JSON(http.StatusOK, gin.H{
		"cardID": card.ID,
	})
}

func Show(r *gin.Context) {
	cardCollection := repository.SelectCollection(repository.DB, "cards")
	c := repository.NewCardRepository(cardCollection)

	cards, err := c.SortCard(r, "")
	if err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	r.JSON(http.StatusOK, gin.H{
		"card_list": cards,
	})
}

func ShowLatest(r *gin.Context) {
	cardCollection := repository.SelectCollection(repository.DB, "cards")
	c := repository.NewCardRepository(cardCollection)

	cards, err := c.SortCard(r, "latest")
	if err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	r.JSON(http.StatusOK, gin.H{
		"card_list": cards,
	})
}

func ShowHottest(r *gin.Context) {
	cardCollection := repository.SelectCollection(repository.DB, "cards")
	c := repository.NewCardRepository(cardCollection)

	cards, err := c.SortCard(r, "hottest")
	if err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	r.JSON(http.StatusOK, gin.H{
		"card_list": cards,
	})
}

func GetCard(r *gin.Context) {
	cardCollection := repository.SelectCollection(repository.DB, "cards")
	c := repository.NewCardRepository(cardCollection)

	cardID := r.Query("cardID")
	_cardID, err := primitive.ObjectIDFromHex(cardID)
	if err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"card":  nil,
		})
		return
	}

	card, err := c.GetCard(r, _cardID)
	if err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"card":  nil,
		})
		return
	}

	r.JSON(http.StatusOK, gin.H{
		"error": nil,
		"card":  card,
	})
}
