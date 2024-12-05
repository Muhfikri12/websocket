package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"project/service"
)

type UserController struct {
	service service.UserService
	logger  *zap.Logger
}

func NewUserController(service service.UserService, logger *zap.Logger) *UserController {
	return &UserController{service: service, logger: logger}
}

// Check Email endpoint
// @Summary Check Email
// @Description email must be valid when users want to reset their passwords
// @Tags Auth
// @Accept  json
// @Produce  json
// @Success 200 {object} handler.Response "email is valid"
// @Failure 404 {object} handler.Response "user not found"
// @Router  /users [post]
func (ctrl *UserController) Get(c *gin.Context) {

}
