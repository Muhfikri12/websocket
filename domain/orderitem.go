package domain

type OrderItem struct {
	ID        uint  `gorm:"primaryKey;autoIncrement"`
	OrderID   uint  `gorm:"not null"`
	Order     Order `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	VariantID uint
	Quantity  uint
	UnitPrice float64 `gorm:"type:float"`
}
