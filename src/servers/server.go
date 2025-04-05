package servers

import (
	"fmt"
	"store/src/configs"
	"store/src/middlewares"
	"store/src/routers"
	"store/src/validation"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func NewServer(cfg configs.Config) {
	engine := gin.New()
	engine.Use(gin.Recovery(), gin.Logger(), middlewares.CheckApiKey)

	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("id", validation.IdValidator, true)
		val.RegisterValidation("password", validation.PasswordValidator, true)
	}

	store := engine.Group("/store")
	{
		routers.Store(store)
	}

	engine.Run(fmt.Sprintf(":%v", cfg.Server.Port))
}