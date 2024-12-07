package handler

import (
	"project/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler struct {
	AuthHandler          AuthController
	Category             CategoryHandler
	OrderHandler         OrderController
	PasswordResetHandler PasswordResetController
	UserHandler          UserController
}

func NewHandler(service service.Service, logger *zap.Logger) *Handler {
	return &Handler{
		AuthHandler:          *NewAuthController(service.Auth, logger),
		Category:             NewCategoryHandler(logger, &service),
		OrderHandler:         *NewOrderController(service.Order, logger),
		PasswordResetHandler: *NewPasswordResetController(service.PasswordReset, logger),
		UserHandler:          *NewUserController(service.User, logger),
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
