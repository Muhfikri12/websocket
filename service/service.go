package service

import (
	"project/repository"
	categoryservice "project/service/category_service"
	productservice "project/service/product_service"

	"go.uber.org/zap"
)

type Service struct {
	Auth          AuthService
	PasswordReset PasswordResetService
	User          UserService
	Category      categoryservice.CategoryService
	Product       productservice.ProductService
}

func NewService(repo repository.Repository, log *zap.Logger) Service {
	return Service{
		Auth:          NewAuthService(repo.Auth),
		PasswordReset: NewPasswordResetService(repo.PasswordReset),
		User:          NewUserService(repo.User),
		Category:      categoryservice.NewCategoryService(&repo, log),
		Product:       productservice.NewProductService(&repo, log),
	}
}
