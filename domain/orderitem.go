package domain

type OrderItem struct {
	ID        uint `gorm:"primaryKey"`
	OrderID   uint
	VariantID uint
	Quantity  uint
	UnitPrice float32 `gorm:"type:money"`
	Order     Order
}
