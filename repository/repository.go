package repository

import (
	"project/config"
	"project/database"
	categoryrepositpry "project/repository/category_repositpry"
	productrepository "project/repository/product_repository"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	Auth          AuthRepository
	Order OrderRepository
	PasswordReset PasswordResetRepository
	User          UserRepository
	Category      categoryrepositpry.CategoryRepo
	Product       productrepository.ProductRepo


}

func NewRepository(db *gorm.DB, cacher database.Cacher, config config.Config, log *zap.Logger) Repository {
	return Repository{
		Category:      categoryrepositpry.NewCategoryRepo(db, log),
		Product:       productrepository.NewProductRepo(db, log),
		Auth:          *NewAuthRepository(db, cacher, config.AppSecret),
		Order: *NewOrderRepository(db),
		PasswordReset: *NewPasswordResetRepository(db),
		User:          *NewUserRepository(db),
	}
}
