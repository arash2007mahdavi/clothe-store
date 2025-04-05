package routers

import (
	"store/src/handlers"

	"github.com/gin-gonic/gin"
)

func Store(router *gin.RouterGroup) {
	helper := handlers.Helper{}
	router.GET("/", helper.MainStore)
	router.GET("/clothes", helper.GetClothes)
	profile := router.Group("/profile")
	{
		profile.POST("/new", helper.ProfileNew)
		profile.GET("/see", helper.ProfileSee)
		profile.POST("/charge/wallet", helper.ProfileChargeWallet)
	}
}