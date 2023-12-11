package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Patch struct {
	Id       string `json:"id" binding:"required"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func UserPatch() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user Patch

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}
