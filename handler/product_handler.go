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
	GetProductByID(c *gin.Context)
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
	if page <= 0 {
		page = 1
	}

	limit, _ := strconv.Atoi(c.Query("limit"))
	if limit < 10 {
		limit = 2
	}

	products, count, totalPages, err := ph.service.Product.ShowAllProduct(page, limit)
	if err != nil {
		BadResponse(c, "Product Not Found", http.StatusNotFound)
		return
	}

	GoodResponseWithPage(c, "Successfully Retrieved Products", http.StatusOK, count, totalPages, page, limit, products)
}

func (ph *productHandler) GetProductByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	product, err := ph.service.Product.GetProductByID(id)
	if err != nil {
		BadResponse(c, "Product Not Found", http.StatusNotFound)
	}

	GoodResponseWithData(c, "Successfully Retrieved Product", http.StatusOK, product)
}
