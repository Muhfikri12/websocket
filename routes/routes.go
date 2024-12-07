package routes

import (
	"log"
	"project/helper"
	"project/infra"
	"sync"

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
		category.POST("/", ctx.Ctl.Category.CreateCategory)
		category.DELETE("/:id", ctx.Ctl.Category.DeleteCategory)
		category.GET("/:id", ctx.Ctl.Category.GetCategoryByID)
		category.PUT("/:id", ctx.Ctl.Category.UpdateCategory)
	}

	products := r.Group("/products")
	{
		products.GET("/", ctx.Ctl.Product.ShowAllProduct)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/cdn-upload", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["images[]"]

		var wg sync.WaitGroup
		responses, _ := helper.Upload(&wg, files)
		log.Println(responses)
	})

	return r
}
