package domain

import "time"

type Review struct {
	ID          uint `gorm:"primaryKey"`
	OrderItemID uint
	OrderItem   OrderItem
	Rating      float32
	Comment     string
	CreatedAt   time.Time
}
