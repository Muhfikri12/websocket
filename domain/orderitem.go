package domain

type OrderItem struct {
	ID        uint `gorm:"primaryKey"`
	OrderID   uint
	VariantID uint
	Quantity  uint
	UnitPrice float64 `gorm:"type:float"`
	Order     Order
}
