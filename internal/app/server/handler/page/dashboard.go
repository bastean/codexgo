package page

import (
	"github.com/gin-gonic/gin"

	"github.com/bastean/codexgo/v4/internal/app/server/component/page/dashboard"
	"github.com/bastean/codexgo/v4/internal/app/server/service/errs"
	"github.com/bastean/codexgo/v4/internal/app/server/service/format"
	"github.com/bastean/codexgo/v4/internal/app/server/service/key"
	"github.com/bastean/codexgo/v4/internal/pkg/adapter/query"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/read"
)

func Dashboard(c *gin.Context) {
	id, exists := c.Get(key.UserID)

	if !exists {
		errs.AbortByErrWithRedirect(c, errs.MissingKey(key.UserID, "Dashboard"), "/")
		return
	}

	attributes := new(read.QueryAttributes)

	attributes.ID = format.String(id)

	response, err := query.Bus.Ask(messages.New(
		read.QueryKey,
		attributes,
		new(read.QueryMeta),
	))

	if err != nil {
		errs.AbortByErrWithRedirect(c, errors.BubbleUp(err, "Dashboard"), "/")
		return
	}

	found, ok := response.Attributes.(*read.ResponseAttributes)

	if !ok {
		errs.AbortByErrWithRedirect(c, errs.Assertion("Dashboard"), "/")
		return
	}

	err = dashboard.Page(found).Render(c.Request.Context(), c.Writer)

	if err != nil {
		errs.AbortByErr(c, errs.Render(err, "Dashboard"))
	}
}
