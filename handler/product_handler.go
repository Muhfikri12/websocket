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
