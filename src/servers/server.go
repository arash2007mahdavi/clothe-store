package servers

import (
	"fmt"
	"store/src/configs"
	"store/src/docs"
	"store/src/routers"
	"store/src/validation"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewServer(cfg configs.Config) {
	engine := gin.New()
	engine.Use(gin.Recovery(), gin.Logger())

	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("id", validation.IdValidator, true)
		val.RegisterValidation("password", validation.PasswordValidator, true)
	}

	store := engine.Group("/store")
	{
		routers.Store(store)
	}

	docs.SwaggerInfo.Title = "clothe store"
	docs.SwaggerInfo.Description = "clothe store"
	docs.SwaggerInfo.Version = "0.4"
	docs.SwaggerInfo.Schemes = []string{"http"}
	docs.SwaggerInfo.BasePath = "/store"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%v", cfg.Server.Port)
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	engine.Run(fmt.Sprintf(":%v", cfg.Server.Port))
}
