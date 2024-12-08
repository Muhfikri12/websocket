package productrepository

import (
	"fmt"
	"log"
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
	CreateProduct(product *domain.Product) error
	DeleteProduct(id int) error
	UpdateProduct(productID uint, product *domain.Product) error
}

type productRepo struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewProductRepo(db *gorm.DB, log *zap.Logger) ProductRepo {
	return &productRepo{db, log}
}

func (pr *productRepo) ShowAllProduct(page, limit int) (*[]domain.Product, int, int, error) {
	productList := []domain.Product{}

	var count int64
	if err := pr.db.Model(&domain.Product{}).Count(&count).Error; err != nil {
		pr.log.Error("Error from Show All Product: " + err.Error())
		return nil, 0, 0, err
	}

	result := pr.db.Scopes(helper.Paginate(uint(page), uint(limit))).
		Preload("ProductVariant").
		Preload("Image").
		Find(&productList)

	if result.Error != nil {
		pr.log.Error("Error: " + result.Error.Error())
		return nil, 0, 0, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, 0, 0, fmt.Errorf("no products found")
	}

	// Hitung total halaman
	totalPages := int(math.Ceil(float64(count) / float64(limit)))

	pr.log.Info("Success Get Data")
	return &productList, int(count), totalPages, nil
}

func (pr *productRepo) GetProductByID(id int) (*domain.Product, error) {

	product := domain.Product{}

	result := pr.db.Model(&product).Where("id = ?", id).
		Preload("ProductVariant").
		Preload("Image").First(&product)

	if result.Error != nil {
		return nil, fmt.Errorf("product not found")
	}

	return &product, nil
}

func (pr *productRepo) CreateProduct(product *domain.Product) error {
	var wg sync.WaitGroup
	var err error

	err = pr.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&product).Error; err != nil {
			return fmt.Errorf("failed to create product: %w", err)
		}

		if err := tx.Create(&domain.ProductVariant{
			ProductID: product.ID,
		}).Error; err != nil {
			return fmt.Errorf("failed to create product variant: %w", err)
		}

		if err := tx.Create(&domain.Image{
			ProductID: product.ID,
		}).Error; err != nil {
			log.Printf("failed to create image: %v", err)
			return fmt.Errorf("failed to create image: %w", err)
		}

		wg.Wait()

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Printf("Transaction failed: %v", err)
	}

	return err
}

func (pr *productRepo) UpdateProduct(productID uint, product *domain.Product) error {

	result := pr.db.Model(&product).
		Where("id = ?", productID).Updates(product)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no record found with shipping_id %d", productID)
	}

	return nil
}

func (pr *productRepo) DeleteProduct(id int) error {

	result := pr.db.Delete(&domain.Product{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("product not found")
	}

	return nil
}
