package domain

import "time"

type Order struct {
	ID         uint `gorm:"primaryKey"`
	CustomerID uint
	Customer   Customer
	Items      []OrderItem
	CreatedAt  time.Time `gorm:"default:now()"`
	UpdatedAt  time.Time
}
