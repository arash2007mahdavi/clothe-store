package routers

import (
	"store/src/handlers"

	"github.com/gin-gonic/gin"
)

func Store(router *gin.RouterGroup) {
	helper := handlers.Helper{}
	router.GET("/", helper.MainStore)
}