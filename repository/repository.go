package repository

import (
	"project/database"
	categoryrepositpry "project/repository/category_repositpry"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	Auth     AuthRepository
	Category categoryrepositpry.CategoryRepo
}

func NewRepository(db *gorm.DB, cacher database.Cacher, log *zap.Logger) Repository {
	return Repository{
		Auth:     *NewAuthRepository(db, cacher),
		Category: categoryrepositpry.NewCategoryRepo(db, log),
	}
}
