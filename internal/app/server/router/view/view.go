package view

import (
	"github.com/gin-gonic/gin"

	"github.com/bastean/codexgo/v4/internal/app/server/handler/page"
	"github.com/bastean/codexgo/v4/internal/app/server/middleware"
)

type View struct {
	*gin.Engine
}

func (v *View) Public() {
	public := v.Group("/")

	home := public.Group("/")

	home.GET("/", page.Home)

	home.GET("/reset", page.Home)
}

func (v *View) Private() {
	private := v.Group("/", middleware.Authentication)

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
