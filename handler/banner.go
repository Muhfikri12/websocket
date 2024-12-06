package handler

import (
	"net/http"
	"project/domain"
	"project/helper"
	"project/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ControllerBanner struct {
	service service.ServiceBanner
	logger  *zap.Logger
}

func NewControllerBanner(service service.ServiceBanner, logger *zap.Logger) *ControllerBanner {
	return &ControllerBanner{service: service, logger: logger}
}

// @Summary Get All Banner
// @Description Endpoint Fetch All Banner
// @Tags GetAllBanner
// @Accept  json
// @Produce  json
// @Success 200 {object} handler.Response{data=[]model.Banner} "Get All Success"
// @Failure 500 {object} handler.Response "server error"
// @Router  /banner [get]
func (ctrl *ControllerBanner) GetAll(c *gin.Context) {
	banners, err := ctrl.service.GetAll()
	if err != nil {
		BadResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	GoodResponseWithData(c, "Get Banners success", http.StatusOK, banners)
}
func (ctrl *ControllerBanner) GetById(c *gin.Context) {
	id, err := helper.Uint(c.Param("id"))
	if err != nil {
		BadResponse(c, "Bad Request (Params)", http.StatusBadRequest)
		return
	}
	banner, err := ctrl.service.GetById(id)
	if err != nil {
		BadResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	GoodResponseWithData(c, "Get Banner success", http.StatusOK, banner)
}
func (ctrl *ControllerBanner) Create(c *gin.Context) {
	err := c.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		BadResponse(c, "Unable to parse form", http.StatusInternalServerError)
		return
	}

	// Ambil file dari form
	file, handler, err := c.Request.FormFile("file")
	if err != nil {
		BadResponse(c, "Error retrieving the file", http.StatusInternalServerError)
		return
	}
	defer file.Close()
	// Menampilkan informasi file
	// fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	// fmt.Printf("File Size: %+v\n", handler.Size)
	// fmt.Printf("MIME Header: %+v\n", handler.Header)

	imageUrl, err := helper.UploadFileThirdPartyAPI(file, handler.Filename)
	if err != nil {
		BadResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}
	banner := domain.Banner{
		Title:     c.DefaultPostForm("title", ""),
		PathPage:  c.DefaultPostForm("pathPage", ""),
		StartDate: c.DefaultPostForm("startDate", ""),
		EndDate:   c.DefaultPostForm("endDate", ""),
		IsPublish: false,
		ImageUrl:  c.DefaultPostForm(imageUrl, ""),
	}
	err = ctrl.service.Create(&banner)
	if err != nil {
		BadResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}
	GoodResponseWithData(c, "This Banner was successfully added", http.StatusCreated, "banner")
}
func (ctrl *ControllerBanner) Edit(c *gin.Context) {
	id, err := helper.Uint(c.Param("id"))
	if err != nil {
		BadResponse(c, "Bad Request (Params)", http.StatusBadRequest)
		return
	}
	err = c.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		BadResponse(c, "Unable to parse form", http.StatusInternalServerError)
		return
	}

	// Ambil file dari form
	file, handler, err := c.Request.FormFile("file")
	if err != nil {
		BadResponse(c, "Error retrieving the file", http.StatusInternalServerError)
		return
	}
	defer file.Close()
	// Menampilkan informasi file
	// fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	// fmt.Printf("File Size: %+v\n", handler.Size)
	// fmt.Printf("MIME Header: %+v\n", handler.Header)

	imageUrl, err := helper.UploadFileThirdPartyAPI(file, handler.Filename)
	if err != nil {
		BadResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}
	banner := domain.Banner{
		ID:        id,
		Title:     c.DefaultPostForm("title", ""),
		PathPage:  c.DefaultPostForm("pathPage", ""),
		StartDate: c.DefaultPostForm("startDate", ""),
		EndDate:   c.DefaultPostForm("endDate", ""),
		IsPublish: false,
		ImageUrl:  c.DefaultPostForm(imageUrl, ""),
	}
	err = ctrl.service.Edit(&banner)
	if err != nil {
		BadResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}
	GoodResponseWithData(c, "This Banner was successfully updated", http.StatusCreated, banner)
}
func (ctrl *ControllerBanner) SetPublish(c *gin.Context) {
	id, err := helper.Uint(c.Param("id"))
	if err != nil {
		BadResponse(c, "Bad Request (Params)", http.StatusBadRequest)
		return
	}
	GoodResponseWithData(c, "This Banner was successfully published", http.StatusCreated, id)
}
func (ctrl *ControllerBanner) Delete(c *gin.Context) {
	id, err := helper.Uint(c.Param("id"))
	if err != nil {
		BadResponse(c, "Bad Request (Params)", http.StatusBadRequest)
		return
	}
	var banner domain.Banner
	banner.ID = id
	err = ctrl.service.Delete(&banner)
	if err != nil {
		BadResponse(c, "Bad Request (Params)", http.StatusBadRequest)
		return
	}
	GoodResponseWithData(c, "This Banner was successfully deleted", http.StatusCreated, banner)
}
