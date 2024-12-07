package service

import (
	"project/repository"
	categoryservice "project/service/category_service"

	"go.uber.org/zap"
)

type Service struct {
	Auth          AuthService
	Category      categoryservice.CategoryService
	Order         OrderService
	PasswordReset PasswordResetService
	User          UserService
}

func NewService(repo repository.Repository, log *zap.Logger) Service {
	return Service{
		Auth:          NewAuthService(repo.Auth),
		Category:      categoryservice.NewCategoryService(&repo, log),
		Order:         NewOrderService(repo.Order),
		PasswordReset: NewPasswordResetService(repo.PasswordReset),
		User:          NewUserService(repo.User),
	}
}
