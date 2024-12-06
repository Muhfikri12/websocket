package routes

import (
	"project/infra"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRoutes(ctx infra.ServiceContext) *gin.Engine {
	r := gin.Default()
	//r.Use(ctx.Middleware.Logger())

	r.POST("/login", ctx.Ctl.AuthHandler.Login)
	r.POST("/register", ctx.Ctl.UserHandler.Registration)
	r.GET("/users", ctx.Ctl.UserHandler.All)
	r.POST("/password-reset", ctx.Ctl.PasswordResetHandler.Create)
	category := r.Group("/category")
	{
		category.GET("/", ctx.Ctl.Category.ShowAllCategory)
		category.DELETE("/:id", ctx.Ctl.Category.DeleteCategory)
		category.GET("/:id", ctx.Ctl.Category.GetCategoryByID)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
