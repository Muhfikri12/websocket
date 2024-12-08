package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"project/domain"
	"project/handler"
	productrepository "project/repository/product_repository"
	"project/service"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

func base() (handler.ProductHandler, *productrepository.ProductRepoMock) {

	log := *zap.NewNop()

	mockService := &productrepository.ProductRepoMock{}

	service := service.Service{
		Product: mockService,
	}

	return handler.NewProductHandler(&service, &log), mockService
}

func TestShowAllProduct(t *testing.T) {

	t.Run("Successfully retrieve all products", func(t *testing.T) {
		handler, mockService := base()
		r := gin.Default()
		r.GET("/products", handler.ShowAllProduct)

		mockProducts := []domain.Product{
			{
				ID:   1,
				Name: "Product 1",
			},
			{
				ID:   2,
				Name: "Product 2",
			},
		}

		mockService.On("ShowAllProduct", 1, 10).Return(&mockProducts, 2, 1, nil)

		req := httptest.NewRequest(http.MethodGet, "/products?page=1&limit=10", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockService.AssertCalled(t, "ShowAllProduct", 1, 10)

		var actualResponse map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &actualResponse)
		assert.NoError(t, err)
		assert.Equal(t, "Successfully Retrieved Products", actualResponse["message"])
		assert.True(t, actualResponse["status"].(bool))
	})

	t.Run("Fail to retrieve all products due to service error", func(t *testing.T) {

		handler, mockService := base()

		r := gin.Default()
		r.GET("/products", handler.ShowAllProduct)

		mockService.On("ShowAllProduct", 1, 10).Return(nil, 0, 0, fmt.Errorf("failed to retrieve products"))

		req := httptest.NewRequest(http.MethodGet, "/products?page=1&limit=10", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
		mockService.AssertCalled(t, "ShowAllProduct", 1, 10)

		expectedResponse := `{"message":"Product Not Found", "status":false}`
		assert.JSONEq(t, expectedResponse, w.Body.String())
	})
}

func TestGetProductByID(t *testing.T) {

	t.Run("Successfully retrieve product by ID", func(t *testing.T) {
		handler, mockService := base()

		r := gin.Default()
		r.GET("/products/:id", handler.GetProductByID)

		mockProduct := domain.Product{
			ID:   1,
			Name: "Product 1",
		}

		mockService.On("GetProductByID", 1).Return(&mockProduct, nil)

		req := httptest.NewRequest(http.MethodGet, "/products/1", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockService.AssertCalled(t, "GetProductByID", 1)

		var actualResponse map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &actualResponse)
		assert.NoError(t, err)
		assert.Equal(t, "Successfully Retrieved Product", actualResponse["message"])
		assert.True(t, actualResponse["status"].(bool))
	})

	t.Run("Fail to retrieve product by ID due to service error", func(t *testing.T) {
		handler, mockService := base()

		r := gin.Default()
		r.GET("/products/:id", handler.GetProductByID)

		mockService.On("GetProductByID", 1).Return(nil, fmt.Errorf("failed to retrieve product"))

		req := httptest.NewRequest(http.MethodGet, "/products/1", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
		mockService.AssertCalled(t, "GetProductByID", 1)

		expectedResponse := `{"message":"Product Not Found", "status":false}`
		assert.JSONEq(t, expectedResponse, w.Body.String())
	})
}

func TestCreateProductWithImage(t *testing.T) {

	t.Run("Successfully create product with image", func(t *testing.T) {
		handler, mockService := base()

		r := gin.Default()
		r.POST("/products", handler.CreateProduct)

		mockService.On("CreateProduct", mock.AnythingOfType("*domain.Product")).Return(nil)

		// Create a test file in memory
		testFile := []byte("dummy image content")
		fileName := "testimage.jpg"

		formData := new(bytes.Buffer)
		writer := multipart.NewWriter(formData)
		writer.WriteField("name", "Product 1")
		writer.WriteField("sku_product", "SKU001")
		writer.WriteField("price", "100")
		writer.WriteField("description", "Product 1 Description")
		writer.WriteField("variants", `[{"size":"L","color":"Red"}]`)

		part, err := writer.CreateFormFile("images", fileName)
		assert.NoError(t, err)
		_, err = part.Write(testFile)
		assert.NoError(t, err)

		writer.Close()

		req := httptest.NewRequest(http.MethodPost, "/products", formData)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		mockService.AssertCalled(t, "CreateProduct", mock.AnythingOfType("*domain.Product"))

		var actualResponse map[string]interface{}
		err = json.Unmarshal(w.Body.Bytes(), &actualResponse)
		assert.NoError(t, err)
		assert.Equal(t, "Product created successfully", actualResponse["message"])
		assert.True(t, actualResponse["status"].(bool))
	})

	t.Run("Fail to create product with image due to invalid price", func(t *testing.T) {
		handler, mockService := base()

		r := gin.Default()
		r.POST("/products", handler.CreateProduct)

		mockService.On("CreateProduct", mock.AnythingOfType("*domain.Product")).Return(fmt.Errorf("failed to create product"))

		// Create a test file in memory
		testFile := []byte("dummy image content")
		fileName := "testimage.jpg"

		formData := new(bytes.Buffer)
		writer := multipart.NewWriter(formData)
		writer.WriteField("name", "Product 1")
		writer.WriteField("sku_product", "SKU001")
		writer.WriteField("price", "invalid_price")
		writer.WriteField("description", "Product 1 Description")
		writer.WriteField("variants", `[{"size":"L","color":"Red"}]`)

		// Simulate uploading an image
		part, err := writer.CreateFormFile("images", fileName)
		assert.NoError(t, err)
		_, err = part.Write(testFile)
		assert.NoError(t, err)

		writer.Close()

		req := httptest.NewRequest(http.MethodPost, "/products", formData)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		mockService.AssertNotCalled(t, "CreateProduct", mock.AnythingOfType("*domain.Product"))

		expectedResponse := `{"message":"Invalid price value", "status":false}`
		assert.JSONEq(t, expectedResponse, w.Body.String())
	})

	t.Run("Fail to create product with image due to service error", func(t *testing.T) {
		handler, mockService := base()

		r := gin.Default()
		r.POST("/products", handler.CreateProduct)

		mockService.On("CreateProduct", mock.AnythingOfType("*domain.Product")).Return(fmt.Errorf("failed to create product"))

		// Create a test file in memory
		testFile := []byte("dummy image content")
		fileName := "testimage.jpg"

		formData := new(bytes.Buffer)
		writer := multipart.NewWriter(formData)
		writer.WriteField("name", "Product 1")
		writer.WriteField("sku_product", "SKU001")
		writer.WriteField("price", "100")
		writer.WriteField("description", "Product 1 Description")
		writer.WriteField("variants", `[{"size":"L","color":"Red"}]`)

		// Simulate uploading an image
		part, err := writer.CreateFormFile("images", fileName)
		assert.NoError(t, err)
		_, err = part.Write(testFile)
		assert.NoError(t, err)

		writer.Close()

		req := httptest.NewRequest(http.MethodPost, "/products", formData)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		mockService.AssertCalled(t, "CreateProduct", mock.AnythingOfType("*domain.Product"))

		expectedResponse := `{"message":"Failed to create product: failed to create product", "status":false}`
		assert.JSONEq(t, expectedResponse, w.Body.String())
	})
}

func TestDeleteProduct(t *testing.T) {

	t.Run("Successfully delete product", func(t *testing.T) {
		handler, mockService := base()

		r := gin.Default()
		r.DELETE("/products/:id", handler.DeleteProduct)

		// Mock service response
		mockService.On("DeleteProduct", mock.AnythingOfType("int")).Return(nil)

		// Perform request to delete product with id 1
		req := httptest.NewRequest(http.MethodDelete, "/products/1", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		// Assert the Response
		assert.Equal(t, http.StatusOK, w.Code)
		mockService.AssertCalled(t, "DeleteProduct", 1)

		// Assert the JSON Response Body
		expectedResponse := `{"message":"Product Deleted successfully","status":true,"data":1}`
		assert.JSONEq(t, expectedResponse, w.Body.String())
	})

	t.Run("Fail to delete product due to service error", func(t *testing.T) {
		handler, mockService := base()

		r := gin.Default()
		r.DELETE("/products/:id", handler.DeleteProduct)

		// Mock service response
		mockService.On("DeleteProduct", mock.AnythingOfType("int")).Return(fmt.Errorf("failed to delete product"))

		// Perform request to delete product with id 1
		req := httptest.NewRequest(http.MethodDelete, "/products/1", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		// Assert the Response
		assert.Equal(t, http.StatusNotFound, w.Code)
		mockService.AssertCalled(t, "DeleteProduct", 1)

		// Assert the JSON Response Body
		expectedResponse := `{"message":"Failed to Delete product: failed to delete product", "status":false}`
		assert.JSONEq(t, expectedResponse, w.Body.String())
	})
}
