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
	ID             uint        `gorm:"primaryKey;autoIncrement" json:"id"`
	CustomerID     uint        `json:"-"`
	Customer       Customer    `json:"customer"`
	PaymentMethod  string      `json:"payment_method"`
	TrackingNumber string      `json:"tracking_number"`
	Status         string      `gorm:"type:orderstatus" json:"status"`
	Items          []OrderItem `gorm:"foreignKey:OrderID" json:"items"`
	CreatedAt      time.Time   `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time   `gorm:"autoUpdateTime" json:"updated_at"`
}
