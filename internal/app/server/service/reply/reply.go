package reply

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/array"
)

type Response struct {
	Success bool
	Message string
	Data    any
}

type Error struct {
	Type    string
	Message string
	Data    any
}

func reply(c *gin.Context, status int, response *Response, payload ...any) {
	data, ok := array.Slice(payload, 0)

	if ok {
		response.Data = data
	}

	c.JSON(status, response)
}

func Success(c *gin.Context, message string, data ...any) {
	reply(c, http.StatusOK, &Response{
		Success: true,
		Message: message,
	}, data...)
}

func failure(c *gin.Context, status int, data ...any) {
	reply(c, status, &Response{
		Message: "Some errors have been found.",
	}, data...)
}

func FailureClient(c *gin.Context, data ...any) {
	failure(c, http.StatusBadRequest, data...)
}

func FailureServer(c *gin.Context) {
	failure(c, http.StatusInternalServerError, []*Error{{Type: "Internal", Message: "Server error. Try again later."}})
}
