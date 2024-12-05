package service

import (
	"project/repository"
	categoryservice "project/service/category_service"

	"go.uber.org/zap"
)

type Service struct {
	Auth          AuthService
	PasswordReset PasswordResetService
	User          UserService
	Category      categoryservice.CategoryService
}

func NewService(repo repository.Repository, log *zap.Logger) Service {
	return Service{
		Auth:          NewAuthService(repo.Auth),
		PasswordReset: NewPasswordResetService(repo.PasswordReset),
		User:          NewUserService(repo.User),
		Category:      categoryservice.NewCategoryService(&repo, log),
	}
}
