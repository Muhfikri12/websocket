package repository

import (
	"project/config"
	"project/database"
	categoryrepositpry "project/repository/category_repositpry"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	Auth          AuthRepository
	Category      categoryrepositpry.CategoryRepo
	Order OrderRepository
	PasswordReset PasswordResetRepository
	User          UserRepository


}

func NewRepository(db *gorm.DB, cacher database.Cacher, config config.Config, log *zap.Logger) Repository {
	return Repository{
		Auth:          *NewAuthRepository(db, cacher, config.AppSecret),
		Category:      categoryrepositpry.NewCategoryRepo(db, log),
		Order: *NewOrderRepository(db),
		PasswordReset: *NewPasswordResetRepository(db),
		User:          *NewUserRepository(db),
	}
}
