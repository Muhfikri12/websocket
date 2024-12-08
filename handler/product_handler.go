package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"project/domain"
	"project/helper"
	"project/service"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ProductHandler interface {
	ShowAllProduct(c *gin.Context)
	GetProductByID(c *gin.Context)
	CreateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
	UpdateProduct(c *gin.Context)
}

type productHandler struct {
	service *service.Service
	log     *zap.Logger
}

func NewProductHandler(service *service.Service, log *zap.Logger) ProductHandler {
	return &productHandler{service, log}
}

// Show All Products endpoint
// @Summary Show all Products
// @Description Get All Products
// @Tags All Products
// @Accept  json
// @Produce  json
// @Success 200 {object} handler.Response "Successfully Retrieved Products"
// @Failure 404 {object} handler.Response "Product Not Found
// @Failure 500 {object} handler.Response "server error"
// @Router  /products [get]
func (ph *productHandler) ShowAllProduct(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	if page <= 0 {
		page = 1
	}

	limit, _ := strconv.Atoi(c.Query("limit"))
	if limit < 10 {
		limit = 10
	}

	products, count, totalPages, err := ph.service.Product.ShowAllProduct(page, limit)
	if err != nil {
		BadResponse(c, "Product Not Found", http.StatusNotFound)
		return
	}

	GoodResponseWithPage(c, "Successfully Retrieved Products", http.StatusOK, count, totalPages, page, limit, products)
}

// Get Product By ID
// @Summary Get Product By ID
// @Description Get Product By ID
// @Tags Get Product By ID
// @Accept  json
// @Produce  json
// @Success 200 {object} handler.Response "Successfully Retrieved Product"
// @Failure 404 {object} handler.Response "Product Not Found"
// @Router  /products/:id [get]
func (ph *productHandler) GetProductByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	product, err := ph.service.Product.GetProductByID(id)
	if err != nil {
		BadResponse(c, "Product Not Found", http.StatusNotFound)
		return
	}

	GoodResponseWithData(c, "Successfully Retrieved Product", http.StatusOK, product)
}

// Create Product
// @Summary Create Product
// @Description Create Product
// @Tags Create Product
// @Accept  json
// @Produce  json
// @Success 201 {object} handler.Response "Product created successfully"
// @Failure 404 {object} handler.Response "Invalid form data"
// @Failure 500 {object} handler.Response "Failed to create product"
// @Router  /products [post]
func (ph *productHandler) CreateProduct(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		log.Println("Error reading form data:", err)
		BadResponse(c, "Invalid form data: "+err.Error(), http.StatusBadRequest)
		return
	}
	files := form.File["images"]
	for _, file := range files {
		log.Println("File size:", file.Size)
	}

	var wg sync.WaitGroup
	responses, err := helper.Upload(&wg, files)
	if err != nil {
		BadResponse(c, "Failed to upload images: "+err.Error(), http.StatusInternalServerError)
		return
	}

	name := c.PostForm("name")
	skuProduct := c.PostForm("sku_product")
	price, err := strconv.Atoi(c.PostForm("price"))
	if err != nil {
		BadResponse(c, "Invalid price value", http.StatusBadRequest)
		return
	}

	var images []*domain.Image
	for _, response := range responses {
		images = append(images, &domain.Image{
			URLPath: response.Data.Url,
		})
	}

	variantData := c.PostForm("variants")

	var productVariants []*domain.ProductVariant
	err = json.Unmarshal([]byte(variantData), &productVariants)
	if err != nil {
		BadResponse(c, "Invalid variant data: "+err.Error(), http.StatusBadRequest)
		return
	}

	product := domain.Product{
		Name:           name,
		SKUProduct:     skuProduct,
		Price:          float64(price),
		Description:    c.PostForm("description"),
		Image:          images,
		ProductVariant: productVariants,
	}

	if err := ph.service.Product.CreateProduct(&product); err != nil {
		BadResponse(c, "Failed to create product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	GoodResponseWithData(c, "Product created successfully", http.StatusCreated, product)
}

// Delete Product
// @Summary Delete Product
// @Description Delete Product
// @Tags Delete Product
// @Accept  json
// @Produce  json
// @Success 200 {object} handler.Response "Product Deleted successfully"
// @Failure 404 {object} handler.Response "Failed to Delete product"
// @Router  /products/:id [delete]
func (ph *productHandler) DeleteProduct(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	if err := ph.service.Product.DeleteProduct(id); err != nil {
		BadResponse(c, "Failed to Delete product: "+err.Error(), http.StatusNotFound)
		return
	}

	GoodResponseWithData(c, "Product Deleted successfully", http.StatusOK, id)
}

// Update Product
// @Summary Update Product
// @Description Update Product
// @Tags Update Product
// @Accept  json
// @Produce  json
// @Success 200 {object} handler.Response "Product Updated successfully"
// @Failure 404 {object} handler.Response "Failed to Update product"
// @Failure 500 {object} handler.Response "Invalid Payload Request"
// @Router  /products/:id [put]
func (ph *productHandler) UpdateProduct(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	product := domain.Product{}

	if err := c.ShouldBindJSON(&product); err != nil {
		BadResponse(c, "Failed to Update product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := ph.service.Product.UpdateProduct(uint(id), &product); err != nil {
		BadResponse(c, "Failed to Update product: "+err.Error(), http.StatusBadRequest)
		return
	}

	GoodResponseWithData(c, "Product Updated successfully", http.StatusOK, product)

}
