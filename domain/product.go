package domain

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          int             `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string          `gorm:"type:varchar(50);not null" json:"name" binding:"required,min=5"`
	SKUProduct  string          `gorm:"type:varchar(100);unique;not null" json:"sku_product" binding:"required"`
	Price       float64         `gorm:"not null" json:"price" binding:"required"`
	Description string          `gorm:"type:text;not null" json:"description" binding:"required"`
	CreatedAt   time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   *gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	Image          []*Image          `gorm:"foreignKey:ProductID" json:"image"`
	ProductVariant []*ProductVariant `gorm:"foreignKey:ProductID" json:"product_variant"`
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

type BestSeller struct {
	ProductID int
	TotalSold int
}

type Revenue struct {
	Month   string `json:"month"`
	Revenue int    `json:"revenue"`
}
