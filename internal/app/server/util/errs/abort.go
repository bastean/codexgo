package errs

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AbortErr(c *gin.Context, err error) {
	c.Error(err)
	c.Abort()
}

func AbortErrWithRedirect(c *gin.Context, err error, route string) {
	AbortErr(c, err)
	c.Redirect(http.StatusFound, route)
}

func AbortWithRedirect(c *gin.Context, route string) {
	c.Abort()
	c.Redirect(http.StatusFound, route)
}
