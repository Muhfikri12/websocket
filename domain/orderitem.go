package domain

type OrderItem struct {
	ID        uint `gorm:"primaryKey"`
	OrderID   uint
	VariantID uint
	Order     Order
}
