package handler

import (
	"net/http"
	"project/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CategoryHandler interface {
	ShowAllCategory(c *gin.Context)
	DeleteCategory(c *gin.Context)
	GetCategoryByID(c *gin.Context)
}

type categoryHandler struct {
	log     *zap.Logger
	service *service.Service
}

func NewCategoryHandler(log *zap.Logger, service *service.Service) CategoryHandler {
	return &categoryHandler{log, service}
}

func (ch *categoryHandler) ShowAllCategory(c *gin.Context) {
	pageStr := c.Query("page")

	page, _ := strconv.Atoi(pageStr)

	categories, err := ch.service.Category.ShowAllCategory(page)
	if err != nil {
		BadResponse(c, "Failed to retrived categories: "+err.Error(), http.StatusNotFound)
		return
	}

	GoodResponseWithData(c, "successfully retrived categories", http.StatusOK, categories)
}

func (ch *categoryHandler) DeleteCategory(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	if err := ch.service.Category.DeleteCategory(id); err != nil {
		BadResponse(c, "Failed to deleted categoriy: "+err.Error(), http.StatusNotFound)
		return
	}

	GoodResponseWithData(c, "successfully deleted category", http.StatusOK, id)

}

func (ch *categoryHandler) GetCategoryByID(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	category, err := ch.service.Category.GetCategoryByID(id)
	if err != nil {
		BadResponse(c, "Failed to Retrieved categoriy: "+err.Error(), http.StatusNotFound)
		return
	}

	GoodResponseWithData(c, "successfully Retrieved category", http.StatusOK, category)
}
