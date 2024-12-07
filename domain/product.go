package domain

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          int             `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string          `gorm:"type:varchar(50)" json:"name"`
	SKUProduct  string          `gorm:"type:varchar(100);unique" json:"sku_product"`
	Price       float64         `gorm:"type:float" json:"price"`
	Description string          `gorm:"type:text" json:"description"`
	CreatedAt   time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   *gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	Image          *[]Image          `json:"image"`
	ProductVariant *[]ProductVariant `json:"product_variant"`
}

func SeedProducts() []Product {
	products := []Product{
		{
			Name:        "Product A",
			SKUProduct:  "SKU-001",
			Price:       100000,
			Description: "This is Product A description.",
		},
		{
			Name:        "Product B",
			SKUProduct:  "SKU-002",
			Price:       150000,
			Description: "This is Product B description.",
		},
		{
			Name:        "Product C",
			SKUProduct:  "SKU-003",
			Price:       200000,
			Description: "This is Product C description.",
		},
	}

	return products
}