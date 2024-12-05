package categoryrepositpry

import (
	"fmt"
	"project/domain"
	"project/helper"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CategoryRepo interface {
	CreateCategory(category *domain.Category) error
	ShowAllCategory(page, limit int) (*[]domain.Category, error)
}

type categoryRepo struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewCategoryRepo(db *gorm.DB, log *zap.Logger) CategoryRepo {
	return &categoryRepo{db, log}
}

func (cr *categoryRepo) ShowAllCategory(page, limit int) (*[]domain.Category, error) {

	category := []domain.Category{}

	result := cr.db.Scopes(helper.Paginate(uint(page), uint(limit))).Find(&category)

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("category not found")
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &category, nil
}

func (cr *categoryRepo) CreateCategory(category *domain.Category) error {

	if err := cr.db.Create(category); err != nil {
		return fmt.Errorf("failed to create category: " + err.Error.Error())
	}

	return nil
}
