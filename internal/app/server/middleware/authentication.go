package middleware

import (
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/bastean/codexgo/v4/internal/app/server/util/errs"
	"github.com/bastean/codexgo/v4/internal/app/server/util/format"
	"github.com/bastean/codexgo/v4/internal/app/server/util/key"
	"github.com/bastean/codexgo/v4/internal/pkg/service/authentication/jwt"
	"github.com/bastean/codexgo/v4/internal/pkg/service/errors"
)

func Authentication(c *gin.Context) {
	session := sessions.Default(c)

	token := session.Get(key.Authorization)

	if token == nil {
		errs.AbortWithRedirect(c, "/")
		return
	}

	signature := strings.Split(format.ToString(token), " ")[1]

	claims, err := jwt.Validate(signature)

	if err != nil {
		errs.AbortErrWithRedirect(c, errors.BubbleUp(err, "Authentication"), "/")
		return
	}

	value, exists := claims[key.UserId]

	if !exists {
		errs.AbortErrWithRedirect(c, errs.MissingKey(key.UserId, "Authentication"), "/")
		return
	}

	c.Set(key.UserId, value)

	c.Next()
}
