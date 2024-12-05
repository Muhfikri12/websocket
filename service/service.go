package service

import (
	"project/repository"
	categoryservice "project/service/category_service"

	"go.uber.org/zap"
)

type Service struct {
	Auth     AuthService
	Category categoryservice.CategoryService
}

func NewService(repo repository.Repository, log *zap.Logger) Service {
	return Service{
		Auth:     NewAuthService(repo.Auth),
		Category: categoryservice.NewCategoryService(&repo, log),
	}
}
