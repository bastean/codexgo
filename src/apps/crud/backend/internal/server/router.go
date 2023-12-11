package server

import (
	"github.com/bastean/codexgo/internal/api/v1/controllers/auth"
	"github.com/bastean/codexgo/internal/api/v1/controllers/public"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

var v1 = router.Group("/v1")

var user = v1.Group("/user")

func loadEndpoints() {
	v1.GET("/status", public.Status())

	user.PUT("/", public.UserPut())
	user.POST("/", public.UserPost())

	user.PATCH("/", auth.UserPatch())
	user.DELETE("/", auth.UserDelete())
}

func Run() {
	loadEndpoints()
	router.Run()
}
