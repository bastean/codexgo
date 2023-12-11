package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Delete struct {
	Password string `json:"password" binding:"required"`
}

func UserDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user Delete

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}
