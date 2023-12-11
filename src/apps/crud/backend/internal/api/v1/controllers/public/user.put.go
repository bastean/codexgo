package public

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Put struct {
	Id       string `json:"id" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func UserPut() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user Put

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}
