package server

import (
	"github.com/bastean/codexgo/backend/internal/api/v1/controllers/auth"
	"github.com/bastean/codexgo/backend/internal/api/v1/controllers/public"
	"github.com/bastean/codexgo/backend/internal/api/v1/middleware"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func loadMiddlewares() {
	router.Use(gin.CustomRecovery(middleware.ErrorHandler))
}

func loadEndpoints() {
	v1 := router.Group("/v1")

	v1.GET("/status", public.Status())

	user := v1.Group("/user")

	user.PUT("/", public.UserPut())
	user.POST("/", public.UserPost())

	user.PATCH("/", auth.UserPatch())
	user.DELETE("/", auth.UserDelete())

}

func Run() {
	loadMiddlewares()
	loadEndpoints()

	router.Run()
}
