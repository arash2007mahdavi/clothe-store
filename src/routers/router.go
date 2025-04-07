package routers

import (
	"store/src/handlers"
	"store/src/middlewares"

	"github.com/gin-gonic/gin"
)

func Store(router *gin.RouterGroup) {
	helper := handlers.Helper{}
	router.GET("/", helper.MainStore)
	
	profile := router.Group("/profile")
	{
		profile.POST("/new", helper.ProfileNew)
		profile.GET("/see", helper.ProfileSee)
		profile.GET("/see/all", middlewares.CheckAdmin, helper.ProfileSeeAll)
		profile.POST("/charge/wallet", helper.ProfileChargeWallet)
	}

	clothes := router.Group("/clothes")
	{
		clothes.GET("/", helper.GetClothes)
		clothes.POST("/buy", helper.BuyClothe)
	}
}