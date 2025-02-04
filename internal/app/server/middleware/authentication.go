package middleware

import (
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/bastean/codexgo/v4/internal/app/server/service/errs"
	"github.com/bastean/codexgo/v4/internal/app/server/service/format"
	"github.com/bastean/codexgo/v4/internal/app/server/service/key"
	"github.com/bastean/codexgo/v4/internal/pkg/service/authentication"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
)

func Authentication(c *gin.Context) {
	session := sessions.Default(c)

	token := session.Get(key.Authorization)

	if token == nil {
		errs.AbortWithRedirect(c, "/")
		return
	}

	signature := strings.Split(format.String(token), " ")[1]

	claims, err := authentication.JWT.Validate(signature)

	if err != nil {
		errs.AbortByErrWithRedirect(c, errors.BubbleUp(err, "Authentication"), "/")
		return
	}

	value, exists := claims[key.UserID]

	if !exists {
		errs.AbortByErrWithRedirect(c, errs.MissingKey(key.UserID, "Authentication"), "/")
		return
	}

	c.Set(key.UserID, value)

	c.Next()
}
