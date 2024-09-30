package errs

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/bastean/codexgo/v4/internal/app/server/service/key"
)

func AbortByErr(c *gin.Context, err error) {
	c.Error(err)
	c.Abort()
}

func AbortByErrWithRedirect(c *gin.Context, err error, route string) {
	AbortByErr(c, err)
	c.Set(key.Redirect, route)
}

func AbortWithRedirect(c *gin.Context, route string) {
	c.Abort()
	c.Redirect(http.StatusFound, route)
}
