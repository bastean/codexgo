package health

import (
	"github.com/gin-gonic/gin"

	"github.com/bastean/codexgo/v4/internal/app/server/service/reply"
)

func Check(c *gin.Context) {
	reply.Success(c, "OK")
}
