package repository

import (
	"project/config"
	"project/database"
	categoryrepositpry "project/repository/category_repositpry"

	"gorm.io/gorm"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	Auth          AuthRepository
	PasswordReset PasswordResetRepository
	User          UserRepository
	Category      categoryrepositpry.CategoryRepo
}

func NewRepository(db *gorm.DB, cacher database.Cacher, config config.Config, log *zap.Logger) Repository {
	return Repository{
		Auth:          *NewAuthRepository(db, cacher, config.AppSecret),
		PasswordReset: *NewPasswordResetRepository(db),
		User:          *NewUserRepository(db),
		Category:      categoryrepositpry.NewCategoryRepo(db, log),
	}
}
