package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"project/domain"
	"project/service"
)

type OrderController struct {
	service service.OrderService
	logger  *zap.Logger
}

func NewOrderController(service service.OrderService, logger *zap.Logger) *OrderController {
	return &OrderController{service: service, logger: logger}
}

// Order endpoint
// @Summary Customer orders
// @Description Get customer orders
// @Tags Order
// @Accept  json
// @Produce  json
// @Success 200 {object} handler.Response "orders retrieved"
// @Failure 404 {object} handler.Response "no data found"
// @Failure 500 {object} handler.Response "server error"
// @Router  /orders [post]
func (ctrl *OrderController) All(c *gin.Context) {
	orders, err := ctrl.service.All(1, 1)
	if err != nil {
		BadResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	GoodResponseWithData(c, "orders retrieved", http.StatusOK, orders)
}

// Order endpoint
// @Summary Customer order
// @Description Update customer order
// @Tags Order
// @Accept  json
// @Produce  json
// @Param domain.Order body domain.Order true " "
// @Success 200 {object} handler.Response "order updated"
// @Failure 422 {object} handler.Response "invalid input"
// @Failure 404 {object} handler.Response "no data found"
// @Failure 500 {object} handler.Response "server error"
// @Router  /orders [put]
func (ctrl *OrderController) Update(c *gin.Context) {
	var order domain.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		BadResponse(c, "invalid request body", http.StatusBadRequest)
		return
	}

	GoodResponseWithData(c, "orders updated", http.StatusOK, order)
}

// Order endpoint
// @Summary Customer order
// @Description Get customer order
// @Tags Order
// @Accept  json
// @Produce  json
// @Param domain.Order body domain.Order true " "
// @Success 200 {object} handler.Response "order retrived"
// @Failure 422 {object} handler.Response "invalid input"
// @Failure 404 {object} handler.Response "no data found"
// @Failure 500 {object} handler.Response "server error"
// @Router  /orders [get]
func (ctrl *OrderController) Get(c *gin.Context) {
	var order domain.Order

	GoodResponseWithData(c, "order retrieved", http.StatusOK, order)
}
