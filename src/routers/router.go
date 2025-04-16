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
		profile.POST("/new", middlewares.CheckApiKey, helper.ProfileNew)
		profile.GET("/see", middlewares.CheckApiKey, helper.ProfileSee)
		profile.GET("/see/all", middlewares.CheckApiKey, middlewares.CheckAdmin, helper.ProfileSeeAll)
		profile.POST("/charge/wallet", middlewares.CheckApiKey, helper.ProfileChargeWallet)
	}

	clothes := router.Group("/clothes")
	{
		clothes.GET("/", middlewares.CheckApiKey, helper.GetClothes)
		clothes.POST("/buy", middlewares.CheckApiKey, helper.BuyClothe)
	}
}