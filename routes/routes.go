package routes

import (
	"project/infra"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRoutes(ctx infra.ServiceContext) *gin.Engine {
	r := gin.Default()

	// endpoint login
	r.POST("/login", ctx.Ctl.AuthHandler.Login)
	category := r.Group("/category")
	{
		category.GET("/", ctx.Ctl.Category.ShowAllCategory)
		category.GET("/:id", ctx.Ctl.Category.DeleteCategory)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
