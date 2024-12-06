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

	bannerGroup := r.Group("/banner")
	{
		bannerGroup.GET("/", ctx.Ctl.HandlerBanner.GetAll)
		bannerGroup.POST("/", ctx.Ctl.HandlerBanner.Create)
		bannerGroup.PUT("/:id", ctx.Ctl.HandlerBanner.Edit)
		bannerGroup.GET("/:id", ctx.Ctl.HandlerBanner.GetById)

		// Mungkin delete bisa di grouping dengan delete yang lain karena hanya admin yang bisa akses
		bannerGroup.DELETE("/:id", ctx.Ctl.HandlerBanner.Delete)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
