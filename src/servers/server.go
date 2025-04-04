package servers

import (
	"fmt"
	"store/src/configs"
	"store/src/middlewares"
	"store/src/routers"

	"github.com/gin-gonic/gin"
)

func NewServer(cfg configs.Config) {
	engine := gin.New()
	engine.Use(gin.Recovery(), gin.Logger(), middlewares.CheckApiKey)

	store := engine.Group("/store")
	{
		routers.Store(store)
	}

	engine.Run(fmt.Sprintf(":%v", cfg.Server.Port))
}