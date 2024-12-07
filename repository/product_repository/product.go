package productrepository

import (
	"fmt"
	"project/domain"
	"project/helper"
	"sync"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ProductRepo interface {
	ShowAllProduct(page, limit int) (*[]domain.Product, error)
}

type productRepo struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewProductRepo(db *gorm.DB, log *zap.Logger) ProductRepo {
	return &productRepo{db, log}
}

func (pr *productRepo) ShowAllProduct(page, limit int) (*[]domain.Product, error) {
	var wg sync.WaitGroup
	productList := []domain.Product{}

	result := pr.db.Scopes(helper.Paginate(uint(page), uint(limit))).Find(&productList)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("no products found")
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

	wg.Wait()

	return &productList, nil
}
