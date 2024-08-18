package redirect

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Default(c *gin.Context) {
	c.Redirect(http.StatusFound, "/")
}
