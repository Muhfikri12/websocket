package domain

import "time"

type Status string

const (
	Created   Status = "created"
	Processed Status = "processed"
	Canceled  Status = "canceled"
	Completed Status = "completed"
)

type Order struct {
	ID            uint `gorm:"primaryKey"`
	CustomerID    uint
	Customer      Customer
	PaymentMethod string
	Status        string `gorm:"type:orderstatus"`
	Items         []OrderItem
	CreatedAt     time.Time `gorm:"default:now()"`
	UpdatedAt     time.Time
}
