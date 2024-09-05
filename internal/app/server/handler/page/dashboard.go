package page

import (
	"github.com/bastean/codexgo/v4/internal/app/server/component/page/dashboard"
	"github.com/bastean/codexgo/v4/internal/app/server/util/errs"
	"github.com/bastean/codexgo/v4/internal/app/server/util/format"
	"github.com/bastean/codexgo/v4/internal/app/server/util/key"
	"github.com/bastean/codexgo/v4/internal/pkg/service/errors"
	"github.com/bastean/codexgo/v4/internal/pkg/service/module/user"
	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {
	id, exists := c.Get(key.UserId)

	if !exists {
		errs.AbortErrWithRedirect(c, errs.MissingKey(key.UserId, "Dashboard"), "/")
		return
	}

	query := new(user.ReadQuery)

	query.Id = format.ToString(id)

	found, err := user.Read.Handle(query)

	if err != nil {
		errs.AbortErrWithRedirect(c, errors.BubbleUp(err, "Dashboard"), "/")
		return
	}

	err = dashboard.Page(found).Render(c.Request.Context(), c.Writer)

	if err != nil {
		errs.AbortErr(c, errs.Render(err, "Dashboard"))
	}
}
