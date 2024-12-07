package handler

import (
	"net/http"
	"project/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DashboardHandler interface {
	GetEerningProduct(c *gin.Context)
}

type dashboardHandler struct {
	service *service.Service
	log     *zap.Logger
}

func NewDashboardHandler(service *service.Service, log *zap.Logger) DashboardHandler {
	return &dashboardHandler{service, log}
}

func (dh *dashboardHandler) GetEerningProduct(c *gin.Context) {

	totalEarning, err := dh.service.Dashboard.GetEarningProduct()
	if err != nil {
		BadResponse(c, "There is no earning yet", http.StatusBadRequest)
	}

	data := make(map[string]interface{})

	data["total_earning"] = totalEarning

	GoodResponseWithData(c, "successfully retrieved earning", http.StatusOK, data)
}
