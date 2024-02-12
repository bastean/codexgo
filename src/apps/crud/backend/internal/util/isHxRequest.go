package util

import (
	"github.com/gin-gonic/gin"
)

func IsHxRequest(c *gin.Context) bool {
	return c.Request.Header["Hx-Request"][0] != ""
}
