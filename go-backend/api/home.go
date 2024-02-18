package api

import (
	"github.com/gin-gonic/gin"
)

func Home(r *gin.Context) {
	r.JSON(200, gin.H{
		"message": "Welcome to the home page!",
	})
}
