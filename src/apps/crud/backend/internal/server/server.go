package server

import (
	"github.com/bastean/codexgo/backend/internal/api/v1/controllers/auth"
	"github.com/bastean/codexgo/backend/internal/api/v1/controllers/public"
	"github.com/bastean/codexgo/backend/internal/api/v1/middleware"
	"github.com/bastean/codexgo/backend/internal/container"
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

	userAuth := user.Group("/auth", middleware.VerifyAuthentication())

	userAuth.PATCH("/", auth.UserPatch())
	userAuth.DELETE("/", auth.UserDelete())
}

func Run() {
	container.Logger.Info("starting server")

	loadMiddlewares()
	loadEndpoints()

	container.Logger.Info("server started!")

	router.Run()
}
