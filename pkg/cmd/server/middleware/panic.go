package middleware

import (
	"net/http"

	"github.com/bastean/codexgo/pkg/cmd/server/service/logger"
	"github.com/bastean/codexgo/pkg/cmd/server/util/reply"
	"github.com/gin-gonic/gin"
)

func PanicHandler(c *gin.Context, err any) {
	logger.Error(err.(error).Error())
	c.AbortWithStatusJSON(http.StatusInternalServerError, reply.JSON(false, "internal server error", reply.Payload{}))
}
