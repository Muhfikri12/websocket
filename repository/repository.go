package repository

import (
	"project/database"

	"gorm.io/gorm"
)

type Repository struct {
	Auth   AuthRepository
	Banner RepositoryBanner
}

func NewRepository(db *gorm.DB, cacher database.Cacher) Repository {
	return Repository{
		Auth:   *NewAuthRepository(db, cacher),
		Banner: *NewRepositoryBanner(db, cacher),
	}
}
