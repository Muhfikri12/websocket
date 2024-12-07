package repository

import (
	"errors"
	"gorm.io/gorm"
	"project/domain"
	"project/helper"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (repo OrderRepository) Update(order *domain.Order) error {
	return repo.db.Create(&order).Error
}

func (repo OrderRepository) All(page, limit uint) ([]domain.Order, error) {
	var orders []domain.Order
	result := repo.db.Scopes(helper.Paginate(page, limit)).Find(&orders)
	if result.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}
	return orders, nil
}

func (repo OrderRepository) Get(orderId uint) (domain.Order, error) {
	var order domain.Order
	result := repo.db.First(&order, orderId)
	return order, result.Error
}
