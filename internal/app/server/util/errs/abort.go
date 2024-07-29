package errs

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Abort(c *gin.Context, err error) {
	c.Error(err)
	c.Abort()
}

func AbortWithRedirect(c *gin.Context, err error, route string) {
	Abort(c, err)
	c.Redirect(http.StatusFound, route)
}
