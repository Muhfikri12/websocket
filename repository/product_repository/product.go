package productrepository

import (
	"fmt"
	"math"
	"project/domain"
	"project/helper"
	"sync"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ProductRepo interface {
	ShowAllProduct(page, limit int) (*[]domain.Product, int, int, error)
	GetProductByID(id int) (*domain.Product, error)
}

type productRepo struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewProductRepo(db *gorm.DB, log *zap.Logger) ProductRepo {
	return &productRepo{db, log}
}

func (pr *productRepo) ShowAllProduct(page, limit int) (*[]domain.Product, int, int, error) {
	var wg sync.WaitGroup
	productList := []domain.Product{}

	var count int64
	if err := pr.db.Model(&domain.Product{}).Count(&count).Error; err != nil {
		return nil, 0, 0, err
	}

	result := pr.db.Scopes(helper.Paginate(uint(page), uint(limit))).Find(&productList)

	if result.Error != nil {
		return nil, 0, 0, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, 0, 0, fmt.Errorf("no products found")
	}

	for i := range productList {
		wg.Add(1)
		go func(product *domain.Product) {
			defer wg.Done()
			var (
				variants []domain.ProductVariant
				images   []domain.Image
			)

			pr.db.Model(&domain.ProductVariant{}).Where("product_id = ?", product.ID).Find(&variants)
			product.ProductVariant = &variants

			pr.db.Model(&domain.Image{}).Where("product_id = ?", product.ID).Find(&images)
			product.Image = &images

		}(&productList[i])
	}

	totalPages := int(math.Ceil(float64(count) / float64(limit)))

	wg.Wait()

	return &productList, int(count), totalPages, nil
}

func (pr *productRepo) GetProductByID(id int) (*domain.Product, error) {

	product := domain.Product{}

	result := pr.db.Model(&product).Where("id = ?", id).First(&product)
	if result.Error != nil {
		return nil, fmt.Errorf("product not found")
	}

	variantChan := make(chan []domain.ProductVariant)
	imageChan := make(chan []domain.Image)
	errChan := make(chan error)

	go func() {
		var variants []domain.ProductVariant
		if err := pr.db.Model(&domain.ProductVariant{}).Where("product_id = ?", product.ID).Find(&variants).Error; err != nil {
			errChan <- err
			return
		}
		variantChan <- variants
	}()

	go func() {
		var images []domain.Image
		if err := pr.db.Model(&domain.Image{}).Where("product_id = ?", product.ID).Find(&images).Error; err != nil {
			errChan <- err
			return
		}
		imageChan <- images
	}()

	var variants []domain.ProductVariant
	var images []domain.Image
	for i := 0; i < 2; i++ {
		select {
		case v := <-variantChan:
			variants = v
		case img := <-imageChan:
			images = img
		case err := <-errChan:
			return nil, err
		}
	}

	product.ProductVariant = &variants
	product.Image = &images

	return &product, nil
}
