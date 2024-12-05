package repository

import (
	"gorm.io/gorm"
	"project/database"
)

type Repository struct {
	Auth AuthRepository
}

func NewRepository(db *gorm.DB, cacher database.Cacher) Repository {
	return Repository{
		Auth: *NewAuthRepository(db, cacher),
	}
}
