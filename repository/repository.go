package repository

import (
	"gorm.io/gorm"
	"project/database"
)

type Repository struct {
	Auth          AuthRepository
	PasswordReset PasswordResetRepository
	User          UserRepository
}

func NewRepository(db *gorm.DB, cacher database.Cacher) Repository {
	return Repository{
		Auth:          *NewAuthRepository(db, cacher),
		PasswordReset: *NewPasswordResetRepository(db),
		User:          *NewUserRepository(db),
	}
}
