package page

import (
	"github.com/gin-gonic/gin"

	"github.com/bastean/codexgo/v4/internal/app/server/component/page/dashboard"
	"github.com/bastean/codexgo/v4/internal/app/server/service/errs"
	"github.com/bastean/codexgo/v4/internal/app/server/service/format"
	"github.com/bastean/codexgo/v4/internal/app/server/service/key"
	"github.com/bastean/codexgo/v4/internal/pkg/service/communication/query"
	"github.com/bastean/codexgo/v4/internal/pkg/service/errors"
	"github.com/bastean/codexgo/v4/internal/pkg/service/module/user"
)

func Dashboard(c *gin.Context) {
	id, exists := c.Get(key.UserID)

	if !exists {
		errs.AbortByErrWithRedirect(c, errs.MissingKey(key.UserID, "Dashboard"), "/")
		return
	}

	attributes := new(user.ReadQueryAttributes)

	attributes.ID = format.ToString(id)

	response, err := query.Bus.Ask(query.New(
		user.ReadQueryKey,
		attributes,
		new(user.ReadQueryMeta),
	))

	if err != nil {
		errs.AbortByErrWithRedirect(c, errors.BubbleUp(err, "Dashboard"), "/")
		return
	}

	found, ok := response.Attributes.(*user.ReadResponseAttributes)

	if !ok {
		errs.AbortByErrWithRedirect(c, errs.Assertion("Dashboard"), "/")
		return
	}

	err = dashboard.Page(found).Render(c.Request.Context(), c.Writer)

	if err != nil {
		errs.AbortByErr(c, errs.Render(err, "Dashboard"))
	}
}
