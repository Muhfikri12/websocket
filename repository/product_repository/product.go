package productrepository

import (
	"project/domain"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ProductRepo interface {
}

type productRepo struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewProductRepo(db *gorm.DB, log *zap.Logger) ProductRepo {
	return &productRepo{db, log}
}

func (pr *productRepo) ShowAllProduct(page, limit int) (*[]domain.Product, error) {

	return nil, nil
}
