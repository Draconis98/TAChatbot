package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-backend/internal/model"
	"go-backend/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

func CheckUserExistence(c *gin.Context) {
	userCollection := repository.SelectCollection(repository.DB, "users")
	u := repository.NewUserRepository(userCollection)

	username := c.Query("username")
	log.Println(username)
	email := c.Query("email")
	log.Println(email)

	// 从数据库中获取用户信息
	user, err := u.GetUser(context.Background(), username, email)
	if err != nil { // 数据库中没有该用户，插入新的用户信息
		newUser := &model.User{
			Username:  username,
			Role:      "student",
			Email:     email,
			Cards:     []primitive.ObjectID{},
			Favorites: []primitive.ObjectID{},
		}
		user, err := u.NewUser(context.Background(), newUser)
		if err != nil { // 插入新用户信息发生错误，返回错误信息
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":  err.Error(),
				"exist":  false,
				"userID": "",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{ // 插入新用户信息成功，返回用户ID
			"error":  nil,
			"exist":  false,
			"userID": user.ID,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{ // 数据库中有该用户，返回用户ID
		"error":  nil,
		"exist":  true,
		"userID": user.ID,
	})
}

func GetUsername(c *gin.Context) {
	userCollection := repository.SelectCollection(repository.DB, "users")
	r := repository.NewUserRepository(userCollection)

	userID := c.Query("userID")
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil { // 用户ID转换失败，是服务器内部错误
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":    err.Error(),
			"username": "",
		})
		return
	}

	username, err := r.GetUsername(context.Background(), id)
	if err != nil { // 没有这个用户ID，属于未授权
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":    err.Error(),
			"username": "",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error":    nil,
		"username": username,
	})
}

func GetUserRole(c *gin.Context) {
	userCollection := repository.SelectCollection(repository.DB, "users")
	r := repository.NewUserRepository(userCollection)

	userID := c.Query("userID")
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil { // 用户ID转换失败，是服务器内部错误
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"role":  "",
		})
		return
	}

	role, err := r.GetUserRole(context.Background(), id)
	if err != nil { // 没有这个用户ID，属于未授权
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
			"role":  "",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error": nil,
		"role":  role,
	})
}
