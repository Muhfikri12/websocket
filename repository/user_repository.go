package repository

import (
	"gorm.io/gorm"
	"project/domain"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo UserRepository) Create(user *domain.User) error {
	return repo.db.Create(&user).Error
}
