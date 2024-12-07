package handler

import (
	"net/http"
	"project/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ProductHandler interface {
	ShowAllProduct(c *gin.Context)
}

type productHandler struct {
	service *service.Service
	log     *zap.Logger
}

func NewProductHandler(service *service.Service, log *zap.Logger) ProductHandler {
	return &productHandler{service, log}
}

func (ph *productHandler) ShowAllProduct(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))

	products, err := ph.service.Product.ShowAllProduct(page)
	if err != nil {
		BadResponse(c, "Product Not Found", http.StatusNotFound)
		return
	}

	GoodResponseWithPage(c, "Successfully Retrieved Products", http.StatusOK, 10, page, 20, products)
}
