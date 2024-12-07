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
	ID             uint `gorm:"primaryKey;autoIncrement"`
	CustomerID     uint
	Customer       Customer
	PaymentMethod  string
	TrackingNumber string
	Status         string      `gorm:"type:orderstatus"`
	Items          []OrderItem `gorm:"foreignKey:OrderID"`
	CreatedAt      time.Time   `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time   `gorm:"autoUpdateTime" json:"updated_at"`
}
