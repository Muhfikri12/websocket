package productservice

import (
	"project/domain"
	"project/repository"

	"go.uber.org/zap"
)

type ProductService interface {
	ShowAllProduct(page int) (*[]domain.Product, error)
}

type productService struct {
	repo *repository.Repository
	log  *zap.Logger
}

func NewProductService(repo *repository.Repository, log *zap.Logger) ProductService {
	return &productService{repo, log}
}

func (ps *productService) ShowAllProduct(page int) (*[]domain.Product, error) {
	limit := 20

	products, err := ps.repo.Product.ShowAllProduct(page, limit)
	if err != nil {
		return nil, err
	}

	return products, nil
}
