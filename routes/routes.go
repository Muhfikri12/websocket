package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"project/helper"
	"project/infra"
	"sync"
)

func NewRoutes(ctx infra.ServiceContext) *gin.Engine {
	r := gin.Default()
	//r.Use(ctx.Middleware.Logger())

	r.POST("/login", ctx.Ctl.AuthHandler.Login)
	r.POST("/register", ctx.Ctl.UserHandler.Registration)
	r.GET("/users", ctx.Ctl.UserHandler.All)
	r.POST("/password-reset", ctx.Ctl.PasswordResetHandler.Create)

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
