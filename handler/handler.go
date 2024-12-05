package handler

import (
	"project/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler struct {
	AuthHandler          AuthController
	PasswordResetHandler PasswordResetController
	UserHandler          UserController
	Category             CategoryHandler
	AuthHandler          AuthController
	PasswordResetHandler PasswordResetController
	UserHandler          UserController
	Category             CategoryHandler
}

func NewHandler(service service.Service, logger *zap.Logger) *Handler {
	return &Handler{
		AuthHandler:          *NewAuthController(service.Auth, logger),
		PasswordResetHandler: *NewPasswordResetController(service.PasswordReset, logger),
		UserHandler:          *NewUserController(service.User, logger),
		Category:             NewCategoryHandler(logger, &service),
	}
}

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func BadResponse(c *gin.Context, message string, statusCode int) {
	c.JSON(statusCode, Response{
		Status:  false,
		Message: message,
	})
}

func GoodResponseWithData(c *gin.Context, message string, statusCode int, data interface{}) {
	c.JSON(statusCode, Response{
		Status:  true,
		Message: message,
		Data:    data,
	})
}
