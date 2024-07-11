package middleware

import (
	"net/http"

	"github.com/bastean/codexgo/internal/app/server/util/reply"
	"github.com/bastean/codexgo/internal/pkg/service/logger/log"
	"github.com/gin-gonic/gin"
)

func PanicHandler(c *gin.Context, err any) {
	log.Error(err.(error).Error())
	c.AbortWithStatusJSON(http.StatusInternalServerError, reply.JSON(false, "internal server error", reply.Payload{}))
}
