package handler

import (
	"project/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler struct {
	AuthHandler   AuthController
	HandlerBanner ControllerBanner
}

func NewHandler(service service.Service, logger *zap.Logger) *Handler {
	return &Handler{
		AuthHandler:   *NewAuthController(service.Auth, logger),
		HandlerBanner: *NewControllerBanner(service.Banner, logger),
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
