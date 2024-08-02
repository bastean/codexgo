package middleware

import (
	"net/http"

	"github.com/bastean/codexgo/v4/internal/app/server/util/reply"
	"github.com/bastean/codexgo/v4/internal/pkg/service/logger/log"
	"github.com/gin-gonic/gin"
)

func Recover(c *gin.Context, err any) {
	log.Error(err.(error).Error())
	c.AbortWithStatusJSON(http.StatusInternalServerError, &reply.JSON{Message: "Server error. Try again later."})
}
