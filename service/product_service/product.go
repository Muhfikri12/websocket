package productservice

import (
	"project/domain"
	"project/repository"

	"go.uber.org/zap"
)

type ProductService interface {
	ShowAllProduct(page, limit int) (*[]domain.Product, int, int, error)
	GetProductByID(id int) (*domain.Product, error)
}

type productService struct {
	repo *repository.Repository
	log  *zap.Logger
}

func NewProductService(repo *repository.Repository, log *zap.Logger) ProductService {
	return &productService{repo, log}
}

func (ps *productService) ShowAllProduct(page, limit int) (*[]domain.Product, int, int, error) {

	products, count, totalPages, err := ps.repo.Product.ShowAllProduct(page, limit)
	if err != nil {
		return nil, 0, 0, err
	}

	return products, count, totalPages, nil
}

func (ps *productService) GetProductByID(id int) (*domain.Product, error) {

	product, err := ps.repo.Product.GetProductByID(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}