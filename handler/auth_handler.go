package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"project/domain"
	"project/service"
)

type AuthController struct {
	service service.AuthService
	logger  *zap.Logger
}

func NewAuthController(service service.AuthService, logger *zap.Logger) *AuthController {
	return &AuthController{service: service, logger: logger}
}

// Login endpoint
// @Summary Admin login
// @Description authenticate user
// @Tags Authentication
// @Accept  json
// @Produce  json
// @Param domain.User body domain.User true " "
// @Success 200 {object} handler.Response "login successfully"
// @Failure 400 {object} handler.Response
// @Failure 401 {object} handler.Response "authentication failed"
// @Failure 500 {object} handler.Response "server error"
// @Router  /login [post]
func (ctrl *AuthController) Login(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		BadResponse(c, "invalid request body", http.StatusBadRequest)
		return
	}

	idKey, token, isAuthenticated, err := ctrl.service.Login(user)
	if err != nil {
		BadResponse(c, err.Error(), http.StatusUnauthorized)
		return
	}

	if !isAuthenticated {
		BadResponse(c, "authentication failed", http.StatusUnauthorized)
		return
	}

	GoodResponseWithData(c, "user authenticated", http.StatusOK, gin.H{"token": token, "id_key": idKey})
}
