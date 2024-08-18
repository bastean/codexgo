package view

import (
	"github.com/bastean/codexgo/v4/internal/app/server/handler/page"
	"github.com/bastean/codexgo/v4/internal/app/server/middleware"
	"github.com/gin-gonic/gin"
)

type View struct {
	*gin.Engine
}

func (view *View) Public() {
	public := view.Group("/")

	home := public.Group("/")

	home.GET("/", page.Home)
}

func (view *View) Private() {
	private := view.Group("/", middleware.Authentication)

	dashboard := private.Group("/dashboard")

	dashboard.GET("/", page.Dashboard)
}

func Use(router *gin.Engine) {
	view := &View{
		Engine: router,
	}

	view.Public()

	view.Private()
}
