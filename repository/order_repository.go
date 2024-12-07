package repository

import (
	"errors"
	"gorm.io/gorm"
	"math"
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

func (repo OrderRepository) All(page, limit uint) (int, int, []domain.Order, error) {
	var count int64
	repo.db.Model(&domain.Order{}).Count(&count)
	pages := int(math.Ceil(float64(count) / float64(limit)))

	var orders []domain.Order
	result := repo.db.Preload("Customer").Scopes(helper.Paginate(page, limit)).Find(&orders)
	if result.RowsAffected == 0 {
		return 0, 0, nil, errors.New("order not found")
	}
	return int(count), pages, orders, nil
}

func (repo OrderRepository) Get(orderId uint) (domain.Order, error) {
	var order domain.Order
	result := repo.db.Preload("Customer").Preload("Items.Variant.Product").First(&order, orderId)
	return order, result.Error
}
