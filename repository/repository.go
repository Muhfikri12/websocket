package repository

import (
	"gorm.io/gorm"
	"project/config"
	"project/database"
)

type Repository struct {
	Auth          AuthRepository
	PasswordReset PasswordResetRepository
	User          UserRepository
}

func NewRepository(db *gorm.DB, cacher database.Cacher, config config.Config) Repository {
	return Repository{
		Auth:          *NewAuthRepository(db, cacher, config.AppSecret),
		PasswordReset: *NewPasswordResetRepository(db),
		User:          *NewUserRepository(db),
	}
}
